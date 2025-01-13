package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type UCharIterator struct {
	Context       unsafe.Pointer
	Length        int32
	Start         int32
	Index         int32
	Limit         int32
	ReservedField int32
	GetIndex      *unsafe.Pointer
	Move          *unsafe.Pointer
	HasNext       *unsafe.Pointer
	HasPrevious   *unsafe.Pointer
	Current       *unsafe.Pointer
	Next          *unsafe.Pointer
	Previous      *unsafe.Pointer
	ReservedFn    *unsafe.Pointer
	GetState      *unsafe.Pointer
	SetState      *unsafe.Pointer
}
type UCharIteratorOrigin c.Int

const (
	UCharIteratorOriginUITERSTART   UCharIteratorOrigin = 0
	UCharIteratorOriginUITERCURRENT UCharIteratorOrigin = 1
	UCharIteratorOriginUITERLIMIT   UCharIteratorOrigin = 2
	UCharIteratorOriginUITERZERO    UCharIteratorOrigin = 3
	UCharIteratorOriginUITERLENGTH  UCharIteratorOrigin = 4
)
const UITERUNKNOWNINDEX c.Int = -2
// llgo:type C
type UCharIteratorGetIndex func(*UCharIterator, UCharIteratorOrigin) int32
// llgo:type C
type UCharIteratorMove func(*UCharIterator, int32, UCharIteratorOrigin) int32
// llgo:type C
type UCharIteratorHasNext func(*UCharIterator) UBool
// llgo:type C
type UCharIteratorHasPrevious func(*UCharIterator) UBool
// llgo:type C
type UCharIteratorCurrent func(*UCharIterator) UChar32
// llgo:type C
type UCharIteratorNext func(*UCharIterator) UChar32
// llgo:type C
type UCharIteratorPrevious func(*UCharIterator) UChar32
// llgo:type C
type UCharIteratorReserved func(*UCharIterator, int32) int32
// llgo:type C
type UCharIteratorGetState func(*UCharIterator) uint32
// llgo:type C
type UCharIteratorSetState func(*UCharIterator, uint32, *UErrorCode)
