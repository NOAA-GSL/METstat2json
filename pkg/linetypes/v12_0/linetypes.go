package v12_0

import (
	"encoding/json"
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

// Document struct definitions

// Represents a complete MODE_CTS document
type MODE_CTS struct {
	util.VxMetadata
	MODE_CTS_header
	Data map[string]MODE_CTS_data `json:"data"`
}

// Represents a complete MODE_OBJ document
type MODE_OBJ struct {
	util.VxMetadata
	MODE_OBJ_header
	Data map[string]MODE_OBJ_data `json:"data"`
}

// Represents a complete MTD_2DSINGLE document
type MTD_2DSINGLE struct {
	util.VxMetadata
	MTD_2DSINGLE_header
	Data map[string]MTD_2DSINGLE_data `json:"data"`
}

// Represents a complete MTD_3DPAIR document
type MTD_3DPAIR struct {
	util.VxMetadata
	MTD_3DPAIR_header
	Data map[string]MTD_3DPAIR_data `json:"data"`
}

// Represents a complete MTD_3DSINGLE document
type MTD_3DSINGLE struct {
	util.VxMetadata
	MTD_3DSINGLE_header
	Data map[string]MTD_3DSINGLE_data `json:"data"`
}

// Represents a complete STAT_CNT document
type STAT_CNT struct {
	util.VxMetadata
	STAT_CNT_header
	Data map[string]STAT_CNT_data `json:"data"`
}

// Represents a complete STAT_CTC document
type STAT_CTC struct {
	util.VxMetadata
	STAT_CTC_header
	Data map[string]STAT_CTC_data `json:"data"`
}

// Represents a complete STAT_CTS document
type STAT_CTS struct {
	util.VxMetadata
	STAT_CTS_header
	Data map[string]STAT_CTS_data `json:"data"`
}

// Represents a complete STAT_DMAP document
type STAT_DMAP struct {
	util.VxMetadata
	STAT_DMAP_header
	Data map[string]STAT_DMAP_data `json:"data"`
}

// Represents a complete STAT_ECLV document
type STAT_ECLV struct {
	util.VxMetadata
	STAT_ECLV_header
	Data map[string]STAT_ECLV_data `json:"data"`
}

// Represents a complete STAT_ECNT document
type STAT_ECNT struct {
	util.VxMetadata
	STAT_ECNT_header
	Data map[string]STAT_ECNT_data `json:"data"`
}

// Represents a complete STAT_FHO document
type STAT_FHO struct {
	util.VxMetadata
	STAT_FHO_header
	Data map[string]STAT_FHO_data `json:"data"`
}

// Represents a complete STAT_GENMPR document
type STAT_GENMPR struct {
	util.VxMetadata
	STAT_GENMPR_header
	Data map[string]STAT_GENMPR_data `json:"data"`
}

// Represents a complete STAT_GRAD document
type STAT_GRAD struct {
	util.VxMetadata
	STAT_GRAD_header
	Data map[string]STAT_GRAD_data `json:"data"`
}

// Represents a complete STAT_ISC document
type STAT_ISC struct {
	util.VxMetadata
	STAT_ISC_header
	Data map[string]STAT_ISC_data `json:"data"`
}

// Represents a complete STAT_MCTC document
type STAT_MCTC struct {
	util.VxMetadata
	STAT_MCTC_header
	Data map[string]STAT_MCTC_data `json:"data"`
}

// Represents a complete STAT_MCTS document
type STAT_MCTS struct {
	util.VxMetadata
	STAT_MCTS_header
	Data map[string]STAT_MCTS_data `json:"data"`
}

// Represents a complete STAT_MPR document
type STAT_MPR struct {
	util.VxMetadata
	STAT_MPR_header
	Data map[string]STAT_MPR_data `json:"data"`
}

// Represents a complete STAT_NBRCNT document
type STAT_NBRCNT struct {
	util.VxMetadata
	STAT_NBRCNT_header
	Data map[string]STAT_NBRCNT_data `json:"data"`
}

// Represents a complete STAT_NBRCTC document
type STAT_NBRCTC struct {
	util.VxMetadata
	STAT_NBRCTC_header
	Data map[string]STAT_NBRCTC_data `json:"data"`
}

// Represents a complete STAT_NBRCTS document
type STAT_NBRCTS struct {
	util.VxMetadata
	STAT_NBRCTS_header
	Data map[string]STAT_NBRCTS_data `json:"data"`
}

// Represents a complete STAT_ORANK document
type STAT_ORANK struct {
	util.VxMetadata
	STAT_ORANK_header
	Data map[string]STAT_ORANK_data `json:"data"`
}

// Represents a complete STAT_PCT document
type STAT_PCT struct {
	util.VxMetadata
	STAT_PCT_header
	Data map[string]STAT_PCT_data `json:"data"`
}

// Represents a complete STAT_PHIST document
type STAT_PHIST struct {
	util.VxMetadata
	STAT_PHIST_header
	Data map[string]STAT_PHIST_data `json:"data"`
}

// Represents a complete STAT_PJC document
type STAT_PJC struct {
	util.VxMetadata
	STAT_PJC_header
	Data map[string]STAT_PJC_data `json:"data"`
}

// Represents a complete STAT_PRC document
type STAT_PRC struct {
	util.VxMetadata
	STAT_PRC_header
	Data map[string]STAT_PRC_data `json:"data"`
}

// Represents a complete STAT_PSTD document
type STAT_PSTD struct {
	util.VxMetadata
	STAT_PSTD_header
	Data map[string]STAT_PSTD_data `json:"data"`
}

// Represents a complete STAT_RELP document
type STAT_RELP struct {
	util.VxMetadata
	STAT_RELP_header
	Data map[string]STAT_RELP_data `json:"data"`
}

// Represents a complete STAT_RHIST document
type STAT_RHIST struct {
	util.VxMetadata
	STAT_RHIST_header
	Data map[string]STAT_RHIST_data `json:"data"`
}

// Represents a complete STAT_RPS document
type STAT_RPS struct {
	util.VxMetadata
	STAT_RPS_header
	Data map[string]STAT_RPS_data `json:"data"`
}

// Represents a complete STAT_SAL1L2 document
type STAT_SAL1L2 struct {
	util.VxMetadata
	STAT_SAL1L2_header
	Data map[string]STAT_SAL1L2_data `json:"data"`
}

// Represents a complete STAT_SEEPS document
type STAT_SEEPS struct {
	util.VxMetadata
	STAT_SEEPS_header
	Data map[string]STAT_SEEPS_data `json:"data"`
}

// Represents a complete STAT_SEEPS_MPR document
type STAT_SEEPS_MPR struct {
	util.VxMetadata
	STAT_SEEPS_MPR_header
	Data map[string]STAT_SEEPS_MPR_data `json:"data"`
}

// Represents a complete STAT_SL1L2 document
type STAT_SL1L2 struct {
	util.VxMetadata
	STAT_SL1L2_header
	Data map[string]STAT_SL1L2_data `json:"data"`
}

// Represents a complete STAT_SSIDX document
type STAT_SSIDX struct {
	util.VxMetadata
	STAT_SSIDX_header
	Data map[string]STAT_SSIDX_data `json:"data"`
}

// Represents a complete STAT_SSVAR document
type STAT_SSVAR struct {
	util.VxMetadata
	STAT_SSVAR_header
	Data map[string]STAT_SSVAR_data `json:"data"`
}

// Represents a complete STAT_VAL1L2 document
type STAT_VAL1L2 struct {
	util.VxMetadata
	STAT_VAL1L2_header
	Data map[string]STAT_VAL1L2_data `json:"data"`
}

// Represents a complete STAT_VCNT document
type STAT_VCNT struct {
	util.VxMetadata
	STAT_VCNT_header
	Data map[string]STAT_VCNT_data `json:"data"`
}

// Represents a complete STAT_VL1L2 document
type STAT_VL1L2 struct {
	util.VxMetadata
	STAT_VL1L2_header
	Data map[string]STAT_VL1L2_data `json:"data"`
}

// Represents a complete TCST_PROBRIRW document
type TCST_PROBRIRW struct {
	util.VxMetadata
	TCST_PROBRIRW_header
	Data map[string]TCST_PROBRIRW_data `json:"data"`
}

// Represents a complete TCST_TCDIAG document
type TCST_TCDIAG struct {
	util.VxMetadata
	TCST_TCDIAG_header
	Data map[string]TCST_TCDIAG_data `json:"data"`
}

// Represents a complete TCST_TCMPR document
type TCST_TCMPR struct {
	util.VxMetadata
	TCST_TCMPR_header
	Data map[string]TCST_TCMPR_data `json:"data"`
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

// Represents the header field of a MTD_2DSINGLE document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type MTD_2DSINGLE_header struct {
	VERSION    validtypes.ValidString `json:"VERSION"`
	MODEL      validtypes.ValidString `json:"MODEL"`
	DESC       validtypes.ValidString `json:"DESC"`
	FCST_LEAD  validtypes.ValidInt    `json:"FCST_LEAD"`
	FCST_VALID validtypes.ValidString `json:"FCST_VALID"`
	OBS_LEAD   validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID  validtypes.ValidString `json:"OBS_VALID"`
	T_DELTA    validtypes.ValidString `json:"T_DELTA"`
	FCST_T_BEG validtypes.ValidInt    `json:"FCST_T_BEG"`
	FCST_T_END validtypes.ValidInt    `json:"FCST_T_END"`
	FCST_RAD   validtypes.ValidInt    `json:"FCST_RAD"`
	FCST_THR   validtypes.ValidString `json:"FCST_THR"`
	OBS_T_BEG  validtypes.ValidInt    `json:"OBS_T_BEG"`
	OBS_T_END  validtypes.ValidInt    `json:"OBS_T_END"`
	OBS_RAD    validtypes.ValidInt    `json:"OBS_RAD"`
	OBS_THR    validtypes.ValidString `json:"OBS_THR"`
	FCST_VAR   validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV   validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR    validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS  validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV    validtypes.ValidString `json:"OBS_LEV"`
	LINE_TYPE  validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a MTD_3DPAIR document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type MTD_3DPAIR_header struct {
	VERSION    validtypes.ValidString `json:"VERSION"`
	MODEL      validtypes.ValidString `json:"MODEL"`
	DESC       validtypes.ValidString `json:"DESC"`
	FCST_LEAD  validtypes.ValidInt    `json:"FCST_LEAD"`
	FCST_VALID validtypes.ValidString `json:"FCST_VALID"`
	OBS_LEAD   validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID  validtypes.ValidString `json:"OBS_VALID"`
	T_DELTA    validtypes.ValidString `json:"T_DELTA"`
	FCST_T_BEG validtypes.ValidInt    `json:"FCST_T_BEG"`
	FCST_T_END validtypes.ValidInt    `json:"FCST_T_END"`
	FCST_RAD   validtypes.ValidInt    `json:"FCST_RAD"`
	FCST_THR   validtypes.ValidString `json:"FCST_THR"`
	OBS_T_BEG  validtypes.ValidInt    `json:"OBS_T_BEG"`
	OBS_T_END  validtypes.ValidInt    `json:"OBS_T_END"`
	OBS_RAD    validtypes.ValidInt    `json:"OBS_RAD"`
	OBS_THR    validtypes.ValidString `json:"OBS_THR"`
	FCST_VAR   validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV   validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR    validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS  validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV    validtypes.ValidString `json:"OBS_LEV"`
	LINE_TYPE  validtypes.ValidString `json:"LINE_TYPE"`
}

// Represents the header field of a MTD_3DSINGLE document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type MTD_3DSINGLE_header struct {
	VERSION    validtypes.ValidString `json:"VERSION"`
	MODEL      validtypes.ValidString `json:"MODEL"`
	DESC       validtypes.ValidString `json:"DESC"`
	FCST_LEAD  validtypes.ValidInt    `json:"FCST_LEAD"`
	FCST_VALID validtypes.ValidString `json:"FCST_VALID"`
	OBS_LEAD   validtypes.ValidInt    `json:"OBS_LEAD"`
	OBS_VALID  validtypes.ValidString `json:"OBS_VALID"`
	T_DELTA    validtypes.ValidString `json:"T_DELTA"`
	FCST_T_BEG validtypes.ValidInt    `json:"FCST_T_BEG"`
	FCST_T_END validtypes.ValidInt    `json:"FCST_T_END"`
	FCST_RAD   validtypes.ValidInt    `json:"FCST_RAD"`
	FCST_THR   validtypes.ValidString `json:"FCST_THR"`
	OBS_T_BEG  validtypes.ValidInt    `json:"OBS_T_BEG"`
	OBS_T_END  validtypes.ValidInt    `json:"OBS_T_END"`
	OBS_RAD    validtypes.ValidInt    `json:"OBS_RAD"`
	OBS_THR    validtypes.ValidString `json:"OBS_THR"`
	FCST_VAR   validtypes.ValidString `json:"FCST_VAR"`
	FCST_UNITS validtypes.ValidString `json:"FCST_UNITS"`
	FCST_LEV   validtypes.ValidString `json:"FCST_LEV"`
	OBS_VAR    validtypes.ValidString `json:"OBS_VAR"`
	OBS_UNITS  validtypes.ValidString `json:"OBS_UNITS"`
	OBS_LEV    validtypes.ValidString `json:"OBS_LEV"`
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

// Represents the header field of a STAT_SEEPS_MPR document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_SEEPS_MPR_header struct {
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

// Represents the header field of a STAT_SEEPS document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type STAT_SEEPS_header struct {
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

// Represents the header field of a TCST_TCDIAG document
// TODO - there are only 4 of these headers - MODE, MTD, STAT, and TCST. This can be represented more efficiently.
type TCST_TCDIAG_header struct {
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
func (s *MODE_CTS_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.N_VALID.UnmarshalText([]byte(fields[2]))
	s.GRID_RES.UnmarshalText([]byte(fields[3]))
	s.DESC.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID.UnmarshalText([]byte(fields[6]))
	s.FCST_ACCUM.UnmarshalText([]byte(fields[7]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[8]))
	s.OBS_VALID.UnmarshalText([]byte(fields[9]))
	s.OBS_ACCUM.UnmarshalText([]byte(fields[10]))
	s.FCST_RAD.UnmarshalText([]byte(fields[11]))
	s.FCST_THR.UnmarshalText([]byte(fields[12]))
	s.OBS_RAD.UnmarshalText([]byte(fields[13]))
	s.OBS_THR.UnmarshalText([]byte(fields[14]))
	s.FCST_VAR.UnmarshalText([]byte(fields[15]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[16]))
	s.FCST_LEV.UnmarshalText([]byte(fields[17]))
	s.OBS_VAR.UnmarshalText([]byte(fields[18]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[19]))
	s.OBS_LEV.UnmarshalText([]byte(fields[20]))
	s.OBTYPE.UnmarshalText([]byte(fields[21]))
	s.LINE_TYPE.UnmarshalText([]byte("MODE_CTS")) // hardcode the LINE_TYPE
}

func (s *MODE_OBJ_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.N_VALID.UnmarshalText([]byte(fields[2]))
	s.GRID_RES.UnmarshalText([]byte(fields[3]))
	s.DESC.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID.UnmarshalText([]byte(fields[6]))
	s.FCST_ACCUM.UnmarshalText([]byte(fields[7]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[8]))
	s.OBS_VALID.UnmarshalText([]byte(fields[9]))
	s.OBS_ACCUM.UnmarshalText([]byte(fields[10]))
	s.FCST_RAD.UnmarshalText([]byte(fields[11]))
	s.FCST_THR.UnmarshalText([]byte(fields[12]))
	s.OBS_RAD.UnmarshalText([]byte(fields[13]))
	s.OBS_THR.UnmarshalText([]byte(fields[14]))
	s.FCST_VAR.UnmarshalText([]byte(fields[15]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[16]))
	s.FCST_LEV.UnmarshalText([]byte(fields[17]))
	s.OBS_VAR.UnmarshalText([]byte(fields[18]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[19]))
	s.OBS_LEV.UnmarshalText([]byte(fields[20]))
	s.OBTYPE.UnmarshalText([]byte(fields[21]))
	s.LINE_TYPE.UnmarshalText([]byte("MODE_OBJ")) // hardcode the LINE_TYPE
}

func (s *MTD_2DSINGLE_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_LEAD.UnmarshalText([]byte(fields[3]))
	s.FCST_VALID.UnmarshalText([]byte(fields[4]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[5]))
	s.OBS_VALID.UnmarshalText([]byte(fields[6]))
	s.T_DELTA.UnmarshalText([]byte(fields[7]))
	s.FCST_T_BEG.UnmarshalText([]byte(fields[8]))
	s.FCST_T_END.UnmarshalText([]byte(fields[9]))
	s.FCST_RAD.UnmarshalText([]byte(fields[10]))
	s.FCST_THR.UnmarshalText([]byte(fields[11]))
	s.OBS_T_BEG.UnmarshalText([]byte(fields[12]))
	s.OBS_T_END.UnmarshalText([]byte(fields[13]))
	s.OBS_RAD.UnmarshalText([]byte(fields[14]))
	s.OBS_THR.UnmarshalText([]byte(fields[15]))
	s.FCST_VAR.UnmarshalText([]byte(fields[16]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[17]))
	s.FCST_LEV.UnmarshalText([]byte(fields[18]))
	s.OBS_VAR.UnmarshalText([]byte(fields[19]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[20]))
	s.OBS_LEV.UnmarshalText([]byte(fields[21]))
	s.LINE_TYPE.UnmarshalText([]byte("MTD_2DSINGLE")) // hardcode the LINE_TYPE
}

func (s *MTD_3DPAIR_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_LEAD.UnmarshalText([]byte(fields[3]))
	s.FCST_VALID.UnmarshalText([]byte(fields[4]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[5]))
	s.OBS_VALID.UnmarshalText([]byte(fields[6]))
	s.T_DELTA.UnmarshalText([]byte(fields[7]))
	s.FCST_T_BEG.UnmarshalText([]byte(fields[8]))
	s.FCST_T_END.UnmarshalText([]byte(fields[9]))
	s.FCST_RAD.UnmarshalText([]byte(fields[10]))
	s.FCST_THR.UnmarshalText([]byte(fields[11]))
	s.OBS_T_BEG.UnmarshalText([]byte(fields[12]))
	s.OBS_T_END.UnmarshalText([]byte(fields[13]))
	s.OBS_RAD.UnmarshalText([]byte(fields[14]))
	s.OBS_THR.UnmarshalText([]byte(fields[15]))
	s.FCST_VAR.UnmarshalText([]byte(fields[16]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[17]))
	s.FCST_LEV.UnmarshalText([]byte(fields[18]))
	s.OBS_VAR.UnmarshalText([]byte(fields[19]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[20]))
	s.OBS_LEV.UnmarshalText([]byte(fields[21]))
	s.LINE_TYPE.UnmarshalText([]byte("MTD_3DPAIR")) // hardcode the LINE_TYPE
}

func (s *MTD_3DSINGLE_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_LEAD.UnmarshalText([]byte(fields[3]))
	s.FCST_VALID.UnmarshalText([]byte(fields[4]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[5]))
	s.OBS_VALID.UnmarshalText([]byte(fields[6]))
	s.T_DELTA.UnmarshalText([]byte(fields[7]))
	s.FCST_T_BEG.UnmarshalText([]byte(fields[8]))
	s.FCST_T_END.UnmarshalText([]byte(fields[9]))
	s.FCST_RAD.UnmarshalText([]byte(fields[10]))
	s.FCST_THR.UnmarshalText([]byte(fields[11]))
	s.OBS_T_BEG.UnmarshalText([]byte(fields[12]))
	s.OBS_T_END.UnmarshalText([]byte(fields[13]))
	s.OBS_RAD.UnmarshalText([]byte(fields[14]))
	s.OBS_THR.UnmarshalText([]byte(fields[15]))
	s.FCST_VAR.UnmarshalText([]byte(fields[16]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[17]))
	s.FCST_LEV.UnmarshalText([]byte(fields[18]))
	s.OBS_VAR.UnmarshalText([]byte(fields[19]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[20]))
	s.OBS_LEV.UnmarshalText([]byte(fields[21]))
	s.LINE_TYPE.UnmarshalText([]byte("MTD_3DSINGLE")) // hardcode the LINE_TYPE
}

func (s *STAT_CNT_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_CTC_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_CTS_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_DMAP_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_ECLV_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_ECNT_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_FHO_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_GENMPR_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_GRAD_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_ISC_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_MCTC_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_MCTS_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_MPR_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_NBRCNT_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_NBRCTC_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_NBRCTS_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_ORANK_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_PCT_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_PHIST_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_PJC_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_PRC_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_PSTD_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_RELP_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_RHIST_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_RPS_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_SAL1L2_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_SEEPS_MPR_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_SEEPS_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_SL1L2_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_SSIDX_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_SSVAR_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_VAL1L2_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_VCNT_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_VL1L2_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.MODEL.UnmarshalText([]byte(fields[1]))
	s.DESC.UnmarshalText([]byte(fields[2]))
	s.FCST_VALID_BEG.UnmarshalText([]byte(fields[4]))
	s.FCST_VALID_END.UnmarshalText([]byte(fields[5]))
	s.OBS_LEAD.UnmarshalText([]byte(fields[6]))
	s.OBS_VALID_BEG.UnmarshalText([]byte(fields[7]))
	s.OBS_VALID_END.UnmarshalText([]byte(fields[8]))
	s.FCST_VAR.UnmarshalText([]byte(fields[9]))
	s.FCST_UNITS.UnmarshalText([]byte(fields[10]))
	s.FCST_LEV.UnmarshalText([]byte(fields[11]))
	s.OBS_VAR.UnmarshalText([]byte(fields[12]))
	s.OBS_UNITS.UnmarshalText([]byte(fields[13]))
	s.OBS_LEV.UnmarshalText([]byte(fields[14]))
	s.OBTYPE.UnmarshalText([]byte(fields[15]))
	s.VX_MASK.UnmarshalText([]byte(fields[16]))
	s.INTERP_MTHD.UnmarshalText([]byte(fields[17]))
	s.INTERP_PNTS.UnmarshalText([]byte(fields[18]))
	s.FCST_THRESH.UnmarshalText([]byte(fields[19]))
	s.OBS_THRESH.UnmarshalText([]byte(fields[20]))
	s.COV_THRESH.UnmarshalText([]byte(fields[21]))
	s.ALPHA.UnmarshalText([]byte(fields[22]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[23]))
}

func (s *TCST_PROBRIRW_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.AMODEL.UnmarshalText([]byte(fields[1]))
	s.BMODEL.UnmarshalText([]byte(fields[2]))
	s.DESC.UnmarshalText([]byte(fields[3]))
	s.STORM_ID.UnmarshalText([]byte(fields[4]))
	s.BASIN.UnmarshalText([]byte(fields[5]))
	s.CYCLONE.UnmarshalText([]byte(fields[6]))
	s.STORM_NAME.UnmarshalText([]byte(fields[7]))
	s.VALID.UnmarshalText([]byte(fields[10]))
	s.INIT_MASK.UnmarshalText([]byte(fields[11]))
	s.VALID_MASK.UnmarshalText([]byte(fields[12]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[13]))
}

func (s *TCST_TCDIAG_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.AMODEL.UnmarshalText([]byte(fields[1]))
	s.BMODEL.UnmarshalText([]byte(fields[2]))
	s.DESC.UnmarshalText([]byte(fields[3]))
	s.STORM_ID.UnmarshalText([]byte(fields[4]))
	s.BASIN.UnmarshalText([]byte(fields[5]))
	s.CYCLONE.UnmarshalText([]byte(fields[6]))
	s.STORM_NAME.UnmarshalText([]byte(fields[7]))
	s.VALID.UnmarshalText([]byte(fields[10]))
	s.INIT_MASK.UnmarshalText([]byte(fields[11]))
	s.VALID_MASK.UnmarshalText([]byte(fields[12]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[13]))
}

func (s *TCST_TCMPR_header) fill(fields []string) {
	s.VERSION.UnmarshalText([]byte(fields[0]))
	s.AMODEL.UnmarshalText([]byte(fields[1]))
	s.BMODEL.UnmarshalText([]byte(fields[2]))
	s.DESC.UnmarshalText([]byte(fields[3]))
	s.STORM_ID.UnmarshalText([]byte(fields[4]))
	s.BASIN.UnmarshalText([]byte(fields[5]))
	s.CYCLONE.UnmarshalText([]byte(fields[6]))
	s.STORM_NAME.UnmarshalText([]byte(fields[7]))
	s.VALID.UnmarshalText([]byte(fields[10]))
	s.INIT_MASK.UnmarshalText([]byte(fields[11]))
	s.VALID_MASK.UnmarshalText([]byte(fields[12]))
	s.LINE_TYPE.UnmarshalText([]byte(fields[13]))
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
	LODDS validtypes.ValidFloat  `json:"LODDS,omitzero"`
	ORSS  validtypes.ValidFloat  `json:"ORSS,omitzero"`
	EDS   validtypes.ValidFloat  `json:"EDS,omitzero"`
	SEDS  validtypes.ValidFloat  `json:"SEDS,omitzero"`
	EDI   validtypes.ValidFloat  `json:"EDI,omitzero"`
	SEDI  validtypes.ValidFloat  `json:"SEDI,omitzero"`
	BAGSS validtypes.ValidFloat  `json:"BAGSS,omitzero"`
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

type MTD_2DSINGLE_data struct {
	OBJECT_ID      validtypes.ValidString `json:"OBJECT_ID,omitzero"`
	OBJECT_CAT     validtypes.ValidString `json:"OBJECT_CAT,omitzero"`
	TIME_INDEX     validtypes.ValidInt    `json:"TIME_INDEX,omitzero"`
	AREA           validtypes.ValidInt    `json:"AREA,omitzero"`
	CENTROID_X     validtypes.ValidFloat  `json:"CENTROID_X,omitzero"`
	CENTROID_Y     validtypes.ValidFloat  `json:"CENTROID_Y,omitzero"`
	CENTROID_LAT   validtypes.ValidFloat  `json:"CENTROID_LAT,omitzero"`
	CENTROID_LON   validtypes.ValidFloat  `json:"CENTROID_LON,omitzero"`
	AXIS_ANG       validtypes.ValidFloat  `json:"AXIS_ANG,omitzero"`
	INTENSITY_10   validtypes.ValidFloat  `json:"INTENSITY_10,omitzero"`
	INTENSITY_25   validtypes.ValidFloat  `json:"INTENSITY_25,omitzero"`
	INTENSITY_50   validtypes.ValidFloat  `json:"INTENSITY_50,omitzero"`
	INTENSITY_75   validtypes.ValidFloat  `json:"INTENSITY_75,omitzero"`
	INTENSITY_90   validtypes.ValidFloat  `json:"INTENSITY_90,omitzero"`
	INTENSITY_USER validtypes.ValidFloat  `json:"INTENSITY_USER,omitzero"`
}

type MTD_3DPAIR_data struct {
	OBJECT_ID           validtypes.ValidString `json:"OBJECT_ID,omitzero"`
	OBJECT_CAT          validtypes.ValidString `json:"OBJECT_CAT,omitzero"`
	SPACE_CENTROID_DIST validtypes.ValidFloat  `json:"SPACE_CENTROID_DIST,omitzero"`
	TIME_CENTROID_DELTA validtypes.ValidFloat  `json:"TIME_CENTROID_DELTA,omitzero"`
	AXIS_DIFF           validtypes.ValidFloat  `json:"AXIS_DIFF,omitzero"`
	SPEED_DELTA         validtypes.ValidFloat  `json:"SPEED_DELTA,omitzero"`
	DIRECTION_DIFF      validtypes.ValidFloat  `json:"DIRECTION_DIFF,omitzero"`
	VOLUME_RATIO        validtypes.ValidFloat  `json:"VOLUME_RATIO,omitzero"`
	START_TIME_DELTA    validtypes.ValidInt    `json:"START_TIME_DELTA,omitzero"`
	END_TIME_DELTA      validtypes.ValidInt    `json:"END_TIME_DELTA,omitzero"`
	INTERSECTION_VOLUME validtypes.ValidFloat  `json:"INTERSECTION_VOLUME,omitzero"`
	DURATION_DIFF       validtypes.ValidFloat  `json:"DURATION_DIFF,omitzero"`
	INTEREST            validtypes.ValidFloat  `json:"INTEREST,omitzero"`
}

type MTD_3DSINGLE_data struct {
	OBJECT_ID       validtypes.ValidString `json:"OBJECT_ID,omitzero"`
	OBJECT_CAT      validtypes.ValidString `json:"OBJECT_CAT,omitzero"`
	CENTROID_X      validtypes.ValidFloat  `json:"CENTROID_X,omitzero"`
	CENTROID_Y      validtypes.ValidFloat  `json:"CENTROID_Y,omitzero"`
	CENTROID_T      validtypes.ValidFloat  `json:"CENTROID_T,omitzero"`
	CENTROID_LAT    validtypes.ValidFloat  `json:"CENTROID_LAT,omitzero"`
	CENTROID_LON    validtypes.ValidFloat  `json:"CENTROID_LON,omitzero"`
	X_DOT           validtypes.ValidFloat  `json:"X_DOT,omitzero"`
	Y_DOT           validtypes.ValidFloat  `json:"Y_DOT,omitzero"`
	AXIS_ANG        validtypes.ValidFloat  `json:"AXIS_ANG,omitzero"`
	VOLUME          validtypes.ValidInt    `json:"VOLUME,omitzero"`
	START_TIME      validtypes.ValidInt    `json:"START_TIME,omitzero"`
	END_TIME        validtypes.ValidInt    `json:"END_TIME,omitzero"`
	CDIST_TRAVELLED validtypes.ValidFloat  `json:"CDIST_TRAVELLED,omitzero"`
	INTENSITY_10    validtypes.ValidFloat  `json:"INTENSITY_10,omitzero"`
	INTENSITY_25    validtypes.ValidFloat  `json:"INTENSITY_25,omitzero"`
	INTENSITY_50    validtypes.ValidFloat  `json:"INTENSITY_50,omitzero"`
	INTENSITY_75    validtypes.ValidFloat  `json:"INTENSITY_75,omitzero"`
	INTENSITY_90    validtypes.ValidFloat  `json:"INTENSITY_90,omitzero"`
	INTENSITY_USER  validtypes.ValidFloat  `json:"INTENSITY_USER,omitzero"`
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
	TOTAL    validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FY_OY    validtypes.ValidFloat `json:"FY_OY,omitzero"`
	FY_ON    validtypes.ValidFloat `json:"FY_ON,omitzero"`
	FN_OY    validtypes.ValidFloat `json:"FN_OY,omitzero"`
	FN_ON    validtypes.ValidFloat `json:"FN_ON,omitzero"`
	EC_VALUE validtypes.ValidFloat `json:"EC_VALUE,omitzero"`
}

type STAT_CTS_data struct {
	TOTAL      validtypes.ValidInt   `json:"TOTAL,omitzero"`
	BASER      validtypes.ValidFloat `json:"BASER,omitzero"`
	BASER_NCL  validtypes.ValidFloat `json:"BASER_NCL,omitzero"`
	BASER_NCU  validtypes.ValidFloat `json:"BASER_NCU,omitzero"`
	BASER_BCL  validtypes.ValidFloat `json:"BASER_BCL,omitzero"`
	BASER_BCU  validtypes.ValidFloat `json:"BASER_BCU,omitzero"`
	FMEAN      validtypes.ValidFloat `json:"FMEAN,omitzero"`
	FMEAN_NCL  validtypes.ValidFloat `json:"FMEAN_NCL,omitzero"`
	FMEAN_NCU  validtypes.ValidFloat `json:"FMEAN_NCU,omitzero"`
	FMEAN_BCL  validtypes.ValidFloat `json:"FMEAN_BCL,omitzero"`
	FMEAN_BCU  validtypes.ValidFloat `json:"FMEAN_BCU,omitzero"`
	ACC        validtypes.ValidFloat `json:"ACC,omitzero"`
	ACC_NCL    validtypes.ValidFloat `json:"ACC_NCL,omitzero"`
	ACC_NCU    validtypes.ValidFloat `json:"ACC_NCU,omitzero"`
	ACC_BCL    validtypes.ValidFloat `json:"ACC_BCL,omitzero"`
	ACC_BCU    validtypes.ValidFloat `json:"ACC_BCU,omitzero"`
	FBIAS      validtypes.ValidFloat `json:"FBIAS,omitzero"`
	FBIAS_BCL  validtypes.ValidFloat `json:"FBIAS_BCL,omitzero"`
	FBIAS_BCU  validtypes.ValidFloat `json:"FBIAS_BCU,omitzero"`
	PODY       validtypes.ValidFloat `json:"PODY,omitzero"`
	PODY_NCL   validtypes.ValidFloat `json:"PODY_NCL,omitzero"`
	PODY_NCU   validtypes.ValidFloat `json:"PODY_NCU,omitzero"`
	PODY_BCL   validtypes.ValidFloat `json:"PODY_BCL,omitzero"`
	PODY_BCU   validtypes.ValidFloat `json:"PODY_BCU,omitzero"`
	PODN       validtypes.ValidFloat `json:"PODN,omitzero"`
	PODN_NCL   validtypes.ValidFloat `json:"PODN_NCL,omitzero"`
	PODN_NCU   validtypes.ValidFloat `json:"PODN_NCU,omitzero"`
	PODN_BCL   validtypes.ValidFloat `json:"PODN_BCL,omitzero"`
	PODN_BCU   validtypes.ValidFloat `json:"PODN_BCU,omitzero"`
	POFD       validtypes.ValidFloat `json:"POFD,omitzero"`
	POFD_NCL   validtypes.ValidFloat `json:"POFD_NCL,omitzero"`
	POFD_NCU   validtypes.ValidFloat `json:"POFD_NCU,omitzero"`
	POFD_BCL   validtypes.ValidFloat `json:"POFD_BCL,omitzero"`
	POFD_BCU   validtypes.ValidFloat `json:"POFD_BCU,omitzero"`
	FAR        validtypes.ValidFloat `json:"FAR,omitzero"`
	FAR_NCL    validtypes.ValidFloat `json:"FAR_NCL,omitzero"`
	FAR_NCU    validtypes.ValidFloat `json:"FAR_NCU,omitzero"`
	FAR_BCL    validtypes.ValidFloat `json:"FAR_BCL,omitzero"`
	FAR_BCU    validtypes.ValidFloat `json:"FAR_BCU,omitzero"`
	CSI        validtypes.ValidFloat `json:"CSI,omitzero"`
	CSI_NCL    validtypes.ValidFloat `json:"CSI_NCL,omitzero"`
	CSI_NCU    validtypes.ValidFloat `json:"CSI_NCU,omitzero"`
	CSI_BCL    validtypes.ValidFloat `json:"CSI_BCL,omitzero"`
	CSI_BCU    validtypes.ValidFloat `json:"CSI_BCU,omitzero"`
	GSS        validtypes.ValidFloat `json:"GSS,omitzero"`
	GSS_BCL    validtypes.ValidFloat `json:"GSS_BCL,omitzero"`
	GSS_BCU    validtypes.ValidFloat `json:"GSS_BCU,omitzero"`
	HK         validtypes.ValidFloat `json:"HK,omitzero"`
	HK_NCL     validtypes.ValidFloat `json:"HK_NCL,omitzero"`
	HK_NCU     validtypes.ValidFloat `json:"HK_NCU,omitzero"`
	HK_BCL     validtypes.ValidFloat `json:"HK_BCL,omitzero"`
	HK_BCU     validtypes.ValidFloat `json:"HK_BCU,omitzero"`
	HSS        validtypes.ValidFloat `json:"HSS,omitzero"`
	HSS_BCL    validtypes.ValidFloat `json:"HSS_BCL,omitzero"`
	HSS_BCU    validtypes.ValidFloat `json:"HSS_BCU,omitzero"`
	ODDS       validtypes.ValidFloat `json:"ODDS,omitzero"`
	ODDS_NCL   validtypes.ValidFloat `json:"ODDS_NCL,omitzero"`
	ODDS_NCU   validtypes.ValidFloat `json:"ODDS_NCU,omitzero"`
	ODDS_BCL   validtypes.ValidFloat `json:"ODDS_BCL,omitzero"`
	ODDS_BCU   validtypes.ValidFloat `json:"ODDS_BCU,omitzero"`
	LODDS      validtypes.ValidFloat `json:"LODDS,omitzero"`
	LODDS_NCL  validtypes.ValidFloat `json:"LODDS_NCL,omitzero"`
	LODDS_NCU  validtypes.ValidFloat `json:"LODDS_NCU,omitzero"`
	LODDS_BCL  validtypes.ValidFloat `json:"LODDS_BCL,omitzero"`
	LODDS_BCU  validtypes.ValidFloat `json:"LODDS_BCU,omitzero"`
	ORSS       validtypes.ValidFloat `json:"ORSS,omitzero"`
	ORSS_NCL   validtypes.ValidFloat `json:"ORSS_NCL,omitzero"`
	ORSS_NCU   validtypes.ValidFloat `json:"ORSS_NCU,omitzero"`
	ORSS_BCL   validtypes.ValidFloat `json:"ORSS_BCL,omitzero"`
	ORSS_BCU   validtypes.ValidFloat `json:"ORSS_BCU,omitzero"`
	EDS        validtypes.ValidFloat `json:"EDS,omitzero"`
	EDS_NCL    validtypes.ValidFloat `json:"EDS_NCL,omitzero"`
	EDS_NCU    validtypes.ValidFloat `json:"EDS_NCU,omitzero"`
	EDS_BCL    validtypes.ValidFloat `json:"EDS_BCL,omitzero"`
	EDS_BCU    validtypes.ValidFloat `json:"EDS_BCU,omitzero"`
	SEDS       validtypes.ValidFloat `json:"SEDS,omitzero"`
	SEDS_NCL   validtypes.ValidFloat `json:"SEDS_NCL,omitzero"`
	SEDS_NCU   validtypes.ValidFloat `json:"SEDS_NCU,omitzero"`
	SEDS_BCL   validtypes.ValidFloat `json:"SEDS_BCL,omitzero"`
	SEDS_BCU   validtypes.ValidFloat `json:"SEDS_BCU,omitzero"`
	EDI        validtypes.ValidFloat `json:"EDI,omitzero"`
	EDI_NCL    validtypes.ValidFloat `json:"EDI_NCL,omitzero"`
	EDI_NCU    validtypes.ValidFloat `json:"EDI_NCU,omitzero"`
	EDI_BCL    validtypes.ValidFloat `json:"EDI_BCL,omitzero"`
	EDI_BCU    validtypes.ValidFloat `json:"EDI_BCU,omitzero"`
	SEDI       validtypes.ValidFloat `json:"SEDI,omitzero"`
	SEDI_NCL   validtypes.ValidFloat `json:"SEDI_NCL,omitzero"`
	SEDI_NCU   validtypes.ValidFloat `json:"SEDI_NCU,omitzero"`
	SEDI_BCL   validtypes.ValidFloat `json:"SEDI_BCL,omitzero"`
	SEDI_BCU   validtypes.ValidFloat `json:"SEDI_BCU,omitzero"`
	BAGSS      validtypes.ValidFloat `json:"BAGSS,omitzero"`
	BAGSS_BCL  validtypes.ValidFloat `json:"BAGSS_BCL,omitzero"`
	BAGSS_BCU  validtypes.ValidFloat `json:"BAGSS_BCU,omitzero"`
	HSS_EC     validtypes.ValidFloat `json:"HSS_EC,omitzero"`
	HSS_EC_BCL validtypes.ValidFloat `json:"HSS_EC_BCL,omitzero"`
	HSS_EC_BCU validtypes.ValidFloat `json:"HSS_EC_BCU,omitzero"`
	EC_VALUE   validtypes.ValidFloat `json:"EC_VALUE,omitzero"`
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
	CRPS_EMP_FAIR    validtypes.ValidFloat `json:"CRPS_EMP_FAIR,omitzero"`
	SPREAD_MD        validtypes.ValidFloat `json:"SPREAD_MD,omitzero"`
	MAE              validtypes.ValidFloat `json:"MAE,omitzero"`
	MAE_OERR         validtypes.ValidFloat `json:"MAE_OERR,omitzero"`
	BIAS_RATIO       validtypes.ValidFloat `json:"BIAS_RATIO,omitzero"`
	N_GE_OBS         validtypes.ValidInt   `json:"N_GE_OBS,omitzero"`
	ME_GE_OBS        validtypes.ValidFloat `json:"ME_GE_OBS,omitzero"`
	N_LT_OBS         validtypes.ValidInt   `json:"N_LT_OBS,omitzero"`
	ME_LT_OBS        validtypes.ValidFloat `json:"ME_LT_OBS,omitzero"`
	IGN_CONV_OERR    validtypes.ValidFloat `json:"IGN_CONV_OERR,omitzero"`
	IGN_CORR_OERR    validtypes.ValidFloat `json:"IGN_CORR_OERR,omitzero"`
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
	TOTAL            validtypes.ValidInt    `json:"TOTAL,omitzero"`
	INDEX            validtypes.ValidInt    `json:"INDEX,omitzero"`
	OBS_SID          validtypes.ValidString `json:"OBS_SID,omitzero"`
	OBS_LAT          validtypes.ValidFloat  `json:"OBS_LAT,omitzero"`
	OBS_LON          validtypes.ValidFloat  `json:"OBS_LON,omitzero"`
	OBS_LVL          validtypes.ValidFloat  `json:"OBS_LVL,omitzero"`
	OBS_ELV          validtypes.ValidFloat  `json:"OBS_ELV,omitzero"`
	FCST             validtypes.ValidFloat  `json:"FCST,omitzero"`
	OBS              validtypes.ValidFloat  `json:"OBS,omitzero"`
	OBS_QC           validtypes.ValidString `json:"OBS_QC,omitzero"`
	OBS_CLIMO_MEAN   validtypes.ValidFloat  `json:"OBS_CLIMO_MEAN,omitzero"`
	OBS_CLIMO_STDEV  validtypes.ValidFloat  `json:"OBS_CLIMO_STDEV,omitzero"`
	OBS_CLIMO_CDF    validtypes.ValidFloat  `json:"OBS_CLIMO_CDF,omitzero"`
	FCST_CLIMO_MEAN  validtypes.ValidFloat  `json:"FCST_CLIMO_MEAN,omitzero"`
	FCST_CLIMO_STDEV validtypes.ValidFloat  `json:"FCST_CLIMO_STDEV,omitzero"`
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
	OBS_CLIMO_MEAN   validtypes.ValidFloat  `json:"OBS_CLIMO_MEAN,omitzero"`
	SPREAD           validtypes.ValidFloat  `json:"SPREAD,omitzero"`
	ENS_MEAN_OERR    validtypes.ValidInt    `json:"ENS_MEAN_OERR,omitzero"`
	SPREAD_OERR      validtypes.ValidFloat  `json:"SPREAD_OERR,omitzero"`
	SPREAD_PLUS_OERR validtypes.ValidFloat  `json:"SPREAD_PLUS_OERR,omitzero"`
	OBS_CLIMO_STDEV  validtypes.ValidFloat  `json:"OBS_CLIMO_STDEV,omitzero"`
	FCST_CLIMO_MEAN  validtypes.ValidFloat  `json:"FCST_CLIMO_MEAN,omitzero"`
	FCST_CLIMO_STDEV validtypes.ValidFloat  `json:"FCST_CLIMO_STDEV,omitzero"`
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

type STAT_SEEPS_MPR_data struct {
	OBS_SID  validtypes.ValidString `json:"OBS_SID,omitzero"`
	OBS_LAT  validtypes.ValidFloat  `json:"OBS_LAT,omitzero"`
	OBS_LON  validtypes.ValidFloat  `json:"OBS_LON,omitzero"`
	FCST     validtypes.ValidFloat  `json:"FCST,omitzero"`
	OBS      validtypes.ValidFloat  `json:"OBS,omitzero"`
	OBS_QC   validtypes.ValidString `json:"OBS_QC,omitzero"`
	FCST_CAT validtypes.ValidInt    `json:"FCST_CAT,omitzero"`
	OBS_CAT  validtypes.ValidInt    `json:"OBS_CAT,omitzero"`
	P1       validtypes.ValidFloat  `json:"P1,omitzero"`
	P2       validtypes.ValidFloat  `json:"P2,omitzero"`
	T1       validtypes.ValidFloat  `json:"T1,omitzero"`
	T2       validtypes.ValidFloat  `json:"T2,omitzero"`
	SEEPS    validtypes.ValidFloat  `json:"SEEPS,omitzero"`
}

type STAT_SEEPS_data struct {
	TOTAL     validtypes.ValidInt   `json:"TOTAL,omitzero"`
	ODFL      validtypes.ValidFloat `json:"ODFL,omitzero"`
	ODFH      validtypes.ValidFloat `json:"ODFH,omitzero"`
	OLFD      validtypes.ValidFloat `json:"OLFD,omitzero"`
	OLFH      validtypes.ValidFloat `json:"OLFH,omitzero"`
	OHFD      validtypes.ValidFloat `json:"OHFD,omitzero"`
	OHFL      validtypes.ValidFloat `json:"OHFL,omitzero"`
	PF1       validtypes.ValidFloat `json:"PF1,omitzero"`
	PF2       validtypes.ValidFloat `json:"PF2,omitzero"`
	PF3       validtypes.ValidFloat `json:"PF3,omitzero"`
	PV1       validtypes.ValidFloat `json:"PV1,omitzero"`
	PV2       validtypes.ValidFloat `json:"PV2,omitzero"`
	PV3       validtypes.ValidFloat `json:"PV3,omitzero"`
	MEAN_FCST validtypes.ValidFloat `json:"MEAN_FCST,omitzero"`
	MEAN_OBS  validtypes.ValidFloat `json:"MEAN_OBS,omitzero"`
	SEEPS     validtypes.ValidFloat `json:"SEEPS,omitzero"`
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
	TOTAL        validtypes.ValidInt   `json:"TOTAL,omitzero"`
	UFABAR       validtypes.ValidFloat `json:"UFABAR,omitzero"`
	VFABAR       validtypes.ValidFloat `json:"VFABAR,omitzero"`
	UOABAR       validtypes.ValidFloat `json:"UOABAR,omitzero"`
	VOABAR       validtypes.ValidFloat `json:"VOABAR,omitzero"`
	UVFOABAR     validtypes.ValidFloat `json:"UVFOABAR,omitzero"`
	UVFFABAR     validtypes.ValidFloat `json:"UVFFABAR,omitzero"`
	UVOOABAR     validtypes.ValidFloat `json:"UVOOABAR,omitzero"`
	FA_SPEED_BAR validtypes.ValidFloat `json:"FA_SPEED_BAR,omitzero"`
	OA_SPEED_BAR validtypes.ValidFloat `json:"OA_SPEED_BAR,omitzero"`
	TOTAL_DIR    validtypes.ValidFloat `json:"TOTAL_DIR,omitzero"`
	DIRA_ME      validtypes.ValidFloat `json:"DIRA_ME,omitzero"`
	DIRA_MAE     validtypes.ValidFloat `json:"DIRA_MAE,omitzero"`
	DIRA_MSE     validtypes.ValidFloat `json:"DIRA_MSE,omitzero"`
}

type STAT_VCNT_data struct {
	TOTAL                validtypes.ValidInt   `json:"TOTAL,omitzero"`
	FBAR                 validtypes.ValidFloat `json:"FBAR,omitzero"`
	FBAR_BCL             validtypes.ValidFloat `json:"FBAR_BCL,omitzero"`
	FBAR_BCU             validtypes.ValidFloat `json:"FBAR_BCU,omitzero"`
	OBAR                 validtypes.ValidFloat `json:"OBAR,omitzero"`
	OBAR_BCL             validtypes.ValidFloat `json:"OBAR_BCL,omitzero"`
	OBAR_BCU             validtypes.ValidFloat `json:"OBAR_BCU,omitzero"`
	FS_RMS               validtypes.ValidFloat `json:"FS_RMS,omitzero"`
	FS_RMS_BCL           validtypes.ValidFloat `json:"FS_RMS_BCL,omitzero"`
	FS_RMS_BCU           validtypes.ValidFloat `json:"FS_RMS_BCU,omitzero"`
	OS_RMS               validtypes.ValidFloat `json:"OS_RMS,omitzero"`
	OS_RMS_BCL           validtypes.ValidFloat `json:"OS_RMS_BCL,omitzero"`
	OS_RMS_BCU           validtypes.ValidFloat `json:"OS_RMS_BCU,omitzero"`
	MSVE                 validtypes.ValidFloat `json:"MSVE,omitzero"`
	MSVE_BCL             validtypes.ValidFloat `json:"MSVE_BCL,omitzero"`
	MSVE_BCU             validtypes.ValidFloat `json:"MSVE_BCU,omitzero"`
	RMSVE                validtypes.ValidFloat `json:"RMSVE,omitzero"`
	RMSVE_BCL            validtypes.ValidFloat `json:"RMSVE_BCL,omitzero"`
	RMSVE_BCU            validtypes.ValidFloat `json:"RMSVE_BCU,omitzero"`
	FSTDEV               validtypes.ValidFloat `json:"FSTDEV,omitzero"`
	FSTDEV_BCL           validtypes.ValidFloat `json:"FSTDEV_BCL,omitzero"`
	FSTDEV_BCU           validtypes.ValidFloat `json:"FSTDEV_BCU,omitzero"`
	OSTDEV               validtypes.ValidFloat `json:"OSTDEV,omitzero"`
	OSTDEV_BCL           validtypes.ValidFloat `json:"OSTDEV_BCL,omitzero"`
	OSTDEV_BCU           validtypes.ValidFloat `json:"OSTDEV_BCU,omitzero"`
	FDIR                 validtypes.ValidFloat `json:"FDIR,omitzero"`
	FDIR_BCL             validtypes.ValidFloat `json:"FDIR_BCL,omitzero"`
	FDIR_BCU             validtypes.ValidFloat `json:"FDIR_BCU,omitzero"`
	ODIR                 validtypes.ValidFloat `json:"ODIR,omitzero"`
	ODIR_BCL             validtypes.ValidFloat `json:"ODIR_BCL,omitzero"`
	ODIR_BCU             validtypes.ValidFloat `json:"ODIR_BCU,omitzero"`
	FBAR_SPEED           validtypes.ValidFloat `json:"FBAR_SPEED,omitzero"`
	FBAR_SPEED_BCL       validtypes.ValidFloat `json:"FBAR_SPEED_BCL,omitzero"`
	FBAR_SPEED_BCU       validtypes.ValidFloat `json:"FBAR_SPEED_BCU,omitzero"`
	OBAR_SPEED           validtypes.ValidFloat `json:"OBAR_SPEED,omitzero"`
	OBAR_SPEED_BCL       validtypes.ValidFloat `json:"OBAR_SPEED_BCL,omitzero"`
	OBAR_SPEED_BCU       validtypes.ValidFloat `json:"OBAR_SPEED_BCU,omitzero"`
	VDIFF_SPEED          validtypes.ValidFloat `json:"VDIFF_SPEED,omitzero"`
	VDIFF_SPEED_BCL      validtypes.ValidFloat `json:"VDIFF_SPEED_BCL,omitzero"`
	VDIFF_SPEED_BCU      validtypes.ValidFloat `json:"VDIFF_SPEED_BCU,omitzero"`
	VDIFF_DIR            validtypes.ValidFloat `json:"VDIFF_DIR,omitzero"`
	VDIFF_DIR_BCL        validtypes.ValidFloat `json:"VDIFF_DIR_BCL,omitzero"`
	VDIFF_DIR_BCU        validtypes.ValidFloat `json:"VDIFF_DIR_BCU,omitzero"`
	SPEED_ERR            validtypes.ValidFloat `json:"SPEED_ERR,omitzero"`
	SPEED_ERR_BCL        validtypes.ValidFloat `json:"SPEED_ERR_BCL,omitzero"`
	SPEED_ERR_BCU        validtypes.ValidFloat `json:"SPEED_ERR_BCU,omitzero"`
	SPEED_ABSERR         validtypes.ValidFloat `json:"SPEED_ABSERR,omitzero"`
	SPEED_ABSERR_BCL     validtypes.ValidFloat `json:"SPEED_ABSERR_BCL,omitzero"`
	SPEED_ABSERR_BCU     validtypes.ValidFloat `json:"SPEED_ABSERR_BCU,omitzero"`
	DIR_ERR              validtypes.ValidFloat `json:"DIR_ERR,omitzero"`
	DIR_ERR_BCL          validtypes.ValidFloat `json:"DIR_ERR_BCL,omitzero"`
	DIR_ERR_BCU          validtypes.ValidFloat `json:"DIR_ERR_BCU,omitzero"`
	DIR_ABSERR           validtypes.ValidFloat `json:"DIR_ABSERR,omitzero"`
	DIR_ABSERR_BCL       validtypes.ValidFloat `json:"DIR_ABSERR_BCL,omitzero"`
	DIR_ABSERR_BCU       validtypes.ValidFloat `json:"DIR_ABSERR_BCU,omitzero"`
	ANOM_CORR            validtypes.ValidFloat `json:"ANOM_CORR,omitzero"`
	ANOM_CORR_NCL        validtypes.ValidFloat `json:"ANOM_CORR_NCL,omitzero"`
	ANOM_CORR_NCU        validtypes.ValidFloat `json:"ANOM_CORR_NCU,omitzero"`
	ANOM_CORR_BCL        validtypes.ValidFloat `json:"ANOM_CORR_BCL,omitzero"`
	ANOM_CORR_BCU        validtypes.ValidFloat `json:"ANOM_CORR_BCU,omitzero"`
	ANOM_CORR_UNCNTR     validtypes.ValidFloat `json:"ANOM_CORR_UNCNTR,omitzero"`
	ANOM_CORR_UNCNTR_BCL validtypes.ValidFloat `json:"ANOM_CORR_UNCNTR_BCL,omitzero"`
	ANOM_CORR_UNCNTR_BCU validtypes.ValidFloat `json:"ANOM_CORR_UNCNTR_BCU,omitzero"`
	TOTAL_DIR            validtypes.ValidFloat `json:"TOTAL_DIR,omitzero"`
	DIR_ME               validtypes.ValidFloat `json:"DIR_ME,omitzero"`
	DIR_ME_BCL           validtypes.ValidFloat `json:"DIR_ME_BCL,omitzero"`
	DIR_ME_BCU           validtypes.ValidFloat `json:"DIR_ME_BCU,omitzero"`
	DIR_MAE              validtypes.ValidFloat `json:"DIR_MAE,omitzero"`
	DIR_MAE_BCL          validtypes.ValidFloat `json:"DIR_MAE_BCL,omitzero"`
	DIR_MAE_BCU          validtypes.ValidFloat `json:"DIR_MAE_BCU,omitzero"`
	DIR_MSE              validtypes.ValidFloat `json:"DIR_MSE,omitzero"`
	DIR_MSE_BCL          validtypes.ValidFloat `json:"DIR_MSE_BCL,omitzero"`
	DIR_MSE_BCU          validtypes.ValidFloat `json:"DIR_MSE_BCU,omitzero"`
	DIR_RMSE             validtypes.ValidFloat `json:"DIR_RMSE,omitzero"`
	DIR_RMSE_BCL         validtypes.ValidFloat `json:"DIR_RMSE_BCL,omitzero"`
	DIR_RMSE_BCU         validtypes.ValidFloat `json:"DIR_RMSE_BCU,omitzero"`
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
	TOTAL_DIR   validtypes.ValidFloat `json:"TOTAL_DIR,omitzero"`
	DIR_ME      validtypes.ValidFloat `json:"DIR_ME,omitzero"`
	DIR_MAE     validtypes.ValidFloat `json:"DIR_MAE,omitzero"`
	DIR_MSE     validtypes.ValidFloat `json:"DIR_MSE,omitzero"`
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

type TCST_TCDIAG_data struct {
	TOTAL        validtypes.ValidInt    `json:"TOTAL,omitzero"`
	INDEX        validtypes.ValidInt    `json:"INDEX,omitzero"`
	DIAG_SOURCE  validtypes.ValidFloat  `json:"DIAG_SOURCE,omitzero"`
	TRACK_SOURCE validtypes.ValidString `json:"TRACK_SOURCE,omitzero"`
	FIELD_SOURCE validtypes.ValidString `json:"FIELD_SOURCE,omitzero"`
	DIAG         map[string]interface{} `json:"DIAG,omitzero"`
	INIT         validtypes.ValidInt    `json:"INIT,omitzero"`
}

type TCST_TCMPR_data struct {
	TOTAL          validtypes.ValidInt    `json:"TOTAL,omitzero"`
	INDEX          validtypes.ValidInt    `json:"INDEX,omitzero"`
	LEVEL          validtypes.ValidString `json:"LEVEL,omitzero"`
	WATCH_WARN     validtypes.ValidString `json:"WATCH_WARN,omitzero"`
	INITIALS       validtypes.ValidString `json:"INITIALS,omitzero"`
	ALAT           validtypes.ValidFloat  `json:"ALAT,omitzero"`
	ALON           validtypes.ValidFloat  `json:"ALON,omitzero"`
	BLAT           validtypes.ValidFloat  `json:"BLAT,omitzero"`
	BLON           validtypes.ValidFloat  `json:"BLON,omitzero"`
	TK_ERR         validtypes.ValidFloat  `json:"TK_ERR,omitzero"`
	X_ERR          validtypes.ValidFloat  `json:"X_ERR,omitzero"`
	Y_ERR          validtypes.ValidFloat  `json:"Y_ERR,omitzero"`
	ALTK_ERR       validtypes.ValidFloat  `json:"ALTK_ERR,omitzero"`
	CRTK_ERR       validtypes.ValidFloat  `json:"CRTK_ERR,omitzero"`
	ADLAND         validtypes.ValidFloat  `json:"ADLAND,omitzero"`
	BDLAND         validtypes.ValidFloat  `json:"BDLAND,omitzero"`
	AMSLP          validtypes.ValidFloat  `json:"AMSLP,omitzero"`
	BMSLP          validtypes.ValidFloat  `json:"BMSLP,omitzero"`
	AMAX_WIND      validtypes.ValidFloat  `json:"AMAX_WIND,omitzero"`
	BMAX_WIND      validtypes.ValidFloat  `json:"BMAX_WIND,omitzero"`
	AAL_WIND_34    validtypes.ValidFloat  `json:"AAL_WIND_34,omitzero"`
	BAL_WIND_34    validtypes.ValidFloat  `json:"BAL_WIND_34,omitzero"`
	ANE_WIND_34    validtypes.ValidFloat  `json:"ANE_WIND_34,omitzero"`
	BNE_WIND_34    validtypes.ValidFloat  `json:"BNE_WIND_34,omitzero"`
	ASE_WIND_34    validtypes.ValidFloat  `json:"ASE_WIND_34,omitzero"`
	BSE_WIND_34    validtypes.ValidFloat  `json:"BSE_WIND_34,omitzero"`
	ASW_WIND_34    validtypes.ValidFloat  `json:"ASW_WIND_34,omitzero"`
	BSW_WIND_34    validtypes.ValidFloat  `json:"BSW_WIND_34,omitzero"`
	ANW_WIND_34    validtypes.ValidFloat  `json:"ANW_WIND_34,omitzero"`
	BNW_WIND_34    validtypes.ValidFloat  `json:"BNW_WIND_34,omitzero"`
	AAL_WIND_50    validtypes.ValidFloat  `json:"AAL_WIND_50,omitzero"`
	BAL_WIND_50    validtypes.ValidFloat  `json:"BAL_WIND_50,omitzero"`
	ANE_WIND_50    validtypes.ValidFloat  `json:"ANE_WIND_50,omitzero"`
	BNE_WIND_50    validtypes.ValidFloat  `json:"BNE_WIND_50,omitzero"`
	ASE_WIND_50    validtypes.ValidFloat  `json:"ASE_WIND_50,omitzero"`
	BSE_WIND_50    validtypes.ValidFloat  `json:"BSE_WIND_50,omitzero"`
	ASW_WIND_50    validtypes.ValidFloat  `json:"ASW_WIND_50,omitzero"`
	BSW_WIND_50    validtypes.ValidFloat  `json:"BSW_WIND_50,omitzero"`
	ANW_WIND_50    validtypes.ValidFloat  `json:"ANW_WIND_50,omitzero"`
	BNW_WIND_50    validtypes.ValidFloat  `json:"BNW_WIND_50,omitzero"`
	AAL_WIND_64    validtypes.ValidFloat  `json:"AAL_WIND_64,omitzero"`
	BAL_WIND_64    validtypes.ValidFloat  `json:"BAL_WIND_64,omitzero"`
	ANE_WIND_64    validtypes.ValidFloat  `json:"ANE_WIND_64,omitzero"`
	BNE_WIND_64    validtypes.ValidFloat  `json:"BNE_WIND_64,omitzero"`
	ASE_WIND_64    validtypes.ValidFloat  `json:"ASE_WIND_64,omitzero"`
	BSE_WIND_64    validtypes.ValidFloat  `json:"BSE_WIND_64,omitzero"`
	ASW_WIND_64    validtypes.ValidFloat  `json:"ASW_WIND_64,omitzero"`
	BSW_WIND_64    validtypes.ValidFloat  `json:"BSW_WIND_64,omitzero"`
	ANW_WIND_64    validtypes.ValidFloat  `json:"ANW_WIND_64,omitzero"`
	BNW_WIND_64    validtypes.ValidFloat  `json:"BNW_WIND_64,omitzero"`
	ARADP          validtypes.ValidString `json:"ARADP,omitzero"`
	BRADP          validtypes.ValidFloat  `json:"BRADP,omitzero"`
	ARRP           validtypes.ValidInt    `json:"ARRP,omitzero"`
	BRRP           validtypes.ValidFloat  `json:"BRRP,omitzero"`
	AMRD           validtypes.ValidInt    `json:"AMRD,omitzero"`
	BMRD           validtypes.ValidFloat  `json:"BMRD,omitzero"`
	AGUSTS         validtypes.ValidInt    `json:"AGUSTS,omitzero"`
	BGUSTS         validtypes.ValidFloat  `json:"BGUSTS,omitzero"`
	AEYE           validtypes.ValidInt    `json:"AEYE,omitzero"`
	BEYE           validtypes.ValidFloat  `json:"BEYE,omitzero"`
	ADIR           validtypes.ValidInt    `json:"ADIR,omitzero"`
	BDIR           validtypes.ValidFloat  `json:"BDIR,omitzero"`
	ASPEED         validtypes.ValidInt    `json:"ASPEED,omitzero"`
	BSPEED         validtypes.ValidFloat  `json:"BSPEED,omitzero"`
	ADEPTH         validtypes.ValidInt    `json:"ADEPTH,omitzero"`
	BDEPTH         validtypes.ValidFloat  `json:"BDEPTH,omitzero"`
	NUM_MEMBERS    validtypes.ValidFloat  `json:"NUM_MEMBERS,omitzero"`
	TRACK_SPREAD   validtypes.ValidFloat  `json:"TRACK_SPREAD,omitzero"`
	TRACK_STDEV    validtypes.ValidFloat  `json:"TRACK_STDEV,omitzero"`
	MSLP_STDEV     validtypes.ValidFloat  `json:"MSLP_STDEV,omitzero"`
	MAX_WIND_STDEV validtypes.ValidFloat  `json:"MAX_WIND_STDEV,omitzero"`
	INIT           validtypes.ValidInt    `json:"INIT,omitzero"`
}

// fillStructure functions
func (s *MODE_CTS_data) fill(fields []string) {
	s.FIELD.UnmarshalText([]byte(fields[0]))
	s.TOTAL.UnmarshalText([]byte(fields[1]))
	s.FY_OY.UnmarshalText([]byte(fields[2]))
	s.FY_ON.UnmarshalText([]byte(fields[3]))
	s.FN_OY.UnmarshalText([]byte(fields[4]))
	s.FN_ON.UnmarshalText([]byte(fields[5]))
	s.BASER.UnmarshalText([]byte(fields[6]))
	s.FMEAN.UnmarshalText([]byte(fields[7]))
	s.ACC.UnmarshalText([]byte(fields[8]))
	s.FBIAS.UnmarshalText([]byte(fields[9]))
	s.PODY.UnmarshalText([]byte(fields[10]))
	s.PODN.UnmarshalText([]byte(fields[11]))
	s.POFD.UnmarshalText([]byte(fields[12]))
	s.FAR.UnmarshalText([]byte(fields[13]))
	s.CSI.UnmarshalText([]byte(fields[14]))
	s.GSS.UnmarshalText([]byte(fields[15]))
	s.HK.UnmarshalText([]byte(fields[16]))
	s.HSS.UnmarshalText([]byte(fields[17]))
	s.ODDS.UnmarshalText([]byte(fields[18]))
	s.LODDS.UnmarshalText([]byte(fields[19]))
	s.ORSS.UnmarshalText([]byte(fields[20]))
	s.EDS.UnmarshalText([]byte(fields[21]))
	s.SEDS.UnmarshalText([]byte(fields[22]))
	s.EDI.UnmarshalText([]byte(fields[23]))
	s.SEDI.UnmarshalText([]byte(fields[24]))
	s.BAGSS.UnmarshalText([]byte(fields[25]))
}

func (s *MODE_OBJ_data) fill(fields []string) {
	s.OBJECT_ID.UnmarshalText([]byte(fields[0]))
	s.OBJECT_CAT.UnmarshalText([]byte(fields[1]))
	s.CENTROID_X.UnmarshalText([]byte(fields[2]))
	s.CENTROID_Y.UnmarshalText([]byte(fields[3]))
	s.CENTROID_LAT.UnmarshalText([]byte(fields[4]))
	s.CENTROID_LON.UnmarshalText([]byte(fields[5]))
	s.AXIS_ANG.UnmarshalText([]byte(fields[6]))
	s.LENGTH.UnmarshalText([]byte(fields[7]))
	s.WIDTH.UnmarshalText([]byte(fields[8]))
	s.AREA.UnmarshalText([]byte(fields[9]))
	s.AREA_THRESH.UnmarshalText([]byte(fields[10]))
	s.CURVATURE.UnmarshalText([]byte(fields[11]))
	s.CURVATURE_X.UnmarshalText([]byte(fields[12]))
	s.CURVATURE_Y.UnmarshalText([]byte(fields[13]))
	s.COMPLEXITY.UnmarshalText([]byte(fields[14]))
	s.INTENSITY_10.UnmarshalText([]byte(fields[15]))
	s.INTENSITY_25.UnmarshalText([]byte(fields[16]))
	s.INTENSITY_50.UnmarshalText([]byte(fields[17]))
	s.INTENSITY_75.UnmarshalText([]byte(fields[18]))
	s.INTENSITY_90.UnmarshalText([]byte(fields[19]))
	s.INTENSITY_USER.UnmarshalText([]byte(fields[20]))
	s.INTENSITY_SUM.UnmarshalText([]byte(fields[21]))
	s.CENTROID_DIST.UnmarshalText([]byte(fields[22]))
	s.BOUNDARY_DIST.UnmarshalText([]byte(fields[23]))
	s.CONVEX_HULL_DIST.UnmarshalText([]byte(fields[24]))
	s.ANGLE_DIFF.UnmarshalText([]byte(fields[25]))
	s.ASPECT_DIFF.UnmarshalText([]byte(fields[26]))
	s.AREA_RATIO.UnmarshalText([]byte(fields[27]))
	s.INTERSECTION_AREA.UnmarshalText([]byte(fields[28]))
	s.UNION_AREA.UnmarshalText([]byte(fields[29]))
	s.SYMMETRIC_DIFF.UnmarshalText([]byte(fields[30]))
	s.INTERSECTION_OVER_AREA.UnmarshalText([]byte(fields[31]))
	s.CURVATURE_RATIO.UnmarshalText([]byte(fields[32]))
	s.COMPLEXITY_RATIO.UnmarshalText([]byte(fields[33]))
	s.PERCENTILE_INTENSITY_RATIO.UnmarshalText([]byte(fields[34]))
	s.INTEREST.UnmarshalText([]byte(fields[35]))
}

func (s *MTD_2DSINGLE_data) fill(fields []string) {
	s.OBJECT_ID.UnmarshalText([]byte(fields[0]))
	s.OBJECT_CAT.UnmarshalText([]byte(fields[1]))
	s.TIME_INDEX.UnmarshalText([]byte(fields[2]))
	s.AREA.UnmarshalText([]byte(fields[3]))
	s.CENTROID_X.UnmarshalText([]byte(fields[4]))
	s.CENTROID_Y.UnmarshalText([]byte(fields[5]))
	s.CENTROID_LAT.UnmarshalText([]byte(fields[6]))
	s.CENTROID_LON.UnmarshalText([]byte(fields[7]))
	s.AXIS_ANG.UnmarshalText([]byte(fields[8]))
	s.INTENSITY_10.UnmarshalText([]byte(fields[9]))
	s.INTENSITY_25.UnmarshalText([]byte(fields[10]))
	s.INTENSITY_50.UnmarshalText([]byte(fields[11]))
	s.INTENSITY_75.UnmarshalText([]byte(fields[12]))
	s.INTENSITY_90.UnmarshalText([]byte(fields[13]))
	s.INTENSITY_USER.UnmarshalText([]byte(fields[14]))
}

func (s *MTD_3DPAIR_data) fill(fields []string) {
	s.OBJECT_ID.UnmarshalText([]byte(fields[0]))
	s.OBJECT_CAT.UnmarshalText([]byte(fields[1]))
	s.SPACE_CENTROID_DIST.UnmarshalText([]byte(fields[2]))
	s.TIME_CENTROID_DELTA.UnmarshalText([]byte(fields[3]))
	s.AXIS_DIFF.UnmarshalText([]byte(fields[4]))
	s.SPEED_DELTA.UnmarshalText([]byte(fields[5]))
	s.DIRECTION_DIFF.UnmarshalText([]byte(fields[6]))
	s.VOLUME_RATIO.UnmarshalText([]byte(fields[7]))
	s.START_TIME_DELTA.UnmarshalText([]byte(fields[8]))
	s.END_TIME_DELTA.UnmarshalText([]byte(fields[9]))
	s.INTERSECTION_VOLUME.UnmarshalText([]byte(fields[10]))
	s.DURATION_DIFF.UnmarshalText([]byte(fields[11]))
	s.INTEREST.UnmarshalText([]byte(fields[12]))
}

func (s *MTD_3DSINGLE_data) fill(fields []string) {
	s.OBJECT_ID.UnmarshalText([]byte(fields[0]))
	s.OBJECT_CAT.UnmarshalText([]byte(fields[1]))
	s.CENTROID_X.UnmarshalText([]byte(fields[2]))
	s.CENTROID_Y.UnmarshalText([]byte(fields[3]))
	s.CENTROID_T.UnmarshalText([]byte(fields[4]))
	s.CENTROID_LAT.UnmarshalText([]byte(fields[5]))
	s.CENTROID_LON.UnmarshalText([]byte(fields[6]))
	s.X_DOT.UnmarshalText([]byte(fields[7]))
	s.Y_DOT.UnmarshalText([]byte(fields[8]))
	s.AXIS_ANG.UnmarshalText([]byte(fields[9]))
	s.VOLUME.UnmarshalText([]byte(fields[10]))
	s.START_TIME.UnmarshalText([]byte(fields[11]))
	s.END_TIME.UnmarshalText([]byte(fields[12]))
	s.CDIST_TRAVELLED.UnmarshalText([]byte(fields[13]))
	s.INTENSITY_10.UnmarshalText([]byte(fields[14]))
	s.INTENSITY_25.UnmarshalText([]byte(fields[15]))
	s.INTENSITY_50.UnmarshalText([]byte(fields[16]))
	s.INTENSITY_75.UnmarshalText([]byte(fields[17]))
	s.INTENSITY_90.UnmarshalText([]byte(fields[18]))
	s.INTENSITY_USER.UnmarshalText([]byte(fields[19]))
}

func (s *STAT_CNT_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.FBAR.UnmarshalText([]byte(fields[1]))
	s.FBAR_NCL.UnmarshalText([]byte(fields[2]))
	s.FBAR_NCU.UnmarshalText([]byte(fields[3]))
	s.FBAR_BCL.UnmarshalText([]byte(fields[4]))
	s.FBAR_BCU.UnmarshalText([]byte(fields[5]))
	s.FSTDEV.UnmarshalText([]byte(fields[6]))
	s.FSTDEV_NCL.UnmarshalText([]byte(fields[7]))
	s.FSTDEV_NCU.UnmarshalText([]byte(fields[8]))
	s.FSTDEV_BCL.UnmarshalText([]byte(fields[9]))
	s.FSTDEV_BCU.UnmarshalText([]byte(fields[10]))
	s.OBAR.UnmarshalText([]byte(fields[11]))
	s.OBAR_NCL.UnmarshalText([]byte(fields[12]))
	s.OBAR_NCU.UnmarshalText([]byte(fields[13]))
	s.OBAR_BCL.UnmarshalText([]byte(fields[14]))
	s.OBAR_BCU.UnmarshalText([]byte(fields[15]))
	s.OSTDEV.UnmarshalText([]byte(fields[16]))
	s.OSTDEV_NCL.UnmarshalText([]byte(fields[17]))
	s.OSTDEV_NCU.UnmarshalText([]byte(fields[18]))
	s.OSTDEV_BCL.UnmarshalText([]byte(fields[19]))
	s.OSTDEV_BCU.UnmarshalText([]byte(fields[20]))
	s.PR_CORR.UnmarshalText([]byte(fields[21]))
	s.PR_CORR_NCL.UnmarshalText([]byte(fields[22]))
	s.PR_CORR_NCU.UnmarshalText([]byte(fields[23]))
	s.PR_CORR_BCL.UnmarshalText([]byte(fields[24]))
	s.PR_CORR_BCU.UnmarshalText([]byte(fields[25]))
	s.SP_CORR.UnmarshalText([]byte(fields[26]))
	s.KT_CORR.UnmarshalText([]byte(fields[27]))
	s.RANKS.UnmarshalText([]byte(fields[28]))
	s.FRANK_TIES.UnmarshalText([]byte(fields[29]))
	s.ORANK_TIES.UnmarshalText([]byte(fields[30]))
	s.ME.UnmarshalText([]byte(fields[31]))
	s.ME_NCL.UnmarshalText([]byte(fields[32]))
	s.ME_NCU.UnmarshalText([]byte(fields[33]))
	s.ME_BCL.UnmarshalText([]byte(fields[34]))
	s.ME_BCU.UnmarshalText([]byte(fields[35]))
	s.ESTDEV.UnmarshalText([]byte(fields[36]))
	s.ESTDEV_NCL.UnmarshalText([]byte(fields[37]))
	s.ESTDEV_NCU.UnmarshalText([]byte(fields[38]))
	s.ESTDEV_BCL.UnmarshalText([]byte(fields[39]))
	s.ESTDEV_BCU.UnmarshalText([]byte(fields[40]))
	s.MBIAS.UnmarshalText([]byte(fields[41]))
	s.MBIAS_BCL.UnmarshalText([]byte(fields[42]))
	s.MBIAS_BCU.UnmarshalText([]byte(fields[43]))
	s.MAE.UnmarshalText([]byte(fields[44]))
	s.MAE_BCL.UnmarshalText([]byte(fields[45]))
	s.MAE_BCU.UnmarshalText([]byte(fields[46]))
	s.MSE.UnmarshalText([]byte(fields[47]))
	s.MSE_BCL.UnmarshalText([]byte(fields[48]))
	s.MSE_BCU.UnmarshalText([]byte(fields[49]))
	s.BCMSE.UnmarshalText([]byte(fields[50]))
	s.BCMSE_BCL.UnmarshalText([]byte(fields[51]))
	s.BCMSE_BCU.UnmarshalText([]byte(fields[52]))
	s.RMSE.UnmarshalText([]byte(fields[53]))
	s.RMSE_BCL.UnmarshalText([]byte(fields[54]))
	s.RMSE_BCU.UnmarshalText([]byte(fields[55]))
	s.E10.UnmarshalText([]byte(fields[56]))
	s.E10_BCL.UnmarshalText([]byte(fields[57]))
	s.E10_BCU.UnmarshalText([]byte(fields[58]))
	s.E25.UnmarshalText([]byte(fields[59]))
	s.E25_BCL.UnmarshalText([]byte(fields[60]))
	s.E25_BCU.UnmarshalText([]byte(fields[61]))
	s.E50.UnmarshalText([]byte(fields[62]))
	s.E50_BCL.UnmarshalText([]byte(fields[63]))
	s.E50_BCU.UnmarshalText([]byte(fields[64]))
	s.E75.UnmarshalText([]byte(fields[65]))
	s.E75_BCL.UnmarshalText([]byte(fields[66]))
	s.E75_BCU.UnmarshalText([]byte(fields[67]))
	s.E90.UnmarshalText([]byte(fields[68]))
	s.E90_BCL.UnmarshalText([]byte(fields[69]))
	s.E90_BCU.UnmarshalText([]byte(fields[70]))
	s.EIQR.UnmarshalText([]byte(fields[71]))
	s.EIQR_BCL.UnmarshalText([]byte(fields[72]))
	s.EIQR_BCU.UnmarshalText([]byte(fields[73]))
	s.MAD.UnmarshalText([]byte(fields[74]))
	s.MAD_BCL.UnmarshalText([]byte(fields[75]))
	s.MAD_BCU.UnmarshalText([]byte(fields[76]))
	s.ANOM_CORR.UnmarshalText([]byte(fields[77]))
	s.ANOM_CORR_NCL.UnmarshalText([]byte(fields[78]))
	s.ANOM_CORR_NCU.UnmarshalText([]byte(fields[79]))
	s.ANOM_CORR_BCL.UnmarshalText([]byte(fields[80]))
	s.ANOM_CORR_BCU.UnmarshalText([]byte(fields[81]))
	s.ME2.UnmarshalText([]byte(fields[82]))
	s.ME2_BCL.UnmarshalText([]byte(fields[83]))
	s.ME2_BCU.UnmarshalText([]byte(fields[84]))
	s.MSESS.UnmarshalText([]byte(fields[85]))
	s.MSESS_BCL.UnmarshalText([]byte(fields[86]))
	s.MSESS_BCU.UnmarshalText([]byte(fields[87]))
	s.RMSFA.UnmarshalText([]byte(fields[88]))
	s.RMSFA_BCL.UnmarshalText([]byte(fields[89]))
	s.RMSFA_BCU.UnmarshalText([]byte(fields[90]))
	s.RMSOA.UnmarshalText([]byte(fields[91]))
	s.RMSOA_BCL.UnmarshalText([]byte(fields[92]))
	s.RMSOA_BCU.UnmarshalText([]byte(fields[93]))
	s.ANOM_CORR_UNCNTR.UnmarshalText([]byte(fields[94]))
	s.ANOM_CORR_UNCNTR_BCL.UnmarshalText([]byte(fields[95]))
	s.ANOM_CORR_UNCNTR_BCU.UnmarshalText([]byte(fields[96]))
	s.SI.UnmarshalText([]byte(fields[97]))
	s.SI_BCL.UnmarshalText([]byte(fields[98]))
	s.SI_BCU.UnmarshalText([]byte(fields[99]))
}

func (s *STAT_CTC_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.FY_OY.UnmarshalText([]byte(fields[1]))
	s.FY_ON.UnmarshalText([]byte(fields[2]))
	s.FN_OY.UnmarshalText([]byte(fields[3]))
	s.FN_ON.UnmarshalText([]byte(fields[4]))
	s.EC_VALUE.UnmarshalText([]byte(fields[5]))
}

func (s *STAT_CTS_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.BASER.UnmarshalText([]byte(fields[1]))
	s.BASER_NCL.UnmarshalText([]byte(fields[2]))
	s.BASER_NCU.UnmarshalText([]byte(fields[3]))
	s.BASER_BCL.UnmarshalText([]byte(fields[4]))
	s.BASER_BCU.UnmarshalText([]byte(fields[5]))
	s.FMEAN.UnmarshalText([]byte(fields[6]))
	s.FMEAN_NCL.UnmarshalText([]byte(fields[7]))
	s.FMEAN_NCU.UnmarshalText([]byte(fields[8]))
	s.FMEAN_BCL.UnmarshalText([]byte(fields[9]))
	s.FMEAN_BCU.UnmarshalText([]byte(fields[10]))
	s.ACC.UnmarshalText([]byte(fields[11]))
	s.ACC_NCL.UnmarshalText([]byte(fields[12]))
	s.ACC_NCU.UnmarshalText([]byte(fields[13]))
	s.ACC_BCL.UnmarshalText([]byte(fields[14]))
	s.ACC_BCU.UnmarshalText([]byte(fields[15]))
	s.FBIAS.UnmarshalText([]byte(fields[16]))
	s.FBIAS_BCL.UnmarshalText([]byte(fields[17]))
	s.FBIAS_BCU.UnmarshalText([]byte(fields[18]))
	s.PODY.UnmarshalText([]byte(fields[19]))
	s.PODY_NCL.UnmarshalText([]byte(fields[20]))
	s.PODY_NCU.UnmarshalText([]byte(fields[21]))
	s.PODY_BCL.UnmarshalText([]byte(fields[22]))
	s.PODY_BCU.UnmarshalText([]byte(fields[23]))
	s.PODN.UnmarshalText([]byte(fields[24]))
	s.PODN_NCL.UnmarshalText([]byte(fields[25]))
	s.PODN_NCU.UnmarshalText([]byte(fields[26]))
	s.PODN_BCL.UnmarshalText([]byte(fields[27]))
	s.PODN_BCU.UnmarshalText([]byte(fields[28]))
	s.POFD.UnmarshalText([]byte(fields[29]))
	s.POFD_NCL.UnmarshalText([]byte(fields[30]))
	s.POFD_NCU.UnmarshalText([]byte(fields[31]))
	s.POFD_BCL.UnmarshalText([]byte(fields[32]))
	s.POFD_BCU.UnmarshalText([]byte(fields[33]))
	s.FAR.UnmarshalText([]byte(fields[34]))
	s.FAR_NCL.UnmarshalText([]byte(fields[35]))
	s.FAR_NCU.UnmarshalText([]byte(fields[36]))
	s.FAR_BCL.UnmarshalText([]byte(fields[37]))
	s.FAR_BCU.UnmarshalText([]byte(fields[38]))
	s.CSI.UnmarshalText([]byte(fields[39]))
	s.CSI_NCL.UnmarshalText([]byte(fields[40]))
	s.CSI_NCU.UnmarshalText([]byte(fields[41]))
	s.CSI_BCL.UnmarshalText([]byte(fields[42]))
	s.CSI_BCU.UnmarshalText([]byte(fields[43]))
	s.GSS.UnmarshalText([]byte(fields[44]))
	s.GSS_BCL.UnmarshalText([]byte(fields[45]))
	s.GSS_BCU.UnmarshalText([]byte(fields[46]))
	s.HK.UnmarshalText([]byte(fields[47]))
	s.HK_NCL.UnmarshalText([]byte(fields[48]))
	s.HK_NCU.UnmarshalText([]byte(fields[49]))
	s.HK_BCL.UnmarshalText([]byte(fields[50]))
	s.HK_BCU.UnmarshalText([]byte(fields[51]))
	s.HSS.UnmarshalText([]byte(fields[52]))
	s.HSS_BCL.UnmarshalText([]byte(fields[53]))
	s.HSS_BCU.UnmarshalText([]byte(fields[54]))
	s.ODDS.UnmarshalText([]byte(fields[55]))
	s.ODDS_NCL.UnmarshalText([]byte(fields[56]))
	s.ODDS_NCU.UnmarshalText([]byte(fields[57]))
	s.ODDS_BCL.UnmarshalText([]byte(fields[58]))
	s.ODDS_BCU.UnmarshalText([]byte(fields[59]))
	s.LODDS.UnmarshalText([]byte(fields[60]))
	s.LODDS_NCL.UnmarshalText([]byte(fields[61]))
	s.LODDS_NCU.UnmarshalText([]byte(fields[62]))
	s.LODDS_BCL.UnmarshalText([]byte(fields[63]))
	s.LODDS_BCU.UnmarshalText([]byte(fields[64]))
	s.ORSS.UnmarshalText([]byte(fields[65]))
	s.ORSS_NCL.UnmarshalText([]byte(fields[66]))
	s.ORSS_NCU.UnmarshalText([]byte(fields[67]))
	s.ORSS_BCL.UnmarshalText([]byte(fields[68]))
	s.ORSS_BCU.UnmarshalText([]byte(fields[69]))
	s.EDS.UnmarshalText([]byte(fields[70]))
	s.EDS_NCL.UnmarshalText([]byte(fields[71]))
	s.EDS_NCU.UnmarshalText([]byte(fields[72]))
	s.EDS_BCL.UnmarshalText([]byte(fields[73]))
	s.EDS_BCU.UnmarshalText([]byte(fields[74]))
	s.SEDS.UnmarshalText([]byte(fields[75]))
	s.SEDS_NCL.UnmarshalText([]byte(fields[76]))
	s.SEDS_NCU.UnmarshalText([]byte(fields[77]))
	s.SEDS_BCL.UnmarshalText([]byte(fields[78]))
	s.SEDS_BCU.UnmarshalText([]byte(fields[79]))
	s.EDI.UnmarshalText([]byte(fields[80]))
	s.EDI_NCL.UnmarshalText([]byte(fields[81]))
	s.EDI_NCU.UnmarshalText([]byte(fields[82]))
	s.EDI_BCL.UnmarshalText([]byte(fields[83]))
	s.EDI_BCU.UnmarshalText([]byte(fields[84]))
	s.SEDI.UnmarshalText([]byte(fields[85]))
	s.SEDI_NCL.UnmarshalText([]byte(fields[86]))
	s.SEDI_NCU.UnmarshalText([]byte(fields[87]))
	s.SEDI_BCL.UnmarshalText([]byte(fields[88]))
	s.SEDI_BCU.UnmarshalText([]byte(fields[89]))
	s.BAGSS.UnmarshalText([]byte(fields[90]))
	s.BAGSS_BCL.UnmarshalText([]byte(fields[91]))
	s.BAGSS_BCU.UnmarshalText([]byte(fields[92]))
	s.HSS_EC.UnmarshalText([]byte(fields[93]))
	s.HSS_EC_BCL.UnmarshalText([]byte(fields[94]))
	s.HSS_EC_BCU.UnmarshalText([]byte(fields[95]))
	s.EC_VALUE.UnmarshalText([]byte(fields[96]))
}

func (s *STAT_DMAP_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.FY.UnmarshalText([]byte(fields[1]))
	s.OY.UnmarshalText([]byte(fields[2]))
	s.FBIAS.UnmarshalText([]byte(fields[3]))
	s.BADDELEY.UnmarshalText([]byte(fields[4]))
	s.HAUSDORFF.UnmarshalText([]byte(fields[5]))
	s.MED_FO.UnmarshalText([]byte(fields[6]))
	s.MED_OF.UnmarshalText([]byte(fields[7]))
	s.MED_MIN.UnmarshalText([]byte(fields[8]))
	s.MED_MAX.UnmarshalText([]byte(fields[9]))
	s.MED_MEAN.UnmarshalText([]byte(fields[10]))
	s.FOM_FO.UnmarshalText([]byte(fields[11]))
	s.FOM_OF.UnmarshalText([]byte(fields[12]))
	s.FOM_MIN.UnmarshalText([]byte(fields[13]))
	s.FOM_MAX.UnmarshalText([]byte(fields[14]))
	s.FOM_MEAN.UnmarshalText([]byte(fields[15]))
	s.ZHU_FO.UnmarshalText([]byte(fields[16]))
	s.ZHU_OF.UnmarshalText([]byte(fields[17]))
	s.ZHU_MIN.UnmarshalText([]byte(fields[18]))
	s.ZHU_MAX.UnmarshalText([]byte(fields[19]))
	s.ZHU_MEAN.UnmarshalText([]byte(fields[20]))
	s.G.UnmarshalText([]byte(fields[21]))
	s.GBETA.UnmarshalText([]byte(fields[22]))
	s.BETA_VALUE.UnmarshalText([]byte(fields[23]))
}

func (s *STAT_ECLV_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.BASER.UnmarshalText([]byte(fields[1]))
	s.VALUE_BASER.UnmarshalText([]byte(fields[2]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.PTS[key] = value
		}
	}
}

func (s *STAT_ECNT_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.N_ENS.UnmarshalText([]byte(fields[1]))
	s.CRPS.UnmarshalText([]byte(fields[2]))
	s.CRPSS.UnmarshalText([]byte(fields[3]))
	s.IGN.UnmarshalText([]byte(fields[4]))
	s.ME.UnmarshalText([]byte(fields[5]))
	s.RMSE.UnmarshalText([]byte(fields[6]))
	s.SPREAD.UnmarshalText([]byte(fields[7]))
	s.ME_OERR.UnmarshalText([]byte(fields[8]))
	s.RMSE_OERR.UnmarshalText([]byte(fields[9]))
	s.SPREAD_OERR.UnmarshalText([]byte(fields[10]))
	s.SPREAD_PLUS_OERR.UnmarshalText([]byte(fields[11]))
	s.CRPSCL.UnmarshalText([]byte(fields[12]))
	s.CRPS_EMP.UnmarshalText([]byte(fields[13]))
	s.CRPSCL_EMP.UnmarshalText([]byte(fields[14]))
	s.CRPSS_EMP.UnmarshalText([]byte(fields[15]))
	s.CRPS_EMP_FAIR.UnmarshalText([]byte(fields[16]))
	s.SPREAD_MD.UnmarshalText([]byte(fields[17]))
	s.MAE.UnmarshalText([]byte(fields[18]))
	s.MAE_OERR.UnmarshalText([]byte(fields[19]))
	s.BIAS_RATIO.UnmarshalText([]byte(fields[20]))
	s.N_GE_OBS.UnmarshalText([]byte(fields[21]))
	s.ME_GE_OBS.UnmarshalText([]byte(fields[22]))
	s.N_LT_OBS.UnmarshalText([]byte(fields[23]))
	s.ME_LT_OBS.UnmarshalText([]byte(fields[24]))
	s.IGN_CONV_OERR.UnmarshalText([]byte(fields[25]))
	s.IGN_CORR_OERR.UnmarshalText([]byte(fields[26]))
}

func (s *STAT_FHO_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.F_RATE.UnmarshalText([]byte(fields[1]))
	s.H_RATE.UnmarshalText([]byte(fields[2]))
	s.O_RATE.UnmarshalText([]byte(fields[3]))
}

func (s *STAT_GENMPR_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.INDEX.UnmarshalText([]byte(fields[1]))
	s.STORM_ID.UnmarshalText([]byte(fields[2]))
	s.PROB_LEAD.UnmarshalText([]byte(fields[3]))
	s.PROB_VAL.UnmarshalText([]byte(fields[4]))
	s.AGEN_INIT.UnmarshalText([]byte(fields[5]))
	s.AGEN_FHR.UnmarshalText([]byte(fields[6]))
	s.AGEN_LAT.UnmarshalText([]byte(fields[7]))
	s.AGEN_LON.UnmarshalText([]byte(fields[8]))
	s.AGEN_DLAND.UnmarshalText([]byte(fields[9]))
	s.BGEN_LAT.UnmarshalText([]byte(fields[10]))
	s.BGEN_LON.UnmarshalText([]byte(fields[11]))
	s.BGEN_DLAND.UnmarshalText([]byte(fields[12]))
	s.GEN_DIST.UnmarshalText([]byte(fields[13]))
	s.GEN_TDIFF.UnmarshalText([]byte(fields[14]))
	s.INIT_TDIFF.UnmarshalText([]byte(fields[15]))
	s.DEV_CAT.UnmarshalText([]byte(fields[16]))
	s.OPS_CAT.UnmarshalText([]byte(fields[17]))
}

func (s *STAT_GRAD_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.FGBAR.UnmarshalText([]byte(fields[1]))
	s.OGBAR.UnmarshalText([]byte(fields[2]))
	s.MGBAR.UnmarshalText([]byte(fields[3]))
	s.EGBAR.UnmarshalText([]byte(fields[4]))
	s.S1.UnmarshalText([]byte(fields[5]))
	s.S1_OG.UnmarshalText([]byte(fields[6]))
	s.FGOG_RATIO.UnmarshalText([]byte(fields[7]))
	s.DX.UnmarshalText([]byte(fields[8]))
	s.DY.UnmarshalText([]byte(fields[9]))
}

func (s *STAT_ISC_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.TILE_DIM.UnmarshalText([]byte(fields[1]))
	s.TILE_XLL.UnmarshalText([]byte(fields[2]))
	s.TILE_YLL.UnmarshalText([]byte(fields[3]))
	s.NSCALE.UnmarshalText([]byte(fields[4]))
	s.ISCALE.UnmarshalText([]byte(fields[5]))
	s.MSE.UnmarshalText([]byte(fields[6]))
	s.ISC.UnmarshalText([]byte(fields[7]))
	s.FENERGY2.UnmarshalText([]byte(fields[8]))
	s.OENERGY2.UnmarshalText([]byte(fields[9]))
	s.BASER.UnmarshalText([]byte(fields[10]))
	s.FBIAS.UnmarshalText([]byte(fields[11]))
}

func (s *STAT_MCTC_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.CAT[key] = value
		}
	}
	s.EC_VALUE.UnmarshalText([]byte(fields[3]))
}

func (s *STAT_MCTS_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.N_CAT.UnmarshalText([]byte(fields[1]))
	s.ACC.UnmarshalText([]byte(fields[2]))
	s.ACC_NCL.UnmarshalText([]byte(fields[3]))
	s.ACC_NCU.UnmarshalText([]byte(fields[4]))
	s.ACC_BCL.UnmarshalText([]byte(fields[5]))
	s.ACC_BCU.UnmarshalText([]byte(fields[6]))
	s.HK.UnmarshalText([]byte(fields[7]))
	s.HK_BCL.UnmarshalText([]byte(fields[8]))
	s.HK_BCU.UnmarshalText([]byte(fields[9]))
	s.HSS.UnmarshalText([]byte(fields[10]))
	s.HSS_BCL.UnmarshalText([]byte(fields[11]))
	s.HSS_BCU.UnmarshalText([]byte(fields[12]))
	s.GER.UnmarshalText([]byte(fields[13]))
	s.GER_BCL.UnmarshalText([]byte(fields[14]))
	s.GER_BCU.UnmarshalText([]byte(fields[15]))
	s.HSS_EC.UnmarshalText([]byte(fields[16]))
	s.HSS_EC_BCL.UnmarshalText([]byte(fields[17]))
	s.HSS_EC_BCU.UnmarshalText([]byte(fields[18]))
	s.EC_VALUE.UnmarshalText([]byte(fields[19]))
}

func (s *STAT_MPR_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.INDEX.UnmarshalText([]byte(fields[1]))
	s.OBS_SID.UnmarshalText([]byte(fields[2]))
	s.OBS_LAT.UnmarshalText([]byte(fields[3]))
	s.OBS_LON.UnmarshalText([]byte(fields[4]))
	s.OBS_LVL.UnmarshalText([]byte(fields[5]))
	s.OBS_ELV.UnmarshalText([]byte(fields[6]))
	s.FCST.UnmarshalText([]byte(fields[7]))
	s.OBS.UnmarshalText([]byte(fields[8]))
	s.OBS_QC.UnmarshalText([]byte(fields[9]))
	s.OBS_CLIMO_MEAN.UnmarshalText([]byte(fields[10]))
	s.OBS_CLIMO_STDEV.UnmarshalText([]byte(fields[11]))
	s.OBS_CLIMO_CDF.UnmarshalText([]byte(fields[12]))
	s.FCST_CLIMO_MEAN.UnmarshalText([]byte(fields[13]))
	s.FCST_CLIMO_STDEV.UnmarshalText([]byte(fields[14]))
}

func (s *STAT_NBRCNT_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.FBS.UnmarshalText([]byte(fields[1]))
	s.FBS_BCL.UnmarshalText([]byte(fields[2]))
	s.FBS_BCU.UnmarshalText([]byte(fields[3]))
	s.FSS.UnmarshalText([]byte(fields[4]))
	s.FSS_BCL.UnmarshalText([]byte(fields[5]))
	s.FSS_BCU.UnmarshalText([]byte(fields[6]))
	s.AFSS.UnmarshalText([]byte(fields[7]))
	s.AFSS_BCL.UnmarshalText([]byte(fields[8]))
	s.AFSS_BCU.UnmarshalText([]byte(fields[9]))
	s.UFSS.UnmarshalText([]byte(fields[10]))
	s.UFSS_BCL.UnmarshalText([]byte(fields[11]))
	s.UFSS_BCU.UnmarshalText([]byte(fields[12]))
	s.F_RATE.UnmarshalText([]byte(fields[13]))
	s.F_RATE_BCL.UnmarshalText([]byte(fields[14]))
	s.F_RATE_BCU.UnmarshalText([]byte(fields[15]))
	s.O_RATE.UnmarshalText([]byte(fields[16]))
	s.O_RATE_BCL.UnmarshalText([]byte(fields[17]))
	s.O_RATE_BCU.UnmarshalText([]byte(fields[18]))
}

func (s *STAT_NBRCTC_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.FY_OY.UnmarshalText([]byte(fields[1]))
	s.FY_ON.UnmarshalText([]byte(fields[2]))
	s.FN_OY.UnmarshalText([]byte(fields[3]))
	s.FN_ON.UnmarshalText([]byte(fields[4]))
}

func (s *STAT_NBRCTS_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.BASER.UnmarshalText([]byte(fields[1]))
	s.BASER_NCL.UnmarshalText([]byte(fields[2]))
	s.BASER_NCU.UnmarshalText([]byte(fields[3]))
	s.BASER_BCL.UnmarshalText([]byte(fields[4]))
	s.BASER_BCU.UnmarshalText([]byte(fields[5]))
	s.FMEAN.UnmarshalText([]byte(fields[6]))
	s.FMEAN_NCL.UnmarshalText([]byte(fields[7]))
	s.FMEAN_NCU.UnmarshalText([]byte(fields[8]))
	s.FMEAN_BCL.UnmarshalText([]byte(fields[9]))
	s.FMEAN_BCU.UnmarshalText([]byte(fields[10]))
	s.ACC.UnmarshalText([]byte(fields[11]))
	s.ACC_NCL.UnmarshalText([]byte(fields[12]))
	s.ACC_NCU.UnmarshalText([]byte(fields[13]))
	s.ACC_BCL.UnmarshalText([]byte(fields[14]))
	s.ACC_BCU.UnmarshalText([]byte(fields[15]))
	s.FBIAS.UnmarshalText([]byte(fields[16]))
	s.FBIAS_BCL.UnmarshalText([]byte(fields[17]))
	s.FBIAS_BCU.UnmarshalText([]byte(fields[18]))
	s.PODY.UnmarshalText([]byte(fields[19]))
	s.PODY_NCL.UnmarshalText([]byte(fields[20]))
	s.PODY_NCU.UnmarshalText([]byte(fields[21]))
	s.PODY_BCL.UnmarshalText([]byte(fields[22]))
	s.PODY_BCU.UnmarshalText([]byte(fields[23]))
	s.PODN.UnmarshalText([]byte(fields[24]))
	s.PODN_NCL.UnmarshalText([]byte(fields[25]))
	s.PODN_NCU.UnmarshalText([]byte(fields[26]))
	s.PODN_BCL.UnmarshalText([]byte(fields[27]))
	s.PODN_BCU.UnmarshalText([]byte(fields[28]))
	s.POFD.UnmarshalText([]byte(fields[29]))
	s.POFD_NCL.UnmarshalText([]byte(fields[30]))
	s.POFD_NCU.UnmarshalText([]byte(fields[31]))
	s.POFD_BCL.UnmarshalText([]byte(fields[32]))
	s.POFD_BCU.UnmarshalText([]byte(fields[33]))
	s.FAR.UnmarshalText([]byte(fields[34]))
	s.FAR_NCL.UnmarshalText([]byte(fields[35]))
	s.FAR_NCU.UnmarshalText([]byte(fields[36]))
	s.FAR_BCL.UnmarshalText([]byte(fields[37]))
	s.FAR_BCU.UnmarshalText([]byte(fields[38]))
	s.CSI.UnmarshalText([]byte(fields[39]))
	s.CSI_NCL.UnmarshalText([]byte(fields[40]))
	s.CSI_NCU.UnmarshalText([]byte(fields[41]))
	s.CSI_BCL.UnmarshalText([]byte(fields[42]))
	s.CSI_BCU.UnmarshalText([]byte(fields[43]))
	s.GSS.UnmarshalText([]byte(fields[44]))
	s.GSS_BCL.UnmarshalText([]byte(fields[45]))
	s.GSS_BCU.UnmarshalText([]byte(fields[46]))
	s.HK.UnmarshalText([]byte(fields[47]))
	s.HK_NCL.UnmarshalText([]byte(fields[48]))
	s.HK_NCU.UnmarshalText([]byte(fields[49]))
	s.HK_BCL.UnmarshalText([]byte(fields[50]))
	s.HK_BCU.UnmarshalText([]byte(fields[51]))
	s.HSS.UnmarshalText([]byte(fields[52]))
	s.HSS_BCL.UnmarshalText([]byte(fields[53]))
	s.HSS_BCU.UnmarshalText([]byte(fields[54]))
	s.ODDS.UnmarshalText([]byte(fields[55]))
	s.ODDS_NCL.UnmarshalText([]byte(fields[56]))
	s.ODDS_NCU.UnmarshalText([]byte(fields[57]))
	s.ODDS_BCL.UnmarshalText([]byte(fields[58]))
	s.ODDS_BCU.UnmarshalText([]byte(fields[59]))
	s.LODDS.UnmarshalText([]byte(fields[60]))
	s.LODDS_NCL.UnmarshalText([]byte(fields[61]))
	s.LODDS_NCU.UnmarshalText([]byte(fields[62]))
	s.LODDS_BCL.UnmarshalText([]byte(fields[63]))
	s.LODDS_BCU.UnmarshalText([]byte(fields[64]))
	s.ORSS.UnmarshalText([]byte(fields[65]))
	s.ORSS_NCL.UnmarshalText([]byte(fields[66]))
	s.ORSS_NCU.UnmarshalText([]byte(fields[67]))
	s.ORSS_BCL.UnmarshalText([]byte(fields[68]))
	s.ORSS_BCU.UnmarshalText([]byte(fields[69]))
	s.EDS.UnmarshalText([]byte(fields[70]))
	s.EDS_NCL.UnmarshalText([]byte(fields[71]))
	s.EDS_NCU.UnmarshalText([]byte(fields[72]))
	s.EDS_BCL.UnmarshalText([]byte(fields[73]))
	s.EDS_BCU.UnmarshalText([]byte(fields[74]))
	s.SEDS.UnmarshalText([]byte(fields[75]))
	s.SEDS_NCL.UnmarshalText([]byte(fields[76]))
	s.SEDS_NCU.UnmarshalText([]byte(fields[77]))
	s.SEDS_BCL.UnmarshalText([]byte(fields[78]))
	s.SEDS_BCU.UnmarshalText([]byte(fields[79]))
	s.EDI.UnmarshalText([]byte(fields[80]))
	s.EDI_NCL.UnmarshalText([]byte(fields[81]))
	s.EDI_NCU.UnmarshalText([]byte(fields[82]))
	s.EDI_BCL.UnmarshalText([]byte(fields[83]))
	s.EDI_BCU.UnmarshalText([]byte(fields[84]))
	s.SEDI.UnmarshalText([]byte(fields[85]))
	s.SEDI_NCL.UnmarshalText([]byte(fields[86]))
	s.SEDI_NCU.UnmarshalText([]byte(fields[87]))
	s.SEDI_BCL.UnmarshalText([]byte(fields[88]))
	s.SEDI_BCU.UnmarshalText([]byte(fields[89]))
	s.BAGSS.UnmarshalText([]byte(fields[90]))
	s.BAGSS_BCL.UnmarshalText([]byte(fields[91]))
	s.BAGSS_BCU.UnmarshalText([]byte(fields[92]))
}

func (s *STAT_ORANK_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.INDEX.UnmarshalText([]byte(fields[1]))
	s.OBS_SID.UnmarshalText([]byte(fields[2]))
	s.OBS_LAT.UnmarshalText([]byte(fields[3]))
	s.OBS_LON.UnmarshalText([]byte(fields[4]))
	s.OBS_LVL.UnmarshalText([]byte(fields[5]))
	s.OBS_ELV.UnmarshalText([]byte(fields[6]))
	s.OBS.UnmarshalText([]byte(fields[7]))
	s.PIT.UnmarshalText([]byte(fields[8]))
	s.RANK.UnmarshalText([]byte(fields[9]))
	s.N_ENS_VLD.UnmarshalText([]byte(fields[10]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.ENS[key] = value
		}
	}
	s.OBS_QC.UnmarshalText([]byte(fields[13]))
	s.ENS_MEAN.UnmarshalText([]byte(fields[14]))
	s.OBS_CLIMO_MEAN.UnmarshalText([]byte(fields[15]))
	s.SPREAD.UnmarshalText([]byte(fields[16]))
	s.ENS_MEAN_OERR.UnmarshalText([]byte(fields[17]))
	s.SPREAD_OERR.UnmarshalText([]byte(fields[18]))
	s.SPREAD_PLUS_OERR.UnmarshalText([]byte(fields[19]))
	s.OBS_CLIMO_STDEV.UnmarshalText([]byte(fields[20]))
	s.FCST_CLIMO_MEAN.UnmarshalText([]byte(fields[21]))
	s.FCST_CLIMO_STDEV.UnmarshalText([]byte(fields[22]))
}

func (s *STAT_PCT_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.THRESH[key] = value
		}
	}
}

func (s *STAT_PHIST_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.BIN_SIZE.UnmarshalText([]byte(fields[1]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.BIN[key] = value
		}
	}
}

func (s *STAT_PJC_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.THRESH[key] = value
		}
	}
}

