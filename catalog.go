package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XmlCatalogPrefer c.Int

const (
	XmlCatalogPreferXMLCATAPREFERNONE   XmlCatalogPrefer = 0
	XmlCatalogPreferXMLCATAPREFERPUBLIC XmlCatalogPrefer = 1
	XmlCatalogPreferXMLCATAPREFERSYSTEM XmlCatalogPrefer = 2
)

type XmlCatalogAllow c.Int

const (
	XmlCatalogAllowXMLCATAALLOWNONE     XmlCatalogAllow = 0
	XmlCatalogAllowXMLCATAALLOWGLOBAL   XmlCatalogAllow = 1
	XmlCatalogAllowXMLCATAALLOWDOCUMENT XmlCatalogAllow = 2
	XmlCatalogAllowXMLCATAALLOWALL      XmlCatalogAllow = 3
)

type X_XmlCatalog struct {
	Unused [8]uint8
}
type XmlCatalog X_XmlCatalog
type XmlCatalogPtr *XmlCatalog
//go:linkname XmlNewCatalog C.xmlNewCatalog
func XmlNewCatalog(sgml c.Int) XmlCatalogPtr
//go:linkname XmlLoadACatalog C.xmlLoadACatalog
func XmlLoadACatalog(filename *int8) XmlCatalogPtr
//go:linkname XmlLoadSGMLSuperCatalog C.xmlLoadSGMLSuperCatalog
func XmlLoadSGMLSuperCatalog(filename *int8) XmlCatalogPtr
//go:linkname XmlConvertSGMLCatalog C.xmlConvertSGMLCatalog
func XmlConvertSGMLCatalog(catal XmlCatalogPtr) c.Int
//go:linkname XmlACatalogAdd C.xmlACatalogAdd
func XmlACatalogAdd(catal XmlCatalogPtr, type_ *XmlChar, orig *XmlChar, replace *XmlChar) c.Int
//go:linkname XmlACatalogRemove C.xmlACatalogRemove
func XmlACatalogRemove(catal XmlCatalogPtr, value *XmlChar) c.Int
//go:linkname XmlACatalogResolve C.xmlACatalogResolve
func XmlACatalogResolve(catal XmlCatalogPtr, pubID *XmlChar, sysID *XmlChar) *XmlChar
//go:linkname XmlACatalogResolveSystem C.xmlACatalogResolveSystem
func XmlACatalogResolveSystem(catal XmlCatalogPtr, sysID *XmlChar) *XmlChar
//go:linkname XmlACatalogResolvePublic C.xmlACatalogResolvePublic
func XmlACatalogResolvePublic(catal XmlCatalogPtr, pubID *XmlChar) *XmlChar
//go:linkname XmlACatalogResolveURI C.xmlACatalogResolveURI
func XmlACatalogResolveURI(catal XmlCatalogPtr, URI *XmlChar) *XmlChar
//go:linkname XmlACatalogDump C.xmlACatalogDump
func XmlACatalogDump(catal XmlCatalogPtr, out *c.FILE)
//go:linkname XmlFreeCatalog C.xmlFreeCatalog
func XmlFreeCatalog(catal XmlCatalogPtr)
//go:linkname XmlCatalogIsEmpty C.xmlCatalogIsEmpty
func XmlCatalogIsEmpty(catal XmlCatalogPtr) c.Int
//go:linkname XmlInitializeCatalog C.xmlInitializeCatalog
func XmlInitializeCatalog()
//go:linkname XmlLoadCatalog C.xmlLoadCatalog
func XmlLoadCatalog(filename *int8) c.Int
//go:linkname XmlLoadCatalogs C.xmlLoadCatalogs
func XmlLoadCatalogs(paths *int8)
//go:linkname XmlCatalogCleanup C.xmlCatalogCleanup
func XmlCatalogCleanup()
//go:linkname XmlCatalogDump C.xmlCatalogDump
func XmlCatalogDump(out *c.FILE)
// llgo:link (*XmlChar).XmlCatalogResolve C.xmlCatalogResolve
func (recv_ *XmlChar) XmlCatalogResolve(sysID *XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlCatalogResolveSystem C.xmlCatalogResolveSystem
func (recv_ *XmlChar) XmlCatalogResolveSystem() *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlCatalogResolvePublic C.xmlCatalogResolvePublic
func (recv_ *XmlChar) XmlCatalogResolvePublic() *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlCatalogResolveURI C.xmlCatalogResolveURI
func (recv_ *XmlChar) XmlCatalogResolveURI() *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlCatalogAdd C.xmlCatalogAdd
func (recv_ *XmlChar) XmlCatalogAdd(orig *XmlChar, replace *XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlCatalogRemove C.xmlCatalogRemove
func (recv_ *XmlChar) XmlCatalogRemove() c.Int {
	return 0
}
//go:linkname XmlParseCatalogFile C.xmlParseCatalogFile
func XmlParseCatalogFile(filename *int8) XmlDocPtr
//go:linkname XmlCatalogConvert C.xmlCatalogConvert
func XmlCatalogConvert() c.Int
//go:linkname XmlCatalogFreeLocal C.xmlCatalogFreeLocal
func XmlCatalogFreeLocal(catalogs unsafe.Pointer)
//go:linkname XmlCatalogAddLocal C.xmlCatalogAddLocal
func XmlCatalogAddLocal(catalogs unsafe.Pointer, URL *XmlChar) unsafe.Pointer
//go:linkname XmlCatalogLocalResolve C.xmlCatalogLocalResolve
func XmlCatalogLocalResolve(catalogs unsafe.Pointer, pubID *XmlChar, sysID *XmlChar) *XmlChar
//go:linkname XmlCatalogLocalResolveURI C.xmlCatalogLocalResolveURI
func XmlCatalogLocalResolveURI(catalogs unsafe.Pointer, URI *XmlChar) *XmlChar
//go:linkname XmlCatalogSetDebug C.xmlCatalogSetDebug
func XmlCatalogSetDebug(level c.Int) c.Int
// llgo:link XmlCatalogPrefer.XmlCatalogSetDefaultPrefer C.xmlCatalogSetDefaultPrefer
func (recv_ XmlCatalogPrefer) XmlCatalogSetDefaultPrefer() XmlCatalogPrefer {
	return 0
}
// llgo:link XmlCatalogAllow.XmlCatalogSetDefaults C.xmlCatalogSetDefaults
func (recv_ XmlCatalogAllow) XmlCatalogSetDefaults() {
}
//go:linkname XmlCatalogGetDefaults C.xmlCatalogGetDefaults
func XmlCatalogGetDefaults() XmlCatalogAllow
// llgo:link (*XmlChar).XmlCatalogGetSystem C.xmlCatalogGetSystem
func (recv_ *XmlChar) XmlCatalogGetSystem() *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlCatalogGetPublic C.xmlCatalogGetPublic
func (recv_ *XmlChar) XmlCatalogGetPublic() *XmlChar {
	return nil
}
