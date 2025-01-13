package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
/**
 * Opaque UListFormatter object for use in C
 * @stable ICU 55
 */

type UListFormatter struct {
	Unused [8]uint8
}

type UFormattedList struct {
	Unused [8]uint8
}
type UListFormatterField c.Int

const (
	UListFormatterFieldULISTFMTLITERALFIELD UListFormatterField = 0
	UListFormatterFieldULISTFMTELEMENTFIELD UListFormatterField = 1
)

type UListFormatterType c.Int

const (
	UListFormatterTypeULISTFMTTYPEAND   UListFormatterType = 0
	UListFormatterTypeULISTFMTTYPEOR    UListFormatterType = 1
	UListFormatterTypeULISTFMTTYPEUNITS UListFormatterType = 2
)

type UListFormatterWidth c.Int

const (
	UListFormatterWidthULISTFMTWIDTHWIDE   UListFormatterWidth = 0
	UListFormatterWidthULISTFMTWIDTHSHORT  UListFormatterWidth = 1
	UListFormatterWidthULISTFMTWIDTHNARROW UListFormatterWidth = 2
)
