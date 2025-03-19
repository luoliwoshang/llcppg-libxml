package libxml

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_xmlRegexp struct {
	Unused [8]uint8
}
type Regexp X_xmlRegexp
type RegexpPtr *Regexp

type X_xmlRegExecCtxt struct {
	Unused [8]uint8
}
type RegExecCtxt X_xmlRegExecCtxt
type RegExecCtxtPtr *RegExecCtxt

/*
 * The POSIX like API
 */
// llgo:link (*Char).RegexpCompile C.xmlRegexpCompile
func (recv_ *Char) RegexpCompile() RegexpPtr {
	return nil
}

//go:linkname RegFreeRegexp C.xmlRegFreeRegexp
func RegFreeRegexp(regexp RegexpPtr)

//go:linkname RegexpExec C.xmlRegexpExec
func RegexpExec(comp RegexpPtr, value *Char) c.Int

//go:linkname RegexpPrint C.xmlRegexpPrint
func RegexpPrint(output *c.FILE, regexp RegexpPtr)

//go:linkname RegexpIsDeterminist C.xmlRegexpIsDeterminist
func RegexpIsDeterminist(comp RegexpPtr) c.Int

// llgo:type C
type RegExecCallbacks func(RegExecCtxtPtr, *Char, unsafe.Pointer, unsafe.Pointer)

/*
 * The progressive API
 */
//go:linkname RegNewExecCtxt C.xmlRegNewExecCtxt
func RegNewExecCtxt(comp RegexpPtr, callback RegExecCallbacks, data unsafe.Pointer) RegExecCtxtPtr

//go:linkname RegFreeExecCtxt C.xmlRegFreeExecCtxt
func RegFreeExecCtxt(exec RegExecCtxtPtr)

//go:linkname RegExecPushString C.xmlRegExecPushString
func RegExecPushString(exec RegExecCtxtPtr, value *Char, data unsafe.Pointer) c.Int

//go:linkname RegExecPushString2 C.xmlRegExecPushString2
func RegExecPushString2(exec RegExecCtxtPtr, value *Char, value2 *Char, data unsafe.Pointer) c.Int

//go:linkname RegExecNextValues C.xmlRegExecNextValues
func RegExecNextValues(exec RegExecCtxtPtr, nbval *c.Int, nbneg *c.Int, values **Char, terminal *c.Int) c.Int

//go:linkname RegExecErrInfo C.xmlRegExecErrInfo
func RegExecErrInfo(exec RegExecCtxtPtr, string **Char, nbval *c.Int, nbneg *c.Int, values **Char, terminal *c.Int) c.Int
