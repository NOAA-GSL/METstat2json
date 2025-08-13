package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"slices"
	"strings"
	"text/template"

	"github.com/dtcenter/METstat2json/pkg/util"
)

/*
NOTES:
See https://github.com/dtcenter/MET/blob/main_v12.0/data/table_files/met_header_columns_V12.0.txt#L24.
also see https://met.readthedocs.io/en/latest/Users_Guide/index.html and look at the data type definitions
for the various MET file types. The data type definitions are in the "Column Definitions" section of the
various MET file type documentation. For example see https://met.readthedocs.io/en/latest/Users_Guide/mode.html#mode-output
for the MODE file type output format.

Delineating header from data sections in the header records:
The field that delineates the header section from the data section in most line types is the LINE_TYPE field
The exceptions are MODE and MTD.
MODE object header records use the "OBJECT_ID" field and MODE cts headers use the "FIELD" field.
MTD headers use the "OBJECT_ID" field.

Recognizing data types:
The way that you can tell if a data file is a mode file is that
MODE object files begin with mode_ and end with _obj.txt and the CTS files begin with mode_
and end with _cts.txt. However, Grid-Stat and Point-Stat also write output files with the
_cts.txt suffix, but they begin with grid_stat_ and point_stat_ , respectively, and they have a LINE_TYPE field.

MTD files start with mtd_ and end with either _2d.txt, _3d_pair_cluster.txt, _3d_pair_simple.txt, _3d_single_cluster.txt, or _3d_single_simple.txt
Other files have a line type field that is used to determine the data types of the fields in the data section of the document.
There can be multiples of these line types in a single file.
*/
type Pattern struct {
	match       *regexp.Regexp
	dType       string
	structField string
	structType  string
}

func getMetHeaderColumnsFileURL(parserVersion string) (string, error) {
	switch parserVersion {
	case "v12_0":
		return util.MetHeaderColumnsFileUrl_v12_0, nil
	case "v11_1":
		return util.MetHeaderColumnsFileUrl_v11_1, nil
	case "v11_0":
		return util.MetHeaderColumnsFileUrl_v11_0, nil
	case "v10_1":
		return util.MetHeaderColumnsFileUrl_v10_1, nil
	case "v10_0":
		return util.MetHeaderColumnsFileUrl_v10_0, nil
	default:
		return "", fmt.Errorf("unsupported MET parserVersion: %s - supported are v12_0, v11_1, v11_0, v10_1, v10_0", parserVersion)
	}
}

/*
The output of this program is a series of structs that can be used to define the header
and data types in the buildHeaderTypes.go file and some parsing routines that are aware of the
header and data types.
There are several repeating fields that are handled in a special way. These fields are handled in the getRepeatingSequenceStructureString
function. See the notes in that function for more information.
*/

func main() {
	// Using the header definitions in the appropriate version of data/table_files/met_header_columns_X.X.txt
	// currently 12.0 to get the required header column definitions and then using
	// using the https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/docs/Users_Guide
	// to get the stat field types. We create a map of field names to field types. Then we have to re-iterate
	// over the header definitions to create the structs and functions to fill the structs.
	// read the header columns file

	// 1. Gather data from MET files
	var version string
	flag.StringVar(&version, "version", "", "Specify the parser version (e.g., -version=v12.0|v11.1|v11.0|v10.1|v10.0)")
	flag.Parse()
	parserVersion := strings.ReplaceAll(version, ".", "_")
	metHeaderColumnsFileUrl, err := getMetHeaderColumnsFileURL(parserVersion)
	if err != nil {
		fmt.Println("error setting MET version: ", err)
		os.Exit(1)
	}
	metLines, fieldNameMap := getColumnLinesAndMapForUrl(metHeaderColumnsFileUrl)
	metDataTypesForLines := fillMetDataMapFromSrcFiles(util.MetSrcFiles, make(map[string]string), fieldNameMap)
	metDataTypesForLines = fillMetDataMapFromUserGuide(util.MetUserDocFiles, metDataTypesForLines, fieldNameMap)

	// 2. Assemble code generation data
	allDocumentStructs := make(map[string]documentStructData)
	allHeaderStructs := make(map[string]headerStructData)
	allDataStructs := make(map[string]dataStructData)
	allHeaderFillFunctions := make(map[string]headerFillMethodData)
	allDataFillFunctions := make(map[string]dataFillMethodData)
	for _, line := range metLines {
		// get the prefix from the line
		fieldStr, fileType, lineType, err := getFileLineType(line)
		if err != nil {
			continue // skip this line - it didn't parse properly
		}
		fileLineType := fileType + "_" + lineType
		// split the line into header and data fields
		headerFields, dataFields := util.SplitColumnDefLine(fileLineType, fieldStr)

		// Create struct names
		structNames := createStructNames(fileType, lineType)

		// Build data for header & data structs
		headerStructFields := getHeaderFields(headerFields, fileType, lineType, metDataTypesForLines)
		dataStructFields := getDataFields(dataFields, fileType, lineType, metDataTypesForLines)

		// get data for Header & Data fill functions
		headerFillFields := getHeaderFillFields(headerFields, fileType, lineType)
		dataFillFields := getDataFillFields(dataFields, metDataTypesForLines, fileType, lineType)
		// data fill can have the weird repeating things

		// Add to data maps
		allDocumentStructs[structNames.DocumentStructName] = structNames
		allHeaderStructs[structNames.HeaderStructName] = headerStructData{
			DocumentStructName: structNames.DocumentStructName,
			HeaderStructName:   structNames.HeaderStructName,
			Fields:             headerStructFields,
		}
		allHeaderFillFunctions[structNames.HeaderStructName] = headerFillMethodData{
			DocumentStructName: structNames.DocumentStructName,
			HeaderStructName:   structNames.HeaderStructName,
			Fields:             headerFillFields,
		}
		allDataStructs[structNames.DataStructName] = dataStructData{
			DataStructName: structNames.DataStructName,
			Fields:         dataStructFields,
		}
		allDataFillFunctions[structNames.DataStructName] = dataFillMethodData{
			DataStructName: structNames.DataStructName,
			Fields:         dataFillFields,
		}
	}

	// 3. Generate code using templates

	// Generate the package declaration & import statements
	printPackageAndImports(parserVersion)

	// Generate the various code sections
	printDocumentStructs(allDocumentStructs)
	printAddDataElementFunctions(allDocumentStructs)
	printGetIDFunctions(allDocumentStructs)
	printHeaderStructs(allHeaderStructs)
	printHeaderFillFunctions(allHeaderFillFunctions)
	printDataStructs(allDataStructs)
	printDataFillFunctions(allDataFillFunctions)
	printNewDocForIDFunction(allDocumentStructs)
}

// --- Code generation helper functions ---

