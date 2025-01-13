package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
/**
 * UResourceBundle is an opaque type for handles for resource bundles in C APIs.
 * @stable ICU 2.0
 */

type UResourceBundle struct {
	Unused [8]uint8
}
type UResType c.Int

const (
	UResTypeURESNONE      UResType = -1
	UResTypeURESSTRING    UResType = 0
	UResTypeURESBINARY    UResType = 1
	UResTypeURESTABLE     UResType = 2
	UResTypeURESALIAS     UResType = 3
	UResTypeURESINT       UResType = 7
	UResTypeURESARRAY     UResType = 8
	UResTypeURESINTVECTOR UResType = 14
	UResTypeRESNONE       UResType = -1
	UResTypeRESSTRING     UResType = 0
	UResTypeRESBINARY     UResType = 1
	UResTypeRESTABLE      UResType = 2
	UResTypeRESALIAS      UResType = 3
	UResTypeRESINT        UResType = 7
	UResTypeRESARRAY      UResType = 8
	UResTypeRESINTVECTOR  UResType = 14
	UResTypeRESRESERVED   UResType = 15
	UResTypeURESLIMIT     UResType = 16
)