func (s *STAT_PRC_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.THRESH[key] = value
		}
	}
}

func (s *STAT_PSTD_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.THRESH[key] = value
		}
	}
	s.BASER_NCL.UnmarshalText([]byte(fields[3]))
	s.BASER_NCU.UnmarshalText([]byte(fields[4]))
	s.RELIABILITY.UnmarshalText([]byte(fields[5]))
	s.RESOLUTION.UnmarshalText([]byte(fields[6]))
	s.UNCERTAINTY.UnmarshalText([]byte(fields[7]))
	s.ROC_AUC.UnmarshalText([]byte(fields[8]))
	s.BRIER.UnmarshalText([]byte(fields[9]))
	s.BRIER_NCL.UnmarshalText([]byte(fields[10]))
	s.BRIER_NCU.UnmarshalText([]byte(fields[11]))
	s.BRIERCL.UnmarshalText([]byte(fields[12]))
	s.BRIERCL_NCL.UnmarshalText([]byte(fields[13]))
	s.BRIERCL_NCU.UnmarshalText([]byte(fields[14]))
	s.BSS.UnmarshalText([]byte(fields[15]))
	s.BSS_SMPL.UnmarshalText([]byte(fields[16]))
	s.THRESH_I.UnmarshalText([]byte(fields[17]))
}

