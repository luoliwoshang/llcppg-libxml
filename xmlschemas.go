package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XmlSchemaValidError c.Int

const (
	XmlSchemaValidErrorXMLSCHEMASERROK             XmlSchemaValidError = 0
	XmlSchemaValidErrorXMLSCHEMASERRNOROOT         XmlSchemaValidError = 1
	XmlSchemaValidErrorXMLSCHEMASERRUNDECLAREDELEM XmlSchemaValidError = 2
	XmlSchemaValidErrorXMLSCHEMASERRNOTTOPLEVEL    XmlSchemaValidError = 3
	XmlSchemaValidErrorXMLSCHEMASERRMISSING        XmlSchemaValidError = 4
	XmlSchemaValidErrorXMLSCHEMASERRWRONGELEM      XmlSchemaValidError = 5
	XmlSchemaValidErrorXMLSCHEMASERRNOTYPE         XmlSchemaValidError = 6
	XmlSchemaValidErrorXMLSCHEMASERRNOROLLBACK     XmlSchemaValidError = 7
	XmlSchemaValidErrorXMLSCHEMASERRISABSTRACT     XmlSchemaValidError = 8
	XmlSchemaValidErrorXMLSCHEMASERRNOTEMPTY       XmlSchemaValidError = 9
	XmlSchemaValidErrorXMLSCHEMASERRELEMCONT       XmlSchemaValidError = 10
	XmlSchemaValidErrorXMLSCHEMASERRHAVEDEFAULT    XmlSchemaValidError = 11
	XmlSchemaValidErrorXMLSCHEMASERRNOTNILLABLE    XmlSchemaValidError = 12
	XmlSchemaValidErrorXMLSCHEMASERREXTRACONTENT   XmlSchemaValidError = 13
	XmlSchemaValidErrorXMLSCHEMASERRINVALIDATTR    XmlSchemaValidError = 14
	XmlSchemaValidErrorXMLSCHEMASERRINVALIDELEM    XmlSchemaValidError = 15
	XmlSchemaValidErrorXMLSCHEMASERRNOTDETERMINIST XmlSchemaValidError = 16
	XmlSchemaValidErrorXMLSCHEMASERRCONSTRUCT      XmlSchemaValidError = 17
	XmlSchemaValidErrorXMLSCHEMASERRINTERNAL       XmlSchemaValidError = 18
	XmlSchemaValidErrorXMLSCHEMASERRNOTSIMPLE      XmlSchemaValidError = 19
	XmlSchemaValidErrorXMLSCHEMASERRATTRUNKNOWN    XmlSchemaValidError = 20
	XmlSchemaValidErrorXMLSCHEMASERRATTRINVALID    XmlSchemaValidError = 21
	XmlSchemaValidErrorXMLSCHEMASERRVALUE          XmlSchemaValidError = 22
	XmlSchemaValidErrorXMLSCHEMASERRFACET          XmlSchemaValidError = 23
	XmlSchemaValidErrorXMLSCHEMASERR               XmlSchemaValidError = 24
	XmlSchemaValidErrorXMLSCHEMASERRXXX            XmlSchemaValidError = 25
)

type XmlSchemaValidOption c.Int

const XmlSchemaValidOptionXMLSCHEMAVALVCICREATE XmlSchemaValidOption = 1

type XmlSchema X_XmlSchema
type XmlSchemaPtr *XmlSchema
// llgo:type C
type XmlSchemaValidityErrorFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})
// llgo:type C
type XmlSchemaValidityWarningFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

type X_XmlSchemaParserCtxt struct {
	Unused [8]uint8
}
type XmlSchemaParserCtxt X_XmlSchemaParserCtxt
type XmlSchemaParserCtxtPtr *XmlSchemaParserCtxt

