package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
//go:linkname XmlXPathPopBoolean C.xmlXPathPopBoolean
func XmlXPathPopBoolean(ctxt XmlXPathParserContextPtr) c.Int
//go:linkname XmlXPathPopNumber C.xmlXPathPopNumber
func XmlXPathPopNumber(ctxt XmlXPathParserContextPtr) float64
//go:linkname XmlXPathPopString C.xmlXPathPopString
func XmlXPathPopString(ctxt XmlXPathParserContextPtr) *XmlChar
//go:linkname XmlXPathPopNodeSet C.xmlXPathPopNodeSet
func XmlXPathPopNodeSet(ctxt XmlXPathParserContextPtr) XmlNodeSetPtr
//go:linkname XmlXPathPopExternal C.xmlXPathPopExternal
func XmlXPathPopExternal(ctxt XmlXPathParserContextPtr) unsafe.Pointer
//go:linkname XmlXPathRegisterVariableLookup C.xmlXPathRegisterVariableLookup
func XmlXPathRegisterVariableLookup(ctxt XmlXPathContextPtr, f XmlXPathVariableLookupFunc, data unsafe.Pointer)
//go:linkname XmlXPathRegisterFuncLookup C.xmlXPathRegisterFuncLookup
func XmlXPathRegisterFuncLookup(ctxt XmlXPathContextPtr, f XmlXPathFuncLookupFunc, funcCtxt unsafe.Pointer)
//go:linkname XmlXPatherror C.xmlXPatherror
func XmlXPatherror(ctxt XmlXPathParserContextPtr, file *int8, line c.Int, no c.Int)
//go:linkname XmlXPathErr C.xmlXPathErr
func XmlXPathErr(ctxt XmlXPathParserContextPtr, error c.Int)
//go:linkname XmlXPathDebugDumpObject C.xmlXPathDebugDumpObject
func XmlXPathDebugDumpObject(output *c.FILE, cur XmlXPathObjectPtr, depth c.Int)
//go:linkname XmlXPathDebugDumpCompExpr C.xmlXPathDebugDumpCompExpr
func XmlXPathDebugDumpCompExpr(output *c.FILE, comp XmlXPathCompExprPtr, depth c.Int)
/**
 * NodeSet handling.
 */
//go:linkname XmlXPathNodeSetContains C.xmlXPathNodeSetContains
func XmlXPathNodeSetContains(cur XmlNodeSetPtr, val XmlNodePtr) c.Int
//go:linkname XmlXPathDifference C.xmlXPathDifference
func XmlXPathDifference(nodes1 XmlNodeSetPtr, nodes2 XmlNodeSetPtr) XmlNodeSetPtr
//go:linkname XmlXPathIntersection C.xmlXPathIntersection
func XmlXPathIntersection(nodes1 XmlNodeSetPtr, nodes2 XmlNodeSetPtr) XmlNodeSetPtr
//go:linkname XmlXPathDistinctSorted C.xmlXPathDistinctSorted
func XmlXPathDistinctSorted(nodes XmlNodeSetPtr) XmlNodeSetPtr
//go:linkname XmlXPathDistinct C.xmlXPathDistinct
func XmlXPathDistinct(nodes XmlNodeSetPtr) XmlNodeSetPtr
//go:linkname XmlXPathHasSameNodes C.xmlXPathHasSameNodes
func XmlXPathHasSameNodes(nodes1 XmlNodeSetPtr, nodes2 XmlNodeSetPtr) c.Int
//go:linkname XmlXPathNodeLeadingSorted C.xmlXPathNodeLeadingSorted
func XmlXPathNodeLeadingSorted(nodes XmlNodeSetPtr, node XmlNodePtr) XmlNodeSetPtr
//go:linkname XmlXPathLeadingSorted C.xmlXPathLeadingSorted
func XmlXPathLeadingSorted(nodes1 XmlNodeSetPtr, nodes2 XmlNodeSetPtr) XmlNodeSetPtr
//go:linkname XmlXPathNodeLeading C.xmlXPathNodeLeading
func XmlXPathNodeLeading(nodes XmlNodeSetPtr, node XmlNodePtr) XmlNodeSetPtr
//go:linkname XmlXPathLeading C.xmlXPathLeading
func XmlXPathLeading(nodes1 XmlNodeSetPtr, nodes2 XmlNodeSetPtr) XmlNodeSetPtr
//go:linkname XmlXPathNodeTrailingSorted C.xmlXPathNodeTrailingSorted
func XmlXPathNodeTrailingSorted(nodes XmlNodeSetPtr, node XmlNodePtr) XmlNodeSetPtr
//go:linkname XmlXPathTrailingSorted C.xmlXPathTrailingSorted
func XmlXPathTrailingSorted(nodes1 XmlNodeSetPtr, nodes2 XmlNodeSetPtr) XmlNodeSetPtr
//go:linkname XmlXPathNodeTrailing C.xmlXPathNodeTrailing
func XmlXPathNodeTrailing(nodes XmlNodeSetPtr, node XmlNodePtr) XmlNodeSetPtr
//go:linkname XmlXPathTrailing C.xmlXPathTrailing
func XmlXPathTrailing(nodes1 XmlNodeSetPtr, nodes2 XmlNodeSetPtr) XmlNodeSetPtr
/**
 * Extending a context.
 */