func (s *STAT_RELP_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.ENS[key] = value
		}
	}
}

func (s *STAT_RHIST_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.RANK[key] = value
		}
	}
}

func (s *STAT_RPS_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.N_PROB.UnmarshalText([]byte(fields[1]))
	s.RPS_REL.UnmarshalText([]byte(fields[2]))
	s.RPS_RES.UnmarshalText([]byte(fields[3]))
	s.RPS_UNC.UnmarshalText([]byte(fields[4]))
	s.RPS.UnmarshalText([]byte(fields[5]))
	s.RPSS.UnmarshalText([]byte(fields[6]))
	s.RPSS_SMPL.UnmarshalText([]byte(fields[7]))
	s.RPS_COMP.UnmarshalText([]byte(fields[8]))
}

func (s *STAT_SAL1L2_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.FABAR.UnmarshalText([]byte(fields[1]))
	s.OABAR.UnmarshalText([]byte(fields[2]))
	s.FOABAR.UnmarshalText([]byte(fields[3]))
	s.FFABAR.UnmarshalText([]byte(fields[4]))
	s.OOABAR.UnmarshalText([]byte(fields[5]))
	s.MAE.UnmarshalText([]byte(fields[6]))
}

func (s *STAT_SEEPS_MPR_data) fill(fields []string) {
	s.OBS_SID.UnmarshalText([]byte(fields[0]))
	s.OBS_LAT.UnmarshalText([]byte(fields[1]))
	s.OBS_LON.UnmarshalText([]byte(fields[2]))
	s.FCST.UnmarshalText([]byte(fields[3]))
	s.OBS.UnmarshalText([]byte(fields[4]))
	s.OBS_QC.UnmarshalText([]byte(fields[5]))
	s.FCST_CAT.UnmarshalText([]byte(fields[6]))
	s.OBS_CAT.UnmarshalText([]byte(fields[7]))
	s.P1.UnmarshalText([]byte(fields[8]))
	s.P2.UnmarshalText([]byte(fields[9]))
	s.T1.UnmarshalText([]byte(fields[10]))
	s.T2.UnmarshalText([]byte(fields[11]))
	s.SEEPS.UnmarshalText([]byte(fields[12]))
}

