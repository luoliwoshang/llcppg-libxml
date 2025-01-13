package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type X_XmlModule struct {
	Unused [8]uint8
}
type XmlModule X_XmlModule
type XmlModulePtr *XmlModule
type XmlModuleOption c.Int

const (
	XmlModuleOptionXMLMODULELAZY  XmlModuleOption = 1
	XmlModuleOptionXMLMODULELOCAL XmlModuleOption = 2
)
//go:linkname XmlModuleOpen C.xmlModuleOpen
func XmlModuleOpen(filename *int8, options c.Int) XmlModulePtr
//go:linkname XmlModuleSymbol C.xmlModuleSymbol
func XmlModuleSymbol(module XmlModulePtr, name *int8, result *unsafe.Pointer) c.Int
//go:linkname XmlModuleClose C.xmlModuleClose
func XmlModuleClose(module XmlModulePtr) c.Int
//go:linkname XmlModuleFree C.xmlModuleFree
func XmlModuleFree(module XmlModulePtr) c.Int