func printPackageAndImports(parserVersion string) {
	fmt.Printf(`package %s
import (
    "errors"
    "fmt"
    "strconv"

	"github.com/dtcenter/METstat2json/pkg/mettypes"
    "github.com/dtcenter/METstat2json/pkg/validtypes"
)
/*
THIS CODE IS AUTOMATICALLY GENERATED - DO NOT EDIT THIS CODE
To modify this code - modify the generator.go file and run the generator.go program
cd  <repo_root>
go run generator -version=v12.0 > pkg/linetypes/v12_0/linetypes.go
*/

// Helper function to reduce boilerplate for using errors.Join()
// appends the given error to the error slice, and adds the field name as context
func appendErrorWithContext(errs *[]error, fieldName string, err error) {
    if err != nil {
        *errs = append(*errs, fmt.Errorf("%%s: %%w", fieldName, err))
    }
}

`, parserVersion)
}

func printDocumentStructs(allDocumentStructs map[string]documentStructData) {
	fmt.Print("\n// Document struct definitions\n")
	for _, key := range getSortedKeys(allDocumentStructs) {
		fmt.Println(createDocumentStruct(allDocumentStructs[key]))
	}
}

func printAddDataElementFunctions(allDocumentStructs map[string]documentStructData) {
	fmt.Print("\n// AddDataElement functions\n")
	for _, key := range getSortedKeys(allDocumentStructs) {
		fmt.Println(createAddDataElementFunction(allDocumentStructs[key]))
	}
}

func printGetIDFunctions(allDocumentStructs map[string]documentStructData) {
	fmt.Print("\n// GetID functions\n")
	for _, key := range getSortedKeys(allDocumentStructs) {
		fmt.Println(createGetIDFunction(allDocumentStructs[key]))
	}
}

func printHeaderStructs(allHeaderStructs map[string]headerStructData) {
	fmt.Print("\n// Header struct definitions\n")
	for _, key := range getSortedKeys(allHeaderStructs) {
		fmt.Println(createHeaderStruct(allHeaderStructs[key]))
	}
}

func printHeaderFillFunctions(allHeaderFillFunctions map[string]headerFillMethodData) {
	fmt.Print("\n// fillHeader functions\n")
	for _, key := range getSortedKeys(allHeaderFillFunctions) {
		fmt.Println(createHeaderFillMethod(allHeaderFillFunctions[key]))
	}
}

func printDataStructs(allDataStructs map[string]dataStructData) {
	fmt.Print("\n// line data struct definitions\n")
	for _, key := range getSortedKeys(allDataStructs) {
		fmt.Println(createDataStruct(allDataStructs[key]))
	}
}

func printDataFillFunctions(allDataFillFunctions map[string]dataFillMethodData) {
	fmt.Print("\n// fillStructure functions\n")
	for _, key := range getSortedKeys(allDataFillFunctions) {
		fmt.Println(createDataFillMethod(allDataFillFunctions[key]))
	}
}

func printNewDocForIDFunction(allDocumentStructs map[string]documentStructData) {
	var data newDocForIDData
	for _, key := range getSortedKeys(allDocumentStructs) {
		data.Documents = append(data.Documents, allDocumentStructs[key])
	}
	fmt.Println(createNewDocForIDFunction(data))
}

// private functions

// returns the keys of the map, sorted alphabetically
func getSortedKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}

func getFileLineType(line string) (string, string, string, error) {
	parts := strings.Split(line, ": VERSION")
	if len(parts) < 2 {
		if line != "" {
			fmt.Println("error parsing line: met_header_columns" + "line:'" + line + "'")
		}
		return "", "", "", fmt.Errorf("error parsing line: met_header_columns line:'%s'", line)
	}
	prefix := parts[0]
	fieldStr := "VERSION " + parts[1]
	// get the version from the line
	parts = strings.Split(prefix, " : ")
	if len(parts) < 3 {
		fmt.Println("error parsing line: " + line)
		return "", "", "", fmt.Errorf("error parsing line: met_header_columns line:'%s'", line)
	}
	fileType := strings.ToUpper(strings.TrimSpace(parts[1]))
	lineType := strings.ToUpper(strings.TrimSpace(parts[2]))
	return fieldStr, fileType, lineType, nil
}

type structField struct {
	Name     string
	Type     string
	JSONName string
}

type documentStructData struct {
	DocumentStructName string
	HeaderStructName   string
	DataStructName     string
}

// Assembles the document, header, and data struct names from the fileType & lineType
func createStructNames(fileType, lineType string) documentStructData {
	docStructName := fmt.Sprintf("%s_%s", fileType, lineType)
	dataStructName := fmt.Sprintf("%s_data", docStructName)
	headerStructName := fmt.Sprintf("%s_header", docStructName)
	return documentStructData{DocumentStructName: docStructName, DataStructName: dataStructName, HeaderStructName: headerStructName}
}

// Generates the "document" struct definitions via template
func createDocumentStruct(data documentStructData) string {
	structTemplate := template.Must(template.New("struct").Parse(`
// Represents a complete {{.DocumentStructName}} document
type {{.DocumentStructName}} struct {
    mettypes.VxMetadata
    {{.HeaderStructName}}
    Data map[string]{{.DataStructName}} ` + "`json:\"data\"` //nolint:tagliatelle // \"data\" is a common JSON field in MATS" + `
}
`))
	var buf bytes.Buffer
	err := structTemplate.Execute(&buf, data)
	if err != nil {
		panic(fmt.Sprintf("template execution failed: %v", err))
	}
	return buf.String()
}

type headerStructData struct {
	DocumentStructName string
	HeaderStructName   string
	Fields             []structField
}

func getHeaderFields(headerFields []string, fileType string, lineType string, metDataTypesForLines map[string]string) []structField {
	headerStructFields := []structField{}
	if fileType == "MODE" || fileType == "MTD" {
		// these file types do not have a LINE_TYPE field in the header definition
		// from the met_header_columns file. We add a LINE_TYPE field to the header struct
		// and the fillHeader function
		headerFields = append(headerFields, `LINE_TYPE`)
	}
	dataKeyMap := util.DataKeyMap[fileType+"_"+lineType]
	keyFields := dataKeyMap.DataKey
	headerDisallowed := dataKeyMap.HeaderDisallow // list of disallowed fields
	for _, term := range headerFields {
		// skip the dataKey fields and disallowed header fields
		if slices.Contains(keyFields, term) || slices.Contains(headerDisallowed, term) {
			continue
		}
		// change regex type terms
		term = strings.ReplaceAll(term, "(", "")
		term = strings.ReplaceAll(term, ")", "")
		term = strings.ReplaceAll(term, "[0-9]*", "i")
		name := strings.ToUpper(term)
		_, dataType := getDataType(term, &metDataTypesForLines)
		jsonName := strings.ToUpper(name)
		headerStructFields = append(headerStructFields, structField{Name: name, Type: dataType, JSONName: jsonName})
	}
	return headerStructFields
}

