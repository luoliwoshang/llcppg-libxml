package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type UFormattableType c.Int

const (
	UFormattableTypeUFMTDATE   UFormattableType = 0
	UFormattableTypeUFMTDOUBLE UFormattableType = 1
	UFormattableTypeUFMTLONG   UFormattableType = 2
	UFormattableTypeUFMTSTRING UFormattableType = 3
	UFormattableTypeUFMTARRAY  UFormattableType = 4
	UFormattableTypeUFMTINT64  UFormattableType = 5
	UFormattableTypeUFMTOBJECT UFormattableType = 6
	UFormattableTypeUFMTCOUNT  UFormattableType = 7
)

type UFormattable unsafe.Pointer
