package libxml_2_0

import _ "unsafe"

type X_XmlGlobalState struct {
	Unused [8]uint8
}
type XmlGlobalState X_XmlGlobalState
type XmlGlobalStatePtr *XmlGlobalState
//go:linkname XmlInitializeGlobalState C.xmlInitializeGlobalState
func XmlInitializeGlobalState(gs XmlGlobalStatePtr)
//go:linkname XmlGetGlobalState C.xmlGetGlobalState
func XmlGetGlobalState() XmlGlobalStatePtr
