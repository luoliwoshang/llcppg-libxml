package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type HtmlParserCtxt XmlParserCtxt
type HtmlParserCtxtPtr XmlParserCtxtPtr
type HtmlParserNodeInfo XmlParserNodeInfo
type HtmlSAXHandler XmlSAXHandler
type HtmlSAXHandlerPtr XmlSAXHandlerPtr
type HtmlParserInput XmlParserInput
type HtmlParserInputPtr XmlParserInputPtr
type HtmlDocPtr XmlDocPtr
type HtmlNodePtr XmlNodePtr

type X_HtmlElemDesc struct {
	Name          *int8
	StartTag      int8
	EndTag        int8
	SaveEndTag    int8
	Empty         int8
	Depr          int8
	Dtd           int8
	Isinline      int8
	Desc          *int8
	Subelts       **int8
	Defaultsubelt *int8
	AttrsOpt      **int8
	AttrsDepr     **int8
	AttrsReq      **int8
}
type HtmlElemDesc X_HtmlElemDesc
type HtmlElemDescPtr *HtmlElemDesc

type X_HtmlEntityDesc struct {
	Value c.Uint
	Name  *int8
	Desc  *int8
}
type HtmlEntityDesc X_HtmlEntityDesc
type HtmlEntityDescPtr *HtmlEntityDesc
//go:linkname X__HtmlDefaultSAXHandler C.__htmlDefaultSAXHandler
func X__HtmlDefaultSAXHandler() *XmlSAXHandlerV1
//go:linkname HtmlInitAutoClose C.htmlInitAutoClose
func HtmlInitAutoClose()
// llgo:link (*XmlChar).HtmlTagLookup C.htmlTagLookup
func (recv_ *XmlChar) HtmlTagLookup() *HtmlElemDesc {
	return nil
}
// llgo:link (*XmlChar).HtmlEntityLookup C.htmlEntityLookup
func (recv_ *XmlChar) HtmlEntityLookup() *HtmlEntityDesc {
	return nil
}
//go:linkname HtmlEntityValueLookup C.htmlEntityValueLookup
func HtmlEntityValueLookup(value c.Uint) *HtmlEntityDesc
//go:linkname HtmlIsAutoClosed C.htmlIsAutoClosed
func HtmlIsAutoClosed(doc HtmlDocPtr, elem HtmlNodePtr) c.Int
//go:linkname HtmlAutoCloseTag C.htmlAutoCloseTag
func HtmlAutoCloseTag(doc HtmlDocPtr, name *XmlChar, elem HtmlNodePtr) c.Int
//go:linkname HtmlParseEntityRef C.htmlParseEntityRef
func HtmlParseEntityRef(ctxt HtmlParserCtxtPtr, str **XmlChar) *HtmlEntityDesc
//go:linkname HtmlParseCharRef C.htmlParseCharRef
func HtmlParseCharRef(ctxt HtmlParserCtxtPtr) c.Int
//go:linkname HtmlParseElement C.htmlParseElement
func HtmlParseElement(ctxt HtmlParserCtxtPtr)
//go:linkname HtmlNewParserCtxt C.htmlNewParserCtxt
func HtmlNewParserCtxt() HtmlParserCtxtPtr
// llgo:link (*HtmlSAXHandler).HtmlNewSAXParserCtxt C.htmlNewSAXParserCtxt
func (recv_ *HtmlSAXHandler) HtmlNewSAXParserCtxt(userData unsafe.Pointer) HtmlParserCtxtPtr {
	return nil
}
//go:linkname HtmlCreateMemoryParserCtxt C.htmlCreateMemoryParserCtxt
func HtmlCreateMemoryParserCtxt(buffer *int8, size c.Int) HtmlParserCtxtPtr
//go:linkname HtmlParseDocument C.htmlParseDocument
func HtmlParseDocument(ctxt HtmlParserCtxtPtr) c.Int
// llgo:link (*XmlChar).HtmlSAXParseDoc C.htmlSAXParseDoc
func (recv_ *XmlChar) HtmlSAXParseDoc(encoding *int8, sax HtmlSAXHandlerPtr, userData unsafe.Pointer) HtmlDocPtr {
	return nil
}
// llgo:link (*XmlChar).HtmlParseDoc C.htmlParseDoc
func (recv_ *XmlChar) HtmlParseDoc(encoding *int8) HtmlDocPtr {
	return nil
}
//go:linkname HtmlCreateFileParserCtxt C.htmlCreateFileParserCtxt
func HtmlCreateFileParserCtxt(filename *int8, encoding *int8) HtmlParserCtxtPtr
//go:linkname HtmlSAXParseFile C.htmlSAXParseFile
func HtmlSAXParseFile(filename *int8, encoding *int8, sax HtmlSAXHandlerPtr, userData unsafe.Pointer) HtmlDocPtr
//go:linkname HtmlParseFile C.htmlParseFile
func HtmlParseFile(filename *int8, encoding *int8) HtmlDocPtr
//go:linkname UTF8ToHtml C.UTF8ToHtml
func UTF8ToHtml(out *int8, outlen *c.Int, in *int8, inlen *c.Int) c.Int
//go:linkname HtmlEncodeEntities C.htmlEncodeEntities
func HtmlEncodeEntities(out *int8, outlen *c.Int, in *int8, inlen *c.Int, quoteChar c.Int) c.Int
// llgo:link (*XmlChar).HtmlIsScriptAttribute C.htmlIsScriptAttribute
func (recv_ *XmlChar) HtmlIsScriptAttribute() c.Int {
	return 0
}
//go:linkname HtmlHandleOmittedElem C.htmlHandleOmittedElem
func HtmlHandleOmittedElem(val c.Int) c.Int
/**
 * Interfaces for the Push mode.
 */
