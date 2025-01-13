package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UDisplayOptionsGrammaticalCase c.Int

const (
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEUNDEFINED          UDisplayOptionsGrammaticalCase = 0
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEABLATIVE           UDisplayOptionsGrammaticalCase = 1
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEACCUSATIVE         UDisplayOptionsGrammaticalCase = 2
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASECOMITATIVE         UDisplayOptionsGrammaticalCase = 3
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEDATIVE             UDisplayOptionsGrammaticalCase = 4
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEERGATIVE           UDisplayOptionsGrammaticalCase = 5
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEGENITIVE           UDisplayOptionsGrammaticalCase = 6
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEINSTRUMENTAL       UDisplayOptionsGrammaticalCase = 7
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASELOCATIVE           UDisplayOptionsGrammaticalCase = 8
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASELOCATIVECOPULATIVE UDisplayOptionsGrammaticalCase = 9
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASENOMINATIVE         UDisplayOptionsGrammaticalCase = 10
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEOBLIQUE            UDisplayOptionsGrammaticalCase = 11
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEPREPOSITIONAL      UDisplayOptionsGrammaticalCase = 12
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASESOCIATIVE          UDisplayOptionsGrammaticalCase = 13
	UDisplayOptionsGrammaticalCaseUDISPOPTGRAMMATICALCASEVOCATIVE           UDisplayOptionsGrammaticalCase = 14
)

type UDisplayOptionsPluralCategory c.Int

const (
	UDisplayOptionsPluralCategoryUDISPOPTPLURALCATEGORYUNDEFINED UDisplayOptionsPluralCategory = 0
	UDisplayOptionsPluralCategoryUDISPOPTPLURALCATEGORYZERO      UDisplayOptionsPluralCategory = 1
	UDisplayOptionsPluralCategoryUDISPOPTPLURALCATEGORYONE       UDisplayOptionsPluralCategory = 2
	UDisplayOptionsPluralCategoryUDISPOPTPLURALCATEGORYTWO       UDisplayOptionsPluralCategory = 3
	UDisplayOptionsPluralCategoryUDISPOPTPLURALCATEGORYFEW       UDisplayOptionsPluralCategory = 4
	UDisplayOptionsPluralCategoryUDISPOPTPLURALCATEGORYMANY      UDisplayOptionsPluralCategory = 5
	UDisplayOptionsPluralCategoryUDISPOPTPLURALCATEGORYOTHER     UDisplayOptionsPluralCategory = 6
)

type UDisplayOptionsNounClass c.Int

const (
	UDisplayOptionsNounClassUDISPOPTNOUNCLASSUNDEFINED UDisplayOptionsNounClass = 0
	UDisplayOptionsNounClassUDISPOPTNOUNCLASSOTHER     UDisplayOptionsNounClass = 1
	UDisplayOptionsNounClassUDISPOPTNOUNCLASSNEUTER    UDisplayOptionsNounClass = 2
	UDisplayOptionsNounClassUDISPOPTNOUNCLASSFEMININE  UDisplayOptionsNounClass = 3
	UDisplayOptionsNounClassUDISPOPTNOUNCLASSMASCULINE UDisplayOptionsNounClass = 4
	UDisplayOptionsNounClassUDISPOPTNOUNCLASSANIMATE   UDisplayOptionsNounClass = 5
	UDisplayOptionsNounClassUDISPOPTNOUNCLASSINANIMATE UDisplayOptionsNounClass = 6
	UDisplayOptionsNounClassUDISPOPTNOUNCLASSPERSONAL  UDisplayOptionsNounClass = 7
	UDisplayOptionsNounClassUDISPOPTNOUNCLASSCOMMON    UDisplayOptionsNounClass = 8
)

type UDisplayOptionsCapitalization c.Int

const (
	UDisplayOptionsCapitalizationUDISPOPTCAPITALIZATIONUNDEFINED           UDisplayOptionsCapitalization = 0
	UDisplayOptionsCapitalizationUDISPOPTCAPITALIZATIONBEGINNINGOFSENTENCE UDisplayOptionsCapitalization = 1
	UDisplayOptionsCapitalizationUDISPOPTCAPITALIZATIONMIDDLEOFSENTENCE    UDisplayOptionsCapitalization = 2
	UDisplayOptionsCapitalizationUDISPOPTCAPITALIZATIONSTANDALONE          UDisplayOptionsCapitalization = 3
	UDisplayOptionsCapitalizationUDISPOPTCAPITALIZATIONUILISTORMENU        UDisplayOptionsCapitalization = 4
)

type UDisplayOptionsNameStyle c.Int

const (
	UDisplayOptionsNameStyleUDISPOPTNAMESTYLEUNDEFINED     UDisplayOptionsNameStyle = 0
	UDisplayOptionsNameStyleUDISPOPTNAMESTYLESTANDARDNAMES UDisplayOptionsNameStyle = 1
	UDisplayOptionsNameStyleUDISPOPTNAMESTYLEDIALECTNAMES  UDisplayOptionsNameStyle = 2
)

type UDisplayOptionsDisplayLength c.Int

const (
	UDisplayOptionsDisplayLengthUDISPOPTDISPLAYLENGTHUNDEFINED UDisplayOptionsDisplayLength = 0
	UDisplayOptionsDisplayLengthUDISPOPTDISPLAYLENGTHFULL      UDisplayOptionsDisplayLength = 1
	UDisplayOptionsDisplayLengthUDISPOPTDISPLAYLENGTHSHORT     UDisplayOptionsDisplayLength = 2
)

type UDisplayOptionsSubstituteHandling c.Int

const (
	UDisplayOptionsSubstituteHandlingUDISPOPTSUBSTITUTEHANDLINGUNDEFINED    UDisplayOptionsSubstituteHandling = 0
	UDisplayOptionsSubstituteHandlingUDISPOPTSUBSTITUTEHANDLINGSUBSTITUTE   UDisplayOptionsSubstituteHandling = 1
	UDisplayOptionsSubstituteHandlingUDISPOPTSUBSTITUTEHANDLINGNOSUBSTITUTE UDisplayOptionsSubstituteHandling = 2
)
