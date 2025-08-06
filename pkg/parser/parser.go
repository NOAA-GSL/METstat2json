package parser

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dtcenter/METstat2json/pkg/linetypes/v10_0"
	"github.com/dtcenter/METstat2json/pkg/linetypes/v10_1"
	"github.com/dtcenter/METstat2json/pkg/linetypes/v11_0"
	"github.com/dtcenter/METstat2json/pkg/linetypes/v11_1"
	"github.com/dtcenter/METstat2json/pkg/linetypes/v12_0"
	"github.com/dtcenter/METstat2json/pkg/util"
)

/*
Supportoing versions FYI, these are all the >= 10.0.0 releases of MET that exist: 10.0 10.1 11.0 11.1 12.0
This package is used to parse the data for MET output files.
The entry point to this package is the ParseLine function that takes a data set name, header line, a data line, a fileType,
and a map of documents indexed by the document id. The document pointer can be an empty document. The data set name is a string that identifies the actual MET dataset.
For example a MET user may run the same thing multiple times and without a unique data set name for
each run the id fields of the JSON documents in the parsed output would be the same and the data would
overwrite itself in the database. So the data set name is required and must be unique. The header line
is the first line of the file and contains the header field names.  The data line is any subsequent line of the file
that contains the header and the data fields.
The fileType is a string that represents the type of file being parsed. The docPtr is a pointer to a map of
documents that are indexed by an id that is derived from the header fields minus the dataKey fields.

A dataKey is an array of header field values. For example most line types have a dataKey of {"Fcst_lead"}
which would have a string representation of the value of the "Fcst_lead" element in the header string.
The dataKey fields are disallowed from the header id and are not included in the headerData. These keys serve
the purpose of actually merging line data with the same dataKey values into a single document. The dataKey is used
to index the data section of the document, which is a map[string]interface{}, where the interface is a specific concrete
data type.

The ParseLine function uses the GetLineType function to determine the lineType of the data line, the headerData,
dataKey, and descIndex. headerData are the ordered data fields for the header section of the line, the dataKey is the actual
dataKey i.e the concatenated dataKey values, and the descIndex is the ordinal index of the desc field.
The descIndex is used to trim the desc field to 10 characters.

The parseLine function also uses the
GetId function to determine the id of the data line. The id is derived from the headerData minus the dataKey fields and
is returned in the form of a VxMetadata struct. The VxMetadata struct is then converted to a map[string]interface{}
so that it can be passed to the NewDocForId function without the NewDocForId function needing to know the VxMetadata struct type.

There are a couple of utility functions that are used to get the headerData without the NA values and to convert the VxMetadata struct.
A document pointer is required as a place to store the parsed data. If the document is nil, a new document is created.
The header line values (minus the dataKey fields) are used to derive the id, with date fields converted to epochs.
If the data section of of the document[id] is nil, a new data section is created. The data section is then populated
with the data fields from the data line. If the data section is not nil, the data fields are added to the existing data map.

The parameter getExternalDocForId is a function pointer that is used to get an external document for a given id. This function
is used to get a document from an external source, such as a database, that is indexed by the id. If the external document
is not nil, it is added to the document map. If the external document is nil, a new document is created for the id.
*/

const DOC_NOT_FOUND = "document not found"

// Gets the MET version from the stat file data line so that we can determine which linetypes module to use.
func getParserVersion(dataLine string) (string, error) {
	metVersion := strings.ToLower(strings.Fields(dataLine)[0])
	metVersionParts := strings.Split(metVersion, ".")
	if len(metVersionParts) != 3 {
		return "", fmt.Errorf("invalid MET version format: %s", metVersion)
	}
	lineVersion := metVersionParts[0] + "_" + metVersionParts[1]
	return lineVersion, nil
}

// Filter out VIM *.swp and MacOS DS_STORE files.
func isValidFileType(filename string) (bool, string) {
	filename = filepath.Base(filename)
	filePathParts := strings.Split(filename, ".")
	fileType := filePathParts[1]
	switch strings.ToUpper(fileType) {
	case "SWP":
		return false, fileType
	case "DS_STORE":
		return false, fileType
	}
	return true, fileType
}

