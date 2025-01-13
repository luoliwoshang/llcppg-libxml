package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type XmlEntityType c.Int

const (
	XmlEntityTypeXMLINTERNALGENERALENTITY         XmlEntityType = 1
	XmlEntityTypeXMLEXTERNALGENERALPARSEDENTITY   XmlEntityType = 2
	XmlEntityTypeXMLEXTERNALGENERALUNPARSEDENTITY XmlEntityType = 3
	XmlEntityTypeXMLINTERNALPARAMETERENTITY       XmlEntityType = 4
	XmlEntityTypeXMLEXTERNALPARAMETERENTITY       XmlEntityType = 5
	XmlEntityTypeXMLINTERNALPREDEFINEDENTITY      XmlEntityType = 6
)

type XmlEntitiesTable X_XmlHashTable
type XmlEntitiesTablePtr *XmlEntitiesTable
//go:linkname XmlInitializePredefinedEntities C.xmlInitializePredefinedEntities
func XmlInitializePredefinedEntities()
//go:linkname XmlNewEntity C.xmlNewEntity
func XmlNewEntity(doc XmlDocPtr, name *XmlChar, type_ c.Int, ExternalID *XmlChar, SystemID *XmlChar, content *XmlChar) XmlEntityPtr
//go:linkname XmlFreeEntity C.xmlFreeEntity
func XmlFreeEntity(entity XmlEntityPtr)
//go:linkname XmlAddEntity C.xmlAddEntity
func XmlAddEntity(doc XmlDocPtr, extSubset c.Int, name *XmlChar, type_ c.Int, ExternalID *XmlChar, SystemID *XmlChar, content *XmlChar, out *XmlEntityPtr) c.Int
//go:linkname XmlAddDocEntity C.xmlAddDocEntity
func XmlAddDocEntity(doc XmlDocPtr, name *XmlChar, type_ c.Int, ExternalID *XmlChar, SystemID *XmlChar, content *XmlChar) XmlEntityPtr
//go:linkname XmlAddDtdEntity C.xmlAddDtdEntity
func XmlAddDtdEntity(doc XmlDocPtr, name *XmlChar, type_ c.Int, ExternalID *XmlChar, SystemID *XmlChar, content *XmlChar) XmlEntityPtr
// llgo:link (*XmlChar).XmlGetPredefinedEntity C.xmlGetPredefinedEntity
func (recv_ *XmlChar) XmlGetPredefinedEntity() XmlEntityPtr {
	return nil
}
// llgo:link (*XmlDoc).XmlGetDocEntity C.xmlGetDocEntity
func (recv_ *XmlDoc) XmlGetDocEntity(name *XmlChar) XmlEntityPtr {
	return nil
}
//go:linkname XmlGetDtdEntity C.xmlGetDtdEntity
func XmlGetDtdEntity(doc XmlDocPtr, name *XmlChar) XmlEntityPtr
//go:linkname XmlGetParameterEntity C.xmlGetParameterEntity
func XmlGetParameterEntity(doc XmlDocPtr, name *XmlChar) XmlEntityPtr
//go:linkname XmlEncodeEntities C.xmlEncodeEntities
func XmlEncodeEntities(doc XmlDocPtr, input *XmlChar) *XmlChar
//go:linkname XmlEncodeEntitiesReentrant C.xmlEncodeEntitiesReentrant
func XmlEncodeEntitiesReentrant(doc XmlDocPtr, input *XmlChar) *XmlChar
// llgo:link (*XmlDoc).XmlEncodeSpecialChars C.xmlEncodeSpecialChars
func (recv_ *XmlDoc) XmlEncodeSpecialChars(input *XmlChar) *XmlChar {
	return nil
}
//go:linkname XmlCreateEntitiesTable C.xmlCreateEntitiesTable
func XmlCreateEntitiesTable() XmlEntitiesTablePtr
//go:linkname XmlCopyEntitiesTable C.xmlCopyEntitiesTable
func XmlCopyEntitiesTable(table XmlEntitiesTablePtr) XmlEntitiesTablePtr
//go:linkname XmlFreeEntitiesTable C.xmlFreeEntitiesTable
func XmlFreeEntitiesTable(table XmlEntitiesTablePtr)
//go:linkname XmlDumpEntitiesTable C.xmlDumpEntitiesTable
func XmlDumpEntitiesTable(buf XmlBufferPtr, table XmlEntitiesTablePtr)
//go:linkname XmlDumpEntityDecl C.xmlDumpEntityDecl
func XmlDumpEntityDecl(buf XmlBufferPtr, ent XmlEntityPtr)
//go:linkname XmlCleanupPredefinedEntities C.xmlCleanupPredefinedEntities
func XmlCleanupPredefinedEntities()
