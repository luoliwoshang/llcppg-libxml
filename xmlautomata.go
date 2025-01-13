package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_XmlAutomata struct {
	Unused [8]uint8
}
type XmlAutomata X_XmlAutomata
type XmlAutomataPtr *XmlAutomata

type X_XmlAutomataState struct {
	Unused [8]uint8
}
type XmlAutomataState X_XmlAutomataState
type XmlAutomataStatePtr *XmlAutomataState
//go:linkname XmlNewAutomata C.xmlNewAutomata
func XmlNewAutomata() XmlAutomataPtr
//go:linkname XmlFreeAutomata C.xmlFreeAutomata
func XmlFreeAutomata(am XmlAutomataPtr)
//go:linkname XmlAutomataGetInitState C.xmlAutomataGetInitState
func XmlAutomataGetInitState(am XmlAutomataPtr) XmlAutomataStatePtr
//go:linkname XmlAutomataSetFinalState C.xmlAutomataSetFinalState
func XmlAutomataSetFinalState(am XmlAutomataPtr, state XmlAutomataStatePtr) c.Int
//go:linkname XmlAutomataNewState C.xmlAutomataNewState
func XmlAutomataNewState(am XmlAutomataPtr) XmlAutomataStatePtr
//go:linkname XmlAutomataNewTransition C.xmlAutomataNewTransition
func XmlAutomataNewTransition(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, token *XmlChar, data unsafe.Pointer) XmlAutomataStatePtr
//go:linkname XmlAutomataNewTransition2 C.xmlAutomataNewTransition2
func XmlAutomataNewTransition2(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, token *XmlChar, token2 *XmlChar, data unsafe.Pointer) XmlAutomataStatePtr
//go:linkname XmlAutomataNewNegTrans C.xmlAutomataNewNegTrans
func XmlAutomataNewNegTrans(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, token *XmlChar, token2 *XmlChar, data unsafe.Pointer) XmlAutomataStatePtr
//go:linkname XmlAutomataNewCountTrans C.xmlAutomataNewCountTrans
func XmlAutomataNewCountTrans(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, token *XmlChar, min c.Int, max c.Int, data unsafe.Pointer) XmlAutomataStatePtr
//go:linkname XmlAutomataNewCountTrans2 C.xmlAutomataNewCountTrans2
func XmlAutomataNewCountTrans2(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, token *XmlChar, token2 *XmlChar, min c.Int, max c.Int, data unsafe.Pointer) XmlAutomataStatePtr
//go:linkname XmlAutomataNewOnceTrans C.xmlAutomataNewOnceTrans
func XmlAutomataNewOnceTrans(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, token *XmlChar, min c.Int, max c.Int, data unsafe.Pointer) XmlAutomataStatePtr
//go:linkname XmlAutomataNewOnceTrans2 C.xmlAutomataNewOnceTrans2
func XmlAutomataNewOnceTrans2(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, token *XmlChar, token2 *XmlChar, min c.Int, max c.Int, data unsafe.Pointer) XmlAutomataStatePtr
//go:linkname XmlAutomataNewAllTrans C.xmlAutomataNewAllTrans
func XmlAutomataNewAllTrans(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, lax c.Int) XmlAutomataStatePtr
//go:linkname XmlAutomataNewEpsilon C.xmlAutomataNewEpsilon
func XmlAutomataNewEpsilon(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr) XmlAutomataStatePtr
//go:linkname XmlAutomataNewCountedTrans C.xmlAutomataNewCountedTrans
func XmlAutomataNewCountedTrans(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, counter c.Int) XmlAutomataStatePtr
//go:linkname XmlAutomataNewCounterTrans C.xmlAutomataNewCounterTrans
func XmlAutomataNewCounterTrans(am XmlAutomataPtr, from XmlAutomataStatePtr, to XmlAutomataStatePtr, counter c.Int) XmlAutomataStatePtr
//go:linkname XmlAutomataNewCounter C.xmlAutomataNewCounter
func XmlAutomataNewCounter(am XmlAutomataPtr, min c.Int, max c.Int) c.Int
//go:linkname XmlAutomataCompile C.xmlAutomataCompile
func XmlAutomataCompile(am XmlAutomataPtr) *X_XmlRegexp
//go:linkname XmlAutomataIsDeterminist C.xmlAutomataIsDeterminist
func XmlAutomataIsDeterminist(am XmlAutomataPtr) c.Int
