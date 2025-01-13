package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
//go:linkname XmlSAX2GetPublicId C.xmlSAX2GetPublicId
func XmlSAX2GetPublicId(ctx unsafe.Pointer) *XmlChar
//go:linkname XmlSAX2GetSystemId C.xmlSAX2GetSystemId
func XmlSAX2GetSystemId(ctx unsafe.Pointer) *XmlChar
//go:linkname XmlSAX2SetDocumentLocator C.xmlSAX2SetDocumentLocator
func XmlSAX2SetDocumentLocator(ctx unsafe.Pointer, loc XmlSAXLocatorPtr)
//go:linkname XmlSAX2GetLineNumber C.xmlSAX2GetLineNumber
func XmlSAX2GetLineNumber(ctx unsafe.Pointer) c.Int
//go:linkname XmlSAX2GetColumnNumber C.xmlSAX2GetColumnNumber
func XmlSAX2GetColumnNumber(ctx unsafe.Pointer) c.Int
//go:linkname XmlSAX2IsStandalone C.xmlSAX2IsStandalone
func XmlSAX2IsStandalone(ctx unsafe.Pointer) c.Int
//go:linkname XmlSAX2HasInternalSubset C.xmlSAX2HasInternalSubset
func XmlSAX2HasInternalSubset(ctx unsafe.Pointer) c.Int
//go:linkname XmlSAX2HasExternalSubset C.xmlSAX2HasExternalSubset
func XmlSAX2HasExternalSubset(ctx unsafe.Pointer) c.Int
//go:linkname XmlSAX2InternalSubset C.xmlSAX2InternalSubset
func XmlSAX2InternalSubset(ctx unsafe.Pointer, name *XmlChar, ExternalID *XmlChar, SystemID *XmlChar)
//go:linkname XmlSAX2ExternalSubset C.xmlSAX2ExternalSubset
func XmlSAX2ExternalSubset(ctx unsafe.Pointer, name *XmlChar, ExternalID *XmlChar, SystemID *XmlChar)
//go:linkname XmlSAX2GetEntity C.xmlSAX2GetEntity
func XmlSAX2GetEntity(ctx unsafe.Pointer, name *XmlChar) XmlEntityPtr
//go:linkname XmlSAX2GetParameterEntity C.xmlSAX2GetParameterEntity
func XmlSAX2GetParameterEntity(ctx unsafe.Pointer, name *XmlChar) XmlEntityPtr
//go:linkname XmlSAX2ResolveEntity C.xmlSAX2ResolveEntity
func XmlSAX2ResolveEntity(ctx unsafe.Pointer, publicId *XmlChar, systemId *XmlChar) XmlParserInputPtr
//go:linkname XmlSAX2EntityDecl C.xmlSAX2EntityDecl
func XmlSAX2EntityDecl(ctx unsafe.Pointer, name *XmlChar, type_ c.Int, publicId *XmlChar, systemId *XmlChar, content *XmlChar)
//go:linkname XmlSAX2AttributeDecl C.xmlSAX2AttributeDecl
func XmlSAX2AttributeDecl(ctx unsafe.Pointer, elem *XmlChar, fullname *XmlChar, type_ c.Int, def c.Int, defaultValue *XmlChar, tree XmlEnumerationPtr)
//go:linkname XmlSAX2ElementDecl C.xmlSAX2ElementDecl
func XmlSAX2ElementDecl(ctx unsafe.Pointer, name *XmlChar, type_ c.Int, content XmlElementContentPtr)
//go:linkname XmlSAX2NotationDecl C.xmlSAX2NotationDecl
func XmlSAX2NotationDecl(ctx unsafe.Pointer, name *XmlChar, publicId *XmlChar, systemId *XmlChar)
//go:linkname XmlSAX2UnparsedEntityDecl C.xmlSAX2UnparsedEntityDecl
func XmlSAX2UnparsedEntityDecl(ctx unsafe.Pointer, name *XmlChar, publicId *XmlChar, systemId *XmlChar, notationName *XmlChar)
//go:linkname XmlSAX2StartDocument C.xmlSAX2StartDocument
func XmlSAX2StartDocument(ctx unsafe.Pointer)
//go:linkname XmlSAX2EndDocument C.xmlSAX2EndDocument
func XmlSAX2EndDocument(ctx unsafe.Pointer)
//go:linkname XmlSAX2StartElement C.xmlSAX2StartElement
func XmlSAX2StartElement(ctx unsafe.Pointer, fullname *XmlChar, atts **XmlChar)
//go:linkname XmlSAX2EndElement C.xmlSAX2EndElement
func XmlSAX2EndElement(ctx unsafe.Pointer, name *XmlChar)
//go:linkname XmlSAX2StartElementNs C.xmlSAX2StartElementNs
func XmlSAX2StartElementNs(ctx unsafe.Pointer, localname *XmlChar, prefix *XmlChar, URI *XmlChar, nb_namespaces c.Int, namespaces **XmlChar, nb_attributes c.Int, nb_defaulted c.Int, attributes **XmlChar)
//go:linkname XmlSAX2EndElementNs C.xmlSAX2EndElementNs
func XmlSAX2EndElementNs(ctx unsafe.Pointer, localname *XmlChar, prefix *XmlChar, URI *XmlChar)
//go:linkname XmlSAX2Reference C.xmlSAX2Reference
func XmlSAX2Reference(ctx unsafe.Pointer, name *XmlChar)
//go:linkname XmlSAX2Characters C.xmlSAX2Characters
func XmlSAX2Characters(ctx unsafe.Pointer, ch *XmlChar, len c.Int)
//go:linkname XmlSAX2IgnorableWhitespace C.xmlSAX2IgnorableWhitespace
func XmlSAX2IgnorableWhitespace(ctx unsafe.Pointer, ch *XmlChar, len c.Int)
//go:linkname XmlSAX2ProcessingInstruction C.xmlSAX2ProcessingInstruction
func XmlSAX2ProcessingInstruction(ctx unsafe.Pointer, target *XmlChar, data *XmlChar)
//go:linkname XmlSAX2Comment C.xmlSAX2Comment
func XmlSAX2Comment(ctx unsafe.Pointer, value *XmlChar)
//go:linkname XmlSAX2CDataBlock C.xmlSAX2CDataBlock
func XmlSAX2CDataBlock(ctx unsafe.Pointer, value *XmlChar, len c.Int)
//go:linkname XmlSAXDefaultVersion C.xmlSAXDefaultVersion
func XmlSAXDefaultVersion(version c.Int) c.Int
// llgo:link (*XmlSAXHandler).XmlSAXVersion C.xmlSAXVersion
func (recv_ *XmlSAXHandler) XmlSAXVersion(version c.Int) c.Int {
	return 0
}
// llgo:link (*XmlSAXHandler).XmlSAX2InitDefaultSAXHandler C.xmlSAX2InitDefaultSAXHandler
func (recv_ *XmlSAXHandler) XmlSAX2InitDefaultSAXHandler(warning c.Int) {
}
// llgo:link (*XmlSAXHandler).XmlSAX2InitHtmlDefaultSAXHandler C.xmlSAX2InitHtmlDefaultSAXHandler
func (recv_ *XmlSAXHandler) XmlSAX2InitHtmlDefaultSAXHandler() {
}
//go:linkname HtmlDefaultSAXHandlerInit C.htmlDefaultSAXHandlerInit
func HtmlDefaultSAXHandlerInit()
//go:linkname XmlDefaultSAXHandlerInit C.xmlDefaultSAXHandlerInit
func XmlDefaultSAXHandlerInit()
