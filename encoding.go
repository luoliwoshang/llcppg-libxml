package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type XmlCharEncError c.Int

const (
	XmlCharEncErrorXMLENCERRSUCCESS  XmlCharEncError = 0
	XmlCharEncErrorXMLENCERRSPACE    XmlCharEncError = -1
	XmlCharEncErrorXMLENCERRINPUT    XmlCharEncError = -2
	XmlCharEncErrorXMLENCERRPARTIAL  XmlCharEncError = -3
	XmlCharEncErrorXMLENCERRINTERNAL XmlCharEncError = -4
	XmlCharEncErrorXMLENCERRMEMORY   XmlCharEncError = -5
)

type XmlCharEncoding c.Int

const (
	XmlCharEncodingXMLCHARENCODINGERROR    XmlCharEncoding = -1
	XmlCharEncodingXMLCHARENCODINGNONE     XmlCharEncoding = 0
	XmlCharEncodingXMLCHARENCODINGUTF8     XmlCharEncoding = 1
	XmlCharEncodingXMLCHARENCODINGUTF16LE  XmlCharEncoding = 2
	XmlCharEncodingXMLCHARENCODINGUTF16BE  XmlCharEncoding = 3
	XmlCharEncodingXMLCHARENCODINGUCS4LE   XmlCharEncoding = 4
	XmlCharEncodingXMLCHARENCODINGUCS4BE   XmlCharEncoding = 5
	XmlCharEncodingXMLCHARENCODINGEBCDIC   XmlCharEncoding = 6
	XmlCharEncodingXMLCHARENCODINGUCS42143 XmlCharEncoding = 7
	XmlCharEncodingXMLCHARENCODINGUCS43412 XmlCharEncoding = 8
	XmlCharEncodingXMLCHARENCODINGUCS2     XmlCharEncoding = 9
	XmlCharEncodingXMLCHARENCODING88591    XmlCharEncoding = 10
	XmlCharEncodingXMLCHARENCODING88592    XmlCharEncoding = 11
	XmlCharEncodingXMLCHARENCODING88593    XmlCharEncoding = 12
	XmlCharEncodingXMLCHARENCODING88594    XmlCharEncoding = 13
	XmlCharEncodingXMLCHARENCODING88595    XmlCharEncoding = 14
	XmlCharEncodingXMLCHARENCODING88596    XmlCharEncoding = 15
	XmlCharEncodingXMLCHARENCODING88597    XmlCharEncoding = 16
	XmlCharEncodingXMLCHARENCODING88598    XmlCharEncoding = 17
	XmlCharEncodingXMLCHARENCODING88599    XmlCharEncoding = 18
	XmlCharEncodingXMLCHARENCODING2022JP   XmlCharEncoding = 19
	XmlCharEncodingXMLCHARENCODINGSHIFTJIS XmlCharEncoding = 20
	XmlCharEncodingXMLCHARENCODINGEUCJP    XmlCharEncoding = 21
	XmlCharEncodingXMLCHARENCODINGASCII    XmlCharEncoding = 22
)
// llgo:type C
type XmlCharEncodingInputFunc func(*int8, *c.Int, *int8, *c.Int) c.Int
// llgo:type C
type XmlCharEncodingOutputFunc func(*int8, *c.Int, *int8, *c.Int) c.Int