// Generates the "header" struct definitions via template
func createHeaderStruct(data headerStructData) string {
	headerStructTemplate := template.Must(template.New("headerStruct").Parse(`
// Represents the header field of a {{.DocumentStructName}} document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type {{.HeaderStructName}} struct {
{{- range .Fields }}
	{{ .Name }} {{ .Type }} ` + "`json:\"{{ .JSONName }}\"`" + `
{{- end }}
}
	`))
	var buf bytes.Buffer
	err := headerStructTemplate.Execute(&buf, data)
	if err != nil {
		panic(fmt.Sprintf("template execution failed: %v", err))
	}
	return buf.String()
}

type headerFillField struct {
	Name             string
	HardCodeLineType bool
	Index            int
}

type headerFillMethodData struct {
	DocumentStructName string
	HeaderStructName   string
	Fields             []headerFillField
}

func getHeaderFillFields(headerFields []string, fileType string, lineType string) []headerFillField {
	var headerFillFields []headerFillField

	if fileType == "MODE" || fileType == "MTD" {
		// these file types do not have a LINE_TYPE field in the header definition
		// from the met_header_columns file. We add a LINE_TYPE field to the header struct
		// and the fillHeader function
		headerFields = append(headerFields, `LINE_TYPE`)
	}
	dataKeyMap := util.DataKeyMap[fileType+"_"+lineType]
	keyFields := dataKeyMap.DataKey
	headerDisallowed := dataKeyMap.HeaderDisallow // list of disallowed fields

	for i, field := range headerFields {
		// skip the dataKey fields and disallowed header fields
		if slices.Contains(keyFields, field) || slices.Contains(headerDisallowed, field) {
			continue
		}

		var headerField headerFillField
		headerField.Index = i
		// change regex type terms
		field = strings.ReplaceAll(field, "(", "")
		field = strings.ReplaceAll(field, ")", "")
		field = strings.ReplaceAll(field, "[0-9]*", "i")
		headerField.Name = strings.ToUpper(field)
		if field == "LINE_TYPE" && (fileType == "MODE" || fileType == "MTD") {
			headerField.HardCodeLineType = true
		} else {
			headerField.HardCodeLineType = false
		}
		headerFillFields = append(headerFillFields, headerField)
	}
	return headerFillFields
}

// Generates the header struct "Fill" methods via template
func createHeaderFillMethod(data headerFillMethodData) string {
	headerFillMethodTemplate := template.Must(template.New("headerStruct").Parse(`
// Sets {{ .HeaderStructName }} struct's fields
func (s *{{ .HeaderStructName }}) fill(fields []string) error {
	{{/*
	// TODO - this doesn't work for dynamic Fields
	expectedNumFields := {{ len .Fields }} // Length of the FieldNames slice from the template
	if len(fields) != expectedNumFields {
		return // TODO - return an error
	}
	*/}}
	var errs []error

	{{- range .Fields }}
	{{- if .HardCodeLineType}}
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte("{{ $.DocumentStructName }}"))) //hardcode the LINE_TYPE
	{{- else}}
	appendErrorWithContext(&errs, "{{ .Name }}", s.{{ .Name }}.UnmarshalText([]byte(fields[{{ .Index }}])))
	{{- end }}
	{{- end }}
	return errors.Join(errs...)
}`))
	var buf bytes.Buffer
	err := headerFillMethodTemplate.Execute(&buf, data)
	if err != nil {
		panic(fmt.Sprintf("template execution failed: %v", err))
	}
	return buf.String()
}

type dataStructData struct {
	DataStructName string
	Fields         []structField
}

func getDataFields(dataFields []string, fileType string, lineType string, metDataTypesForLines map[string]string) []structField {
	dataStructFields := []structField{}
	// reference getFillStructureString
	// add disallowed field terms to the dataFields and the associated data to the (embedded) fields array
	dataFields = append(dataFields, util.DataKeyMap[fileType+"_"+lineType].HeaderDisallow...)
	for index := 0; index < len(dataFields); index++ {
		rawFieldName := dataFields[index]
		cleanedFieldName, fieldType := getDataType(rawFieldName, &metDataTypesForLines)

		// We need to skip some fields as they end up in a map[string]interface{} instead of as direct
		// struct members
		indexOffset := 0
		if fieldType == "map[string]interface{}" {
			if rawFieldName == "(N_CAT)" {
				indexOffset = 1
			} else {
				keyPrefixes, _ := getRepeatingKeysAndTypes(rawFieldName, fileType, lineType)
				indexOffset = len(keyPrefixes)
			}
		}
		jsonTerm := strings.ToUpper(cleanedFieldName)
		dataStructFields = append(dataStructFields, structField{Name: cleanedFieldName, Type: fieldType, JSONName: jsonTerm})
		index = index + indexOffset
	}
	return dataStructFields
}

// Generates the "data" struct definitions via template
func createDataStruct(data dataStructData) string {
	dataStructTemplate := template.Must(template.New("dataStruct").Parse(`
type {{.DataStructName}} struct {
	{{- range .Fields }}
	{{ .Name }} {{ .Type }} ` + "`json:\"{{ .JSONName }},omitzero\"`" + `
	{{- end }}
}
	`))
	var buf bytes.Buffer
	err := dataStructTemplate.Execute(&buf, data)
	if err != nil {
		panic(fmt.Sprintf("template execution failed: %v", err))
	}
	return buf.String()
}

type dataFillField struct {
	Name             string
	Type             string
	IsNCAT           bool
	IsRepeatingField bool
	KeyPrefixes      []string
	Index            int
}

type dataFillMethodData struct {
	DataStructName string
	Fields         []dataFillField
}

func getDataFillFields(dataFields []string, metDataTypesForLines map[string]string, fileType string, lineType string) []dataFillField {
	dataFillFields := []dataFillField{}

	// add disallowed field terms to the dataFields and the associated data to the (embedded) fields array
	dataFields = append(dataFields, util.DataKeyMap[fileType+"_"+lineType].HeaderDisallow...)
	// iterate through the data fields to get the information needed for the data fill functions
	// use a manual for loop instead of a range as the index occasionally needs to be
	// updated multiple times inside an iteration. (TODO - could we range on this and
	// just have index as a regular variable?)
	for index := 0; index < len(dataFields); index++ {
		// Handle normal fields, we just need the fieldName & fieldType
		rawFieldName := dataFields[index]
		cleanedFieldName, fieldType := getDataType(rawFieldName, &metDataTypesForLines)
		var isNCAT, isRepeatingField bool
		var keyPrefixes []string
		indexOffset := 0 // These special cases can add more than one field at a time and will need to increment the index
		if fieldType == "map[string]interface{}" {
			// This is a repeating field, set some fields for special handling
			if rawFieldName == "(N_CAT)" {
				// special case for N_CAT fields
				isNCAT = true
				indexOffset = 1
			} else {
				// Handle all other repeating fields
				isRepeatingField = true
				keyPrefixes, fieldType = getRepeatingKeysAndTypes(rawFieldName, fileType, lineType)
				indexOffset = len(keyPrefixes)
			}
		}
		field := dataFillField{
			Name:             cleanedFieldName,
			Type:             fieldType,
			Index:            index,
			IsNCAT:           isNCAT,
			IsRepeatingField: isRepeatingField,
			KeyPrefixes:      keyPrefixes,
		}
		dataFillFields = append(dataFillFields, field)
		index = index + indexOffset
	}

	return dataFillFields
}

