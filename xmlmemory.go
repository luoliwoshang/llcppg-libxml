package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
// llgo:type C
type XmlFreeFunc func(unsafe.Pointer)
// llgo:type C
type XmlMallocFunc func(uintptr) unsafe.Pointer
// llgo:type C
type XmlReallocFunc func(unsafe.Pointer, uintptr) unsafe.Pointer
// llgo:type C
type XmlStrdupFunc func(*int8) *int8
//go:linkname XmlMemSetup C.xmlMemSetup
func XmlMemSetup(freeFunc XmlFreeFunc, mallocFunc XmlMallocFunc, reallocFunc XmlReallocFunc, strdupFunc XmlStrdupFunc) c.Int
//go:linkname XmlMemGet C.xmlMemGet
func XmlMemGet(freeFunc XmlFreeFunc, mallocFunc XmlMallocFunc, reallocFunc XmlReallocFunc, strdupFunc XmlStrdupFunc) c.Int
//go:linkname XmlGcMemSetup C.xmlGcMemSetup
func XmlGcMemSetup(freeFunc XmlFreeFunc, mallocFunc XmlMallocFunc, mallocAtomicFunc XmlMallocFunc, reallocFunc XmlReallocFunc, strdupFunc XmlStrdupFunc) c.Int
//go:linkname XmlGcMemGet C.xmlGcMemGet
func XmlGcMemGet(freeFunc XmlFreeFunc, mallocFunc XmlMallocFunc, mallocAtomicFunc XmlMallocFunc, reallocFunc XmlReallocFunc, strdupFunc XmlStrdupFunc) c.Int
//go:linkname XmlInitMemory C.xmlInitMemory
func XmlInitMemory() c.Int
//go:linkname XmlCleanupMemory C.xmlCleanupMemory
func XmlCleanupMemory()
//go:linkname XmlMemSize C.xmlMemSize
func XmlMemSize(ptr unsafe.Pointer) uintptr
//go:linkname XmlMemUsed C.xmlMemUsed
func XmlMemUsed() c.Int
//go:linkname XmlMemBlocks C.xmlMemBlocks
func XmlMemBlocks() c.Int
//go:linkname XmlMemDisplay C.xmlMemDisplay
func XmlMemDisplay(fp *c.FILE)
//go:linkname XmlMemDisplayLast C.xmlMemDisplayLast
func XmlMemDisplayLast(fp *c.FILE, nbBytes c.Long)
//go:linkname XmlMemShow C.xmlMemShow
func XmlMemShow(fp *c.FILE, nr c.Int)
//go:linkname XmlMemoryDump C.xmlMemoryDump
func XmlMemoryDump()
//go:linkname XmlMemMalloc C.xmlMemMalloc
func XmlMemMalloc(size uintptr) unsafe.Pointer
//go:linkname XmlMemRealloc C.xmlMemRealloc
func XmlMemRealloc(ptr unsafe.Pointer, size uintptr) unsafe.Pointer
//go:linkname XmlMemFree C.xmlMemFree
func XmlMemFree(ptr unsafe.Pointer)
//go:linkname XmlMemoryStrdup C.xmlMemoryStrdup
func XmlMemoryStrdup(str *int8) *int8
//go:linkname XmlMallocLoc C.xmlMallocLoc
func XmlMallocLoc(size uintptr, file *int8, line c.Int) unsafe.Pointer
//go:linkname XmlReallocLoc C.xmlReallocLoc
func XmlReallocLoc(ptr unsafe.Pointer, size uintptr, file *int8, line c.Int) unsafe.Pointer
//go:linkname XmlMallocAtomicLoc C.xmlMallocAtomicLoc
func XmlMallocAtomicLoc(size uintptr, file *int8, line c.Int) unsafe.Pointer
//go:linkname XmlMemStrdupLoc C.xmlMemStrdupLoc
func XmlMemStrdupLoc(str *int8, file *int8, line c.Int) *int8