type X_XmlCharEncodingHandler struct {
	Unused [8]uint8
}
type XmlCharEncodingHandler X_XmlCharEncodingHandler
type XmlCharEncodingHandlerPtr *XmlCharEncodingHandler
//go:linkname XmlInitCharEncodingHandlers C.xmlInitCharEncodingHandlers
func XmlInitCharEncodingHandlers()
//go:linkname XmlCleanupCharEncodingHandlers C.xmlCleanupCharEncodingHandlers
func XmlCleanupCharEncodingHandlers()
//go:linkname XmlRegisterCharEncodingHandler C.xmlRegisterCharEncodingHandler
func XmlRegisterCharEncodingHandler(handler XmlCharEncodingHandlerPtr)
// llgo:link XmlCharEncoding.XmlLookupCharEncodingHandler C.xmlLookupCharEncodingHandler
func (recv_ XmlCharEncoding) XmlLookupCharEncodingHandler(out *XmlCharEncodingHandlerPtr) c.Int {
	return 0
}
//go:linkname XmlOpenCharEncodingHandler C.xmlOpenCharEncodingHandler
func XmlOpenCharEncodingHandler(name *int8, output c.Int, out *XmlCharEncodingHandlerPtr) c.Int
// llgo:link XmlCharEncoding.XmlGetCharEncodingHandler C.xmlGetCharEncodingHandler
func (recv_ XmlCharEncoding) XmlGetCharEncodingHandler() XmlCharEncodingHandlerPtr {
	return nil
}
//go:linkname XmlFindCharEncodingHandler C.xmlFindCharEncodingHandler
func XmlFindCharEncodingHandler(name *int8) XmlCharEncodingHandlerPtr
//go:linkname XmlNewCharEncodingHandler C.xmlNewCharEncodingHandler
func XmlNewCharEncodingHandler(name *int8, input XmlCharEncodingInputFunc, output XmlCharEncodingOutputFunc) XmlCharEncodingHandlerPtr
//go:linkname XmlAddEncodingAlias C.xmlAddEncodingAlias
func XmlAddEncodingAlias(name *int8, alias *int8) c.Int
//go:linkname XmlDelEncodingAlias C.xmlDelEncodingAlias
func XmlDelEncodingAlias(alias *int8) c.Int
//go:linkname XmlGetEncodingAlias C.xmlGetEncodingAlias
func XmlGetEncodingAlias(alias *int8) *int8
//go:linkname XmlCleanupEncodingAliases C.xmlCleanupEncodingAliases
func XmlCleanupEncodingAliases()
//go:linkname XmlParseCharEncoding C.xmlParseCharEncoding
func XmlParseCharEncoding(name *int8) XmlCharEncoding
// llgo:link XmlCharEncoding.XmlGetCharEncodingName C.xmlGetCharEncodingName
func (recv_ XmlCharEncoding) XmlGetCharEncodingName() *int8 {
	return nil
}
//go:linkname XmlDetectCharEncoding C.xmlDetectCharEncoding
func XmlDetectCharEncoding(in *int8, len c.Int) XmlCharEncoding
/** DOC_ENABLE */
// llgo:link (*XmlCharEncodingHandler).XmlCharEncOutFunc C.xmlCharEncOutFunc
func (recv_ *XmlCharEncodingHandler) XmlCharEncOutFunc(out *X_XmlBuffer, in *X_XmlBuffer) c.Int {
	return 0
}
// llgo:link (*XmlCharEncodingHandler).XmlCharEncInFunc C.xmlCharEncInFunc
func (recv_ *XmlCharEncodingHandler) XmlCharEncInFunc(out *X_XmlBuffer, in *X_XmlBuffer) c.Int {
	return 0
}
// llgo:link (*XmlCharEncodingHandler).XmlCharEncFirstLine C.xmlCharEncFirstLine
func (recv_ *XmlCharEncodingHandler) XmlCharEncFirstLine(out *X_XmlBuffer, in *X_XmlBuffer) c.Int {
	return 0
}
// llgo:link (*XmlCharEncodingHandler).XmlCharEncCloseFunc C.xmlCharEncCloseFunc
func (recv_ *XmlCharEncodingHandler) XmlCharEncCloseFunc() c.Int {
	return 0
}
//go:linkname UTF8Toisolat1 C.UTF8Toisolat1
func UTF8Toisolat1(out *int8, outlen *c.Int, in *int8, inlen *c.Int) c.Int
//go:linkname Isolat1ToUTF8 C.isolat1ToUTF8
func Isolat1ToUTF8(out *int8, outlen *c.Int, in *int8, inlen *c.Int) c.Int
