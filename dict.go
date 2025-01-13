package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type XmlDict X_XmlDict
type XmlDictPtr *XmlDict
//go:linkname XmlInitializeDict C.xmlInitializeDict
func XmlInitializeDict() c.Int
//go:linkname XmlDictCreate C.xmlDictCreate
func XmlDictCreate() XmlDictPtr
//go:linkname XmlDictSetLimit C.xmlDictSetLimit
func XmlDictSetLimit(dict XmlDictPtr, limit uintptr) uintptr
//go:linkname XmlDictGetUsage C.xmlDictGetUsage
func XmlDictGetUsage(dict XmlDictPtr) uintptr
//go:linkname XmlDictCreateSub C.xmlDictCreateSub
func XmlDictCreateSub(sub XmlDictPtr) XmlDictPtr
//go:linkname XmlDictReference C.xmlDictReference
func XmlDictReference(dict XmlDictPtr) c.Int
//go:linkname XmlDictFree C.xmlDictFree
func XmlDictFree(dict XmlDictPtr)
//go:linkname XmlDictLookup C.xmlDictLookup
func XmlDictLookup(dict XmlDictPtr, name *XmlChar, len c.Int) *XmlChar
//go:linkname XmlDictExists C.xmlDictExists
func XmlDictExists(dict XmlDictPtr, name *XmlChar, len c.Int) *XmlChar
//go:linkname XmlDictQLookup C.xmlDictQLookup
func XmlDictQLookup(dict XmlDictPtr, prefix *XmlChar, name *XmlChar) *XmlChar
//go:linkname XmlDictOwns C.xmlDictOwns
func XmlDictOwns(dict XmlDictPtr, str *XmlChar) c.Int
//go:linkname XmlDictSize C.xmlDictSize
func XmlDictSize(dict XmlDictPtr) c.Int
//go:linkname XmlDictCleanup C.xmlDictCleanup
func XmlDictCleanup()
