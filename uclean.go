package libxml_2_0

import "unsafe"
// llgo:type C
type UMemAllocFn func(unsafe.Pointer, uintptr) unsafe.Pointer
// llgo:type C
type UMemReallocFn func(unsafe.Pointer, unsafe.Pointer, uintptr) unsafe.Pointer
// llgo:type C
type UMemFreeFn func(unsafe.Pointer, unsafe.Pointer)
type UMTX unsafe.Pointer
// llgo:type C
type UMtxInitFn func(unsafe.Pointer, *UMTX, *UErrorCode)
// llgo:type C
type UMtxFn func(unsafe.Pointer, *UMTX)
// llgo:type C
type UMtxAtomicFn func(unsafe.Pointer, *int32) int32