func add(a, b int) int {
	return a + b
}

func createDataFillMethod(data dataFillMethodData) string {
	dataFillMethodTemplate := template.New("dataStruct")
	dataFillMethodTemplate.Funcs(template.FuncMap{"add": add}) // Add the custom "add" function
	dataFillMethodTemplate = template.Must(dataFillMethodTemplate.Parse(`
// Sets {{ .DataStructName }} struct's fields
func (s *{{ .DataStructName }}) fill(fields []string) error {
	var errs []error
	{{- range .Fields }}
	{{- if .IsNCAT }}
	// these values seem to always be ints (or "NA")
	var value validtypes.ValidInt
	count, err := strconv.Atoi(fields[1])
	if err != nil {
		count = 0
	}
	s.{{ .Name }} = make(map[string]interface{})
	for i1 := {{ .Index }}; i1 <= count; i1++ {
		for i2 := 1; i2 <= count; i2++ {
			// generate the particular key for the map i.e. F1_O1, F1_O2, F1_O3, F1_O4, F2_O1, F2_O2, F2_O3, F2_O4, etc.
			key := fmt.Sprintf("F%d_O%d", i1, i2)
			index := (i1-1)*count + i2
			if index >= len(fields) {
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "{{ .Name }}", value.UnmarshalText([]byte(fields[index])))
			}
			s.{{ .Name }}[key] = value
		}
	}
	{{- else if .IsRepeatingField}}
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value {{ .Type }}
	count, err := strconv.Atoi(fields[{{ .Index }}])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{ {{- range .KeyPrefixes -}} "{{ . }}", {{- end -}}}
	s.{{ .Name }} = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := {{ add .Index 1 }}; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "{{ .Name }}", value.UnmarshalText([]byte(fields[index])))
			}
			s.{{ .Name }}[key] = value
		}
	}
	{{- else}}
	appendErrorWithContext(&errs, "{{ .Name }}", s.{{ .Name }}.UnmarshalText([]byte(fields[{{ .Index }}])))
	{{- end }}
	{{- end }}
	return errors.Join(errs...)
}`))
	var buf bytes.Buffer
	err := dataFillMethodTemplate.Execute(&buf, data)
	if err != nil {
		panic(fmt.Sprintf("template execution failed: %v", err))
	}
	return buf.String()
}

type newDocForIDData struct {
	Documents []documentStructData
}

// Generates the NewDocForID function via template
func createNewDocForIDFunction(data newDocForIDData) string {
	newDocForIDTemplate := template.Must(template.New("NewDocForID").Parse(`
// Creates an appropriate MET document struct based on the fileLineType. The MET document struct is filled in with
// vx team metadata, MET header data, and MET "data" data. The MET "data" entry is associated with the dataKey provided.
func NewDocForId(fileLineType string, metaData mettypes.VxMetadata, headerData []string, dataData []string, dataKey string) (mettypes.METdocument, error) {
	var statDoc mettypes.METdocument
	var errs []error

	switch fileLineType {
	{{- range .Documents }}
	case "{{ .DocumentStructName }}":
		elem_header := {{ .HeaderStructName }}{}
		appendErrorWithContext(&errs, "{{ .HeaderStructName }}", elem_header.fill(headerData))
		elem_data := {{ .DataStructName }}{}
		appendErrorWithContext(&errs, "{{ .DataStructName }}", elem_data.fill(dataData))

		tmp := {{ .DocumentStructName }}{
			VxMetadata:        metaData,
			{{ .HeaderStructName }}: elem_header,
			Data:              make(map[string]{{ .DataStructName }}),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	{{- end }}
	default:
		return nil, errors.New("NewDocForId: Unknown file_line type:" + fileLineType)
	}
	return statDoc, errors.Join(errs...)
}
	`))
	var buf bytes.Buffer
	err := newDocForIDTemplate.Execute(&buf, data)
	if err != nil {
		panic(fmt.Sprintf("template execution failed: %v", err))
	}
	return buf.String()
}

// Generates the AddDataElement method via template
func createAddDataElementFunction(data documentStructData) string {
	addDataElementTemplate := template.Must(template.New("AddDataElement").Parse(`
// Adds a new "data" element to {{ .DocumentStructName }}
func (doc *{{.DocumentStructName}}) AddDataElement(dataKey string, dataData []string) error {
	data := {{.DataStructName}}{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}
	`))
	var buf bytes.Buffer
	err := addDataElementTemplate.Execute(&buf, data)
	if err != nil {
		panic(fmt.Sprintf("template execution failed: %v", err))
	}
	return buf.String()
}

// Generates the GetID method via template
func createGetIDFunction(data documentStructData) string {
	addGetIDTemplate := template.Must(template.New("GetID").Parse(`
// Returns the ID field of {{ .DocumentStructName }}
func (doc *{{.DocumentStructName}}) GetID() string {
	return doc.ID
}
	`))
	var buf bytes.Buffer
	err := addGetIDTemplate.Execute(&buf, data)
	if err != nil {
		panic(fmt.Sprintf("template execution failed: %v", err))
	}
	return buf.String()
}

