package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const UCOPYRIGHTSTRINGLENGTH c.Int = 128
const UMAXVERSIONLENGTH c.Int = 4
const UVERSIONDELIMITER string = "."
const UMAXVERSIONSTRINGLENGTH c.Int = 20

type UVersionInfo [4]uint8
