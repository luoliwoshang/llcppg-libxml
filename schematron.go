package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XmlSchematronValidOptions c.Int

const (
	XmlSchematronValidOptionsXMLSCHEMATRONOUTQUIET  XmlSchematronValidOptions = 1
	XmlSchematronValidOptionsXMLSCHEMATRONOUTTEXT   XmlSchematronValidOptions = 2
	XmlSchematronValidOptionsXMLSCHEMATRONOUTXML    XmlSchematronValidOptions = 4
	XmlSchematronValidOptionsXMLSCHEMATRONOUTERROR  XmlSchematronValidOptions = 8
	XmlSchematronValidOptionsXMLSCHEMATRONOUTFILE   XmlSchematronValidOptions = 256
	XmlSchematronValidOptionsXMLSCHEMATRONOUTBUFFER XmlSchematronValidOptions = 512
	XmlSchematronValidOptionsXMLSCHEMATRONOUTIO     XmlSchematronValidOptions = 1024
)

type X_XmlSchematron struct {
	Unused [8]uint8
}
type XmlSchematron X_XmlSchematron
type XmlSchematronPtr *XmlSchematron
// llgo:type C
type XmlSchematronValidityErrorFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})
// llgo:type C
type XmlSchematronValidityWarningFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

type X_XmlSchematronParserCtxt struct {
	Unused [8]uint8
}
type XmlSchematronParserCtxt X_XmlSchematronParserCtxt
type XmlSchematronParserCtxtPtr *XmlSchematronParserCtxt

type X_XmlSchematronValidCtxt struct {
	Unused [8]uint8
}
type XmlSchematronValidCtxt X_XmlSchematronValidCtxt
type XmlSchematronValidCtxtPtr *XmlSchematronValidCtxt
//go:linkname XmlSchematronNewParserCtxt C.xmlSchematronNewParserCtxt
func XmlSchematronNewParserCtxt(URL *int8) XmlSchematronParserCtxtPtr
//go:linkname XmlSchematronNewMemParserCtxt C.xmlSchematronNewMemParserCtxt
func XmlSchematronNewMemParserCtxt(buffer *int8, size c.Int) XmlSchematronParserCtxtPtr
//go:linkname XmlSchematronNewDocParserCtxt C.xmlSchematronNewDocParserCtxt
func XmlSchematronNewDocParserCtxt(doc XmlDocPtr) XmlSchematronParserCtxtPtr
//go:linkname XmlSchematronFreeParserCtxt C.xmlSchematronFreeParserCtxt
func XmlSchematronFreeParserCtxt(ctxt XmlSchematronParserCtxtPtr)
/*****
XMLPUBFUN void
	    xmlSchematronSetParserErrors(xmlSchematronParserCtxtPtr ctxt,
					 xmlSchematronValidityErrorFunc err,
					 xmlSchematronValidityWarningFunc warn,
					 void *ctx);
XMLPUBFUN int
		xmlSchematronGetParserErrors(xmlSchematronParserCtxtPtr ctxt,
					xmlSchematronValidityErrorFunc * err,
					xmlSchematronValidityWarningFunc * warn,
					void **ctx);
XMLPUBFUN int
		xmlSchematronIsValid	(xmlSchematronValidCtxtPtr ctxt);
 *****/
//go:linkname XmlSchematronParse C.xmlSchematronParse
func XmlSchematronParse(ctxt XmlSchematronParserCtxtPtr) XmlSchematronPtr
//go:linkname XmlSchematronFree C.xmlSchematronFree
func XmlSchematronFree(schema XmlSchematronPtr)
//go:linkname XmlSchematronSetValidStructuredErrors C.xmlSchematronSetValidStructuredErrors
func XmlSchematronSetValidStructuredErrors(ctxt XmlSchematronValidCtxtPtr, serror XmlStructuredErrorFunc, ctx unsafe.Pointer)
/******
XMLPUBFUN void
	    xmlSchematronSetValidErrors	(xmlSchematronValidCtxtPtr ctxt,
					 xmlSchematronValidityErrorFunc err,
					 xmlSchematronValidityWarningFunc warn,
					 void *ctx);
XMLPUBFUN int
	    xmlSchematronGetValidErrors	(xmlSchematronValidCtxtPtr ctxt,
					 xmlSchematronValidityErrorFunc *err,
					 xmlSchematronValidityWarningFunc *warn,
					 void **ctx);
XMLPUBFUN int
	    xmlSchematronSetValidOptions(xmlSchematronValidCtxtPtr ctxt,
					 int options);
XMLPUBFUN int
	    xmlSchematronValidCtxtGetOptions(xmlSchematronValidCtxtPtr ctxt);
XMLPUBFUN int
            xmlSchematronValidateOneElement (xmlSchematronValidCtxtPtr ctxt,
			                 xmlNodePtr elem);
 *******/
//go:linkname XmlSchematronNewValidCtxt C.xmlSchematronNewValidCtxt
func XmlSchematronNewValidCtxt(schema XmlSchematronPtr, options c.Int) XmlSchematronValidCtxtPtr
//go:linkname XmlSchematronFreeValidCtxt C.xmlSchematronFreeValidCtxt
func XmlSchematronFreeValidCtxt(ctxt XmlSchematronValidCtxtPtr)
//go:linkname XmlSchematronValidateDoc C.xmlSchematronValidateDoc
func XmlSchematronValidateDoc(ctxt XmlSchematronValidCtxtPtr, instance XmlDocPtr) c.Int
