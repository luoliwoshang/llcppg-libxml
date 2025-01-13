package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const XMLSCHEMASANYATTRSKIP c.Int = 1
const XMLSCHEMASANYATTRLAX c.Int = 2
const XMLSCHEMASANYATTRSTRICT c.Int = 3
const XMLSCHEMASANYSKIP c.Int = 1
const XMLSCHEMASANYLAX c.Int = 2
const XMLSCHEMASANYSTRICT c.Int = 3
const XMLSCHEMASATTRUSEPROHIBITED c.Int = 0
const XMLSCHEMASATTRUSEREQUIRED c.Int = 1
const XMLSCHEMASATTRUSEOPTIONAL c.Int = 2
const XMLSCHEMASFACETUNKNOWN c.Int = 0
const XMLSCHEMASFACETPRESERVE c.Int = 1
const XMLSCHEMASFACETREPLACE c.Int = 2
const XMLSCHEMASFACETCOLLAPSE c.Int = 3

type XmlSchemaValType c.Int

const (
	XmlSchemaValTypeXMLSCHEMASUNKNOWN       XmlSchemaValType = 0
	XmlSchemaValTypeXMLSCHEMASSTRING        XmlSchemaValType = 1
	XmlSchemaValTypeXMLSCHEMASNORMSTRING    XmlSchemaValType = 2
	XmlSchemaValTypeXMLSCHEMASDECIMAL       XmlSchemaValType = 3
	XmlSchemaValTypeXMLSCHEMASTIME          XmlSchemaValType = 4
	XmlSchemaValTypeXMLSCHEMASGDAY          XmlSchemaValType = 5
	XmlSchemaValTypeXMLSCHEMASGMONTH        XmlSchemaValType = 6
	XmlSchemaValTypeXMLSCHEMASGMONTHDAY     XmlSchemaValType = 7
	XmlSchemaValTypeXMLSCHEMASGYEAR         XmlSchemaValType = 8
	XmlSchemaValTypeXMLSCHEMASGYEARMONTH    XmlSchemaValType = 9
	XmlSchemaValTypeXMLSCHEMASDATE          XmlSchemaValType = 10
	XmlSchemaValTypeXMLSCHEMASDATETIME      XmlSchemaValType = 11
	XmlSchemaValTypeXMLSCHEMASDURATION      XmlSchemaValType = 12
	XmlSchemaValTypeXMLSCHEMASFLOAT         XmlSchemaValType = 13
	XmlSchemaValTypeXMLSCHEMASDOUBLE        XmlSchemaValType = 14
	XmlSchemaValTypeXMLSCHEMASBOOLEAN       XmlSchemaValType = 15
	XmlSchemaValTypeXMLSCHEMASTOKEN         XmlSchemaValType = 16
	XmlSchemaValTypeXMLSCHEMASLANGUAGE      XmlSchemaValType = 17
	XmlSchemaValTypeXMLSCHEMASNMTOKEN       XmlSchemaValType = 18
	XmlSchemaValTypeXMLSCHEMASNMTOKENS      XmlSchemaValType = 19
	XmlSchemaValTypeXMLSCHEMASNAME          XmlSchemaValType = 20
	XmlSchemaValTypeXMLSCHEMASQNAME         XmlSchemaValType = 21
	XmlSchemaValTypeXMLSCHEMASNCNAME        XmlSchemaValType = 22
	XmlSchemaValTypeXMLSCHEMASID            XmlSchemaValType = 23
	XmlSchemaValTypeXMLSCHEMASIDREF         XmlSchemaValType = 24
	XmlSchemaValTypeXMLSCHEMASIDREFS        XmlSchemaValType = 25
	XmlSchemaValTypeXMLSCHEMASENTITY        XmlSchemaValType = 26
	XmlSchemaValTypeXMLSCHEMASENTITIES      XmlSchemaValType = 27
	XmlSchemaValTypeXMLSCHEMASNOTATION      XmlSchemaValType = 28
	XmlSchemaValTypeXMLSCHEMASANYURI        XmlSchemaValType = 29
	XmlSchemaValTypeXMLSCHEMASINTEGER       XmlSchemaValType = 30
	XmlSchemaValTypeXMLSCHEMASNPINTEGER     XmlSchemaValType = 31
	XmlSchemaValTypeXMLSCHEMASNINTEGER      XmlSchemaValType = 32
	XmlSchemaValTypeXMLSCHEMASNNINTEGER     XmlSchemaValType = 33
	XmlSchemaValTypeXMLSCHEMASPINTEGER      XmlSchemaValType = 34
	XmlSchemaValTypeXMLSCHEMASINT           XmlSchemaValType = 35
	XmlSchemaValTypeXMLSCHEMASUINT          XmlSchemaValType = 36
	XmlSchemaValTypeXMLSCHEMASLONG          XmlSchemaValType = 37
	XmlSchemaValTypeXMLSCHEMASULONG         XmlSchemaValType = 38
	XmlSchemaValTypeXMLSCHEMASSHORT         XmlSchemaValType = 39
	XmlSchemaValTypeXMLSCHEMASUSHORT        XmlSchemaValType = 40
	XmlSchemaValTypeXMLSCHEMASBYTE          XmlSchemaValType = 41
	XmlSchemaValTypeXMLSCHEMASUBYTE         XmlSchemaValType = 42
	XmlSchemaValTypeXMLSCHEMASHEXBINARY     XmlSchemaValType = 43
	XmlSchemaValTypeXMLSCHEMASBASE64BINARY  XmlSchemaValType = 44
	XmlSchemaValTypeXMLSCHEMASANYTYPE       XmlSchemaValType = 45
	XmlSchemaValTypeXMLSCHEMASANYSIMPLETYPE XmlSchemaValType = 46
)

