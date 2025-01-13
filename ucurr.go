package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type UCurrencyUsage c.Int

const (
	UCurrencyUsageUCURRUSAGESTANDARD UCurrencyUsage = 0
	UCurrencyUsageUCURRUSAGECASH     UCurrencyUsage = 1
	UCurrencyUsageUCURRUSAGECOUNT    UCurrencyUsage = 2
)

type UCurrNameStyle c.Int

const (
	UCurrNameStyleUCURRSYMBOLNAME        UCurrNameStyle = 0
	UCurrNameStyleUCURRLONGNAME          UCurrNameStyle = 1
	UCurrNameStyleUCURRNARROWSYMBOLNAME  UCurrNameStyle = 2
	UCurrNameStyleUCURRFORMALSYMBOLNAME  UCurrNameStyle = 3
	UCurrNameStyleUCURRVARIANTSYMBOLNAME UCurrNameStyle = 4
)

type UCurrRegistryKey unsafe.Pointer
type UCurrCurrencyType c.Int

const (
	UCurrCurrencyTypeUCURRALL           UCurrCurrencyType = 2147483647
	UCurrCurrencyTypeUCURRCOMMON        UCurrCurrencyType = 1
	UCurrCurrencyTypeUCURRUNCOMMON      UCurrCurrencyType = 2
	UCurrCurrencyTypeUCURRDEPRECATED    UCurrCurrencyType = 4
	UCurrCurrencyTypeUCURRNONDEPRECATED UCurrCurrencyType = 8
)