func (s *STAT_SEEPS_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.ODFL.UnmarshalText([]byte(fields[1]))
	s.ODFH.UnmarshalText([]byte(fields[2]))
	s.OLFD.UnmarshalText([]byte(fields[3]))
	s.OLFH.UnmarshalText([]byte(fields[4]))
	s.OHFD.UnmarshalText([]byte(fields[5]))
	s.OHFL.UnmarshalText([]byte(fields[6]))
	s.PF1.UnmarshalText([]byte(fields[7]))
	s.PF2.UnmarshalText([]byte(fields[8]))
	s.PF3.UnmarshalText([]byte(fields[9]))
	s.PV1.UnmarshalText([]byte(fields[10]))
	s.PV2.UnmarshalText([]byte(fields[11]))
	s.PV3.UnmarshalText([]byte(fields[12]))
	s.MEAN_FCST.UnmarshalText([]byte(fields[13]))
	s.MEAN_OBS.UnmarshalText([]byte(fields[14]))
	s.SEEPS.UnmarshalText([]byte(fields[15]))
}

func (s *STAT_SL1L2_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.FBAR.UnmarshalText([]byte(fields[1]))
	s.OBAR.UnmarshalText([]byte(fields[2]))
	s.FOBAR.UnmarshalText([]byte(fields[3]))
	s.FFBAR.UnmarshalText([]byte(fields[4]))
	s.OOBAR.UnmarshalText([]byte(fields[5]))
	s.MAE.UnmarshalText([]byte(fields[6]))
}

