package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const XMLMAXTEXTLENGTH c.Int = 10000000
const XMLMAXHUGELENGTH c.Int = 1000000000
const XMLMAXNAMELENGTH c.Int = 50000
const XMLMAXDICTIONARYLIMIT c.Int = 10000000
const XMLMAXLOOKUPLIMIT c.Int = 10000000
const XMLMAXNAMELEN c.Int = 100
const INPUTCHUNK c.Int = 250
const XMLSUBSTITUTENONE c.Int = 0
const XMLSUBSTITUTEREF c.Int = 1
const XMLSUBSTITUTEPEREF c.Int = 2
const XMLSUBSTITUTEBOTH c.Int = 3
//go:linkname XmlIsLetter C.xmlIsLetter
func XmlIsLetter(c c.Int) c.Int
/**
 * Parser context.
 */
//go:linkname XmlCreateFileParserCtxt C.xmlCreateFileParserCtxt
func XmlCreateFileParserCtxt(filename *int8) XmlParserCtxtPtr
//go:linkname XmlCreateURLParserCtxt C.xmlCreateURLParserCtxt
func XmlCreateURLParserCtxt(filename *int8, options c.Int) XmlParserCtxtPtr
//go:linkname XmlCreateMemoryParserCtxt C.xmlCreateMemoryParserCtxt
func XmlCreateMemoryParserCtxt(buffer *int8, size c.Int) XmlParserCtxtPtr
// llgo:link (*XmlChar).XmlCreateEntityParserCtxt C.xmlCreateEntityParserCtxt
func (recv_ *XmlChar) XmlCreateEntityParserCtxt(ID *XmlChar, base *XmlChar) XmlParserCtxtPtr {
	return nil
}
//go:linkname XmlCtxtErrMemory C.xmlCtxtErrMemory
func XmlCtxtErrMemory(ctxt XmlParserCtxtPtr)
//go:linkname XmlSwitchEncoding C.xmlSwitchEncoding
func XmlSwitchEncoding(ctxt XmlParserCtxtPtr, enc XmlCharEncoding) c.Int
//go:linkname XmlSwitchEncodingName C.xmlSwitchEncodingName
func XmlSwitchEncodingName(ctxt XmlParserCtxtPtr, encoding *int8) c.Int
//go:linkname XmlSwitchToEncoding C.xmlSwitchToEncoding
func XmlSwitchToEncoding(ctxt XmlParserCtxtPtr, handler XmlCharEncodingHandlerPtr) c.Int
//go:linkname XmlSwitchInputEncoding C.xmlSwitchInputEncoding
func XmlSwitchInputEncoding(ctxt XmlParserCtxtPtr, input XmlParserInputPtr, handler XmlCharEncodingHandlerPtr) c.Int
/**
 * Input Streams.
 */
//go:linkname XmlNewStringInputStream C.xmlNewStringInputStream
func XmlNewStringInputStream(ctxt XmlParserCtxtPtr, buffer *XmlChar) XmlParserInputPtr
//go:linkname XmlNewEntityInputStream C.xmlNewEntityInputStream
func XmlNewEntityInputStream(ctxt XmlParserCtxtPtr, entity XmlEntityPtr) XmlParserInputPtr
//go:linkname XmlPushInput C.xmlPushInput
func XmlPushInput(ctxt XmlParserCtxtPtr, input XmlParserInputPtr) c.Int
//go:linkname XmlPopInput C.xmlPopInput
func XmlPopInput(ctxt XmlParserCtxtPtr) XmlChar
//go:linkname XmlFreeInputStream C.xmlFreeInputStream
func XmlFreeInputStream(input XmlParserInputPtr)
//go:linkname XmlNewInputFromFile C.xmlNewInputFromFile
func XmlNewInputFromFile(ctxt XmlParserCtxtPtr, filename *int8) XmlParserInputPtr
//go:linkname XmlNewInputStream C.xmlNewInputStream
func XmlNewInputStream(ctxt XmlParserCtxtPtr) XmlParserInputPtr
/**
 * Namespaces.
 */
//go:linkname XmlSplitQName C.xmlSplitQName
func XmlSplitQName(ctxt XmlParserCtxtPtr, name *XmlChar, prefix **XmlChar) *XmlChar
/**
 * Generic production rules.
 */
