package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
//go:linkname XmlDebugDumpString C.xmlDebugDumpString
func XmlDebugDumpString(output *c.FILE, str *XmlChar)
//go:linkname XmlDebugDumpAttr C.xmlDebugDumpAttr
func XmlDebugDumpAttr(output *c.FILE, attr XmlAttrPtr, depth c.Int)
//go:linkname XmlDebugDumpAttrList C.xmlDebugDumpAttrList
func XmlDebugDumpAttrList(output *c.FILE, attr XmlAttrPtr, depth c.Int)
//go:linkname XmlDebugDumpOneNode C.xmlDebugDumpOneNode
func XmlDebugDumpOneNode(output *c.FILE, node XmlNodePtr, depth c.Int)
//go:linkname XmlDebugDumpNode C.xmlDebugDumpNode
func XmlDebugDumpNode(output *c.FILE, node XmlNodePtr, depth c.Int)
//go:linkname XmlDebugDumpNodeList C.xmlDebugDumpNodeList
func XmlDebugDumpNodeList(output *c.FILE, node XmlNodePtr, depth c.Int)
//go:linkname XmlDebugDumpDocumentHead C.xmlDebugDumpDocumentHead
func XmlDebugDumpDocumentHead(output *c.FILE, doc XmlDocPtr)
//go:linkname XmlDebugDumpDocument C.xmlDebugDumpDocument
func XmlDebugDumpDocument(output *c.FILE, doc XmlDocPtr)
//go:linkname XmlDebugDumpDTD C.xmlDebugDumpDTD
func XmlDebugDumpDTD(output *c.FILE, dtd XmlDtdPtr)
//go:linkname XmlDebugDumpEntities C.xmlDebugDumpEntities
func XmlDebugDumpEntities(output *c.FILE, doc XmlDocPtr)
/****************************************************************
 *								*
 *			Checking routines			*
 *								*
 ****************************************************************/
//go:linkname XmlDebugCheckDocument C.xmlDebugCheckDocument
func XmlDebugCheckDocument(output *c.FILE, doc XmlDocPtr) c.Int
/****************************************************************
 *								*
 *			XML shell helpers			*
 *								*
 ****************************************************************/
//go:linkname XmlLsOneNode C.xmlLsOneNode
func XmlLsOneNode(output *c.FILE, node XmlNodePtr)
//go:linkname XmlLsCountNode C.xmlLsCountNode
func XmlLsCountNode(node XmlNodePtr) c.Int
//go:linkname XmlBoolToText C.xmlBoolToText
func XmlBoolToText(boolval c.Int) *int8
// llgo:type C
type XmlShellReadlineFunc func(*int8) *int8

type X_XmlShellCtxt struct {
	Filename *int8
	Doc      XmlDocPtr
	Node     XmlNodePtr
	Pctxt    XmlXPathContextPtr
	Loaded   c.Int
	Output   *c.FILE
	Input    unsafe.Pointer
}
type XmlShellCtxt X_XmlShellCtxt
type XmlShellCtxtPtr *XmlShellCtxt
// llgo:type C
type XmlShellCmd func(XmlShellCtxtPtr, *int8, XmlNodePtr, XmlNodePtr) c.Int
//go:linkname XmlShellPrintXPathError C.xmlShellPrintXPathError
func XmlShellPrintXPathError(errorType c.Int, arg *int8)
//go:linkname XmlShellPrintXPathResult C.xmlShellPrintXPathResult
func XmlShellPrintXPathResult(list XmlXPathObjectPtr)
//go:linkname XmlShellList C.xmlShellList
func XmlShellList(ctxt XmlShellCtxtPtr, arg *int8, node XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShellBase C.xmlShellBase
func XmlShellBase(ctxt XmlShellCtxtPtr, arg *int8, node XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShellDir C.xmlShellDir
func XmlShellDir(ctxt XmlShellCtxtPtr, arg *int8, node XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShellLoad C.xmlShellLoad
func XmlShellLoad(ctxt XmlShellCtxtPtr, filename *int8, node XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShellPrintNode C.xmlShellPrintNode
func XmlShellPrintNode(node XmlNodePtr)
//go:linkname XmlShellCat C.xmlShellCat
func XmlShellCat(ctxt XmlShellCtxtPtr, arg *int8, node XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShellWrite C.xmlShellWrite
func XmlShellWrite(ctxt XmlShellCtxtPtr, filename *int8, node XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShellSave C.xmlShellSave
func XmlShellSave(ctxt XmlShellCtxtPtr, filename *int8, node XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShellValidate C.xmlShellValidate
func XmlShellValidate(ctxt XmlShellCtxtPtr, dtd *int8, node XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShellDu C.xmlShellDu
func XmlShellDu(ctxt XmlShellCtxtPtr, arg *int8, tree XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShellPwd C.xmlShellPwd
func XmlShellPwd(ctxt XmlShellCtxtPtr, buffer *int8, node XmlNodePtr, node2 XmlNodePtr) c.Int
//go:linkname XmlShell C.xmlShell
func XmlShell(doc XmlDocPtr, filename *int8, input XmlShellReadlineFunc, output *c.FILE)