func (s *STAT_SSIDX_data) fill(fields []string) {
	s.FCST_MODEL.UnmarshalText([]byte(fields[0]))
	s.REF_MODEL.UnmarshalText([]byte(fields[1]))
	s.N_INIT.UnmarshalText([]byte(fields[2]))
	s.N_TERM.UnmarshalText([]byte(fields[3]))
	s.N_VLD.UnmarshalText([]byte(fields[4]))
	s.SS_INDEX.UnmarshalText([]byte(fields[5]))
}

func (s *STAT_SSVAR_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.N_BIN.UnmarshalText([]byte(fields[1]))
	s.BIN_I.UnmarshalText([]byte(fields[2]))
	s.BIN_N.UnmarshalText([]byte(fields[3]))
	s.VAR_MIN.UnmarshalText([]byte(fields[4]))
	s.VAR_MAX.UnmarshalText([]byte(fields[5]))
	s.VAR_MEAN.UnmarshalText([]byte(fields[6]))
	s.FBAR.UnmarshalText([]byte(fields[7]))
	s.OBAR.UnmarshalText([]byte(fields[8]))
	s.FOBAR.UnmarshalText([]byte(fields[9]))
	s.FFBAR.UnmarshalText([]byte(fields[10]))
	s.OOBAR.UnmarshalText([]byte(fields[11]))
	s.FBAR_NCL.UnmarshalText([]byte(fields[12]))
	s.FBAR_NCU.UnmarshalText([]byte(fields[13]))
	s.FSTDEV.UnmarshalText([]byte(fields[14]))
	s.FSTDEV_NCL.UnmarshalText([]byte(fields[15]))
	s.FSTDEV_NCU.UnmarshalText([]byte(fields[16]))
	s.OBAR_NCL.UnmarshalText([]byte(fields[17]))
	s.OBAR_NCU.UnmarshalText([]byte(fields[18]))
	s.OSTDEV.UnmarshalText([]byte(fields[19]))
	s.OSTDEV_NCL.UnmarshalText([]byte(fields[20]))
	s.OSTDEV_NCU.UnmarshalText([]byte(fields[21]))
	s.PR_CORR.UnmarshalText([]byte(fields[22]))
	s.PR_CORR_NCL.UnmarshalText([]byte(fields[23]))
	s.PR_CORR_NCU.UnmarshalText([]byte(fields[24]))
	s.ME.UnmarshalText([]byte(fields[25]))
	s.ME_NCL.UnmarshalText([]byte(fields[26]))
	s.ME_NCU.UnmarshalText([]byte(fields[27]))
	s.ESTDEV.UnmarshalText([]byte(fields[28]))
	s.ESTDEV_NCL.UnmarshalText([]byte(fields[29]))
	s.ESTDEV_NCU.UnmarshalText([]byte(fields[30]))
	s.MBIAS.UnmarshalText([]byte(fields[31]))
	s.MSE.UnmarshalText([]byte(fields[32]))
	s.BCMSE.UnmarshalText([]byte(fields[33]))
	s.RMSE.UnmarshalText([]byte(fields[34]))
}

