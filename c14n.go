package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XmlC14NMode c.Int

const (
	XmlC14NModeXMLC14N10          XmlC14NMode = 0
	XmlC14NModeXMLC14NEXCLUSIVE10 XmlC14NMode = 1
	XmlC14NModeXMLC14N11          XmlC14NMode = 2
)
//go:linkname XmlC14NDocSaveTo C.xmlC14NDocSaveTo
func XmlC14NDocSaveTo(doc XmlDocPtr, nodes XmlNodeSetPtr, mode c.Int, inclusive_ns_prefixes **XmlChar, with_comments c.Int, buf XmlOutputBufferPtr) c.Int
//go:linkname XmlC14NDocDumpMemory C.xmlC14NDocDumpMemory
func XmlC14NDocDumpMemory(doc XmlDocPtr, nodes XmlNodeSetPtr, mode c.Int, inclusive_ns_prefixes **XmlChar, with_comments c.Int, doc_txt_ptr **XmlChar) c.Int
//go:linkname XmlC14NDocSave C.xmlC14NDocSave
func XmlC14NDocSave(doc XmlDocPtr, nodes XmlNodeSetPtr, mode c.Int, inclusive_ns_prefixes **XmlChar, with_comments c.Int, filename *int8, compression c.Int) c.Int
// llgo:type C
type XmlC14NIsVisibleCallback func(unsafe.Pointer, XmlNodePtr, XmlNodePtr) c.Int
//go:linkname XmlC14NExecute C.xmlC14NExecute
func XmlC14NExecute(doc XmlDocPtr, is_visible_callback XmlC14NIsVisibleCallback, user_data unsafe.Pointer, mode c.Int, inclusive_ns_prefixes **XmlChar, with_comments c.Int, buf XmlOutputBufferPtr) c.Int