//go:linkname XmlParseName C.xmlParseName
func XmlParseName(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseNmtoken C.xmlParseNmtoken
func XmlParseNmtoken(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseEntityValue C.xmlParseEntityValue
func XmlParseEntityValue(ctxt XmlParserCtxtPtr, orig **XmlChar) *XmlChar
//go:linkname XmlParseAttValue C.xmlParseAttValue
func XmlParseAttValue(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseSystemLiteral C.xmlParseSystemLiteral
func XmlParseSystemLiteral(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParsePubidLiteral C.xmlParsePubidLiteral
func XmlParsePubidLiteral(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseCharData C.xmlParseCharData
func XmlParseCharData(ctxt XmlParserCtxtPtr, cdata c.Int)
//go:linkname XmlParseExternalID C.xmlParseExternalID
func XmlParseExternalID(ctxt XmlParserCtxtPtr, publicID **XmlChar, strict c.Int) *XmlChar
//go:linkname XmlParseComment C.xmlParseComment
func XmlParseComment(ctxt XmlParserCtxtPtr)
//go:linkname XmlParsePITarget C.xmlParsePITarget
func XmlParsePITarget(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParsePI C.xmlParsePI
func XmlParsePI(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseNotationDecl C.xmlParseNotationDecl
func XmlParseNotationDecl(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseEntityDecl C.xmlParseEntityDecl
func XmlParseEntityDecl(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseDefaultDecl C.xmlParseDefaultDecl
func XmlParseDefaultDecl(ctxt XmlParserCtxtPtr, value **XmlChar) c.Int
//go:linkname XmlParseNotationType C.xmlParseNotationType
func XmlParseNotationType(ctxt XmlParserCtxtPtr) XmlEnumerationPtr
//go:linkname XmlParseEnumerationType C.xmlParseEnumerationType
func XmlParseEnumerationType(ctxt XmlParserCtxtPtr) XmlEnumerationPtr
//go:linkname XmlParseEnumeratedType C.xmlParseEnumeratedType
func XmlParseEnumeratedType(ctxt XmlParserCtxtPtr, tree *XmlEnumerationPtr) c.Int
//go:linkname XmlParseAttributeType C.xmlParseAttributeType
func XmlParseAttributeType(ctxt XmlParserCtxtPtr, tree *XmlEnumerationPtr) c.Int
//go:linkname XmlParseAttributeListDecl C.xmlParseAttributeListDecl
func XmlParseAttributeListDecl(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseElementMixedContentDecl C.xmlParseElementMixedContentDecl
func XmlParseElementMixedContentDecl(ctxt XmlParserCtxtPtr, inputchk c.Int) XmlElementContentPtr
//go:linkname XmlParseElementChildrenContentDecl C.xmlParseElementChildrenContentDecl
func XmlParseElementChildrenContentDecl(ctxt XmlParserCtxtPtr, inputchk c.Int) XmlElementContentPtr
//go:linkname XmlParseElementContentDecl C.xmlParseElementContentDecl
func XmlParseElementContentDecl(ctxt XmlParserCtxtPtr, name *XmlChar, result *XmlElementContentPtr) c.Int
//go:linkname XmlParseElementDecl C.xmlParseElementDecl
func XmlParseElementDecl(ctxt XmlParserCtxtPtr) c.Int
//go:linkname XmlParseMarkupDecl C.xmlParseMarkupDecl
func XmlParseMarkupDecl(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseCharRef C.xmlParseCharRef
func XmlParseCharRef(ctxt XmlParserCtxtPtr) c.Int
//go:linkname XmlParseEntityRef C.xmlParseEntityRef
func XmlParseEntityRef(ctxt XmlParserCtxtPtr) XmlEntityPtr
//go:linkname XmlParseReference C.xmlParseReference
func XmlParseReference(ctxt XmlParserCtxtPtr)
//go:linkname XmlParsePEReference C.xmlParsePEReference
func XmlParsePEReference(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseDocTypeDecl C.xmlParseDocTypeDecl
func XmlParseDocTypeDecl(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseAttribute C.xmlParseAttribute
func XmlParseAttribute(ctxt XmlParserCtxtPtr, value **XmlChar) *XmlChar
//go:linkname XmlParseStartTag C.xmlParseStartTag
func XmlParseStartTag(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseEndTag C.xmlParseEndTag
func XmlParseEndTag(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseCDSect C.xmlParseCDSect
func XmlParseCDSect(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseContent C.xmlParseContent
func XmlParseContent(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseElement C.xmlParseElement
func XmlParseElement(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseVersionNum C.xmlParseVersionNum
func XmlParseVersionNum(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseVersionInfo C.xmlParseVersionInfo
func XmlParseVersionInfo(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseEncName C.xmlParseEncName
func XmlParseEncName(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseEncodingDecl C.xmlParseEncodingDecl
func XmlParseEncodingDecl(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseSDDecl C.xmlParseSDDecl
func XmlParseSDDecl(ctxt XmlParserCtxtPtr) c.Int
//go:linkname XmlParseXMLDecl C.xmlParseXMLDecl
func XmlParseXMLDecl(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseTextDecl C.xmlParseTextDecl
func XmlParseTextDecl(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseMisc C.xmlParseMisc
func XmlParseMisc(ctxt XmlParserCtxtPtr)
//go:linkname XmlParseExternalSubset C.xmlParseExternalSubset
func XmlParseExternalSubset(ctxt XmlParserCtxtPtr, ExternalID *XmlChar, SystemID *XmlChar)
//go:linkname XmlStringDecodeEntities C.xmlStringDecodeEntities
func XmlStringDecodeEntities(ctxt XmlParserCtxtPtr, str *XmlChar, what c.Int, end XmlChar, end2 XmlChar, end3 XmlChar) *XmlChar
//go:linkname XmlStringLenDecodeEntities C.xmlStringLenDecodeEntities
func XmlStringLenDecodeEntities(ctxt XmlParserCtxtPtr, str *XmlChar, len c.Int, what c.Int, end XmlChar, end2 XmlChar, end3 XmlChar) *XmlChar
//go:linkname NodePush C.nodePush
func NodePush(ctxt XmlParserCtxtPtr, value XmlNodePtr) c.Int
//go:linkname NodePop C.nodePop
func NodePop(ctxt XmlParserCtxtPtr) XmlNodePtr
//go:linkname InputPush C.inputPush
func InputPush(ctxt XmlParserCtxtPtr, value XmlParserInputPtr) c.Int
//go:linkname InputPop C.inputPop
func InputPop(ctxt XmlParserCtxtPtr) XmlParserInputPtr
//go:linkname NamePop C.namePop
func NamePop(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname NamePush C.namePush
func NamePush(ctxt XmlParserCtxtPtr, value *XmlChar) c.Int
//go:linkname XmlSkipBlankChars C.xmlSkipBlankChars
func XmlSkipBlankChars(ctxt XmlParserCtxtPtr) c.Int
//go:linkname XmlStringCurrentChar C.xmlStringCurrentChar
func XmlStringCurrentChar(ctxt XmlParserCtxtPtr, cur *XmlChar, len *c.Int) c.Int
//go:linkname XmlParserHandlePEReference C.xmlParserHandlePEReference
func XmlParserHandlePEReference(ctxt XmlParserCtxtPtr)
// llgo:link (*XmlChar).XmlCheckLanguageID C.xmlCheckLanguageID
func (recv_ *XmlChar) XmlCheckLanguageID() c.Int {
	return 0
}
//go:linkname XmlCurrentChar C.xmlCurrentChar
func XmlCurrentChar(ctxt XmlParserCtxtPtr, len *c.Int) c.Int
// llgo:link (*XmlChar).XmlCopyCharMultiByte C.xmlCopyCharMultiByte
func (recv_ *XmlChar) XmlCopyCharMultiByte(val c.Int) c.Int {
	return 0
}
//go:linkname XmlCopyChar C.xmlCopyChar
func XmlCopyChar(len c.Int, out *XmlChar, val c.Int) c.Int
//go:linkname XmlNextChar C.xmlNextChar
func XmlNextChar(ctxt XmlParserCtxtPtr)
//go:linkname XmlParserInputShrink C.xmlParserInputShrink
func XmlParserInputShrink(in XmlParserInputPtr)
// llgo:type C
type XmlEntityReferenceFunc func(XmlEntityPtr, XmlNodePtr, XmlNodePtr)
//go:linkname XmlSetEntityReferenceFunc C.xmlSetEntityReferenceFunc
func XmlSetEntityReferenceFunc(func_ XmlEntityReferenceFunc)
//go:linkname XmlParseQuotedString C.xmlParseQuotedString
func XmlParseQuotedString(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParseNamespace C.xmlParseNamespace
func XmlParseNamespace(ctxt XmlParserCtxtPtr)
//go:linkname XmlNamespaceParseNSDef C.xmlNamespaceParseNSDef
func XmlNamespaceParseNSDef(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlScanName C.xmlScanName
func XmlScanName(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlNamespaceParseNCName C.xmlNamespaceParseNCName
func XmlNamespaceParseNCName(ctxt XmlParserCtxtPtr) *XmlChar
//go:linkname XmlParserHandleReference C.xmlParserHandleReference
func XmlParserHandleReference(ctxt XmlParserCtxtPtr)
//go:linkname XmlNamespaceParseQName C.xmlNamespaceParseQName
func XmlNamespaceParseQName(ctxt XmlParserCtxtPtr, prefix **XmlChar) *XmlChar
/**
 * Entities
 */
//go:linkname XmlDecodeEntities C.xmlDecodeEntities
func XmlDecodeEntities(ctxt XmlParserCtxtPtr, len c.Int, what c.Int, end XmlChar, end2 XmlChar, end3 XmlChar) *XmlChar
//go:linkname XmlHandleEntity C.xmlHandleEntity
func XmlHandleEntity(ctxt XmlParserCtxtPtr, entity XmlEntityPtr)
