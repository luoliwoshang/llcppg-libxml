package libxml_2_0

import _ "unsafe"
//go:linkname XmlXPtrNewContext C.xmlXPtrNewContext
func XmlXPtrNewContext(doc XmlDocPtr, here XmlNodePtr, origin XmlNodePtr) XmlXPathContextPtr
// llgo:link (*XmlChar).XmlXPtrEval C.xmlXPtrEval
func (recv_ *XmlChar) XmlXPtrEval(ctx XmlXPathContextPtr) XmlXPathObjectPtr {
	return nil
}
