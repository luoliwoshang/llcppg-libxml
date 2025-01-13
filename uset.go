package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const (
	USETIGNORESPACE           c.Int = 1
	USETCASEINSENSITIVE       c.Int = 2
	USETADDCASEMAPPINGS       c.Int = 4
	USETSIMPLECASEINSENSITIVE c.Int = 6
)

type USetSpanCondition c.Int

const (
	USetSpanConditionUSETSPANNOTCONTAINED   USetSpanCondition = 0
	USetSpanConditionUSETSPANCONTAINED      USetSpanCondition = 1
	USetSpanConditionUSETSPANSIMPLE         USetSpanCondition = 2
	USetSpanConditionUSETSPANCONDITIONCOUNT USetSpanCondition = 3
)
const USETSERIALIZEDSTATICARRAYCAPACITY c.Int = 8
/**
 * A serialized form of a Unicode set.  Limited manipulations are
 * possible directly on a serialized set.  See below.
 * @stable ICU 2.4
 */

type USerializedSet struct {
	Array       *uint16
	BmpLength   int32
	Length      int32
	StaticArray [8]uint16
}
