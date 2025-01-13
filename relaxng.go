package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_XmlRelaxNG struct {
	Unused [8]uint8
}
type XmlRelaxNG X_XmlRelaxNG
type XmlRelaxNGPtr *XmlRelaxNG
// llgo:type C
type XmlRelaxNGValidityErrorFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})
// llgo:type C
type XmlRelaxNGValidityWarningFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

type X_XmlRelaxNGParserCtxt struct {
	Unused [8]uint8
}
type XmlRelaxNGParserCtxt X_XmlRelaxNGParserCtxt
type XmlRelaxNGParserCtxtPtr *XmlRelaxNGParserCtxt

type X_XmlRelaxNGValidCtxt struct {
	Unused [8]uint8
}
type XmlRelaxNGValidCtxt X_XmlRelaxNGValidCtxt
type XmlRelaxNGValidCtxtPtr *XmlRelaxNGValidCtxt
type XmlRelaxNGValidErr c.Int

const (
	XmlRelaxNGValidErrXMLRELAXNGOK              XmlRelaxNGValidErr = 0
	XmlRelaxNGValidErrXMLRELAXNGERRMEMORY       XmlRelaxNGValidErr = 1
	XmlRelaxNGValidErrXMLRELAXNGERRTYPE         XmlRelaxNGValidErr = 2
	XmlRelaxNGValidErrXMLRELAXNGERRTYPEVAL      XmlRelaxNGValidErr = 3
	XmlRelaxNGValidErrXMLRELAXNGERRDUPID        XmlRelaxNGValidErr = 4
	XmlRelaxNGValidErrXMLRELAXNGERRTYPECMP      XmlRelaxNGValidErr = 5
	XmlRelaxNGValidErrXMLRELAXNGERRNOSTATE      XmlRelaxNGValidErr = 6
	XmlRelaxNGValidErrXMLRELAXNGERRNODEFINE     XmlRelaxNGValidErr = 7
	XmlRelaxNGValidErrXMLRELAXNGERRLISTEXTRA    XmlRelaxNGValidErr = 8
	XmlRelaxNGValidErrXMLRELAXNGERRLISTEMPTY    XmlRelaxNGValidErr = 9
	XmlRelaxNGValidErrXMLRELAXNGERRINTERNODATA  XmlRelaxNGValidErr = 10
	XmlRelaxNGValidErrXMLRELAXNGERRINTERSEQ     XmlRelaxNGValidErr = 11
	XmlRelaxNGValidErrXMLRELAXNGERRINTEREXTRA   XmlRelaxNGValidErr = 12
	XmlRelaxNGValidErrXMLRELAXNGERRELEMNAME     XmlRelaxNGValidErr = 13
	XmlRelaxNGValidErrXMLRELAXNGERRATTRNAME     XmlRelaxNGValidErr = 14
	XmlRelaxNGValidErrXMLRELAXNGERRELEMNONS     XmlRelaxNGValidErr = 15
	XmlRelaxNGValidErrXMLRELAXNGERRATTRNONS     XmlRelaxNGValidErr = 16
	XmlRelaxNGValidErrXMLRELAXNGERRELEMWRONGNS  XmlRelaxNGValidErr = 17
	XmlRelaxNGValidErrXMLRELAXNGERRATTRWRONGNS  XmlRelaxNGValidErr = 18
	XmlRelaxNGValidErrXMLRELAXNGERRELEMEXTRANS  XmlRelaxNGValidErr = 19
	XmlRelaxNGValidErrXMLRELAXNGERRATTREXTRANS  XmlRelaxNGValidErr = 20
	XmlRelaxNGValidErrXMLRELAXNGERRELEMNOTEMPTY XmlRelaxNGValidErr = 21
	XmlRelaxNGValidErrXMLRELAXNGERRNOELEM       XmlRelaxNGValidErr = 22
	XmlRelaxNGValidErrXMLRELAXNGERRNOTELEM      XmlRelaxNGValidErr = 23
	XmlRelaxNGValidErrXMLRELAXNGERRATTRVALID    XmlRelaxNGValidErr = 24
	XmlRelaxNGValidErrXMLRELAXNGERRCONTENTVALID XmlRelaxNGValidErr = 25
	XmlRelaxNGValidErrXMLRELAXNGERREXTRACONTENT XmlRelaxNGValidErr = 26
	XmlRelaxNGValidErrXMLRELAXNGERRINVALIDATTR  XmlRelaxNGValidErr = 27
	XmlRelaxNGValidErrXMLRELAXNGERRDATAELEM     XmlRelaxNGValidErr = 28
	XmlRelaxNGValidErrXMLRELAXNGERRVALELEM      XmlRelaxNGValidErr = 29
	XmlRelaxNGValidErrXMLRELAXNGERRLISTELEM     XmlRelaxNGValidErr = 30
	XmlRelaxNGValidErrXMLRELAXNGERRDATATYPE     XmlRelaxNGValidErr = 31
	XmlRelaxNGValidErrXMLRELAXNGERRVALUE        XmlRelaxNGValidErr = 32
	XmlRelaxNGValidErrXMLRELAXNGERRLIST         XmlRelaxNGValidErr = 33
	XmlRelaxNGValidErrXMLRELAXNGERRNOGRAMMAR    XmlRelaxNGValidErr = 34
	XmlRelaxNGValidErrXMLRELAXNGERREXTRADATA    XmlRelaxNGValidErr = 35
	XmlRelaxNGValidErrXMLRELAXNGERRLACKDATA     XmlRelaxNGValidErr = 36
	XmlRelaxNGValidErrXMLRELAXNGERRINTERNAL     XmlRelaxNGValidErr = 37
	XmlRelaxNGValidErrXMLRELAXNGERRELEMWRONG    XmlRelaxNGValidErr = 38
	XmlRelaxNGValidErrXMLRELAXNGERRTEXTWRONG    XmlRelaxNGValidErr = 39
)

