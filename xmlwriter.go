package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type X_XmlTextWriter struct {
	Unused [8]uint8
}
type XmlTextWriter X_XmlTextWriter
type XmlTextWriterPtr *XmlTextWriter
//go:linkname XmlNewTextWriter C.xmlNewTextWriter
func XmlNewTextWriter(out XmlOutputBufferPtr) XmlTextWriterPtr
//go:linkname XmlNewTextWriterFilename C.xmlNewTextWriterFilename
func XmlNewTextWriterFilename(uri *int8, compression c.Int) XmlTextWriterPtr
//go:linkname XmlNewTextWriterMemory C.xmlNewTextWriterMemory
func XmlNewTextWriterMemory(buf XmlBufferPtr, compression c.Int) XmlTextWriterPtr
//go:linkname XmlNewTextWriterPushParser C.xmlNewTextWriterPushParser
func XmlNewTextWriterPushParser(ctxt XmlParserCtxtPtr, compression c.Int) XmlTextWriterPtr
//go:linkname XmlNewTextWriterDoc C.xmlNewTextWriterDoc
func XmlNewTextWriterDoc(doc *XmlDocPtr, compression c.Int) XmlTextWriterPtr
//go:linkname XmlNewTextWriterTree C.xmlNewTextWriterTree
func XmlNewTextWriterTree(doc XmlDocPtr, node XmlNodePtr, compression c.Int) XmlTextWriterPtr
//go:linkname XmlFreeTextWriter C.xmlFreeTextWriter
func XmlFreeTextWriter(writer XmlTextWriterPtr)
//go:linkname XmlTextWriterStartDocument C.xmlTextWriterStartDocument
func XmlTextWriterStartDocument(writer XmlTextWriterPtr, version *int8, encoding *int8, standalone *int8) c.Int
//go:linkname XmlTextWriterEndDocument C.xmlTextWriterEndDocument
func XmlTextWriterEndDocument(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterStartComment C.xmlTextWriterStartComment
func XmlTextWriterStartComment(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterEndComment C.xmlTextWriterEndComment
func XmlTextWriterEndComment(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterWriteFormatComment C.xmlTextWriterWriteFormatComment
func XmlTextWriterWriteFormatComment(writer XmlTextWriterPtr, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteComment C.xmlTextWriterWriteComment
func XmlTextWriterWriteComment(writer XmlTextWriterPtr, content *XmlChar) c.Int
//go:linkname XmlTextWriterStartElement C.xmlTextWriterStartElement
func XmlTextWriterStartElement(writer XmlTextWriterPtr, name *XmlChar) c.Int
//go:linkname XmlTextWriterStartElementNS C.xmlTextWriterStartElementNS
func XmlTextWriterStartElementNS(writer XmlTextWriterPtr, prefix *XmlChar, name *XmlChar, namespaceURI *XmlChar) c.Int
//go:linkname XmlTextWriterEndElement C.xmlTextWriterEndElement
func XmlTextWriterEndElement(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterFullEndElement C.xmlTextWriterFullEndElement
func XmlTextWriterFullEndElement(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterWriteFormatElement C.xmlTextWriterWriteFormatElement
func XmlTextWriterWriteFormatElement(writer XmlTextWriterPtr, name *XmlChar, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteElement C.xmlTextWriterWriteElement
func XmlTextWriterWriteElement(writer XmlTextWriterPtr, name *XmlChar, content *XmlChar) c.Int
//go:linkname XmlTextWriterWriteFormatElementNS C.xmlTextWriterWriteFormatElementNS
func XmlTextWriterWriteFormatElementNS(writer XmlTextWriterPtr, prefix *XmlChar, name *XmlChar, namespaceURI *XmlChar, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteElementNS C.xmlTextWriterWriteElementNS
func XmlTextWriterWriteElementNS(writer XmlTextWriterPtr, prefix *XmlChar, name *XmlChar, namespaceURI *XmlChar, content *XmlChar) c.Int
//go:linkname XmlTextWriterWriteFormatRaw C.xmlTextWriterWriteFormatRaw
func XmlTextWriterWriteFormatRaw(writer XmlTextWriterPtr, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteRawLen C.xmlTextWriterWriteRawLen
func XmlTextWriterWriteRawLen(writer XmlTextWriterPtr, content *XmlChar, len c.Int) c.Int
//go:linkname XmlTextWriterWriteRaw C.xmlTextWriterWriteRaw
func XmlTextWriterWriteRaw(writer XmlTextWriterPtr, content *XmlChar) c.Int
//go:linkname XmlTextWriterWriteFormatString C.xmlTextWriterWriteFormatString
func XmlTextWriterWriteFormatString(writer XmlTextWriterPtr, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteString C.xmlTextWriterWriteString
func XmlTextWriterWriteString(writer XmlTextWriterPtr, content *XmlChar) c.Int
//go:linkname XmlTextWriterWriteBase64 C.xmlTextWriterWriteBase64
func XmlTextWriterWriteBase64(writer XmlTextWriterPtr, data *int8, start c.Int, len c.Int) c.Int
//go:linkname XmlTextWriterWriteBinHex C.xmlTextWriterWriteBinHex
func XmlTextWriterWriteBinHex(writer XmlTextWriterPtr, data *int8, start c.Int, len c.Int) c.Int
//go:linkname XmlTextWriterStartAttribute C.xmlTextWriterStartAttribute
func XmlTextWriterStartAttribute(writer XmlTextWriterPtr, name *XmlChar) c.Int
//go:linkname XmlTextWriterStartAttributeNS C.xmlTextWriterStartAttributeNS
func XmlTextWriterStartAttributeNS(writer XmlTextWriterPtr, prefix *XmlChar, name *XmlChar, namespaceURI *XmlChar) c.Int
//go:linkname XmlTextWriterEndAttribute C.xmlTextWriterEndAttribute
func XmlTextWriterEndAttribute(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterWriteFormatAttribute C.xmlTextWriterWriteFormatAttribute
func XmlTextWriterWriteFormatAttribute(writer XmlTextWriterPtr, name *XmlChar, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteAttribute C.xmlTextWriterWriteAttribute
func XmlTextWriterWriteAttribute(writer XmlTextWriterPtr, name *XmlChar, content *XmlChar) c.Int
//go:linkname XmlTextWriterWriteFormatAttributeNS C.xmlTextWriterWriteFormatAttributeNS
func XmlTextWriterWriteFormatAttributeNS(writer XmlTextWriterPtr, prefix *XmlChar, name *XmlChar, namespaceURI *XmlChar, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteAttributeNS C.xmlTextWriterWriteAttributeNS
func XmlTextWriterWriteAttributeNS(writer XmlTextWriterPtr, prefix *XmlChar, name *XmlChar, namespaceURI *XmlChar, content *XmlChar) c.Int
//go:linkname XmlTextWriterStartPI C.xmlTextWriterStartPI
func XmlTextWriterStartPI(writer XmlTextWriterPtr, target *XmlChar) c.Int
//go:linkname XmlTextWriterEndPI C.xmlTextWriterEndPI
func XmlTextWriterEndPI(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterWriteFormatPI C.xmlTextWriterWriteFormatPI
func XmlTextWriterWriteFormatPI(writer XmlTextWriterPtr, target *XmlChar, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWritePI C.xmlTextWriterWritePI
func XmlTextWriterWritePI(writer XmlTextWriterPtr, target *XmlChar, content *XmlChar) c.Int
//go:linkname XmlTextWriterStartCDATA C.xmlTextWriterStartCDATA
func XmlTextWriterStartCDATA(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterEndCDATA C.xmlTextWriterEndCDATA
func XmlTextWriterEndCDATA(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterWriteFormatCDATA C.xmlTextWriterWriteFormatCDATA
func XmlTextWriterWriteFormatCDATA(writer XmlTextWriterPtr, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteCDATA C.xmlTextWriterWriteCDATA
func XmlTextWriterWriteCDATA(writer XmlTextWriterPtr, content *XmlChar) c.Int
//go:linkname XmlTextWriterStartDTD C.xmlTextWriterStartDTD
func XmlTextWriterStartDTD(writer XmlTextWriterPtr, name *XmlChar, pubid *XmlChar, sysid *XmlChar) c.Int
//go:linkname XmlTextWriterEndDTD C.xmlTextWriterEndDTD
func XmlTextWriterEndDTD(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterWriteFormatDTD C.xmlTextWriterWriteFormatDTD
func XmlTextWriterWriteFormatDTD(writer XmlTextWriterPtr, name *XmlChar, pubid *XmlChar, sysid *XmlChar, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteDTD C.xmlTextWriterWriteDTD
func XmlTextWriterWriteDTD(writer XmlTextWriterPtr, name *XmlChar, pubid *XmlChar, sysid *XmlChar, subset *XmlChar) c.Int
//go:linkname XmlTextWriterStartDTDElement C.xmlTextWriterStartDTDElement
func XmlTextWriterStartDTDElement(writer XmlTextWriterPtr, name *XmlChar) c.Int
//go:linkname XmlTextWriterEndDTDElement C.xmlTextWriterEndDTDElement
func XmlTextWriterEndDTDElement(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterWriteFormatDTDElement C.xmlTextWriterWriteFormatDTDElement
func XmlTextWriterWriteFormatDTDElement(writer XmlTextWriterPtr, name *XmlChar, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteDTDElement C.xmlTextWriterWriteDTDElement
func XmlTextWriterWriteDTDElement(writer XmlTextWriterPtr, name *XmlChar, content *XmlChar) c.Int
//go:linkname XmlTextWriterStartDTDAttlist C.xmlTextWriterStartDTDAttlist
func XmlTextWriterStartDTDAttlist(writer XmlTextWriterPtr, name *XmlChar) c.Int
//go:linkname XmlTextWriterEndDTDAttlist C.xmlTextWriterEndDTDAttlist
func XmlTextWriterEndDTDAttlist(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterWriteFormatDTDAttlist C.xmlTextWriterWriteFormatDTDAttlist
func XmlTextWriterWriteFormatDTDAttlist(writer XmlTextWriterPtr, name *XmlChar, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteDTDAttlist C.xmlTextWriterWriteDTDAttlist
func XmlTextWriterWriteDTDAttlist(writer XmlTextWriterPtr, name *XmlChar, content *XmlChar) c.Int
//go:linkname XmlTextWriterStartDTDEntity C.xmlTextWriterStartDTDEntity
func XmlTextWriterStartDTDEntity(writer XmlTextWriterPtr, pe c.Int, name *XmlChar) c.Int
//go:linkname XmlTextWriterEndDTDEntity C.xmlTextWriterEndDTDEntity
func XmlTextWriterEndDTDEntity(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterWriteFormatDTDInternalEntity C.xmlTextWriterWriteFormatDTDInternalEntity
func XmlTextWriterWriteFormatDTDInternalEntity(writer XmlTextWriterPtr, pe c.Int, name *XmlChar, format *int8, __llgo_va_list ...interface{}) c.Int
//go:linkname XmlTextWriterWriteDTDInternalEntity C.xmlTextWriterWriteDTDInternalEntity
func XmlTextWriterWriteDTDInternalEntity(writer XmlTextWriterPtr, pe c.Int, name *XmlChar, content *XmlChar) c.Int
//go:linkname XmlTextWriterWriteDTDExternalEntity C.xmlTextWriterWriteDTDExternalEntity
func XmlTextWriterWriteDTDExternalEntity(writer XmlTextWriterPtr, pe c.Int, name *XmlChar, pubid *XmlChar, sysid *XmlChar, ndataid *XmlChar) c.Int
//go:linkname XmlTextWriterWriteDTDExternalEntityContents C.xmlTextWriterWriteDTDExternalEntityContents
func XmlTextWriterWriteDTDExternalEntityContents(writer XmlTextWriterPtr, pubid *XmlChar, sysid *XmlChar, ndataid *XmlChar) c.Int
//go:linkname XmlTextWriterWriteDTDEntity C.xmlTextWriterWriteDTDEntity
func XmlTextWriterWriteDTDEntity(writer XmlTextWriterPtr, pe c.Int, name *XmlChar, pubid *XmlChar, sysid *XmlChar, ndataid *XmlChar, content *XmlChar) c.Int
//go:linkname XmlTextWriterWriteDTDNotation C.xmlTextWriterWriteDTDNotation
func XmlTextWriterWriteDTDNotation(writer XmlTextWriterPtr, name *XmlChar, pubid *XmlChar, sysid *XmlChar) c.Int
//go:linkname XmlTextWriterSetIndent C.xmlTextWriterSetIndent
func XmlTextWriterSetIndent(writer XmlTextWriterPtr, indent c.Int) c.Int
//go:linkname XmlTextWriterSetIndentString C.xmlTextWriterSetIndentString
func XmlTextWriterSetIndentString(writer XmlTextWriterPtr, str *XmlChar) c.Int
//go:linkname XmlTextWriterSetQuoteChar C.xmlTextWriterSetQuoteChar
func XmlTextWriterSetQuoteChar(writer XmlTextWriterPtr, quotechar XmlChar) c.Int
//go:linkname XmlTextWriterFlush C.xmlTextWriterFlush
func XmlTextWriterFlush(writer XmlTextWriterPtr) c.Int
//go:linkname XmlTextWriterClose C.xmlTextWriterClose
func XmlTextWriterClose(writer XmlTextWriterPtr) c.Int