//go:linkname HtmlCreatePushParserCtxt C.htmlCreatePushParserCtxt
func HtmlCreatePushParserCtxt(sax HtmlSAXHandlerPtr, user_data unsafe.Pointer, chunk *int8, size c.Int, filename *int8, enc XmlCharEncoding) HtmlParserCtxtPtr
//go:linkname HtmlParseChunk C.htmlParseChunk
func HtmlParseChunk(ctxt HtmlParserCtxtPtr, chunk *int8, size c.Int, terminate c.Int) c.Int
//go:linkname HtmlFreeParserCtxt C.htmlFreeParserCtxt
func HtmlFreeParserCtxt(ctxt HtmlParserCtxtPtr)

type HtmlParserOption c.Int

const (
	HtmlParserOptionHTMLPARSERECOVER   HtmlParserOption = 1
	HtmlParserOptionHTMLPARSENODEFDTD  HtmlParserOption = 4
	HtmlParserOptionHTMLPARSENOERROR   HtmlParserOption = 32
	HtmlParserOptionHTMLPARSENOWARNING HtmlParserOption = 64
	HtmlParserOptionHTMLPARSEPEDANTIC  HtmlParserOption = 128
	HtmlParserOptionHTMLPARSENOBLANKS  HtmlParserOption = 256
	HtmlParserOptionHTMLPARSENONET     HtmlParserOption = 2048
	HtmlParserOptionHTMLPARSENOIMPLIED HtmlParserOption = 8192
	HtmlParserOptionHTMLPARSECOMPACT   HtmlParserOption = 65536
	HtmlParserOptionHTMLPARSEIGNOREENC HtmlParserOption = 2097152
)
//go:linkname HtmlCtxtReset C.htmlCtxtReset
func HtmlCtxtReset(ctxt HtmlParserCtxtPtr)
//go:linkname HtmlCtxtUseOptions C.htmlCtxtUseOptions
func HtmlCtxtUseOptions(ctxt HtmlParserCtxtPtr, options c.Int) c.Int
// llgo:link (*XmlChar).HtmlReadDoc C.htmlReadDoc
func (recv_ *XmlChar) HtmlReadDoc(URL *int8, encoding *int8, options c.Int) HtmlDocPtr {
	return nil
}
//go:linkname HtmlReadFile C.htmlReadFile
func HtmlReadFile(URL *int8, encoding *int8, options c.Int) HtmlDocPtr
//go:linkname HtmlReadMemory C.htmlReadMemory
func HtmlReadMemory(buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) HtmlDocPtr
//go:linkname HtmlReadFd C.htmlReadFd
func HtmlReadFd(fd c.Int, URL *int8, encoding *int8, options c.Int) HtmlDocPtr
//go:linkname HtmlReadIO C.htmlReadIO
func HtmlReadIO(ioread XmlInputReadCallback, ioclose XmlInputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) HtmlDocPtr
//go:linkname HtmlCtxtParseDocument C.htmlCtxtParseDocument
func HtmlCtxtParseDocument(ctxt HtmlParserCtxtPtr, input XmlParserInputPtr) HtmlDocPtr
//go:linkname HtmlCtxtReadDoc C.htmlCtxtReadDoc
func HtmlCtxtReadDoc(ctxt XmlParserCtxtPtr, cur *XmlChar, URL *int8, encoding *int8, options c.Int) HtmlDocPtr
//go:linkname HtmlCtxtReadFile C.htmlCtxtReadFile
func HtmlCtxtReadFile(ctxt XmlParserCtxtPtr, filename *int8, encoding *int8, options c.Int) HtmlDocPtr
//go:linkname HtmlCtxtReadMemory C.htmlCtxtReadMemory
func HtmlCtxtReadMemory(ctxt XmlParserCtxtPtr, buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) HtmlDocPtr
//go:linkname HtmlCtxtReadFd C.htmlCtxtReadFd
func HtmlCtxtReadFd(ctxt XmlParserCtxtPtr, fd c.Int, URL *int8, encoding *int8, options c.Int) HtmlDocPtr
//go:linkname HtmlCtxtReadIO C.htmlCtxtReadIO
func HtmlCtxtReadIO(ctxt XmlParserCtxtPtr, ioread XmlInputReadCallback, ioclose XmlInputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) HtmlDocPtr

type HtmlStatus c.Int

const (
	HtmlStatusHTMLNA         HtmlStatus = 0
	HtmlStatusHTMLINVALID    HtmlStatus = 1
	HtmlStatusHTMLDEPRECATED HtmlStatus = 2
	HtmlStatusHTMLVALID      HtmlStatus = 4
	HtmlStatusHTMLREQUIRED   HtmlStatus = 12
)
// llgo:link (*HtmlElemDesc).HtmlAttrAllowed C.htmlAttrAllowed
func (recv_ *HtmlElemDesc) HtmlAttrAllowed(*XmlChar, c.Int) HtmlStatus {
	return 0
}
// llgo:link (*HtmlElemDesc).HtmlElementAllowedHere C.htmlElementAllowedHere
func (recv_ *HtmlElemDesc) HtmlElementAllowedHere(*XmlChar) c.Int {
	return 0
}
// llgo:link (*HtmlElemDesc).HtmlElementStatusHere C.htmlElementStatusHere
func (recv_ *HtmlElemDesc) HtmlElementStatusHere(*HtmlElemDesc) HtmlStatus {
	return 0
}
//go:linkname HtmlNodeStatus C.htmlNodeStatus
func HtmlNodeStatus(HtmlNodePtr, c.Int) HtmlStatus
