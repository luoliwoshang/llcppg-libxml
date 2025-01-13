package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
// llgo:type C
type XmlInputMatchCallback func(*int8) c.Int
// llgo:type C
type XmlInputOpenCallback func(*int8) unsafe.Pointer
// llgo:type C
type XmlInputReadCallback func(unsafe.Pointer, *int8, c.Int) c.Int
// llgo:type C
type XmlInputCloseCallback func(unsafe.Pointer) c.Int
// llgo:type C
type XmlOutputMatchCallback func(*int8) c.Int
// llgo:type C
type XmlOutputOpenCallback func(*int8) unsafe.Pointer
// llgo:type C
type XmlOutputWriteCallback func(unsafe.Pointer, *int8, c.Int) c.Int
// llgo:type C
type XmlOutputCloseCallback func(unsafe.Pointer) c.Int
// llgo:type C
type XmlParserInputBufferCreateFilenameFunc func(*int8, XmlCharEncoding) XmlParserInputBufferPtr
// llgo:type C
type XmlOutputBufferCreateFilenameFunc func(*int8, XmlCharEncodingHandlerPtr, c.Int) XmlOutputBufferPtr
//go:linkname X__XmlParserInputBufferCreateFilenameValue C.__xmlParserInputBufferCreateFilenameValue
func X__XmlParserInputBufferCreateFilenameValue() XmlParserInputBufferCreateFilenameFunc
//go:linkname X__XmlOutputBufferCreateFilenameValue C.__xmlOutputBufferCreateFilenameValue
func X__XmlOutputBufferCreateFilenameValue() XmlOutputBufferCreateFilenameFunc
/** DOC_ENABLE */
//go:linkname XmlCleanupInputCallbacks C.xmlCleanupInputCallbacks
func XmlCleanupInputCallbacks()
//go:linkname XmlPopInputCallbacks C.xmlPopInputCallbacks
func XmlPopInputCallbacks() c.Int
//go:linkname XmlRegisterDefaultInputCallbacks C.xmlRegisterDefaultInputCallbacks
func XmlRegisterDefaultInputCallbacks()
// llgo:link XmlCharEncoding.XmlAllocParserInputBuffer C.xmlAllocParserInputBuffer
func (recv_ XmlCharEncoding) XmlAllocParserInputBuffer() XmlParserInputBufferPtr {
	return nil
}
//go:linkname XmlParserInputBufferCreateFilename C.xmlParserInputBufferCreateFilename
func XmlParserInputBufferCreateFilename(URI *int8, enc XmlCharEncoding) XmlParserInputBufferPtr
//go:linkname XmlParserInputBufferCreateFile C.xmlParserInputBufferCreateFile
func XmlParserInputBufferCreateFile(file *c.FILE, enc XmlCharEncoding) XmlParserInputBufferPtr
//go:linkname XmlParserInputBufferCreateFd C.xmlParserInputBufferCreateFd
func XmlParserInputBufferCreateFd(fd c.Int, enc XmlCharEncoding) XmlParserInputBufferPtr
//go:linkname XmlParserInputBufferCreateMem C.xmlParserInputBufferCreateMem
func XmlParserInputBufferCreateMem(mem *int8, size c.Int, enc XmlCharEncoding) XmlParserInputBufferPtr
//go:linkname XmlParserInputBufferCreateStatic C.xmlParserInputBufferCreateStatic
func XmlParserInputBufferCreateStatic(mem *int8, size c.Int, enc XmlCharEncoding) XmlParserInputBufferPtr
//go:linkname XmlParserInputBufferCreateIO C.xmlParserInputBufferCreateIO
func XmlParserInputBufferCreateIO(ioread XmlInputReadCallback, ioclose XmlInputCloseCallback, ioctx unsafe.Pointer, enc XmlCharEncoding) XmlParserInputBufferPtr
//go:linkname XmlParserInputBufferRead C.xmlParserInputBufferRead
func XmlParserInputBufferRead(in XmlParserInputBufferPtr, len c.Int) c.Int
//go:linkname XmlParserInputBufferGrow C.xmlParserInputBufferGrow
func XmlParserInputBufferGrow(in XmlParserInputBufferPtr, len c.Int) c.Int
//go:linkname XmlParserInputBufferPush C.xmlParserInputBufferPush
func XmlParserInputBufferPush(in XmlParserInputBufferPtr, len c.Int, buf *int8) c.Int
//go:linkname XmlFreeParserInputBuffer C.xmlFreeParserInputBuffer
func XmlFreeParserInputBuffer(in XmlParserInputBufferPtr)
//go:linkname XmlParserGetDirectory C.xmlParserGetDirectory
func XmlParserGetDirectory(filename *int8) *int8
//go:linkname XmlRegisterInputCallbacks C.xmlRegisterInputCallbacks
func XmlRegisterInputCallbacks(matchFunc XmlInputMatchCallback, openFunc XmlInputOpenCallback, readFunc XmlInputReadCallback, closeFunc XmlInputCloseCallback) c.Int
//go:linkname X__XmlParserInputBufferCreateFilename C.__xmlParserInputBufferCreateFilename
func X__XmlParserInputBufferCreateFilename(URI *int8, enc XmlCharEncoding) XmlParserInputBufferPtr
//go:linkname XmlCleanupOutputCallbacks C.xmlCleanupOutputCallbacks
func XmlCleanupOutputCallbacks()
//go:linkname XmlPopOutputCallbacks C.xmlPopOutputCallbacks
func XmlPopOutputCallbacks() c.Int
//go:linkname XmlRegisterDefaultOutputCallbacks C.xmlRegisterDefaultOutputCallbacks
func XmlRegisterDefaultOutputCallbacks()
//go:linkname XmlAllocOutputBuffer C.xmlAllocOutputBuffer
func XmlAllocOutputBuffer(encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr
//go:linkname XmlOutputBufferCreateFilename C.xmlOutputBufferCreateFilename
func XmlOutputBufferCreateFilename(URI *int8, encoder XmlCharEncodingHandlerPtr, compression c.Int) XmlOutputBufferPtr
//go:linkname XmlOutputBufferCreateFile C.xmlOutputBufferCreateFile
func XmlOutputBufferCreateFile(file *c.FILE, encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr
//go:linkname XmlOutputBufferCreateBuffer C.xmlOutputBufferCreateBuffer
func XmlOutputBufferCreateBuffer(buffer XmlBufferPtr, encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr
//go:linkname XmlOutputBufferCreateFd C.xmlOutputBufferCreateFd
func XmlOutputBufferCreateFd(fd c.Int, encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr
//go:linkname XmlOutputBufferCreateIO C.xmlOutputBufferCreateIO
func XmlOutputBufferCreateIO(iowrite XmlOutputWriteCallback, ioclose XmlOutputCloseCallback, ioctx unsafe.Pointer, encoder XmlCharEncodingHandlerPtr) XmlOutputBufferPtr
//go:linkname XmlOutputBufferGetContent C.xmlOutputBufferGetContent
func XmlOutputBufferGetContent(out XmlOutputBufferPtr) *XmlChar
//go:linkname XmlOutputBufferGetSize C.xmlOutputBufferGetSize
func XmlOutputBufferGetSize(out XmlOutputBufferPtr) uintptr
//go:linkname XmlOutputBufferWrite C.xmlOutputBufferWrite
func XmlOutputBufferWrite(out XmlOutputBufferPtr, len c.Int, buf *int8) c.Int
//go:linkname XmlOutputBufferWriteString C.xmlOutputBufferWriteString
func XmlOutputBufferWriteString(out XmlOutputBufferPtr, str *int8) c.Int
//go:linkname XmlOutputBufferWriteEscape C.xmlOutputBufferWriteEscape
func XmlOutputBufferWriteEscape(out XmlOutputBufferPtr, str *XmlChar, escaping XmlCharEncodingOutputFunc) c.Int
//go:linkname XmlOutputBufferFlush C.xmlOutputBufferFlush
func XmlOutputBufferFlush(out XmlOutputBufferPtr) c.Int
//go:linkname XmlOutputBufferClose C.xmlOutputBufferClose
func XmlOutputBufferClose(out XmlOutputBufferPtr) c.Int
//go:linkname XmlRegisterOutputCallbacks C.xmlRegisterOutputCallbacks
func XmlRegisterOutputCallbacks(matchFunc XmlOutputMatchCallback, openFunc XmlOutputOpenCallback, writeFunc XmlOutputWriteCallback, closeFunc XmlOutputCloseCallback) c.Int
//go:linkname X__XmlOutputBufferCreateFilename C.__xmlOutputBufferCreateFilename
func X__XmlOutputBufferCreateFilename(URI *int8, encoder XmlCharEncodingHandlerPtr, compression c.Int) XmlOutputBufferPtr
//go:linkname XmlRegisterHTTPPostCallbacks C.xmlRegisterHTTPPostCallbacks
func XmlRegisterHTTPPostCallbacks()
//go:linkname XmlCheckHTTPInput C.xmlCheckHTTPInput
func XmlCheckHTTPInput(ctxt XmlParserCtxtPtr, ret XmlParserInputPtr) XmlParserInputPtr
//go:linkname XmlNoNetExternalEntityLoader C.xmlNoNetExternalEntityLoader
func XmlNoNetExternalEntityLoader(URL *int8, ID *int8, ctxt XmlParserCtxtPtr) XmlParserInputPtr
// llgo:link (*XmlChar).XmlNormalizeWindowsPath C.xmlNormalizeWindowsPath
func (recv_ *XmlChar) XmlNormalizeWindowsPath() *XmlChar {
	return nil
}
//go:linkname XmlCheckFilename C.xmlCheckFilename
func XmlCheckFilename(path *int8) c.Int
/**
 * Default 'file://' protocol callbacks
 */
//go:linkname XmlFileMatch C.xmlFileMatch
func XmlFileMatch(filename *int8) c.Int
//go:linkname XmlFileOpen C.xmlFileOpen
func XmlFileOpen(filename *int8) unsafe.Pointer
//go:linkname XmlFileRead C.xmlFileRead
func XmlFileRead(context unsafe.Pointer, buffer *int8, len c.Int) c.Int
//go:linkname XmlFileClose C.xmlFileClose
func XmlFileClose(context unsafe.Pointer) c.Int
//go:linkname XmlIOHTTPMatch C.xmlIOHTTPMatch
func XmlIOHTTPMatch(filename *int8) c.Int
//go:linkname XmlIOHTTPOpen C.xmlIOHTTPOpen
func XmlIOHTTPOpen(filename *int8) unsafe.Pointer
//go:linkname XmlIOHTTPOpenW C.xmlIOHTTPOpenW
func XmlIOHTTPOpenW(post_uri *int8, compression c.Int) unsafe.Pointer
//go:linkname XmlIOHTTPRead C.xmlIOHTTPRead
func XmlIOHTTPRead(context unsafe.Pointer, buffer *int8, len c.Int) c.Int
//go:linkname XmlIOHTTPClose C.xmlIOHTTPClose
func XmlIOHTTPClose(context unsafe.Pointer) c.Int
//go:linkname XmlParserInputBufferCreateFilenameDefault C.xmlParserInputBufferCreateFilenameDefault
func XmlParserInputBufferCreateFilenameDefault(func_ XmlParserInputBufferCreateFilenameFunc) XmlParserInputBufferCreateFilenameFunc
//go:linkname XmlOutputBufferCreateFilenameDefault C.xmlOutputBufferCreateFilenameDefault
func XmlOutputBufferCreateFilenameDefault(func_ XmlOutputBufferCreateFilenameFunc) XmlOutputBufferCreateFilenameFunc
//go:linkname XmlThrDefOutputBufferCreateFilenameDefault C.xmlThrDefOutputBufferCreateFilenameDefault
func XmlThrDefOutputBufferCreateFilenameDefault(func_ XmlOutputBufferCreateFilenameFunc) XmlOutputBufferCreateFilenameFunc
//go:linkname XmlThrDefParserInputBufferCreateFilenameDefault C.xmlThrDefParserInputBufferCreateFilenameDefault
func XmlThrDefParserInputBufferCreateFilenameDefault(func_ XmlParserInputBufferCreateFilenameFunc) XmlParserInputBufferCreateFilenameFunc
