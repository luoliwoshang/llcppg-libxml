package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_XmlRegexp struct {
	Unused [8]uint8
}
type XmlRegexp X_XmlRegexp
type XmlRegexpPtr *XmlRegexp

type X_XmlRegExecCtxt struct {
	Unused [8]uint8
}
type XmlRegExecCtxt X_XmlRegExecCtxt
type XmlRegExecCtxtPtr *XmlRegExecCtxt
// llgo:link (*XmlChar).XmlRegexpCompile C.xmlRegexpCompile
func (recv_ *XmlChar) XmlRegexpCompile() XmlRegexpPtr {
	return nil
}
//go:linkname XmlRegFreeRegexp C.xmlRegFreeRegexp
func XmlRegFreeRegexp(regexp XmlRegexpPtr)
//go:linkname XmlRegexpExec C.xmlRegexpExec
func XmlRegexpExec(comp XmlRegexpPtr, value *XmlChar) c.Int
//go:linkname XmlRegexpPrint C.xmlRegexpPrint
func XmlRegexpPrint(output *c.FILE, regexp XmlRegexpPtr)
//go:linkname XmlRegexpIsDeterminist C.xmlRegexpIsDeterminist
func XmlRegexpIsDeterminist(comp XmlRegexpPtr) c.Int
// llgo:type C
type XmlRegExecCallbacks func(XmlRegExecCtxtPtr, *XmlChar, unsafe.Pointer, unsafe.Pointer)
//go:linkname XmlRegNewExecCtxt C.xmlRegNewExecCtxt
func XmlRegNewExecCtxt(comp XmlRegexpPtr, callback XmlRegExecCallbacks, data unsafe.Pointer) XmlRegExecCtxtPtr
//go:linkname XmlRegFreeExecCtxt C.xmlRegFreeExecCtxt
func XmlRegFreeExecCtxt(exec XmlRegExecCtxtPtr)
//go:linkname XmlRegExecPushString C.xmlRegExecPushString
func XmlRegExecPushString(exec XmlRegExecCtxtPtr, value *XmlChar, data unsafe.Pointer) c.Int
//go:linkname XmlRegExecPushString2 C.xmlRegExecPushString2
func XmlRegExecPushString2(exec XmlRegExecCtxtPtr, value *XmlChar, value2 *XmlChar, data unsafe.Pointer) c.Int
//go:linkname XmlRegExecNextValues C.xmlRegExecNextValues
func XmlRegExecNextValues(exec XmlRegExecCtxtPtr, nbval *c.Int, nbneg *c.Int, values **XmlChar, terminal *c.Int) c.Int
//go:linkname XmlRegExecErrInfo C.xmlRegExecErrInfo
func XmlRegExecErrInfo(exec XmlRegExecCtxtPtr, string **XmlChar, nbval *c.Int, nbneg *c.Int, values **XmlChar, terminal *c.Int) c.Int
