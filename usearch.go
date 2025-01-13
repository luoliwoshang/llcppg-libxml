package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
/**
* Data structure for searching
* @stable ICU 2.4
*/
type UStringSearch struct {
	Unused [8]uint8
}
type USearchAttribute c.Int

const (
	USearchAttributeUSEARCHOVERLAP           USearchAttribute = 0
	USearchAttributeUSEARCHCANONICALMATCH    USearchAttribute = 1
	USearchAttributeUSEARCHELEMENTCOMPARISON USearchAttribute = 2
	USearchAttributeUSEARCHATTRIBUTECOUNT    USearchAttribute = 3
)

type USearchAttributeValue c.Int

const (
	USearchAttributeValueUSEARCHDEFAULT                     USearchAttributeValue = -1
	USearchAttributeValueUSEARCHOFF                         USearchAttributeValue = 0
	USearchAttributeValueUSEARCHON                          USearchAttributeValue = 1
	USearchAttributeValueUSEARCHSTANDARDELEMENTCOMPARISON   USearchAttributeValue = 2
	USearchAttributeValueUSEARCHPATTERNBASEWEIGHTISWILDCARD USearchAttributeValue = 3
	USearchAttributeValueUSEARCHANYBASEWEIGHTISWILDCARD     USearchAttributeValue = 4
	USearchAttributeValueUSEARCHATTRIBUTEVALUECOUNT         USearchAttributeValue = 5
)
