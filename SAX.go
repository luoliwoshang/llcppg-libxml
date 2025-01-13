package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
//go:linkname GetPublicId C.getPublicId
func GetPublicId(ctx unsafe.Pointer) *XmlChar
//go:linkname GetSystemId C.getSystemId
func GetSystemId(ctx unsafe.Pointer) *XmlChar
//go:linkname SetDocumentLocator C.setDocumentLocator
func SetDocumentLocator(ctx unsafe.Pointer, loc XmlSAXLocatorPtr)
//go:linkname GetLineNumber C.getLineNumber
func GetLineNumber(ctx unsafe.Pointer) c.Int
//go:linkname GetColumnNumber C.getColumnNumber
func GetColumnNumber(ctx unsafe.Pointer) c.Int
//go:linkname IsStandalone C.isStandalone
func IsStandalone(ctx unsafe.Pointer) c.Int
//go:linkname HasInternalSubset C.hasInternalSubset
func HasInternalSubset(ctx unsafe.Pointer) c.Int
//go:linkname HasExternalSubset C.hasExternalSubset
func HasExternalSubset(ctx unsafe.Pointer) c.Int
//go:linkname InternalSubset C.internalSubset
func InternalSubset(ctx unsafe.Pointer, name *XmlChar, ExternalID *XmlChar, SystemID *XmlChar)
//go:linkname ExternalSubset C.externalSubset
func ExternalSubset(ctx unsafe.Pointer, name *XmlChar, ExternalID *XmlChar, SystemID *XmlChar)
//go:linkname GetEntity C.getEntity
func GetEntity(ctx unsafe.Pointer, name *XmlChar) XmlEntityPtr
//go:linkname GetParameterEntity C.getParameterEntity
func GetParameterEntity(ctx unsafe.Pointer, name *XmlChar) XmlEntityPtr
//go:linkname ResolveEntity C.resolveEntity
func ResolveEntity(ctx unsafe.Pointer, publicId *XmlChar, systemId *XmlChar) XmlParserInputPtr
//go:linkname EntityDecl C.entityDecl
func EntityDecl(ctx unsafe.Pointer, name *XmlChar, type_ c.Int, publicId *XmlChar, systemId *XmlChar, content *XmlChar)
//go:linkname AttributeDecl C.attributeDecl
func AttributeDecl(ctx unsafe.Pointer, elem *XmlChar, fullname *XmlChar, type_ c.Int, def c.Int, defaultValue *XmlChar, tree XmlEnumerationPtr)
//go:linkname ElementDecl C.elementDecl
func ElementDecl(ctx unsafe.Pointer, name *XmlChar, type_ c.Int, content XmlElementContentPtr)
//go:linkname NotationDecl C.notationDecl
func NotationDecl(ctx unsafe.Pointer, name *XmlChar, publicId *XmlChar, systemId *XmlChar)
//go:linkname UnparsedEntityDecl C.unparsedEntityDecl
func UnparsedEntityDecl(ctx unsafe.Pointer, name *XmlChar, publicId *XmlChar, systemId *XmlChar, notationName *XmlChar)
//go:linkname StartDocument C.startDocument
func StartDocument(ctx unsafe.Pointer)
//go:linkname EndDocument C.endDocument
func EndDocument(ctx unsafe.Pointer)
//go:linkname Attribute C.attribute
func Attribute(ctx unsafe.Pointer, fullname *XmlChar, value *XmlChar)
//go:linkname StartElement C.startElement
func StartElement(ctx unsafe.Pointer, fullname *XmlChar, atts **XmlChar)
//go:linkname EndElement C.endElement
func EndElement(ctx unsafe.Pointer, name *XmlChar)
//go:linkname Reference C.reference
func Reference(ctx unsafe.Pointer, name *XmlChar)
//go:linkname Characters C.characters
func Characters(ctx unsafe.Pointer, ch *XmlChar, len c.Int)
//go:linkname IgnorableWhitespace C.ignorableWhitespace
func IgnorableWhitespace(ctx unsafe.Pointer, ch *XmlChar, len c.Int)
//go:linkname ProcessingInstruction C.processingInstruction
func ProcessingInstruction(ctx unsafe.Pointer, target *XmlChar, data *XmlChar)
//go:linkname GlobalNamespace C.globalNamespace
func GlobalNamespace(ctx unsafe.Pointer, href *XmlChar, prefix *XmlChar)
//go:linkname SetNamespace C.setNamespace
func SetNamespace(ctx unsafe.Pointer, name *XmlChar)
//go:linkname GetNamespace C.getNamespace
func GetNamespace(ctx unsafe.Pointer) XmlNsPtr
//go:linkname CheckNamespace C.checkNamespace
func CheckNamespace(ctx unsafe.Pointer, nameSpace *XmlChar) c.Int
//go:linkname NamespaceDecl C.namespaceDecl
func NamespaceDecl(ctx unsafe.Pointer, href *XmlChar, prefix *XmlChar)
//go:linkname Comment C.comment
func Comment(ctx unsafe.Pointer, value *XmlChar)
//go:linkname CdataBlock C.cdataBlock
func CdataBlock(ctx unsafe.Pointer, value *XmlChar, len c.Int)
// llgo:link (*XmlSAXHandlerV1).InitxmlDefaultSAXHandler C.initxmlDefaultSAXHandler
func (recv_ *XmlSAXHandlerV1) InitxmlDefaultSAXHandler(warning c.Int) {
}
// llgo:link (*XmlSAXHandlerV1).InithtmlDefaultSAXHandler C.inithtmlDefaultSAXHandler
func (recv_ *XmlSAXHandlerV1) InithtmlDefaultSAXHandler() {
}
