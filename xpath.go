package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const XPATHPOINT c.Int = 5
const XPATHRANGE c.Int = 6
const XPATHLOCATIONSET c.Int = 7

type X_XmlXPathContext struct {
	Doc                XmlDocPtr
	Node               XmlNodePtr
	NbVariablesUnused  c.Int
	MaxVariablesUnused c.Int
	VarHash            XmlHashTablePtr
	NbTypes            c.Int
	MaxTypes           c.Int
	Types              XmlXPathTypePtr
	NbFuncsUnused      c.Int
	MaxFuncsUnused     c.Int
	FuncHash           XmlHashTablePtr
	NbAxis             c.Int
	MaxAxis            c.Int
	Axis               XmlXPathAxisPtr
	Namespaces         *XmlNsPtr
	NsNr               c.Int
	User               unsafe.Pointer
	ContextSize        c.Int
	ProximityPosition  c.Int
	Xptr               c.Int
	Here               XmlNodePtr
	Origin             XmlNodePtr
	NsHash             XmlHashTablePtr
	VarLookupFunc      unsafe.Pointer
	VarLookupData      unsafe.Pointer
	Extra              unsafe.Pointer
	Function           *XmlChar
	FunctionURI        *XmlChar
	FuncLookupFunc     unsafe.Pointer
	FuncLookupData     unsafe.Pointer
	TmpNsList          *XmlNsPtr
	TmpNsNr            c.Int
	UserData           unsafe.Pointer
	Error              unsafe.Pointer
	LastError          XmlError
	DebugNode          XmlNodePtr
	Dict               XmlDictPtr
	Flags              c.Int
	Cache              unsafe.Pointer
	OpLimit            c.Ulong
	OpCount            c.Ulong
	Depth              c.Int
}
type XmlXPathContext X_XmlXPathContext
type XmlXPathContextPtr *XmlXPathContext

type X_XmlXPathParserContext struct {
	Cur        *XmlChar
	Base       *XmlChar
	Error      c.Int
	Context    XmlXPathContextPtr
	Value      XmlXPathObjectPtr
	ValueNr    c.Int
	ValueMax   c.Int
	ValueTab   *XmlXPathObjectPtr
	Comp       XmlXPathCompExprPtr
	Xptr       c.Int
	Ancestor   XmlNodePtr
	ValueFrame c.Int
}
type XmlXPathParserContext X_XmlXPathParserContext
type XmlXPathParserContextPtr *XmlXPathParserContext
type XmlXPathError c.Int

const (
	XmlXPathErrorXPATHEXPRESSIONOK           XmlXPathError = 0
	XmlXPathErrorXPATHNUMBERERROR            XmlXPathError = 1
	XmlXPathErrorXPATHUNFINISHEDLITERALERROR XmlXPathError = 2
	XmlXPathErrorXPATHSTARTLITERALERROR      XmlXPathError = 3
	XmlXPathErrorXPATHVARIABLEREFERROR       XmlXPathError = 4
	XmlXPathErrorXPATHUNDEFVARIABLEERROR     XmlXPathError = 5
	XmlXPathErrorXPATHINVALIDPREDICATEERROR  XmlXPathError = 6
	XmlXPathErrorXPATHEXPRERROR              XmlXPathError = 7
	XmlXPathErrorXPATHUNCLOSEDERROR          XmlXPathError = 8
	XmlXPathErrorXPATHUNKNOWNFUNCERROR       XmlXPathError = 9
	XmlXPathErrorXPATHINVALIDOPERAND         XmlXPathError = 10
	XmlXPathErrorXPATHINVALIDTYPE            XmlXPathError = 11
	XmlXPathErrorXPATHINVALIDARITY           XmlXPathError = 12
	XmlXPathErrorXPATHINVALIDCTXTSIZE        XmlXPathError = 13
	XmlXPathErrorXPATHINVALIDCTXTPOSITION    XmlXPathError = 14
	XmlXPathErrorXPATHMEMORYERROR            XmlXPathError = 15
	XmlXPathErrorXPTRSYNTAXERROR             XmlXPathError = 16
	XmlXPathErrorXPTRRESOURCEERROR           XmlXPathError = 17
	XmlXPathErrorXPTRSUBRESOURCEERROR        XmlXPathError = 18
	XmlXPathErrorXPATHUNDEFPREFIXERROR       XmlXPathError = 19
	XmlXPathErrorXPATHENCODINGERROR          XmlXPathError = 20
	XmlXPathErrorXPATHINVALIDCHARERROR       XmlXPathError = 21
	XmlXPathErrorXPATHINVALIDCTXT            XmlXPathError = 22
	XmlXPathErrorXPATHSTACKERROR             XmlXPathError = 23
	XmlXPathErrorXPATHFORBIDVARIABLEERROR    XmlXPathError = 24
	XmlXPathErrorXPATHOPLIMITEXCEEDED        XmlXPathError = 25
	XmlXPathErrorXPATHRECURSIONLIMITEXCEEDED XmlXPathError = 26
)

