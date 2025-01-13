package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XmlParserSeverities c.Int

const (
	XmlParserSeveritiesXMLPARSERSEVERITYVALIDITYWARNING XmlParserSeverities = 1
	XmlParserSeveritiesXMLPARSERSEVERITYVALIDITYERROR   XmlParserSeverities = 2
	XmlParserSeveritiesXMLPARSERSEVERITYWARNING         XmlParserSeverities = 3
	XmlParserSeveritiesXMLPARSERSEVERITYERROR           XmlParserSeverities = 4
)

type XmlTextReaderMode c.Int

const (
	XmlTextReaderModeXMLTEXTREADERMODEINITIAL     XmlTextReaderMode = 0
	XmlTextReaderModeXMLTEXTREADERMODEINTERACTIVE XmlTextReaderMode = 1
	XmlTextReaderModeXMLTEXTREADERMODEERROR       XmlTextReaderMode = 2
	XmlTextReaderModeXMLTEXTREADERMODEEOF         XmlTextReaderMode = 3
	XmlTextReaderModeXMLTEXTREADERMODECLOSED      XmlTextReaderMode = 4
	XmlTextReaderModeXMLTEXTREADERMODEREADING     XmlTextReaderMode = 5
)

type XmlParserProperties c.Int

const (
	XmlParserPropertiesXMLPARSERLOADDTD       XmlParserProperties = 1
	XmlParserPropertiesXMLPARSERDEFAULTATTRS  XmlParserProperties = 2
	XmlParserPropertiesXMLPARSERVALIDATE      XmlParserProperties = 3
	XmlParserPropertiesXMLPARSERSUBSTENTITIES XmlParserProperties = 4
)

type XmlReaderTypes c.Int

const (
	XmlReaderTypesXMLREADERTYPENONE                  XmlReaderTypes = 0
	XmlReaderTypesXMLREADERTYPEELEMENT               XmlReaderTypes = 1
	XmlReaderTypesXMLREADERTYPEATTRIBUTE             XmlReaderTypes = 2
	XmlReaderTypesXMLREADERTYPETEXT                  XmlReaderTypes = 3
	XmlReaderTypesXMLREADERTYPECDATA                 XmlReaderTypes = 4
	XmlReaderTypesXMLREADERTYPEENTITYREFERENCE       XmlReaderTypes = 5
	XmlReaderTypesXMLREADERTYPEENTITY                XmlReaderTypes = 6
	XmlReaderTypesXMLREADERTYPEPROCESSINGINSTRUCTION XmlReaderTypes = 7
	XmlReaderTypesXMLREADERTYPECOMMENT               XmlReaderTypes = 8
	XmlReaderTypesXMLREADERTYPEDOCUMENT              XmlReaderTypes = 9
	XmlReaderTypesXMLREADERTYPEDOCUMENTTYPE          XmlReaderTypes = 10
	XmlReaderTypesXMLREADERTYPEDOCUMENTFRAGMENT      XmlReaderTypes = 11
	XmlReaderTypesXMLREADERTYPENOTATION              XmlReaderTypes = 12
	XmlReaderTypesXMLREADERTYPEWHITESPACE            XmlReaderTypes = 13
	XmlReaderTypesXMLREADERTYPESIGNIFICANTWHITESPACE XmlReaderTypes = 14
	XmlReaderTypesXMLREADERTYPEENDELEMENT            XmlReaderTypes = 15
	XmlReaderTypesXMLREADERTYPEENDENTITY             XmlReaderTypes = 16
	XmlReaderTypesXMLREADERTYPEXMLDECLARATION        XmlReaderTypes = 17
)

