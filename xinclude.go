package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_XmlXIncludeCtxt struct {
	Unused [8]uint8
}
type XmlXIncludeCtxt X_XmlXIncludeCtxt
type XmlXIncludeCtxtPtr *XmlXIncludeCtxt
//go:linkname XmlXIncludeProcess C.xmlXIncludeProcess
func XmlXIncludeProcess(doc XmlDocPtr) c.Int
//go:linkname XmlXIncludeProcessFlags C.xmlXIncludeProcessFlags
func XmlXIncludeProcessFlags(doc XmlDocPtr, flags c.Int) c.Int
//go:linkname XmlXIncludeProcessFlagsData C.xmlXIncludeProcessFlagsData
func XmlXIncludeProcessFlagsData(doc XmlDocPtr, flags c.Int, data unsafe.Pointer) c.Int
//go:linkname XmlXIncludeProcessTreeFlagsData C.xmlXIncludeProcessTreeFlagsData
func XmlXIncludeProcessTreeFlagsData(tree XmlNodePtr, flags c.Int, data unsafe.Pointer) c.Int
//go:linkname XmlXIncludeProcessTree C.xmlXIncludeProcessTree
func XmlXIncludeProcessTree(tree XmlNodePtr) c.Int
//go:linkname XmlXIncludeProcessTreeFlags C.xmlXIncludeProcessTreeFlags
func XmlXIncludeProcessTreeFlags(tree XmlNodePtr, flags c.Int) c.Int
//go:linkname XmlXIncludeNewContext C.xmlXIncludeNewContext
func XmlXIncludeNewContext(doc XmlDocPtr) XmlXIncludeCtxtPtr
//go:linkname XmlXIncludeSetFlags C.xmlXIncludeSetFlags
func XmlXIncludeSetFlags(ctxt XmlXIncludeCtxtPtr, flags c.Int) c.Int
//go:linkname XmlXIncludeSetErrorHandler C.xmlXIncludeSetErrorHandler
func XmlXIncludeSetErrorHandler(ctxt XmlXIncludeCtxtPtr, handler XmlStructuredErrorFunc, data unsafe.Pointer)
//go:linkname XmlXIncludeGetLastError C.xmlXIncludeGetLastError
func XmlXIncludeGetLastError(ctxt XmlXIncludeCtxtPtr) c.Int
//go:linkname XmlXIncludeFreeContext C.xmlXIncludeFreeContext
func XmlXIncludeFreeContext(ctxt XmlXIncludeCtxtPtr)
//go:linkname XmlXIncludeProcessNode C.xmlXIncludeProcessNode
func XmlXIncludeProcessNode(ctxt XmlXIncludeCtxtPtr, tree XmlNodePtr) c.Int