func getRepeatingKeysAndTypes(dataField string, fileType string, lineType string) (keyPrefixes []string, fieldType string) {
	/*

		The function definition is already in the fillStructureString in the caller so here we just need to return the
		part of the fillStructure string that does the logic of filling the repeating struct fields which are of
		differing lengths depending on the data file. This string will get appended to the caller's string. We also
		have to return the newIndex which is the number of fields that get processed in this part of the function,
		because there may be single fields remaining in the data line after the repeating fields are swallowed up, and
		the caller needs to know what the index of those fields is.

		There are several possible repeating fields depending on line and file type.
			Repeated sequences:
			Some line types can have repeated sequences of fields.
			These fields must be represented in the data sections of documents as embedded maps.
			These are identified by the N_ prefix in the field name and the field name being surrounded by parenthesis.
			The key for the map is patterned after the field name with the N_ prefix removed.

			(N_CAT) for MCTC files, the repeated sequence is F[0-9]*_O[0-9]* that will be contained in a map[string]int.

			(N_THRESH) for PCT which is a sequence of THRESH_n OY_n ON_n key/values that will be contained in a map[string]interface{}
			where the keys are THRESH_n, OY_n, or ON_n with the n being the threshold number.

			(N_THRESH) for PJC which is a sequence of THRESH_n, OY_TP_n, ON_TP_n, CALIBRATION_n, REFINEMENT_n, LIKELIHOOD_n, BASER_n key/values
			that will be contained in a map[string]interface{}
			where the keys are THRESH_n, OY_TP_n, ON_TP_n, CALIBRATION_n, REFINEMENT_n, LIKELIHOOD_n, and BASER_n.

			(N_THRESH) for PRC which is a sequence of THRESH_n PODY_n POFD_n key/values that will be contained in a map[string]interface{}
			where the keys are THRESH_n PODY_n POFD_n.

			(N_THRESH) for PSTD which is a map[string]interface{}
			where the keys are all THRESH_n

			(N_THRESH) for PROBRIRW which is a sequence of THRESH_n PROB_n that will be contained in a map[string]interface{}
			where the keys are THRESH_n and PROB_n.

			(N_PTS) for ECLV files which is a sequence of CL_n VALUE_n key/values that will be contained in a map[string]interface{}
			where the keys are CL_n and VALUE_n e.g. CL_1, VALUE_1, CL_2, VALUE_2, etc.

			(N_RANK) for RHIST files, the repeated sequence is RANK_n that will be contained in a map[string]int
			where the keys are RANK_n e.g. RANK_1, RANK_2, etc.

			(N_BIN) for PHIST files, the repeated sequence is BIN_n that will be contained in a map[string]int
			where the keys are BIN_n e.g. BIN_1, BIN_2, etc.

			(N_ENS) for ORANK files, the repeated sequence is ENS_n that will be contained in a map[string]interface{}
			where the keys are ENS_n e.g. ENS_1, ENS_2, etc.

			(N_ENS) for RELP files, the repeated sequence is RELP_n that will be contained in a map[string]interface{}
			where the keys are RELP_N e.g. RELP_1, RELP_2, etc.

			(N_DIAG) for TCDIAG files, the repeated sequence is DIAG_n VALUE_n that will be contained in a map[string]interface{}
			where the keys are DIAG_n and VALUE_n e.g. DIAG_1, VALUE_1, DIAG_2, VALUE_2 etc.
	*/
	fileLineType := fileType + "_" + lineType
	switch dataField {
	case "(N_THRESH)": // PCT, PJC, PRC, PSTD, PROBRIRW files
		switch lineType {
		case "PCT":
			// THRESH PCT which is a sequence of THRESH_n OY_n ON_n which are float64 values
			return []string{"THRESH_", "OY_", "ON_"}, "validtypes.ValidFloat"
		case "PJC":
			// PJC files have a sequence of THRESH_n OY_TP_n ON_TP_n CALIBRATION_n REFINEMENT_n LIKELIHOOD_n BASER_n which are float64 values
			return []string{"THRESH_", "OY_TP_", "ON_TP_", "CALIBRATION_", "REFINEMENT", "LIKELIHOOD_", "BASER_"}, "validtypes.ValidFloat"
		case "PRC":
			// PRC files have a sequence of THRESH_n PODY_n POFD_n which are float64 values
			return []string{"THRESH_", "PODY_", "POFD_"}, "validtypes.ValidFloat"
		case "PSTD":
			// PSTD files have a sequence of THRESH_n which are float64 values
			return []string{"THRESH_"}, "validtypes.ValidFloat"
		case "PROBRIRW":
			// PROBRIRW files have a sequence of THRESH_n PROB_n which are int values
			return []string{"THRESH_", "PROB_"}, "validtypes.ValidInt"
		}

	case "(N_PTS)": // ECLV files (N_PTS) for ECLV files which is a sequence of CL_n VALUE_n which are float64 values
		return []string{"CL_", "VALUE_"}, "validtypes.ValidFloat"
	case "(N_RANK)": // RHIST files (N_RANK) for RHIST files, the repeated sequence is RANK_n which are int values
		return []string{"RANK_"}, "validtypes.ValidInt"
	case "(N_BIN)": // PHIST files (N_BIN) for PHIST files, the repeated sequence is BIN_n which are int values
		return []string{"BIN_"}, "validtypes.ValidInt"
	case "(N_ENS)": // ORANK, RELP files (N_ENS) for ORANK files, the repeated sequence is ENS_n
		if fileLineType == "STAT_ORANK" {
			// (N_ENS) for ORANK files, the repeated sequence is ENS_n which are ints (can have NA values)
			return []string{"ENS_"}, "validtypes.ValidInt"
		}
		if fileLineType == "STAT_RELP" {
			// (N_ENS) for RELP files, the repeated sequence is RELP_n which are float64 values
			return []string{"RELP_"}, "validtypes.ValidFloat"
		}
	case "(N_DIAG)": // TCDIAG files (no sample data for this type)
		return []string{"DIAG_", "VALUE_"}, "validtypes.ValidString"
	}
	return []string{}, ""
}

// Fetches & scans the given MET source files for specific C/C++ type conversion fields, and then sets the token type based on the results
func fillMetDataMapFromSrcFiles(srcFileURLs []string, metDataTypesForLines map[string]string, fieldNameMap map[string]string) map[string]string {
	// use a map (atoLines) as a set to avoid duplicate lines
	atoLines := make(map[string]bool)
	// iterate through the srcFileURLs to get the data types for the fields
	matchConvertLine := regexp.MustCompile(`= ato[fi]\(l.get_item\(`)
	for _, url := range srcFileURLs {
		lines := getLinesForUrl(url)
		// iterate through the lines to find the data types
		for _, line := range lines {
			parts := matchConvertLine.Split(line, -1)
			if len(parts) > 1 {
				atoLines[line] = true
			}
		}
	}
	// iterate the fieldNames to see if we can find data types in the atoLines from the src code files
	for token := range fieldNameMap {
		// iterate through the atoLines to get any data types
		for line := range atoLines {
			if matched, err := regexp.Match(token, []byte(line)); err == nil && matched {
				parts := strings.Split(line, " =")
				if len(parts) > 1 {
					atPart := strings.TrimSpace(strings.Split(parts[1], "(")[0])
					var dataType string
					switch atPart {
					case "atoi":
						dataType = "validtypes.ValidInt"
					case "atof":
						dataType = "validtypes.ValidFloat"
					default:
						dataType = "validtypes.ValidString"
					}
					fieldNameMap[token] = dataType
					metDataTypesForLines[token] = dataType
				}
			}
		}
	}
	return metDataTypesForLines
}