type XmlSchemaTypeType c.Int

const (
	XmlSchemaTypeTypeXMLSCHEMATYPEBASIC           XmlSchemaTypeType = 1
	XmlSchemaTypeTypeXMLSCHEMATYPEANY             XmlSchemaTypeType = 2
	XmlSchemaTypeTypeXMLSCHEMATYPEFACET           XmlSchemaTypeType = 3
	XmlSchemaTypeTypeXMLSCHEMATYPESIMPLE          XmlSchemaTypeType = 4
	XmlSchemaTypeTypeXMLSCHEMATYPECOMPLEX         XmlSchemaTypeType = 5
	XmlSchemaTypeTypeXMLSCHEMATYPESEQUENCE        XmlSchemaTypeType = 6
	XmlSchemaTypeTypeXMLSCHEMATYPECHOICE          XmlSchemaTypeType = 7
	XmlSchemaTypeTypeXMLSCHEMATYPEALL             XmlSchemaTypeType = 8
	XmlSchemaTypeTypeXMLSCHEMATYPESIMPLECONTENT   XmlSchemaTypeType = 9
	XmlSchemaTypeTypeXMLSCHEMATYPECOMPLEXCONTENT  XmlSchemaTypeType = 10
	XmlSchemaTypeTypeXMLSCHEMATYPEUR              XmlSchemaTypeType = 11
	XmlSchemaTypeTypeXMLSCHEMATYPERESTRICTION     XmlSchemaTypeType = 12
	XmlSchemaTypeTypeXMLSCHEMATYPEEXTENSION       XmlSchemaTypeType = 13
	XmlSchemaTypeTypeXMLSCHEMATYPEELEMENT         XmlSchemaTypeType = 14
	XmlSchemaTypeTypeXMLSCHEMATYPEATTRIBUTE       XmlSchemaTypeType = 15
	XmlSchemaTypeTypeXMLSCHEMATYPEATTRIBUTEGROUP  XmlSchemaTypeType = 16
	XmlSchemaTypeTypeXMLSCHEMATYPEGROUP           XmlSchemaTypeType = 17
	XmlSchemaTypeTypeXMLSCHEMATYPENOTATION        XmlSchemaTypeType = 18
	XmlSchemaTypeTypeXMLSCHEMATYPELIST            XmlSchemaTypeType = 19
	XmlSchemaTypeTypeXMLSCHEMATYPEUNION           XmlSchemaTypeType = 20
	XmlSchemaTypeTypeXMLSCHEMATYPEANYATTRIBUTE    XmlSchemaTypeType = 21
	XmlSchemaTypeTypeXMLSCHEMATYPEIDCUNIQUE       XmlSchemaTypeType = 22
	XmlSchemaTypeTypeXMLSCHEMATYPEIDCKEY          XmlSchemaTypeType = 23
	XmlSchemaTypeTypeXMLSCHEMATYPEIDCKEYREF       XmlSchemaTypeType = 24
	XmlSchemaTypeTypeXMLSCHEMATYPEPARTICLE        XmlSchemaTypeType = 25
	XmlSchemaTypeTypeXMLSCHEMATYPEATTRIBUTEUSE    XmlSchemaTypeType = 26
	XmlSchemaTypeTypeXMLSCHEMAFACETMININCLUSIVE   XmlSchemaTypeType = 1000
	XmlSchemaTypeTypeXMLSCHEMAFACETMINEXCLUSIVE   XmlSchemaTypeType = 1001
	XmlSchemaTypeTypeXMLSCHEMAFACETMAXINCLUSIVE   XmlSchemaTypeType = 1002
	XmlSchemaTypeTypeXMLSCHEMAFACETMAXEXCLUSIVE   XmlSchemaTypeType = 1003
	XmlSchemaTypeTypeXMLSCHEMAFACETTOTALDIGITS    XmlSchemaTypeType = 1004
	XmlSchemaTypeTypeXMLSCHEMAFACETFRACTIONDIGITS XmlSchemaTypeType = 1005
	XmlSchemaTypeTypeXMLSCHEMAFACETPATTERN        XmlSchemaTypeType = 1006
	XmlSchemaTypeTypeXMLSCHEMAFACETENUMERATION    XmlSchemaTypeType = 1007
	XmlSchemaTypeTypeXMLSCHEMAFACETWHITESPACE     XmlSchemaTypeType = 1008
	XmlSchemaTypeTypeXMLSCHEMAFACETLENGTH         XmlSchemaTypeType = 1009
	XmlSchemaTypeTypeXMLSCHEMAFACETMAXLENGTH      XmlSchemaTypeType = 1010
	XmlSchemaTypeTypeXMLSCHEMAFACETMINLENGTH      XmlSchemaTypeType = 1011
	XmlSchemaTypeTypeXMLSCHEMAEXTRAQNAMEREF       XmlSchemaTypeType = 2000
	XmlSchemaTypeTypeXMLSCHEMAEXTRAATTRUSEPROHIB  XmlSchemaTypeType = 2001
)

