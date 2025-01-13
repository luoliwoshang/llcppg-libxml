package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XmlSaveOption c.Int

const (
	XmlSaveOptionXMLSAVEFORMAT   XmlSaveOption = 1
	XmlSaveOptionXMLSAVENODECL   XmlSaveOption = 2
	XmlSaveOptionXMLSAVENOEMPTY  XmlSaveOption = 4
	XmlSaveOptionXMLSAVENOXHTML  XmlSaveOption = 8
	XmlSaveOptionXMLSAVEXHTML    XmlSaveOption = 16
	XmlSaveOptionXMLSAVEASXML    XmlSaveOption = 32
	XmlSaveOptionXMLSAVEASHTML   XmlSaveOption = 64
	XmlSaveOptionXMLSAVEWSNONSIG XmlSaveOption = 128
)

type X_XmlSaveCtxt struct {
	Unused [8]uint8
}
type XmlSaveCtxt X_XmlSaveCtxt
type XmlSaveCtxtPtr *XmlSaveCtxt
//go:linkname XmlSaveToFd C.xmlSaveToFd
func XmlSaveToFd(fd c.Int, encoding *int8, options c.Int) XmlSaveCtxtPtr
//go:linkname XmlSaveToFilename C.xmlSaveToFilename
func XmlSaveToFilename(filename *int8, encoding *int8, options c.Int) XmlSaveCtxtPtr
//go:linkname XmlSaveToBuffer C.xmlSaveToBuffer
func XmlSaveToBuffer(buffer XmlBufferPtr, encoding *int8, options c.Int) XmlSaveCtxtPtr
//go:linkname XmlSaveToIO C.xmlSaveToIO
func XmlSaveToIO(iowrite XmlOutputWriteCallback, ioclose XmlOutputCloseCallback, ioctx unsafe.Pointer, encoding *int8, options c.Int) XmlSaveCtxtPtr
//go:linkname XmlSaveDoc C.xmlSaveDoc
func XmlSaveDoc(ctxt XmlSaveCtxtPtr, doc XmlDocPtr) c.Long
//go:linkname XmlSaveTree C.xmlSaveTree
func XmlSaveTree(ctxt XmlSaveCtxtPtr, node XmlNodePtr) c.Long
//go:linkname XmlSaveFlush C.xmlSaveFlush
func XmlSaveFlush(ctxt XmlSaveCtxtPtr) c.Int
//go:linkname XmlSaveClose C.xmlSaveClose
func XmlSaveClose(ctxt XmlSaveCtxtPtr) c.Int
//go:linkname XmlSaveFinish C.xmlSaveFinish
func XmlSaveFinish(ctxt XmlSaveCtxtPtr) c.Int
//go:linkname XmlSaveSetEscape C.xmlSaveSetEscape
func XmlSaveSetEscape(ctxt XmlSaveCtxtPtr, escape XmlCharEncodingOutputFunc) c.Int
//go:linkname XmlSaveSetAttrEscape C.xmlSaveSetAttrEscape
func XmlSaveSetAttrEscape(ctxt XmlSaveCtxtPtr, escape XmlCharEncodingOutputFunc) c.Int
//go:linkname XmlThrDefIndentTreeOutput C.xmlThrDefIndentTreeOutput
func XmlThrDefIndentTreeOutput(v c.Int) c.Int
//go:linkname XmlThrDefTreeIndentString C.xmlThrDefTreeIndentString
func XmlThrDefTreeIndentString(v *int8) *int8
//go:linkname XmlThrDefSaveNoEmptyTags C.xmlThrDefSaveNoEmptyTags
func XmlThrDefSaveNoEmptyTags(v c.Int) c.Int
