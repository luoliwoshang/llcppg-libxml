package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type X_XmlMutex struct {
	Unused [8]uint8
}
type XmlMutex X_XmlMutex
type XmlMutexPtr *XmlMutex

type X_XmlRMutex struct {
	Unused [8]uint8
}
type XmlRMutex X_XmlRMutex
type XmlRMutexPtr *XmlRMutex
//go:linkname XmlCheckThreadLocalStorage C.xmlCheckThreadLocalStorage
func XmlCheckThreadLocalStorage() c.Int
//go:linkname XmlNewMutex C.xmlNewMutex
func XmlNewMutex() XmlMutexPtr
//go:linkname XmlMutexLock C.xmlMutexLock
func XmlMutexLock(tok XmlMutexPtr)
//go:linkname XmlMutexUnlock C.xmlMutexUnlock
func XmlMutexUnlock(tok XmlMutexPtr)
//go:linkname XmlFreeMutex C.xmlFreeMutex
func XmlFreeMutex(tok XmlMutexPtr)
//go:linkname XmlNewRMutex C.xmlNewRMutex
func XmlNewRMutex() XmlRMutexPtr
//go:linkname XmlRMutexLock C.xmlRMutexLock
func XmlRMutexLock(tok XmlRMutexPtr)
//go:linkname XmlRMutexUnlock C.xmlRMutexUnlock
func XmlRMutexUnlock(tok XmlRMutexPtr)
//go:linkname XmlFreeRMutex C.xmlFreeRMutex
func XmlFreeRMutex(tok XmlRMutexPtr)
//go:linkname XmlInitThreads C.xmlInitThreads
func XmlInitThreads()
//go:linkname XmlLockLibrary C.xmlLockLibrary
func XmlLockLibrary()
//go:linkname XmlUnlockLibrary C.xmlUnlockLibrary
func XmlUnlockLibrary()
//go:linkname XmlGetThreadId C.xmlGetThreadId
func XmlGetThreadId() c.Int
//go:linkname XmlIsMainThread C.xmlIsMainThread
func XmlIsMainThread() c.Int
//go:linkname XmlCleanupThreads C.xmlCleanupThreads
func XmlCleanupThreads()