//go:linkname XmlXPathRegisterNs C.xmlXPathRegisterNs
func XmlXPathRegisterNs(ctxt XmlXPathContextPtr, prefix *XmlChar, ns_uri *XmlChar) c.Int
//go:linkname XmlXPathNsLookup C.xmlXPathNsLookup
func XmlXPathNsLookup(ctxt XmlXPathContextPtr, prefix *XmlChar) *XmlChar
//go:linkname XmlXPathRegisteredNsCleanup C.xmlXPathRegisteredNsCleanup
func XmlXPathRegisteredNsCleanup(ctxt XmlXPathContextPtr)
//go:linkname XmlXPathRegisterFunc C.xmlXPathRegisterFunc
func XmlXPathRegisterFunc(ctxt XmlXPathContextPtr, name *XmlChar, f XmlXPathFunction) c.Int
//go:linkname XmlXPathRegisterFuncNS C.xmlXPathRegisterFuncNS
func XmlXPathRegisterFuncNS(ctxt XmlXPathContextPtr, name *XmlChar, ns_uri *XmlChar, f XmlXPathFunction) c.Int
//go:linkname XmlXPathRegisterVariable C.xmlXPathRegisterVariable
func XmlXPathRegisterVariable(ctxt XmlXPathContextPtr, name *XmlChar, value XmlXPathObjectPtr) c.Int
//go:linkname XmlXPathRegisterVariableNS C.xmlXPathRegisterVariableNS
func XmlXPathRegisterVariableNS(ctxt XmlXPathContextPtr, name *XmlChar, ns_uri *XmlChar, value XmlXPathObjectPtr) c.Int
//go:linkname XmlXPathFunctionLookup C.xmlXPathFunctionLookup
func XmlXPathFunctionLookup(ctxt XmlXPathContextPtr, name *XmlChar) XmlXPathFunction
//go:linkname XmlXPathFunctionLookupNS C.xmlXPathFunctionLookupNS
func XmlXPathFunctionLookupNS(ctxt XmlXPathContextPtr, name *XmlChar, ns_uri *XmlChar) XmlXPathFunction
//go:linkname XmlXPathRegisteredFuncsCleanup C.xmlXPathRegisteredFuncsCleanup
func XmlXPathRegisteredFuncsCleanup(ctxt XmlXPathContextPtr)
//go:linkname XmlXPathVariableLookup C.xmlXPathVariableLookup
func XmlXPathVariableLookup(ctxt XmlXPathContextPtr, name *XmlChar) XmlXPathObjectPtr
//go:linkname XmlXPathVariableLookupNS C.xmlXPathVariableLookupNS
func XmlXPathVariableLookupNS(ctxt XmlXPathContextPtr, name *XmlChar, ns_uri *XmlChar) XmlXPathObjectPtr
//go:linkname XmlXPathRegisteredVariablesCleanup C.xmlXPathRegisteredVariablesCleanup
func XmlXPathRegisteredVariablesCleanup(ctxt XmlXPathContextPtr)
/**
 * Utilities to extend XPath.
 */
