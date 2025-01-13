package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
//go:linkname XmlNanoHTTPInit C.xmlNanoHTTPInit
func XmlNanoHTTPInit()
//go:linkname XmlNanoHTTPCleanup C.xmlNanoHTTPCleanup
func XmlNanoHTTPCleanup()
//go:linkname XmlNanoHTTPScanProxy C.xmlNanoHTTPScanProxy
func XmlNanoHTTPScanProxy(URL *int8)
//go:linkname XmlNanoHTTPFetch C.xmlNanoHTTPFetch
func XmlNanoHTTPFetch(URL *int8, filename *int8, contentType **int8) c.Int
//go:linkname XmlNanoHTTPMethod C.xmlNanoHTTPMethod
func XmlNanoHTTPMethod(URL *int8, method *int8, input *int8, contentType **int8, headers *int8, ilen c.Int) unsafe.Pointer
//go:linkname XmlNanoHTTPMethodRedir C.xmlNanoHTTPMethodRedir
func XmlNanoHTTPMethodRedir(URL *int8, method *int8, input *int8, contentType **int8, redir **int8, headers *int8, ilen c.Int) unsafe.Pointer
//go:linkname XmlNanoHTTPOpen C.xmlNanoHTTPOpen
func XmlNanoHTTPOpen(URL *int8, contentType **int8) unsafe.Pointer
//go:linkname XmlNanoHTTPOpenRedir C.xmlNanoHTTPOpenRedir
func XmlNanoHTTPOpenRedir(URL *int8, contentType **int8, redir **int8) unsafe.Pointer
//go:linkname XmlNanoHTTPReturnCode C.xmlNanoHTTPReturnCode
func XmlNanoHTTPReturnCode(ctx unsafe.Pointer) c.Int
//go:linkname XmlNanoHTTPAuthHeader C.xmlNanoHTTPAuthHeader
func XmlNanoHTTPAuthHeader(ctx unsafe.Pointer) *int8
//go:linkname XmlNanoHTTPRedir C.xmlNanoHTTPRedir
func XmlNanoHTTPRedir(ctx unsafe.Pointer) *int8
//go:linkname XmlNanoHTTPContentLength C.xmlNanoHTTPContentLength
func XmlNanoHTTPContentLength(ctx unsafe.Pointer) c.Int
//go:linkname XmlNanoHTTPEncoding C.xmlNanoHTTPEncoding
func XmlNanoHTTPEncoding(ctx unsafe.Pointer) *int8
//go:linkname XmlNanoHTTPMimeType C.xmlNanoHTTPMimeType
func XmlNanoHTTPMimeType(ctx unsafe.Pointer) *int8
//go:linkname XmlNanoHTTPRead C.xmlNanoHTTPRead
func XmlNanoHTTPRead(ctx unsafe.Pointer, dest unsafe.Pointer, len c.Int) c.Int
//go:linkname XmlNanoHTTPSave C.xmlNanoHTTPSave
func XmlNanoHTTPSave(ctxt unsafe.Pointer, filename *int8) c.Int
//go:linkname XmlNanoHTTPClose C.xmlNanoHTTPClose
func XmlNanoHTTPClose(ctx unsafe.Pointer)
