package v10_1

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dtcenter/METstat2json/pkg/util"
	"github.com/dtcenter/METstat2json/pkg/validtypes"
)

/*
THIS CODE IS AUTOMATICALLY GENERATED - DO NOT EDIT THIS CODE
To modify this code - modify the generator.go file and run the generator.go program
cd  <repo_root>
go run generator -version=v12.0 > pkg/linetypes/v12_0/linetypes.go
*/

// Helper function to reduce boilerplate for using errors.Join()
func appendErrorWithContext(errs *[]error, fieldName string, err error) {
	if err != nil {
		*errs = append(*errs, fmt.Errorf("%s: %w", fieldName, err))
	}
}

// Document struct definitions

// Represents a complete MODE_CTS document
type MODE_CTS struct {
	util.VxMetadata
	MODE_CTS_header
	Data map[string]MODE_CTS_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete MODE_OBJ document
type MODE_OBJ struct {
	util.VxMetadata
	MODE_OBJ_header
	Data map[string]MODE_OBJ_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_CNT document
type STAT_CNT struct {
	util.VxMetadata
	STAT_CNT_header
	Data map[string]STAT_CNT_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_CTC document
type STAT_CTC struct {
	util.VxMetadata
	STAT_CTC_header
	Data map[string]STAT_CTC_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_CTS document
type STAT_CTS struct {
	util.VxMetadata
	STAT_CTS_header
	Data map[string]STAT_CTS_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_DMAP document
type STAT_DMAP struct {
	util.VxMetadata
	STAT_DMAP_header
	Data map[string]STAT_DMAP_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_ECLV document
type STAT_ECLV struct {
	util.VxMetadata
	STAT_ECLV_header
	Data map[string]STAT_ECLV_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_ECNT document
type STAT_ECNT struct {
	util.VxMetadata
	STAT_ECNT_header
	Data map[string]STAT_ECNT_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_FHO document
type STAT_FHO struct {
	util.VxMetadata
	STAT_FHO_header
	Data map[string]STAT_FHO_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_GENMPR document
type STAT_GENMPR struct {
	util.VxMetadata
	STAT_GENMPR_header
	Data map[string]STAT_GENMPR_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_GRAD document
type STAT_GRAD struct {
	util.VxMetadata
	STAT_GRAD_header
	Data map[string]STAT_GRAD_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_ISC document
type STAT_ISC struct {
	util.VxMetadata
	STAT_ISC_header
	Data map[string]STAT_ISC_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_MCTC document
type STAT_MCTC struct {
	util.VxMetadata
	STAT_MCTC_header
	Data map[string]STAT_MCTC_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_MCTS document
type STAT_MCTS struct {
	util.VxMetadata
	STAT_MCTS_header
	Data map[string]STAT_MCTS_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_MPR document
type STAT_MPR struct {
	util.VxMetadata
	STAT_MPR_header
	Data map[string]STAT_MPR_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_NBRCNT document
type STAT_NBRCNT struct {
	util.VxMetadata
	STAT_NBRCNT_header
	Data map[string]STAT_NBRCNT_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_NBRCTC document
type STAT_NBRCTC struct {
	util.VxMetadata
	STAT_NBRCTC_header
	Data map[string]STAT_NBRCTC_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_NBRCTS document
type STAT_NBRCTS struct {
	util.VxMetadata
	STAT_NBRCTS_header
	Data map[string]STAT_NBRCTS_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_ORANK document
type STAT_ORANK struct {
	util.VxMetadata
	STAT_ORANK_header
	Data map[string]STAT_ORANK_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_PCT document
type STAT_PCT struct {
	util.VxMetadata
	STAT_PCT_header
	Data map[string]STAT_PCT_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_PHIST document
type STAT_PHIST struct {
	util.VxMetadata
	STAT_PHIST_header
	Data map[string]STAT_PHIST_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_PJC document
type STAT_PJC struct {
	util.VxMetadata
	STAT_PJC_header
	Data map[string]STAT_PJC_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_PRC document
type STAT_PRC struct {
	util.VxMetadata
	STAT_PRC_header
	Data map[string]STAT_PRC_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_PSTD document
type STAT_PSTD struct {
	util.VxMetadata
	STAT_PSTD_header
	Data map[string]STAT_PSTD_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_RELP document
type STAT_RELP struct {
	util.VxMetadata
	STAT_RELP_header
	Data map[string]STAT_RELP_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_RHIST document
type STAT_RHIST struct {
	util.VxMetadata
	STAT_RHIST_header
	Data map[string]STAT_RHIST_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_RPS document
type STAT_RPS struct {
	util.VxMetadata
	STAT_RPS_header
	Data map[string]STAT_RPS_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_SAL1L2 document
type STAT_SAL1L2 struct {
	util.VxMetadata
	STAT_SAL1L2_header
	Data map[string]STAT_SAL1L2_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_SL1L2 document
type STAT_SL1L2 struct {
	util.VxMetadata
	STAT_SL1L2_header
	Data map[string]STAT_SL1L2_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_SSIDX document
type STAT_SSIDX struct {
	util.VxMetadata
	STAT_SSIDX_header
	Data map[string]STAT_SSIDX_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_SSVAR document
type STAT_SSVAR struct {
	util.VxMetadata
	STAT_SSVAR_header
	Data map[string]STAT_SSVAR_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_VAL1L2 document
type STAT_VAL1L2 struct {
	util.VxMetadata
	STAT_VAL1L2_header
	Data map[string]STAT_VAL1L2_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_VCNT document
type STAT_VCNT struct {
	util.VxMetadata
	STAT_VCNT_header
	Data map[string]STAT_VCNT_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete STAT_VL1L2 document
type STAT_VL1L2 struct {
	util.VxMetadata
	STAT_VL1L2_header
	Data map[string]STAT_VL1L2_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete TCST_PROBRIRW document
type TCST_PROBRIRW struct {
	util.VxMetadata
	TCST_PROBRIRW_header
	Data map[string]TCST_PROBRIRW_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// Represents a complete TCST_TCMPR document
type TCST_TCMPR struct {
	util.VxMetadata
	TCST_TCMPR_header
	Data map[string]TCST_TCMPR_data `json:"data"` //nolint:tagliatelle // "data" is a common JSON field in MATS
}

// AddDataElement functions

// Adds a new "data" element to MODE_CTS
func (doc *MODE_CTS) AddDataElement(dataKey string, dataData []string) error {
	data := MODE_CTS_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to MODE_OBJ
func (doc *MODE_OBJ) AddDataElement(dataKey string, dataData []string) error {
	data := MODE_OBJ_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_CNT
func (doc *STAT_CNT) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_CNT_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_CTC
func (doc *STAT_CTC) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_CTC_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_CTS
func (doc *STAT_CTS) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_CTS_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_DMAP
func (doc *STAT_DMAP) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_DMAP_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_ECLV
func (doc *STAT_ECLV) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_ECLV_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_ECNT
func (doc *STAT_ECNT) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_ECNT_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_FHO
func (doc *STAT_FHO) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_FHO_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_GENMPR
func (doc *STAT_GENMPR) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_GENMPR_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_GRAD
func (doc *STAT_GRAD) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_GRAD_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_ISC
func (doc *STAT_ISC) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_ISC_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_MCTC
func (doc *STAT_MCTC) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_MCTC_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_MCTS
func (doc *STAT_MCTS) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_MCTS_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_MPR
func (doc *STAT_MPR) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_MPR_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_NBRCNT
func (doc *STAT_NBRCNT) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_NBRCNT_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_NBRCTC
func (doc *STAT_NBRCTC) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_NBRCTC_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_NBRCTS
func (doc *STAT_NBRCTS) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_NBRCTS_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_ORANK
func (doc *STAT_ORANK) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_ORANK_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_PCT
func (doc *STAT_PCT) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_PCT_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_PHIST
func (doc *STAT_PHIST) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_PHIST_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_PJC
func (doc *STAT_PJC) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_PJC_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_PRC
func (doc *STAT_PRC) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_PRC_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_PSTD
func (doc *STAT_PSTD) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_PSTD_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_RELP
func (doc *STAT_RELP) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_RELP_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_RHIST
func (doc *STAT_RHIST) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_RHIST_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_RPS
func (doc *STAT_RPS) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_RPS_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_SAL1L2
func (doc *STAT_SAL1L2) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_SAL1L2_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_SL1L2
func (doc *STAT_SL1L2) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_SL1L2_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_SSIDX
func (doc *STAT_SSIDX) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_SSIDX_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_SSVAR
func (doc *STAT_SSVAR) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_SSVAR_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_VAL1L2
func (doc *STAT_VAL1L2) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_VAL1L2_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_VCNT
func (doc *STAT_VCNT) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_VCNT_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to STAT_VL1L2
func (doc *STAT_VL1L2) AddDataElement(dataKey string, dataData []string) error {
	data := STAT_VL1L2_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to TCST_PROBRIRW
func (doc *TCST_PROBRIRW) AddDataElement(dataKey string, dataData []string) error {
	data := TCST_PROBRIRW_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Adds a new "data" element to TCST_TCMPR
func (doc *TCST_TCMPR) AddDataElement(dataKey string, dataData []string) error {
	data := TCST_TCMPR_data{}
	if err := data.fill(dataData); err != nil {
		return err
	}
	doc.Data[dataKey] = data

	return nil
}

// Header struct definitions

// Represents the header field of a MODE_CTS document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type MODE_CTS_header struct {
	VERSION    validtypes.ValidString `json:"VERSION"`
	MODEL      validtypes.ValidString `json:"MODEL"`
	N_VALID    validtypes.ValidInt    `json:"N_VALID"`
	GRID_RES   validtypes.ValidFloat  `json:"GRID_RES"`
	DESC       validtypes.ValidString `json:"DESC"`
	FCST_VALID validtypes.ValidString `json:"FCST_VALID"`
	FCST_ACCUM validtypes.ValidString `json:"FCST_ACCUM"`
	OBS_LEAD   validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID  validtypes.ValidString `json:"OBS_VALID"`
	OBS_ACCUM  validtypes.ValidString `json:"OBS_ACCUM"`
	FCST_RAD   validtypes.ValidInt    `json:"FCST_RAD"`
	FCST_THR   validtypes.ValidString `json:"FCST_THR"`
	OBS_RAD    validtypes.ValidInt    `json:"OBS_RAD"`
	OBS_THR    validtypes.ValidString `json:"OBS_THR"`
	FCST_VAR   validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV   validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR    validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS  validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV    validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE     validtypes.ValidString `json:"OBTYPE"`
	LINE_TYPE  validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a MODE_OBJ document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type MODE_OBJ_header struct {
	VERSION    validtypes.ValidString `json:"VERSION"`
	MODEL      validtypes.ValidString `json:"MODEL"`
	N_VALID    validtypes.ValidInt    `json:"N_VALID"`
	GRID_RES   validtypes.ValidFloat  `json:"GRID_RES"`
	DESC       validtypes.ValidString `json:"DESC"`
	FCST_VALID validtypes.ValidString `json:"FCST_VALID"`
	FCST_ACCUM validtypes.ValidString `json:"FCST_ACCUM"`
	OBS_LEAD   validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID  validtypes.ValidString `json:"OBS_VALID"`
	OBS_ACCUM  validtypes.ValidString `json:"OBS_ACCUM"`
	FCST_RAD   validtypes.ValidInt    `json:"FCST_RAD"`
	FCST_THR   validtypes.ValidString `json:"FCST_THR"`
	OBS_RAD    validtypes.ValidInt    `json:"OBS_RAD"`
	OBS_THR    validtypes.ValidString `json:"OBS_THR"`
	FCST_VAR   validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV   validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR    validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS  validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV    validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE     validtypes.ValidString `json:"OBTYPE"`
	LINE_TYPE  validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_CNT document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_CNT_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_CTC document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_CTC_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_CTS document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_CTS_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_DMAP document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_DMAP_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_ECLV document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_ECLV_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_ECNT document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_ECNT_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_FHO document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_FHO_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_GENMPR document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_GENMPR_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_GRAD document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_GRAD_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_ISC document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_ISC_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_MCTC document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_MCTC_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_MCTS document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_MCTS_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_MPR document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_MPR_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_NBRCNT document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_NBRCNT_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_NBRCTC document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_NBRCTC_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_NBRCTS document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_NBRCTS_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_ORANK document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_ORANK_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_PCT document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_PCT_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_PHIST document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_PHIST_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_PJC document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_PJC_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_PRC document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_PRC_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_PSTD document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_PSTD_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_RELP document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_RELP_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_RHIST document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_RHIST_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_RPS document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_RPS_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_SAL1L2 document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_SAL1L2_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_SL1L2 document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_SL1L2_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_SSIDX document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_SSIDX_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_SSVAR document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_SSVAR_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_VAL1L2 document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_VAL1L2_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_VCNT document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_VCNT_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a STAT_VL1L2 document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_VL1L2_header struct {
	VERSION        validtypes.ValidString `json:"VERSION"`
	MODEL          validtypes.ValidString `json:"MODEL"`
	DESC           validtypes.ValidString `json:"DESC"`
	FCST_VALID_BEG validtypes.ValidInt    `json:"FCST_VALID_BEG"`
	FCST_VALID_END validtypes.ValidInt    `json:"FCST_VALID_END"`
	OBS_LEAD       validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID_BEG  validtypes.ValidInt    `json:"OBS_VALID_BEG"`
	OBS_VALID_END  validtypes.ValidInt    `json:"OBS_VALID_END"`
	FCST_VAR       validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS     validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV       validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR        validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS      validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV        validtypes.ValidString `json:"OBS_LEV"`
	OBTYPE         validtypes.ValidString `json:"OBTYPE"`
	VX_MASK        validtypes.ValidString `json:"VX_MASK"`
	INTERP_MTHD    validtypes.ValidString `json:"INTERP_MTHD"`
	INTERP_PNTS    validtypes.ValidInt    `json:"INTERP_PNTS"`
	FCST_THRESH    validtypes.ValidString `json:"FCST_THRESH"`
	OBS_THRESH     validtypes.ValidString `json:"OBS_THRESH"`
	COV_THRESH     validtypes.ValidString `json:"COV_THRESH"`
	ALPHA          validtypes.ValidFloat  `json:"ALPHA"`
	LINE_TYPE      validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a TCST_PROBRIRW document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type TCST_PROBRIRW_header struct {
	VERSION    validtypes.ValidString `json:"VERSION"`
	AMODEL     validtypes.ValidString `json:"AMODEL"`
	BMODEL     validtypes.ValidString `json:"BMODEL"`
	DESC       validtypes.ValidString `json:"DESC"`
	STORM_ID   validtypes.ValidString `json:"STORM_ID"`
	BASIN      validtypes.ValidString `json:"BASIN"`
	CYCLONE    validtypes.ValidString `json:"CYCLONE"`
	STORM_NAME validtypes.ValidString `json:"STORM_NAME"`
	VALID      validtypes.ValidInt    `json:"VALID"`
	INIT_MASK  validtypes.ValidString `json:"INIT_MASK"`
	VALID_MASK validtypes.ValidString `json:"VALID_MASK"`
	LINE_TYPE  validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a TCST_TCMPR document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type TCST_TCMPR_header struct {
	VERSION    validtypes.ValidString `json:"VERSION"`
	AMODEL     validtypes.ValidString `json:"AMODEL"`
	BMODEL     validtypes.ValidString `json:"BMODEL"`
	DESC       validtypes.ValidString `json:"DESC"`
	STORM_ID   validtypes.ValidString `json:"STORM_ID"`
	BASIN      validtypes.ValidString `json:"BASIN"`
	CYCLONE    validtypes.ValidString `json:"CYCLONE"`
	STORM_NAME validtypes.ValidString `json:"STORM_NAME"`
	VALID      validtypes.ValidInt    `json:"VALID"`
	INIT_MASK  validtypes.ValidString `json:"INIT_MASK"`
	VALID_MASK validtypes.ValidString `json:"VALID_MASK"`
	LINE_TYPE  validtypes.ValidString `json:"LINE_TYPE"`
}

// fillHeader functions

// Sets MODE_CTS_header struct's fields
func (s *MODE_CTS_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "N_VALID", s.N_VALID.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "GRID_RES", s.GRID_RES.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID", s.FCST_VALID.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FCST_ACCUM", s.FCST_ACCUM.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "OBS_VALID", s.OBS_VALID.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "OBS_ACCUM", s.OBS_ACCUM.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_RAD", s.FCST_RAD.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "FCST_THR", s.FCST_THR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_RAD", s.OBS_RAD.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_THR", s.OBS_THR.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte("MODE_CTS"))) // hardcode the LINE_TYPE
	return errors.Join(errs...)
}

// Sets MODE_OBJ_header struct's fields
func (s *MODE_OBJ_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "N_VALID", s.N_VALID.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "GRID_RES", s.GRID_RES.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID", s.FCST_VALID.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FCST_ACCUM", s.FCST_ACCUM.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "OBS_VALID", s.OBS_VALID.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "OBS_ACCUM", s.OBS_ACCUM.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_RAD", s.FCST_RAD.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "FCST_THR", s.FCST_THR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_RAD", s.OBS_RAD.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_THR", s.OBS_THR.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte("MODE_OBJ"))) // hardcode the LINE_TYPE
	return errors.Join(errs...)
}

// Sets STAT_CNT_header struct's fields
func (s *STAT_CNT_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_CTC_header struct's fields
func (s *STAT_CTC_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_CTS_header struct's fields
func (s *STAT_CTS_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_DMAP_header struct's fields
func (s *STAT_DMAP_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_ECLV_header struct's fields
func (s *STAT_ECLV_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_ECNT_header struct's fields
func (s *STAT_ECNT_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_FHO_header struct's fields
func (s *STAT_FHO_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_GENMPR_header struct's fields
func (s *STAT_GENMPR_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_GRAD_header struct's fields
func (s *STAT_GRAD_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_ISC_header struct's fields
func (s *STAT_ISC_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_MCTC_header struct's fields
func (s *STAT_MCTC_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_MCTS_header struct's fields
func (s *STAT_MCTS_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_MPR_header struct's fields
func (s *STAT_MPR_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_NBRCNT_header struct's fields
func (s *STAT_NBRCNT_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_NBRCTC_header struct's fields
func (s *STAT_NBRCTC_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_NBRCTS_header struct's fields
func (s *STAT_NBRCTS_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_ORANK_header struct's fields
func (s *STAT_ORANK_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_PCT_header struct's fields
func (s *STAT_PCT_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_PHIST_header struct's fields
func (s *STAT_PHIST_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_PJC_header struct's fields
func (s *STAT_PJC_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_PRC_header struct's fields
func (s *STAT_PRC_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_PSTD_header struct's fields
func (s *STAT_PSTD_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_RELP_header struct's fields
func (s *STAT_RELP_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_RHIST_header struct's fields
func (s *STAT_RHIST_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_RPS_header struct's fields
func (s *STAT_RPS_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_SAL1L2_header struct's fields
func (s *STAT_SAL1L2_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_SL1L2_header struct's fields
func (s *STAT_SL1L2_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_SSIDX_header struct's fields
func (s *STAT_SSIDX_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_SSVAR_header struct's fields
func (s *STAT_SSVAR_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_VAL1L2_header struct's fields
func (s *STAT_VAL1L2_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_VCNT_header struct's fields
func (s *STAT_VCNT_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_VL1L2_header struct's fields
func (s *STAT_VL1L2_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "MODEL", s.MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FCST_VALID_BEG", s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FCST_VALID_END", s.FCST_VALID_END.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_LEAD", s.OBS_LEAD.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS_VALID_BEG", s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS_VALID_END", s.OBS_VALID_END.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FCST_VAR", s.FCST_VAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FCST_UNITS", s.FCST_UNITS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FCST_LEV", s.FCST_LEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBS_VAR", s.OBS_VAR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBS_UNITS", s.OBS_UNITS.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBS_LEV", s.OBS_LEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBTYPE", s.OBTYPE.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "VX_MASK", s.VX_MASK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTERP_MTHD", s.INTERP_MTHD.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTERP_PNTS", s.INTERP_PNTS.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FCST_THRESH", s.FCST_THRESH.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OBS_THRESH", s.OBS_THRESH.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "COV_THRESH", s.COV_THRESH.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ALPHA", s.ALPHA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets TCST_PROBRIRW_header struct's fields
func (s *TCST_PROBRIRW_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "AMODEL", s.AMODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "BMODEL", s.BMODEL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "STORM_ID", s.STORM_ID.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "BASIN", s.BASIN.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "CYCLONE", s.CYCLONE.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "STORM_NAME", s.STORM_NAME.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "VALID", s.VALID.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "INIT_MASK", s.INIT_MASK.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "VALID_MASK", s.VALID_MASK.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[13])))
	return errors.Join(errs...)
}

// Sets TCST_TCMPR_header struct's fields
func (s *TCST_TCMPR_header) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "VERSION", s.VERSION.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "AMODEL", s.AMODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "BMODEL", s.BMODEL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "DESC", s.DESC.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "STORM_ID", s.STORM_ID.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "BASIN", s.BASIN.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "CYCLONE", s.CYCLONE.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "STORM_NAME", s.STORM_NAME.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "VALID", s.VALID.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "INIT_MASK", s.INIT_MASK.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "VALID_MASK", s.VALID_MASK.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "LINE_TYPE", s.LINE_TYPE.UnmarshalText([]byte(fields[13])))
	return errors.Join(errs...)
}

// line data struct definitions

type MODE_CTS_data struct {
	FIELD validtypes.ValidString `json:"FIELD,omitzero"`
	TOTAL validtypes.ValidInt    `json:"TOTAL,omitzero"`
	FY_OY validtypes.ValidFloat  `json:"FY_OY,omitzero"`
	FY_ON validtypes.ValidFloat  `json:"FY_ON,omitzero"`
	FN_OY validtypes.ValidFloat  `json:"FN_OY,omitzero"`
	FN_ON validtypes.ValidFloat  `json:"FN_ON,omitzero"`
	BASER validtypes.ValidFloat  `json:"BASER,omitzero"`
	FMEAN validtypes.ValidFloat  `json:"FMEAN,omitzero"`
	ACC   validtypes.ValidFloat  `json:"ACC,omitzero"`
	FBIAS validtypes.ValidFloat  `json:"FBIAS,omitzero"`
	PODY  validtypes.ValidFloat  `json:"PODY,omitzero"`
	PODN  validtypes.ValidFloat  `json:"PODN,omitzero"`
	POFD  validtypes.ValidFloat  `json:"POFD,omitzero"`
	FAR   validtypes.ValidFloat  `json:"FAR,omitzero"`
	CSI   validtypes.ValidFloat  `json:"CSI,omitzero"`
	GSS   validtypes.ValidFloat  `json:"GSS,omitzero"`
	HK    validtypes.ValidFloat  `json:"HK,omitzero"`
	HSS   validtypes.ValidFloat  `json:"HSS,omitzero"`
	ODDS  validtypes.ValidFloat  `json:"ODDS,omitzero"`
}

type MODE_OBJ_data struct {
	OBJECT_ID                  validtypes.ValidString `json:"OBJECT_ID,omitzero"`
	OBJECT_CAT                 validtypes.ValidString `json:"OBJECT_CAT,omitzero"`
	CENTROID_X                 validtypes.ValidFloat  `json:"CENTROID_X,omitzero"`
	CENTROID_Y                 validtypes.ValidFloat  `json:"CENTROID_Y,omitzero"`
	CENTROID_LAT               validtypes.ValidFloat  `json:"CENTROID_LAT,omitzero"`
	CENTROID_LON               validtypes.ValidFloat  `json:"CENTROID_LON,omitzero"`
	AXIS_ANG                   validtypes.ValidFloat  `json:"AXIS_ANG,omitzero"`
	LENGTH                     validtypes.ValidFloat  `json:"LENGTH,omitzero"`
	WIDTH                      validtypes.ValidFloat  `json:"WIDTH,omitzero"`
	AREA                       validtypes.ValidInt    `json:"AREA,omitzero"`
	AREA_THRESH                validtypes.ValidInt    `json:"AREA_THRESH,omitzero"`
	CURVATURE                  validtypes.ValidFloat  `json:"CURVATURE,omitzero"`
	CURVATURE_X                validtypes.ValidFloat  `json:"CURVATURE_X,omitzero"`
	CURVATURE_Y                validtypes.ValidFloat  `json:"CURVATURE_Y,omitzero"`
	COMPLEXITY                 validtypes.ValidFloat  `json:"COMPLEXITY,omitzero"`
	INTENSITY_10               validtypes.ValidFloat  `json:"INTENSITY_10,omitzero"`
	INTENSITY_25               validtypes.ValidFloat  `json:"INTENSITY_25,omitzero"`
	INTENSITY_50               validtypes.ValidFloat  `json:"INTENSITY_50,omitzero"`
	INTENSITY_75               validtypes.ValidFloat  `json:"INTENSITY_75,omitzero"`
	INTENSITY_90               validtypes.ValidFloat  `json:"INTENSITY_90,omitzero"`
	INTENSITY_USER             validtypes.ValidFloat  `json:"INTENSITY_USER,omitzero"`
	INTENSITY_SUM              validtypes.ValidFloat  `json:"INTENSITY_SUM,omitzero"`
	CENTROID_DIST              validtypes.ValidFloat  `json:"CENTROID_DIST,omitzero"`
	BOUNDARY_DIST              validtypes.ValidFloat  `json:"BOUNDARY_DIST,omitzero"`
	CONVEX_HULL_DIST           validtypes.ValidFloat  `json:"CONVEX_HULL_DIST,omitzero"`
	ANGLE_DIFF                 validtypes.ValidFloat  `json:"ANGLE_DIFF,omitzero"`
	ASPECT_DIFF                validtypes.ValidFloat  `json:"ASPECT_DIFF,omitzero"`
	AREA_RATIO                 validtypes.ValidFloat  `json:"AREA_RATIO,omitzero"`
	INTERSECTION_AREA          validtypes.ValidFloat  `json:"INTERSECTION_AREA,omitzero"`
	UNION_AREA                 validtypes.ValidFloat  `json:"UNION_AREA,omitzero"`
	SYMMETRIC_DIFF             validtypes.ValidFloat  `json:"SYMMETRIC_DIFF,omitzero"`
	INTERSECTION_OVER_AREA     validtypes.ValidFloat  `json:"INTERSECTION_OVER_AREA,omitzero"`
	CURVATURE_RATIO            validtypes.ValidFloat  `json:"CURVATURE_RATIO,omitzero"`
	COMPLEXITY_RATIO           validtypes.ValidFloat  `json:"COMPLEXITY_RATIO,omitzero"`
	PERCENTILE_INTENSITY_RATIO validtypes.ValidFloat  `json:"PERCENTILE_INTENSITY_RATIO,omitzero"`
	INTEREST                   validtypes.ValidFloat  `json:"INTEREST,omitzero"`
}

type STAT_CNT_data struct {
	TOTAL                validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FBAR                 validtypes.ValidFloat `json:"FBAR,omitzero"`
	FBAR_NCL             validtypes.ValidFloat `json:"FBAR_NCL,omitzero"`
	FBAR_NCU             validtypes.ValidFloat `json:"FBAR_NCU,omitzero"`
	FBAR_BCL             validtypes.ValidFloat `json:"FBAR_BCL,omitzero"`
	FBAR_BCU             validtypes.ValidFloat `json:"FBAR_BCU,omitzero"`
	FSTDEV               validtypes.ValidFloat `json:"FSTDEV,omitzero"`
	FSTDEV_NCL           validtypes.ValidFloat `json:"FSTDEV_NCL,omitzero"`
	FSTDEV_NCU           validtypes.ValidFloat `json:"FSTDEV_NCU,omitzero"`
	FSTDEV_BCL           validtypes.ValidFloat `json:"FSTDEV_BCL,omitzero"`
	FSTDEV_BCU           validtypes.ValidFloat `json:"FSTDEV_BCU,omitzero"`
	OBAR                 validtypes.ValidFloat `json:"OBAR,omitzero"`
	OBAR_NCL             validtypes.ValidFloat `json:"OBAR_NCL,omitzero"`
	OBAR_NCU             validtypes.ValidFloat `json:"OBAR_NCU,omitzero"`
	OBAR_BCL             validtypes.ValidFloat `json:"OBAR_BCL,omitzero"`
	OBAR_BCU             validtypes.ValidFloat `json:"OBAR_BCU,omitzero"`
	OSTDEV               validtypes.ValidFloat `json:"OSTDEV,omitzero"`
	OSTDEV_NCL           validtypes.ValidFloat `json:"OSTDEV_NCL,omitzero"`
	OSTDEV_NCU           validtypes.ValidFloat `json:"OSTDEV_NCU,omitzero"`
	OSTDEV_BCL           validtypes.ValidFloat `json:"OSTDEV_BCL,omitzero"`
	OSTDEV_BCU           validtypes.ValidFloat `json:"OSTDEV_BCU,omitzero"`
	PR_CORR              validtypes.ValidFloat `json:"PR_CORR,omitzero"`
	PR_CORR_NCL          validtypes.ValidFloat `json:"PR_CORR_NCL,omitzero"`
	PR_CORR_NCU          validtypes.ValidFloat `json:"PR_CORR_NCU,omitzero"`
	PR_CORR_BCL          validtypes.ValidFloat `json:"PR_CORR_BCL,omitzero"`
	PR_CORR_BCU          validtypes.ValidFloat `json:"PR_CORR_BCU,omitzero"`
	SP_CORR              validtypes.ValidFloat `json:"SP_CORR,omitzero"`
	KT_CORR              validtypes.ValidFloat `json:"KT_CORR,omitzero"`
	RANKS                validtypes.ValidInt   `json:"RANKS,omitzero"`
	FRANK_TIES           validtypes.ValidInt   `json:"FRANK_TIES,omitzero"`
	ORANK_TIES           validtypes.ValidInt   `json:"ORANK_TIES,omitzero"`
	ME                   validtypes.ValidFloat `json:"ME,omitzero"`
	ME_NCL               validtypes.ValidFloat `json:"ME_NCL,omitzero"`
	ME_NCU               validtypes.ValidFloat `json:"ME_NCU,omitzero"`
	ME_BCL               validtypes.ValidFloat `json:"ME_BCL,omitzero"`
	ME_BCU               validtypes.ValidFloat `json:"ME_BCU,omitzero"`
	ESTDEV               validtypes.ValidFloat `json:"ESTDEV,omitzero"`
	ESTDEV_NCL           validtypes.ValidFloat `json:"ESTDEV_NCL,omitzero"`
	ESTDEV_NCU           validtypes.ValidFloat `json:"ESTDEV_NCU,omitzero"`
	ESTDEV_BCL           validtypes.ValidFloat `json:"ESTDEV_BCL,omitzero"`
	ESTDEV_BCU           validtypes.ValidFloat `json:"ESTDEV_BCU,omitzero"`
	MBIAS                validtypes.ValidFloat `json:"MBIAS,omitzero"`
	MBIAS_BCL            validtypes.ValidFloat `json:"MBIAS_BCL,omitzero"`
	MBIAS_BCU            validtypes.ValidFloat `json:"MBIAS_BCU,omitzero"`
	MAE                  validtypes.ValidFloat `json:"MAE,omitzero"`
	MAE_BCL              validtypes.ValidFloat `json:"MAE_BCL,omitzero"`
	MAE_BCU              validtypes.ValidFloat `json:"MAE_BCU,omitzero"`
	MSE                  validtypes.ValidFloat `json:"MSE,omitzero"`
	MSE_BCL              validtypes.ValidFloat `json:"MSE_BCL,omitzero"`
	MSE_BCU              validtypes.ValidFloat `json:"MSE_BCU,omitzero"`
	BCMSE                validtypes.ValidFloat `json:"BCMSE,omitzero"`
	BCMSE_BCL            validtypes.ValidFloat `json:"BCMSE_BCL,omitzero"`
	BCMSE_BCU            validtypes.ValidFloat `json:"BCMSE_BCU,omitzero"`
	RMSE                 validtypes.ValidFloat `json:"RMSE,omitzero"`
	RMSE_BCL             validtypes.ValidFloat `json:"RMSE_BCL,omitzero"`
	RMSE_BCU             validtypes.ValidFloat `json:"RMSE_BCU,omitzero"`
	E10                  validtypes.ValidFloat `json:"E10,omitzero"`
	E10_BCL              validtypes.ValidFloat `json:"E10_BCL,omitzero"`
	E10_BCU              validtypes.ValidFloat `json:"E10_BCU,omitzero"`
	E25                  validtypes.ValidFloat `json:"E25,omitzero"`
	E25_BCL              validtypes.ValidFloat `json:"E25_BCL,omitzero"`
	E25_BCU              validtypes.ValidFloat `json:"E25_BCU,omitzero"`
	E50                  validtypes.ValidFloat `json:"E50,omitzero"`
	E50_BCL              validtypes.ValidFloat `json:"E50_BCL,omitzero"`
	E50_BCU              validtypes.ValidFloat `json:"E50_BCU,omitzero"`
	E75                  validtypes.ValidFloat `json:"E75,omitzero"`
	E75_BCL              validtypes.ValidFloat `json:"E75_BCL,omitzero"`
	E75_BCU              validtypes.ValidFloat `json:"E75_BCU,omitzero"`
	E90                  validtypes.ValidFloat `json:"E90,omitzero"`
	E90_BCL              validtypes.ValidFloat `json:"E90_BCL,omitzero"`
	E90_BCU              validtypes.ValidFloat `json:"E90_BCU,omitzero"`
	EIQR                 validtypes.ValidFloat `json:"EIQR,omitzero"`
	EIQR_BCL             validtypes.ValidFloat `json:"EIQR_BCL,omitzero"`
	EIQR_BCU             validtypes.ValidFloat `json:"EIQR_BCU,omitzero"`
	MAD                  validtypes.ValidFloat `json:"MAD,omitzero"`
	MAD_BCL              validtypes.ValidFloat `json:"MAD_BCL,omitzero"`
	MAD_BCU              validtypes.ValidFloat `json:"MAD_BCU,omitzero"`
	ANOM_CORR            validtypes.ValidFloat `json:"ANOM_CORR,omitzero"`
	ANOM_CORR_NCL        validtypes.ValidFloat `json:"ANOM_CORR_NCL,omitzero"`
	ANOM_CORR_NCU        validtypes.ValidFloat `json:"ANOM_CORR_NCU,omitzero"`
	ANOM_CORR_BCL        validtypes.ValidFloat `json:"ANOM_CORR_BCL,omitzero"`
	ANOM_CORR_BCU        validtypes.ValidFloat `json:"ANOM_CORR_BCU,omitzero"`
	ME2                  validtypes.ValidFloat `json:"ME2,omitzero"`
	ME2_BCL              validtypes.ValidFloat `json:"ME2_BCL,omitzero"`
	ME2_BCU              validtypes.ValidFloat `json:"ME2_BCU,omitzero"`
	MSESS                validtypes.ValidFloat `json:"MSESS,omitzero"`
	MSESS_BCL            validtypes.ValidFloat `json:"MSESS_BCL,omitzero"`
	MSESS_BCU            validtypes.ValidFloat `json:"MSESS_BCU,omitzero"`
	RMSFA                validtypes.ValidFloat `json:"RMSFA,omitzero"`
	RMSFA_BCL            validtypes.ValidFloat `json:"RMSFA_BCL,omitzero"`
	RMSFA_BCU            validtypes.ValidFloat `json:"RMSFA_BCU,omitzero"`
	RMSOA                validtypes.ValidFloat `json:"RMSOA,omitzero"`
	RMSOA_BCL            validtypes.ValidFloat `json:"RMSOA_BCL,omitzero"`
	RMSOA_BCU            validtypes.ValidFloat `json:"RMSOA_BCU,omitzero"`
	ANOM_CORR_UNCNTR     validtypes.ValidFloat `json:"ANOM_CORR_UNCNTR,omitzero"`
	ANOM_CORR_UNCNTR_BCL validtypes.ValidFloat `json:"ANOM_CORR_UNCNTR_BCL,omitzero"`
	ANOM_CORR_UNCNTR_BCU validtypes.ValidFloat `json:"ANOM_CORR_UNCNTR_BCU,omitzero"`
	SI                   validtypes.ValidFloat `json:"SI,omitzero"`
	SI_BCL               validtypes.ValidFloat `json:"SI_BCL,omitzero"`
	SI_BCU               validtypes.ValidFloat `json:"SI_BCU,omitzero"`
}

type STAT_CTC_data struct {
	TOTAL validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FY_OY validtypes.ValidFloat `json:"FY_OY,omitzero"`
	FY_ON validtypes.ValidFloat `json:"FY_ON,omitzero"`
	FN_OY validtypes.ValidFloat `json:"FN_OY,omitzero"`
	FN_ON validtypes.ValidFloat `json:"FN_ON,omitzero"`
}

type STAT_CTS_data struct {
	TOTAL     validtypes.ValidInt   `json:"TOTAL,omitzero"`
	BASER     validtypes.ValidFloat `json:"BASER,omitzero"`
	BASER_NCL validtypes.ValidFloat `json:"BASER_NCL,omitzero"`
	BASER_NCU validtypes.ValidFloat `json:"BASER_NCU,omitzero"`
	BASER_BCL validtypes.ValidFloat `json:"BASER_BCL,omitzero"`
	BASER_BCU validtypes.ValidFloat `json:"BASER_BCU,omitzero"`
	FMEAN     validtypes.ValidFloat `json:"FMEAN,omitzero"`
	FMEAN_NCL validtypes.ValidFloat `json:"FMEAN_NCL,omitzero"`
	FMEAN_NCU validtypes.ValidFloat `json:"FMEAN_NCU,omitzero"`
	FMEAN_BCL validtypes.ValidFloat `json:"FMEAN_BCL,omitzero"`
	FMEAN_BCU validtypes.ValidFloat `json:"FMEAN_BCU,omitzero"`
	ACC       validtypes.ValidFloat `json:"ACC,omitzero"`
	ACC_NCL   validtypes.ValidFloat `json:"ACC_NCL,omitzero"`
	ACC_NCU   validtypes.ValidFloat `json:"ACC_NCU,omitzero"`
	ACC_BCL   validtypes.ValidFloat `json:"ACC_BCL,omitzero"`
	ACC_BCU   validtypes.ValidFloat `json:"ACC_BCU,omitzero"`
	FBIAS     validtypes.ValidFloat `json:"FBIAS,omitzero"`
	FBIAS_BCL validtypes.ValidFloat `json:"FBIAS_BCL,omitzero"`
	FBIAS_BCU validtypes.ValidFloat `json:"FBIAS_BCU,omitzero"`
	PODY      validtypes.ValidFloat `json:"PODY,omitzero"`
	PODY_NCL  validtypes.ValidFloat `json:"PODY_NCL,omitzero"`
	PODY_NCU  validtypes.ValidFloat `json:"PODY_NCU,omitzero"`
	PODY_BCL  validtypes.ValidFloat `json:"PODY_BCL,omitzero"`
	PODY_BCU  validtypes.ValidFloat `json:"PODY_BCU,omitzero"`
	PODN      validtypes.ValidFloat `json:"PODN,omitzero"`
	PODN_NCL  validtypes.ValidFloat `json:"PODN_NCL,omitzero"`
	PODN_NCU  validtypes.ValidFloat `json:"PODN_NCU,omitzero"`
	PODN_BCL  validtypes.ValidFloat `json:"PODN_BCL,omitzero"`
	PODN_BCU  validtypes.ValidFloat `json:"PODN_BCU,omitzero"`
	POFD      validtypes.ValidFloat `json:"POFD,omitzero"`
	POFD_NCL  validtypes.ValidFloat `json:"POFD_NCL,omitzero"`
	POFD_NCU  validtypes.ValidFloat `json:"POFD_NCU,omitzero"`
	POFD_BCL  validtypes.ValidFloat `json:"POFD_BCL,omitzero"`
	POFD_BCU  validtypes.ValidFloat `json:"POFD_BCU,omitzero"`
	FAR       validtypes.ValidFloat `json:"FAR,omitzero"`
	FAR_NCL   validtypes.ValidFloat `json:"FAR_NCL,omitzero"`
	FAR_NCU   validtypes.ValidFloat `json:"FAR_NCU,omitzero"`
	FAR_BCL   validtypes.ValidFloat `json:"FAR_BCL,omitzero"`
	FAR_BCU   validtypes.ValidFloat `json:"FAR_BCU,omitzero"`
	CSI       validtypes.ValidFloat `json:"CSI,omitzero"`
	CSI_NCL   validtypes.ValidFloat `json:"CSI_NCL,omitzero"`
	CSI_NCU   validtypes.ValidFloat `json:"CSI_NCU,omitzero"`
	CSI_BCL   validtypes.ValidFloat `json:"CSI_BCL,omitzero"`
	CSI_BCU   validtypes.ValidFloat `json:"CSI_BCU,omitzero"`
	GSS       validtypes.ValidFloat `json:"GSS,omitzero"`
	GSS_BCL   validtypes.ValidFloat `json:"GSS_BCL,omitzero"`
	GSS_BCU   validtypes.ValidFloat `json:"GSS_BCU,omitzero"`
	HK        validtypes.ValidFloat `json:"HK,omitzero"`
	HK_NCL    validtypes.ValidFloat `json:"HK_NCL,omitzero"`
	HK_NCU    validtypes.ValidFloat `json:"HK_NCU,omitzero"`
	HK_BCL    validtypes.ValidFloat `json:"HK_BCL,omitzero"`
	HK_BCU    validtypes.ValidFloat `json:"HK_BCU,omitzero"`
	HSS       validtypes.ValidFloat `json:"HSS,omitzero"`
	HSS_BCL   validtypes.ValidFloat `json:"HSS_BCL,omitzero"`
	HSS_BCU   validtypes.ValidFloat `json:"HSS_BCU,omitzero"`
	ODDS      validtypes.ValidFloat `json:"ODDS,omitzero"`
	ODDS_NCL  validtypes.ValidFloat `json:"ODDS_NCL,omitzero"`
	ODDS_NCU  validtypes.ValidFloat `json:"ODDS_NCU,omitzero"`
	ODDS_BCL  validtypes.ValidFloat `json:"ODDS_BCL,omitzero"`
	ODDS_BCU  validtypes.ValidFloat `json:"ODDS_BCU,omitzero"`
	LODDS     validtypes.ValidFloat `json:"LODDS,omitzero"`
	LODDS_NCL validtypes.ValidFloat `json:"LODDS_NCL,omitzero"`
	LODDS_NCU validtypes.ValidFloat `json:"LODDS_NCU,omitzero"`
	LODDS_BCL validtypes.ValidFloat `json:"LODDS_BCL,omitzero"`
	LODDS_BCU validtypes.ValidFloat `json:"LODDS_BCU,omitzero"`
	ORSS      validtypes.ValidFloat `json:"ORSS,omitzero"`
	ORSS_NCL  validtypes.ValidFloat `json:"ORSS_NCL,omitzero"`
	ORSS_NCU  validtypes.ValidFloat `json:"ORSS_NCU,omitzero"`
	ORSS_BCL  validtypes.ValidFloat `json:"ORSS_BCL,omitzero"`
	ORSS_BCU  validtypes.ValidFloat `json:"ORSS_BCU,omitzero"`
	EDS       validtypes.ValidFloat `json:"EDS,omitzero"`
	EDS_NCL   validtypes.ValidFloat `json:"EDS_NCL,omitzero"`
	EDS_NCU   validtypes.ValidFloat `json:"EDS_NCU,omitzero"`
	EDS_BCL   validtypes.ValidFloat `json:"EDS_BCL,omitzero"`
	EDS_BCU   validtypes.ValidFloat `json:"EDS_BCU,omitzero"`
	SEDS      validtypes.ValidFloat `json:"SEDS,omitzero"`
	SEDS_NCL  validtypes.ValidFloat `json:"SEDS_NCL,omitzero"`
	SEDS_NCU  validtypes.ValidFloat `json:"SEDS_NCU,omitzero"`
	SEDS_BCL  validtypes.ValidFloat `json:"SEDS_BCL,omitzero"`
	SEDS_BCU  validtypes.ValidFloat `json:"SEDS_BCU,omitzero"`
	EDI       validtypes.ValidFloat `json:"EDI,omitzero"`
	EDI_NCL   validtypes.ValidFloat `json:"EDI_NCL,omitzero"`
	EDI_NCU   validtypes.ValidFloat `json:"EDI_NCU,omitzero"`
	EDI_BCL   validtypes.ValidFloat `json:"EDI_BCL,omitzero"`
	EDI_BCU   validtypes.ValidFloat `json:"EDI_BCU,omitzero"`
	SEDI      validtypes.ValidFloat `json:"SEDI,omitzero"`
	SEDI_NCL  validtypes.ValidFloat `json:"SEDI_NCL,omitzero"`
	SEDI_NCU  validtypes.ValidFloat `json:"SEDI_NCU,omitzero"`
	SEDI_BCL  validtypes.ValidFloat `json:"SEDI_BCL,omitzero"`
	SEDI_BCU  validtypes.ValidFloat `json:"SEDI_BCU,omitzero"`
	BAGSS     validtypes.ValidFloat `json:"BAGSS,omitzero"`
	BAGSS_BCL validtypes.ValidFloat `json:"BAGSS_BCL,omitzero"`
	BAGSS_BCU validtypes.ValidFloat `json:"BAGSS_BCU,omitzero"`
}

type STAT_DMAP_data struct {
	TOTAL      validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FY         validtypes.ValidInt   `json:"FY,omitzero"`
	OY         validtypes.ValidInt   `json:"OY,omitzero"`
	FBIAS      validtypes.ValidFloat `json:"FBIAS,omitzero"`
	BADDELEY   validtypes.ValidFloat `json:"BADDELEY,omitzero"`
	HAUSDORFF  validtypes.ValidFloat `json:"HAUSDORFF,omitzero"`
	MED_FO     validtypes.ValidFloat `json:"MED_FO,omitzero"`
	MED_OF     validtypes.ValidFloat `json:"MED_OF,omitzero"`
	MED_MIN    validtypes.ValidFloat `json:"MED_MIN,omitzero"`
	MED_MAX    validtypes.ValidFloat `json:"MED_MAX,omitzero"`
	MED_MEAN   validtypes.ValidFloat `json:"MED_MEAN,omitzero"`
	FOM_FO     validtypes.ValidFloat `json:"FOM_FO,omitzero"`
	FOM_OF     validtypes.ValidFloat `json:"FOM_OF,omitzero"`
	FOM_MIN    validtypes.ValidFloat `json:"FOM_MIN,omitzero"`
	FOM_MAX    validtypes.ValidFloat `json:"FOM_MAX,omitzero"`
	FOM_MEAN   validtypes.ValidFloat `json:"FOM_MEAN,omitzero"`
	ZHU_FO     validtypes.ValidFloat `json:"ZHU_FO,omitzero"`
	ZHU_OF     validtypes.ValidFloat `json:"ZHU_OF,omitzero"`
	ZHU_MIN    validtypes.ValidFloat `json:"ZHU_MIN,omitzero"`
	ZHU_MAX    validtypes.ValidFloat `json:"ZHU_MAX,omitzero"`
	ZHU_MEAN   validtypes.ValidFloat `json:"ZHU_MEAN,omitzero"`
	G          validtypes.ValidFloat `json:"G,omitzero"`
	GBETA      validtypes.ValidFloat `json:"GBETA,omitzero"`
	BETA_VALUE validtypes.ValidFloat `json:"BETA_VALUE,omitzero"`
}

type STAT_ECLV_data struct {
	TOTAL       validtypes.ValidInt    `json:"TOTAL,omitzero"`
	BASER       validtypes.ValidFloat  `json:"BASER,omitzero"`
	VALUE_BASER validtypes.ValidInt    `json:"VALUE_BASER,omitzero"`
	PTS         map[string]interface{} `json:"PTS,omitzero"`
}

type STAT_ECNT_data struct {
	TOTAL            validtypes.ValidInt   `json:"TOTAL,omitzero"`
	N_ENS            validtypes.ValidInt   `json:"N_ENS,omitzero"`
	CRPS             validtypes.ValidFloat `json:"CRPS,omitzero"`
	CRPSS            validtypes.ValidFloat `json:"CRPSS,omitzero"`
	IGN              validtypes.ValidFloat `json:"IGN,omitzero"`
	ME               validtypes.ValidFloat `json:"ME,omitzero"`
	RMSE             validtypes.ValidFloat `json:"RMSE,omitzero"`
	SPREAD           validtypes.ValidFloat `json:"SPREAD,omitzero"`
	ME_OERR          validtypes.ValidFloat `json:"ME_OERR,omitzero"`
	RMSE_OERR        validtypes.ValidFloat `json:"RMSE_OERR,omitzero"`
	SPREAD_OERR      validtypes.ValidFloat `json:"SPREAD_OERR,omitzero"`
	SPREAD_PLUS_OERR validtypes.ValidFloat `json:"SPREAD_PLUS_OERR,omitzero"`
	CRPSCL           validtypes.ValidFloat `json:"CRPSCL,omitzero"`
	CRPS_EMP         validtypes.ValidFloat `json:"CRPS_EMP,omitzero"`
	CRPSCL_EMP       validtypes.ValidFloat `json:"CRPSCL_EMP,omitzero"`
	CRPSS_EMP        validtypes.ValidFloat `json:"CRPSS_EMP,omitzero"`
}

type STAT_FHO_data struct {
	TOTAL  validtypes.ValidInt   `json:"TOTAL,omitzero"`
	F_RATE validtypes.ValidFloat `json:"F_RATE,omitzero"`
	H_RATE validtypes.ValidFloat `json:"H_RATE,omitzero"`
	O_RATE validtypes.ValidFloat `json:"O_RATE,omitzero"`
}

type STAT_GENMPR_data struct {
	TOTAL      validtypes.ValidInt    `json:"TOTAL,omitzero"`
	INDEX      validtypes.ValidInt    `json:"INDEX,omitzero"`
	STORM_ID   validtypes.ValidString `json:"STORM_ID,omitzero"`
	PROB_LEAD  validtypes.ValidFloat  `json:"PROB_LEAD,omitzero"`
	PROB_VAL   validtypes.ValidFloat  `json:"PROB_VAL,omitzero"`
	AGEN_INIT  validtypes.ValidString `json:"AGEN_INIT,omitzero"`
	AGEN_FHR   validtypes.ValidString `json:"AGEN_FHR,omitzero"`
	AGEN_LAT   validtypes.ValidFloat  `json:"AGEN_LAT,omitzero"`
	AGEN_LON   validtypes.ValidFloat  `json:"AGEN_LON,omitzero"`
	AGEN_DLAND validtypes.ValidFloat  `json:"AGEN_DLAND,omitzero"`
	BGEN_LAT   validtypes.ValidFloat  `json:"BGEN_LAT,omitzero"`
	BGEN_LON   validtypes.ValidFloat  `json:"BGEN_LON,omitzero"`
	BGEN_DLAND validtypes.ValidFloat  `json:"BGEN_DLAND,omitzero"`
	GEN_DIST   validtypes.ValidFloat  `json:"GEN_DIST,omitzero"`
	GEN_TDIFF  validtypes.ValidString `json:"GEN_TDIFF,omitzero"`
	INIT_TDIFF validtypes.ValidString `json:"INIT_TDIFF,omitzero"`
	DEV_CAT    validtypes.ValidString `json:"DEV_CAT,omitzero"`
	OPS_CAT    validtypes.ValidString `json:"OPS_CAT,omitzero"`
}

type STAT_GRAD_data struct {
	TOTAL      validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FGBAR      validtypes.ValidFloat `json:"FGBAR,omitzero"`
	OGBAR      validtypes.ValidFloat `json:"OGBAR,omitzero"`
	MGBAR      validtypes.ValidFloat `json:"MGBAR,omitzero"`
	EGBAR      validtypes.ValidFloat `json:"EGBAR,omitzero"`
	S1         validtypes.ValidFloat `json:"S1,omitzero"`
	S1_OG      validtypes.ValidFloat `json:"S1_OG,omitzero"`
	FGOG_RATIO validtypes.ValidFloat `json:"FGOG_RATIO,omitzero"`
	DX         validtypes.ValidFloat `json:"DX,omitzero"`
	DY         validtypes.ValidFloat `json:"DY,omitzero"`
}

type STAT_ISC_data struct {
	TOTAL    validtypes.ValidInt   `json:"TOTAL,omitzero"`
	TILE_DIM validtypes.ValidInt   `json:"TILE_DIM,omitzero"`
	TILE_XLL validtypes.ValidInt   `json:"TILE_XLL,omitzero"`
	TILE_YLL validtypes.ValidInt   `json:"TILE_YLL,omitzero"`
	NSCALE   validtypes.ValidInt   `json:"NSCALE,omitzero"`
	ISCALE   validtypes.ValidInt   `json:"ISCALE,omitzero"`
	MSE      validtypes.ValidFloat `json:"MSE,omitzero"`
	ISC      validtypes.ValidFloat `json:"ISC,omitzero"`
	FENERGY2 validtypes.ValidFloat `json:"FENERGY2,omitzero"`
	OENERGY2 validtypes.ValidFloat `json:"OENERGY2,omitzero"`
	BASER    validtypes.ValidFloat `json:"BASER,omitzero"`
	FBIAS    validtypes.ValidFloat `json:"FBIAS,omitzero"`
}

type STAT_MCTC_data struct {
	TOTAL    validtypes.ValidInt    `json:"TOTAL,omitzero"`
	CAT      map[string]interface{} `json:"CAT,omitzero"`
	EC_VALUE validtypes.ValidFloat  `json:"EC_VALUE,omitzero"`
}

type STAT_MCTS_data struct {
	TOTAL      validtypes.ValidInt   `json:"TOTAL,omitzero"`
	N_CAT      validtypes.ValidInt   `json:"N_CAT,omitzero"`
	ACC        validtypes.ValidFloat `json:"ACC,omitzero"`
	ACC_NCL    validtypes.ValidFloat `json:"ACC_NCL,omitzero"`
	ACC_NCU    validtypes.ValidFloat `json:"ACC_NCU,omitzero"`
	ACC_BCL    validtypes.ValidFloat `json:"ACC_BCL,omitzero"`
	ACC_BCU    validtypes.ValidFloat `json:"ACC_BCU,omitzero"`
	HK         validtypes.ValidFloat `json:"HK,omitzero"`
	HK_BCL     validtypes.ValidFloat `json:"HK_BCL,omitzero"`
	HK_BCU     validtypes.ValidFloat `json:"HK_BCU,omitzero"`
	HSS        validtypes.ValidFloat `json:"HSS,omitzero"`
	HSS_BCL    validtypes.ValidFloat `json:"HSS_BCL,omitzero"`
	HSS_BCU    validtypes.ValidFloat `json:"HSS_BCU,omitzero"`
	GER        validtypes.ValidFloat `json:"GER,omitzero"`
	GER_BCL    validtypes.ValidFloat `json:"GER_BCL,omitzero"`
	GER_BCU    validtypes.ValidFloat `json:"GER_BCU,omitzero"`
	HSS_EC     validtypes.ValidFloat `json:"HSS_EC,omitzero"`
	HSS_EC_BCL validtypes.ValidFloat `json:"HSS_EC_BCL,omitzero"`
	HSS_EC_BCU validtypes.ValidFloat `json:"HSS_EC_BCU,omitzero"`
	EC_VALUE   validtypes.ValidFloat `json:"EC_VALUE,omitzero"`
}

type STAT_MPR_data struct {
	TOTAL       validtypes.ValidInt    `json:"TOTAL,omitzero"`
	INDEX       validtypes.ValidInt    `json:"INDEX,omitzero"`
	OBS_SID     validtypes.ValidString `json:"OBS_SID,omitzero"`
	OBS_LAT     validtypes.ValidFloat  `json:"OBS_LAT,omitzero"`
	OBS_LON     validtypes.ValidFloat  `json:"OBS_LON,omitzero"`
	OBS_LVL     validtypes.ValidFloat  `json:"OBS_LVL,omitzero"`
	OBS_ELV     validtypes.ValidFloat  `json:"OBS_ELV,omitzero"`
	FCST        validtypes.ValidFloat  `json:"FCST,omitzero"`
	OBS         validtypes.ValidFloat  `json:"OBS,omitzero"`
	OBS_QC      validtypes.ValidString `json:"OBS_QC,omitzero"`
	CLIMO_MEAN  validtypes.ValidFloat  `json:"CLIMO_MEAN,omitzero"`
	CLIMO_STDEV validtypes.ValidFloat  `json:"CLIMO_STDEV,omitzero"`
	CLIMO_CDF   validtypes.ValidFloat  `json:"CLIMO_CDF,omitzero"`
}

type STAT_NBRCNT_data struct {
	TOTAL      validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FBS        validtypes.ValidFloat `json:"FBS,omitzero"`
	FBS_BCL    validtypes.ValidFloat `json:"FBS_BCL,omitzero"`
	FBS_BCU    validtypes.ValidFloat `json:"FBS_BCU,omitzero"`
	FSS        validtypes.ValidFloat `json:"FSS,omitzero"`
	FSS_BCL    validtypes.ValidFloat `json:"FSS_BCL,omitzero"`
	FSS_BCU    validtypes.ValidFloat `json:"FSS_BCU,omitzero"`
	AFSS       validtypes.ValidFloat `json:"AFSS,omitzero"`
	AFSS_BCL   validtypes.ValidFloat `json:"AFSS_BCL,omitzero"`
	AFSS_BCU   validtypes.ValidFloat `json:"AFSS_BCU,omitzero"`
	UFSS       validtypes.ValidFloat `json:"UFSS,omitzero"`
	UFSS_BCL   validtypes.ValidFloat `json:"UFSS_BCL,omitzero"`
	UFSS_BCU   validtypes.ValidFloat `json:"UFSS_BCU,omitzero"`
	F_RATE     validtypes.ValidFloat `json:"F_RATE,omitzero"`
	F_RATE_BCL validtypes.ValidFloat `json:"F_RATE_BCL,omitzero"`
	F_RATE_BCU validtypes.ValidFloat `json:"F_RATE_BCU,omitzero"`
	O_RATE     validtypes.ValidFloat `json:"O_RATE,omitzero"`
	O_RATE_BCL validtypes.ValidFloat `json:"O_RATE_BCL,omitzero"`
	O_RATE_BCU validtypes.ValidFloat `json:"O_RATE_BCU,omitzero"`
}

type STAT_NBRCTC_data struct {
	TOTAL validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FY_OY validtypes.ValidFloat `json:"FY_OY,omitzero"`
	FY_ON validtypes.ValidFloat `json:"FY_ON,omitzero"`
	FN_OY validtypes.ValidFloat `json:"FN_OY,omitzero"`
	FN_ON validtypes.ValidFloat `json:"FN_ON,omitzero"`
}

type STAT_NBRCTS_data struct {
	TOTAL     validtypes.ValidInt   `json:"TOTAL,omitzero"`
	BASER     validtypes.ValidFloat `json:"BASER,omitzero"`
	BASER_NCL validtypes.ValidFloat `json:"BASER_NCL,omitzero"`
	BASER_NCU validtypes.ValidFloat `json:"BASER_NCU,omitzero"`
	BASER_BCL validtypes.ValidFloat `json:"BASER_BCL,omitzero"`
	BASER_BCU validtypes.ValidFloat `json:"BASER_BCU,omitzero"`
	FMEAN     validtypes.ValidFloat `json:"FMEAN,omitzero"`
	FMEAN_NCL validtypes.ValidFloat `json:"FMEAN_NCL,omitzero"`
	FMEAN_NCU validtypes.ValidFloat `json:"FMEAN_NCU,omitzero"`
	FMEAN_BCL validtypes.ValidFloat `json:"FMEAN_BCL,omitzero"`
	FMEAN_BCU validtypes.ValidFloat `json:"FMEAN_BCU,omitzero"`
	ACC       validtypes.ValidFloat `json:"ACC,omitzero"`
	ACC_NCL   validtypes.ValidFloat `json:"ACC_NCL,omitzero"`
	ACC_NCU   validtypes.ValidFloat `json:"ACC_NCU,omitzero"`
	ACC_BCL   validtypes.ValidFloat `json:"ACC_BCL,omitzero"`
	ACC_BCU   validtypes.ValidFloat `json:"ACC_BCU,omitzero"`
	FBIAS     validtypes.ValidFloat `json:"FBIAS,omitzero"`
	FBIAS_BCL validtypes.ValidFloat `json:"FBIAS_BCL,omitzero"`
	FBIAS_BCU validtypes.ValidFloat `json:"FBIAS_BCU,omitzero"`
	PODY      validtypes.ValidFloat `json:"PODY,omitzero"`
	PODY_NCL  validtypes.ValidFloat `json:"PODY_NCL,omitzero"`
	PODY_NCU  validtypes.ValidFloat `json:"PODY_NCU,omitzero"`
	PODY_BCL  validtypes.ValidFloat `json:"PODY_BCL,omitzero"`
	PODY_BCU  validtypes.ValidFloat `json:"PODY_BCU,omitzero"`
	PODN      validtypes.ValidFloat `json:"PODN,omitzero"`
	PODN_NCL  validtypes.ValidFloat `json:"PODN_NCL,omitzero"`
	PODN_NCU  validtypes.ValidFloat `json:"PODN_NCU,omitzero"`
	PODN_BCL  validtypes.ValidFloat `json:"PODN_BCL,omitzero"`
	PODN_BCU  validtypes.ValidFloat `json:"PODN_BCU,omitzero"`
	POFD      validtypes.ValidFloat `json:"POFD,omitzero"`
	POFD_NCL  validtypes.ValidFloat `json:"POFD_NCL,omitzero"`
	POFD_NCU  validtypes.ValidFloat `json:"POFD_NCU,omitzero"`
	POFD_BCL  validtypes.ValidFloat `json:"POFD_BCL,omitzero"`
	POFD_BCU  validtypes.ValidFloat `json:"POFD_BCU,omitzero"`
	FAR       validtypes.ValidFloat `json:"FAR,omitzero"`
	FAR_NCL   validtypes.ValidFloat `json:"FAR_NCL,omitzero"`
	FAR_NCU   validtypes.ValidFloat `json:"FAR_NCU,omitzero"`
	FAR_BCL   validtypes.ValidFloat `json:"FAR_BCL,omitzero"`
	FAR_BCU   validtypes.ValidFloat `json:"FAR_BCU,omitzero"`
	CSI       validtypes.ValidFloat `json:"CSI,omitzero"`
	CSI_NCL   validtypes.ValidFloat `json:"CSI_NCL,omitzero"`
	CSI_NCU   validtypes.ValidFloat `json:"CSI_NCU,omitzero"`
	CSI_BCL   validtypes.ValidFloat `json:"CSI_BCL,omitzero"`
	CSI_BCU   validtypes.ValidFloat `json:"CSI_BCU,omitzero"`
	GSS       validtypes.ValidFloat `json:"GSS,omitzero"`
	GSS_BCL   validtypes.ValidFloat `json:"GSS_BCL,omitzero"`
	GSS_BCU   validtypes.ValidFloat `json:"GSS_BCU,omitzero"`
	HK        validtypes.ValidFloat `json:"HK,omitzero"`
	HK_NCL    validtypes.ValidFloat `json:"HK_NCL,omitzero"`
	HK_NCU    validtypes.ValidFloat `json:"HK_NCU,omitzero"`
	HK_BCL    validtypes.ValidFloat `json:"HK_BCL,omitzero"`
	HK_BCU    validtypes.ValidFloat `json:"HK_BCU,omitzero"`
	HSS       validtypes.ValidFloat `json:"HSS,omitzero"`
	HSS_BCL   validtypes.ValidFloat `json:"HSS_BCL,omitzero"`
	HSS_BCU   validtypes.ValidFloat `json:"HSS_BCU,omitzero"`
	ODDS      validtypes.ValidFloat `json:"ODDS,omitzero"`
	ODDS_NCL  validtypes.ValidFloat `json:"ODDS_NCL,omitzero"`
	ODDS_NCU  validtypes.ValidFloat `json:"ODDS_NCU,omitzero"`
	ODDS_BCL  validtypes.ValidFloat `json:"ODDS_BCL,omitzero"`
	ODDS_BCU  validtypes.ValidFloat `json:"ODDS_BCU,omitzero"`
	LODDS     validtypes.ValidFloat `json:"LODDS,omitzero"`
	LODDS_NCL validtypes.ValidFloat `json:"LODDS_NCL,omitzero"`
	LODDS_NCU validtypes.ValidFloat `json:"LODDS_NCU,omitzero"`
	LODDS_BCL validtypes.ValidFloat `json:"LODDS_BCL,omitzero"`
	LODDS_BCU validtypes.ValidFloat `json:"LODDS_BCU,omitzero"`
	ORSS      validtypes.ValidFloat `json:"ORSS,omitzero"`
	ORSS_NCL  validtypes.ValidFloat `json:"ORSS_NCL,omitzero"`
	ORSS_NCU  validtypes.ValidFloat `json:"ORSS_NCU,omitzero"`
	ORSS_BCL  validtypes.ValidFloat `json:"ORSS_BCL,omitzero"`
	ORSS_BCU  validtypes.ValidFloat `json:"ORSS_BCU,omitzero"`
	EDS       validtypes.ValidFloat `json:"EDS,omitzero"`
	EDS_NCL   validtypes.ValidFloat `json:"EDS_NCL,omitzero"`
	EDS_NCU   validtypes.ValidFloat `json:"EDS_NCU,omitzero"`
	EDS_BCL   validtypes.ValidFloat `json:"EDS_BCL,omitzero"`
	EDS_BCU   validtypes.ValidFloat `json:"EDS_BCU,omitzero"`
	SEDS      validtypes.ValidFloat `json:"SEDS,omitzero"`
	SEDS_NCL  validtypes.ValidFloat `json:"SEDS_NCL,omitzero"`
	SEDS_NCU  validtypes.ValidFloat `json:"SEDS_NCU,omitzero"`
	SEDS_BCL  validtypes.ValidFloat `json:"SEDS_BCL,omitzero"`
	SEDS_BCU  validtypes.ValidFloat `json:"SEDS_BCU,omitzero"`
	EDI       validtypes.ValidFloat `json:"EDI,omitzero"`
	EDI_NCL   validtypes.ValidFloat `json:"EDI_NCL,omitzero"`
	EDI_NCU   validtypes.ValidFloat `json:"EDI_NCU,omitzero"`
	EDI_BCL   validtypes.ValidFloat `json:"EDI_BCL,omitzero"`
	EDI_BCU   validtypes.ValidFloat `json:"EDI_BCU,omitzero"`
	SEDI      validtypes.ValidFloat `json:"SEDI,omitzero"`
	SEDI_NCL  validtypes.ValidFloat `json:"SEDI_NCL,omitzero"`
	SEDI_NCU  validtypes.ValidFloat `json:"SEDI_NCU,omitzero"`
	SEDI_BCL  validtypes.ValidFloat `json:"SEDI_BCL,omitzero"`
	SEDI_BCU  validtypes.ValidFloat `json:"SEDI_BCU,omitzero"`
	BAGSS     validtypes.ValidFloat `json:"BAGSS,omitzero"`
	BAGSS_BCL validtypes.ValidFloat `json:"BAGSS_BCL,omitzero"`
	BAGSS_BCU validtypes.ValidFloat `json:"BAGSS_BCU,omitzero"`
}

type STAT_ORANK_data struct {
	TOTAL            validtypes.ValidInt    `json:"TOTAL,omitzero"`
	INDEX            validtypes.ValidInt    `json:"INDEX,omitzero"`
	OBS_SID          validtypes.ValidString `json:"OBS_SID,omitzero"`
	OBS_LAT          validtypes.ValidFloat  `json:"OBS_LAT,omitzero"`
	OBS_LON          validtypes.ValidFloat  `json:"OBS_LON,omitzero"`
	OBS_LVL          validtypes.ValidFloat  `json:"OBS_LVL,omitzero"`
	OBS_ELV          validtypes.ValidFloat  `json:"OBS_ELV,omitzero"`
	OBS              validtypes.ValidFloat  `json:"OBS,omitzero"`
	PIT              validtypes.ValidFloat  `json:"PIT,omitzero"`
	RANK             validtypes.ValidInt    `json:"RANK,omitzero"`
	N_ENS_VLD        validtypes.ValidInt    `json:"N_ENS_VLD,omitzero"`
	ENS              map[string]interface{} `json:"ENS,omitzero"`
	OBS_QC           validtypes.ValidString `json:"OBS_QC,omitzero"`
	ENS_MEAN         validtypes.ValidInt    `json:"ENS_MEAN,omitzero"`
	CLIMO_MEAN       validtypes.ValidFloat  `json:"CLIMO_MEAN,omitzero"`
	SPREAD           validtypes.ValidFloat  `json:"SPREAD,omitzero"`
	ENS_MEAN_OERR    validtypes.ValidInt    `json:"ENS_MEAN_OERR,omitzero"`
	SPREAD_OERR      validtypes.ValidFloat  `json:"SPREAD_OERR,omitzero"`
	SPREAD_PLUS_OERR validtypes.ValidFloat  `json:"SPREAD_PLUS_OERR,omitzero"`
	CLIMO_STDEV      validtypes.ValidFloat  `json:"CLIMO_STDEV,omitzero"`
}

type STAT_PCT_data struct {
	TOTAL  validtypes.ValidInt    `json:"TOTAL,omitzero"`
	THRESH map[string]interface{} `json:"THRESH,omitzero"`
}

type STAT_PHIST_data struct {
	TOTAL    validtypes.ValidInt    `json:"TOTAL,omitzero"`
	BIN_SIZE validtypes.ValidInt    `json:"BIN_SIZE,omitzero"`
	BIN      map[string]interface{} `json:"BIN,omitzero"`
}

type STAT_PJC_data struct {
	TOTAL  validtypes.ValidInt    `json:"TOTAL,omitzero"`
	THRESH map[string]interface{} `json:"THRESH,omitzero"`
}

type STAT_PRC_data struct {
	TOTAL  validtypes.ValidInt    `json:"TOTAL,omitzero"`
	THRESH map[string]interface{} `json:"THRESH,omitzero"`
}

type STAT_PSTD_data struct {
	TOTAL       validtypes.ValidInt    `json:"TOTAL,omitzero"`
	THRESH      map[string]interface{} `json:"THRESH,omitzero"`
	BASER_NCL   validtypes.ValidFloat  `json:"BASER_NCL,omitzero"`
	BASER_NCU   validtypes.ValidFloat  `json:"BASER_NCU,omitzero"`
	RELIABILITY validtypes.ValidFloat  `json:"RELIABILITY,omitzero"`
	RESOLUTION  validtypes.ValidFloat  `json:"RESOLUTION,omitzero"`
	UNCERTAINTY validtypes.ValidFloat  `json:"UNCERTAINTY,omitzero"`
	ROC_AUC     validtypes.ValidFloat  `json:"ROC_AUC,omitzero"`
	BRIER       validtypes.ValidFloat  `json:"BRIER,omitzero"`
	BRIER_NCL   validtypes.ValidFloat  `json:"BRIER_NCL,omitzero"`
	BRIER_NCU   validtypes.ValidFloat  `json:"BRIER_NCU,omitzero"`
	BRIERCL     validtypes.ValidFloat  `json:"BRIERCL,omitzero"`
	BRIERCL_NCL validtypes.ValidFloat  `json:"BRIERCL_NCL,omitzero"`
	BRIERCL_NCU validtypes.ValidFloat  `json:"BRIERCL_NCU,omitzero"`
	BSS         validtypes.ValidFloat  `json:"BSS,omitzero"`
	BSS_SMPL    validtypes.ValidFloat  `json:"BSS_SMPL,omitzero"`
	THRESH_I    validtypes.ValidInt    `json:"THRESH_I,omitzero"`
}

type STAT_RELP_data struct {
	TOTAL validtypes.ValidInt    `json:"TOTAL,omitzero"`
	ENS   map[string]interface{} `json:"ENS,omitzero"`
}

type STAT_RHIST_data struct {
	TOTAL validtypes.ValidInt    `json:"TOTAL,omitzero"`
	RANK  map[string]interface{} `json:"RANK,omitzero"`
}

type STAT_RPS_data struct {
	TOTAL     validtypes.ValidInt   `json:"TOTAL,omitzero"`
	N_PROB    validtypes.ValidInt   `json:"N_PROB,omitzero"`
	RPS_REL   validtypes.ValidFloat `json:"RPS_REL,omitzero"`
	RPS_RES   validtypes.ValidFloat `json:"RPS_RES,omitzero"`
	RPS_UNC   validtypes.ValidFloat `json:"RPS_UNC,omitzero"`
	RPS       validtypes.ValidFloat `json:"RPS,omitzero"`
	RPSS      validtypes.ValidFloat `json:"RPSS,omitzero"`
	RPSS_SMPL validtypes.ValidFloat `json:"RPSS_SMPL,omitzero"`
	RPS_COMP  validtypes.ValidFloat `json:"RPS_COMP,omitzero"`
}

type STAT_SAL1L2_data struct {
	TOTAL  validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FABAR  validtypes.ValidFloat `json:"FABAR,omitzero"`
	OABAR  validtypes.ValidFloat `json:"OABAR,omitzero"`
	FOABAR validtypes.ValidFloat `json:"FOABAR,omitzero"`
	FFABAR validtypes.ValidFloat `json:"FFABAR,omitzero"`
	OOABAR validtypes.ValidFloat `json:"OOABAR,omitzero"`
	MAE    validtypes.ValidFloat `json:"MAE,omitzero"`
}

type STAT_SL1L2_data struct {
	TOTAL validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FBAR  validtypes.ValidFloat `json:"FBAR,omitzero"`
	OBAR  validtypes.ValidFloat `json:"OBAR,omitzero"`
	FOBAR validtypes.ValidFloat `json:"FOBAR,omitzero"`
	FFBAR validtypes.ValidFloat `json:"FFBAR,omitzero"`
	OOBAR validtypes.ValidFloat `json:"OOBAR,omitzero"`
	MAE   validtypes.ValidFloat `json:"MAE,omitzero"`
}

type STAT_SSIDX_data struct {
	FCST_MODEL validtypes.ValidString `json:"FCST_MODEL,omitzero"`
	REF_MODEL  validtypes.ValidString `json:"REF_MODEL,omitzero"`
	N_INIT     validtypes.ValidInt    `json:"N_INIT,omitzero"`
	N_TERM     validtypes.ValidInt    `json:"N_TERM,omitzero"`
	N_VLD      validtypes.ValidInt    `json:"N_VLD,omitzero"`
	SS_INDEX   validtypes.ValidFloat  `json:"SS_INDEX,omitzero"`
}

type STAT_SSVAR_data struct {
	TOTAL       validtypes.ValidInt   `json:"TOTAL,omitzero"`
	N_BIN       validtypes.ValidInt   `json:"N_BIN,omitzero"`
	BIN_I       validtypes.ValidInt   `json:"BIN_I,omitzero"`
	BIN_N       validtypes.ValidInt   `json:"BIN_N,omitzero"`
	VAR_MIN     validtypes.ValidFloat `json:"VAR_MIN,omitzero"`
	VAR_MAX     validtypes.ValidFloat `json:"VAR_MAX,omitzero"`
	VAR_MEAN    validtypes.ValidFloat `json:"VAR_MEAN,omitzero"`
	FBAR        validtypes.ValidFloat `json:"FBAR,omitzero"`
	OBAR        validtypes.ValidFloat `json:"OBAR,omitzero"`
	FOBAR       validtypes.ValidFloat `json:"FOBAR,omitzero"`
	FFBAR       validtypes.ValidFloat `json:"FFBAR,omitzero"`
	OOBAR       validtypes.ValidFloat `json:"OOBAR,omitzero"`
	FBAR_NCL    validtypes.ValidFloat `json:"FBAR_NCL,omitzero"`
	FBAR_NCU    validtypes.ValidFloat `json:"FBAR_NCU,omitzero"`
	FSTDEV      validtypes.ValidFloat `json:"FSTDEV,omitzero"`
	FSTDEV_NCL  validtypes.ValidFloat `json:"FSTDEV_NCL,omitzero"`
	FSTDEV_NCU  validtypes.ValidFloat `json:"FSTDEV_NCU,omitzero"`
	OBAR_NCL    validtypes.ValidFloat `json:"OBAR_NCL,omitzero"`
	OBAR_NCU    validtypes.ValidFloat `json:"OBAR_NCU,omitzero"`
	OSTDEV      validtypes.ValidFloat `json:"OSTDEV,omitzero"`
	OSTDEV_NCL  validtypes.ValidFloat `json:"OSTDEV_NCL,omitzero"`
	OSTDEV_NCU  validtypes.ValidFloat `json:"OSTDEV_NCU,omitzero"`
	PR_CORR     validtypes.ValidFloat `json:"PR_CORR,omitzero"`
	PR_CORR_NCL validtypes.ValidFloat `json:"PR_CORR_NCL,omitzero"`
	PR_CORR_NCU validtypes.ValidFloat `json:"PR_CORR_NCU,omitzero"`
	ME          validtypes.ValidFloat `json:"ME,omitzero"`
	ME_NCL      validtypes.ValidFloat `json:"ME_NCL,omitzero"`
	ME_NCU      validtypes.ValidFloat `json:"ME_NCU,omitzero"`
	ESTDEV      validtypes.ValidFloat `json:"ESTDEV,omitzero"`
	ESTDEV_NCL  validtypes.ValidFloat `json:"ESTDEV_NCL,omitzero"`
	ESTDEV_NCU  validtypes.ValidFloat `json:"ESTDEV_NCU,omitzero"`
	MBIAS       validtypes.ValidFloat `json:"MBIAS,omitzero"`
	MSE         validtypes.ValidFloat `json:"MSE,omitzero"`
	BCMSE       validtypes.ValidFloat `json:"BCMSE,omitzero"`
	RMSE        validtypes.ValidFloat `json:"RMSE,omitzero"`
}

type STAT_VAL1L2_data struct {
	TOTAL    validtypes.ValidInt   `json:"TOTAL,omitzero"`
	UFABAR   validtypes.ValidFloat `json:"UFABAR,omitzero"`
	VFABAR   validtypes.ValidFloat `json:"VFABAR,omitzero"`
	UOABAR   validtypes.ValidFloat `json:"UOABAR,omitzero"`
	VOABAR   validtypes.ValidFloat `json:"VOABAR,omitzero"`
	UVFOABAR validtypes.ValidFloat `json:"UVFOABAR,omitzero"`
	UVFFABAR validtypes.ValidFloat `json:"UVFFABAR,omitzero"`
	UVOOABAR validtypes.ValidFloat `json:"UVOOABAR,omitzero"`
}

type STAT_VCNT_data struct {
	TOTAL            validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FBAR             validtypes.ValidFloat `json:"FBAR,omitzero"`
	FBAR_BCL         validtypes.ValidFloat `json:"FBAR_BCL,omitzero"`
	FBAR_BCU         validtypes.ValidFloat `json:"FBAR_BCU,omitzero"`
	OBAR             validtypes.ValidFloat `json:"OBAR,omitzero"`
	OBAR_BCL         validtypes.ValidFloat `json:"OBAR_BCL,omitzero"`
	OBAR_BCU         validtypes.ValidFloat `json:"OBAR_BCU,omitzero"`
	FS_RMS           validtypes.ValidFloat `json:"FS_RMS,omitzero"`
	FS_RMS_BCL       validtypes.ValidFloat `json:"FS_RMS_BCL,omitzero"`
	FS_RMS_BCU       validtypes.ValidFloat `json:"FS_RMS_BCU,omitzero"`
	OS_RMS           validtypes.ValidFloat `json:"OS_RMS,omitzero"`
	OS_RMS_BCL       validtypes.ValidFloat `json:"OS_RMS_BCL,omitzero"`
	OS_RMS_BCU       validtypes.ValidFloat `json:"OS_RMS_BCU,omitzero"`
	MSVE             validtypes.ValidFloat `json:"MSVE,omitzero"`
	MSVE_BCL         validtypes.ValidFloat `json:"MSVE_BCL,omitzero"`
	MSVE_BCU         validtypes.ValidFloat `json:"MSVE_BCU,omitzero"`
	RMSVE            validtypes.ValidFloat `json:"RMSVE,omitzero"`
	RMSVE_BCL        validtypes.ValidFloat `json:"RMSVE_BCL,omitzero"`
	RMSVE_BCU        validtypes.ValidFloat `json:"RMSVE_BCU,omitzero"`
	FSTDEV           validtypes.ValidFloat `json:"FSTDEV,omitzero"`
	FSTDEV_BCL       validtypes.ValidFloat `json:"FSTDEV_BCL,omitzero"`
	FSTDEV_BCU       validtypes.ValidFloat `json:"FSTDEV_BCU,omitzero"`
	OSTDEV           validtypes.ValidFloat `json:"OSTDEV,omitzero"`
	OSTDEV_BCL       validtypes.ValidFloat `json:"OSTDEV_BCL,omitzero"`
	OSTDEV_BCU       validtypes.ValidFloat `json:"OSTDEV_BCU,omitzero"`
	FDIR             validtypes.ValidFloat `json:"FDIR,omitzero"`
	FDIR_BCL         validtypes.ValidFloat `json:"FDIR_BCL,omitzero"`
	FDIR_BCU         validtypes.ValidFloat `json:"FDIR_BCU,omitzero"`
	ODIR             validtypes.ValidFloat `json:"ODIR,omitzero"`
	ODIR_BCL         validtypes.ValidFloat `json:"ODIR_BCL,omitzero"`
	ODIR_BCU         validtypes.ValidFloat `json:"ODIR_BCU,omitzero"`
	FBAR_SPEED       validtypes.ValidFloat `json:"FBAR_SPEED,omitzero"`
	FBAR_SPEED_BCL   validtypes.ValidFloat `json:"FBAR_SPEED_BCL,omitzero"`
	FBAR_SPEED_BCU   validtypes.ValidFloat `json:"FBAR_SPEED_BCU,omitzero"`
	OBAR_SPEED       validtypes.ValidFloat `json:"OBAR_SPEED,omitzero"`
	OBAR_SPEED_BCL   validtypes.ValidFloat `json:"OBAR_SPEED_BCL,omitzero"`
	OBAR_SPEED_BCU   validtypes.ValidFloat `json:"OBAR_SPEED_BCU,omitzero"`
	VDIFF_SPEED      validtypes.ValidFloat `json:"VDIFF_SPEED,omitzero"`
	VDIFF_SPEED_BCL  validtypes.ValidFloat `json:"VDIFF_SPEED_BCL,omitzero"`
	VDIFF_SPEED_BCU  validtypes.ValidFloat `json:"VDIFF_SPEED_BCU,omitzero"`
	VDIFF_DIR        validtypes.ValidFloat `json:"VDIFF_DIR,omitzero"`
	VDIFF_DIR_BCL    validtypes.ValidFloat `json:"VDIFF_DIR_BCL,omitzero"`
	VDIFF_DIR_BCU    validtypes.ValidFloat `json:"VDIFF_DIR_BCU,omitzero"`
	SPEED_ERR        validtypes.ValidFloat `json:"SPEED_ERR,omitzero"`
	SPEED_ERR_BCL    validtypes.ValidFloat `json:"SPEED_ERR_BCL,omitzero"`
	SPEED_ERR_BCU    validtypes.ValidFloat `json:"SPEED_ERR_BCU,omitzero"`
	SPEED_ABSERR     validtypes.ValidFloat `json:"SPEED_ABSERR,omitzero"`
	SPEED_ABSERR_BCL validtypes.ValidFloat `json:"SPEED_ABSERR_BCL,omitzero"`
	SPEED_ABSERR_BCU validtypes.ValidFloat `json:"SPEED_ABSERR_BCU,omitzero"`
	DIR_ERR          validtypes.ValidFloat `json:"DIR_ERR,omitzero"`
	DIR_ERR_BCL      validtypes.ValidFloat `json:"DIR_ERR_BCL,omitzero"`
	DIR_ERR_BCU      validtypes.ValidFloat `json:"DIR_ERR_BCU,omitzero"`
	DIR_ABSERR       validtypes.ValidFloat `json:"DIR_ABSERR,omitzero"`
	DIR_ABSERR_BCL   validtypes.ValidFloat `json:"DIR_ABSERR_BCL,omitzero"`
	DIR_ABSERR_BCU   validtypes.ValidFloat `json:"DIR_ABSERR_BCU,omitzero"`
}

type STAT_VL1L2_data struct {
	TOTAL       validtypes.ValidInt   `json:"TOTAL,omitzero"`
	UFBAR       validtypes.ValidFloat `json:"UFBAR,omitzero"`
	VFBAR       validtypes.ValidFloat `json:"VFBAR,omitzero"`
	UOBAR       validtypes.ValidFloat `json:"UOBAR,omitzero"`
	VOBAR       validtypes.ValidFloat `json:"VOBAR,omitzero"`
	UVFOBAR     validtypes.ValidFloat `json:"UVFOBAR,omitzero"`
	UVFFBAR     validtypes.ValidFloat `json:"UVFFBAR,omitzero"`
	UVOOBAR     validtypes.ValidFloat `json:"UVOOBAR,omitzero"`
	F_SPEED_BAR validtypes.ValidFloat `json:"F_SPEED_BAR,omitzero"`
	O_SPEED_BAR validtypes.ValidFloat `json:"O_SPEED_BAR,omitzero"`
}

type TCST_PROBRIRW_data struct {
	ALAT        validtypes.ValidFloat  `json:"ALAT,omitzero"`
	ALON        validtypes.ValidFloat  `json:"ALON,omitzero"`
	BLAT        validtypes.ValidFloat  `json:"BLAT,omitzero"`
	BLON        validtypes.ValidFloat  `json:"BLON,omitzero"`
	INITIALS    validtypes.ValidString `json:"INITIALS,omitzero"`
	TK_ERR      validtypes.ValidFloat  `json:"TK_ERR,omitzero"`
	X_ERR       validtypes.ValidFloat  `json:"X_ERR,omitzero"`
	Y_ERR       validtypes.ValidFloat  `json:"Y_ERR,omitzero"`
	ADLAND      validtypes.ValidFloat  `json:"ADLAND,omitzero"`
	BDLAND      validtypes.ValidFloat  `json:"BDLAND,omitzero"`
	RIRW_BEG    validtypes.ValidInt    `json:"RIRW_BEG,omitzero"`
	RIRW_END    validtypes.ValidInt    `json:"RIRW_END,omitzero"`
	RIRW_WINDOW validtypes.ValidInt    `json:"RIRW_WINDOW,omitzero"`
	AWIND_END   validtypes.ValidFloat  `json:"AWIND_END,omitzero"`
	BWIND_BEG   validtypes.ValidFloat  `json:"BWIND_BEG,omitzero"`
	BWIND_END   validtypes.ValidFloat  `json:"BWIND_END,omitzero"`
	BDELTA      validtypes.ValidFloat  `json:"BDELTA,omitzero"`
	BDELTA_MAX  validtypes.ValidFloat  `json:"BDELTA_MAX,omitzero"`
	BLEVEL_BEG  validtypes.ValidString `json:"BLEVEL_BEG,omitzero"`
	BLEVEL_END  validtypes.ValidString `json:"BLEVEL_END,omitzero"`
	THRESH      map[string]interface{} `json:"THRESH,omitzero"`
	INIT        validtypes.ValidInt    `json:"INIT,omitzero"`
}

type TCST_TCMPR_data struct {
	TOTAL       validtypes.ValidInt    `json:"TOTAL,omitzero"`
	INDEX       validtypes.ValidInt    `json:"INDEX,omitzero"`
	LEVEL       validtypes.ValidString `json:"LEVEL,omitzero"`
	WATCH_WARN  validtypes.ValidString `json:"WATCH_WARN,omitzero"`
	INITIALS    validtypes.ValidString `json:"INITIALS,omitzero"`
	ALAT        validtypes.ValidFloat  `json:"ALAT,omitzero"`
	ALON        validtypes.ValidFloat  `json:"ALON,omitzero"`
	BLAT        validtypes.ValidFloat  `json:"BLAT,omitzero"`
	BLON        validtypes.ValidFloat  `json:"BLON,omitzero"`
	TK_ERR      validtypes.ValidFloat  `json:"TK_ERR,omitzero"`
	X_ERR       validtypes.ValidFloat  `json:"X_ERR,omitzero"`
	Y_ERR       validtypes.ValidFloat  `json:"Y_ERR,omitzero"`
	ALTK_ERR    validtypes.ValidFloat  `json:"ALTK_ERR,omitzero"`
	CRTK_ERR    validtypes.ValidFloat  `json:"CRTK_ERR,omitzero"`
	ADLAND      validtypes.ValidFloat  `json:"ADLAND,omitzero"`
	BDLAND      validtypes.ValidFloat  `json:"BDLAND,omitzero"`
	AMSLP       validtypes.ValidFloat  `json:"AMSLP,omitzero"`
	BMSLP       validtypes.ValidFloat  `json:"BMSLP,omitzero"`
	AMAX_WIND   validtypes.ValidFloat  `json:"AMAX_WIND,omitzero"`
	BMAX_WIND   validtypes.ValidFloat  `json:"BMAX_WIND,omitzero"`
	AAL_WIND_34 validtypes.ValidFloat  `json:"AAL_WIND_34,omitzero"`
	BAL_WIND_34 validtypes.ValidFloat  `json:"BAL_WIND_34,omitzero"`
	ANE_WIND_34 validtypes.ValidFloat  `json:"ANE_WIND_34,omitzero"`
	BNE_WIND_34 validtypes.ValidFloat  `json:"BNE_WIND_34,omitzero"`
	ASE_WIND_34 validtypes.ValidFloat  `json:"ASE_WIND_34,omitzero"`
	BSE_WIND_34 validtypes.ValidFloat  `json:"BSE_WIND_34,omitzero"`
	ASW_WIND_34 validtypes.ValidFloat  `json:"ASW_WIND_34,omitzero"`
	BSW_WIND_34 validtypes.ValidFloat  `json:"BSW_WIND_34,omitzero"`
	ANW_WIND_34 validtypes.ValidFloat  `json:"ANW_WIND_34,omitzero"`
	BNW_WIND_34 validtypes.ValidFloat  `json:"BNW_WIND_34,omitzero"`
	AAL_WIND_50 validtypes.ValidFloat  `json:"AAL_WIND_50,omitzero"`
	BAL_WIND_50 validtypes.ValidFloat  `json:"BAL_WIND_50,omitzero"`
	ANE_WIND_50 validtypes.ValidFloat  `json:"ANE_WIND_50,omitzero"`
	BNE_WIND_50 validtypes.ValidFloat  `json:"BNE_WIND_50,omitzero"`
	ASE_WIND_50 validtypes.ValidFloat  `json:"ASE_WIND_50,omitzero"`
	BSE_WIND_50 validtypes.ValidFloat  `json:"BSE_WIND_50,omitzero"`
	ASW_WIND_50 validtypes.ValidFloat  `json:"ASW_WIND_50,omitzero"`
	BSW_WIND_50 validtypes.ValidFloat  `json:"BSW_WIND_50,omitzero"`
	ANW_WIND_50 validtypes.ValidFloat  `json:"ANW_WIND_50,omitzero"`
	BNW_WIND_50 validtypes.ValidFloat  `json:"BNW_WIND_50,omitzero"`
	AAL_WIND_64 validtypes.ValidFloat  `json:"AAL_WIND_64,omitzero"`
	BAL_WIND_64 validtypes.ValidFloat  `json:"BAL_WIND_64,omitzero"`
	ANE_WIND_64 validtypes.ValidFloat  `json:"ANE_WIND_64,omitzero"`
	BNE_WIND_64 validtypes.ValidFloat  `json:"BNE_WIND_64,omitzero"`
	ASE_WIND_64 validtypes.ValidFloat  `json:"ASE_WIND_64,omitzero"`
	BSE_WIND_64 validtypes.ValidFloat  `json:"BSE_WIND_64,omitzero"`
	ASW_WIND_64 validtypes.ValidFloat  `json:"ASW_WIND_64,omitzero"`
	BSW_WIND_64 validtypes.ValidFloat  `json:"BSW_WIND_64,omitzero"`
	ANW_WIND_64 validtypes.ValidFloat  `json:"ANW_WIND_64,omitzero"`
	BNW_WIND_64 validtypes.ValidFloat  `json:"BNW_WIND_64,omitzero"`
	ARADP       validtypes.ValidString `json:"ARADP,omitzero"`
	BRADP       validtypes.ValidFloat  `json:"BRADP,omitzero"`
	ARRP        validtypes.ValidInt    `json:"ARRP,omitzero"`
	BRRP        validtypes.ValidFloat  `json:"BRRP,omitzero"`
	AMRD        validtypes.ValidInt    `json:"AMRD,omitzero"`
	BMRD        validtypes.ValidFloat  `json:"BMRD,omitzero"`
	AGUSTS      validtypes.ValidInt    `json:"AGUSTS,omitzero"`
	BGUSTS      validtypes.ValidFloat  `json:"BGUSTS,omitzero"`
	AEYE        validtypes.ValidInt    `json:"AEYE,omitzero"`
	BEYE        validtypes.ValidFloat  `json:"BEYE,omitzero"`
	ADIR        validtypes.ValidInt    `json:"ADIR,omitzero"`
	BDIR        validtypes.ValidFloat  `json:"BDIR,omitzero"`
	ASPEED      validtypes.ValidInt    `json:"ASPEED,omitzero"`
	BSPEED      validtypes.ValidFloat  `json:"BSPEED,omitzero"`
	ADEPTH      validtypes.ValidInt    `json:"ADEPTH,omitzero"`
	BDEPTH      validtypes.ValidFloat  `json:"BDEPTH,omitzero"`
	INIT        validtypes.ValidInt    `json:"INIT,omitzero"`
}

// fillStructure functions

// Sets MODE_CTS_data struct's fields
func (s *MODE_CTS_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "FIELD", s.FIELD.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "FY_OY", s.FY_OY.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FY_ON", s.FY_ON.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "FN_OY", s.FN_OY.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FN_ON", s.FN_ON.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "BASER", s.BASER.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FMEAN", s.FMEAN.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "ACC", s.ACC.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FBIAS", s.FBIAS.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "PODY", s.PODY.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "PODN", s.PODN.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "POFD", s.POFD.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "FAR", s.FAR.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "CSI", s.CSI.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "GSS", s.GSS.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "HK", s.HK.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "HSS", s.HSS.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "ODDS", s.ODDS.UnmarshalText([]byte(fields[18])))
	return errors.Join(errs...)
}

// Sets MODE_OBJ_data struct's fields
func (s *MODE_OBJ_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "OBJECT_ID", s.OBJECT_ID.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "OBJECT_CAT", s.OBJECT_CAT.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "CENTROID_X", s.CENTROID_X.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "CENTROID_Y", s.CENTROID_Y.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "CENTROID_LAT", s.CENTROID_LAT.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "CENTROID_LON", s.CENTROID_LON.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "AXIS_ANG", s.AXIS_ANG.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "LENGTH", s.LENGTH.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "WIDTH", s.WIDTH.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "AREA", s.AREA.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "AREA_THRESH", s.AREA_THRESH.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "CURVATURE", s.CURVATURE.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "CURVATURE_X", s.CURVATURE_X.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "CURVATURE_Y", s.CURVATURE_Y.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "COMPLEXITY", s.COMPLEXITY.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "INTENSITY_10", s.INTENSITY_10.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "INTENSITY_25", s.INTENSITY_25.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "INTENSITY_50", s.INTENSITY_50.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "INTENSITY_75", s.INTENSITY_75.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "INTENSITY_90", s.INTENSITY_90.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "INTENSITY_USER", s.INTENSITY_USER.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "INTENSITY_SUM", s.INTENSITY_SUM.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "CENTROID_DIST", s.CENTROID_DIST.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "BOUNDARY_DIST", s.BOUNDARY_DIST.UnmarshalText([]byte(fields[23])))
	appendErrorWithContext(&errs, "CONVEX_HULL_DIST", s.CONVEX_HULL_DIST.UnmarshalText([]byte(fields[24])))
	appendErrorWithContext(&errs, "ANGLE_DIFF", s.ANGLE_DIFF.UnmarshalText([]byte(fields[25])))
	appendErrorWithContext(&errs, "ASPECT_DIFF", s.ASPECT_DIFF.UnmarshalText([]byte(fields[26])))
	appendErrorWithContext(&errs, "AREA_RATIO", s.AREA_RATIO.UnmarshalText([]byte(fields[27])))
	appendErrorWithContext(&errs, "INTERSECTION_AREA", s.INTERSECTION_AREA.UnmarshalText([]byte(fields[28])))
	appendErrorWithContext(&errs, "UNION_AREA", s.UNION_AREA.UnmarshalText([]byte(fields[29])))
	appendErrorWithContext(&errs, "SYMMETRIC_DIFF", s.SYMMETRIC_DIFF.UnmarshalText([]byte(fields[30])))
	appendErrorWithContext(&errs, "INTERSECTION_OVER_AREA", s.INTERSECTION_OVER_AREA.UnmarshalText([]byte(fields[31])))
	appendErrorWithContext(&errs, "CURVATURE_RATIO", s.CURVATURE_RATIO.UnmarshalText([]byte(fields[32])))
	appendErrorWithContext(&errs, "COMPLEXITY_RATIO", s.COMPLEXITY_RATIO.UnmarshalText([]byte(fields[33])))
	appendErrorWithContext(&errs, "PERCENTILE_INTENSITY_RATIO", s.PERCENTILE_INTENSITY_RATIO.UnmarshalText([]byte(fields[34])))
	appendErrorWithContext(&errs, "INTEREST", s.INTEREST.UnmarshalText([]byte(fields[35])))
	return errors.Join(errs...)
}

// Sets STAT_CNT_data struct's fields
func (s *STAT_CNT_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "FBAR", s.FBAR.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "FBAR_NCL", s.FBAR_NCL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FBAR_NCU", s.FBAR_NCU.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "FBAR_BCL", s.FBAR_BCL.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FBAR_BCU", s.FBAR_BCU.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "FSTDEV", s.FSTDEV.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FSTDEV_NCL", s.FSTDEV_NCL.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "FSTDEV_NCU", s.FSTDEV_NCU.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FSTDEV_BCL", s.FSTDEV_BCL.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FSTDEV_BCU", s.FSTDEV_BCU.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "OBAR", s.OBAR.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OBAR_NCL", s.OBAR_NCL.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "OBAR_NCU", s.OBAR_NCU.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "OBAR_BCL", s.OBAR_BCL.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "OBAR_BCU", s.OBAR_BCU.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "OSTDEV", s.OSTDEV.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "OSTDEV_NCL", s.OSTDEV_NCL.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "OSTDEV_NCU", s.OSTDEV_NCU.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "OSTDEV_BCL", s.OSTDEV_BCL.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OSTDEV_BCU", s.OSTDEV_BCU.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "PR_CORR", s.PR_CORR.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "PR_CORR_NCL", s.PR_CORR_NCL.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "PR_CORR_NCU", s.PR_CORR_NCU.UnmarshalText([]byte(fields[23])))
	appendErrorWithContext(&errs, "PR_CORR_BCL", s.PR_CORR_BCL.UnmarshalText([]byte(fields[24])))
	appendErrorWithContext(&errs, "PR_CORR_BCU", s.PR_CORR_BCU.UnmarshalText([]byte(fields[25])))
	appendErrorWithContext(&errs, "SP_CORR", s.SP_CORR.UnmarshalText([]byte(fields[26])))
	appendErrorWithContext(&errs, "KT_CORR", s.KT_CORR.UnmarshalText([]byte(fields[27])))
	appendErrorWithContext(&errs, "RANKS", s.RANKS.UnmarshalText([]byte(fields[28])))
	appendErrorWithContext(&errs, "FRANK_TIES", s.FRANK_TIES.UnmarshalText([]byte(fields[29])))
	appendErrorWithContext(&errs, "ORANK_TIES", s.ORANK_TIES.UnmarshalText([]byte(fields[30])))
	appendErrorWithContext(&errs, "ME", s.ME.UnmarshalText([]byte(fields[31])))
	appendErrorWithContext(&errs, "ME_NCL", s.ME_NCL.UnmarshalText([]byte(fields[32])))
	appendErrorWithContext(&errs, "ME_NCU", s.ME_NCU.UnmarshalText([]byte(fields[33])))
	appendErrorWithContext(&errs, "ME_BCL", s.ME_BCL.UnmarshalText([]byte(fields[34])))
	appendErrorWithContext(&errs, "ME_BCU", s.ME_BCU.UnmarshalText([]byte(fields[35])))
	appendErrorWithContext(&errs, "ESTDEV", s.ESTDEV.UnmarshalText([]byte(fields[36])))
	appendErrorWithContext(&errs, "ESTDEV_NCL", s.ESTDEV_NCL.UnmarshalText([]byte(fields[37])))
	appendErrorWithContext(&errs, "ESTDEV_NCU", s.ESTDEV_NCU.UnmarshalText([]byte(fields[38])))
	appendErrorWithContext(&errs, "ESTDEV_BCL", s.ESTDEV_BCL.UnmarshalText([]byte(fields[39])))
	appendErrorWithContext(&errs, "ESTDEV_BCU", s.ESTDEV_BCU.UnmarshalText([]byte(fields[40])))
	appendErrorWithContext(&errs, "MBIAS", s.MBIAS.UnmarshalText([]byte(fields[41])))
	appendErrorWithContext(&errs, "MBIAS_BCL", s.MBIAS_BCL.UnmarshalText([]byte(fields[42])))
	appendErrorWithContext(&errs, "MBIAS_BCU", s.MBIAS_BCU.UnmarshalText([]byte(fields[43])))
	appendErrorWithContext(&errs, "MAE", s.MAE.UnmarshalText([]byte(fields[44])))
	appendErrorWithContext(&errs, "MAE_BCL", s.MAE_BCL.UnmarshalText([]byte(fields[45])))
	appendErrorWithContext(&errs, "MAE_BCU", s.MAE_BCU.UnmarshalText([]byte(fields[46])))
	appendErrorWithContext(&errs, "MSE", s.MSE.UnmarshalText([]byte(fields[47])))
	appendErrorWithContext(&errs, "MSE_BCL", s.MSE_BCL.UnmarshalText([]byte(fields[48])))
	appendErrorWithContext(&errs, "MSE_BCU", s.MSE_BCU.UnmarshalText([]byte(fields[49])))
	appendErrorWithContext(&errs, "BCMSE", s.BCMSE.UnmarshalText([]byte(fields[50])))
	appendErrorWithContext(&errs, "BCMSE_BCL", s.BCMSE_BCL.UnmarshalText([]byte(fields[51])))
	appendErrorWithContext(&errs, "BCMSE_BCU", s.BCMSE_BCU.UnmarshalText([]byte(fields[52])))
	appendErrorWithContext(&errs, "RMSE", s.RMSE.UnmarshalText([]byte(fields[53])))
	appendErrorWithContext(&errs, "RMSE_BCL", s.RMSE_BCL.UnmarshalText([]byte(fields[54])))
	appendErrorWithContext(&errs, "RMSE_BCU", s.RMSE_BCU.UnmarshalText([]byte(fields[55])))
	appendErrorWithContext(&errs, "E10", s.E10.UnmarshalText([]byte(fields[56])))
	appendErrorWithContext(&errs, "E10_BCL", s.E10_BCL.UnmarshalText([]byte(fields[57])))
	appendErrorWithContext(&errs, "E10_BCU", s.E10_BCU.UnmarshalText([]byte(fields[58])))
	appendErrorWithContext(&errs, "E25", s.E25.UnmarshalText([]byte(fields[59])))
	appendErrorWithContext(&errs, "E25_BCL", s.E25_BCL.UnmarshalText([]byte(fields[60])))
	appendErrorWithContext(&errs, "E25_BCU", s.E25_BCU.UnmarshalText([]byte(fields[61])))
	appendErrorWithContext(&errs, "E50", s.E50.UnmarshalText([]byte(fields[62])))
	appendErrorWithContext(&errs, "E50_BCL", s.E50_BCL.UnmarshalText([]byte(fields[63])))
	appendErrorWithContext(&errs, "E50_BCU", s.E50_BCU.UnmarshalText([]byte(fields[64])))
	appendErrorWithContext(&errs, "E75", s.E75.UnmarshalText([]byte(fields[65])))
	appendErrorWithContext(&errs, "E75_BCL", s.E75_BCL.UnmarshalText([]byte(fields[66])))
	appendErrorWithContext(&errs, "E75_BCU", s.E75_BCU.UnmarshalText([]byte(fields[67])))
	appendErrorWithContext(&errs, "E90", s.E90.UnmarshalText([]byte(fields[68])))
	appendErrorWithContext(&errs, "E90_BCL", s.E90_BCL.UnmarshalText([]byte(fields[69])))
	appendErrorWithContext(&errs, "E90_BCU", s.E90_BCU.UnmarshalText([]byte(fields[70])))
	appendErrorWithContext(&errs, "EIQR", s.EIQR.UnmarshalText([]byte(fields[71])))
	appendErrorWithContext(&errs, "EIQR_BCL", s.EIQR_BCL.UnmarshalText([]byte(fields[72])))
	appendErrorWithContext(&errs, "EIQR_BCU", s.EIQR_BCU.UnmarshalText([]byte(fields[73])))
	appendErrorWithContext(&errs, "MAD", s.MAD.UnmarshalText([]byte(fields[74])))
	appendErrorWithContext(&errs, "MAD_BCL", s.MAD_BCL.UnmarshalText([]byte(fields[75])))
	appendErrorWithContext(&errs, "MAD_BCU", s.MAD_BCU.UnmarshalText([]byte(fields[76])))
	appendErrorWithContext(&errs, "ANOM_CORR", s.ANOM_CORR.UnmarshalText([]byte(fields[77])))
	appendErrorWithContext(&errs, "ANOM_CORR_NCL", s.ANOM_CORR_NCL.UnmarshalText([]byte(fields[78])))
	appendErrorWithContext(&errs, "ANOM_CORR_NCU", s.ANOM_CORR_NCU.UnmarshalText([]byte(fields[79])))
	appendErrorWithContext(&errs, "ANOM_CORR_BCL", s.ANOM_CORR_BCL.UnmarshalText([]byte(fields[80])))
	appendErrorWithContext(&errs, "ANOM_CORR_BCU", s.ANOM_CORR_BCU.UnmarshalText([]byte(fields[81])))
	appendErrorWithContext(&errs, "ME2", s.ME2.UnmarshalText([]byte(fields[82])))
	appendErrorWithContext(&errs, "ME2_BCL", s.ME2_BCL.UnmarshalText([]byte(fields[83])))
	appendErrorWithContext(&errs, "ME2_BCU", s.ME2_BCU.UnmarshalText([]byte(fields[84])))
	appendErrorWithContext(&errs, "MSESS", s.MSESS.UnmarshalText([]byte(fields[85])))
	appendErrorWithContext(&errs, "MSESS_BCL", s.MSESS_BCL.UnmarshalText([]byte(fields[86])))
	appendErrorWithContext(&errs, "MSESS_BCU", s.MSESS_BCU.UnmarshalText([]byte(fields[87])))
	appendErrorWithContext(&errs, "RMSFA", s.RMSFA.UnmarshalText([]byte(fields[88])))
	appendErrorWithContext(&errs, "RMSFA_BCL", s.RMSFA_BCL.UnmarshalText([]byte(fields[89])))
	appendErrorWithContext(&errs, "RMSFA_BCU", s.RMSFA_BCU.UnmarshalText([]byte(fields[90])))
	appendErrorWithContext(&errs, "RMSOA", s.RMSOA.UnmarshalText([]byte(fields[91])))
	appendErrorWithContext(&errs, "RMSOA_BCL", s.RMSOA_BCL.UnmarshalText([]byte(fields[92])))
	appendErrorWithContext(&errs, "RMSOA_BCU", s.RMSOA_BCU.UnmarshalText([]byte(fields[93])))
	appendErrorWithContext(&errs, "ANOM_CORR_UNCNTR", s.ANOM_CORR_UNCNTR.UnmarshalText([]byte(fields[94])))
	appendErrorWithContext(&errs, "ANOM_CORR_UNCNTR_BCL", s.ANOM_CORR_UNCNTR_BCL.UnmarshalText([]byte(fields[95])))
	appendErrorWithContext(&errs, "ANOM_CORR_UNCNTR_BCU", s.ANOM_CORR_UNCNTR_BCU.UnmarshalText([]byte(fields[96])))
	appendErrorWithContext(&errs, "SI", s.SI.UnmarshalText([]byte(fields[97])))
	appendErrorWithContext(&errs, "SI_BCL", s.SI_BCL.UnmarshalText([]byte(fields[98])))
	appendErrorWithContext(&errs, "SI_BCU", s.SI_BCU.UnmarshalText([]byte(fields[99])))
	return errors.Join(errs...)
}

// Sets STAT_CTC_data struct's fields
func (s *STAT_CTC_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "FY_OY", s.FY_OY.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "FY_ON", s.FY_ON.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FN_OY", s.FN_OY.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "FN_ON", s.FN_ON.UnmarshalText([]byte(fields[4])))
	return errors.Join(errs...)
}

// Sets STAT_CTS_data struct's fields
func (s *STAT_CTS_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "BASER", s.BASER.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "BASER_NCL", s.BASER_NCL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "BASER_NCU", s.BASER_NCU.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "BASER_BCL", s.BASER_BCL.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "BASER_BCU", s.BASER_BCU.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "FMEAN", s.FMEAN.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FMEAN_NCL", s.FMEAN_NCL.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "FMEAN_NCU", s.FMEAN_NCU.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FMEAN_BCL", s.FMEAN_BCL.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FMEAN_BCU", s.FMEAN_BCU.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "ACC", s.ACC.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "ACC_NCL", s.ACC_NCL.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "ACC_NCU", s.ACC_NCU.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "ACC_BCL", s.ACC_BCL.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "ACC_BCU", s.ACC_BCU.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "FBIAS", s.FBIAS.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "FBIAS_BCL", s.FBIAS_BCL.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "FBIAS_BCU", s.FBIAS_BCU.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "PODY", s.PODY.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "PODY_NCL", s.PODY_NCL.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "PODY_NCU", s.PODY_NCU.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "PODY_BCL", s.PODY_BCL.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "PODY_BCU", s.PODY_BCU.UnmarshalText([]byte(fields[23])))
	appendErrorWithContext(&errs, "PODN", s.PODN.UnmarshalText([]byte(fields[24])))
	appendErrorWithContext(&errs, "PODN_NCL", s.PODN_NCL.UnmarshalText([]byte(fields[25])))
	appendErrorWithContext(&errs, "PODN_NCU", s.PODN_NCU.UnmarshalText([]byte(fields[26])))
	appendErrorWithContext(&errs, "PODN_BCL", s.PODN_BCL.UnmarshalText([]byte(fields[27])))
	appendErrorWithContext(&errs, "PODN_BCU", s.PODN_BCU.UnmarshalText([]byte(fields[28])))
	appendErrorWithContext(&errs, "POFD", s.POFD.UnmarshalText([]byte(fields[29])))
	appendErrorWithContext(&errs, "POFD_NCL", s.POFD_NCL.UnmarshalText([]byte(fields[30])))
	appendErrorWithContext(&errs, "POFD_NCU", s.POFD_NCU.UnmarshalText([]byte(fields[31])))
	appendErrorWithContext(&errs, "POFD_BCL", s.POFD_BCL.UnmarshalText([]byte(fields[32])))
	appendErrorWithContext(&errs, "POFD_BCU", s.POFD_BCU.UnmarshalText([]byte(fields[33])))
	appendErrorWithContext(&errs, "FAR", s.FAR.UnmarshalText([]byte(fields[34])))
	appendErrorWithContext(&errs, "FAR_NCL", s.FAR_NCL.UnmarshalText([]byte(fields[35])))
	appendErrorWithContext(&errs, "FAR_NCU", s.FAR_NCU.UnmarshalText([]byte(fields[36])))
	appendErrorWithContext(&errs, "FAR_BCL", s.FAR_BCL.UnmarshalText([]byte(fields[37])))
	appendErrorWithContext(&errs, "FAR_BCU", s.FAR_BCU.UnmarshalText([]byte(fields[38])))
	appendErrorWithContext(&errs, "CSI", s.CSI.UnmarshalText([]byte(fields[39])))
	appendErrorWithContext(&errs, "CSI_NCL", s.CSI_NCL.UnmarshalText([]byte(fields[40])))
	appendErrorWithContext(&errs, "CSI_NCU", s.CSI_NCU.UnmarshalText([]byte(fields[41])))
	appendErrorWithContext(&errs, "CSI_BCL", s.CSI_BCL.UnmarshalText([]byte(fields[42])))
	appendErrorWithContext(&errs, "CSI_BCU", s.CSI_BCU.UnmarshalText([]byte(fields[43])))
	appendErrorWithContext(&errs, "GSS", s.GSS.UnmarshalText([]byte(fields[44])))
	appendErrorWithContext(&errs, "GSS_BCL", s.GSS_BCL.UnmarshalText([]byte(fields[45])))
	appendErrorWithContext(&errs, "GSS_BCU", s.GSS_BCU.UnmarshalText([]byte(fields[46])))
	appendErrorWithContext(&errs, "HK", s.HK.UnmarshalText([]byte(fields[47])))
	appendErrorWithContext(&errs, "HK_NCL", s.HK_NCL.UnmarshalText([]byte(fields[48])))
	appendErrorWithContext(&errs, "HK_NCU", s.HK_NCU.UnmarshalText([]byte(fields[49])))
	appendErrorWithContext(&errs, "HK_BCL", s.HK_BCL.UnmarshalText([]byte(fields[50])))
	appendErrorWithContext(&errs, "HK_BCU", s.HK_BCU.UnmarshalText([]byte(fields[51])))
	appendErrorWithContext(&errs, "HSS", s.HSS.UnmarshalText([]byte(fields[52])))
	appendErrorWithContext(&errs, "HSS_BCL", s.HSS_BCL.UnmarshalText([]byte(fields[53])))
	appendErrorWithContext(&errs, "HSS_BCU", s.HSS_BCU.UnmarshalText([]byte(fields[54])))
	appendErrorWithContext(&errs, "ODDS", s.ODDS.UnmarshalText([]byte(fields[55])))
	appendErrorWithContext(&errs, "ODDS_NCL", s.ODDS_NCL.UnmarshalText([]byte(fields[56])))
	appendErrorWithContext(&errs, "ODDS_NCU", s.ODDS_NCU.UnmarshalText([]byte(fields[57])))
	appendErrorWithContext(&errs, "ODDS_BCL", s.ODDS_BCL.UnmarshalText([]byte(fields[58])))
	appendErrorWithContext(&errs, "ODDS_BCU", s.ODDS_BCU.UnmarshalText([]byte(fields[59])))
	appendErrorWithContext(&errs, "LODDS", s.LODDS.UnmarshalText([]byte(fields[60])))
	appendErrorWithContext(&errs, "LODDS_NCL", s.LODDS_NCL.UnmarshalText([]byte(fields[61])))
	appendErrorWithContext(&errs, "LODDS_NCU", s.LODDS_NCU.UnmarshalText([]byte(fields[62])))
	appendErrorWithContext(&errs, "LODDS_BCL", s.LODDS_BCL.UnmarshalText([]byte(fields[63])))
	appendErrorWithContext(&errs, "LODDS_BCU", s.LODDS_BCU.UnmarshalText([]byte(fields[64])))
	appendErrorWithContext(&errs, "ORSS", s.ORSS.UnmarshalText([]byte(fields[65])))
	appendErrorWithContext(&errs, "ORSS_NCL", s.ORSS_NCL.UnmarshalText([]byte(fields[66])))
	appendErrorWithContext(&errs, "ORSS_NCU", s.ORSS_NCU.UnmarshalText([]byte(fields[67])))
	appendErrorWithContext(&errs, "ORSS_BCL", s.ORSS_BCL.UnmarshalText([]byte(fields[68])))
	appendErrorWithContext(&errs, "ORSS_BCU", s.ORSS_BCU.UnmarshalText([]byte(fields[69])))
	appendErrorWithContext(&errs, "EDS", s.EDS.UnmarshalText([]byte(fields[70])))
	appendErrorWithContext(&errs, "EDS_NCL", s.EDS_NCL.UnmarshalText([]byte(fields[71])))
	appendErrorWithContext(&errs, "EDS_NCU", s.EDS_NCU.UnmarshalText([]byte(fields[72])))
	appendErrorWithContext(&errs, "EDS_BCL", s.EDS_BCL.UnmarshalText([]byte(fields[73])))
	appendErrorWithContext(&errs, "EDS_BCU", s.EDS_BCU.UnmarshalText([]byte(fields[74])))
	appendErrorWithContext(&errs, "SEDS", s.SEDS.UnmarshalText([]byte(fields[75])))
	appendErrorWithContext(&errs, "SEDS_NCL", s.SEDS_NCL.UnmarshalText([]byte(fields[76])))
	appendErrorWithContext(&errs, "SEDS_NCU", s.SEDS_NCU.UnmarshalText([]byte(fields[77])))
	appendErrorWithContext(&errs, "SEDS_BCL", s.SEDS_BCL.UnmarshalText([]byte(fields[78])))
	appendErrorWithContext(&errs, "SEDS_BCU", s.SEDS_BCU.UnmarshalText([]byte(fields[79])))
	appendErrorWithContext(&errs, "EDI", s.EDI.UnmarshalText([]byte(fields[80])))
	appendErrorWithContext(&errs, "EDI_NCL", s.EDI_NCL.UnmarshalText([]byte(fields[81])))
	appendErrorWithContext(&errs, "EDI_NCU", s.EDI_NCU.UnmarshalText([]byte(fields[82])))
	appendErrorWithContext(&errs, "EDI_BCL", s.EDI_BCL.UnmarshalText([]byte(fields[83])))
	appendErrorWithContext(&errs, "EDI_BCU", s.EDI_BCU.UnmarshalText([]byte(fields[84])))
	appendErrorWithContext(&errs, "SEDI", s.SEDI.UnmarshalText([]byte(fields[85])))
	appendErrorWithContext(&errs, "SEDI_NCL", s.SEDI_NCL.UnmarshalText([]byte(fields[86])))
	appendErrorWithContext(&errs, "SEDI_NCU", s.SEDI_NCU.UnmarshalText([]byte(fields[87])))
	appendErrorWithContext(&errs, "SEDI_BCL", s.SEDI_BCL.UnmarshalText([]byte(fields[88])))
	appendErrorWithContext(&errs, "SEDI_BCU", s.SEDI_BCU.UnmarshalText([]byte(fields[89])))
	appendErrorWithContext(&errs, "BAGSS", s.BAGSS.UnmarshalText([]byte(fields[90])))
	appendErrorWithContext(&errs, "BAGSS_BCL", s.BAGSS_BCL.UnmarshalText([]byte(fields[91])))
	appendErrorWithContext(&errs, "BAGSS_BCU", s.BAGSS_BCU.UnmarshalText([]byte(fields[92])))
	return errors.Join(errs...)
}

// Sets STAT_DMAP_data struct's fields
func (s *STAT_DMAP_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "FY", s.FY.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "OY", s.OY.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FBIAS", s.FBIAS.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "BADDELEY", s.BADDELEY.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "HAUSDORFF", s.HAUSDORFF.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "MED_FO", s.MED_FO.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "MED_OF", s.MED_OF.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "MED_MIN", s.MED_MIN.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "MED_MAX", s.MED_MAX.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "MED_MEAN", s.MED_MEAN.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FOM_FO", s.FOM_FO.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "FOM_OF", s.FOM_OF.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "FOM_MIN", s.FOM_MIN.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "FOM_MAX", s.FOM_MAX.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "FOM_MEAN", s.FOM_MEAN.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "ZHU_FO", s.ZHU_FO.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "ZHU_OF", s.ZHU_OF.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "ZHU_MIN", s.ZHU_MIN.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "ZHU_MAX", s.ZHU_MAX.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "ZHU_MEAN", s.ZHU_MEAN.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "G", s.G.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "GBETA", s.GBETA.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "BETA_VALUE", s.BETA_VALUE.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets STAT_ECLV_data struct's fields
func (s *STAT_ECLV_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "BASER", s.BASER.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "VALUE_BASER", s.VALUE_BASER.UnmarshalText([]byte(fields[2])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidFloat
	count, err := strconv.Atoi(fields[3])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"CL_", "VALUE_"}
	s.PTS = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 4; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "PTS", value.UnmarshalText([]byte(fields[index])))
			}
			s.PTS[key] = value
		}
	}
	return errors.Join(errs...)
}

// Sets STAT_ECNT_data struct's fields
func (s *STAT_ECNT_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "N_ENS", s.N_ENS.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "CRPS", s.CRPS.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "CRPSS", s.CRPSS.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "IGN", s.IGN.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "ME", s.ME.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "RMSE", s.RMSE.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "SPREAD", s.SPREAD.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "ME_OERR", s.ME_OERR.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "RMSE_OERR", s.RMSE_OERR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "SPREAD_OERR", s.SPREAD_OERR.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "SPREAD_PLUS_OERR", s.SPREAD_PLUS_OERR.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "CRPSCL", s.CRPSCL.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "CRPS_EMP", s.CRPS_EMP.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "CRPSCL_EMP", s.CRPSCL_EMP.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "CRPSS_EMP", s.CRPSS_EMP.UnmarshalText([]byte(fields[15])))
	return errors.Join(errs...)
}

// Sets STAT_FHO_data struct's fields
func (s *STAT_FHO_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "F_RATE", s.F_RATE.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "H_RATE", s.H_RATE.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "O_RATE", s.O_RATE.UnmarshalText([]byte(fields[3])))
	return errors.Join(errs...)
}

// Sets STAT_GENMPR_data struct's fields
func (s *STAT_GENMPR_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "INDEX", s.INDEX.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "STORM_ID", s.STORM_ID.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "PROB_LEAD", s.PROB_LEAD.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "PROB_VAL", s.PROB_VAL.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "AGEN_INIT", s.AGEN_INIT.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "AGEN_FHR", s.AGEN_FHR.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "AGEN_LAT", s.AGEN_LAT.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "AGEN_LON", s.AGEN_LON.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "AGEN_DLAND", s.AGEN_DLAND.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "BGEN_LAT", s.BGEN_LAT.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "BGEN_LON", s.BGEN_LON.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "BGEN_DLAND", s.BGEN_DLAND.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "GEN_DIST", s.GEN_DIST.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "GEN_TDIFF", s.GEN_TDIFF.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "INIT_TDIFF", s.INIT_TDIFF.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "DEV_CAT", s.DEV_CAT.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "OPS_CAT", s.OPS_CAT.UnmarshalText([]byte(fields[17])))
	return errors.Join(errs...)
}

// Sets STAT_GRAD_data struct's fields
func (s *STAT_GRAD_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "FGBAR", s.FGBAR.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "OGBAR", s.OGBAR.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "MGBAR", s.MGBAR.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "EGBAR", s.EGBAR.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "S1", s.S1.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "S1_OG", s.S1_OG.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FGOG_RATIO", s.FGOG_RATIO.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "DX", s.DX.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "DY", s.DY.UnmarshalText([]byte(fields[9])))
	return errors.Join(errs...)
}

// Sets STAT_ISC_data struct's fields
func (s *STAT_ISC_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "TILE_DIM", s.TILE_DIM.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "TILE_XLL", s.TILE_XLL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "TILE_YLL", s.TILE_YLL.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "NSCALE", s.NSCALE.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "ISCALE", s.ISCALE.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "MSE", s.MSE.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "ISC", s.ISC.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "FENERGY2", s.FENERGY2.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "OENERGY2", s.OENERGY2.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "BASER", s.BASER.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "FBIAS", s.FBIAS.UnmarshalText([]byte(fields[11])))
	return errors.Join(errs...)
}

// Sets STAT_MCTC_data struct's fields
func (s *STAT_MCTC_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	// these values seem to always be ints (or "NA")
	var value validtypes.ValidInt
	count, err := strconv.Atoi(fields[1])
	if err != nil {
		count = 0
	}
	s.CAT = make(map[string]interface{})
	for i1 := 1; i1 <= count; i1++ {
		for i2 := 1; i2 <= count; i2++ {
			// generate the particular key for the map i.e. F1_O1, F1_O2, F1_O3, F1_O4, F2_O1, F2_O2, F2_O3, F2_O4, etc.
			key := fmt.Sprintf("F%d_O%d", i1, i2)
			index := (i1-1)*count + i2
			if index >= len(fields) {
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "CAT", value.UnmarshalText([]byte(fields[index])))
			}
			s.CAT[key] = value
		}
	}
	appendErrorWithContext(&errs, "EC_VALUE", s.EC_VALUE.UnmarshalText([]byte(fields[3])))
	return errors.Join(errs...)
}

// Sets STAT_MCTS_data struct's fields
func (s *STAT_MCTS_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "N_CAT", s.N_CAT.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "ACC", s.ACC.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "ACC_NCL", s.ACC_NCL.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "ACC_NCU", s.ACC_NCU.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "ACC_BCL", s.ACC_BCL.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "ACC_BCU", s.ACC_BCU.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "HK", s.HK.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "HK_BCL", s.HK_BCL.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "HK_BCU", s.HK_BCU.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "HSS", s.HSS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "HSS_BCL", s.HSS_BCL.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "HSS_BCU", s.HSS_BCU.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "GER", s.GER.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "GER_BCL", s.GER_BCL.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "GER_BCU", s.GER_BCU.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "HSS_EC", s.HSS_EC.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "HSS_EC_BCL", s.HSS_EC_BCL.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "HSS_EC_BCU", s.HSS_EC_BCU.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "EC_VALUE", s.EC_VALUE.UnmarshalText([]byte(fields[19])))
	return errors.Join(errs...)
}

// Sets STAT_MPR_data struct's fields
func (s *STAT_MPR_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "INDEX", s.INDEX.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "OBS_SID", s.OBS_SID.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "OBS_LAT", s.OBS_LAT.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "OBS_LON", s.OBS_LON.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "OBS_LVL", s.OBS_LVL.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_ELV", s.OBS_ELV.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FCST", s.FCST.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBS", s.OBS.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "OBS_QC", s.OBS_QC.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "CLIMO_MEAN", s.CLIMO_MEAN.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "CLIMO_STDEV", s.CLIMO_STDEV.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "CLIMO_CDF", s.CLIMO_CDF.UnmarshalText([]byte(fields[12])))
	return errors.Join(errs...)
}

// Sets STAT_NBRCNT_data struct's fields
func (s *STAT_NBRCNT_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "FBS", s.FBS.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "FBS_BCL", s.FBS_BCL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FBS_BCU", s.FBS_BCU.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "FSS", s.FSS.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "FSS_BCL", s.FSS_BCL.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "FSS_BCU", s.FSS_BCU.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "AFSS", s.AFSS.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "AFSS_BCL", s.AFSS_BCL.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "AFSS_BCU", s.AFSS_BCU.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "UFSS", s.UFSS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "UFSS_BCL", s.UFSS_BCL.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "UFSS_BCU", s.UFSS_BCU.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "F_RATE", s.F_RATE.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "F_RATE_BCL", s.F_RATE_BCL.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "F_RATE_BCU", s.F_RATE_BCU.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "O_RATE", s.O_RATE.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "O_RATE_BCL", s.O_RATE_BCL.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "O_RATE_BCU", s.O_RATE_BCU.UnmarshalText([]byte(fields[18])))
	return errors.Join(errs...)
}

// Sets STAT_NBRCTC_data struct's fields
func (s *STAT_NBRCTC_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "FY_OY", s.FY_OY.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "FY_ON", s.FY_ON.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FN_OY", s.FN_OY.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "FN_ON", s.FN_ON.UnmarshalText([]byte(fields[4])))
	return errors.Join(errs...)
}

// Sets STAT_NBRCTS_data struct's fields
func (s *STAT_NBRCTS_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "BASER", s.BASER.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "BASER_NCL", s.BASER_NCL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "BASER_NCU", s.BASER_NCU.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "BASER_BCL", s.BASER_BCL.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "BASER_BCU", s.BASER_BCU.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "FMEAN", s.FMEAN.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FMEAN_NCL", s.FMEAN_NCL.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "FMEAN_NCU", s.FMEAN_NCU.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FMEAN_BCL", s.FMEAN_BCL.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FMEAN_BCU", s.FMEAN_BCU.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "ACC", s.ACC.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "ACC_NCL", s.ACC_NCL.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "ACC_NCU", s.ACC_NCU.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "ACC_BCL", s.ACC_BCL.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "ACC_BCU", s.ACC_BCU.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "FBIAS", s.FBIAS.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "FBIAS_BCL", s.FBIAS_BCL.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "FBIAS_BCU", s.FBIAS_BCU.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "PODY", s.PODY.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "PODY_NCL", s.PODY_NCL.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "PODY_NCU", s.PODY_NCU.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "PODY_BCL", s.PODY_BCL.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "PODY_BCU", s.PODY_BCU.UnmarshalText([]byte(fields[23])))
	appendErrorWithContext(&errs, "PODN", s.PODN.UnmarshalText([]byte(fields[24])))
	appendErrorWithContext(&errs, "PODN_NCL", s.PODN_NCL.UnmarshalText([]byte(fields[25])))
	appendErrorWithContext(&errs, "PODN_NCU", s.PODN_NCU.UnmarshalText([]byte(fields[26])))
	appendErrorWithContext(&errs, "PODN_BCL", s.PODN_BCL.UnmarshalText([]byte(fields[27])))
	appendErrorWithContext(&errs, "PODN_BCU", s.PODN_BCU.UnmarshalText([]byte(fields[28])))
	appendErrorWithContext(&errs, "POFD", s.POFD.UnmarshalText([]byte(fields[29])))
	appendErrorWithContext(&errs, "POFD_NCL", s.POFD_NCL.UnmarshalText([]byte(fields[30])))
	appendErrorWithContext(&errs, "POFD_NCU", s.POFD_NCU.UnmarshalText([]byte(fields[31])))
	appendErrorWithContext(&errs, "POFD_BCL", s.POFD_BCL.UnmarshalText([]byte(fields[32])))
	appendErrorWithContext(&errs, "POFD_BCU", s.POFD_BCU.UnmarshalText([]byte(fields[33])))
	appendErrorWithContext(&errs, "FAR", s.FAR.UnmarshalText([]byte(fields[34])))
	appendErrorWithContext(&errs, "FAR_NCL", s.FAR_NCL.UnmarshalText([]byte(fields[35])))
	appendErrorWithContext(&errs, "FAR_NCU", s.FAR_NCU.UnmarshalText([]byte(fields[36])))
	appendErrorWithContext(&errs, "FAR_BCL", s.FAR_BCL.UnmarshalText([]byte(fields[37])))
	appendErrorWithContext(&errs, "FAR_BCU", s.FAR_BCU.UnmarshalText([]byte(fields[38])))
	appendErrorWithContext(&errs, "CSI", s.CSI.UnmarshalText([]byte(fields[39])))
	appendErrorWithContext(&errs, "CSI_NCL", s.CSI_NCL.UnmarshalText([]byte(fields[40])))
	appendErrorWithContext(&errs, "CSI_NCU", s.CSI_NCU.UnmarshalText([]byte(fields[41])))
	appendErrorWithContext(&errs, "CSI_BCL", s.CSI_BCL.UnmarshalText([]byte(fields[42])))
	appendErrorWithContext(&errs, "CSI_BCU", s.CSI_BCU.UnmarshalText([]byte(fields[43])))
	appendErrorWithContext(&errs, "GSS", s.GSS.UnmarshalText([]byte(fields[44])))
	appendErrorWithContext(&errs, "GSS_BCL", s.GSS_BCL.UnmarshalText([]byte(fields[45])))
	appendErrorWithContext(&errs, "GSS_BCU", s.GSS_BCU.UnmarshalText([]byte(fields[46])))
	appendErrorWithContext(&errs, "HK", s.HK.UnmarshalText([]byte(fields[47])))
	appendErrorWithContext(&errs, "HK_NCL", s.HK_NCL.UnmarshalText([]byte(fields[48])))
	appendErrorWithContext(&errs, "HK_NCU", s.HK_NCU.UnmarshalText([]byte(fields[49])))
	appendErrorWithContext(&errs, "HK_BCL", s.HK_BCL.UnmarshalText([]byte(fields[50])))
	appendErrorWithContext(&errs, "HK_BCU", s.HK_BCU.UnmarshalText([]byte(fields[51])))
	appendErrorWithContext(&errs, "HSS", s.HSS.UnmarshalText([]byte(fields[52])))
	appendErrorWithContext(&errs, "HSS_BCL", s.HSS_BCL.UnmarshalText([]byte(fields[53])))
	appendErrorWithContext(&errs, "HSS_BCU", s.HSS_BCU.UnmarshalText([]byte(fields[54])))
	appendErrorWithContext(&errs, "ODDS", s.ODDS.UnmarshalText([]byte(fields[55])))
	appendErrorWithContext(&errs, "ODDS_NCL", s.ODDS_NCL.UnmarshalText([]byte(fields[56])))
	appendErrorWithContext(&errs, "ODDS_NCU", s.ODDS_NCU.UnmarshalText([]byte(fields[57])))
	appendErrorWithContext(&errs, "ODDS_BCL", s.ODDS_BCL.UnmarshalText([]byte(fields[58])))
	appendErrorWithContext(&errs, "ODDS_BCU", s.ODDS_BCU.UnmarshalText([]byte(fields[59])))
	appendErrorWithContext(&errs, "LODDS", s.LODDS.UnmarshalText([]byte(fields[60])))
	appendErrorWithContext(&errs, "LODDS_NCL", s.LODDS_NCL.UnmarshalText([]byte(fields[61])))
	appendErrorWithContext(&errs, "LODDS_NCU", s.LODDS_NCU.UnmarshalText([]byte(fields[62])))
	appendErrorWithContext(&errs, "LODDS_BCL", s.LODDS_BCL.UnmarshalText([]byte(fields[63])))
	appendErrorWithContext(&errs, "LODDS_BCU", s.LODDS_BCU.UnmarshalText([]byte(fields[64])))
	appendErrorWithContext(&errs, "ORSS", s.ORSS.UnmarshalText([]byte(fields[65])))
	appendErrorWithContext(&errs, "ORSS_NCL", s.ORSS_NCL.UnmarshalText([]byte(fields[66])))
	appendErrorWithContext(&errs, "ORSS_NCU", s.ORSS_NCU.UnmarshalText([]byte(fields[67])))
	appendErrorWithContext(&errs, "ORSS_BCL", s.ORSS_BCL.UnmarshalText([]byte(fields[68])))
	appendErrorWithContext(&errs, "ORSS_BCU", s.ORSS_BCU.UnmarshalText([]byte(fields[69])))
	appendErrorWithContext(&errs, "EDS", s.EDS.UnmarshalText([]byte(fields[70])))
	appendErrorWithContext(&errs, "EDS_NCL", s.EDS_NCL.UnmarshalText([]byte(fields[71])))
	appendErrorWithContext(&errs, "EDS_NCU", s.EDS_NCU.UnmarshalText([]byte(fields[72])))
	appendErrorWithContext(&errs, "EDS_BCL", s.EDS_BCL.UnmarshalText([]byte(fields[73])))
	appendErrorWithContext(&errs, "EDS_BCU", s.EDS_BCU.UnmarshalText([]byte(fields[74])))
	appendErrorWithContext(&errs, "SEDS", s.SEDS.UnmarshalText([]byte(fields[75])))
	appendErrorWithContext(&errs, "SEDS_NCL", s.SEDS_NCL.UnmarshalText([]byte(fields[76])))
	appendErrorWithContext(&errs, "SEDS_NCU", s.SEDS_NCU.UnmarshalText([]byte(fields[77])))
	appendErrorWithContext(&errs, "SEDS_BCL", s.SEDS_BCL.UnmarshalText([]byte(fields[78])))
	appendErrorWithContext(&errs, "SEDS_BCU", s.SEDS_BCU.UnmarshalText([]byte(fields[79])))
	appendErrorWithContext(&errs, "EDI", s.EDI.UnmarshalText([]byte(fields[80])))
	appendErrorWithContext(&errs, "EDI_NCL", s.EDI_NCL.UnmarshalText([]byte(fields[81])))
	appendErrorWithContext(&errs, "EDI_NCU", s.EDI_NCU.UnmarshalText([]byte(fields[82])))
	appendErrorWithContext(&errs, "EDI_BCL", s.EDI_BCL.UnmarshalText([]byte(fields[83])))
	appendErrorWithContext(&errs, "EDI_BCU", s.EDI_BCU.UnmarshalText([]byte(fields[84])))
	appendErrorWithContext(&errs, "SEDI", s.SEDI.UnmarshalText([]byte(fields[85])))
	appendErrorWithContext(&errs, "SEDI_NCL", s.SEDI_NCL.UnmarshalText([]byte(fields[86])))
	appendErrorWithContext(&errs, "SEDI_NCU", s.SEDI_NCU.UnmarshalText([]byte(fields[87])))
	appendErrorWithContext(&errs, "SEDI_BCL", s.SEDI_BCL.UnmarshalText([]byte(fields[88])))
	appendErrorWithContext(&errs, "SEDI_BCU", s.SEDI_BCU.UnmarshalText([]byte(fields[89])))
	appendErrorWithContext(&errs, "BAGSS", s.BAGSS.UnmarshalText([]byte(fields[90])))
	appendErrorWithContext(&errs, "BAGSS_BCL", s.BAGSS_BCL.UnmarshalText([]byte(fields[91])))
	appendErrorWithContext(&errs, "BAGSS_BCU", s.BAGSS_BCU.UnmarshalText([]byte(fields[92])))
	return errors.Join(errs...)
}

// Sets STAT_ORANK_data struct's fields
func (s *STAT_ORANK_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "INDEX", s.INDEX.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "OBS_SID", s.OBS_SID.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "OBS_LAT", s.OBS_LAT.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "OBS_LON", s.OBS_LON.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "OBS_LVL", s.OBS_LVL.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBS_ELV", s.OBS_ELV.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "OBS", s.OBS.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "PIT", s.PIT.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "RANK", s.RANK.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "N_ENS_VLD", s.N_ENS_VLD.UnmarshalText([]byte(fields[10])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidInt
	count, err := strconv.Atoi(fields[11])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"ENS_"}
	s.ENS = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 12; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "ENS", value.UnmarshalText([]byte(fields[index])))
			}
			s.ENS[key] = value
		}
	}
	appendErrorWithContext(&errs, "OBS_QC", s.OBS_QC.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "ENS_MEAN", s.ENS_MEAN.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "CLIMO_MEAN", s.CLIMO_MEAN.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "SPREAD", s.SPREAD.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "ENS_MEAN_OERR", s.ENS_MEAN_OERR.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "SPREAD_OERR", s.SPREAD_OERR.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "SPREAD_PLUS_OERR", s.SPREAD_PLUS_OERR.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "CLIMO_STDEV", s.CLIMO_STDEV.UnmarshalText([]byte(fields[20])))
	return errors.Join(errs...)
}

// Sets STAT_PCT_data struct's fields
func (s *STAT_PCT_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidFloat
	count, err := strconv.Atoi(fields[1])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"THRESH_", "OY_", "ON_"}
	s.THRESH = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 2; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "THRESH", value.UnmarshalText([]byte(fields[index])))
			}
			s.THRESH[key] = value
		}
	}
	return errors.Join(errs...)
}

// Sets STAT_PHIST_data struct's fields
func (s *STAT_PHIST_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "BIN_SIZE", s.BIN_SIZE.UnmarshalText([]byte(fields[1])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidInt
	count, err := strconv.Atoi(fields[2])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"BIN_"}
	s.BIN = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 3; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "BIN", value.UnmarshalText([]byte(fields[index])))
			}
			s.BIN[key] = value
		}
	}
	return errors.Join(errs...)
}

// Sets STAT_PJC_data struct's fields
func (s *STAT_PJC_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidFloat
	count, err := strconv.Atoi(fields[1])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"THRESH_", "OY_TP_", "ON_TP_", "CALIBRATION_", "REFINEMENT", "LIKELIHOOD_", "BASER_"}
	s.THRESH = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 2; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "THRESH", value.UnmarshalText([]byte(fields[index])))
			}
			s.THRESH[key] = value
		}
	}
	return errors.Join(errs...)
}

// Sets STAT_PRC_data struct's fields
func (s *STAT_PRC_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidFloat
	count, err := strconv.Atoi(fields[1])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"THRESH_", "PODY_", "POFD_"}
	s.THRESH = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 2; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "THRESH", value.UnmarshalText([]byte(fields[index])))
			}
			s.THRESH[key] = value
		}
	}
	return errors.Join(errs...)
}

// Sets STAT_PSTD_data struct's fields
func (s *STAT_PSTD_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidFloat
	count, err := strconv.Atoi(fields[1])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"THRESH_"}
	s.THRESH = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 2; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "THRESH", value.UnmarshalText([]byte(fields[index])))
			}
			s.THRESH[key] = value
		}
	}
	appendErrorWithContext(&errs, "BASER_NCL", s.BASER_NCL.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "BASER_NCU", s.BASER_NCU.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "RELIABILITY", s.RELIABILITY.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "RESOLUTION", s.RESOLUTION.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "UNCERTAINTY", s.UNCERTAINTY.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "ROC_AUC", s.ROC_AUC.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "BRIER", s.BRIER.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "BRIER_NCL", s.BRIER_NCL.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "BRIER_NCU", s.BRIER_NCU.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "BRIERCL", s.BRIERCL.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "BRIERCL_NCL", s.BRIERCL_NCL.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "BRIERCL_NCU", s.BRIERCL_NCU.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "BSS", s.BSS.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "BSS_SMPL", s.BSS_SMPL.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "THRESH_I", s.THRESH_I.UnmarshalText([]byte(fields[17])))
	return errors.Join(errs...)
}

// Sets STAT_RELP_data struct's fields
func (s *STAT_RELP_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidFloat
	count, err := strconv.Atoi(fields[1])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"RELP_"}
	s.ENS = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 2; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "ENS", value.UnmarshalText([]byte(fields[index])))
			}
			s.ENS[key] = value
		}
	}
	return errors.Join(errs...)
}

// Sets STAT_RHIST_data struct's fields
func (s *STAT_RHIST_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidInt
	count, err := strconv.Atoi(fields[1])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"RANK_"}
	s.RANK = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 2; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "RANK", value.UnmarshalText([]byte(fields[index])))
			}
			s.RANK[key] = value
		}
	}
	return errors.Join(errs...)
}

// Sets STAT_RPS_data struct's fields
func (s *STAT_RPS_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "N_PROB", s.N_PROB.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "RPS_REL", s.RPS_REL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "RPS_RES", s.RPS_RES.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "RPS_UNC", s.RPS_UNC.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "RPS", s.RPS.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "RPSS", s.RPSS.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "RPSS_SMPL", s.RPSS_SMPL.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "RPS_COMP", s.RPS_COMP.UnmarshalText([]byte(fields[8])))
	return errors.Join(errs...)
}

// Sets STAT_SAL1L2_data struct's fields
func (s *STAT_SAL1L2_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "FABAR", s.FABAR.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "OABAR", s.OABAR.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FOABAR", s.FOABAR.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "FFABAR", s.FFABAR.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "OOABAR", s.OOABAR.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "MAE", s.MAE.UnmarshalText([]byte(fields[6])))
	return errors.Join(errs...)
}

// Sets STAT_SL1L2_data struct's fields
func (s *STAT_SL1L2_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "FBAR", s.FBAR.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "OBAR", s.OBAR.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FOBAR", s.FOBAR.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "FFBAR", s.FFBAR.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "OOBAR", s.OOBAR.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "MAE", s.MAE.UnmarshalText([]byte(fields[6])))
	return errors.Join(errs...)
}

// Sets STAT_SSIDX_data struct's fields
func (s *STAT_SSIDX_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "FCST_MODEL", s.FCST_MODEL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "REF_MODEL", s.REF_MODEL.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "N_INIT", s.N_INIT.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "N_TERM", s.N_TERM.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "N_VLD", s.N_VLD.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "SS_INDEX", s.SS_INDEX.UnmarshalText([]byte(fields[5])))
	return errors.Join(errs...)
}

// Sets STAT_SSVAR_data struct's fields
func (s *STAT_SSVAR_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "N_BIN", s.N_BIN.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "BIN_I", s.BIN_I.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "BIN_N", s.BIN_N.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "VAR_MIN", s.VAR_MIN.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "VAR_MAX", s.VAR_MAX.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "VAR_MEAN", s.VAR_MEAN.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FBAR", s.FBAR.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "OBAR", s.OBAR.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FOBAR", s.FOBAR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "FFBAR", s.FFBAR.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "OOBAR", s.OOBAR.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "FBAR_NCL", s.FBAR_NCL.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "FBAR_NCU", s.FBAR_NCU.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "FSTDEV", s.FSTDEV.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "FSTDEV_NCL", s.FSTDEV_NCL.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "FSTDEV_NCU", s.FSTDEV_NCU.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "OBAR_NCL", s.OBAR_NCL.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "OBAR_NCU", s.OBAR_NCU.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "OSTDEV", s.OSTDEV.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "OSTDEV_NCL", s.OSTDEV_NCL.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "OSTDEV_NCU", s.OSTDEV_NCU.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "PR_CORR", s.PR_CORR.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "PR_CORR_NCL", s.PR_CORR_NCL.UnmarshalText([]byte(fields[23])))
	appendErrorWithContext(&errs, "PR_CORR_NCU", s.PR_CORR_NCU.UnmarshalText([]byte(fields[24])))
	appendErrorWithContext(&errs, "ME", s.ME.UnmarshalText([]byte(fields[25])))
	appendErrorWithContext(&errs, "ME_NCL", s.ME_NCL.UnmarshalText([]byte(fields[26])))
	appendErrorWithContext(&errs, "ME_NCU", s.ME_NCU.UnmarshalText([]byte(fields[27])))
	appendErrorWithContext(&errs, "ESTDEV", s.ESTDEV.UnmarshalText([]byte(fields[28])))
	appendErrorWithContext(&errs, "ESTDEV_NCL", s.ESTDEV_NCL.UnmarshalText([]byte(fields[29])))
	appendErrorWithContext(&errs, "ESTDEV_NCU", s.ESTDEV_NCU.UnmarshalText([]byte(fields[30])))
	appendErrorWithContext(&errs, "MBIAS", s.MBIAS.UnmarshalText([]byte(fields[31])))
	appendErrorWithContext(&errs, "MSE", s.MSE.UnmarshalText([]byte(fields[32])))
	appendErrorWithContext(&errs, "BCMSE", s.BCMSE.UnmarshalText([]byte(fields[33])))
	appendErrorWithContext(&errs, "RMSE", s.RMSE.UnmarshalText([]byte(fields[34])))
	return errors.Join(errs...)
}

// Sets STAT_VAL1L2_data struct's fields
func (s *STAT_VAL1L2_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "UFABAR", s.UFABAR.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "VFABAR", s.VFABAR.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "UOABAR", s.UOABAR.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "VOABAR", s.VOABAR.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "UVFOABAR", s.UVFOABAR.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "UVFFABAR", s.UVFFABAR.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "UVOOABAR", s.UVOOABAR.UnmarshalText([]byte(fields[7])))
	return errors.Join(errs...)
}

// Sets STAT_VCNT_data struct's fields
func (s *STAT_VCNT_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "FBAR", s.FBAR.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "FBAR_BCL", s.FBAR_BCL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "FBAR_BCU", s.FBAR_BCU.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "OBAR", s.OBAR.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "OBAR_BCL", s.OBAR_BCL.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "OBAR_BCU", s.OBAR_BCU.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "FS_RMS", s.FS_RMS.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "FS_RMS_BCL", s.FS_RMS_BCL.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "FS_RMS_BCU", s.FS_RMS_BCU.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "OS_RMS", s.OS_RMS.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "OS_RMS_BCL", s.OS_RMS_BCL.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "OS_RMS_BCU", s.OS_RMS_BCU.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "MSVE", s.MSVE.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "MSVE_BCL", s.MSVE_BCL.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "MSVE_BCU", s.MSVE_BCU.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "RMSVE", s.RMSVE.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "RMSVE_BCL", s.RMSVE_BCL.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "RMSVE_BCU", s.RMSVE_BCU.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "FSTDEV", s.FSTDEV.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "FSTDEV_BCL", s.FSTDEV_BCL.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "FSTDEV_BCU", s.FSTDEV_BCU.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "OSTDEV", s.OSTDEV.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "OSTDEV_BCL", s.OSTDEV_BCL.UnmarshalText([]byte(fields[23])))
	appendErrorWithContext(&errs, "OSTDEV_BCU", s.OSTDEV_BCU.UnmarshalText([]byte(fields[24])))
	appendErrorWithContext(&errs, "FDIR", s.FDIR.UnmarshalText([]byte(fields[25])))
	appendErrorWithContext(&errs, "FDIR_BCL", s.FDIR_BCL.UnmarshalText([]byte(fields[26])))
	appendErrorWithContext(&errs, "FDIR_BCU", s.FDIR_BCU.UnmarshalText([]byte(fields[27])))
	appendErrorWithContext(&errs, "ODIR", s.ODIR.UnmarshalText([]byte(fields[28])))
	appendErrorWithContext(&errs, "ODIR_BCL", s.ODIR_BCL.UnmarshalText([]byte(fields[29])))
	appendErrorWithContext(&errs, "ODIR_BCU", s.ODIR_BCU.UnmarshalText([]byte(fields[30])))
	appendErrorWithContext(&errs, "FBAR_SPEED", s.FBAR_SPEED.UnmarshalText([]byte(fields[31])))
	appendErrorWithContext(&errs, "FBAR_SPEED_BCL", s.FBAR_SPEED_BCL.UnmarshalText([]byte(fields[32])))
	appendErrorWithContext(&errs, "FBAR_SPEED_BCU", s.FBAR_SPEED_BCU.UnmarshalText([]byte(fields[33])))
	appendErrorWithContext(&errs, "OBAR_SPEED", s.OBAR_SPEED.UnmarshalText([]byte(fields[34])))
	appendErrorWithContext(&errs, "OBAR_SPEED_BCL", s.OBAR_SPEED_BCL.UnmarshalText([]byte(fields[35])))
	appendErrorWithContext(&errs, "OBAR_SPEED_BCU", s.OBAR_SPEED_BCU.UnmarshalText([]byte(fields[36])))
	appendErrorWithContext(&errs, "VDIFF_SPEED", s.VDIFF_SPEED.UnmarshalText([]byte(fields[37])))
	appendErrorWithContext(&errs, "VDIFF_SPEED_BCL", s.VDIFF_SPEED_BCL.UnmarshalText([]byte(fields[38])))
	appendErrorWithContext(&errs, "VDIFF_SPEED_BCU", s.VDIFF_SPEED_BCU.UnmarshalText([]byte(fields[39])))
	appendErrorWithContext(&errs, "VDIFF_DIR", s.VDIFF_DIR.UnmarshalText([]byte(fields[40])))
	appendErrorWithContext(&errs, "VDIFF_DIR_BCL", s.VDIFF_DIR_BCL.UnmarshalText([]byte(fields[41])))
	appendErrorWithContext(&errs, "VDIFF_DIR_BCU", s.VDIFF_DIR_BCU.UnmarshalText([]byte(fields[42])))
	appendErrorWithContext(&errs, "SPEED_ERR", s.SPEED_ERR.UnmarshalText([]byte(fields[43])))
	appendErrorWithContext(&errs, "SPEED_ERR_BCL", s.SPEED_ERR_BCL.UnmarshalText([]byte(fields[44])))
	appendErrorWithContext(&errs, "SPEED_ERR_BCU", s.SPEED_ERR_BCU.UnmarshalText([]byte(fields[45])))
	appendErrorWithContext(&errs, "SPEED_ABSERR", s.SPEED_ABSERR.UnmarshalText([]byte(fields[46])))
	appendErrorWithContext(&errs, "SPEED_ABSERR_BCL", s.SPEED_ABSERR_BCL.UnmarshalText([]byte(fields[47])))
	appendErrorWithContext(&errs, "SPEED_ABSERR_BCU", s.SPEED_ABSERR_BCU.UnmarshalText([]byte(fields[48])))
	appendErrorWithContext(&errs, "DIR_ERR", s.DIR_ERR.UnmarshalText([]byte(fields[49])))
	appendErrorWithContext(&errs, "DIR_ERR_BCL", s.DIR_ERR_BCL.UnmarshalText([]byte(fields[50])))
	appendErrorWithContext(&errs, "DIR_ERR_BCU", s.DIR_ERR_BCU.UnmarshalText([]byte(fields[51])))
	appendErrorWithContext(&errs, "DIR_ABSERR", s.DIR_ABSERR.UnmarshalText([]byte(fields[52])))
	appendErrorWithContext(&errs, "DIR_ABSERR_BCL", s.DIR_ABSERR_BCL.UnmarshalText([]byte(fields[53])))
	appendErrorWithContext(&errs, "DIR_ABSERR_BCU", s.DIR_ABSERR_BCU.UnmarshalText([]byte(fields[54])))
	return errors.Join(errs...)
}

// Sets STAT_VL1L2_data struct's fields
func (s *STAT_VL1L2_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "UFBAR", s.UFBAR.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "VFBAR", s.VFBAR.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "UOBAR", s.UOBAR.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "VOBAR", s.VOBAR.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "UVFOBAR", s.UVFOBAR.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "UVFFBAR", s.UVFFBAR.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "UVOOBAR", s.UVOOBAR.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "F_SPEED_BAR", s.F_SPEED_BAR.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "O_SPEED_BAR", s.O_SPEED_BAR.UnmarshalText([]byte(fields[9])))
	return errors.Join(errs...)
}

// Sets TCST_PROBRIRW_data struct's fields
func (s *TCST_PROBRIRW_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "ALAT", s.ALAT.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "ALON", s.ALON.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "BLAT", s.BLAT.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "BLON", s.BLON.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "INITIALS", s.INITIALS.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "TK_ERR", s.TK_ERR.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "X_ERR", s.X_ERR.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "Y_ERR", s.Y_ERR.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "ADLAND", s.ADLAND.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "BDLAND", s.BDLAND.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "RIRW_BEG", s.RIRW_BEG.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "RIRW_END", s.RIRW_END.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "RIRW_WINDOW", s.RIRW_WINDOW.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "AWIND_END", s.AWIND_END.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "BWIND_BEG", s.BWIND_BEG.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "BWIND_END", s.BWIND_END.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "BDELTA", s.BDELTA.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "BDELTA_MAX", s.BDELTA_MAX.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "BLEVEL_BEG", s.BLEVEL_BEG.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "BLEVEL_END", s.BLEVEL_END.UnmarshalText([]byte(fields[19])))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidInt
	count, err := strconv.Atoi(fields[20])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"THRESH_", "PROB_"}
	s.THRESH = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 21; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				appendErrorWithContext(&errs, "THRESH", value.UnmarshalText([]byte(fields[index])))
			}
			s.THRESH[key] = value
		}
	}
	appendErrorWithContext(&errs, "INIT", s.INIT.UnmarshalText([]byte(fields[23])))
	return errors.Join(errs...)
}

// Sets TCST_TCMPR_data struct's fields
func (s *TCST_TCMPR_data) fill(fields []string) error {
	var errs []error
	appendErrorWithContext(&errs, "TOTAL", s.TOTAL.UnmarshalText([]byte(fields[0])))
	appendErrorWithContext(&errs, "INDEX", s.INDEX.UnmarshalText([]byte(fields[1])))
	appendErrorWithContext(&errs, "LEVEL", s.LEVEL.UnmarshalText([]byte(fields[2])))
	appendErrorWithContext(&errs, "WATCH_WARN", s.WATCH_WARN.UnmarshalText([]byte(fields[3])))
	appendErrorWithContext(&errs, "INITIALS", s.INITIALS.UnmarshalText([]byte(fields[4])))
	appendErrorWithContext(&errs, "ALAT", s.ALAT.UnmarshalText([]byte(fields[5])))
	appendErrorWithContext(&errs, "ALON", s.ALON.UnmarshalText([]byte(fields[6])))
	appendErrorWithContext(&errs, "BLAT", s.BLAT.UnmarshalText([]byte(fields[7])))
	appendErrorWithContext(&errs, "BLON", s.BLON.UnmarshalText([]byte(fields[8])))
	appendErrorWithContext(&errs, "TK_ERR", s.TK_ERR.UnmarshalText([]byte(fields[9])))
	appendErrorWithContext(&errs, "X_ERR", s.X_ERR.UnmarshalText([]byte(fields[10])))
	appendErrorWithContext(&errs, "Y_ERR", s.Y_ERR.UnmarshalText([]byte(fields[11])))
	appendErrorWithContext(&errs, "ALTK_ERR", s.ALTK_ERR.UnmarshalText([]byte(fields[12])))
	appendErrorWithContext(&errs, "CRTK_ERR", s.CRTK_ERR.UnmarshalText([]byte(fields[13])))
	appendErrorWithContext(&errs, "ADLAND", s.ADLAND.UnmarshalText([]byte(fields[14])))
	appendErrorWithContext(&errs, "BDLAND", s.BDLAND.UnmarshalText([]byte(fields[15])))
	appendErrorWithContext(&errs, "AMSLP", s.AMSLP.UnmarshalText([]byte(fields[16])))
	appendErrorWithContext(&errs, "BMSLP", s.BMSLP.UnmarshalText([]byte(fields[17])))
	appendErrorWithContext(&errs, "AMAX_WIND", s.AMAX_WIND.UnmarshalText([]byte(fields[18])))
	appendErrorWithContext(&errs, "BMAX_WIND", s.BMAX_WIND.UnmarshalText([]byte(fields[19])))
	appendErrorWithContext(&errs, "AAL_WIND_34", s.AAL_WIND_34.UnmarshalText([]byte(fields[20])))
	appendErrorWithContext(&errs, "BAL_WIND_34", s.BAL_WIND_34.UnmarshalText([]byte(fields[21])))
	appendErrorWithContext(&errs, "ANE_WIND_34", s.ANE_WIND_34.UnmarshalText([]byte(fields[22])))
	appendErrorWithContext(&errs, "BNE_WIND_34", s.BNE_WIND_34.UnmarshalText([]byte(fields[23])))
	appendErrorWithContext(&errs, "ASE_WIND_34", s.ASE_WIND_34.UnmarshalText([]byte(fields[24])))
	appendErrorWithContext(&errs, "BSE_WIND_34", s.BSE_WIND_34.UnmarshalText([]byte(fields[25])))
	appendErrorWithContext(&errs, "ASW_WIND_34", s.ASW_WIND_34.UnmarshalText([]byte(fields[26])))
	appendErrorWithContext(&errs, "BSW_WIND_34", s.BSW_WIND_34.UnmarshalText([]byte(fields[27])))
	appendErrorWithContext(&errs, "ANW_WIND_34", s.ANW_WIND_34.UnmarshalText([]byte(fields[28])))
	appendErrorWithContext(&errs, "BNW_WIND_34", s.BNW_WIND_34.UnmarshalText([]byte(fields[29])))
	appendErrorWithContext(&errs, "AAL_WIND_50", s.AAL_WIND_50.UnmarshalText([]byte(fields[30])))
	appendErrorWithContext(&errs, "BAL_WIND_50", s.BAL_WIND_50.UnmarshalText([]byte(fields[31])))
	appendErrorWithContext(&errs, "ANE_WIND_50", s.ANE_WIND_50.UnmarshalText([]byte(fields[32])))
	appendErrorWithContext(&errs, "BNE_WIND_50", s.BNE_WIND_50.UnmarshalText([]byte(fields[33])))
	appendErrorWithContext(&errs, "ASE_WIND_50", s.ASE_WIND_50.UnmarshalText([]byte(fields[34])))
	appendErrorWithContext(&errs, "BSE_WIND_50", s.BSE_WIND_50.UnmarshalText([]byte(fields[35])))
	appendErrorWithContext(&errs, "ASW_WIND_50", s.ASW_WIND_50.UnmarshalText([]byte(fields[36])))
	appendErrorWithContext(&errs, "BSW_WIND_50", s.BSW_WIND_50.UnmarshalText([]byte(fields[37])))
	appendErrorWithContext(&errs, "ANW_WIND_50", s.ANW_WIND_50.UnmarshalText([]byte(fields[38])))
	appendErrorWithContext(&errs, "BNW_WIND_50", s.BNW_WIND_50.UnmarshalText([]byte(fields[39])))
	appendErrorWithContext(&errs, "AAL_WIND_64", s.AAL_WIND_64.UnmarshalText([]byte(fields[40])))
	appendErrorWithContext(&errs, "BAL_WIND_64", s.BAL_WIND_64.UnmarshalText([]byte(fields[41])))
	appendErrorWithContext(&errs, "ANE_WIND_64", s.ANE_WIND_64.UnmarshalText([]byte(fields[42])))
	appendErrorWithContext(&errs, "BNE_WIND_64", s.BNE_WIND_64.UnmarshalText([]byte(fields[43])))
	appendErrorWithContext(&errs, "ASE_WIND_64", s.ASE_WIND_64.UnmarshalText([]byte(fields[44])))
	appendErrorWithContext(&errs, "BSE_WIND_64", s.BSE_WIND_64.UnmarshalText([]byte(fields[45])))
	appendErrorWithContext(&errs, "ASW_WIND_64", s.ASW_WIND_64.UnmarshalText([]byte(fields[46])))
	appendErrorWithContext(&errs, "BSW_WIND_64", s.BSW_WIND_64.UnmarshalText([]byte(fields[47])))
	appendErrorWithContext(&errs, "ANW_WIND_64", s.ANW_WIND_64.UnmarshalText([]byte(fields[48])))
	appendErrorWithContext(&errs, "BNW_WIND_64", s.BNW_WIND_64.UnmarshalText([]byte(fields[49])))
	appendErrorWithContext(&errs, "ARADP", s.ARADP.UnmarshalText([]byte(fields[50])))
	appendErrorWithContext(&errs, "BRADP", s.BRADP.UnmarshalText([]byte(fields[51])))
	appendErrorWithContext(&errs, "ARRP", s.ARRP.UnmarshalText([]byte(fields[52])))
	appendErrorWithContext(&errs, "BRRP", s.BRRP.UnmarshalText([]byte(fields[53])))
	appendErrorWithContext(&errs, "AMRD", s.AMRD.UnmarshalText([]byte(fields[54])))
	appendErrorWithContext(&errs, "BMRD", s.BMRD.UnmarshalText([]byte(fields[55])))
	appendErrorWithContext(&errs, "AGUSTS", s.AGUSTS.UnmarshalText([]byte(fields[56])))
	appendErrorWithContext(&errs, "BGUSTS", s.BGUSTS.UnmarshalText([]byte(fields[57])))
	appendErrorWithContext(&errs, "AEYE", s.AEYE.UnmarshalText([]byte(fields[58])))
	appendErrorWithContext(&errs, "BEYE", s.BEYE.UnmarshalText([]byte(fields[59])))
	appendErrorWithContext(&errs, "ADIR", s.ADIR.UnmarshalText([]byte(fields[60])))
	appendErrorWithContext(&errs, "BDIR", s.BDIR.UnmarshalText([]byte(fields[61])))
	appendErrorWithContext(&errs, "ASPEED", s.ASPEED.UnmarshalText([]byte(fields[62])))
	appendErrorWithContext(&errs, "BSPEED", s.BSPEED.UnmarshalText([]byte(fields[63])))
	appendErrorWithContext(&errs, "ADEPTH", s.ADEPTH.UnmarshalText([]byte(fields[64])))
	appendErrorWithContext(&errs, "BDEPTH", s.BDEPTH.UnmarshalText([]byte(fields[65])))
	appendErrorWithContext(&errs, "INIT", s.INIT.UnmarshalText([]byte(fields[66])))
	return errors.Join(errs...)
}

// Creates a new doc, header functions and all.
func GetDocForId(fileLineType string, metaData util.VxMetadata, headerData []string, dataData []string, dataKey string) (util.METdocument, error) {
	var statDoc util.METdocument
	var errs []error

	switch fileLineType {
	case "MODE_CTS":
		elem_header := MODE_CTS_header{}
		appendErrorWithContext(&errs, "MODE_CTS_header", elem_header.fill(headerData))
		elem_data := MODE_CTS_data{}
		appendErrorWithContext(&errs, "MODE_CTS_data", elem_data.fill(dataData))

		tmp := MODE_CTS{
			VxMetadata:      metaData,
			MODE_CTS_header: elem_header,
			Data:            make(map[string]MODE_CTS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "MODE_OBJ":
		elem_header := MODE_OBJ_header{}
		appendErrorWithContext(&errs, "MODE_OBJ_header", elem_header.fill(headerData))
		elem_data := MODE_OBJ_data{}
		appendErrorWithContext(&errs, "MODE_OBJ_data", elem_data.fill(dataData))

		tmp := MODE_OBJ{
			VxMetadata:      metaData,
			MODE_OBJ_header: elem_header,
			Data:            make(map[string]MODE_OBJ_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_CNT":
		elem_header := STAT_CNT_header{}
		appendErrorWithContext(&errs, "STAT_CNT_header", elem_header.fill(headerData))
		elem_data := STAT_CNT_data{}
		appendErrorWithContext(&errs, "STAT_CNT_data", elem_data.fill(dataData))

		tmp := STAT_CNT{
			VxMetadata:      metaData,
			STAT_CNT_header: elem_header,
			Data:            make(map[string]STAT_CNT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_CTC":
		elem_header := STAT_CTC_header{}
		appendErrorWithContext(&errs, "STAT_CTC_header", elem_header.fill(headerData))
		elem_data := STAT_CTC_data{}
		appendErrorWithContext(&errs, "STAT_CTC_data", elem_data.fill(dataData))

		tmp := STAT_CTC{
			VxMetadata:      metaData,
			STAT_CTC_header: elem_header,
			Data:            make(map[string]STAT_CTC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_CTS":
		elem_header := STAT_CTS_header{}
		appendErrorWithContext(&errs, "STAT_CTS_header", elem_header.fill(headerData))
		elem_data := STAT_CTS_data{}
		appendErrorWithContext(&errs, "STAT_CTS_data", elem_data.fill(dataData))

		tmp := STAT_CTS{
			VxMetadata:      metaData,
			STAT_CTS_header: elem_header,
			Data:            make(map[string]STAT_CTS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_DMAP":
		elem_header := STAT_DMAP_header{}
		appendErrorWithContext(&errs, "STAT_DMAP_header", elem_header.fill(headerData))
		elem_data := STAT_DMAP_data{}
		appendErrorWithContext(&errs, "STAT_DMAP_data", elem_data.fill(dataData))

		tmp := STAT_DMAP{
			VxMetadata:       metaData,
			STAT_DMAP_header: elem_header,
			Data:             make(map[string]STAT_DMAP_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_ECLV":
		elem_header := STAT_ECLV_header{}
		appendErrorWithContext(&errs, "STAT_ECLV_header", elem_header.fill(headerData))
		elem_data := STAT_ECLV_data{}
		appendErrorWithContext(&errs, "STAT_ECLV_data", elem_data.fill(dataData))

		tmp := STAT_ECLV{
			VxMetadata:       metaData,
			STAT_ECLV_header: elem_header,
			Data:             make(map[string]STAT_ECLV_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_ECNT":
		elem_header := STAT_ECNT_header{}
		appendErrorWithContext(&errs, "STAT_ECNT_header", elem_header.fill(headerData))
		elem_data := STAT_ECNT_data{}
		appendErrorWithContext(&errs, "STAT_ECNT_data", elem_data.fill(dataData))

		tmp := STAT_ECNT{
			VxMetadata:       metaData,
			STAT_ECNT_header: elem_header,
			Data:             make(map[string]STAT_ECNT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_FHO":
		elem_header := STAT_FHO_header{}
		appendErrorWithContext(&errs, "STAT_FHO_header", elem_header.fill(headerData))
		elem_data := STAT_FHO_data{}
		appendErrorWithContext(&errs, "STAT_FHO_data", elem_data.fill(dataData))

		tmp := STAT_FHO{
			VxMetadata:      metaData,
			STAT_FHO_header: elem_header,
			Data:            make(map[string]STAT_FHO_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_GENMPR":
		elem_header := STAT_GENMPR_header{}
		appendErrorWithContext(&errs, "STAT_GENMPR_header", elem_header.fill(headerData))
		elem_data := STAT_GENMPR_data{}
		appendErrorWithContext(&errs, "STAT_GENMPR_data", elem_data.fill(dataData))

		tmp := STAT_GENMPR{
			VxMetadata:         metaData,
			STAT_GENMPR_header: elem_header,
			Data:               make(map[string]STAT_GENMPR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_GRAD":
		elem_header := STAT_GRAD_header{}
		appendErrorWithContext(&errs, "STAT_GRAD_header", elem_header.fill(headerData))
		elem_data := STAT_GRAD_data{}
		appendErrorWithContext(&errs, "STAT_GRAD_data", elem_data.fill(dataData))

		tmp := STAT_GRAD{
			VxMetadata:       metaData,
			STAT_GRAD_header: elem_header,
			Data:             make(map[string]STAT_GRAD_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_ISC":
		elem_header := STAT_ISC_header{}
		appendErrorWithContext(&errs, "STAT_ISC_header", elem_header.fill(headerData))
		elem_data := STAT_ISC_data{}
		appendErrorWithContext(&errs, "STAT_ISC_data", elem_data.fill(dataData))

		tmp := STAT_ISC{
			VxMetadata:      metaData,
			STAT_ISC_header: elem_header,
			Data:            make(map[string]STAT_ISC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_MCTC":
		elem_header := STAT_MCTC_header{}
		appendErrorWithContext(&errs, "STAT_MCTC_header", elem_header.fill(headerData))
		elem_data := STAT_MCTC_data{}
		appendErrorWithContext(&errs, "STAT_MCTC_data", elem_data.fill(dataData))

		tmp := STAT_MCTC{
			VxMetadata:       metaData,
			STAT_MCTC_header: elem_header,
			Data:             make(map[string]STAT_MCTC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_MCTS":
		elem_header := STAT_MCTS_header{}
		appendErrorWithContext(&errs, "STAT_MCTS_header", elem_header.fill(headerData))
		elem_data := STAT_MCTS_data{}
		appendErrorWithContext(&errs, "STAT_MCTS_data", elem_data.fill(dataData))

		tmp := STAT_MCTS{
			VxMetadata:       metaData,
			STAT_MCTS_header: elem_header,
			Data:             make(map[string]STAT_MCTS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_MPR":
		elem_header := STAT_MPR_header{}
		appendErrorWithContext(&errs, "STAT_MPR_header", elem_header.fill(headerData))
		elem_data := STAT_MPR_data{}
		appendErrorWithContext(&errs, "STAT_MPR_data", elem_data.fill(dataData))

		tmp := STAT_MPR{
			VxMetadata:      metaData,
			STAT_MPR_header: elem_header,
			Data:            make(map[string]STAT_MPR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_NBRCNT":
		elem_header := STAT_NBRCNT_header{}
		appendErrorWithContext(&errs, "STAT_NBRCNT_header", elem_header.fill(headerData))
		elem_data := STAT_NBRCNT_data{}
		appendErrorWithContext(&errs, "STAT_NBRCNT_data", elem_data.fill(dataData))

		tmp := STAT_NBRCNT{
			VxMetadata:         metaData,
			STAT_NBRCNT_header: elem_header,
			Data:               make(map[string]STAT_NBRCNT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_NBRCTC":
		elem_header := STAT_NBRCTC_header{}
		appendErrorWithContext(&errs, "STAT_NBRCTC_header", elem_header.fill(headerData))
		elem_data := STAT_NBRCTC_data{}
		appendErrorWithContext(&errs, "STAT_NBRCTC_data", elem_data.fill(dataData))

		tmp := STAT_NBRCTC{
			VxMetadata:         metaData,
			STAT_NBRCTC_header: elem_header,
			Data:               make(map[string]STAT_NBRCTC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_NBRCTS":
		elem_header := STAT_NBRCTS_header{}
		appendErrorWithContext(&errs, "STAT_NBRCTS_header", elem_header.fill(headerData))
		elem_data := STAT_NBRCTS_data{}
		appendErrorWithContext(&errs, "STAT_NBRCTS_data", elem_data.fill(dataData))

		tmp := STAT_NBRCTS{
			VxMetadata:         metaData,
			STAT_NBRCTS_header: elem_header,
			Data:               make(map[string]STAT_NBRCTS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_ORANK":
		elem_header := STAT_ORANK_header{}
		appendErrorWithContext(&errs, "STAT_ORANK_header", elem_header.fill(headerData))
		elem_data := STAT_ORANK_data{}
		appendErrorWithContext(&errs, "STAT_ORANK_data", elem_data.fill(dataData))

		tmp := STAT_ORANK{
			VxMetadata:        metaData,
			STAT_ORANK_header: elem_header,
			Data:              make(map[string]STAT_ORANK_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_PCT":
		elem_header := STAT_PCT_header{}
		appendErrorWithContext(&errs, "STAT_PCT_header", elem_header.fill(headerData))
		elem_data := STAT_PCT_data{}
		appendErrorWithContext(&errs, "STAT_PCT_data", elem_data.fill(dataData))

		tmp := STAT_PCT{
			VxMetadata:      metaData,
			STAT_PCT_header: elem_header,
			Data:            make(map[string]STAT_PCT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_PHIST":
		elem_header := STAT_PHIST_header{}
		appendErrorWithContext(&errs, "STAT_PHIST_header", elem_header.fill(headerData))
		elem_data := STAT_PHIST_data{}
		appendErrorWithContext(&errs, "STAT_PHIST_data", elem_data.fill(dataData))

		tmp := STAT_PHIST{
			VxMetadata:        metaData,
			STAT_PHIST_header: elem_header,
			Data:              make(map[string]STAT_PHIST_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_PJC":
		elem_header := STAT_PJC_header{}
		appendErrorWithContext(&errs, "STAT_PJC_header", elem_header.fill(headerData))
		elem_data := STAT_PJC_data{}
		appendErrorWithContext(&errs, "STAT_PJC_data", elem_data.fill(dataData))

		tmp := STAT_PJC{
			VxMetadata:      metaData,
			STAT_PJC_header: elem_header,
			Data:            make(map[string]STAT_PJC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_PRC":
		elem_header := STAT_PRC_header{}
		appendErrorWithContext(&errs, "STAT_PRC_header", elem_header.fill(headerData))
		elem_data := STAT_PRC_data{}
		appendErrorWithContext(&errs, "STAT_PRC_data", elem_data.fill(dataData))

		tmp := STAT_PRC{
			VxMetadata:      metaData,
			STAT_PRC_header: elem_header,
			Data:            make(map[string]STAT_PRC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_PSTD":
		elem_header := STAT_PSTD_header{}
		appendErrorWithContext(&errs, "STAT_PSTD_header", elem_header.fill(headerData))
		elem_data := STAT_PSTD_data{}
		appendErrorWithContext(&errs, "STAT_PSTD_data", elem_data.fill(dataData))

		tmp := STAT_PSTD{
			VxMetadata:       metaData,
			STAT_PSTD_header: elem_header,
			Data:             make(map[string]STAT_PSTD_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_RELP":
		elem_header := STAT_RELP_header{}
		appendErrorWithContext(&errs, "STAT_RELP_header", elem_header.fill(headerData))
		elem_data := STAT_RELP_data{}
		appendErrorWithContext(&errs, "STAT_RELP_data", elem_data.fill(dataData))

		tmp := STAT_RELP{
			VxMetadata:       metaData,
			STAT_RELP_header: elem_header,
			Data:             make(map[string]STAT_RELP_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_RHIST":
		elem_header := STAT_RHIST_header{}
		appendErrorWithContext(&errs, "STAT_RHIST_header", elem_header.fill(headerData))
		elem_data := STAT_RHIST_data{}
		appendErrorWithContext(&errs, "STAT_RHIST_data", elem_data.fill(dataData))

		tmp := STAT_RHIST{
			VxMetadata:        metaData,
			STAT_RHIST_header: elem_header,
			Data:              make(map[string]STAT_RHIST_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_RPS":
		elem_header := STAT_RPS_header{}
		appendErrorWithContext(&errs, "STAT_RPS_header", elem_header.fill(headerData))
		elem_data := STAT_RPS_data{}
		appendErrorWithContext(&errs, "STAT_RPS_data", elem_data.fill(dataData))

		tmp := STAT_RPS{
			VxMetadata:      metaData,
			STAT_RPS_header: elem_header,
			Data:            make(map[string]STAT_RPS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_SAL1L2":
		elem_header := STAT_SAL1L2_header{}
		appendErrorWithContext(&errs, "STAT_SAL1L2_header", elem_header.fill(headerData))
		elem_data := STAT_SAL1L2_data{}
		appendErrorWithContext(&errs, "STAT_SAL1L2_data", elem_data.fill(dataData))

		tmp := STAT_SAL1L2{
			VxMetadata:         metaData,
			STAT_SAL1L2_header: elem_header,
			Data:               make(map[string]STAT_SAL1L2_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_SL1L2":
		elem_header := STAT_SL1L2_header{}
		appendErrorWithContext(&errs, "STAT_SL1L2_header", elem_header.fill(headerData))
		elem_data := STAT_SL1L2_data{}
		appendErrorWithContext(&errs, "STAT_SL1L2_data", elem_data.fill(dataData))

		tmp := STAT_SL1L2{
			VxMetadata:        metaData,
			STAT_SL1L2_header: elem_header,
			Data:              make(map[string]STAT_SL1L2_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_SSIDX":
		elem_header := STAT_SSIDX_header{}
		appendErrorWithContext(&errs, "STAT_SSIDX_header", elem_header.fill(headerData))
		elem_data := STAT_SSIDX_data{}
		appendErrorWithContext(&errs, "STAT_SSIDX_data", elem_data.fill(dataData))

		tmp := STAT_SSIDX{
			VxMetadata:        metaData,
			STAT_SSIDX_header: elem_header,
			Data:              make(map[string]STAT_SSIDX_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_SSVAR":
		elem_header := STAT_SSVAR_header{}
		appendErrorWithContext(&errs, "STAT_SSVAR_header", elem_header.fill(headerData))
		elem_data := STAT_SSVAR_data{}
		appendErrorWithContext(&errs, "STAT_SSVAR_data", elem_data.fill(dataData))

		tmp := STAT_SSVAR{
			VxMetadata:        metaData,
			STAT_SSVAR_header: elem_header,
			Data:              make(map[string]STAT_SSVAR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_VAL1L2":
		elem_header := STAT_VAL1L2_header{}
		appendErrorWithContext(&errs, "STAT_VAL1L2_header", elem_header.fill(headerData))
		elem_data := STAT_VAL1L2_data{}
		appendErrorWithContext(&errs, "STAT_VAL1L2_data", elem_data.fill(dataData))

		tmp := STAT_VAL1L2{
			VxMetadata:         metaData,
			STAT_VAL1L2_header: elem_header,
			Data:               make(map[string]STAT_VAL1L2_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_VCNT":
		elem_header := STAT_VCNT_header{}
		appendErrorWithContext(&errs, "STAT_VCNT_header", elem_header.fill(headerData))
		elem_data := STAT_VCNT_data{}
		appendErrorWithContext(&errs, "STAT_VCNT_data", elem_data.fill(dataData))

		tmp := STAT_VCNT{
			VxMetadata:       metaData,
			STAT_VCNT_header: elem_header,
			Data:             make(map[string]STAT_VCNT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "STAT_VL1L2":
		elem_header := STAT_VL1L2_header{}
		appendErrorWithContext(&errs, "STAT_VL1L2_header", elem_header.fill(headerData))
		elem_data := STAT_VL1L2_data{}
		appendErrorWithContext(&errs, "STAT_VL1L2_data", elem_data.fill(dataData))

		tmp := STAT_VL1L2{
			VxMetadata:        metaData,
			STAT_VL1L2_header: elem_header,
			Data:              make(map[string]STAT_VL1L2_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "TCST_PROBRIRW":
		elem_header := TCST_PROBRIRW_header{}
		appendErrorWithContext(&errs, "TCST_PROBRIRW_header", elem_header.fill(headerData))
		elem_data := TCST_PROBRIRW_data{}
		appendErrorWithContext(&errs, "TCST_PROBRIRW_data", elem_data.fill(dataData))

		tmp := TCST_PROBRIRW{
			VxMetadata:           metaData,
			TCST_PROBRIRW_header: elem_header,
			Data:                 make(map[string]TCST_PROBRIRW_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	case "TCST_TCMPR":
		elem_header := TCST_TCMPR_header{}
		appendErrorWithContext(&errs, "TCST_TCMPR_header", elem_header.fill(headerData))
		elem_data := TCST_TCMPR_data{}
		appendErrorWithContext(&errs, "TCST_TCMPR_data", elem_data.fill(dataData))

		tmp := TCST_TCMPR{
			VxMetadata:        metaData,
			TCST_TCMPR_header: elem_header,
			Data:              make(map[string]TCST_TCMPR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = &tmp
	default:
		return nil, errors.New("GetDocForId: Unknown file_line type:" + fileLineType)
	}
	return statDoc, errors.Join(errs...)
}
