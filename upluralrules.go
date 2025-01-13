package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UFormattedNumber struct {
	Unused [8]uint8
}

type UFormattedNumberRange struct {
	Unused [8]uint8
}
type UPluralType c.Int

const (
	UPluralTypeUPLURALTYPECARDINAL UPluralType = 0
	UPluralTypeUPLURALTYPEORDINAL  UPluralType = 1
	UPluralTypeUPLURALTYPECOUNT    UPluralType = 2
)
/**
 * Opaque UPluralRules object for use in C programs.
 * @stable ICU 4.8
 */

type UPluralRules struct {
	Unused [8]uint8
}