type XmlSchemaContentType c.Int

const (
	XmlSchemaContentTypeXMLSCHEMACONTENTUNKNOWN         XmlSchemaContentType = 0
	XmlSchemaContentTypeXMLSCHEMACONTENTEMPTY           XmlSchemaContentType = 1
	XmlSchemaContentTypeXMLSCHEMACONTENTELEMENTS        XmlSchemaContentType = 2
	XmlSchemaContentTypeXMLSCHEMACONTENTMIXED           XmlSchemaContentType = 3
	XmlSchemaContentTypeXMLSCHEMACONTENTSIMPLE          XmlSchemaContentType = 4
	XmlSchemaContentTypeXMLSCHEMACONTENTMIXEDORELEMENTS XmlSchemaContentType = 5
	XmlSchemaContentTypeXMLSCHEMACONTENTBASIC           XmlSchemaContentType = 6
	XmlSchemaContentTypeXMLSCHEMACONTENTANY             XmlSchemaContentType = 7
)

type X_XmlSchemaVal struct {
	Unused [8]uint8
}
type XmlSchemaVal X_XmlSchemaVal
type XmlSchemaValPtr *XmlSchemaVal

type X_XmlSchemaType struct {
	Type              XmlSchemaTypeType
	Next              *X_XmlSchemaType
	Name              *XmlChar
	Id                *XmlChar
	Ref               *XmlChar
	RefNs             *XmlChar
	Annot             XmlSchemaAnnotPtr
	Subtypes          XmlSchemaTypePtr
	Attributes        XmlSchemaAttributePtr
	Node              XmlNodePtr
	MinOccurs         c.Int
	MaxOccurs         c.Int
	Flags             c.Int
	ContentType       XmlSchemaContentType
	Base              *XmlChar
	BaseNs            *XmlChar
	BaseType          XmlSchemaTypePtr
	Facets            XmlSchemaFacetPtr
	Redef             *X_XmlSchemaType
	Recurse           c.Int
	AttributeUses     *XmlSchemaAttributeLinkPtr
	AttributeWildcard XmlSchemaWildcardPtr
	BuiltInType       c.Int
	MemberTypes       XmlSchemaTypeLinkPtr
	FacetSet          XmlSchemaFacetLinkPtr
	RefPrefix         *XmlChar
	ContentTypeDef    XmlSchemaTypePtr
	ContModel         XmlRegexpPtr
	TargetNamespace   *XmlChar
	AttrUses          unsafe.Pointer
}
type XmlSchemaType X_XmlSchemaType
type XmlSchemaTypePtr *XmlSchemaType

