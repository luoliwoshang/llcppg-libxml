package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
//go:linkname XmlCheckVersion C.xmlCheckVersion
func XmlCheckVersion(version c.Int)