func (s *STAT_VAL1L2_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.UFABAR.UnmarshalText([]byte(fields[1]))
	s.VFABAR.UnmarshalText([]byte(fields[2]))
	s.UOABAR.UnmarshalText([]byte(fields[3]))
	s.VOABAR.UnmarshalText([]byte(fields[4]))
	s.UVFOABAR.UnmarshalText([]byte(fields[5]))
	s.UVFFABAR.UnmarshalText([]byte(fields[6]))
	s.UVOOABAR.UnmarshalText([]byte(fields[7]))
	s.FA_SPEED_BAR.UnmarshalText([]byte(fields[8]))
	s.OA_SPEED_BAR.UnmarshalText([]byte(fields[9]))
	s.TOTAL_DIR.UnmarshalText([]byte(fields[10]))
	s.DIRA_ME.UnmarshalText([]byte(fields[11]))
	s.DIRA_MAE.UnmarshalText([]byte(fields[12]))
	s.DIRA_MSE.UnmarshalText([]byte(fields[13]))
}

func (s *STAT_VCNT_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.FBAR.UnmarshalText([]byte(fields[1]))
	s.FBAR_BCL.UnmarshalText([]byte(fields[2]))
	s.FBAR_BCU.UnmarshalText([]byte(fields[3]))
	s.OBAR.UnmarshalText([]byte(fields[4]))
	s.OBAR_BCL.UnmarshalText([]byte(fields[5]))
	s.OBAR_BCU.UnmarshalText([]byte(fields[6]))
	s.FS_RMS.UnmarshalText([]byte(fields[7]))
	s.FS_RMS_BCL.UnmarshalText([]byte(fields[8]))
	s.FS_RMS_BCU.UnmarshalText([]byte(fields[9]))
	s.OS_RMS.UnmarshalText([]byte(fields[10]))
	s.OS_RMS_BCL.UnmarshalText([]byte(fields[11]))
	s.OS_RMS_BCU.UnmarshalText([]byte(fields[12]))
	s.MSVE.UnmarshalText([]byte(fields[13]))
	s.MSVE_BCL.UnmarshalText([]byte(fields[14]))
	s.MSVE_BCU.UnmarshalText([]byte(fields[15]))
	s.RMSVE.UnmarshalText([]byte(fields[16]))
	s.RMSVE_BCL.UnmarshalText([]byte(fields[17]))
	s.RMSVE_BCU.UnmarshalText([]byte(fields[18]))
	s.FSTDEV.UnmarshalText([]byte(fields[19]))
	s.FSTDEV_BCL.UnmarshalText([]byte(fields[20]))
	s.FSTDEV_BCU.UnmarshalText([]byte(fields[21]))
	s.OSTDEV.UnmarshalText([]byte(fields[22]))
	s.OSTDEV_BCL.UnmarshalText([]byte(fields[23]))
	s.OSTDEV_BCU.UnmarshalText([]byte(fields[24]))
	s.FDIR.UnmarshalText([]byte(fields[25]))
	s.FDIR_BCL.UnmarshalText([]byte(fields[26]))
	s.FDIR_BCU.UnmarshalText([]byte(fields[27]))
	s.ODIR.UnmarshalText([]byte(fields[28]))
	s.ODIR_BCL.UnmarshalText([]byte(fields[29]))
	s.ODIR_BCU.UnmarshalText([]byte(fields[30]))
	s.FBAR_SPEED.UnmarshalText([]byte(fields[31]))
	s.FBAR_SPEED_BCL.UnmarshalText([]byte(fields[32]))
	s.FBAR_SPEED_BCU.UnmarshalText([]byte(fields[33]))
	s.OBAR_SPEED.UnmarshalText([]byte(fields[34]))
	s.OBAR_SPEED_BCL.UnmarshalText([]byte(fields[35]))
	s.OBAR_SPEED_BCU.UnmarshalText([]byte(fields[36]))
	s.VDIFF_SPEED.UnmarshalText([]byte(fields[37]))
	s.VDIFF_SPEED_BCL.UnmarshalText([]byte(fields[38]))
	s.VDIFF_SPEED_BCU.UnmarshalText([]byte(fields[39]))
	s.VDIFF_DIR.UnmarshalText([]byte(fields[40]))
	s.VDIFF_DIR_BCL.UnmarshalText([]byte(fields[41]))
	s.VDIFF_DIR_BCU.UnmarshalText([]byte(fields[42]))
	s.SPEED_ERR.UnmarshalText([]byte(fields[43]))
	s.SPEED_ERR_BCL.UnmarshalText([]byte(fields[44]))
	s.SPEED_ERR_BCU.UnmarshalText([]byte(fields[45]))
	s.SPEED_ABSERR.UnmarshalText([]byte(fields[46]))
	s.SPEED_ABSERR_BCL.UnmarshalText([]byte(fields[47]))
	s.SPEED_ABSERR_BCU.UnmarshalText([]byte(fields[48]))
	s.DIR_ERR.UnmarshalText([]byte(fields[49]))
	s.DIR_ERR_BCL.UnmarshalText([]byte(fields[50]))
	s.DIR_ERR_BCU.UnmarshalText([]byte(fields[51]))
	s.DIR_ABSERR.UnmarshalText([]byte(fields[52]))
	s.DIR_ABSERR_BCL.UnmarshalText([]byte(fields[53]))
	s.DIR_ABSERR_BCU.UnmarshalText([]byte(fields[54]))
	s.ANOM_CORR.UnmarshalText([]byte(fields[55]))
	s.ANOM_CORR_NCL.UnmarshalText([]byte(fields[56]))
	s.ANOM_CORR_NCU.UnmarshalText([]byte(fields[57]))
	s.ANOM_CORR_BCL.UnmarshalText([]byte(fields[58]))
	s.ANOM_CORR_BCU.UnmarshalText([]byte(fields[59]))
	s.ANOM_CORR_UNCNTR.UnmarshalText([]byte(fields[60]))
	s.ANOM_CORR_UNCNTR_BCL.UnmarshalText([]byte(fields[61]))
	s.ANOM_CORR_UNCNTR_BCU.UnmarshalText([]byte(fields[62]))
	s.TOTAL_DIR.UnmarshalText([]byte(fields[63]))
	s.DIR_ME.UnmarshalText([]byte(fields[64]))
	s.DIR_ME_BCL.UnmarshalText([]byte(fields[65]))
	s.DIR_ME_BCU.UnmarshalText([]byte(fields[66]))
	s.DIR_MAE.UnmarshalText([]byte(fields[67]))
	s.DIR_MAE_BCL.UnmarshalText([]byte(fields[68]))
	s.DIR_MAE_BCU.UnmarshalText([]byte(fields[69]))
	s.DIR_MSE.UnmarshalText([]byte(fields[70]))
	s.DIR_MSE_BCL.UnmarshalText([]byte(fields[71]))
	s.DIR_MSE_BCU.UnmarshalText([]byte(fields[72]))
	s.DIR_RMSE.UnmarshalText([]byte(fields[73]))
	s.DIR_RMSE_BCL.UnmarshalText([]byte(fields[74]))
	s.DIR_RMSE_BCU.UnmarshalText([]byte(fields[75]))
}