type X_XmlSchemaFacet struct {
	Type       XmlSchemaTypeType
	Next       *X_XmlSchemaFacet
	Value      *XmlChar
	Id         *XmlChar
	Annot      XmlSchemaAnnotPtr
	Node       XmlNodePtr
	Fixed      c.Int
	Whitespace c.Int
	Val        XmlSchemaValPtr
	Regexp     XmlRegexpPtr
}
type XmlSchemaFacet X_XmlSchemaFacet
type XmlSchemaFacetPtr *XmlSchemaFacet

type X_XmlSchemaAnnot struct {
	Next    *X_XmlSchemaAnnot
	Content XmlNodePtr
}
type XmlSchemaAnnot X_XmlSchemaAnnot
type XmlSchemaAnnotPtr *XmlSchemaAnnot

type X_XmlSchemaAttribute struct {
	Type            XmlSchemaTypeType
	Next            *X_XmlSchemaAttribute
	Name            *XmlChar
	Id              *XmlChar
	Ref             *XmlChar
	RefNs           *XmlChar
	TypeName        *XmlChar
	TypeNs          *XmlChar
	Annot           XmlSchemaAnnotPtr
	Base            XmlSchemaTypePtr
	Occurs          c.Int
	DefValue        *XmlChar
	Subtypes        XmlSchemaTypePtr
	Node            XmlNodePtr
	TargetNamespace *XmlChar
	Flags           c.Int
	RefPrefix       *XmlChar
	DefVal          XmlSchemaValPtr
	RefDecl         XmlSchemaAttributePtr
}
type XmlSchemaAttribute X_XmlSchemaAttribute
type XmlSchemaAttributePtr *XmlSchemaAttribute

type X_XmlSchemaAttributeLink struct {
	Next *X_XmlSchemaAttributeLink
	Attr *X_XmlSchemaAttribute
}
type XmlSchemaAttributeLink X_XmlSchemaAttributeLink
type XmlSchemaAttributeLinkPtr *XmlSchemaAttributeLink

type X_XmlSchemaWildcardNs struct {
	Next  *X_XmlSchemaWildcardNs
	Value *XmlChar
}
type XmlSchemaWildcardNs X_XmlSchemaWildcardNs
type XmlSchemaWildcardNsPtr *XmlSchemaWildcardNs

type X_XmlSchemaWildcard struct {
	Type            XmlSchemaTypeType
	Id              *XmlChar
	Annot           XmlSchemaAnnotPtr
	Node            XmlNodePtr
	MinOccurs       c.Int
	MaxOccurs       c.Int
	ProcessContents c.Int
	Any             c.Int
	NsSet           XmlSchemaWildcardNsPtr
	NegNsSet        XmlSchemaWildcardNsPtr
	Flags           c.Int
}
type XmlSchemaWildcard X_XmlSchemaWildcard
type XmlSchemaWildcardPtr *XmlSchemaWildcard

