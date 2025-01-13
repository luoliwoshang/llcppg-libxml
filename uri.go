package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type X_XmlURI struct {
	Scheme    *int8
	Opaque    *int8
	Authority *int8
	Server    *int8
	User      *int8
	Port      c.Int
	Path      *int8
	Query     *int8
	Fragment  *int8
	Cleanup   c.Int
	QueryRaw  *int8
}
type XmlURI X_XmlURI
type XmlURIPtr *XmlURI
//go:linkname XmlCreateURI C.xmlCreateURI
func XmlCreateURI() XmlURIPtr
// llgo:link (*XmlChar).XmlBuildURISafe C.xmlBuildURISafe
func (recv_ *XmlChar) XmlBuildURISafe(base *XmlChar, out **XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlBuildURI C.xmlBuildURI
func (recv_ *XmlChar) XmlBuildURI(base *XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlBuildRelativeURISafe C.xmlBuildRelativeURISafe
func (recv_ *XmlChar) XmlBuildRelativeURISafe(base *XmlChar, out **XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlBuildRelativeURI C.xmlBuildRelativeURI
func (recv_ *XmlChar) XmlBuildRelativeURI(base *XmlChar) *XmlChar {
	return nil
}
//go:linkname XmlParseURI C.xmlParseURI
func XmlParseURI(str *int8) XmlURIPtr
//go:linkname XmlParseURISafe C.xmlParseURISafe
func XmlParseURISafe(str *int8, uri *XmlURIPtr) c.Int
//go:linkname XmlParseURIRaw C.xmlParseURIRaw
func XmlParseURIRaw(str *int8, raw c.Int) XmlURIPtr
//go:linkname XmlParseURIReference C.xmlParseURIReference
func XmlParseURIReference(uri XmlURIPtr, str *int8) c.Int
//go:linkname XmlSaveUri C.xmlSaveUri
func XmlSaveUri(uri XmlURIPtr) *XmlChar
//go:linkname XmlPrintURI C.xmlPrintURI
func XmlPrintURI(stream *c.FILE, uri XmlURIPtr)
// llgo:link (*XmlChar).XmlURIEscapeStr C.xmlURIEscapeStr
func (recv_ *XmlChar) XmlURIEscapeStr(list *XmlChar) *XmlChar {
	return nil
}
//go:linkname XmlURIUnescapeString C.xmlURIUnescapeString
func XmlURIUnescapeString(str *int8, len c.Int, target *int8) *int8
//go:linkname XmlNormalizeURIPath C.xmlNormalizeURIPath
func XmlNormalizeURIPath(path *int8) c.Int
// llgo:link (*XmlChar).XmlURIEscape C.xmlURIEscape
func (recv_ *XmlChar) XmlURIEscape() *XmlChar {
	return nil
}
//go:linkname XmlFreeURI C.xmlFreeURI
func XmlFreeURI(uri XmlURIPtr)
// llgo:link (*XmlChar).XmlCanonicPath C.xmlCanonicPath
func (recv_ *XmlChar) XmlCanonicPath() *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlPathToURI C.xmlPathToURI
func (recv_ *XmlChar) XmlPathToURI() *XmlChar {
	return nil
}