type X_XmlNodeSet struct {
	NodeNr  c.Int
	NodeMax c.Int
	NodeTab *XmlNodePtr
}
type XmlNodeSet X_XmlNodeSet
type XmlNodeSetPtr *XmlNodeSet
type XmlXPathObjectType c.Int

const (
	XmlXPathObjectTypeXPATHUNDEFINED XmlXPathObjectType = 0
	XmlXPathObjectTypeXPATHNODESET   XmlXPathObjectType = 1
	XmlXPathObjectTypeXPATHBOOLEAN   XmlXPathObjectType = 2
	XmlXPathObjectTypeXPATHNUMBER    XmlXPathObjectType = 3
	XmlXPathObjectTypeXPATHSTRING    XmlXPathObjectType = 4
	XmlXPathObjectTypeXPATHUSERS     XmlXPathObjectType = 8
	XmlXPathObjectTypeXPATHXSLTTREE  XmlXPathObjectType = 9
)

type X_XmlXPathObject struct {
	Type       XmlXPathObjectType
	Nodesetval XmlNodeSetPtr
	Boolval    c.Int
	Floatval   float64
	Stringval  *XmlChar
	User       unsafe.Pointer
	Index      c.Int
	User2      unsafe.Pointer
	Index2     c.Int
}
type XmlXPathObject X_XmlXPathObject
type XmlXPathObjectPtr *XmlXPathObject
// llgo:type C
type XmlXPathConvertFunc func(XmlXPathObjectPtr, c.Int) c.Int

type X_XmlXPathType struct {
	Name *XmlChar
	Func unsafe.Pointer
}
type XmlXPathType X_XmlXPathType
type XmlXPathTypePtr *XmlXPathType

type X_XmlXPathVariable struct {
	Name  *XmlChar
	Value XmlXPathObjectPtr
}
type XmlXPathVariable X_XmlXPathVariable
type XmlXPathVariablePtr *XmlXPathVariable
// llgo:type C
type XmlXPathEvalFunc func(XmlXPathParserContextPtr, c.Int)

type X_XmlXPathFunct struct {
	Name *XmlChar
	Func unsafe.Pointer
}
type XmlXPathFunct X_XmlXPathFunct
type XmlXPathFuncPtr *XmlXPathFunct
// llgo:type C
type XmlXPathAxisFunc func(XmlXPathParserContextPtr, XmlXPathObjectPtr) XmlXPathObjectPtr

type X_XmlXPathAxis struct {
	Name *XmlChar
	Func unsafe.Pointer
}
type XmlXPathAxis X_XmlXPathAxis
type XmlXPathAxisPtr *XmlXPathAxis
// llgo:type C
type XmlXPathFunction func(XmlXPathParserContextPtr, c.Int)
// llgo:type C
type XmlXPathVariableLookupFunc func(unsafe.Pointer, *XmlChar, *XmlChar) XmlXPathObjectPtr
// llgo:type C
type XmlXPathFuncLookupFunc func(unsafe.Pointer, *XmlChar, *XmlChar) XmlXPathFunction

