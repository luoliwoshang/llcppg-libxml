package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const UDEFINEFALSEANDTRUE c.Int = 0
const USIZEOFUCHAR c.Int = 2
const UCHAR16ISTYPEDEF c.Int = 0

type UBool int8
type UChar Char16T
type OldUChar uint16
type UChar32 int32
