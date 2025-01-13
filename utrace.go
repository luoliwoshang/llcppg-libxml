package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type UTraceLevel c.Int

const (
	UTraceLevelUTRACEOFF       UTraceLevel = -1
	UTraceLevelUTRACEERROR     UTraceLevel = 0
	UTraceLevelUTRACEWARNING   UTraceLevel = 3
	UTraceLevelUTRACEOPENCLOSE UTraceLevel = 5
	UTraceLevelUTRACEINFO      UTraceLevel = 7
	UTraceLevelUTRACEVERBOSE   UTraceLevel = 9
)

type UTraceFunctionNumber c.Int

const (
	UTraceFunctionNumberUTRACEFUNCTIONSTART           UTraceFunctionNumber = 0
	UTraceFunctionNumberUTRACEUINIT                   UTraceFunctionNumber = 0
	UTraceFunctionNumberUTRACEUCLEANUP                UTraceFunctionNumber = 1
	UTraceFunctionNumberUTRACEFUNCTIONLIMIT           UTraceFunctionNumber = 2
	UTraceFunctionNumberUTRACECONVERSIONSTART         UTraceFunctionNumber = 4096
	UTraceFunctionNumberUTRACEUCNVOPEN                UTraceFunctionNumber = 4096
	UTraceFunctionNumberUTRACEUCNVOPENPACKAGE         UTraceFunctionNumber = 4097
	UTraceFunctionNumberUTRACEUCNVOPENALGORITHMIC     UTraceFunctionNumber = 4098
	UTraceFunctionNumberUTRACEUCNVCLONE               UTraceFunctionNumber = 4099
	UTraceFunctionNumberUTRACEUCNVCLOSE               UTraceFunctionNumber = 4100
	UTraceFunctionNumberUTRACEUCNVFLUSHCACHE          UTraceFunctionNumber = 4101
	UTraceFunctionNumberUTRACEUCNVLOAD                UTraceFunctionNumber = 4102
	UTraceFunctionNumberUTRACEUCNVUNLOAD              UTraceFunctionNumber = 4103
	UTraceFunctionNumberUTRACECONVERSIONLIMIT         UTraceFunctionNumber = 4104
	UTraceFunctionNumberUTRACECOLLATIONSTART          UTraceFunctionNumber = 8192
	UTraceFunctionNumberUTRACEUCOLOPEN                UTraceFunctionNumber = 8192
	UTraceFunctionNumberUTRACEUCOLCLOSE               UTraceFunctionNumber = 8193
	UTraceFunctionNumberUTRACEUCOLSTRCOLL             UTraceFunctionNumber = 8194
	UTraceFunctionNumberUTRACEUCOLGETSORTKEY          UTraceFunctionNumber = 8195
	UTraceFunctionNumberUTRACEUCOLGETLOCALE           UTraceFunctionNumber = 8196
	UTraceFunctionNumberUTRACEUCOLNEXTSORTKEYPART     UTraceFunctionNumber = 8197
	UTraceFunctionNumberUTRACEUCOLSTRCOLLITER         UTraceFunctionNumber = 8198
	UTraceFunctionNumberUTRACEUCOLOPENFROMSHORTSTRING UTraceFunctionNumber = 8199
	UTraceFunctionNumberUTRACEUCOLSTRCOLLUTF8         UTraceFunctionNumber = 8200
	UTraceFunctionNumberUTRACECOLLATIONLIMIT          UTraceFunctionNumber = 8201
	UTraceFunctionNumberUTRACEUDATASTART              UTraceFunctionNumber = 12288
	UTraceFunctionNumberUTRACEUDATARESOURCE           UTraceFunctionNumber = 12288
	UTraceFunctionNumberUTRACEUDATABUNDLE             UTraceFunctionNumber = 12289
	UTraceFunctionNumberUTRACEUDATADATAFILE           UTraceFunctionNumber = 12290
	UTraceFunctionNumberUTRACEUDATARESFILE            UTraceFunctionNumber = 12291
	UTraceFunctionNumberUTRACERESDATALIMIT            UTraceFunctionNumber = 12292
	UTraceFunctionNumberUTRACEUBRKSTART               UTraceFunctionNumber = 16384
	UTraceFunctionNumberUTRACEUBRKCREATECHARACTER     UTraceFunctionNumber = 16384
	UTraceFunctionNumberUTRACEUBRKCREATEWORD          UTraceFunctionNumber = 16385
	UTraceFunctionNumberUTRACEUBRKCREATELINE          UTraceFunctionNumber = 16386
	UTraceFunctionNumberUTRACEUBRKCREATESENTENCE      UTraceFunctionNumber = 16387
	UTraceFunctionNumberUTRACEUBRKCREATETITLE         UTraceFunctionNumber = 16388
	UTraceFunctionNumberUTRACEUBRKCREATEBREAKENGINE   UTraceFunctionNumber = 16389
	UTraceFunctionNumberUTRACEUBRKLIMIT               UTraceFunctionNumber = 16390
)
// llgo:type C
type UTraceEntry func(unsafe.Pointer, int32)
// llgo:type C
type UTraceExit func(unsafe.Pointer, int32, *int8, c.Int)
// llgo:type C
type UTraceData func(unsafe.Pointer, int32, int32, *int8, c.Int)
