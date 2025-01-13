package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XlinkHRef *XmlChar
type XlinkRole *XmlChar
type XlinkTitle *XmlChar
type XlinkType c.Int

const (
	XlinkTypeXLINKTYPENONE        XlinkType = 0
	XlinkTypeXLINKTYPESIMPLE      XlinkType = 1
	XlinkTypeXLINKTYPEEXTENDED    XlinkType = 2
	XlinkTypeXLINKTYPEEXTENDEDSET XlinkType = 3
)

type XlinkShow c.Int

const (
	XlinkShowXLINKSHOWNONE    XlinkShow = 0
	XlinkShowXLINKSHOWNEW     XlinkShow = 1
	XlinkShowXLINKSHOWEMBED   XlinkShow = 2
	XlinkShowXLINKSHOWREPLACE XlinkShow = 3
)

type XlinkActuate c.Int

const (
	XlinkActuateXLINKACTUATENONE      XlinkActuate = 0
	XlinkActuateXLINKACTUATEAUTO      XlinkActuate = 1
	XlinkActuateXLINKACTUATEONREQUEST XlinkActuate = 2
)
// llgo:type C
type XlinkNodeDetectFunc func(unsafe.Pointer, XmlNodePtr)
// llgo:type C
type XlinkSimpleLinkFunk func(unsafe.Pointer, XmlNodePtr, XlinkHRef, XlinkRole, XlinkTitle)
// llgo:type C
type XlinkExtendedLinkFunk func(unsafe.Pointer, XmlNodePtr, c.Int, *XlinkHRef, *XlinkRole, c.Int, *XlinkRole, *XlinkRole, *XlinkShow, *XlinkActuate, c.Int, *XlinkTitle, **XmlChar)
// llgo:type C
type XlinkExtendedLinkSetFunk func(unsafe.Pointer, XmlNodePtr, c.Int, *XlinkHRef, *XlinkRole, c.Int, *XlinkTitle, **XmlChar)

type X_XlinkHandler struct {
	Simple   unsafe.Pointer
	Extended unsafe.Pointer
	Set      unsafe.Pointer
}
type XlinkHandler X_XlinkHandler
type XlinkHandlerPtr *XlinkHandler
//go:linkname XlinkGetDefaultDetect C.xlinkGetDefaultDetect
func XlinkGetDefaultDetect() XlinkNodeDetectFunc
//go:linkname XlinkSetDefaultDetect C.xlinkSetDefaultDetect
func XlinkSetDefaultDetect(func_ XlinkNodeDetectFunc)
//go:linkname XlinkGetDefaultHandler C.xlinkGetDefaultHandler
func XlinkGetDefaultHandler() XlinkHandlerPtr
//go:linkname XlinkSetDefaultHandler C.xlinkSetDefaultHandler
func XlinkSetDefaultHandler(handler XlinkHandlerPtr)
//go:linkname XlinkIsLink C.xlinkIsLink
func XlinkIsLink(doc XmlDocPtr, node XmlNodePtr) XlinkType