type X_XmlXPathCompExpr struct {
	Unused [8]uint8
}
type XmlXPathCompExpr X_XmlXPathCompExpr
type XmlXPathCompExprPtr *XmlXPathCompExpr
//go:linkname XmlXPathFreeObject C.xmlXPathFreeObject
func XmlXPathFreeObject(obj XmlXPathObjectPtr)
//go:linkname XmlXPathNodeSetCreate C.xmlXPathNodeSetCreate
func XmlXPathNodeSetCreate(val XmlNodePtr) XmlNodeSetPtr
//go:linkname XmlXPathFreeNodeSetList C.xmlXPathFreeNodeSetList
func XmlXPathFreeNodeSetList(obj XmlXPathObjectPtr)
//go:linkname XmlXPathFreeNodeSet C.xmlXPathFreeNodeSet
func XmlXPathFreeNodeSet(obj XmlNodeSetPtr)
//go:linkname XmlXPathObjectCopy C.xmlXPathObjectCopy
func XmlXPathObjectCopy(val XmlXPathObjectPtr) XmlXPathObjectPtr
//go:linkname XmlXPathCmpNodes C.xmlXPathCmpNodes
func XmlXPathCmpNodes(node1 XmlNodePtr, node2 XmlNodePtr) c.Int
/**
 * Conversion functions to basic types.
 */
//go:linkname XmlXPathCastNumberToBoolean C.xmlXPathCastNumberToBoolean
func XmlXPathCastNumberToBoolean(val float64) c.Int
// llgo:link (*XmlChar).XmlXPathCastStringToBoolean C.xmlXPathCastStringToBoolean
func (recv_ *XmlChar) XmlXPathCastStringToBoolean() c.Int {
	return 0
}
//go:linkname XmlXPathCastNodeSetToBoolean C.xmlXPathCastNodeSetToBoolean
func XmlXPathCastNodeSetToBoolean(ns XmlNodeSetPtr) c.Int
//go:linkname XmlXPathCastToBoolean C.xmlXPathCastToBoolean
func XmlXPathCastToBoolean(val XmlXPathObjectPtr) c.Int
//go:linkname XmlXPathCastBooleanToNumber C.xmlXPathCastBooleanToNumber
func XmlXPathCastBooleanToNumber(val c.Int) float64
// llgo:link (*XmlChar).XmlXPathCastStringToNumber C.xmlXPathCastStringToNumber
func (recv_ *XmlChar) XmlXPathCastStringToNumber() float64 {
	return 0
}
//go:linkname XmlXPathCastNodeToNumber C.xmlXPathCastNodeToNumber
func XmlXPathCastNodeToNumber(node XmlNodePtr) float64
//go:linkname XmlXPathCastNodeSetToNumber C.xmlXPathCastNodeSetToNumber
func XmlXPathCastNodeSetToNumber(ns XmlNodeSetPtr) float64
//go:linkname XmlXPathCastToNumber C.xmlXPathCastToNumber
func XmlXPathCastToNumber(val XmlXPathObjectPtr) float64
//go:linkname XmlXPathCastBooleanToString C.xmlXPathCastBooleanToString
func XmlXPathCastBooleanToString(val c.Int) *XmlChar
//go:linkname XmlXPathCastNumberToString C.xmlXPathCastNumberToString
func XmlXPathCastNumberToString(val float64) *XmlChar
//go:linkname XmlXPathCastNodeToString C.xmlXPathCastNodeToString
func XmlXPathCastNodeToString(node XmlNodePtr) *XmlChar
//go:linkname XmlXPathCastNodeSetToString C.xmlXPathCastNodeSetToString
func XmlXPathCastNodeSetToString(ns XmlNodeSetPtr) *XmlChar
//go:linkname XmlXPathCastToString C.xmlXPathCastToString
func XmlXPathCastToString(val XmlXPathObjectPtr) *XmlChar
//go:linkname XmlXPathConvertBoolean C.xmlXPathConvertBoolean
func XmlXPathConvertBoolean(val XmlXPathObjectPtr) XmlXPathObjectPtr
//go:linkname XmlXPathConvertNumber C.xmlXPathConvertNumber
func XmlXPathConvertNumber(val XmlXPathObjectPtr) XmlXPathObjectPtr
//go:linkname XmlXPathConvertString C.xmlXPathConvertString
func XmlXPathConvertString(val XmlXPathObjectPtr) XmlXPathObjectPtr
/**
 * Context handling.
 */
