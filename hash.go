package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_XmlHashTable struct {
	Unused [8]uint8
}
type XmlHashTable X_XmlHashTable
type XmlHashTablePtr *XmlHashTable
// llgo:type C
type XmlHashDeallocator func(unsafe.Pointer, *XmlChar)
// llgo:type C
type XmlHashCopier func(unsafe.Pointer, *XmlChar) unsafe.Pointer
// llgo:type C
type XmlHashScanner func(unsafe.Pointer, unsafe.Pointer, *XmlChar)
// llgo:type C
type XmlHashScannerFull func(unsafe.Pointer, unsafe.Pointer, *XmlChar, *XmlChar, *XmlChar)
//go:linkname XmlHashCreate C.xmlHashCreate
func XmlHashCreate(size c.Int) XmlHashTablePtr
//go:linkname XmlHashCreateDict C.xmlHashCreateDict
func XmlHashCreateDict(size c.Int, dict XmlDictPtr) XmlHashTablePtr
//go:linkname XmlHashFree C.xmlHashFree
func XmlHashFree(hash XmlHashTablePtr, dealloc XmlHashDeallocator)
//go:linkname XmlHashDefaultDeallocator C.xmlHashDefaultDeallocator
func XmlHashDefaultDeallocator(entry unsafe.Pointer, name *XmlChar)
//go:linkname XmlHashAdd C.xmlHashAdd
func XmlHashAdd(hash XmlHashTablePtr, name *XmlChar, userdata unsafe.Pointer) c.Int
//go:linkname XmlHashAddEntry C.xmlHashAddEntry
func XmlHashAddEntry(hash XmlHashTablePtr, name *XmlChar, userdata unsafe.Pointer) c.Int
//go:linkname XmlHashUpdateEntry C.xmlHashUpdateEntry
func XmlHashUpdateEntry(hash XmlHashTablePtr, name *XmlChar, userdata unsafe.Pointer, dealloc XmlHashDeallocator) c.Int
//go:linkname XmlHashAdd2 C.xmlHashAdd2
func XmlHashAdd2(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, userdata unsafe.Pointer) c.Int
//go:linkname XmlHashAddEntry2 C.xmlHashAddEntry2
func XmlHashAddEntry2(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, userdata unsafe.Pointer) c.Int
//go:linkname XmlHashUpdateEntry2 C.xmlHashUpdateEntry2
func XmlHashUpdateEntry2(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, userdata unsafe.Pointer, dealloc XmlHashDeallocator) c.Int
//go:linkname XmlHashAdd3 C.xmlHashAdd3
func XmlHashAdd3(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, name3 *XmlChar, userdata unsafe.Pointer) c.Int
//go:linkname XmlHashAddEntry3 C.xmlHashAddEntry3
func XmlHashAddEntry3(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, name3 *XmlChar, userdata unsafe.Pointer) c.Int
//go:linkname XmlHashUpdateEntry3 C.xmlHashUpdateEntry3
func XmlHashUpdateEntry3(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, name3 *XmlChar, userdata unsafe.Pointer, dealloc XmlHashDeallocator) c.Int
//go:linkname XmlHashRemoveEntry C.xmlHashRemoveEntry
func XmlHashRemoveEntry(hash XmlHashTablePtr, name *XmlChar, dealloc XmlHashDeallocator) c.Int
//go:linkname XmlHashRemoveEntry2 C.xmlHashRemoveEntry2
func XmlHashRemoveEntry2(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, dealloc XmlHashDeallocator) c.Int
//go:linkname XmlHashRemoveEntry3 C.xmlHashRemoveEntry3
func XmlHashRemoveEntry3(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, name3 *XmlChar, dealloc XmlHashDeallocator) c.Int
//go:linkname XmlHashLookup C.xmlHashLookup
func XmlHashLookup(hash XmlHashTablePtr, name *XmlChar) unsafe.Pointer
//go:linkname XmlHashLookup2 C.xmlHashLookup2
func XmlHashLookup2(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar) unsafe.Pointer
//go:linkname XmlHashLookup3 C.xmlHashLookup3
func XmlHashLookup3(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, name3 *XmlChar) unsafe.Pointer
//go:linkname XmlHashQLookup C.xmlHashQLookup
func XmlHashQLookup(hash XmlHashTablePtr, prefix *XmlChar, name *XmlChar) unsafe.Pointer
//go:linkname XmlHashQLookup2 C.xmlHashQLookup2
func XmlHashQLookup2(hash XmlHashTablePtr, prefix *XmlChar, name *XmlChar, prefix2 *XmlChar, name2 *XmlChar) unsafe.Pointer
//go:linkname XmlHashQLookup3 C.xmlHashQLookup3
func XmlHashQLookup3(hash XmlHashTablePtr, prefix *XmlChar, name *XmlChar, prefix2 *XmlChar, name2 *XmlChar, prefix3 *XmlChar, name3 *XmlChar) unsafe.Pointer
//go:linkname XmlHashCopySafe C.xmlHashCopySafe
func XmlHashCopySafe(hash XmlHashTablePtr, copy XmlHashCopier, dealloc XmlHashDeallocator) XmlHashTablePtr
//go:linkname XmlHashCopy C.xmlHashCopy
func XmlHashCopy(hash XmlHashTablePtr, copy XmlHashCopier) XmlHashTablePtr
//go:linkname XmlHashSize C.xmlHashSize
func XmlHashSize(hash XmlHashTablePtr) c.Int
//go:linkname XmlHashScan C.xmlHashScan
func XmlHashScan(hash XmlHashTablePtr, scan XmlHashScanner, data unsafe.Pointer)
//go:linkname XmlHashScan3 C.xmlHashScan3
func XmlHashScan3(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, name3 *XmlChar, scan XmlHashScanner, data unsafe.Pointer)
//go:linkname XmlHashScanFull C.xmlHashScanFull
func XmlHashScanFull(hash XmlHashTablePtr, scan XmlHashScannerFull, data unsafe.Pointer)
//go:linkname XmlHashScanFull3 C.xmlHashScanFull3
func XmlHashScanFull3(hash XmlHashTablePtr, name *XmlChar, name2 *XmlChar, name3 *XmlChar, scan XmlHashScannerFull, data unsafe.Pointer)