func (s *STAT_VL1L2_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.UFBAR.UnmarshalText([]byte(fields[1]))
	s.VFBAR.UnmarshalText([]byte(fields[2]))
	s.UOBAR.UnmarshalText([]byte(fields[3]))
	s.VOBAR.UnmarshalText([]byte(fields[4]))
	s.UVFOBAR.UnmarshalText([]byte(fields[5]))
	s.UVFFBAR.UnmarshalText([]byte(fields[6]))
	s.UVOOBAR.UnmarshalText([]byte(fields[7]))
	s.F_SPEED_BAR.UnmarshalText([]byte(fields[8]))
	s.O_SPEED_BAR.UnmarshalText([]byte(fields[9]))
	s.TOTAL_DIR.UnmarshalText([]byte(fields[10]))
	s.DIR_ME.UnmarshalText([]byte(fields[11]))
	s.DIR_MAE.UnmarshalText([]byte(fields[12]))
	s.DIR_MSE.UnmarshalText([]byte(fields[13]))
}

func (s *TCST_PROBRIRW_data) fill(fields []string) {
	s.ALAT.UnmarshalText([]byte(fields[0]))
	s.ALON.UnmarshalText([]byte(fields[1]))
	s.BLAT.UnmarshalText([]byte(fields[2]))
	s.BLON.UnmarshalText([]byte(fields[3]))
	s.INITIALS.UnmarshalText([]byte(fields[4]))
	s.TK_ERR.UnmarshalText([]byte(fields[5]))
	s.X_ERR.UnmarshalText([]byte(fields[6]))
	s.Y_ERR.UnmarshalText([]byte(fields[7]))
	s.ADLAND.UnmarshalText([]byte(fields[8]))
	s.BDLAND.UnmarshalText([]byte(fields[9]))
	s.RIRW_BEG.UnmarshalText([]byte(fields[10]))
	s.RIRW_END.UnmarshalText([]byte(fields[11]))
	s.RIRW_WINDOW.UnmarshalText([]byte(fields[12]))
	s.AWIND_END.UnmarshalText([]byte(fields[13]))
	s.BWIND_BEG.UnmarshalText([]byte(fields[14]))
	s.BWIND_END.UnmarshalText([]byte(fields[15]))
	s.BDELTA.UnmarshalText([]byte(fields[16]))
	s.BDELTA_MAX.UnmarshalText([]byte(fields[17]))
	s.BLEVEL_BEG.UnmarshalText([]byte(fields[18]))
	s.BLEVEL_END.UnmarshalText([]byte(fields[19]))
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
				value.UnmarshalText([]byte(fields[index]))
			}
			s.THRESH[key] = value
		}
	}
	s.INIT.UnmarshalText([]byte(fields[23]))
}

func (s *TCST_TCDIAG_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.INDEX.UnmarshalText([]byte(fields[1]))
	s.DIAG_SOURCE.UnmarshalText([]byte(fields[2]))
	s.TRACK_SOURCE.UnmarshalText([]byte(fields[3]))
	s.FIELD_SOURCE.UnmarshalText([]byte(fields[4]))
	// the first field of the repeating fields is the TOTAL, the second field is the 1st dimenSion of the 1st sequence (there might be only one sequence)
	var value validtypes.ValidString
	count, err := strconv.Atoi(fields[5])
	if err != nil {
		count = 0
	}
	keyPrefixes := []string{"DIAG_", "VALUE_"}
	s.DIAG = make(map[string]interface{})
	for group := 1; group <= count; group++ {
		for index := 6; index <= len(keyPrefixes); index++ {
			key := fmt.Sprintf("%s_%d", keyPrefixes[index-1], index)
			if index > len(fields) { // sometimes the data line is truncated - invalidate that field
				value.Reset()
			} else {
				value.UnmarshalText([]byte(fields[index]))
			}
			s.DIAG[key] = value
		}
	}
	s.INIT.UnmarshalText([]byte(fields[8]))
}

func (s *TCST_TCMPR_data) fill(fields []string) {
	s.TOTAL.UnmarshalText([]byte(fields[0]))
	s.INDEX.UnmarshalText([]byte(fields[1]))
	s.LEVEL.UnmarshalText([]byte(fields[2]))
	s.WATCH_WARN.UnmarshalText([]byte(fields[3]))
	s.INITIALS.UnmarshalText([]byte(fields[4]))
	s.ALAT.UnmarshalText([]byte(fields[5]))
	s.ALON.UnmarshalText([]byte(fields[6]))
	s.BLAT.UnmarshalText([]byte(fields[7]))
	s.BLON.UnmarshalText([]byte(fields[8]))
	s.TK_ERR.UnmarshalText([]byte(fields[9]))
	s.X_ERR.UnmarshalText([]byte(fields[10]))
	s.Y_ERR.UnmarshalText([]byte(fields[11]))
	s.ALTK_ERR.UnmarshalText([]byte(fields[12]))
	s.CRTK_ERR.UnmarshalText([]byte(fields[13]))
	s.ADLAND.UnmarshalText([]byte(fields[14]))
	s.BDLAND.UnmarshalText([]byte(fields[15]))
	s.AMSLP.UnmarshalText([]byte(fields[16]))
	s.BMSLP.UnmarshalText([]byte(fields[17]))
	s.AMAX_WIND.UnmarshalText([]byte(fields[18]))
	s.BMAX_WIND.UnmarshalText([]byte(fields[19]))
	s.AAL_WIND_34.UnmarshalText([]byte(fields[20]))
	s.BAL_WIND_34.UnmarshalText([]byte(fields[21]))
	s.ANE_WIND_34.UnmarshalText([]byte(fields[22]))
	s.BNE_WIND_34.UnmarshalText([]byte(fields[23]))
	s.ASE_WIND_34.UnmarshalText([]byte(fields[24]))
	s.BSE_WIND_34.UnmarshalText([]byte(fields[25]))
	s.ASW_WIND_34.UnmarshalText([]byte(fields[26]))
	s.BSW_WIND_34.UnmarshalText([]byte(fields[27]))
	s.ANW_WIND_34.UnmarshalText([]byte(fields[28]))
	s.BNW_WIND_34.UnmarshalText([]byte(fields[29]))
	s.AAL_WIND_50.UnmarshalText([]byte(fields[30]))
	s.BAL_WIND_50.UnmarshalText([]byte(fields[31]))
	s.ANE_WIND_50.UnmarshalText([]byte(fields[32]))
	s.BNE_WIND_50.UnmarshalText([]byte(fields[33]))
	s.ASE_WIND_50.UnmarshalText([]byte(fields[34]))
	s.BSE_WIND_50.UnmarshalText([]byte(fields[35]))
	s.ASW_WIND_50.UnmarshalText([]byte(fields[36]))
	s.BSW_WIND_50.UnmarshalText([]byte(fields[37]))
	s.ANW_WIND_50.UnmarshalText([]byte(fields[38]))
	s.BNW_WIND_50.UnmarshalText([]byte(fields[39]))
	s.AAL_WIND_64.UnmarshalText([]byte(fields[40]))
	s.BAL_WIND_64.UnmarshalText([]byte(fields[41]))
	s.ANE_WIND_64.UnmarshalText([]byte(fields[42]))
	s.BNE_WIND_64.UnmarshalText([]byte(fields[43]))
	s.ASE_WIND_64.UnmarshalText([]byte(fields[44]))
	s.BSE_WIND_64.UnmarshalText([]byte(fields[45]))
	s.ASW_WIND_64.UnmarshalText([]byte(fields[46]))
	s.BSW_WIND_64.UnmarshalText([]byte(fields[47]))
	s.ANW_WIND_64.UnmarshalText([]byte(fields[48]))
	s.BNW_WIND_64.UnmarshalText([]byte(fields[49]))
	s.ARADP.UnmarshalText([]byte(fields[50]))
	s.BRADP.UnmarshalText([]byte(fields[51]))
	s.ARRP.UnmarshalText([]byte(fields[52]))
	s.BRRP.UnmarshalText([]byte(fields[53]))
	s.AMRD.UnmarshalText([]byte(fields[54]))
	s.BMRD.UnmarshalText([]byte(fields[55]))
	s.AGUSTS.UnmarshalText([]byte(fields[56]))
	s.BGUSTS.UnmarshalText([]byte(fields[57]))
	s.AEYE.UnmarshalText([]byte(fields[58]))
	s.BEYE.UnmarshalText([]byte(fields[59]))
	s.ADIR.UnmarshalText([]byte(fields[60]))
	s.BDIR.UnmarshalText([]byte(fields[61]))
	s.ASPEED.UnmarshalText([]byte(fields[62]))
	s.BSPEED.UnmarshalText([]byte(fields[63]))
	s.ADEPTH.UnmarshalText([]byte(fields[64]))
	s.BDEPTH.UnmarshalText([]byte(fields[65]))
	s.NUM_MEMBERS.UnmarshalText([]byte(fields[66]))
	s.TRACK_SPREAD.UnmarshalText([]byte(fields[67]))
	s.TRACK_STDEV.UnmarshalText([]byte(fields[68]))
	s.MSLP_STDEV.UnmarshalText([]byte(fields[69]))
	s.MAX_WIND_STDEV.UnmarshalText([]byte(fields[70]))
	s.INIT.UnmarshalText([]byte(fields[71]))
}