type X_XmlSchemaValidCtxt struct {
	Unused [8]uint8
}
type XmlSchemaValidCtxt X_XmlSchemaValidCtxt
type XmlSchemaValidCtxtPtr *XmlSchemaValidCtxt
// llgo:type C
type XmlSchemaValidityLocatorFunc func(unsafe.Pointer, **int8, *c.Ulong) c.Int
//go:linkname XmlSchemaNewParserCtxt C.xmlSchemaNewParserCtxt
func XmlSchemaNewParserCtxt(URL *int8) XmlSchemaParserCtxtPtr
//go:linkname XmlSchemaNewMemParserCtxt C.xmlSchemaNewMemParserCtxt
func XmlSchemaNewMemParserCtxt(buffer *int8, size c.Int) XmlSchemaParserCtxtPtr
//go:linkname XmlSchemaNewDocParserCtxt C.xmlSchemaNewDocParserCtxt
func XmlSchemaNewDocParserCtxt(doc XmlDocPtr) XmlSchemaParserCtxtPtr
//go:linkname XmlSchemaFreeParserCtxt C.xmlSchemaFreeParserCtxt
func XmlSchemaFreeParserCtxt(ctxt XmlSchemaParserCtxtPtr)
//go:linkname XmlSchemaSetParserErrors C.xmlSchemaSetParserErrors
func XmlSchemaSetParserErrors(ctxt XmlSchemaParserCtxtPtr, err XmlSchemaValidityErrorFunc, warn XmlSchemaValidityWarningFunc, ctx unsafe.Pointer)
//go:linkname XmlSchemaSetParserStructuredErrors C.xmlSchemaSetParserStructuredErrors
func XmlSchemaSetParserStructuredErrors(ctxt XmlSchemaParserCtxtPtr, serror XmlStructuredErrorFunc, ctx unsafe.Pointer)
//go:linkname XmlSchemaGetParserErrors C.xmlSchemaGetParserErrors
func XmlSchemaGetParserErrors(ctxt XmlSchemaParserCtxtPtr, err XmlSchemaValidityErrorFunc, warn XmlSchemaValidityWarningFunc, ctx *unsafe.Pointer) c.Int
//go:linkname XmlSchemaIsValid C.xmlSchemaIsValid
func XmlSchemaIsValid(ctxt XmlSchemaValidCtxtPtr) c.Int
//go:linkname XmlSchemaParse C.xmlSchemaParse
func XmlSchemaParse(ctxt XmlSchemaParserCtxtPtr) XmlSchemaPtr
//go:linkname XmlSchemaFree C.xmlSchemaFree
func XmlSchemaFree(schema XmlSchemaPtr)
//go:linkname XmlSchemaDump C.xmlSchemaDump
func XmlSchemaDump(output *c.FILE, schema XmlSchemaPtr)
//go:linkname XmlSchemaSetValidErrors C.xmlSchemaSetValidErrors
func XmlSchemaSetValidErrors(ctxt XmlSchemaValidCtxtPtr, err XmlSchemaValidityErrorFunc, warn XmlSchemaValidityWarningFunc, ctx unsafe.Pointer)
//go:linkname XmlSchemaSetValidStructuredErrors C.xmlSchemaSetValidStructuredErrors
func XmlSchemaSetValidStructuredErrors(ctxt XmlSchemaValidCtxtPtr, serror XmlStructuredErrorFunc, ctx unsafe.Pointer)
//go:linkname XmlSchemaGetValidErrors C.xmlSchemaGetValidErrors
func XmlSchemaGetValidErrors(ctxt XmlSchemaValidCtxtPtr, err XmlSchemaValidityErrorFunc, warn XmlSchemaValidityWarningFunc, ctx *unsafe.Pointer) c.Int
//go:linkname XmlSchemaSetValidOptions C.xmlSchemaSetValidOptions
func XmlSchemaSetValidOptions(ctxt XmlSchemaValidCtxtPtr, options c.Int) c.Int
//go:linkname XmlSchemaValidateSetFilename C.xmlSchemaValidateSetFilename
func XmlSchemaValidateSetFilename(vctxt XmlSchemaValidCtxtPtr, filename *int8)
//go:linkname XmlSchemaValidCtxtGetOptions C.xmlSchemaValidCtxtGetOptions
func XmlSchemaValidCtxtGetOptions(ctxt XmlSchemaValidCtxtPtr) c.Int
//go:linkname XmlSchemaNewValidCtxt C.xmlSchemaNewValidCtxt
func XmlSchemaNewValidCtxt(schema XmlSchemaPtr) XmlSchemaValidCtxtPtr
//go:linkname XmlSchemaFreeValidCtxt C.xmlSchemaFreeValidCtxt
func XmlSchemaFreeValidCtxt(ctxt XmlSchemaValidCtxtPtr)
//go:linkname XmlSchemaValidateDoc C.xmlSchemaValidateDoc
func XmlSchemaValidateDoc(ctxt XmlSchemaValidCtxtPtr, instance XmlDocPtr) c.Int
//go:linkname XmlSchemaValidateOneElement C.xmlSchemaValidateOneElement
func XmlSchemaValidateOneElement(ctxt XmlSchemaValidCtxtPtr, elem XmlNodePtr) c.Int
//go:linkname XmlSchemaValidateStream C.xmlSchemaValidateStream
func XmlSchemaValidateStream(ctxt XmlSchemaValidCtxtPtr, input XmlParserInputBufferPtr, enc XmlCharEncoding, sax XmlSAXHandlerPtr, user_data unsafe.Pointer) c.Int
//go:linkname XmlSchemaValidateFile C.xmlSchemaValidateFile
func XmlSchemaValidateFile(ctxt XmlSchemaValidCtxtPtr, filename *int8, options c.Int) c.Int
//go:linkname XmlSchemaValidCtxtGetParserCtxt C.xmlSchemaValidCtxtGetParserCtxt
func XmlSchemaValidCtxtGetParserCtxt(ctxt XmlSchemaValidCtxtPtr) XmlParserCtxtPtr

type X_XmlSchemaSAXPlug struct {
	Unused [8]uint8
}
type XmlSchemaSAXPlugStruct X_XmlSchemaSAXPlug
type XmlSchemaSAXPlugPtr *XmlSchemaSAXPlugStruct
//go:linkname XmlSchemaSAXPlug C.xmlSchemaSAXPlug
func XmlSchemaSAXPlug(ctxt XmlSchemaValidCtxtPtr, sax *XmlSAXHandlerPtr, user_data *unsafe.Pointer) XmlSchemaSAXPlugPtr
//go:linkname XmlSchemaSAXUnplug C.xmlSchemaSAXUnplug
func XmlSchemaSAXUnplug(plug XmlSchemaSAXPlugPtr) c.Int
//go:linkname XmlSchemaValidateSetLocator C.xmlSchemaValidateSetLocator
func XmlSchemaValidateSetLocator(vctxt XmlSchemaValidCtxtPtr, f XmlSchemaValidityLocatorFunc, ctxt unsafe.Pointer)