type X_XmlTextReader struct {
	Unused [8]uint8
}
type XmlTextReader X_XmlTextReader
type XmlTextReaderPtr *XmlTextReader
//go:linkname XmlNewTextReader C.xmlNewTextReader
func XmlNewTextReader(input XmlParserInputBufferPtr, URI *int8) XmlTextReaderPtr
//go:linkname XmlNewTextReaderFilename C.xmlNewTextReaderFilename
func XmlNewTextReaderFilename(URI *int8) XmlTextReaderPtr
//go:linkname XmlFreeTextReader C.xmlFreeTextReader
func XmlFreeTextReader(reader XmlTextReaderPtr)
//go:linkname XmlTextReaderSetup C.xmlTextReaderSetup
func XmlTextReaderSetup(reader XmlTextReaderPtr, input XmlParserInputBufferPtr, URL *int8, encoding *int8, options c.Int) c.Int
//go:linkname XmlTextReaderSetMaxAmplification C.xmlTextReaderSetMaxAmplification
func XmlTextReaderSetMaxAmplification(reader XmlTextReaderPtr, maxAmpl c.Uint)
//go:linkname XmlTextReaderGetLastError C.xmlTextReaderGetLastError
func XmlTextReaderGetLastError(reader XmlTextReaderPtr) *XmlError
//go:linkname XmlTextReaderRead C.xmlTextReaderRead
func XmlTextReaderRead(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderReadInnerXml C.xmlTextReaderReadInnerXml
func XmlTextReaderReadInnerXml(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderReadOuterXml C.xmlTextReaderReadOuterXml
func XmlTextReaderReadOuterXml(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderReadString C.xmlTextReaderReadString
func XmlTextReaderReadString(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderReadAttributeValue C.xmlTextReaderReadAttributeValue
func XmlTextReaderReadAttributeValue(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderAttributeCount C.xmlTextReaderAttributeCount
func XmlTextReaderAttributeCount(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderDepth C.xmlTextReaderDepth
func XmlTextReaderDepth(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderHasAttributes C.xmlTextReaderHasAttributes
func XmlTextReaderHasAttributes(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderHasValue C.xmlTextReaderHasValue
func XmlTextReaderHasValue(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderIsDefault C.xmlTextReaderIsDefault
func XmlTextReaderIsDefault(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderIsEmptyElement C.xmlTextReaderIsEmptyElement
func XmlTextReaderIsEmptyElement(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderNodeType C.xmlTextReaderNodeType
func XmlTextReaderNodeType(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderQuoteChar C.xmlTextReaderQuoteChar
func XmlTextReaderQuoteChar(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderReadState C.xmlTextReaderReadState
func XmlTextReaderReadState(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderIsNamespaceDecl C.xmlTextReaderIsNamespaceDecl
func XmlTextReaderIsNamespaceDecl(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderConstBaseUri C.xmlTextReaderConstBaseUri
func XmlTextReaderConstBaseUri(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderConstLocalName C.xmlTextReaderConstLocalName
func XmlTextReaderConstLocalName(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderConstName C.xmlTextReaderConstName
func XmlTextReaderConstName(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderConstNamespaceUri C.xmlTextReaderConstNamespaceUri
func XmlTextReaderConstNamespaceUri(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderConstPrefix C.xmlTextReaderConstPrefix
func XmlTextReaderConstPrefix(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderConstXmlLang C.xmlTextReaderConstXmlLang
func XmlTextReaderConstXmlLang(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderConstString C.xmlTextReaderConstString
func XmlTextReaderConstString(reader XmlTextReaderPtr, str *XmlChar) *XmlChar
//go:linkname XmlTextReaderConstValue C.xmlTextReaderConstValue
func XmlTextReaderConstValue(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderBaseUri C.xmlTextReaderBaseUri
func XmlTextReaderBaseUri(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderLocalName C.xmlTextReaderLocalName
func XmlTextReaderLocalName(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderName C.xmlTextReaderName
func XmlTextReaderName(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderNamespaceUri C.xmlTextReaderNamespaceUri
func XmlTextReaderNamespaceUri(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderPrefix C.xmlTextReaderPrefix
func XmlTextReaderPrefix(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderXmlLang C.xmlTextReaderXmlLang
func XmlTextReaderXmlLang(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderValue C.xmlTextReaderValue
func XmlTextReaderValue(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderClose C.xmlTextReaderClose
func XmlTextReaderClose(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderGetAttributeNo C.xmlTextReaderGetAttributeNo
func XmlTextReaderGetAttributeNo(reader XmlTextReaderPtr, no c.Int) *XmlChar
//go:linkname XmlTextReaderGetAttribute C.xmlTextReaderGetAttribute
func XmlTextReaderGetAttribute(reader XmlTextReaderPtr, name *XmlChar) *XmlChar
//go:linkname XmlTextReaderGetAttributeNs C.xmlTextReaderGetAttributeNs
func XmlTextReaderGetAttributeNs(reader XmlTextReaderPtr, localName *XmlChar, namespaceURI *XmlChar) *XmlChar
//go:linkname XmlTextReaderGetRemainder C.xmlTextReaderGetRemainder
func XmlTextReaderGetRemainder(reader XmlTextReaderPtr) XmlParserInputBufferPtr
//go:linkname XmlTextReaderLookupNamespace C.xmlTextReaderLookupNamespace
func XmlTextReaderLookupNamespace(reader XmlTextReaderPtr, prefix *XmlChar) *XmlChar
//go:linkname XmlTextReaderMoveToAttributeNo C.xmlTextReaderMoveToAttributeNo
func XmlTextReaderMoveToAttributeNo(reader XmlTextReaderPtr, no c.Int) c.Int
//go:linkname XmlTextReaderMoveToAttribute C.xmlTextReaderMoveToAttribute
func XmlTextReaderMoveToAttribute(reader XmlTextReaderPtr, name *XmlChar) c.Int
//go:linkname XmlTextReaderMoveToAttributeNs C.xmlTextReaderMoveToAttributeNs
func XmlTextReaderMoveToAttributeNs(reader XmlTextReaderPtr, localName *XmlChar, namespaceURI *XmlChar) c.Int
//go:linkname XmlTextReaderMoveToFirstAttribute C.xmlTextReaderMoveToFirstAttribute
func XmlTextReaderMoveToFirstAttribute(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderMoveToNextAttribute C.xmlTextReaderMoveToNextAttribute
func XmlTextReaderMoveToNextAttribute(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderMoveToElement C.xmlTextReaderMoveToElement
func XmlTextReaderMoveToElement(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderNormalization C.xmlTextReaderNormalization
func XmlTextReaderNormalization(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderConstEncoding C.xmlTextReaderConstEncoding
func XmlTextReaderConstEncoding(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderSetParserProp C.xmlTextReaderSetParserProp
func XmlTextReaderSetParserProp(reader XmlTextReaderPtr, prop c.Int, value c.Int) c.Int
//go:linkname XmlTextReaderGetParserProp C.xmlTextReaderGetParserProp
func XmlTextReaderGetParserProp(reader XmlTextReaderPtr, prop c.Int) c.Int
//go:linkname XmlTextReaderCurrentNode C.xmlTextReaderCurrentNode
func XmlTextReaderCurrentNode(reader XmlTextReaderPtr) XmlNodePtr
//go:linkname XmlTextReaderGetParserLineNumber C.xmlTextReaderGetParserLineNumber
func XmlTextReaderGetParserLineNumber(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderGetParserColumnNumber C.xmlTextReaderGetParserColumnNumber
func XmlTextReaderGetParserColumnNumber(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderPreserve C.xmlTextReaderPreserve
func XmlTextReaderPreserve(reader XmlTextReaderPtr) XmlNodePtr
//go:linkname XmlTextReaderPreservePattern C.xmlTextReaderPreservePattern
func XmlTextReaderPreservePattern(reader XmlTextReaderPtr, pattern *XmlChar, namespaces **XmlChar) c.Int
//go:linkname XmlTextReaderCurrentDoc C.xmlTextReaderCurrentDoc
func XmlTextReaderCurrentDoc(reader XmlTextReaderPtr) XmlDocPtr
//go:linkname XmlTextReaderExpand C.xmlTextReaderExpand
func XmlTextReaderExpand(reader XmlTextReaderPtr) XmlNodePtr
//go:linkname XmlTextReaderNext C.xmlTextReaderNext
func XmlTextReaderNext(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderNextSibling C.xmlTextReaderNextSibling
func XmlTextReaderNextSibling(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderIsValid C.xmlTextReaderIsValid
func XmlTextReaderIsValid(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderRelaxNGValidate C.xmlTextReaderRelaxNGValidate
func XmlTextReaderRelaxNGValidate(reader XmlTextReaderPtr, rng *int8) c.Int
//go:linkname XmlTextReaderRelaxNGValidateCtxt C.xmlTextReaderRelaxNGValidateCtxt
func XmlTextReaderRelaxNGValidateCtxt(reader XmlTextReaderPtr, ctxt XmlRelaxNGValidCtxtPtr, options c.Int) c.Int
//go:linkname XmlTextReaderRelaxNGSetSchema C.xmlTextReaderRelaxNGSetSchema
func XmlTextReaderRelaxNGSetSchema(reader XmlTextReaderPtr, schema XmlRelaxNGPtr) c.Int
//go:linkname XmlTextReaderSchemaValidate C.xmlTextReaderSchemaValidate
func XmlTextReaderSchemaValidate(reader XmlTextReaderPtr, xsd *int8) c.Int
//go:linkname XmlTextReaderSchemaValidateCtxt C.xmlTextReaderSchemaValidateCtxt
func XmlTextReaderSchemaValidateCtxt(reader XmlTextReaderPtr, ctxt XmlSchemaValidCtxtPtr, options c.Int) c.Int
//go:linkname XmlTextReaderSetSchema C.xmlTextReaderSetSchema
func XmlTextReaderSetSchema(reader XmlTextReaderPtr, schema XmlSchemaPtr) c.Int
//go:linkname XmlTextReaderConstXmlVersion C.xmlTextReaderConstXmlVersion
func XmlTextReaderConstXmlVersion(reader XmlTextReaderPtr) *XmlChar
//go:linkname XmlTextReaderStandalone C.xmlTextReaderStandalone
func XmlTextReaderStandalone(reader XmlTextReaderPtr) c.Int
//go:linkname XmlTextReaderByteConsumed C.xmlTextReaderByteConsumed
func XmlTextReaderByteConsumed(reader XmlTextReaderPtr) c.Long
//go:linkname XmlReaderWalker C.xmlReaderWalker
func XmlReaderWalker(doc XmlDocPtr) XmlTextReaderPtr
// llgo:link (*XmlChar).XmlReaderForDoc C.xmlReaderForDoc
func (recv_ *XmlChar) XmlReaderForDoc(URL *int8, encoding *int8, options c.Int) XmlTextReaderPtr {
	return nil
}
//go:linkname XmlReaderForFile C.xmlReaderForFile
func XmlReaderForFile(filename *int8, encoding *int8, options c.Int) XmlTextReaderPtr
//go:linkname XmlReaderForMemory C.xmlReaderForMemory
func XmlReaderForMemory(buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) XmlTextReaderPtr
//go:linkname XmlReaderForFd C.xmlReaderForFd
func XmlReaderForFd(fd c.Int, URL *int8, encoding *int8, options c.Int) XmlTextReaderPtr
//go:linkname XmlReaderForIO C.xmlReaderForIO
func XmlReaderForIO(ioread XmlInputReadCallback, ioclose XmlInputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) XmlTextReaderPtr
//go:linkname XmlReaderNewWalker C.xmlReaderNewWalker
func XmlReaderNewWalker(reader XmlTextReaderPtr, doc XmlDocPtr) c.Int
//go:linkname XmlReaderNewDoc C.xmlReaderNewDoc
func XmlReaderNewDoc(reader XmlTextReaderPtr, cur *XmlChar, URL *int8, encoding *int8, options c.Int) c.Int
//go:linkname XmlReaderNewFile C.xmlReaderNewFile
func XmlReaderNewFile(reader XmlTextReaderPtr, filename *int8, encoding *int8, options c.Int) c.Int
//go:linkname XmlReaderNewMemory C.xmlReaderNewMemory
func XmlReaderNewMemory(reader XmlTextReaderPtr, buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) c.Int
//go:linkname XmlReaderNewFd C.xmlReaderNewFd
func XmlReaderNewFd(reader XmlTextReaderPtr, fd c.Int, URL *int8, encoding *int8, options c.Int) c.Int
//go:linkname XmlReaderNewIO C.xmlReaderNewIO
func XmlReaderNewIO(reader XmlTextReaderPtr, ioread XmlInputReadCallback, ioclose XmlInputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) c.Int

type XmlTextReaderLocatorPtr unsafe.Pointer
// llgo:type C
type XmlTextReaderErrorFunc func(unsafe.Pointer, *int8, XmlParserSeverities, XmlTextReaderLocatorPtr)
//go:linkname XmlTextReaderLocatorLineNumber C.xmlTextReaderLocatorLineNumber
func XmlTextReaderLocatorLineNumber(locator XmlTextReaderLocatorPtr) c.Int
//go:linkname XmlTextReaderLocatorBaseURI C.xmlTextReaderLocatorBaseURI
func XmlTextReaderLocatorBaseURI(locator XmlTextReaderLocatorPtr) *XmlChar
//go:linkname XmlTextReaderSetErrorHandler C.xmlTextReaderSetErrorHandler
func XmlTextReaderSetErrorHandler(reader XmlTextReaderPtr, f XmlTextReaderErrorFunc, arg unsafe.Pointer)
//go:linkname XmlTextReaderSetStructuredErrorHandler C.xmlTextReaderSetStructuredErrorHandler
func XmlTextReaderSetStructuredErrorHandler(reader XmlTextReaderPtr, f XmlStructuredErrorFunc, arg unsafe.Pointer)
//go:linkname XmlTextReaderGetErrorHandler C.xmlTextReaderGetErrorHandler
func XmlTextReaderGetErrorHandler(reader XmlTextReaderPtr, f XmlTextReaderErrorFunc, arg *unsafe.Pointer)