// llgo:link (*XmlChar).XmlXPathNewParserContext C.xmlXPathNewParserContext
func (recv_ *XmlChar) XmlXPathNewParserContext(ctxt XmlXPathContextPtr) XmlXPathParserContextPtr {
	return nil
}
//go:linkname XmlXPathFreeParserContext C.xmlXPathFreeParserContext
func XmlXPathFreeParserContext(ctxt XmlXPathParserContextPtr)
//go:linkname ValuePop C.valuePop
func ValuePop(ctxt XmlXPathParserContextPtr) XmlXPathObjectPtr
//go:linkname ValuePush C.valuePush
func ValuePush(ctxt XmlXPathParserContextPtr, value XmlXPathObjectPtr) c.Int
// llgo:link (*XmlChar).XmlXPathNewString C.xmlXPathNewString
func (recv_ *XmlChar) XmlXPathNewString() XmlXPathObjectPtr {
	return nil
}
//go:linkname XmlXPathNewCString C.xmlXPathNewCString
func XmlXPathNewCString(val *int8) XmlXPathObjectPtr
// llgo:link (*XmlChar).XmlXPathWrapString C.xmlXPathWrapString
func (recv_ *XmlChar) XmlXPathWrapString() XmlXPathObjectPtr {
	return nil
}
//go:linkname XmlXPathWrapCString C.xmlXPathWrapCString
func XmlXPathWrapCString(val *int8) XmlXPathObjectPtr
//go:linkname XmlXPathNewFloat C.xmlXPathNewFloat
func XmlXPathNewFloat(val float64) XmlXPathObjectPtr
//go:linkname XmlXPathNewBoolean C.xmlXPathNewBoolean
func XmlXPathNewBoolean(val c.Int) XmlXPathObjectPtr
//go:linkname XmlXPathNewNodeSet C.xmlXPathNewNodeSet
func XmlXPathNewNodeSet(val XmlNodePtr) XmlXPathObjectPtr
//go:linkname XmlXPathNewValueTree C.xmlXPathNewValueTree
func XmlXPathNewValueTree(val XmlNodePtr) XmlXPathObjectPtr
//go:linkname XmlXPathNodeSetAdd C.xmlXPathNodeSetAdd
func XmlXPathNodeSetAdd(cur XmlNodeSetPtr, val XmlNodePtr) c.Int
//go:linkname XmlXPathNodeSetAddUnique C.xmlXPathNodeSetAddUnique
func XmlXPathNodeSetAddUnique(cur XmlNodeSetPtr, val XmlNodePtr) c.Int
//go:linkname XmlXPathNodeSetAddNs C.xmlXPathNodeSetAddNs
func XmlXPathNodeSetAddNs(cur XmlNodeSetPtr, node XmlNodePtr, ns XmlNsPtr) c.Int
//go:linkname XmlXPathNodeSetSort C.xmlXPathNodeSetSort
func XmlXPathNodeSetSort(set XmlNodeSetPtr)
//go:linkname XmlXPathRoot C.xmlXPathRoot
func XmlXPathRoot(ctxt XmlXPathParserContextPtr)
//go:linkname XmlXPathEvalExpr C.xmlXPathEvalExpr
func XmlXPathEvalExpr(ctxt XmlXPathParserContextPtr)
//go:linkname XmlXPathParseName C.xmlXPathParseName
func XmlXPathParseName(ctxt XmlXPathParserContextPtr) *XmlChar
//go:linkname XmlXPathParseNCName C.xmlXPathParseNCName
func XmlXPathParseNCName(ctxt XmlXPathParserContextPtr) *XmlChar
// llgo:link (*XmlChar).XmlXPathStringEvalNumber C.xmlXPathStringEvalNumber
func (recv_ *XmlChar) XmlXPathStringEvalNumber() float64 {
	return 0
}
//go:linkname XmlXPathEvaluatePredicateResult C.xmlXPathEvaluatePredicateResult
func XmlXPathEvaluatePredicateResult(ctxt XmlXPathParserContextPtr, res XmlXPathObjectPtr) c.Int
//go:linkname XmlXPathRegisterAllFunctions C.xmlXPathRegisterAllFunctions
func XmlXPathRegisterAllFunctions(ctxt XmlXPathContextPtr)
//go:linkname XmlXPathNodeSetMerge C.xmlXPathNodeSetMerge
func XmlXPathNodeSetMerge(val1 XmlNodeSetPtr, val2 XmlNodeSetPtr) XmlNodeSetPtr
//go:linkname XmlXPathNodeSetDel C.xmlXPathNodeSetDel
func XmlXPathNodeSetDel(cur XmlNodeSetPtr, val XmlNodePtr)
//go:linkname XmlXPathNodeSetRemove C.xmlXPathNodeSetRemove
func XmlXPathNodeSetRemove(cur XmlNodeSetPtr, val c.Int)
//go:linkname XmlXPathNewNodeSetList C.xmlXPathNewNodeSetList
func XmlXPathNewNodeSetList(val XmlNodeSetPtr) XmlXPathObjectPtr
//go:linkname XmlXPathWrapNodeSet C.xmlXPathWrapNodeSet
func XmlXPathWrapNodeSet(val XmlNodeSetPtr) XmlXPathObjectPtr
//go:linkname XmlXPathWrapExternal C.xmlXPathWrapExternal
func XmlXPathWrapExternal(val unsafe.Pointer) XmlXPathObjectPtr
//go:linkname XmlXPathEqualValues C.xmlXPathEqualValues
func XmlXPathEqualValues(ctxt XmlXPathParserContextPtr) c.Int
//go:linkname XmlXPathNotEqualValues C.xmlXPathNotEqualValues
func XmlXPathNotEqualValues(ctxt XmlXPathParserContextPtr) c.Int
//go:linkname XmlXPathCompareValues C.xmlXPathCompareValues
func XmlXPathCompareValues(ctxt XmlXPathParserContextPtr, inf c.Int, strict c.Int) c.Int
//go:linkname XmlXPathValueFlipSign C.xmlXPathValueFlipSign
func XmlXPathValueFlipSign(ctxt XmlXPathParserContextPtr)
//go:linkname XmlXPathAddValues C.xmlXPathAddValues
func XmlXPathAddValues(ctxt XmlXPathParserContextPtr)
//go:linkname XmlXPathSubValues C.xmlXPathSubValues
func XmlXPathSubValues(ctxt XmlXPathParserContextPtr)
//go:linkname XmlXPathMultValues C.xmlXPathMultValues
func XmlXPathMultValues(ctxt XmlXPathParserContextPtr)
//go:linkname XmlXPathDivValues C.xmlXPathDivValues
func XmlXPathDivValues(ctxt XmlXPathParserContextPtr)
//go:linkname XmlXPathModValues C.xmlXPathModValues
func XmlXPathModValues(ctxt XmlXPathParserContextPtr)
// llgo:link (*XmlChar).XmlXPathIsNodeType C.xmlXPathIsNodeType
func (recv_ *XmlChar) XmlXPathIsNodeType() c.Int {
	return 0
}
//go:linkname XmlXPathNextSelf C.xmlXPathNextSelf
func XmlXPathNextSelf(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextChild C.xmlXPathNextChild
func XmlXPathNextChild(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextDescendant C.xmlXPathNextDescendant
func XmlXPathNextDescendant(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextDescendantOrSelf C.xmlXPathNextDescendantOrSelf
func XmlXPathNextDescendantOrSelf(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextParent C.xmlXPathNextParent
func XmlXPathNextParent(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextAncestorOrSelf C.xmlXPathNextAncestorOrSelf
func XmlXPathNextAncestorOrSelf(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextFollowingSibling C.xmlXPathNextFollowingSibling
func XmlXPathNextFollowingSibling(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextFollowing C.xmlXPathNextFollowing
func XmlXPathNextFollowing(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextNamespace C.xmlXPathNextNamespace
func XmlXPathNextNamespace(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextAttribute C.xmlXPathNextAttribute
func XmlXPathNextAttribute(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextPreceding C.xmlXPathNextPreceding
func XmlXPathNextPreceding(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextAncestor C.xmlXPathNextAncestor
func XmlXPathNextAncestor(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathNextPrecedingSibling C.xmlXPathNextPrecedingSibling
func XmlXPathNextPrecedingSibling(ctxt XmlXPathParserContextPtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlXPathLastFunction C.xmlXPathLastFunction
func XmlXPathLastFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathPositionFunction C.xmlXPathPositionFunction
func XmlXPathPositionFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathCountFunction C.xmlXPathCountFunction
func XmlXPathCountFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathIdFunction C.xmlXPathIdFunction
func XmlXPathIdFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathLocalNameFunction C.xmlXPathLocalNameFunction
func XmlXPathLocalNameFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathNamespaceURIFunction C.xmlXPathNamespaceURIFunction
func XmlXPathNamespaceURIFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathStringFunction C.xmlXPathStringFunction
func XmlXPathStringFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathStringLengthFunction C.xmlXPathStringLengthFunction
func XmlXPathStringLengthFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathConcatFunction C.xmlXPathConcatFunction
func XmlXPathConcatFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathContainsFunction C.xmlXPathContainsFunction
func XmlXPathContainsFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathStartsWithFunction C.xmlXPathStartsWithFunction
func XmlXPathStartsWithFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathSubstringFunction C.xmlXPathSubstringFunction
func XmlXPathSubstringFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathSubstringBeforeFunction C.xmlXPathSubstringBeforeFunction
func XmlXPathSubstringBeforeFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathSubstringAfterFunction C.xmlXPathSubstringAfterFunction
func XmlXPathSubstringAfterFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathNormalizeFunction C.xmlXPathNormalizeFunction
func XmlXPathNormalizeFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathTranslateFunction C.xmlXPathTranslateFunction
func XmlXPathTranslateFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathNotFunction C.xmlXPathNotFunction
func XmlXPathNotFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathTrueFunction C.xmlXPathTrueFunction
func XmlXPathTrueFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathFalseFunction C.xmlXPathFalseFunction
func XmlXPathFalseFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathLangFunction C.xmlXPathLangFunction
func XmlXPathLangFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathNumberFunction C.xmlXPathNumberFunction
func XmlXPathNumberFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathSumFunction C.xmlXPathSumFunction
func XmlXPathSumFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathFloorFunction C.xmlXPathFloorFunction
func XmlXPathFloorFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathCeilingFunction C.xmlXPathCeilingFunction
func XmlXPathCeilingFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathRoundFunction C.xmlXPathRoundFunction
func XmlXPathRoundFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
//go:linkname XmlXPathBooleanFunction C.xmlXPathBooleanFunction
func XmlXPathBooleanFunction(ctxt XmlXPathParserContextPtr, nargs c.Int)
/**
 * Really internal functions
 */
//go:linkname XmlXPathNodeSetFreeNs C.xmlXPathNodeSetFreeNs
func XmlXPathNodeSetFreeNs(ns XmlNsPtr)
