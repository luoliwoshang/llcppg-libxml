package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type XmlSchemaWhitespaceValueType c.Int

const (
	XmlSchemaWhitespaceValueTypeXMLSCHEMAWHITESPACEUNKNOWN  XmlSchemaWhitespaceValueType = 0
	XmlSchemaWhitespaceValueTypeXMLSCHEMAWHITESPACEPRESERVE XmlSchemaWhitespaceValueType = 1
	XmlSchemaWhitespaceValueTypeXMLSCHEMAWHITESPACEREPLACE  XmlSchemaWhitespaceValueType = 2
	XmlSchemaWhitespaceValueTypeXMLSCHEMAWHITESPACECOLLAPSE XmlSchemaWhitespaceValueType = 3
)
//go:linkname XmlSchemaInitTypes C.xmlSchemaInitTypes
func XmlSchemaInitTypes() c.Int
//go:linkname XmlSchemaCleanupTypes C.xmlSchemaCleanupTypes
func XmlSchemaCleanupTypes()
// llgo:link (*XmlChar).XmlSchemaGetPredefinedType C.xmlSchemaGetPredefinedType
func (recv_ *XmlChar) XmlSchemaGetPredefinedType(ns *XmlChar) XmlSchemaTypePtr {
	return nil
}
//go:linkname XmlSchemaValidatePredefinedType C.xmlSchemaValidatePredefinedType
func XmlSchemaValidatePredefinedType(type_ XmlSchemaTypePtr, value *XmlChar, val *XmlSchemaValPtr) c.Int
//go:linkname XmlSchemaValPredefTypeNode C.xmlSchemaValPredefTypeNode
func XmlSchemaValPredefTypeNode(type_ XmlSchemaTypePtr, value *XmlChar, val *XmlSchemaValPtr, node XmlNodePtr) c.Int
//go:linkname XmlSchemaValidateFacet C.xmlSchemaValidateFacet
func XmlSchemaValidateFacet(base XmlSchemaTypePtr, facet XmlSchemaFacetPtr, value *XmlChar, val XmlSchemaValPtr) c.Int
//go:linkname XmlSchemaValidateFacetWhtsp C.xmlSchemaValidateFacetWhtsp
func XmlSchemaValidateFacetWhtsp(facet XmlSchemaFacetPtr, fws XmlSchemaWhitespaceValueType, valType XmlSchemaValType, value *XmlChar, val XmlSchemaValPtr, ws XmlSchemaWhitespaceValueType) c.Int
//go:linkname XmlSchemaFreeValue C.xmlSchemaFreeValue
func XmlSchemaFreeValue(val XmlSchemaValPtr)
//go:linkname XmlSchemaNewFacet C.xmlSchemaNewFacet
func XmlSchemaNewFacet() XmlSchemaFacetPtr
//go:linkname XmlSchemaCheckFacet C.xmlSchemaCheckFacet
func XmlSchemaCheckFacet(facet XmlSchemaFacetPtr, typeDecl XmlSchemaTypePtr, ctxt XmlSchemaParserCtxtPtr, name *XmlChar) c.Int
//go:linkname XmlSchemaFreeFacet C.xmlSchemaFreeFacet
func XmlSchemaFreeFacet(facet XmlSchemaFacetPtr)
//go:linkname XmlSchemaCompareValues C.xmlSchemaCompareValues
func XmlSchemaCompareValues(x XmlSchemaValPtr, y XmlSchemaValPtr) c.Int
//go:linkname XmlSchemaGetBuiltInListSimpleTypeItemType C.xmlSchemaGetBuiltInListSimpleTypeItemType
func XmlSchemaGetBuiltInListSimpleTypeItemType(type_ XmlSchemaTypePtr) XmlSchemaTypePtr
//go:linkname XmlSchemaValidateListSimpleTypeFacet C.xmlSchemaValidateListSimpleTypeFacet
func XmlSchemaValidateListSimpleTypeFacet(facet XmlSchemaFacetPtr, value *XmlChar, actualLen c.Ulong, expectedLen *c.Ulong) c.Int
// llgo:link XmlSchemaValType.XmlSchemaGetBuiltInType C.xmlSchemaGetBuiltInType
func (recv_ XmlSchemaValType) XmlSchemaGetBuiltInType() XmlSchemaTypePtr {
	return nil
}
//go:linkname XmlSchemaIsBuiltInTypeFacet C.xmlSchemaIsBuiltInTypeFacet
func XmlSchemaIsBuiltInTypeFacet(type_ XmlSchemaTypePtr, facetType c.Int) c.Int
// llgo:link (*XmlChar).XmlSchemaCollapseString C.xmlSchemaCollapseString
func (recv_ *XmlChar) XmlSchemaCollapseString() *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlSchemaWhiteSpaceReplace C.xmlSchemaWhiteSpaceReplace
func (recv_ *XmlChar) XmlSchemaWhiteSpaceReplace() *XmlChar {
	return nil
}
//go:linkname XmlSchemaGetFacetValueAsULong C.xmlSchemaGetFacetValueAsULong
func XmlSchemaGetFacetValueAsULong(facet XmlSchemaFacetPtr) c.Ulong
//go:linkname XmlSchemaValidateLengthFacet C.xmlSchemaValidateLengthFacet
func XmlSchemaValidateLengthFacet(type_ XmlSchemaTypePtr, facet XmlSchemaFacetPtr, value *XmlChar, val XmlSchemaValPtr, length *c.Ulong) c.Int
//go:linkname XmlSchemaValidateLengthFacetWhtsp C.xmlSchemaValidateLengthFacetWhtsp
func XmlSchemaValidateLengthFacetWhtsp(facet XmlSchemaFacetPtr, valType XmlSchemaValType, value *XmlChar, val XmlSchemaValPtr, length *c.Ulong, ws XmlSchemaWhitespaceValueType) c.Int
//go:linkname XmlSchemaValPredefTypeNodeNoNorm C.xmlSchemaValPredefTypeNodeNoNorm
func XmlSchemaValPredefTypeNodeNoNorm(type_ XmlSchemaTypePtr, value *XmlChar, val *XmlSchemaValPtr, node XmlNodePtr) c.Int
//go:linkname XmlSchemaGetCanonValue C.xmlSchemaGetCanonValue
func XmlSchemaGetCanonValue(val XmlSchemaValPtr, retValue **XmlChar) c.Int
//go:linkname XmlSchemaGetCanonValueWhtsp C.xmlSchemaGetCanonValueWhtsp
func XmlSchemaGetCanonValueWhtsp(val XmlSchemaValPtr, retValue **XmlChar, ws XmlSchemaWhitespaceValueType) c.Int
//go:linkname XmlSchemaValueAppend C.xmlSchemaValueAppend
func XmlSchemaValueAppend(prev XmlSchemaValPtr, cur XmlSchemaValPtr) c.Int
//go:linkname XmlSchemaValueGetNext C.xmlSchemaValueGetNext
func XmlSchemaValueGetNext(cur XmlSchemaValPtr) XmlSchemaValPtr
//go:linkname XmlSchemaValueGetAsString C.xmlSchemaValueGetAsString
func XmlSchemaValueGetAsString(val XmlSchemaValPtr) *XmlChar
//go:linkname XmlSchemaValueGetAsBoolean C.xmlSchemaValueGetAsBoolean
func XmlSchemaValueGetAsBoolean(val XmlSchemaValPtr) c.Int
// llgo:link XmlSchemaValType.XmlSchemaNewStringValue C.xmlSchemaNewStringValue
func (recv_ XmlSchemaValType) XmlSchemaNewStringValue(value *XmlChar) XmlSchemaValPtr {
	return nil
}
// llgo:link (*XmlChar).XmlSchemaNewNOTATIONValue C.xmlSchemaNewNOTATIONValue
func (recv_ *XmlChar) XmlSchemaNewNOTATIONValue(ns *XmlChar) XmlSchemaValPtr {
	return nil
}
// llgo:link (*XmlChar).XmlSchemaNewQNameValue C.xmlSchemaNewQNameValue
func (recv_ *XmlChar) XmlSchemaNewQNameValue(localName *XmlChar) XmlSchemaValPtr {
	return nil
}
//go:linkname XmlSchemaCompareValuesWhtsp C.xmlSchemaCompareValuesWhtsp
func XmlSchemaCompareValuesWhtsp(x XmlSchemaValPtr, xws XmlSchemaWhitespaceValueType, y XmlSchemaValPtr, yws XmlSchemaWhitespaceValueType) c.Int
//go:linkname XmlSchemaCopyValue C.xmlSchemaCopyValue
func XmlSchemaCopyValue(val XmlSchemaValPtr) XmlSchemaValPtr
//go:linkname XmlSchemaGetValType C.xmlSchemaGetValType
func XmlSchemaGetValType(val XmlSchemaValPtr) XmlSchemaValType