type X_XmlSchemaAttributeGroup struct {
	Type              XmlSchemaTypeType
	Next              *X_XmlSchemaAttribute
	Name              *XmlChar
	Id                *XmlChar
	Ref               *XmlChar
	RefNs             *XmlChar
	Annot             XmlSchemaAnnotPtr
	Attributes        XmlSchemaAttributePtr
	Node              XmlNodePtr
	Flags             c.Int
	AttributeWildcard XmlSchemaWildcardPtr
	RefPrefix         *XmlChar
	RefItem           XmlSchemaAttributeGroupPtr
	TargetNamespace   *XmlChar
	AttrUses          unsafe.Pointer
}
type XmlSchemaAttributeGroup X_XmlSchemaAttributeGroup
type XmlSchemaAttributeGroupPtr *XmlSchemaAttributeGroup

type X_XmlSchemaTypeLink struct {
	Next *X_XmlSchemaTypeLink
	Type XmlSchemaTypePtr
}
type XmlSchemaTypeLink X_XmlSchemaTypeLink
type XmlSchemaTypeLinkPtr *XmlSchemaTypeLink

type X_XmlSchemaFacetLink struct {
	Next  *X_XmlSchemaFacetLink
	Facet XmlSchemaFacetPtr
}
type XmlSchemaFacetLink X_XmlSchemaFacetLink
type XmlSchemaFacetLinkPtr *XmlSchemaFacetLink

type X_XmlSchemaElement struct {
	Type            XmlSchemaTypeType
	Next            *X_XmlSchemaType
	Name            *XmlChar
	Id              *XmlChar
	Ref             *XmlChar
	RefNs           *XmlChar
	Annot           XmlSchemaAnnotPtr
	Subtypes        XmlSchemaTypePtr
	Attributes      XmlSchemaAttributePtr
	Node            XmlNodePtr
	MinOccurs       c.Int
	MaxOccurs       c.Int
	Flags           c.Int
	TargetNamespace *XmlChar
	NamedType       *XmlChar
	NamedTypeNs     *XmlChar
	SubstGroup      *XmlChar
	SubstGroupNs    *XmlChar
	Scope           *XmlChar
	Value           *XmlChar
	RefDecl         *X_XmlSchemaElement
	ContModel       XmlRegexpPtr
	ContentType     XmlSchemaContentType
	RefPrefix       *XmlChar
	DefVal          XmlSchemaValPtr
	Idcs            unsafe.Pointer
}
type XmlSchemaElement X_XmlSchemaElement
type XmlSchemaElementPtr *XmlSchemaElement

type X_XmlSchemaNotation struct {
	Type            XmlSchemaTypeType
	Name            *XmlChar
	Annot           XmlSchemaAnnotPtr
	Identifier      *XmlChar
	TargetNamespace *XmlChar
}
type XmlSchemaNotation X_XmlSchemaNotation
type XmlSchemaNotationPtr *XmlSchemaNotation
/**
 * _xmlSchema:
 *
 * A Schemas definition
 */

type X_XmlSchema struct {
	Name            *XmlChar
	TargetNamespace *XmlChar
	Version         *XmlChar
	Id              *XmlChar
	Doc             XmlDocPtr
	Annot           XmlSchemaAnnotPtr
	Flags           c.Int
	TypeDecl        XmlHashTablePtr
	AttrDecl        XmlHashTablePtr
	AttrgrpDecl     XmlHashTablePtr
	ElemDecl        XmlHashTablePtr
	NotaDecl        XmlHashTablePtr
	SchemasImports  XmlHashTablePtr
	X_Private       unsafe.Pointer
	GroupDecl       XmlHashTablePtr
	Dict            XmlDictPtr
	Includes        unsafe.Pointer
	Preserve        c.Int
	Counter         c.Int
	IdcDef          XmlHashTablePtr
	Volatiles       unsafe.Pointer
}
//go:linkname XmlSchemaFreeType C.xmlSchemaFreeType
func XmlSchemaFreeType(type_ XmlSchemaTypePtr)
//go:linkname XmlSchemaFreeWildcard C.xmlSchemaFreeWildcard
func XmlSchemaFreeWildcard(wildcard XmlSchemaWildcardPtr)