func fillMetDataMapFromUserGuide(userGuideURLS []string, metDataTypesForLines, fieldNameMap map[string]string) map[string]string {
	// MET user guide files with data type definitions
	// Using the slower regexp instead of a string match because I don't know if the line will have
	// extra leading spaces or not. These documents might get reformatted and the leading spaces might
	// change. They are less likely to change the order of the fields in the table.
	// The regex to split the line into fields
	linePrefix := regexp.MustCompile(`^ *- `)
	// The regexp to identify the start of a column header
	lineColumnStart := regexp.MustCompile(`^\s*\* - Column`)
	for _, url := range userGuideURLS {
		docFileLines := getLinesForUrl(url)
		var parts []string
		for i := 0; i < len(docFileLines)-1; i++ {
			line := docFileLines[i]
			if line == "" {
				continue
			}
			/*
					look for the lines that have data type table entries - they all start with "  * - " and
					have Column in the first line. If a line starts with "  * - " and does not have Column in it
					then it is the end of the data type table.
					* - Column Number       - start of table (this is the header)
					- Header Column Name
					- Description
					- Data Type
				* - 1                       - 1st entry of the table (this is the column number)
					- VERSION               - field name
					- Version number        - field description
					- String                - data type
			*/
			// skip the lines that are not the start of the data type table i.e. don't start with '* - Column Number'
			if strings.Contains(line, "* - Column") {
				if !strings.Contains(docFileLines[i+1], "Name") ||
					!strings.Contains(docFileLines[i+2], "Description") ||
					!strings.Contains(docFileLines[i+3], "Data Type") {
					continue // this is not the start of a table we are interested in
				}
				// process the table
				for ; i < len(docFileLines)-1; i++ {
					line = docFileLines[i]
					if line == "" {
						break // this table is ended, go to the next one.
					}
					// is this a Column Header Start line (header of table)?
					if lineColumnStart.MatchString(line) {
						// skip to the end of the table header
						i = i + 4
					}
					// skip the first line - it is the column number
					i++
					line = docFileLines[i]
					// remove possible embedded html - ugh!
					line = strings.ReplaceAll(line, ":raw-html:`<br />`", "")
					line = strings.ReplaceAll(line, `\`, "")
					if line == "    -" {
						break // this table is formatted poorly (empty fieldName), go to the next one.
					}
					parts = linePrefix.Split(line, -1)
					fieldName := strings.ReplaceAll(parts[1], " ", "") // remove extra spaces
					parts := strings.Split(fieldName, "/")             // remove any front slashes
					if len(parts) > 1 {
						fieldName = parts[1]
					}
					// It seems that the actual fieldNames that are specified in the doc as abc_i are labeled as abc_[0-9]* in the code
					fieldName = strings.ReplaceAll(fieldName, `_i`, `_[0-9]*`) // replace _i with _[0-9]
					// skip the description line
					i = i + 2
					line = docFileLines[i]
					// skip any extra line (for some reason these lines can have notes and things that take up multiple lines) that do not start with "* - "
					if !linePrefix.MatchString(line) {
						i = i + 1
						line = docFileLines[i]
					}
					parts = linePrefix.Split(line, -1)
					dataType := strings.TrimSpace(parts[1]) // get the data type
					// any unknown data types will default to string
					switch dataType {
					case "Integer":
						dataType = "validtypes.ValidInt"
					case "Double":
						dataType = "validtypes.ValidFloat"
					case "String":
						dataType = "validtypes.ValidString"
					default:
						dataType = "validtypes.ValidString"
					}
					fields := []string{}
					// INTENSITY fields seem to be sort of one-off fields that have partial subfields i.e. "INTENSITY_10, _25, _50, _75, _90"
					// These must be handled differently to get the actual subfields.
					if strings.Contains(fieldName, "INTENSITY") ||
						strings.Contains(fieldName, "CURVATURE") ||
						strings.Contains(fieldName, "CENTROID") {
						parts = strings.Split(fieldName, ",")
						fields = append(fields, parts[0])
						if len(parts) > 1 {
							pre := strings.Split(parts[0], "_")[0]
							for _, p := range parts[1:] {
								subFieldName := pre + strings.TrimSpace(p)
								fields = append(fields, subFieldName)
							}
						} else {
							fields = append(fields, fieldName)
						}
					} else {
						// some fields have multiple subfields i.e. "FCST_LEAD, FCST_LEV"
						// These also must be handled differently to get the actual subfields.
						majorFields := strings.Split(fieldName, ",")
						if len(majorFields) > 1 {
							for _, majorField := range majorFields {
								// remove any  remaining spaces - some fields have nonsense spaces in them
								subfield := strings.ReplaceAll(majorField, " ", "")
								fields = append(fields, subfield)
							}
						} else {
							fields = append(fields, fieldName)
						}
					}
					// now fields should have all the field names including subfields
					for _, fieldName := range fields {
						if fieldName != "" {
							/*
								NOTE: many of the fieldNameMap entries remain "UNDEFINED" because the
								data type is not found in the MET user guide files. The data type
								will default to "string" in the generated code. This is the best that
								can be done without a more comprehensive data type definition. I am
								leaving the fieldNameMap entries as a way to track down
								the missing data types, but it is not used in the code generation.
							*/
							metDataTypesForLines[fieldName] = dataType
							fieldNameMap[fieldName] = dataType
						}
					}
				}
			}
		}
	}
	// Fill undefined's that I have not found in the MET user guide files in text (not column tables), or in data files themselves.
	// NOTE: These will overwrite any previous data types for specific named fields that were found in the MET source files.
	metDataTypesForLines, _ = overRideDefinedMetDataTypes(metDataTypesForLines, fieldNameMap)

	return metDataTypesForLines
}

// overRideDefinedMetDataTypes manually overrides the data types for specific fields in the fieldNameMap & metDataTypesForLines that are incorrectly
// defined in, or missing from the MET source code & documentation that we reference.
func overRideDefinedMetDataTypes(metDataTypesForLines map[string]string, fieldNameMap map[string]string) (map[string]string, map[string]string) {
	metDataTypesForLines["RIRW_WINDOW"] = "validtypes.ValidInt"
	metDataTypesForLines["F[0-9]*_O[0-9]*"] = "validtypes.ValidString"
	metDataTypesForLines["INTENSITY_USER"] = "validtypes.ValidFloat"
	metDataTypesForLines["INTENSITY_USER_MIN"] = "validtypes.ValidFloat"
	metDataTypesForLines["INTENSITY_USER_MAX"] = "validtypes.ValidFloat"
	metDataTypesForLines["RPS_COMP"] = "validtypes.ValidFloat"
	metDataTypesForLines["ARADP"] = "validtypes.ValidString"
	metDataTypesForLines["AMRD"] = "validtypes.ValidInt"
	metDataTypesForLines["AGUSTS"] = "validtypes.ValidInt"
	metDataTypesForLines["ADIR"] = "validtypes.ValidInt"
	metDataTypesForLines["AEYE"] = "validtypes.ValidInt"
	metDataTypesForLines["EIQR_BCL"] = "validtypes.ValidFloat"
	metDataTypesForLines["EIQR_BCU"] = "validtypes.ValidFloat"
	metDataTypesForLines["ASPEED"] = "validtypes.ValidInt"
	metDataTypesForLines["ARRP"] = "validtypes.ValidInt"
	metDataTypesForLines["ADEPTH"] = "validtypes.ValidString"         // TCST_TCMPR - Should be one of "D,M,S,X" per the link from the met-tc overview: https://science.nrlmry.navy.mil/atcf/docs/database/new/abdeck.txt
	metDataTypesForLines["BDEPTH"] = "validtypes.ValidString"         // TCST_TCMPR - Should be one of "D,M,S,X" per the link from the met-tc overview: https://science.nrlmry.navy.mil/atcf/docs/database/new/abdeck.txt
	metDataTypesForLines["DIAG_SOURCE"] = "validtypes.ValidString"    // TCST_TCDIAG
	metDataTypesForLines["OBS_CLIMO_MEAN"] = "validtypes.ValidString" // STAT_ORANK

	fieldNameMap["RIRW_WINDOW"] = "validtypes.ValidInt"
	fieldNameMap["F[0-9]*_O[0-9]*"] = "validtypes.ValidString"
	fieldNameMap["INTENSITY_USER"] = "validtypes.ValidFloat"
	fieldNameMap["INTENSITY_USER_MIN"] = "validtypes.ValidFloat"
	fieldNameMap["INTENSITY_USER_MAX"] = "validtypes.ValidFloat"
	fieldNameMap["RPS_COMP"] = "validtypes.ValidFloat"
	fieldNameMap["ARADP"] = "validtypes.ValidString"
	fieldNameMap["AMRD"] = "validtypes.ValidInt"
	fieldNameMap["AGUSTS"] = "validtypes.ValidInt"
	fieldNameMap["ADIR"] = "validtypes.ValidInt"
	fieldNameMap["AEYE"] = "validtypes.ValidInt"
	fieldNameMap["EIQR_BCL"] = "validtypes.ValidFloat"
	fieldNameMap["EIQR_BCU"] = "validtypes.ValidFloat"
	fieldNameMap["ASPEED"] = "validtypes.ValidInt"
	fieldNameMap["ARRP"] = "validtypes.ValidInt"
	fieldNameMap["ADEPTH"] = "validtypes.ValidString"         // TCST_TCMPR - Should be one of "D,M,S,X" per the link from the met-tc overview: https://science.nrlmry.navy.mil/atcf/docs/database/new/abdeck.txt
	fieldNameMap["BDEPTH"] = "validtypes.ValidString"         // TCST_TCMPR - Should be one of "D,M,S,X" per the link from the met-tc overview: https://science.nrlmry.navy.mil/atcf/docs/database/new/abdeck.txt
	fieldNameMap["DIAG_SOURCE"] = "validtypes.ValidString"    // TCST_TCDIAG
	fieldNameMap["OBS_CLIMO_MEAN"] = "validtypes.ValidString" // STAT_ORANK

	// Uncomment the following to look for missing data types in the MET user guide files.
	var found bool
	undefineds := []string{}
	for k, v := range fieldNameMap {
		if v == "UNDEFINED" {
			found = false
			for _, v1 := range getPatterns() {
				if v1.match.MatchString(strings.ToUpper(k)) {
					found = true
					break
				}
			}
			if !found {
				undefineds = append(undefineds, k)
			}
		}
	}
	if len(undefineds) > 0 {
		slices.Sort(undefineds)
		fmt.Printf(`
/*
The following data types were not found in the MET user guide files or the MET source code files.
For simplicity, the values of these data types will be treated as strings in the generated code.

    Undefined data types: %v

To resolve this, consult the github.com/dtcenter/MET repo to determine if there is a more appropriate type, 
and, if there is, add an override to the overRideDefinedMetDataTypes function in generator/generator.go.
*/
`, undefineds)
	}
	return metDataTypesForLines, fieldNameMap
}

// Returns a slice of regex patterns that identify dynamic field sequences in MET data
// files, distinguishing between count fields (like N_THRESH) and their corresponding
// numbered data fields (like THRESH_1, PROB_1), along with their Go data types for
// code generation.
func getPatterns() []Pattern {
	// PROBRIRW example
	// VERSION AMODEL BMODEL DESC STORM_ID BASIN CYCLONE STORM_NAME INIT            LEAD   VALID           INIT_MASK VALID_MASK LINE_TYPE ALAT ALON  BLAT BLON  INITIALS TK_ERR    X_ERR      Y_ERR ADLAND     BDLAND     RIRW_BEG RIRW_END RIRW_WINDOW AWIND_END BWIND_BEG BWIND_END BDELTA BDELTA_MAX BLEVEL_BEG BLEVEL_END N_THRESH THRESH_1 PROB_1 THRESH_2 PROB_2 THRESH_3 PROB_3 THRESH_4 PROB_4 THRESH_5 PROB_5
	// V12.0.0 GPMI   BEST   NA   AL012015 AL    01      ANA        20150508_120000 240000 20150509_120000 NA        NA         PROBRIRW  31.6 -77.7 32.5 -77.8       NA  54.23894    5.08551   -54  135.63956   80.31028        0       24          24        44        40        50     10         10         TS         TS        5      -30      0      -10      0        0    100       10      0       30      0

	// these `\(N_[A-Z]+\)` are repeated sequences of data fields - a map of data fields
	// that indicate repeated sequences of data fields in the data sections.
	// These do not have direct corresponding data elements, but they have
	// corresponding sequences of data elements in the data sections.
	// Those elements are identified in the actual data section headers by
	// things like THRESH_[0-9]* PROB_[0-9]* for PROBRIRW. There are associated data
	// elements in the data sections that are of different types and they are
	// defined by tokens like THRESH_1 PROB_1 THRESH_2 PROB_2 THRESH_3 PROB_3
	// Since these are essentially key value pairs of varying length (per individual data file)
	// they are not directly represented in the data structures as a map.
	// The data structures are created dynamically based on the data files and the types
	// of the values in the map are specific to each pattern e.g. bin_n values are ints
	// but cl_n values are float64s.
	// structField is for repeating fields only. These patterns are surrounded by "()" i.e. (N_CAT).
	// The structure generator needs to know what it is that is repeating. Most likely a map of some sort.

	return []Pattern{
		// repeating patterns
		{match: regexp.MustCompile(`(N_CAT)`), dType: "validtypes.ValidInt", structField: "CAT", structType: "map[string]interface{}"},
		{match: regexp.MustCompile(`(N_THRESH)`), dType: "validtypes.ValidInt", structField: "THRESH", structType: "map[string]interface{}"},
		{match: regexp.MustCompile(`(N_PTS)`), dType: "validtypes.ValidInt", structField: "PTS", structType: "map[string]interface{}"},
		{match: regexp.MustCompile(`(N_ENS)`), dType: "validtypes.ValidInt", structField: "ENS", structType: "map[string]interface{}"},
		{match: regexp.MustCompile(`(N_RANK)`), dType: "validtypes.ValidInt", structField: "RANK", structType: "map[string]interface{}"},
		{match: regexp.MustCompile(`(N_BIN)`), dType: "validtypes.ValidInt", structField: "BIN", structType: "map[string]interface{}"},
		{match: regexp.MustCompile(`(N_DIAG)`), dType: "validtypes.ValidInt", structField: "DIAG", structType: "map[string]interface{}"},
		// single patterns
		{match: regexp.MustCompile("BASER_[0-9]*"), dType: "validtypes.ValidFloat", structField: "BASER_I", structType: "map[string]interface{}"},
		{match: regexp.MustCompile("BIN_[0-9]*"), dType: "validtypes.ValidFloat", structField: "BIN_I", structType: "validtypes.ValidInt"}, // float? STAT_PHIST - BIN_SIZE 0.05 - see getRepeatingKeysAndType as well
		{match: regexp.MustCompile("CALIBRATION_[0-9]*"), dType: "validtypes.ValidFloat", structField: "CALIBRATION_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("CL_[0-9]*"), dType: "validtypes.ValidFloat", structField: "CL_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("DIAG_[0-9]*"), dType: "validtypes.ValidFloat", structField: "DIAG_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("ENS_[0-9]*"), dType: "validtypes.ValidFloat", structField: "ENS_I", structType: "validtypes.ValidFloat"}, // float? STAT_ORANK - ENS_MEAN 272.52432
		{match: regexp.MustCompile("F[0-9]*_O[0-9]*"), dType: "validtypes.ValidString", structField: "FI_OI", structType: "validtypes.ValidString"},
		{match: regexp.MustCompile("[A-Z]F[0-9]*_[A-Z]O[0-9]*"), dType: "validtypes.ValidString", structField: "AZFI_AZOI", structType: "validtypes.ValidString"},
		{match: regexp.MustCompile("LIKELIHOOD_[0-9]*"), dType: "validtypes.ValidFloat", structField: "LIKELIHOOD_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("AAL_WIND_[0-9]*"), dType: "validtypes.ValidFloat", structField: "AAL_WIND_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("ASE_WIND_[0-9]*"), dType: "validtypes.ValidFloat", structField: "ASE_WIND_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("ASW_WIND_[0-9]*"), dType: "validtypes.ValidFloat", structField: "ASW_WIND_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("ANE_WIND_[0-9]*"), dType: "validtypes.ValidFloat", structField: "ANE_WIND_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("ANW_WIND_[0-9]*"), dType: "validtypes.ValidFloat", structField: "ANW_WIND_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("ON_TP_[0-9]*"), dType: "validtypes.ValidFloat", structField: "ON_TP_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("ON_[0-9]*"), dType: "validtypes.ValidFloat", structField: "ON_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("OY_TP_[0-9]*"), dType: "validtypes.ValidFloat", structField: "OY_TP_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("OY_[0-9]*"), dType: "validtypes.ValidFloat", structField: "OY_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("PODY_[0-9]*"), dType: "validtypes.ValidFloat", structField: "PODY_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("POFD_[0-9]*"), dType: "validtypes.ValidFloat", structField: "POFD_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("PROB_[0-9]*"), dType: "validtypes.ValidFloat", structField: "PROB_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("RANK_[0-9]*"), dType: "validtypes.ValidInt", structField: "RANK_I", structType: "validtypes.ValidInt"},
		{match: regexp.MustCompile("REFINEMENT_[0-9]*"), dType: "validtypes.ValidFloat", structField: "REFINEMENT_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("RELP_[0-9]*"), dType: "validtypes.ValidFloat", structField: "RELP_I", structType: "validtypes.ValidFloat"},
		{match: regexp.MustCompile("THRESH_[0-9]*"), dType: "validtypes.ValidInt", structField: "THRESH_I", structType: "validtypes.ValidInt"},
		{match: regexp.MustCompile("VALUE_[0-9]*"), dType: "validtypes.ValidFloat", structField: "VALUE_I", structType: "validtypes.ValidFloat"},
	}
}

// reads the met_header_columns plaintext files published by the MET team, and returns the raw lines plus a map of each token in
// the lines, with values initialized to UNDEFINED.
//
// Assumptions - each entry in the map is assumed to be unique or, if conflicts arise, is assumed to be of the same type.
//
// Returns:
// - headerLines: All raw lines from the met_header_file
// - fieldTypeByToken: map[token]type initialized to "UNDEFINED" for each token in the met_header_file
func getColumnLinesAndMapForUrl(fileUrl string) ([]string, map[string]string) {
	fieldTypeByToken := make(map[string]string)
	headerLines := getLinesForUrl(fileUrl)
	// split out all the fields to get a map of required fields
	for _, line := range headerLines {
		// get the prefix from the line
		parts := strings.Split(line, ": VERSION")
		if len(parts) < 2 {
			if line != "" {
				fmt.Println("error parsing line: met_header_columns" + "line:'" + line + "'")
			}
			continue
		}
		// get all the required fields for this line
		fields := strings.Fields(parts[1])
		// use a map as a SET to get unique field names
		// this makes an assumption the field names are unique or that
		// the field names have the same data type all the different structs i.e. columnDef lines.
		for _, field := range fields {
			name := strings.ToLower(field)
			fieldTypeByToken[strings.ToUpper(name)] = "UNDEFINED"
		}
	}
	return headerLines, fieldTypeByToken
}

// Fetches a plain-text file from a URL, and returns its contents split by newline
//
// E.g. - https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt
// E.g. - https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/src/libcode/vx_analysis_util/mode_line.cc
//
// returns:
//   - lines: File content split on '\n'
func getLinesForUrl(fileUrl string) []string {
	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println("error getting met_header_columns file" + err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	rawColumnsBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		// Using the header definitions in the appropriate version of data/table_files/met_header_columns_X.X.txt
		os.Exit(1)
	}
	lines := strings.Split(string(rawColumnsBytes), "\n")
	return lines
}

func getDataType(name string, metDataTypesLines *map[string]string) (key string, dType string) {
	/*
		The getDataType function is used to get the data type for a field name. The field name is used to look up the data type in the metDataTypesLines map.
		If the data type is not found in the map then the data type is set to "string". The field name is returned as the key and the data type is returned as the dType.
		There are some fields that are arrays of values and these are handled differently. The field name is modified to reflect the array of values and the data type is set to the array type.
		Some interpretations of the patterns depend upon the line type.
	*/
	uName := strings.ToUpper(name)
	// is it exactly the same as a pattern? This is the case for a structure repeating field.
	for _, v := range getPatterns() {
		if v.match.String() == uName {
			return v.structField, v.structType
		}
	}
	// does it match a pattern? This is the case for a structure simple field.
	for _, v := range getPatterns() {
		if v.match.MatchString(uName) {
			return uName, v.dType
		}
	}

	// it wasn't one of the patterns so try to look it up in the metDataTypesLines
	dataType := (*metDataTypesLines)[uName]
	// it could still be undefined - if it is then set it to string
	if dataType == "" {
		dataType = "validtypes.ValidString"
	}
	// convert the type of a field that is a date to "int"
	if slices.Contains(util.DateFieldNames, uName) {
		dataType = "validtypes.ValidInt"
	}
	// convert the type of a field that is an int to "int"
	if slices.Contains(util.IntFieldNames, uName) {
		dataType = "validtypes.ValidInt"
	}
	return uName, dataType
}
