package libxml

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

// llgo:link (*Char).HtmlNewDoc C.htmlNewDoc
func (recv_ *Char) HtmlNewDoc(ExternalID *Char) HtmlDocPtr {
	return nil
}

// llgo:link (*Char).HtmlNewDocNoDtD C.htmlNewDocNoDtD
func (recv_ *Char) HtmlNewDocNoDtD(ExternalID *Char) HtmlDocPtr {
	return nil
}

//go:linkname HtmlGetMetaEncoding C.htmlGetMetaEncoding
func HtmlGetMetaEncoding(doc HtmlDocPtr) *Char

//go:linkname HtmlSetMetaEncoding C.htmlSetMetaEncoding
func HtmlSetMetaEncoding(doc HtmlDocPtr, encoding *Char) c.Int

//go:linkname HtmlDocDumpMemory C.htmlDocDumpMemory
func HtmlDocDumpMemory(cur DocPtr, mem **Char, size *c.Int)

//go:linkname HtmlDocDumpMemoryFormat C.htmlDocDumpMemoryFormat
func HtmlDocDumpMemoryFormat(cur DocPtr, mem **Char, size *c.Int, format c.Int)

//go:linkname HtmlDocDump C.htmlDocDump
func HtmlDocDump(f *c.FILE, cur DocPtr) c.Int

//go:linkname HtmlSaveFile C.htmlSaveFile
func HtmlSaveFile(filename *int8, cur DocPtr) c.Int

//go:linkname HtmlNodeDump C.htmlNodeDump
func HtmlNodeDump(buf BufferPtr, doc DocPtr, cur NodePtr) c.Int

//go:linkname HtmlNodeDumpFile C.htmlNodeDumpFile
func HtmlNodeDumpFile(out *c.FILE, doc DocPtr, cur NodePtr)

//go:linkname HtmlNodeDumpFileFormat C.htmlNodeDumpFileFormat
func HtmlNodeDumpFileFormat(out *c.FILE, doc DocPtr, cur NodePtr, encoding *int8, format c.Int) c.Int

//go:linkname HtmlSaveFileEnc C.htmlSaveFileEnc
func HtmlSaveFileEnc(filename *int8, cur DocPtr, encoding *int8) c.Int

//go:linkname HtmlSaveFileFormat C.htmlSaveFileFormat
func HtmlSaveFileFormat(filename *int8, cur DocPtr, encoding *int8, format c.Int) c.Int

//go:linkname HtmlNodeDumpFormatOutput C.htmlNodeDumpFormatOutput
func HtmlNodeDumpFormatOutput(buf OutputBufferPtr, doc DocPtr, cur NodePtr, encoding *int8, format c.Int)

//go:linkname HtmlDocContentDumpOutput C.htmlDocContentDumpOutput
func HtmlDocContentDumpOutput(buf OutputBufferPtr, cur DocPtr, encoding *int8)

//go:linkname HtmlDocContentDumpFormatOutput C.htmlDocContentDumpFormatOutput
func HtmlDocContentDumpFormatOutput(buf OutputBufferPtr, cur DocPtr, encoding *int8, format c.Int)

//go:linkname HtmlNodeDumpOutput C.htmlNodeDumpOutput
func HtmlNodeDumpOutput(buf OutputBufferPtr, doc DocPtr, cur NodePtr, encoding *int8)

// llgo:link (*Char).HtmlIsBooleanAttr C.htmlIsBooleanAttr
func (recv_ *Char) HtmlIsBooleanAttr() c.Int {
	return 0
}
