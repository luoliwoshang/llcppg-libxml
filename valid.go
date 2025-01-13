package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_XmlValidState struct {
	Unused [8]uint8
}
type XmlValidState X_XmlValidState
type XmlValidStatePtr *XmlValidState
// llgo:type C
type XmlValidityErrorFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})
// llgo:type C
type XmlValidityWarningFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

type X_XmlValidCtxt struct {
	UserData  unsafe.Pointer
	Error     unsafe.Pointer
	Warning   unsafe.Pointer
	Node      XmlNodePtr
	NodeNr    c.Int
	NodeMax   c.Int
	NodeTab   *XmlNodePtr
	Flags     c.Uint
	Doc       XmlDocPtr
	Valid     c.Int
	Vstate    *XmlValidState
	VstateNr  c.Int
	VstateMax c.Int
	VstateTab *XmlValidState
	Am        XmlAutomataPtr
	State     XmlAutomataStatePtr
}
type XmlValidCtxt X_XmlValidCtxt
type XmlValidCtxtPtr *XmlValidCtxt
type XmlNotationTable X_XmlHashTable
type XmlNotationTablePtr *XmlNotationTable
type XmlElementTable X_XmlHashTable
type XmlElementTablePtr *XmlElementTable
type XmlAttributeTable X_XmlHashTable
type XmlAttributeTablePtr *XmlAttributeTable
type XmlIDTable X_XmlHashTable
type XmlIDTablePtr *XmlIDTable
type XmlRefTable X_XmlHashTable
type XmlRefTablePtr *XmlRefTable
//go:linkname XmlAddNotationDecl C.xmlAddNotationDecl
func XmlAddNotationDecl(ctxt XmlValidCtxtPtr, dtd XmlDtdPtr, name *XmlChar, PublicID *XmlChar, SystemID *XmlChar) XmlNotationPtr
//go:linkname XmlCopyNotationTable C.xmlCopyNotationTable
func XmlCopyNotationTable(table XmlNotationTablePtr) XmlNotationTablePtr
//go:linkname XmlFreeNotationTable C.xmlFreeNotationTable
func XmlFreeNotationTable(table XmlNotationTablePtr)
//go:linkname XmlDumpNotationDecl C.xmlDumpNotationDecl
func XmlDumpNotationDecl(buf XmlBufferPtr, nota XmlNotationPtr)
//go:linkname XmlDumpNotationTable C.xmlDumpNotationTable
func XmlDumpNotationTable(buf XmlBufferPtr, table XmlNotationTablePtr)
// llgo:link (*XmlChar).XmlNewElementContent C.xmlNewElementContent
func (recv_ *XmlChar) XmlNewElementContent(type_ XmlElementContentType) XmlElementContentPtr {
	return nil
}
//go:linkname XmlCopyElementContent C.xmlCopyElementContent
func XmlCopyElementContent(content XmlElementContentPtr) XmlElementContentPtr
//go:linkname XmlFreeElementContent C.xmlFreeElementContent
func XmlFreeElementContent(cur XmlElementContentPtr)
//go:linkname XmlNewDocElementContent C.xmlNewDocElementContent
func XmlNewDocElementContent(doc XmlDocPtr, name *XmlChar, type_ XmlElementContentType) XmlElementContentPtr
//go:linkname XmlCopyDocElementContent C.xmlCopyDocElementContent
func XmlCopyDocElementContent(doc XmlDocPtr, content XmlElementContentPtr) XmlElementContentPtr
//go:linkname XmlFreeDocElementContent C.xmlFreeDocElementContent
func XmlFreeDocElementContent(doc XmlDocPtr, cur XmlElementContentPtr)
//go:linkname XmlSnprintfElementContent C.xmlSnprintfElementContent
func XmlSnprintfElementContent(buf *int8, size c.Int, content XmlElementContentPtr, englob c.Int)
//go:linkname XmlSprintfElementContent C.xmlSprintfElementContent
func XmlSprintfElementContent(buf *int8, content XmlElementContentPtr, englob c.Int)
//go:linkname XmlAddElementDecl C.xmlAddElementDecl
func XmlAddElementDecl(ctxt XmlValidCtxtPtr, dtd XmlDtdPtr, name *XmlChar, type_ XmlElementTypeVal, content XmlElementContentPtr) XmlElementPtr
//go:linkname XmlCopyElementTable C.xmlCopyElementTable
func XmlCopyElementTable(table XmlElementTablePtr) XmlElementTablePtr
//go:linkname XmlFreeElementTable C.xmlFreeElementTable
func XmlFreeElementTable(table XmlElementTablePtr)
//go:linkname XmlDumpElementTable C.xmlDumpElementTable
func XmlDumpElementTable(buf XmlBufferPtr, table XmlElementTablePtr)
//go:linkname XmlDumpElementDecl C.xmlDumpElementDecl
func XmlDumpElementDecl(buf XmlBufferPtr, elem XmlElementPtr)
// llgo:link (*XmlChar).XmlCreateEnumeration C.xmlCreateEnumeration
func (recv_ *XmlChar) XmlCreateEnumeration() XmlEnumerationPtr {
	return nil
}
//go:linkname XmlFreeEnumeration C.xmlFreeEnumeration
func XmlFreeEnumeration(cur XmlEnumerationPtr)
//go:linkname XmlCopyEnumeration C.xmlCopyEnumeration
func XmlCopyEnumeration(cur XmlEnumerationPtr) XmlEnumerationPtr
//go:linkname XmlAddAttributeDecl C.xmlAddAttributeDecl
func XmlAddAttributeDecl(ctxt XmlValidCtxtPtr, dtd XmlDtdPtr, elem *XmlChar, name *XmlChar, ns *XmlChar, type_ XmlAttributeType, def XmlAttributeDefault, defaultValue *XmlChar, tree XmlEnumerationPtr) XmlAttributePtr
//go:linkname XmlCopyAttributeTable C.xmlCopyAttributeTable
func XmlCopyAttributeTable(table XmlAttributeTablePtr) XmlAttributeTablePtr
//go:linkname XmlFreeAttributeTable C.xmlFreeAttributeTable
func XmlFreeAttributeTable(table XmlAttributeTablePtr)
//go:linkname XmlDumpAttributeTable C.xmlDumpAttributeTable
func XmlDumpAttributeTable(buf XmlBufferPtr, table XmlAttributeTablePtr)
//go:linkname XmlDumpAttributeDecl C.xmlDumpAttributeDecl
func XmlDumpAttributeDecl(buf XmlBufferPtr, attr XmlAttributePtr)
//go:linkname XmlAddIDSafe C.xmlAddIDSafe
func XmlAddIDSafe(attr XmlAttrPtr, value *XmlChar) c.Int
//go:linkname XmlAddID C.xmlAddID
func XmlAddID(ctxt XmlValidCtxtPtr, doc XmlDocPtr, value *XmlChar, attr XmlAttrPtr) XmlIDPtr
//go:linkname XmlFreeIDTable C.xmlFreeIDTable
func XmlFreeIDTable(table XmlIDTablePtr)
//go:linkname XmlGetID C.xmlGetID
func XmlGetID(doc XmlDocPtr, ID *XmlChar) XmlAttrPtr
//go:linkname XmlIsID C.xmlIsID
func XmlIsID(doc XmlDocPtr, elem XmlNodePtr, attr XmlAttrPtr) c.Int
//go:linkname XmlRemoveID C.xmlRemoveID
func XmlRemoveID(doc XmlDocPtr, attr XmlAttrPtr) c.Int
//go:linkname XmlAddRef C.xmlAddRef
func XmlAddRef(ctxt XmlValidCtxtPtr, doc XmlDocPtr, value *XmlChar, attr XmlAttrPtr) XmlRefPtr
//go:linkname XmlFreeRefTable C.xmlFreeRefTable
func XmlFreeRefTable(table XmlRefTablePtr)
//go:linkname XmlIsRef C.xmlIsRef
func XmlIsRef(doc XmlDocPtr, elem XmlNodePtr, attr XmlAttrPtr) c.Int
//go:linkname XmlRemoveRef C.xmlRemoveRef
func XmlRemoveRef(doc XmlDocPtr, attr XmlAttrPtr) c.Int
//go:linkname XmlGetRefs C.xmlGetRefs
func XmlGetRefs(doc XmlDocPtr, ID *XmlChar) XmlListPtr
//go:linkname XmlNewValidCtxt C.xmlNewValidCtxt
func XmlNewValidCtxt() XmlValidCtxtPtr
//go:linkname XmlFreeValidCtxt C.xmlFreeValidCtxt
func XmlFreeValidCtxt(XmlValidCtxtPtr)
//go:linkname XmlValidateRoot C.xmlValidateRoot
func XmlValidateRoot(ctxt XmlValidCtxtPtr, doc XmlDocPtr) c.Int
//go:linkname XmlValidateElementDecl C.xmlValidateElementDecl
func XmlValidateElementDecl(ctxt XmlValidCtxtPtr, doc XmlDocPtr, elem XmlElementPtr) c.Int
//go:linkname XmlValidNormalizeAttributeValue C.xmlValidNormalizeAttributeValue
func XmlValidNormalizeAttributeValue(doc XmlDocPtr, elem XmlNodePtr, name *XmlChar, value *XmlChar) *XmlChar
//go:linkname XmlValidCtxtNormalizeAttributeValue C.xmlValidCtxtNormalizeAttributeValue
func XmlValidCtxtNormalizeAttributeValue(ctxt XmlValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr, name *XmlChar, value *XmlChar) *XmlChar
//go:linkname XmlValidateAttributeDecl C.xmlValidateAttributeDecl
func XmlValidateAttributeDecl(ctxt XmlValidCtxtPtr, doc XmlDocPtr, attr XmlAttributePtr) c.Int
// llgo:link XmlAttributeType.XmlValidateAttributeValue C.xmlValidateAttributeValue
func (recv_ XmlAttributeType) XmlValidateAttributeValue(value *XmlChar) c.Int {
	return 0
}
//go:linkname XmlValidateNotationDecl C.xmlValidateNotationDecl
func XmlValidateNotationDecl(ctxt XmlValidCtxtPtr, doc XmlDocPtr, nota XmlNotationPtr) c.Int
//go:linkname XmlValidateDtd C.xmlValidateDtd
func XmlValidateDtd(ctxt XmlValidCtxtPtr, doc XmlDocPtr, dtd XmlDtdPtr) c.Int
//go:linkname XmlValidateDtdFinal C.xmlValidateDtdFinal
func XmlValidateDtdFinal(ctxt XmlValidCtxtPtr, doc XmlDocPtr) c.Int
//go:linkname XmlValidateDocument C.xmlValidateDocument
func XmlValidateDocument(ctxt XmlValidCtxtPtr, doc XmlDocPtr) c.Int
//go:linkname XmlValidateElement C.xmlValidateElement
func XmlValidateElement(ctxt XmlValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr) c.Int
//go:linkname XmlValidateOneElement C.xmlValidateOneElement
func XmlValidateOneElement(ctxt XmlValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr) c.Int
//go:linkname XmlValidateOneAttribute C.xmlValidateOneAttribute
func XmlValidateOneAttribute(ctxt XmlValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr, attr XmlAttrPtr, value *XmlChar) c.Int
//go:linkname XmlValidateOneNamespace C.xmlValidateOneNamespace
func XmlValidateOneNamespace(ctxt XmlValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr, prefix *XmlChar, ns XmlNsPtr, value *XmlChar) c.Int
//go:linkname XmlValidateDocumentFinal C.xmlValidateDocumentFinal
func XmlValidateDocumentFinal(ctxt XmlValidCtxtPtr, doc XmlDocPtr) c.Int
//go:linkname XmlValidateNotationUse C.xmlValidateNotationUse
func XmlValidateNotationUse(ctxt XmlValidCtxtPtr, doc XmlDocPtr, notationName *XmlChar) c.Int
//go:linkname XmlIsMixedElement C.xmlIsMixedElement
func XmlIsMixedElement(doc XmlDocPtr, name *XmlChar) c.Int
//go:linkname XmlGetDtdAttrDesc C.xmlGetDtdAttrDesc
func XmlGetDtdAttrDesc(dtd XmlDtdPtr, elem *XmlChar, name *XmlChar) XmlAttributePtr
//go:linkname XmlGetDtdQAttrDesc C.xmlGetDtdQAttrDesc
func XmlGetDtdQAttrDesc(dtd XmlDtdPtr, elem *XmlChar, name *XmlChar, prefix *XmlChar) XmlAttributePtr
//go:linkname XmlGetDtdNotationDesc C.xmlGetDtdNotationDesc
func XmlGetDtdNotationDesc(dtd XmlDtdPtr, name *XmlChar) XmlNotationPtr
//go:linkname XmlGetDtdQElementDesc C.xmlGetDtdQElementDesc
func XmlGetDtdQElementDesc(dtd XmlDtdPtr, name *XmlChar, prefix *XmlChar) XmlElementPtr
//go:linkname XmlGetDtdElementDesc C.xmlGetDtdElementDesc
func XmlGetDtdElementDesc(dtd XmlDtdPtr, name *XmlChar) XmlElementPtr
// llgo:link (*XmlElementContent).XmlValidGetPotentialChildren C.xmlValidGetPotentialChildren
func (recv_ *XmlElementContent) XmlValidGetPotentialChildren(names **XmlChar, len *c.Int, max c.Int) c.Int {
	return 0
}
// llgo:link (*XmlNode).XmlValidGetValidElements C.xmlValidGetValidElements
func (recv_ *XmlNode) XmlValidGetValidElements(next *XmlNode, names **XmlChar, max c.Int) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlValidateNameValue C.xmlValidateNameValue
func (recv_ *XmlChar) XmlValidateNameValue() c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlValidateNamesValue C.xmlValidateNamesValue
func (recv_ *XmlChar) XmlValidateNamesValue() c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlValidateNmtokenValue C.xmlValidateNmtokenValue
func (recv_ *XmlChar) XmlValidateNmtokenValue() c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlValidateNmtokensValue C.xmlValidateNmtokensValue
func (recv_ *XmlChar) XmlValidateNmtokensValue() c.Int {
	return 0
}
//go:linkname XmlValidBuildContentModel C.xmlValidBuildContentModel
func XmlValidBuildContentModel(ctxt XmlValidCtxtPtr, elem XmlElementPtr) c.Int
//go:linkname XmlValidatePushElement C.xmlValidatePushElement
func XmlValidatePushElement(ctxt XmlValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr, qname *XmlChar) c.Int
//go:linkname XmlValidatePushCData C.xmlValidatePushCData
func XmlValidatePushCData(ctxt XmlValidCtxtPtr, data *XmlChar, len c.Int) c.Int
//go:linkname XmlValidatePopElement C.xmlValidatePopElement
func XmlValidatePopElement(ctxt XmlValidCtxtPtr, doc XmlDocPtr, elem XmlNodePtr, qname *XmlChar) c.Int
