package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
// llgo:link (*XmlChar).HtmlNewDoc C.htmlNewDoc
func (recv_ *XmlChar) HtmlNewDoc(ExternalID *XmlChar) HtmlDocPtr {
	return nil
}
// llgo:link (*XmlChar).HtmlNewDocNoDtD C.htmlNewDocNoDtD
func (recv_ *XmlChar) HtmlNewDocNoDtD(ExternalID *XmlChar) HtmlDocPtr {
	return nil
}
//go:linkname HtmlGetMetaEncoding C.htmlGetMetaEncoding
func HtmlGetMetaEncoding(doc HtmlDocPtr) *XmlChar
//go:linkname HtmlSetMetaEncoding C.htmlSetMetaEncoding
func HtmlSetMetaEncoding(doc HtmlDocPtr, encoding *XmlChar) c.Int
//go:linkname HtmlDocDumpMemory C.htmlDocDumpMemory
func HtmlDocDumpMemory(cur XmlDocPtr, mem **XmlChar, size *c.Int)
//go:linkname HtmlDocDumpMemoryFormat C.htmlDocDumpMemoryFormat
func HtmlDocDumpMemoryFormat(cur XmlDocPtr, mem **XmlChar, size *c.Int, format c.Int)
//go:linkname HtmlDocDump C.htmlDocDump
func HtmlDocDump(f *c.FILE, cur XmlDocPtr) c.Int
//go:linkname HtmlSaveFile C.htmlSaveFile
func HtmlSaveFile(filename *int8, cur XmlDocPtr) c.Int
//go:linkname HtmlNodeDump C.htmlNodeDump
func HtmlNodeDump(buf XmlBufferPtr, doc XmlDocPtr, cur XmlNodePtr) c.Int
//go:linkname HtmlNodeDumpFile C.htmlNodeDumpFile
func HtmlNodeDumpFile(out *c.FILE, doc XmlDocPtr, cur XmlNodePtr)
//go:linkname HtmlNodeDumpFileFormat C.htmlNodeDumpFileFormat
func HtmlNodeDumpFileFormat(out *c.FILE, doc XmlDocPtr, cur XmlNodePtr, encoding *int8, format c.Int) c.Int
//go:linkname HtmlSaveFileEnc C.htmlSaveFileEnc
func HtmlSaveFileEnc(filename *int8, cur XmlDocPtr, encoding *int8) c.Int
//go:linkname HtmlSaveFileFormat C.htmlSaveFileFormat
func HtmlSaveFileFormat(filename *int8, cur XmlDocPtr, encoding *int8, format c.Int) c.Int
//go:linkname HtmlNodeDumpFormatOutput C.htmlNodeDumpFormatOutput
func HtmlNodeDumpFormatOutput(buf XmlOutputBufferPtr, doc XmlDocPtr, cur XmlNodePtr, encoding *int8, format c.Int)
//go:linkname HtmlDocContentDumpOutput C.htmlDocContentDumpOutput
func HtmlDocContentDumpOutput(buf XmlOutputBufferPtr, cur XmlDocPtr, encoding *int8)
//go:linkname HtmlDocContentDumpFormatOutput C.htmlDocContentDumpFormatOutput
func HtmlDocContentDumpFormatOutput(buf XmlOutputBufferPtr, cur XmlDocPtr, encoding *int8, format c.Int)
//go:linkname HtmlNodeDumpOutput C.htmlNodeDumpOutput
func HtmlNodeDumpOutput(buf XmlOutputBufferPtr, doc XmlDocPtr, cur XmlNodePtr, encoding *int8)
// llgo:link (*XmlChar).HtmlIsBooleanAttr C.htmlIsBooleanAttr
func (recv_ *XmlChar) HtmlIsBooleanAttr() c.Int {
	return 0
}
