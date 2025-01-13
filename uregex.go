package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type URegularExpression struct {
	Unused [8]uint8
}
type URegexpFlag c.Int

const (
	URegexpFlagUREGEXCANONEQ               URegexpFlag = 128
	URegexpFlagUREGEXCASEINSENSITIVE       URegexpFlag = 2
	URegexpFlagUREGEXCOMMENTS              URegexpFlag = 4
	URegexpFlagUREGEXDOTALL                URegexpFlag = 32
	URegexpFlagUREGEXLITERAL               URegexpFlag = 16
	URegexpFlagUREGEXMULTILINE             URegexpFlag = 8
	URegexpFlagUREGEXUNIXLINES             URegexpFlag = 1
	URegexpFlagUREGEXUWORD                 URegexpFlag = 256
	URegexpFlagUREGEXERRORONUNKNOWNESCAPES URegexpFlag = 512
)
// llgo:type C
type URegexMatchCallback func(unsafe.Pointer, int32) UBool
// llgo:type C
type URegexFindProgressCallback func(unsafe.Pointer, int64) UBool
