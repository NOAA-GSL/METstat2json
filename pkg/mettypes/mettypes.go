package mettypes

// interface to describe MET document structs in the various linetypes.go files
type METdocument interface {
	AddDataElement(key string, dataData []string) error
	GetID() string
}

// vxMetadata struct definition
//
//nolint:tagliatelle // We need these fields to match other metadata fields in our documents.
type VxMetadata struct {
	ID          string `json:"id"`
	Subset      string `json:"subset"`
	Type        string `json:"type"`
	SubType     string `json:"subtype"`
	DataSetName string `json:"dataSetName"`
}
