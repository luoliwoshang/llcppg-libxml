package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const U8LEAD3T1BITS string = " 000000000000\x1000"
const U8LEAD4T1BITS string = "\x00\x00\x00\x00\x00\x00\x00\x00\x1e\x0f\x0f\x0f\x00\x00\x00\x00"
const U8MAXLENGTH c.Int = 4