// getDocForId function
// Creates a new doc, header functions and all.
func GetDocForId(fileLineType string, metaData util.VxMetadata, headerData []string, dataData []string, dataKey string) (map[string]interface{}, error) {
	var statDoc any
	switch fileLineType {
	case "STAT_CNT":
		elem_header := STAT_CNT_header{}
		elem_header.fill(headerData)
		elem_data := STAT_CNT_data{}
		elem_data.fill(dataData)

		tmp := STAT_CNT{
			VxMetadata:      metaData,
			STAT_CNT_header: elem_header,
			Data:            make(map[string]STAT_CNT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_CTC":
		elem_header := STAT_CTC_header{}
		elem_header.fill(headerData)
		elem_data := STAT_CTC_data{}
		elem_data.fill(dataData)

		tmp := STAT_CTC{
			VxMetadata:      metaData,
			STAT_CTC_header: elem_header,
			Data:            make(map[string]STAT_CTC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_CTS":
		elem_header := STAT_CTS_header{}
		elem_header.fill(headerData)
		elem_data := STAT_CTS_data{}
		elem_data.fill(dataData)

		tmp := STAT_CTS{
			VxMetadata:      metaData,
			STAT_CTS_header: elem_header,
			Data:            make(map[string]STAT_CTS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_FHO":
		elem_header := STAT_FHO_header{}
		elem_header.fill(headerData)
		elem_data := STAT_FHO_data{}
		elem_data.fill(dataData)

		tmp := STAT_FHO{
			VxMetadata:      metaData,
			STAT_FHO_header: elem_header,
			Data:            make(map[string]STAT_FHO_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_ISC":
		elem_header := STAT_ISC_header{}
		elem_header.fill(headerData)
		elem_data := STAT_ISC_data{}
		elem_data.fill(dataData)

		tmp := STAT_ISC{
			VxMetadata:      metaData,
			STAT_ISC_header: elem_header,
			Data:            make(map[string]STAT_ISC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_MCTC":
		elem_header := STAT_MCTC_header{}
		elem_header.fill(headerData)
		elem_data := STAT_MCTC_data{}
		elem_data.fill(dataData)

		tmp := STAT_MCTC{
			VxMetadata:       metaData,
			STAT_MCTC_header: elem_header,
			Data:             make(map[string]STAT_MCTC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_MCTS":
		elem_header := STAT_MCTS_header{}
		elem_header.fill(headerData)
		elem_data := STAT_MCTS_data{}
		elem_data.fill(dataData)

		tmp := STAT_MCTS{
			VxMetadata:       metaData,
			STAT_MCTS_header: elem_header,
			Data:             make(map[string]STAT_MCTS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_MPR":
		elem_header := STAT_MPR_header{}
		elem_header.fill(headerData)
		elem_data := STAT_MPR_data{}
		elem_data.fill(dataData)

		tmp := STAT_MPR{
			VxMetadata:      metaData,
			STAT_MPR_header: elem_header,
			Data:            make(map[string]STAT_MPR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_SEEPS":
		elem_header := STAT_SEEPS_header{}
		elem_header.fill(headerData)
		elem_data := STAT_SEEPS_data{}
		elem_data.fill(dataData)

		tmp := STAT_SEEPS{
			VxMetadata:        metaData,
			STAT_SEEPS_header: elem_header,
			Data:              make(map[string]STAT_SEEPS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_SEEPS_MPR":
		elem_header := STAT_SEEPS_MPR_header{}
		elem_header.fill(headerData)
		elem_data := STAT_SEEPS_MPR_data{}
		elem_data.fill(dataData)

		tmp := STAT_SEEPS_MPR{
			VxMetadata:            metaData,
			STAT_SEEPS_MPR_header: elem_header,
			Data:                  make(map[string]STAT_SEEPS_MPR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_NBRCNT":
		elem_header := STAT_NBRCNT_header{}
		elem_header.fill(headerData)
		elem_data := STAT_NBRCNT_data{}
		elem_data.fill(dataData)

		tmp := STAT_NBRCNT{
			VxMetadata:         metaData,
			STAT_NBRCNT_header: elem_header,
			Data:               make(map[string]STAT_NBRCNT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_NBRCTC":
		elem_header := STAT_NBRCTC_header{}
		elem_header.fill(headerData)
		elem_data := STAT_NBRCTC_data{}
		elem_data.fill(dataData)

		tmp := STAT_NBRCTC{
			VxMetadata:         metaData,
			STAT_NBRCTC_header: elem_header,
			Data:               make(map[string]STAT_NBRCTC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_NBRCTS":
		elem_header := STAT_NBRCTS_header{}
		elem_header.fill(headerData)
		elem_data := STAT_NBRCTS_data{}
		elem_data.fill(dataData)

		tmp := STAT_NBRCTS{
			VxMetadata:         metaData,
			STAT_NBRCTS_header: elem_header,
			Data:               make(map[string]STAT_NBRCTS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_GRAD":
		elem_header := STAT_GRAD_header{}
		elem_header.fill(headerData)
		elem_data := STAT_GRAD_data{}
		elem_data.fill(dataData)

		tmp := STAT_GRAD{
			VxMetadata:       metaData,
			STAT_GRAD_header: elem_header,
			Data:             make(map[string]STAT_GRAD_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_DMAP":
		elem_header := STAT_DMAP_header{}
		elem_header.fill(headerData)
		elem_data := STAT_DMAP_data{}
		elem_data.fill(dataData)

		tmp := STAT_DMAP{
			VxMetadata:       metaData,
			STAT_DMAP_header: elem_header,
			Data:             make(map[string]STAT_DMAP_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_ORANK":
		elem_header := STAT_ORANK_header{}
		elem_header.fill(headerData)
		elem_data := STAT_ORANK_data{}
		elem_data.fill(dataData)

		tmp := STAT_ORANK{
			VxMetadata:        metaData,
			STAT_ORANK_header: elem_header,
			Data:              make(map[string]STAT_ORANK_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_PCT":
		elem_header := STAT_PCT_header{}
		elem_header.fill(headerData)
		elem_data := STAT_PCT_data{}
		elem_data.fill(dataData)

		tmp := STAT_PCT{
			VxMetadata:      metaData,
			STAT_PCT_header: elem_header,
			Data:            make(map[string]STAT_PCT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_PJC":
		elem_header := STAT_PJC_header{}
		elem_header.fill(headerData)
		elem_data := STAT_PJC_data{}
		elem_data.fill(dataData)

		tmp := STAT_PJC{
			VxMetadata:      metaData,
			STAT_PJC_header: elem_header,
			Data:            make(map[string]STAT_PJC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_PRC":
		elem_header := STAT_PRC_header{}
		elem_header.fill(headerData)
		elem_data := STAT_PRC_data{}
		elem_data.fill(dataData)

		tmp := STAT_PRC{
			VxMetadata:      metaData,
			STAT_PRC_header: elem_header,
			Data:            make(map[string]STAT_PRC_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_PSTD":
		elem_header := STAT_PSTD_header{}
		elem_header.fill(headerData)
		elem_data := STAT_PSTD_data{}
		elem_data.fill(dataData)

		tmp := STAT_PSTD{
			VxMetadata:       metaData,
			STAT_PSTD_header: elem_header,
			Data:             make(map[string]STAT_PSTD_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_ECLV":
		elem_header := STAT_ECLV_header{}
		elem_header.fill(headerData)
		elem_data := STAT_ECLV_data{}
		elem_data.fill(dataData)

		tmp := STAT_ECLV{
			VxMetadata:       metaData,
			STAT_ECLV_header: elem_header,
			Data:             make(map[string]STAT_ECLV_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_ECNT":
		elem_header := STAT_ECNT_header{}
		elem_header.fill(headerData)
		elem_data := STAT_ECNT_data{}
		elem_data.fill(dataData)

		tmp := STAT_ECNT{
			VxMetadata:       metaData,
			STAT_ECNT_header: elem_header,
			Data:             make(map[string]STAT_ECNT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_RPS":
		elem_header := STAT_RPS_header{}
		elem_header.fill(headerData)
		elem_data := STAT_RPS_data{}
		elem_data.fill(dataData)

		tmp := STAT_RPS{
			VxMetadata:      metaData,
			STAT_RPS_header: elem_header,
			Data:            make(map[string]STAT_RPS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_RHIST":
		elem_header := STAT_RHIST_header{}
		elem_header.fill(headerData)
		elem_data := STAT_RHIST_data{}
		elem_data.fill(dataData)

		tmp := STAT_RHIST{
			VxMetadata:        metaData,
			STAT_RHIST_header: elem_header,
			Data:              make(map[string]STAT_RHIST_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_PHIST":
		elem_header := STAT_PHIST_header{}
		elem_header.fill(headerData)
		elem_data := STAT_PHIST_data{}
		elem_data.fill(dataData)

		tmp := STAT_PHIST{
			VxMetadata:        metaData,
			STAT_PHIST_header: elem_header,
			Data:              make(map[string]STAT_PHIST_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_RELP":
		elem_header := STAT_RELP_header{}
		elem_header.fill(headerData)
		elem_data := STAT_RELP_data{}
		elem_data.fill(dataData)

		tmp := STAT_RELP{
			VxMetadata:       metaData,
			STAT_RELP_header: elem_header,
			Data:             make(map[string]STAT_RELP_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_SAL1L2":
		elem_header := STAT_SAL1L2_header{}
		elem_header.fill(headerData)
		elem_data := STAT_SAL1L2_data{}
		elem_data.fill(dataData)

		tmp := STAT_SAL1L2{
			VxMetadata:         metaData,
			STAT_SAL1L2_header: elem_header,
			Data:               make(map[string]STAT_SAL1L2_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_SL1L2":
		elem_header := STAT_SL1L2_header{}
		elem_header.fill(headerData)
		elem_data := STAT_SL1L2_data{}
		elem_data.fill(dataData)

		tmp := STAT_SL1L2{
			VxMetadata:        metaData,
			STAT_SL1L2_header: elem_header,
			Data:              make(map[string]STAT_SL1L2_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_SSVAR":
		elem_header := STAT_SSVAR_header{}
		elem_header.fill(headerData)
		elem_data := STAT_SSVAR_data{}
		elem_data.fill(dataData)

		tmp := STAT_SSVAR{
			VxMetadata:        metaData,
			STAT_SSVAR_header: elem_header,
			Data:              make(map[string]STAT_SSVAR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_VAL1L2":
		elem_header := STAT_VAL1L2_header{}
		elem_header.fill(headerData)
		elem_data := STAT_VAL1L2_data{}
		elem_data.fill(dataData)

		tmp := STAT_VAL1L2{
			VxMetadata:         metaData,
			STAT_VAL1L2_header: elem_header,
			Data:               make(map[string]STAT_VAL1L2_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_VL1L2":
		elem_header := STAT_VL1L2_header{}
		elem_header.fill(headerData)
		elem_data := STAT_VL1L2_data{}
		elem_data.fill(dataData)

		tmp := STAT_VL1L2{
			VxMetadata:        metaData,
			STAT_VL1L2_header: elem_header,
			Data:              make(map[string]STAT_VL1L2_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_VCNT":
		elem_header := STAT_VCNT_header{}
		elem_header.fill(headerData)
		elem_data := STAT_VCNT_data{}
		elem_data.fill(dataData)

		tmp := STAT_VCNT{
			VxMetadata:       metaData,
			STAT_VCNT_header: elem_header,
			Data:             make(map[string]STAT_VCNT_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_GENMPR":
		elem_header := STAT_GENMPR_header{}
		elem_header.fill(headerData)
		elem_data := STAT_GENMPR_data{}
		elem_data.fill(dataData)

		tmp := STAT_GENMPR{
			VxMetadata:         metaData,
			STAT_GENMPR_header: elem_header,
			Data:               make(map[string]STAT_GENMPR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "STAT_SSIDX":
		elem_header := STAT_SSIDX_header{}
		elem_header.fill(headerData)
		elem_data := STAT_SSIDX_data{}
		elem_data.fill(dataData)

		tmp := STAT_SSIDX{
			VxMetadata:        metaData,
			STAT_SSIDX_header: elem_header,
			Data:              make(map[string]STAT_SSIDX_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "MODE_OBJ":
		elem_header := MODE_OBJ_header{}
		elem_header.fill(headerData)
		elem_data := MODE_OBJ_data{}
		elem_data.fill(dataData)

		tmp := MODE_OBJ{
			VxMetadata:      metaData,
			MODE_OBJ_header: elem_header,
			Data:            make(map[string]MODE_OBJ_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "MODE_CTS":
		elem_header := MODE_CTS_header{}
		elem_header.fill(headerData)
		elem_data := MODE_CTS_data{}
		elem_data.fill(dataData)

		tmp := MODE_CTS{
			VxMetadata:      metaData,
			MODE_CTS_header: elem_header,
			Data:            make(map[string]MODE_CTS_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "MTD_2DSINGLE":
		elem_header := MTD_2DSINGLE_header{}
		elem_header.fill(headerData)
		elem_data := MTD_2DSINGLE_data{}
		elem_data.fill(dataData)

		tmp := MTD_2DSINGLE{
			VxMetadata:          metaData,
			MTD_2DSINGLE_header: elem_header,
			Data:                make(map[string]MTD_2DSINGLE_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "MTD_3DSINGLE":
		elem_header := MTD_3DSINGLE_header{}
		elem_header.fill(headerData)
		elem_data := MTD_3DSINGLE_data{}
		elem_data.fill(dataData)

		tmp := MTD_3DSINGLE{
			VxMetadata:          metaData,
			MTD_3DSINGLE_header: elem_header,
			Data:                make(map[string]MTD_3DSINGLE_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "MTD_3DPAIR":
		elem_header := MTD_3DPAIR_header{}
		elem_header.fill(headerData)
		elem_data := MTD_3DPAIR_data{}
		elem_data.fill(dataData)

		tmp := MTD_3DPAIR{
			VxMetadata:        metaData,
			MTD_3DPAIR_header: elem_header,
			Data:              make(map[string]MTD_3DPAIR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "TCST_TCMPR":
		elem_header := TCST_TCMPR_header{}
		elem_header.fill(headerData)
		elem_data := TCST_TCMPR_data{}
		elem_data.fill(dataData)

		tmp := TCST_TCMPR{
			VxMetadata:        metaData,
			TCST_TCMPR_header: elem_header,
			Data:              make(map[string]TCST_TCMPR_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "TCST_TCDIAG":
		elem_header := TCST_TCDIAG_header{}
		elem_header.fill(headerData)
		elem_data := TCST_TCDIAG_data{}
		elem_data.fill(dataData)

		tmp := TCST_TCDIAG{
			VxMetadata:         metaData,
			TCST_TCDIAG_header: elem_header,
			Data:               make(map[string]TCST_TCDIAG_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	case "TCST_PROBRIRW":
		elem_header := TCST_PROBRIRW_header{}
		elem_header.fill(headerData)
		elem_data := TCST_PROBRIRW_data{}
		elem_data.fill(dataData)

		tmp := TCST_PROBRIRW{
			VxMetadata:           metaData,
			TCST_PROBRIRW_header: elem_header,
			Data:                 make(map[string]TCST_PROBRIRW_data),
		}
		tmp.Data[dataKey] = elem_data
		statDoc = tmp
	default:
		return nil, errors.New("GetDocForId: Unknown file_line type:" + fileLineType)
	}
	// Convert our types to a map[string]any by marshaling & unmarshaling through JSON
	// TODO - would it be advantageous to keep the type longer, e.g. for AddDataElement?
	jsonBytes, err := json.Marshal(statDoc)
	if err != nil {
		return nil, fmt.Errorf("error marshalling TCST_PROBRIRW struct: %w", err)
	}
	var doc map[string]any
	if err := json.Unmarshal(jsonBytes, &doc); err != nil {
		return nil, fmt.Errorf("error unmarshalling TCST_PROBRIRW to map: %w", err)
	}
	return doc, nil
}

// addDataElement function
// Header info has already been set by GetDocForId. Solely adds a new "data" element to the map.
// doc is expected to be a map representing the "base" struct (E.g. "STAT_CNT") with header, metadata, & data info
func AddDataElement(dataKey string, fileLineType string, dataData []string, doc *map[string]interface{}) (map[string]interface{}, error) {
	switch fileLineType {
	case "STAT_CNT":
		elem_data := STAT_CNT_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_CNT_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_CTC":
		elem_data := STAT_CTC_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_CTC_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_CTS":
		elem_data := STAT_CTS_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_CTS_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_FHO":
		elem_data := STAT_FHO_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_FHO_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_ISC":
		elem_data := STAT_ISC_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_ISC_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_MCTC":
		elem_data := STAT_MCTC_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_MCTC_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_MCTS":
		elem_data := STAT_MCTS_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_MCTS_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_MPR":
		elem_data := STAT_MPR_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_MPR_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_SEEPS":
		elem_data := STAT_SEEPS_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SEEPS_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_SEEPS_MPR":
		elem_data := STAT_SEEPS_MPR_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SEEPS_MPR_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_NBRCNT":
		elem_data := STAT_NBRCNT_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_NBRCNT_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_NBRCTC":
		elem_data := STAT_NBRCTC_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_NBRCTC_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_NBRCTS":
		elem_data := STAT_NBRCTS_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_NBRCTS_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_GRAD":
		elem_data := STAT_GRAD_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_GRAD_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_DMAP":
		elem_data := STAT_DMAP_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_DMAP_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_ORANK":
		elem_data := STAT_ORANK_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_ORANK_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_PCT":
		elem_data := STAT_PCT_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PCT_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_PJC":
		elem_data := STAT_PJC_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PJC_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_PRC":
		elem_data := STAT_PRC_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PRC_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_PSTD":
		elem_data := STAT_PSTD_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PSTD_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_ECLV":
		elem_data := STAT_ECLV_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_ECLV_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_ECNT":
		elem_data := STAT_ECNT_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_ECNT_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_RPS":
		elem_data := STAT_RPS_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_RPS_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_RHIST":
		elem_data := STAT_RHIST_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_RHIST_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_PHIST":
		elem_data := STAT_PHIST_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_PHIST_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_RELP":
		elem_data := STAT_RELP_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_RELP_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_SAL1L2":
		elem_data := STAT_SAL1L2_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SAL1L2_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_SL1L2":
		elem_data := STAT_SL1L2_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SL1L2_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_SSVAR":
		elem_data := STAT_SSVAR_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SSVAR_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_VAL1L2":
		elem_data := STAT_VAL1L2_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_VAL1L2_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_VL1L2":
		elem_data := STAT_VL1L2_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_VL1L2_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_VCNT":
		elem_data := STAT_VCNT_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_VCNT_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_GENMPR":
		elem_data := STAT_GENMPR_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_GENMPR_data); ok {
			val[dataKey] = elem_data
		}
	case "STAT_SSIDX":
		elem_data := STAT_SSIDX_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]STAT_SSIDX_data); ok {
			val[dataKey] = elem_data
		}
	case "MODE_OBJ":
		elem_data := MODE_OBJ_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]MODE_OBJ_data); ok {
			val[dataKey] = elem_data
		}
	case "MODE_CTS":
		elem_data := MODE_CTS_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]MODE_CTS_data); ok {
			val[dataKey] = elem_data
		}
	case "MTD_2DSINGLE":
		elem_data := MTD_2DSINGLE_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]MTD_2DSINGLE_data); ok {
			val[dataKey] = elem_data
		}
	case "MTD_3DSINGLE":
		elem_data := MTD_3DSINGLE_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]MTD_3DSINGLE_data); ok {
			val[dataKey] = elem_data
		}
	case "MTD_3DPAIR":
		elem_data := MTD_3DPAIR_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]MTD_3DPAIR_data); ok {
			val[dataKey] = elem_data
		}
	case "TCST_TCMPR":
		elem_data := TCST_TCMPR_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]TCST_TCMPR_data); ok {
			val[dataKey] = elem_data
		}
	case "TCST_TCDIAG":
		elem_data := TCST_TCDIAG_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]TCST_TCDIAG_data); ok {
			val[dataKey] = elem_data
		}
	case "TCST_PROBRIRW":
		elem_data := TCST_PROBRIRW_data{}
		elem_data.fill(dataData)
		if val, ok := (*doc)["data"].(map[string]TCST_PROBRIRW_data); ok {
			val[dataKey] = elem_data
		}
	default:
		return nil, errors.New("AddDataElement: Unknown file_line type:" + fileLineType)
	}
	return *doc, nil
}

var MetHeaderColumnsFileUrl = "https://raw.githubusercontent.com/dtcenter/MET/refs/heads/main_v12.0/data/table_files/met_header_columns_V12.0.txt"
