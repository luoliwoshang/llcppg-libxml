package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_XmlLink struct {
	Unused [8]uint8
}
type XmlLink X_XmlLink
type XmlLinkPtr *XmlLink

type X_XmlList struct {
	Unused [8]uint8
}
type XmlList X_XmlList
type XmlListPtr *XmlList
// llgo:type C
type XmlListDeallocator func(XmlLinkPtr)
// llgo:type C
type XmlListDataCompare func(unsafe.Pointer, unsafe.Pointer) c.Int
// llgo:type C
type XmlListWalker func(unsafe.Pointer, unsafe.Pointer) c.Int
//go:linkname XmlListCreate C.xmlListCreate
func XmlListCreate(deallocator XmlListDeallocator, compare XmlListDataCompare) XmlListPtr
//go:linkname XmlListDelete C.xmlListDelete
func XmlListDelete(l XmlListPtr)
//go:linkname XmlListSearch C.xmlListSearch
func XmlListSearch(l XmlListPtr, data unsafe.Pointer) unsafe.Pointer
//go:linkname XmlListReverseSearch C.xmlListReverseSearch
func XmlListReverseSearch(l XmlListPtr, data unsafe.Pointer) unsafe.Pointer
//go:linkname XmlListInsert C.xmlListInsert
func XmlListInsert(l XmlListPtr, data unsafe.Pointer) c.Int
//go:linkname XmlListAppend C.xmlListAppend
func XmlListAppend(l XmlListPtr, data unsafe.Pointer) c.Int
//go:linkname XmlListRemoveFirst C.xmlListRemoveFirst
func XmlListRemoveFirst(l XmlListPtr, data unsafe.Pointer) c.Int
//go:linkname XmlListRemoveLast C.xmlListRemoveLast
func XmlListRemoveLast(l XmlListPtr, data unsafe.Pointer) c.Int
//go:linkname XmlListRemoveAll C.xmlListRemoveAll
func XmlListRemoveAll(l XmlListPtr, data unsafe.Pointer) c.Int
//go:linkname XmlListClear C.xmlListClear
func XmlListClear(l XmlListPtr)
//go:linkname XmlListEmpty C.xmlListEmpty
func XmlListEmpty(l XmlListPtr) c.Int
//go:linkname XmlListFront C.xmlListFront
func XmlListFront(l XmlListPtr) XmlLinkPtr
//go:linkname XmlListEnd C.xmlListEnd
func XmlListEnd(l XmlListPtr) XmlLinkPtr
//go:linkname XmlListSize C.xmlListSize
func XmlListSize(l XmlListPtr) c.Int
//go:linkname XmlListPopFront C.xmlListPopFront
func XmlListPopFront(l XmlListPtr)
//go:linkname XmlListPopBack C.xmlListPopBack
func XmlListPopBack(l XmlListPtr)
//go:linkname XmlListPushFront C.xmlListPushFront
func XmlListPushFront(l XmlListPtr, data unsafe.Pointer) c.Int
//go:linkname XmlListPushBack C.xmlListPushBack
func XmlListPushBack(l XmlListPtr, data unsafe.Pointer) c.Int
//go:linkname XmlListReverse C.xmlListReverse
func XmlListReverse(l XmlListPtr)
//go:linkname XmlListSort C.xmlListSort
func XmlListSort(l XmlListPtr)
//go:linkname XmlListWalk C.xmlListWalk
func XmlListWalk(l XmlListPtr, walker XmlListWalker, user unsafe.Pointer)
//go:linkname XmlListReverseWalk C.xmlListReverseWalk
func XmlListReverseWalk(l XmlListPtr, walker XmlListWalker, user unsafe.Pointer)
//go:linkname XmlListMerge C.xmlListMerge
func XmlListMerge(l1 XmlListPtr, l2 XmlListPtr)
//go:linkname XmlListDup C.xmlListDup
func XmlListDup(old XmlListPtr) XmlListPtr
//go:linkname XmlListCopy C.xmlListCopy
func XmlListCopy(cur XmlListPtr, old XmlListPtr) c.Int
//go:linkname XmlLinkGetData C.xmlLinkGetData
func XmlLinkGetData(lk XmlLinkPtr) unsafe.Pointer