//go:linkname XmlXPathNewContext C.xmlXPathNewContext
func XmlXPathNewContext(doc XmlDocPtr) XmlXPathContextPtr
//go:linkname XmlXPathFreeContext C.xmlXPathFreeContext
func XmlXPathFreeContext(ctxt XmlXPathContextPtr)
//go:linkname XmlXPathSetErrorHandler C.xmlXPathSetErrorHandler
func XmlXPathSetErrorHandler(ctxt XmlXPathContextPtr, handler XmlStructuredErrorFunc, context unsafe.Pointer)
//go:linkname XmlXPathContextSetCache C.xmlXPathContextSetCache
func XmlXPathContextSetCache(ctxt XmlXPathContextPtr, active c.Int, value c.Int, options c.Int) c.Int
/**
 * Evaluation functions.
 */
//go:linkname XmlXPathOrderDocElems C.xmlXPathOrderDocElems
func XmlXPathOrderDocElems(doc XmlDocPtr) c.Long
//go:linkname XmlXPathSetContextNode C.xmlXPathSetContextNode
func XmlXPathSetContextNode(node XmlNodePtr, ctx XmlXPathContextPtr) c.Int
//go:linkname XmlXPathNodeEval C.xmlXPathNodeEval
func XmlXPathNodeEval(node XmlNodePtr, str *XmlChar, ctx XmlXPathContextPtr) XmlXPathObjectPtr
// llgo:link (*XmlChar).XmlXPathEval C.xmlXPathEval
func (recv_ *XmlChar) XmlXPathEval(ctx XmlXPathContextPtr) XmlXPathObjectPtr {
	return nil
}
// llgo:link (*XmlChar).XmlXPathEvalExpression C.xmlXPathEvalExpression
func (recv_ *XmlChar) XmlXPathEvalExpression(ctxt XmlXPathContextPtr) XmlXPathObjectPtr {
	return nil
}
//go:linkname XmlXPathEvalPredicate C.xmlXPathEvalPredicate
func XmlXPathEvalPredicate(ctxt XmlXPathContextPtr, res XmlXPathObjectPtr) c.Int
/**
 * Separate compilation/evaluation entry points.
 */
// llgo:link (*XmlChar).XmlXPathCompile C.xmlXPathCompile
func (recv_ *XmlChar) XmlXPathCompile() XmlXPathCompExprPtr {
	return nil
}
//go:linkname XmlXPathCtxtCompile C.xmlXPathCtxtCompile
func XmlXPathCtxtCompile(ctxt XmlXPathContextPtr, str *XmlChar) XmlXPathCompExprPtr
//go:linkname XmlXPathCompiledEval C.xmlXPathCompiledEval
func XmlXPathCompiledEval(comp XmlXPathCompExprPtr, ctx XmlXPathContextPtr) XmlXPathObjectPtr
//go:linkname XmlXPathCompiledEvalToBoolean C.xmlXPathCompiledEvalToBoolean
func XmlXPathCompiledEvalToBoolean(comp XmlXPathCompExprPtr, ctxt XmlXPathContextPtr) c.Int
//go:linkname XmlXPathFreeCompExpr C.xmlXPathFreeCompExpr
func XmlXPathFreeCompExpr(comp XmlXPathCompExprPtr)
//go:linkname XmlXPathInit C.xmlXPathInit
func XmlXPathInit()
//go:linkname XmlXPathIsNaN C.xmlXPathIsNaN
func XmlXPathIsNaN(val float64) c.Int
//go:linkname XmlXPathIsInf C.xmlXPathIsInf
func XmlXPathIsInf(val float64) c.Int
