package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const UCOLSAFECLONEBUFFERSIZE c.Int = 1
/** A collator.
*  For usage in C programs.
*/
type UCollator struct {
	Unused [8]uint8
}
type UCollationResult c.Int

const (
	UCollationResultUCOLEQUAL   UCollationResult = 0
	UCollationResultUCOLGREATER UCollationResult = 1
	UCollationResultUCOLLESS    UCollationResult = -1
)

type UColAttributeValue c.Int

const (
	UColAttributeValueUCOLDEFAULT             UColAttributeValue = -1
	UColAttributeValueUCOLPRIMARY             UColAttributeValue = 0
	UColAttributeValueUCOLSECONDARY           UColAttributeValue = 1
	UColAttributeValueUCOLTERTIARY            UColAttributeValue = 2
	UColAttributeValueUCOLDEFAULTSTRENGTH     UColAttributeValue = 2
	UColAttributeValueUCOLCESTRENGTHLIMIT     UColAttributeValue = 3
	UColAttributeValueUCOLQUATERNARY          UColAttributeValue = 3
	UColAttributeValueUCOLIDENTICAL           UColAttributeValue = 15
	UColAttributeValueUCOLSTRENGTHLIMIT       UColAttributeValue = 16
	UColAttributeValueUCOLOFF                 UColAttributeValue = 16
	UColAttributeValueUCOLON                  UColAttributeValue = 17
	UColAttributeValueUCOLSHIFTED             UColAttributeValue = 20
	UColAttributeValueUCOLNONIGNORABLE        UColAttributeValue = 21
	UColAttributeValueUCOLLOWERFIRST          UColAttributeValue = 24
	UColAttributeValueUCOLUPPERFIRST          UColAttributeValue = 25
	UColAttributeValueUCOLATTRIBUTEVALUECOUNT UColAttributeValue = 26
)

type UColReorderCode c.Int

const (
	UColReorderCodeUCOLREORDERCODEDEFAULT     UColReorderCode = -1
	UColReorderCodeUCOLREORDERCODENONE        UColReorderCode = 103
	UColReorderCodeUCOLREORDERCODEOTHERS      UColReorderCode = 103
	UColReorderCodeUCOLREORDERCODESPACE       UColReorderCode = 4096
	UColReorderCodeUCOLREORDERCODEFIRST       UColReorderCode = 4096
	UColReorderCodeUCOLREORDERCODEPUNCTUATION UColReorderCode = 4097
	UColReorderCodeUCOLREORDERCODESYMBOL      UColReorderCode = 4098
	UColReorderCodeUCOLREORDERCODECURRENCY    UColReorderCode = 4099
	UColReorderCodeUCOLREORDERCODEDIGIT       UColReorderCode = 4100
	UColReorderCodeUCOLREORDERCODELIMIT       UColReorderCode = 4101
)

type UCollationStrength UColAttributeValue
type UColAttribute c.Int

const (
	UColAttributeUCOLFRENCHCOLLATION        UColAttribute = 0
	UColAttributeUCOLALTERNATEHANDLING      UColAttribute = 1
	UColAttributeUCOLCASEFIRST              UColAttribute = 2
	UColAttributeUCOLCASELEVEL              UColAttribute = 3
	UColAttributeUCOLNORMALIZATIONMODE      UColAttribute = 4
	UColAttributeUCOLDECOMPOSITIONMODE      UColAttribute = 4
	UColAttributeUCOLSTRENGTH               UColAttribute = 5
	UColAttributeUCOLHIRAGANAQUATERNARYMODE UColAttribute = 6
	UColAttributeUCOLNUMERICCOLLATION       UColAttribute = 7
	UColAttributeUCOLATTRIBUTECOUNT         UColAttribute = 8
)

type UColRuleOption c.Int

const (
	UColRuleOptionUCOLTAILORINGONLY UColRuleOption = 0
	UColRuleOptionUCOLFULLRULES     UColRuleOption = 1
)

type UColBoundMode c.Int

const (
	UColBoundModeUCOLBOUNDLOWER      UColBoundMode = 0
	UColBoundModeUCOLBOUNDUPPER      UColBoundMode = 1
	UColBoundModeUCOLBOUNDUPPERLONG  UColBoundMode = 2
	UColBoundModeUCOLBOUNDVALUECOUNT UColBoundMode = 3
)