// Main entrypoint to the library. Ideally, this should be the only thing consumers need to call.
// Parses the headerLine & dataLine passed to it and adds it to the collection of JSON docs pointed to by docPtr.
// (docPtr is a map[string]interface where key = docID & value = doc struct with header & data)
// Uses the fileName to deduce the doc type.
// dataSetName should be a <=10 char name which identifies the dataset.
func ParseLine(dataSetName string, headerLine string, dataLine string, docs *map[string]util.METdocument, fileName string, getExternalDocForId func(id string) (util.METdocument, error)) (map[string]util.METdocument, error) {
	// recover from unexpected errors
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %v for fileName %s\n", r, fileName)
		}
	}()

	if dataSetName == "" {
		return *docs, fmt.Errorf("dataSetName is empty")
	}
	if len(dataSetName) > 10 {
		return *docs, fmt.Errorf("dataSetName is too long - must be <= 10 characters")
	}
	// get line version e.g. V12.0.0 -> v12_0
	parserVersion, err := getParserVersion(dataLine)
	if err != nil {
		return *docs, fmt.Errorf("error getting parser version from line %s: %w", dataLine, err)
	}
	if headerLine == "" {
		return *docs, fmt.Errorf("empty header line")
	}
	if dataLine == "" {
		return *docs, fmt.Errorf("empty data line")
	}

	// Filter out undesired files.
	valid, filetype := isValidFileType(fileName)
	if !valid {
		// Skip undesired files
		return *docs, fmt.Errorf("skipping %s file", filetype)
	}

	// get the lineType
	fileLineType, headerData, dataData, dataKey, descIndex, err := util.GetLineType(headerLine, dataLine, fileName, parserVersion) // Doesn't access file - uses fileName to deduce file type.
	if err != nil {
		// cannot process this line so return the docPtr as is - it is probably a truncated line
		fmt.Println("Error getting line type: ", err)
		return *docs, err
	}
	// if there are any disallowed fields in this linetype then add the disallowed data to the dataData array - in order
	disallowedFields := util.DataKeyMap[fileLineType].HeaderDisallow
	if len(disallowedFields) > 0 {
		for _, disallowedField := range disallowedFields {
			disAllowedFieldValue, err := util.GetHeaderValue(strings.Fields(headerLine), strings.Fields(dataLine), disallowedField)
			// if there is an error getting the disallowed field, just append "" to the dataData array
			if err != nil {
				fmt.Println("Error getting disallowed field: ", err)
			}
			// if an err it appends ""
			dataData = append(dataData, disAllowedFieldValue)
		}
	}

	// get the tmpHeaderData without the NA values
	tmpHeaderData := getTmpHeaderSanNA(headerData, descIndex)
	if *docs == nil {
		newDoc := make(map[string]util.METdocument)
		docs = &newDoc
	}

	// metadata doesn't change between versions, we just use the latest one. Same with DOC
	docID, err := util.BuildId("MET", "DD", "MET", dataSetName, tmpHeaderData)
	if err != nil {
		return *docs, fmt.Errorf("error getting id from line %s: %w", dataLine, err)
	}

	metadata := util.VxMetadata{
		Subset:      "MET",
		Type:        "DD",
		SubType:     "MET",
		DataSetName: dataSetName,
		ID:          docID,
	}

	_, exists := (*docs)[metadata.ID]
	if !exists {
		// check to see if there is an existing external document for this id
		externalExistingDoc, err := (getExternalDocForId)(metadata.ID)
		if err != nil && !strings.HasPrefix(err.Error(), DOC_NOT_FOUND) {
			return *docs, err
		}
		// if there is an external document for this id, use it, we will add the data into it
		if externalExistingDoc != nil {
			(*docs)[metadata.ID] = externalExistingDoc
		} else {
			// have to create a new document for this id

			// create a new document for the new metaData.ID
			// This function will also fill in the headerData fields
			// indexed by dataKey value in the document.
			// The document needs to be of the correct version.
			switch parserVersion {
			case "v10_0":
				(*docs)[metadata.ID], err = v10_0.NewDocForId(fileLineType, metadata, headerData, dataData, dataKey)
			case "v10_1":
				(*docs)[metadata.ID], err = v10_1.NewDocForId(fileLineType, metadata, headerData, dataData, dataKey)
			case "v11_0":
				(*docs)[metadata.ID], err = v11_0.NewDocForId(fileLineType, metadata, headerData, dataData, dataKey)
			case "v11_1":
				(*docs)[metadata.ID], err = v11_1.NewDocForId(fileLineType, metadata, headerData, dataData, dataKey)
			case "v12_0":
				(*docs)[metadata.ID], err = v12_0.NewDocForId(fileLineType, metadata, headerData, dataData, dataKey)
			default:
				return *docs, fmt.Errorf("unsupported version %s", parserVersion)
			}
			if err != nil || (*docs)[metadata.ID] == nil {
				return *docs, fmt.Errorf("error creating doc for file: %s error: %w", fileName, err)
			}
			// return the new doc - the doc was created and the data was added to it
			return *docs, err
		}
	} else {
		// we either had the doc already, got it externally, or created it
		// now we need to add the data to the document
		doc := (*docs)[metadata.ID]
		if err := doc.AddDataElement(dataKey, dataData); err != nil {
			return *docs, fmt.Errorf("problem adding data to document %w", err)
		}
		return *docs, err
	}
	return *docs, err
}

/*
create a tmpHeaderData and remove the "" and the NA values fromm the headerData.
This also has to be done in the NewDocForId i.e. (fill_XXXX_Header) functions,
and trim the desc field data to 10 chars, if it isn't empty ("")
*/
func getTmpHeaderSanNA(headerData []string, descIndex int) []string {
	tmpHeaderData := []string{}
	for i, h := range headerData {
		if h != "NA" && h != "" {
			if i == descIndex {
				if len(h) > 10 {
					h = h[:10]
				}
			}
			tmpHeaderData = append(tmpHeaderData, h)
		}
	}
	// make the headerData NA values into "" so that the fillXXXX_Header functions make those values empty
	for i, h := range headerData {
		if h == "NA" {
			headerData[i] = ""
		}
	}
	return tmpHeaderData
}

func WriteJsonToCompressedFile(docs map[string]util.METdocument, filename string) error {
	// get the documents as a list
	// Defines the Slice capacity to match the Map elements count
	docList := make([]util.METdocument, 0, len(docs))

	for _, doc := range docs {
		docList = append(docList, doc)
	}
	// Marshal the document struct to JSON
	jsonBytes, err := json.Marshal(docList)
	if err != nil {
		return err
	}
	// Write the JSON bytes to a file
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err = w.Write(jsonBytes)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	// Write the compressed data to a file
	err = os.WriteFile(filename, b.Bytes(), 0o644)
	if err != nil {
		fmt.Printf("Failed to write file: %v", err)
	}
	return nil
}
