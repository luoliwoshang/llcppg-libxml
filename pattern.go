package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type X_XmlPattern struct {
	Unused [8]uint8
}
type XmlPattern X_XmlPattern
type XmlPatternPtr *XmlPattern
type XmlPatternFlags c.Int

const (
	XmlPatternFlagsXMLPATTERNDEFAULT XmlPatternFlags = 0
	XmlPatternFlagsXMLPATTERNXPATH   XmlPatternFlags = 1
	XmlPatternFlagsXMLPATTERNXSSEL   XmlPatternFlags = 2
	XmlPatternFlagsXMLPATTERNXSFIELD XmlPatternFlags = 4
)
//go:linkname XmlFreePattern C.xmlFreePattern
func XmlFreePattern(comp XmlPatternPtr)
//go:linkname XmlFreePatternList C.xmlFreePatternList
func XmlFreePatternList(comp XmlPatternPtr)
// llgo:link (*XmlChar).XmlPatterncompile C.xmlPatterncompile
func (recv_ *XmlChar) XmlPatterncompile(dict *XmlDict, flags c.Int, namespaces **XmlChar) XmlPatternPtr {
	return nil
}
// llgo:link (*XmlChar).XmlPatternCompileSafe C.xmlPatternCompileSafe
func (recv_ *XmlChar) XmlPatternCompileSafe(dict *XmlDict, flags c.Int, namespaces **XmlChar, patternOut *XmlPatternPtr) c.Int {
	return 0
}
//go:linkname XmlPatternMatch C.xmlPatternMatch
func XmlPatternMatch(comp XmlPatternPtr, node XmlNodePtr) c.Int

type X_XmlStreamCtxt struct {
	Unused [8]uint8
}
type XmlStreamCtxt X_XmlStreamCtxt
type XmlStreamCtxtPtr *XmlStreamCtxt
//go:linkname XmlPatternStreamable C.xmlPatternStreamable
func XmlPatternStreamable(comp XmlPatternPtr) c.Int
//go:linkname XmlPatternMaxDepth C.xmlPatternMaxDepth
func XmlPatternMaxDepth(comp XmlPatternPtr) c.Int
//go:linkname XmlPatternMinDepth C.xmlPatternMinDepth
func XmlPatternMinDepth(comp XmlPatternPtr) c.Int
//go:linkname XmlPatternFromRoot C.xmlPatternFromRoot
func XmlPatternFromRoot(comp XmlPatternPtr) c.Int
//go:linkname XmlPatternGetStreamCtxt C.xmlPatternGetStreamCtxt
func XmlPatternGetStreamCtxt(comp XmlPatternPtr) XmlStreamCtxtPtr
//go:linkname XmlFreeStreamCtxt C.xmlFreeStreamCtxt
func XmlFreeStreamCtxt(stream XmlStreamCtxtPtr)
//go:linkname XmlStreamPushNode C.xmlStreamPushNode
func XmlStreamPushNode(stream XmlStreamCtxtPtr, name *XmlChar, ns *XmlChar, nodeType c.Int) c.Int
//go:linkname XmlStreamPush C.xmlStreamPush
func XmlStreamPush(stream XmlStreamCtxtPtr, name *XmlChar, ns *XmlChar) c.Int
//go:linkname XmlStreamPushAttr C.xmlStreamPushAttr
func XmlStreamPushAttr(stream XmlStreamCtxtPtr, name *XmlChar, ns *XmlChar) c.Int
//go:linkname XmlStreamPop C.xmlStreamPop
func XmlStreamPop(stream XmlStreamCtxtPtr) c.Int
//go:linkname XmlStreamWantsAnyNode C.xmlStreamWantsAnyNode
func XmlStreamWantsAnyNode(stream XmlStreamCtxtPtr) c.Int