type XmlRelaxNGParserFlag c.Int

const (
	XmlRelaxNGParserFlagXMLRELAXNGPNONE    XmlRelaxNGParserFlag = 0
	XmlRelaxNGParserFlagXMLRELAXNGPFREEDOC XmlRelaxNGParserFlag = 1
	XmlRelaxNGParserFlagXMLRELAXNGPCRNG    XmlRelaxNGParserFlag = 2
)
//go:linkname XmlRelaxNGInitTypes C.xmlRelaxNGInitTypes
func XmlRelaxNGInitTypes() c.Int
//go:linkname XmlRelaxNGCleanupTypes C.xmlRelaxNGCleanupTypes
func XmlRelaxNGCleanupTypes()
//go:linkname XmlRelaxNGNewParserCtxt C.xmlRelaxNGNewParserCtxt
func XmlRelaxNGNewParserCtxt(URL *int8) XmlRelaxNGParserCtxtPtr
//go:linkname XmlRelaxNGNewMemParserCtxt C.xmlRelaxNGNewMemParserCtxt
func XmlRelaxNGNewMemParserCtxt(buffer *int8, size c.Int) XmlRelaxNGParserCtxtPtr
//go:linkname XmlRelaxNGNewDocParserCtxt C.xmlRelaxNGNewDocParserCtxt
func XmlRelaxNGNewDocParserCtxt(doc XmlDocPtr) XmlRelaxNGParserCtxtPtr
//go:linkname XmlRelaxParserSetFlag C.xmlRelaxParserSetFlag
func XmlRelaxParserSetFlag(ctxt XmlRelaxNGParserCtxtPtr, flag c.Int) c.Int
//go:linkname XmlRelaxNGFreeParserCtxt C.xmlRelaxNGFreeParserCtxt
func XmlRelaxNGFreeParserCtxt(ctxt XmlRelaxNGParserCtxtPtr)
//go:linkname XmlRelaxNGSetParserErrors C.xmlRelaxNGSetParserErrors
func XmlRelaxNGSetParserErrors(ctxt XmlRelaxNGParserCtxtPtr, err XmlRelaxNGValidityErrorFunc, warn XmlRelaxNGValidityWarningFunc, ctx unsafe.Pointer)
//go:linkname XmlRelaxNGGetParserErrors C.xmlRelaxNGGetParserErrors
func XmlRelaxNGGetParserErrors(ctxt XmlRelaxNGParserCtxtPtr, err XmlRelaxNGValidityErrorFunc, warn XmlRelaxNGValidityWarningFunc, ctx *unsafe.Pointer) c.Int
//go:linkname XmlRelaxNGSetParserStructuredErrors C.xmlRelaxNGSetParserStructuredErrors
func XmlRelaxNGSetParserStructuredErrors(ctxt XmlRelaxNGParserCtxtPtr, serror XmlStructuredErrorFunc, ctx unsafe.Pointer)
//go:linkname XmlRelaxNGParse C.xmlRelaxNGParse
func XmlRelaxNGParse(ctxt XmlRelaxNGParserCtxtPtr) XmlRelaxNGPtr
//go:linkname XmlRelaxNGFree C.xmlRelaxNGFree
func XmlRelaxNGFree(schema XmlRelaxNGPtr)
//go:linkname XmlRelaxNGDump C.xmlRelaxNGDump
func XmlRelaxNGDump(output *c.FILE, schema XmlRelaxNGPtr)
//go:linkname XmlRelaxNGDumpTree C.xmlRelaxNGDumpTree
func XmlRelaxNGDumpTree(output *c.FILE, schema XmlRelaxNGPtr)
//go:linkname XmlRelaxNGSetValidErrors C.xmlRelaxNGSetValidErrors
func XmlRelaxNGSetValidErrors(ctxt XmlRelaxNGValidCtxtPtr, err XmlRelaxNGValidityErrorFunc, warn XmlRelaxNGValidityWarningFunc, ctx unsafe.Pointer)
//go:linkname XmlRelaxNGGetValidErrors C.xmlRelaxNGGetValidErrors
func XmlRelaxNGGetValidErrors(ctxt XmlRelaxNGValidCtxtPtr, err XmlRelaxNGValidityErrorFunc, warn XmlRelaxNGValidityWarningFunc, ctx *unsafe.Pointer) c.Int
//go:linkname XmlRelaxNGSetValidStructuredErrors C.xmlRelaxNGSetValidStructuredErrors
func XmlRelaxNGSetValidStructuredErrors(ctxt XmlRelaxNGValidCtxtPtr, serror XmlStructuredErrorFunc, ctx unsafe.Pointer)
//go:linkname XmlRelaxNGNewValidCtxt C.xmlRelaxNGNewValidCtxt
func XmlRelaxNGNewValidCtxt(schema XmlRelaxNGPtr) XmlRelaxNGValidCtxtPtr
//go:linkname XmlRelaxNGFreeValidCtxt C.xmlRelaxNGFreeValidCtxt
func XmlRelaxNGFreeValidCtxt(ctxt XmlRelaxNGValidCtxtPtr)
//go:linkname XmlRelaxNGValidateDoc C.xmlRelaxNGValidateDoc
func XmlRelaxNGValidateDoc(ctxt XmlRelaxNGValidCtxtPtr, doc XmlDocPtr) c.Int
//go:linkname XmlRelaxNGValidatePushElement C.xmlRelaxNGValidatePushElement
func XmlRelaxNGValidatePushElement(ctxt XmlRelaxNGValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr) c.Int
//go:linkname XmlRelaxNGValidatePushCData C.xmlRelaxNGValidatePushCData
func XmlRelaxNGValidatePushCData(ctxt XmlRelaxNGValidCtxtPtr, data *XmlChar, len c.Int) c.Int
//go:linkname XmlRelaxNGValidatePopElement C.xmlRelaxNGValidatePopElement
func XmlRelaxNGValidatePopElement(ctxt XmlRelaxNGValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr) c.Int
//go:linkname XmlRelaxNGValidateFullElement C.xmlRelaxNGValidateFullElement
func XmlRelaxNGValidateFullElement(ctxt XmlRelaxNGValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr) c.Int
