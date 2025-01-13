package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const BASEBUFFERSIZE c.Int = 4096
const XMLDOCBDOCUMENTNODE c.Int = 21

type X_XmlParserInputBuffer struct {
	Context       unsafe.Pointer
	Readcallback  unsafe.Pointer
	Closecallback unsafe.Pointer
	Encoder       XmlCharEncodingHandlerPtr
	Buffer        XmlBufPtr
	Raw           XmlBufPtr
	Compressed    c.Int
	Error         c.Int
	Rawconsumed   c.Ulong
}
type XmlParserInputBuffer X_XmlParserInputBuffer
type XmlParserInputBufferPtr *XmlParserInputBuffer

type X_XmlOutputBuffer struct {
	Context       unsafe.Pointer
	Writecallback unsafe.Pointer
	Closecallback unsafe.Pointer
	Encoder       XmlCharEncodingHandlerPtr
	Buffer        XmlBufPtr
	Conv          XmlBufPtr
	Written       c.Int
	Error         c.Int
}
type XmlOutputBuffer X_XmlOutputBuffer
type XmlOutputBufferPtr *XmlOutputBuffer

type X_XmlParserInput struct {
	Buf            XmlParserInputBufferPtr
	Filename       *int8
	Directory      *int8
	Base           *XmlChar
	Cur            *XmlChar
	End            *XmlChar
	Length         c.Int
	Line           c.Int
	Col            c.Int
	Consumed       c.Ulong
	Free           unsafe.Pointer
	Encoding       *XmlChar
	Version        *XmlChar
	Flags          c.Int
	Id             c.Int
	ParentConsumed c.Ulong
	Entity         XmlEntityPtr
}
type XmlParserInput X_XmlParserInput
type XmlParserInputPtr *XmlParserInput

type X_XmlParserCtxt struct {
	Sax               *X_XmlSAXHandler
	UserData          unsafe.Pointer
	MyDoc             XmlDocPtr
	WellFormed        c.Int
	ReplaceEntities   c.Int
	Version           *XmlChar
	Encoding          *XmlChar
	Standalone        c.Int
	Html              c.Int
	Input             XmlParserInputPtr
	InputNr           c.Int
	InputMax          c.Int
	InputTab          *XmlParserInputPtr
	Node              XmlNodePtr
	NodeNr            c.Int
	NodeMax           c.Int
	NodeTab           *XmlNodePtr
	RecordInfo        c.Int
	NodeSeq           XmlParserNodeInfoSeq
	ErrNo             c.Int
	HasExternalSubset c.Int
	HasPErefs         c.Int
	External          c.Int
	Valid             c.Int
	Validate          c.Int
	Vctxt             XmlValidCtxt
	Instate           XmlParserInputState
	Token             c.Int
	Directory         *int8
	Name              *XmlChar
	NameNr            c.Int
	NameMax           c.Int
	NameTab           **XmlChar
	NbChars           c.Long
	CheckIndex        c.Long
	KeepBlanks        c.Int
	DisableSAX        c.Int
	InSubset          c.Int
	IntSubName        *XmlChar
	ExtSubURI         *XmlChar
	ExtSubSystem      *XmlChar
	Space             *c.Int
	SpaceNr           c.Int
	SpaceMax          c.Int
	SpaceTab          *c.Int
	Depth             c.Int
	Entity            XmlParserInputPtr
	Charset           c.Int
	Nodelen           c.Int
	Nodemem           c.Int
	Pedantic          c.Int
	X_Private         unsafe.Pointer
	Loadsubset        c.Int
	Linenumbers       c.Int
	Catalogs          unsafe.Pointer
	Recovery          c.Int
	Progressive       c.Int
	Dict              XmlDictPtr
	Atts              **XmlChar
	Maxatts           c.Int
	Docdict           c.Int
	StrXml            *XmlChar
	StrXmlns          *XmlChar
	StrXmlNs          *XmlChar
	Sax2              c.Int
	NsNr              c.Int
	NsMax             c.Int
	NsTab             **XmlChar
	Attallocs         *c.Uint
	PushTab           *XmlStartTag
	AttsDefault       XmlHashTablePtr
	AttsSpecial       XmlHashTablePtr
	NsWellFormed      c.Int
	Options           c.Int
	DictNames         c.Int
	FreeElemsNr       c.Int
	FreeElems         XmlNodePtr
	FreeAttrsNr       c.Int
	FreeAttrs         XmlAttrPtr
	LastError         XmlError
	ParseMode         XmlParserMode
	Nbentities        c.Ulong
	Sizeentities      c.Ulong
	NodeInfo          *XmlParserNodeInfo
	NodeInfoNr        c.Int
	NodeInfoMax       c.Int
	NodeInfoTab       *XmlParserNodeInfo
	InputId           c.Int
	Sizeentcopy       c.Ulong
	EndCheckState     c.Int
	NbErrors          uint16
	NbWarnings        uint16
	MaxAmpl           c.Uint
	Nsdb              *XmlParserNsData
	AttrHashMax       c.Uint
	AttrHash          *XmlAttrHashBucket
	ErrorHandler      unsafe.Pointer
	ErrorCtxt         unsafe.Pointer
}
type XmlParserCtxt X_XmlParserCtxt
type XmlParserCtxtPtr *XmlParserCtxt

type X_XmlSAXLocator struct {
	GetPublicId     unsafe.Pointer
	GetSystemId     unsafe.Pointer
	GetLineNumber   unsafe.Pointer
	GetColumnNumber unsafe.Pointer
}
type XmlSAXLocator X_XmlSAXLocator
type XmlSAXLocatorPtr *XmlSAXLocator

type X_XmlSAXHandler struct {
	InternalSubset        unsafe.Pointer
	IsStandalone          unsafe.Pointer
	HasInternalSubset     unsafe.Pointer
	HasExternalSubset     unsafe.Pointer
	ResolveEntity         unsafe.Pointer
	GetEntity             unsafe.Pointer
	EntityDecl            unsafe.Pointer
	NotationDecl          unsafe.Pointer
	AttributeDecl         unsafe.Pointer
	ElementDecl           unsafe.Pointer
	UnparsedEntityDecl    unsafe.Pointer
	SetDocumentLocator    unsafe.Pointer
	StartDocument         unsafe.Pointer
	EndDocument           unsafe.Pointer
	StartElement          unsafe.Pointer
	EndElement            unsafe.Pointer
	Reference             unsafe.Pointer
	Characters            unsafe.Pointer
	IgnorableWhitespace   unsafe.Pointer
	ProcessingInstruction unsafe.Pointer
	Comment               unsafe.Pointer
	Warning               unsafe.Pointer
	Error                 unsafe.Pointer
	FatalError            unsafe.Pointer
	GetParameterEntity    unsafe.Pointer
	CdataBlock            unsafe.Pointer
	ExternalSubset        unsafe.Pointer
	Initialized           c.Uint
	X_Private             unsafe.Pointer
	StartElementNs        unsafe.Pointer
	EndElementNs          unsafe.Pointer
	Serror                unsafe.Pointer
}
type XmlSAXHandler X_XmlSAXHandler
type XmlSAXHandlerPtr *XmlSAXHandler

type X_XmlEntity struct {
	X_Private    unsafe.Pointer
	Type         XmlElementType
	Name         *XmlChar
	Children     *X_XmlNode
	Last         *X_XmlNode
	Parent       *X_XmlDtd
	Next         *X_XmlNode
	Prev         *X_XmlNode
	Doc          *X_XmlDoc
	Orig         *XmlChar
	Content      *XmlChar
	Length       c.Int
	Etype        XmlEntityType
	ExternalID   *XmlChar
	SystemID     *XmlChar
	Nexte        *X_XmlEntity
	URI          *XmlChar
	Owner        c.Int
	Flags        c.Int
	ExpandedSize c.Ulong
}
type XmlEntity X_XmlEntity
type XmlEntityPtr *XmlEntity
type XmlBufferAllocationScheme c.Int

const (
	XmlBufferAllocationSchemeXMLBUFFERALLOCDOUBLEIT  XmlBufferAllocationScheme = 0
	XmlBufferAllocationSchemeXMLBUFFERALLOCEXACT     XmlBufferAllocationScheme = 1
	XmlBufferAllocationSchemeXMLBUFFERALLOCIMMUTABLE XmlBufferAllocationScheme = 2
	XmlBufferAllocationSchemeXMLBUFFERALLOCIO        XmlBufferAllocationScheme = 3
	XmlBufferAllocationSchemeXMLBUFFERALLOCHYBRID    XmlBufferAllocationScheme = 4
	XmlBufferAllocationSchemeXMLBUFFERALLOCBOUNDED   XmlBufferAllocationScheme = 5
)

type X_XmlBuffer struct {
	Content   *XmlChar
	Use       c.Uint
	Size      c.Uint
	Alloc     XmlBufferAllocationScheme
	ContentIO *XmlChar
}
type XmlBuffer X_XmlBuffer
type XmlBufferPtr *XmlBuffer

type X_XmlBuf struct {
	Unused [8]uint8
}
type XmlBuf X_XmlBuf
type XmlBufPtr *XmlBuf
// llgo:link (*XmlBuf).XmlBufContent C.xmlBufContent
func (recv_ *XmlBuf) XmlBufContent() *XmlChar {
	return nil
}
//go:linkname XmlBufEnd C.xmlBufEnd
func XmlBufEnd(buf XmlBufPtr) *XmlChar
//go:linkname XmlBufUse C.xmlBufUse
func XmlBufUse(buf XmlBufPtr) uintptr
//go:linkname XmlBufShrink C.xmlBufShrink
func XmlBufShrink(buf XmlBufPtr, len uintptr) uintptr

type XmlElementType c.Int

const (
	XmlElementTypeXMLELEMENTNODE      XmlElementType = 1
	XmlElementTypeXMLATTRIBUTENODE    XmlElementType = 2
	XmlElementTypeXMLTEXTNODE         XmlElementType = 3
	XmlElementTypeXMLCDATASECTIONNODE XmlElementType = 4
	XmlElementTypeXMLENTITYREFNODE    XmlElementType = 5
	XmlElementTypeXMLENTITYNODE       XmlElementType = 6
	XmlElementTypeXMLPINODE           XmlElementType = 7
	XmlElementTypeXMLCOMMENTNODE      XmlElementType = 8
	XmlElementTypeXMLDOCUMENTNODE     XmlElementType = 9
	XmlElementTypeXMLDOCUMENTTYPENODE XmlElementType = 10
	XmlElementTypeXMLDOCUMENTFRAGNODE XmlElementType = 11
	XmlElementTypeXMLNOTATIONNODE     XmlElementType = 12
	XmlElementTypeXMLHTMLDOCUMENTNODE XmlElementType = 13
	XmlElementTypeXMLDTDNODE          XmlElementType = 14
	XmlElementTypeXMLELEMENTDECL      XmlElementType = 15
	XmlElementTypeXMLATTRIBUTEDECL    XmlElementType = 16
	XmlElementTypeXMLENTITYDECL       XmlElementType = 17
	XmlElementTypeXMLNAMESPACEDECL    XmlElementType = 18
	XmlElementTypeXMLXINCLUDESTART    XmlElementType = 19
	XmlElementTypeXMLXINCLUDEEND      XmlElementType = 20
)

type X_XmlNotation struct {
	Name     *XmlChar
	PublicID *XmlChar
	SystemID *XmlChar
}
type XmlNotation X_XmlNotation
type XmlNotationPtr *XmlNotation
type XmlAttributeType c.Int

const (
	XmlAttributeTypeXMLATTRIBUTECDATA       XmlAttributeType = 1
	XmlAttributeTypeXMLATTRIBUTEID          XmlAttributeType = 2
	XmlAttributeTypeXMLATTRIBUTEIDREF       XmlAttributeType = 3
	XmlAttributeTypeXMLATTRIBUTEIDREFS      XmlAttributeType = 4
	XmlAttributeTypeXMLATTRIBUTEENTITY      XmlAttributeType = 5
	XmlAttributeTypeXMLATTRIBUTEENTITIES    XmlAttributeType = 6
	XmlAttributeTypeXMLATTRIBUTENMTOKEN     XmlAttributeType = 7
	XmlAttributeTypeXMLATTRIBUTENMTOKENS    XmlAttributeType = 8
	XmlAttributeTypeXMLATTRIBUTEENUMERATION XmlAttributeType = 9
	XmlAttributeTypeXMLATTRIBUTENOTATION    XmlAttributeType = 10
)

type XmlAttributeDefault c.Int

const (
	XmlAttributeDefaultXMLATTRIBUTENONE     XmlAttributeDefault = 1
	XmlAttributeDefaultXMLATTRIBUTEREQUIRED XmlAttributeDefault = 2
	XmlAttributeDefaultXMLATTRIBUTEIMPLIED  XmlAttributeDefault = 3
	XmlAttributeDefaultXMLATTRIBUTEFIXED    XmlAttributeDefault = 4
)

type X_XmlEnumeration struct {
	Next *X_XmlEnumeration
	Name *XmlChar
}
type XmlEnumeration X_XmlEnumeration
type XmlEnumerationPtr *XmlEnumeration

type X_XmlAttribute struct {
	X_Private    unsafe.Pointer
	Type         XmlElementType
	Name         *XmlChar
	Children     *X_XmlNode
	Last         *X_XmlNode
	Parent       *X_XmlDtd
	Next         *X_XmlNode
	Prev         *X_XmlNode
	Doc          *X_XmlDoc
	Nexth        *X_XmlAttribute
	Atype        XmlAttributeType
	Def          XmlAttributeDefault
	DefaultValue *XmlChar
	Tree         XmlEnumerationPtr
	Prefix       *XmlChar
	Elem         *XmlChar
}
type XmlAttribute X_XmlAttribute
type XmlAttributePtr *XmlAttribute

type X_XmlNode struct {
	X_Private  unsafe.Pointer
	Type       XmlElementType
	Name       *XmlChar
	Children   *X_XmlNode
	Last       *X_XmlNode
	Parent     *X_XmlNode
	Next       *X_XmlNode
	Prev       *X_XmlNode
	Doc        *X_XmlDoc
	Ns         *XmlNs
	Content    *XmlChar
	Properties *X_XmlAttr
	NsDef      *XmlNs
	Psvi       unsafe.Pointer
	Line       uint16
	Extra      uint16
}

type X_XmlDtd struct {
	X_Private  unsafe.Pointer
	Type       XmlElementType
	Name       *XmlChar
	Children   *X_XmlNode
	Last       *X_XmlNode
	Parent     *X_XmlDoc
	Next       *X_XmlNode
	Prev       *X_XmlNode
	Doc        *X_XmlDoc
	Notations  unsafe.Pointer
	Elements   unsafe.Pointer
	Attributes unsafe.Pointer
	Entities   unsafe.Pointer
	ExternalID *XmlChar
	SystemID   *XmlChar
	Pentities  unsafe.Pointer
}

type X_XmlDoc struct {
	X_Private   unsafe.Pointer
	Type        XmlElementType
	Name        *int8
	Children    *X_XmlNode
	Last        *X_XmlNode
	Parent      *X_XmlNode
	Next        *X_XmlNode
	Prev        *X_XmlNode
	Doc         *X_XmlDoc
	Compression c.Int
	Standalone  c.Int
	IntSubset   *X_XmlDtd
	ExtSubset   *X_XmlDtd
	OldNs       *X_XmlNs
	Version     *XmlChar
	Encoding    *XmlChar
	Ids         unsafe.Pointer
	Refs        unsafe.Pointer
	URL         *XmlChar
	Charset     c.Int
	Dict        *X_XmlDict
	Psvi        unsafe.Pointer
	ParseFlags  c.Int
	Properties  c.Int
}
type XmlElementContentType c.Int

const (
	XmlElementContentTypeXMLELEMENTCONTENTPCDATA  XmlElementContentType = 1
	XmlElementContentTypeXMLELEMENTCONTENTELEMENT XmlElementContentType = 2
	XmlElementContentTypeXMLELEMENTCONTENTSEQ     XmlElementContentType = 3
	XmlElementContentTypeXMLELEMENTCONTENTOR      XmlElementContentType = 4
)

type XmlElementContentOccur c.Int

const (
	XmlElementContentOccurXMLELEMENTCONTENTONCE XmlElementContentOccur = 1
	XmlElementContentOccurXMLELEMENTCONTENTOPT  XmlElementContentOccur = 2
	XmlElementContentOccurXMLELEMENTCONTENTMULT XmlElementContentOccur = 3
	XmlElementContentOccurXMLELEMENTCONTENTPLUS XmlElementContentOccur = 4
)

type X_XmlElementContent struct {
	Type   XmlElementContentType
	Ocur   XmlElementContentOccur
	Name   *XmlChar
	C1     *X_XmlElementContent
	C2     *X_XmlElementContent
	Parent *X_XmlElementContent
	Prefix *XmlChar
}
type XmlElementContent X_XmlElementContent
type XmlElementContentPtr *XmlElementContent
type XmlElementTypeVal c.Int

const (
	XmlElementTypeValXMLELEMENTTYPEUNDEFINED XmlElementTypeVal = 0
	XmlElementTypeValXMLELEMENTTYPEEMPTY     XmlElementTypeVal = 1
	XmlElementTypeValXMLELEMENTTYPEANY       XmlElementTypeVal = 2
	XmlElementTypeValXMLELEMENTTYPEMIXED     XmlElementTypeVal = 3
	XmlElementTypeValXMLELEMENTTYPEELEMENT   XmlElementTypeVal = 4
)

type X_XmlElement struct {
	X_Private  unsafe.Pointer
	Type       XmlElementType
	Name       *XmlChar
	Children   *X_XmlNode
	Last       *X_XmlNode
	Parent     *X_XmlDtd
	Next       *X_XmlNode
	Prev       *X_XmlNode
	Doc        *X_XmlDoc
	Etype      XmlElementTypeVal
	Content    XmlElementContentPtr
	Attributes XmlAttributePtr
	Prefix     *XmlChar
	ContModel  XmlRegexpPtr
}
type XmlElement X_XmlElement
type XmlElementPtr *XmlElement
type XmlNsType XmlElementType

type X_XmlNs struct {
	Next      *X_XmlNs
	Type      XmlNsType
	Href      *XmlChar
	Prefix    *XmlChar
	X_Private unsafe.Pointer
	Context   *X_XmlDoc
}
type XmlNs X_XmlNs
type XmlNsPtr *XmlNs
type XmlDtd X_XmlDtd
type XmlDtdPtr *XmlDtd

type X_XmlAttr struct {
	X_Private unsafe.Pointer
	Type      XmlElementType
	Name      *XmlChar
	Children  *X_XmlNode
	Last      *X_XmlNode
	Parent    *X_XmlNode
	Next      *X_XmlAttr
	Prev      *X_XmlAttr
	Doc       *X_XmlDoc
	Ns        *XmlNs
	Atype     XmlAttributeType
	Psvi      unsafe.Pointer
	Id        *X_XmlID
}
type XmlAttr X_XmlAttr
type XmlAttrPtr *XmlAttr

type X_XmlID struct {
	Next   *X_XmlID
	Value  *XmlChar
	Attr   XmlAttrPtr
	Name   *XmlChar
	Lineno c.Int
	Doc    *X_XmlDoc
}
type XmlID X_XmlID
type XmlIDPtr *XmlID

type X_XmlRef struct {
	Next   *X_XmlRef
	Value  *XmlChar
	Attr   XmlAttrPtr
	Name   *XmlChar
	Lineno c.Int
}
type XmlRef X_XmlRef
type XmlRefPtr *XmlRef
type XmlNode X_XmlNode
type XmlNodePtr *XmlNode
type XmlDocProperties c.Int

const (
	XmlDocPropertiesXMLDOCWELLFORMED XmlDocProperties = 1
	XmlDocPropertiesXMLDOCNSVALID    XmlDocProperties = 2
	XmlDocPropertiesXMLDOCOLD10      XmlDocProperties = 4
	XmlDocPropertiesXMLDOCDTDVALID   XmlDocProperties = 8
	XmlDocPropertiesXMLDOCXINCLUDE   XmlDocProperties = 16
	XmlDocPropertiesXMLDOCUSERBUILT  XmlDocProperties = 32
	XmlDocPropertiesXMLDOCINTERNAL   XmlDocProperties = 64
	XmlDocPropertiesXMLDOCHTML       XmlDocProperties = 128
)

type XmlDoc X_XmlDoc
type XmlDocPtr *XmlDoc

type X_XmlDict struct {
	Unused [8]uint8
}

type X_XmlDOMWrapCtxt struct {
	X_Private        unsafe.Pointer
	Type             c.Int
	NamespaceMap     unsafe.Pointer
	GetNsForNodeFunc unsafe.Pointer
}
type XmlDOMWrapCtxt X_XmlDOMWrapCtxt
type XmlDOMWrapCtxtPtr *XmlDOMWrapCtxt
// llgo:type C
type XmlDOMWrapAcquireNsFunction func(XmlDOMWrapCtxtPtr, XmlNodePtr, *XmlChar, *XmlChar) XmlNsPtr
// llgo:type C
type XmlRegisterNodeFunc func(XmlNodePtr)
// llgo:type C
type XmlDeregisterNodeFunc func(XmlNodePtr)
//go:linkname X__XmlBufferAllocScheme C.__xmlBufferAllocScheme
func X__XmlBufferAllocScheme() *XmlBufferAllocationScheme
//go:linkname X__XmlDefaultBufferSize C.__xmlDefaultBufferSize
func X__XmlDefaultBufferSize() *c.Int
//go:linkname X__XmlRegisterNodeDefaultValue C.__xmlRegisterNodeDefaultValue
func X__XmlRegisterNodeDefaultValue() XmlRegisterNodeFunc
//go:linkname X__XmlDeregisterNodeDefaultValue C.__xmlDeregisterNodeDefaultValue
func X__XmlDeregisterNodeDefaultValue() XmlDeregisterNodeFunc
/** DOC_ENABLE */
// llgo:link (*XmlChar).XmlValidateNCName C.xmlValidateNCName
func (recv_ *XmlChar) XmlValidateNCName(space c.Int) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlValidateQName C.xmlValidateQName
func (recv_ *XmlChar) XmlValidateQName(space c.Int) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlValidateName C.xmlValidateName
func (recv_ *XmlChar) XmlValidateName(space c.Int) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlValidateNMToken C.xmlValidateNMToken
func (recv_ *XmlChar) XmlValidateNMToken(space c.Int) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlBuildQName C.xmlBuildQName
func (recv_ *XmlChar) XmlBuildQName(prefix *XmlChar, memory *XmlChar, len c.Int) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlSplitQName2 C.xmlSplitQName2
func (recv_ *XmlChar) XmlSplitQName2(prefix **XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlSplitQName3 C.xmlSplitQName3
func (recv_ *XmlChar) XmlSplitQName3(len *c.Int) *XmlChar {
	return nil
}
// llgo:link XmlBufferAllocationScheme.XmlSetBufferAllocationScheme C.xmlSetBufferAllocationScheme
func (recv_ XmlBufferAllocationScheme) XmlSetBufferAllocationScheme() {
}
//go:linkname XmlGetBufferAllocationScheme C.xmlGetBufferAllocationScheme
func XmlGetBufferAllocationScheme() XmlBufferAllocationScheme
//go:linkname XmlBufferCreate C.xmlBufferCreate
func XmlBufferCreate() XmlBufferPtr
//go:linkname XmlBufferCreateSize C.xmlBufferCreateSize
func XmlBufferCreateSize(size uintptr) XmlBufferPtr
//go:linkname XmlBufferCreateStatic C.xmlBufferCreateStatic
func XmlBufferCreateStatic(mem unsafe.Pointer, size uintptr) XmlBufferPtr
//go:linkname XmlBufferResize C.xmlBufferResize
func XmlBufferResize(buf XmlBufferPtr, size c.Uint) c.Int
//go:linkname XmlBufferFree C.xmlBufferFree
func XmlBufferFree(buf XmlBufferPtr)
//go:linkname XmlBufferDump C.xmlBufferDump
func XmlBufferDump(file *c.FILE, buf XmlBufferPtr) c.Int
//go:linkname XmlBufferAdd C.xmlBufferAdd
func XmlBufferAdd(buf XmlBufferPtr, str *XmlChar, len c.Int) c.Int
//go:linkname XmlBufferAddHead C.xmlBufferAddHead
func XmlBufferAddHead(buf XmlBufferPtr, str *XmlChar, len c.Int) c.Int
//go:linkname XmlBufferCat C.xmlBufferCat
func XmlBufferCat(buf XmlBufferPtr, str *XmlChar) c.Int
//go:linkname XmlBufferCCat C.xmlBufferCCat
func XmlBufferCCat(buf XmlBufferPtr, str *int8) c.Int
//go:linkname XmlBufferShrink C.xmlBufferShrink
func XmlBufferShrink(buf XmlBufferPtr, len c.Uint) c.Int
//go:linkname XmlBufferGrow C.xmlBufferGrow
func XmlBufferGrow(buf XmlBufferPtr, len c.Uint) c.Int
//go:linkname XmlBufferEmpty C.xmlBufferEmpty
func XmlBufferEmpty(buf XmlBufferPtr)
// llgo:link (*XmlBuffer).XmlBufferContent C.xmlBufferContent
func (recv_ *XmlBuffer) XmlBufferContent() *XmlChar {
	return nil
}
//go:linkname XmlBufferDetach C.xmlBufferDetach
func XmlBufferDetach(buf XmlBufferPtr) *XmlChar
//go:linkname XmlBufferSetAllocationScheme C.xmlBufferSetAllocationScheme
func XmlBufferSetAllocationScheme(buf XmlBufferPtr, scheme XmlBufferAllocationScheme)
// llgo:link (*XmlBuffer).XmlBufferLength C.xmlBufferLength
func (recv_ *XmlBuffer) XmlBufferLength() c.Int {
	return 0
}
//go:linkname XmlCreateIntSubset C.xmlCreateIntSubset
func XmlCreateIntSubset(doc XmlDocPtr, name *XmlChar, ExternalID *XmlChar, SystemID *XmlChar) XmlDtdPtr
//go:linkname XmlNewDtd C.xmlNewDtd
func XmlNewDtd(doc XmlDocPtr, name *XmlChar, ExternalID *XmlChar, SystemID *XmlChar) XmlDtdPtr
// llgo:link (*XmlDoc).XmlGetIntSubset C.xmlGetIntSubset
func (recv_ *XmlDoc) XmlGetIntSubset() XmlDtdPtr {
	return nil
}
//go:linkname XmlFreeDtd C.xmlFreeDtd
func XmlFreeDtd(cur XmlDtdPtr)
//go:linkname XmlNewGlobalNs C.xmlNewGlobalNs
func XmlNewGlobalNs(doc XmlDocPtr, href *XmlChar, prefix *XmlChar) XmlNsPtr
//go:linkname XmlNewNs C.xmlNewNs
func XmlNewNs(node XmlNodePtr, href *XmlChar, prefix *XmlChar) XmlNsPtr
//go:linkname XmlFreeNs C.xmlFreeNs
func XmlFreeNs(cur XmlNsPtr)
//go:linkname XmlFreeNsList C.xmlFreeNsList
func XmlFreeNsList(cur XmlNsPtr)
// llgo:link (*XmlChar).XmlNewDoc C.xmlNewDoc
func (recv_ *XmlChar) XmlNewDoc() XmlDocPtr {
	return nil
}
//go:linkname XmlFreeDoc C.xmlFreeDoc
func XmlFreeDoc(cur XmlDocPtr)
//go:linkname XmlNewDocProp C.xmlNewDocProp
func XmlNewDocProp(doc XmlDocPtr, name *XmlChar, value *XmlChar) XmlAttrPtr
//go:linkname XmlNewProp C.xmlNewProp
func XmlNewProp(node XmlNodePtr, name *XmlChar, value *XmlChar) XmlAttrPtr
//go:linkname XmlNewNsProp C.xmlNewNsProp
func XmlNewNsProp(node XmlNodePtr, ns XmlNsPtr, name *XmlChar, value *XmlChar) XmlAttrPtr
//go:linkname XmlNewNsPropEatName C.xmlNewNsPropEatName
func XmlNewNsPropEatName(node XmlNodePtr, ns XmlNsPtr, name *XmlChar, value *XmlChar) XmlAttrPtr
//go:linkname XmlFreePropList C.xmlFreePropList
func XmlFreePropList(cur XmlAttrPtr)
//go:linkname XmlFreeProp C.xmlFreeProp
func XmlFreeProp(cur XmlAttrPtr)
//go:linkname XmlCopyProp C.xmlCopyProp
func XmlCopyProp(target XmlNodePtr, cur XmlAttrPtr) XmlAttrPtr
//go:linkname XmlCopyPropList C.xmlCopyPropList
func XmlCopyPropList(target XmlNodePtr, cur XmlAttrPtr) XmlAttrPtr
//go:linkname XmlCopyDtd C.xmlCopyDtd
func XmlCopyDtd(dtd XmlDtdPtr) XmlDtdPtr
//go:linkname XmlCopyDoc C.xmlCopyDoc
func XmlCopyDoc(doc XmlDocPtr, recursive c.Int) XmlDocPtr
//go:linkname XmlNewDocNode C.xmlNewDocNode
func XmlNewDocNode(doc XmlDocPtr, ns XmlNsPtr, name *XmlChar, content *XmlChar) XmlNodePtr
//go:linkname XmlNewDocNodeEatName C.xmlNewDocNodeEatName
func XmlNewDocNodeEatName(doc XmlDocPtr, ns XmlNsPtr, name *XmlChar, content *XmlChar) XmlNodePtr
//go:linkname XmlNewNode C.xmlNewNode
func XmlNewNode(ns XmlNsPtr, name *XmlChar) XmlNodePtr
//go:linkname XmlNewNodeEatName C.xmlNewNodeEatName
func XmlNewNodeEatName(ns XmlNsPtr, name *XmlChar) XmlNodePtr
//go:linkname XmlNewChild C.xmlNewChild
func XmlNewChild(parent XmlNodePtr, ns XmlNsPtr, name *XmlChar, content *XmlChar) XmlNodePtr
// llgo:link (*XmlDoc).XmlNewDocText C.xmlNewDocText
func (recv_ *XmlDoc) XmlNewDocText(content *XmlChar) XmlNodePtr {
	return nil
}
// llgo:link (*XmlChar).XmlNewText C.xmlNewText
func (recv_ *XmlChar) XmlNewText() XmlNodePtr {
	return nil
}
//go:linkname XmlNewDocPI C.xmlNewDocPI
func XmlNewDocPI(doc XmlDocPtr, name *XmlChar, content *XmlChar) XmlNodePtr
// llgo:link (*XmlChar).XmlNewPI C.xmlNewPI
func (recv_ *XmlChar) XmlNewPI(content *XmlChar) XmlNodePtr {
	return nil
}
//go:linkname XmlNewDocTextLen C.xmlNewDocTextLen
func XmlNewDocTextLen(doc XmlDocPtr, content *XmlChar, len c.Int) XmlNodePtr
// llgo:link (*XmlChar).XmlNewTextLen C.xmlNewTextLen
func (recv_ *XmlChar) XmlNewTextLen(len c.Int) XmlNodePtr {
	return nil
}
//go:linkname XmlNewDocComment C.xmlNewDocComment
func XmlNewDocComment(doc XmlDocPtr, content *XmlChar) XmlNodePtr
// llgo:link (*XmlChar).XmlNewComment C.xmlNewComment
func (recv_ *XmlChar) XmlNewComment() XmlNodePtr {
	return nil
}
//go:linkname XmlNewCDataBlock C.xmlNewCDataBlock
func XmlNewCDataBlock(doc XmlDocPtr, content *XmlChar, len c.Int) XmlNodePtr
//go:linkname XmlNewCharRef C.xmlNewCharRef
func XmlNewCharRef(doc XmlDocPtr, name *XmlChar) XmlNodePtr
// llgo:link (*XmlDoc).XmlNewReference C.xmlNewReference
func (recv_ *XmlDoc) XmlNewReference(name *XmlChar) XmlNodePtr {
	return nil
}
//go:linkname XmlCopyNode C.xmlCopyNode
func XmlCopyNode(node XmlNodePtr, recursive c.Int) XmlNodePtr
//go:linkname XmlDocCopyNode C.xmlDocCopyNode
func XmlDocCopyNode(node XmlNodePtr, doc XmlDocPtr, recursive c.Int) XmlNodePtr
//go:linkname XmlDocCopyNodeList C.xmlDocCopyNodeList
func XmlDocCopyNodeList(doc XmlDocPtr, node XmlNodePtr) XmlNodePtr
//go:linkname XmlCopyNodeList C.xmlCopyNodeList
func XmlCopyNodeList(node XmlNodePtr) XmlNodePtr
//go:linkname XmlNewTextChild C.xmlNewTextChild
func XmlNewTextChild(parent XmlNodePtr, ns XmlNsPtr, name *XmlChar, content *XmlChar) XmlNodePtr
//go:linkname XmlNewDocRawNode C.xmlNewDocRawNode
func XmlNewDocRawNode(doc XmlDocPtr, ns XmlNsPtr, name *XmlChar, content *XmlChar) XmlNodePtr
//go:linkname XmlNewDocFragment C.xmlNewDocFragment
func XmlNewDocFragment(doc XmlDocPtr) XmlNodePtr
// llgo:link (*XmlNode).XmlGetLineNo C.xmlGetLineNo
func (recv_ *XmlNode) XmlGetLineNo() c.Long {
	return 0
}
// llgo:link (*XmlNode).XmlGetNodePath C.xmlGetNodePath
func (recv_ *XmlNode) XmlGetNodePath() *XmlChar {
	return nil
}
// llgo:link (*XmlDoc).XmlDocGetRootElement C.xmlDocGetRootElement
func (recv_ *XmlDoc) XmlDocGetRootElement() XmlNodePtr {
	return nil
}
// llgo:link (*XmlNode).XmlGetLastChild C.xmlGetLastChild
func (recv_ *XmlNode) XmlGetLastChild() XmlNodePtr {
	return nil
}
// llgo:link (*XmlNode).XmlNodeIsText C.xmlNodeIsText
func (recv_ *XmlNode) XmlNodeIsText() c.Int {
	return 0
}
// llgo:link (*XmlNode).XmlIsBlankNode C.xmlIsBlankNode
func (recv_ *XmlNode) XmlIsBlankNode() c.Int {
	return 0
}
//go:linkname XmlDocSetRootElement C.xmlDocSetRootElement
func XmlDocSetRootElement(doc XmlDocPtr, root XmlNodePtr) XmlNodePtr
//go:linkname XmlNodeSetName C.xmlNodeSetName
func XmlNodeSetName(cur XmlNodePtr, name *XmlChar)
//go:linkname XmlAddChild C.xmlAddChild
func XmlAddChild(parent XmlNodePtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlAddChildList C.xmlAddChildList
func XmlAddChildList(parent XmlNodePtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlReplaceNode C.xmlReplaceNode
func XmlReplaceNode(old XmlNodePtr, cur XmlNodePtr) XmlNodePtr
//go:linkname XmlAddPrevSibling C.xmlAddPrevSibling
func XmlAddPrevSibling(cur XmlNodePtr, elem XmlNodePtr) XmlNodePtr
//go:linkname XmlAddSibling C.xmlAddSibling
func XmlAddSibling(cur XmlNodePtr, elem XmlNodePtr) XmlNodePtr
//go:linkname XmlAddNextSibling C.xmlAddNextSibling
func XmlAddNextSibling(cur XmlNodePtr, elem XmlNodePtr) XmlNodePtr
//go:linkname XmlUnlinkNode C.xmlUnlinkNode
func XmlUnlinkNode(cur XmlNodePtr)
//go:linkname XmlTextMerge C.xmlTextMerge
func XmlTextMerge(first XmlNodePtr, second XmlNodePtr) XmlNodePtr
//go:linkname XmlTextConcat C.xmlTextConcat
func XmlTextConcat(node XmlNodePtr, content *XmlChar, len c.Int) c.Int
//go:linkname XmlFreeNodeList C.xmlFreeNodeList
func XmlFreeNodeList(cur XmlNodePtr)
//go:linkname XmlFreeNode C.xmlFreeNode
func XmlFreeNode(cur XmlNodePtr)
//go:linkname XmlSetTreeDoc C.xmlSetTreeDoc
func XmlSetTreeDoc(tree XmlNodePtr, doc XmlDocPtr) c.Int
//go:linkname XmlSetListDoc C.xmlSetListDoc
func XmlSetListDoc(list XmlNodePtr, doc XmlDocPtr) c.Int
//go:linkname XmlSearchNs C.xmlSearchNs
func XmlSearchNs(doc XmlDocPtr, node XmlNodePtr, nameSpace *XmlChar) XmlNsPtr
//go:linkname XmlSearchNsByHref C.xmlSearchNsByHref
func XmlSearchNsByHref(doc XmlDocPtr, node XmlNodePtr, href *XmlChar) XmlNsPtr
// llgo:link (*XmlDoc).XmlGetNsListSafe C.xmlGetNsListSafe
func (recv_ *XmlDoc) XmlGetNsListSafe(node *XmlNode, out **XmlNsPtr) c.Int {
	return 0
}
// llgo:link (*XmlDoc).XmlGetNsList C.xmlGetNsList
func (recv_ *XmlDoc) XmlGetNsList(node *XmlNode) *XmlNsPtr {
	return nil
}
//go:linkname XmlSetNs C.xmlSetNs
func XmlSetNs(node XmlNodePtr, ns XmlNsPtr)
//go:linkname XmlCopyNamespace C.xmlCopyNamespace
func XmlCopyNamespace(cur XmlNsPtr) XmlNsPtr
//go:linkname XmlCopyNamespaceList C.xmlCopyNamespaceList
func XmlCopyNamespaceList(cur XmlNsPtr) XmlNsPtr
//go:linkname XmlSetProp C.xmlSetProp
func XmlSetProp(node XmlNodePtr, name *XmlChar, value *XmlChar) XmlAttrPtr
//go:linkname XmlSetNsProp C.xmlSetNsProp
func XmlSetNsProp(node XmlNodePtr, ns XmlNsPtr, name *XmlChar, value *XmlChar) XmlAttrPtr
// llgo:link (*XmlNode).XmlNodeGetAttrValue C.xmlNodeGetAttrValue
func (recv_ *XmlNode) XmlNodeGetAttrValue(name *XmlChar, nsUri *XmlChar, out **XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlNode).XmlGetNoNsProp C.xmlGetNoNsProp
func (recv_ *XmlNode) XmlGetNoNsProp(name *XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlNode).XmlGetProp C.xmlGetProp
func (recv_ *XmlNode) XmlGetProp(name *XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlNode).XmlHasProp C.xmlHasProp
func (recv_ *XmlNode) XmlHasProp(name *XmlChar) XmlAttrPtr {
	return nil
}
// llgo:link (*XmlNode).XmlHasNsProp C.xmlHasNsProp
func (recv_ *XmlNode) XmlHasNsProp(name *XmlChar, nameSpace *XmlChar) XmlAttrPtr {
	return nil
}
// llgo:link (*XmlNode).XmlGetNsProp C.xmlGetNsProp
func (recv_ *XmlNode) XmlGetNsProp(name *XmlChar, nameSpace *XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlDoc).XmlStringGetNodeList C.xmlStringGetNodeList
func (recv_ *XmlDoc) XmlStringGetNodeList(value *XmlChar) XmlNodePtr {
	return nil
}
// llgo:link (*XmlDoc).XmlStringLenGetNodeList C.xmlStringLenGetNodeList
func (recv_ *XmlDoc) XmlStringLenGetNodeList(value *XmlChar, len c.Int) XmlNodePtr {
	return nil
}
//go:linkname XmlNodeListGetString C.xmlNodeListGetString
func XmlNodeListGetString(doc XmlDocPtr, list *XmlNode, inLine c.Int) *XmlChar
// llgo:link (*XmlDoc).XmlNodeListGetRawString C.xmlNodeListGetRawString
func (recv_ *XmlDoc) XmlNodeListGetRawString(list *XmlNode, inLine c.Int) *XmlChar {
	return nil
}
//go:linkname XmlNodeSetContent C.xmlNodeSetContent
func XmlNodeSetContent(cur XmlNodePtr, content *XmlChar) c.Int
//go:linkname XmlNodeSetContentLen C.xmlNodeSetContentLen
func XmlNodeSetContentLen(cur XmlNodePtr, content *XmlChar, len c.Int) c.Int
//go:linkname XmlNodeAddContent C.xmlNodeAddContent
func XmlNodeAddContent(cur XmlNodePtr, content *XmlChar) c.Int
//go:linkname XmlNodeAddContentLen C.xmlNodeAddContentLen
func XmlNodeAddContentLen(cur XmlNodePtr, content *XmlChar, len c.Int) c.Int
// llgo:link (*XmlNode).XmlNodeGetContent C.xmlNodeGetContent
func (recv_ *XmlNode) XmlNodeGetContent() *XmlChar {
	return nil
}
//go:linkname XmlNodeBufGetContent C.xmlNodeBufGetContent
func XmlNodeBufGetContent(buffer XmlBufferPtr, cur *XmlNode) c.Int
//go:linkname XmlBufGetNodeContent C.xmlBufGetNodeContent
func XmlBufGetNodeContent(buf XmlBufPtr, cur *XmlNode) c.Int
// llgo:link (*XmlNode).XmlNodeGetLang C.xmlNodeGetLang
func (recv_ *XmlNode) XmlNodeGetLang() *XmlChar {
	return nil
}
// llgo:link (*XmlNode).XmlNodeGetSpacePreserve C.xmlNodeGetSpacePreserve
func (recv_ *XmlNode) XmlNodeGetSpacePreserve() c.Int {
	return 0
}
//go:linkname XmlNodeSetLang C.xmlNodeSetLang
func XmlNodeSetLang(cur XmlNodePtr, lang *XmlChar) c.Int
//go:linkname XmlNodeSetSpacePreserve C.xmlNodeSetSpacePreserve
func XmlNodeSetSpacePreserve(cur XmlNodePtr, val c.Int) c.Int
// llgo:link (*XmlDoc).XmlNodeGetBaseSafe C.xmlNodeGetBaseSafe
func (recv_ *XmlDoc) XmlNodeGetBaseSafe(cur *XmlNode, baseOut **XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlDoc).XmlNodeGetBase C.xmlNodeGetBase
func (recv_ *XmlDoc) XmlNodeGetBase(cur *XmlNode) *XmlChar {
	return nil
}
//go:linkname XmlNodeSetBase C.xmlNodeSetBase
func XmlNodeSetBase(cur XmlNodePtr, uri *XmlChar) c.Int
//go:linkname XmlRemoveProp C.xmlRemoveProp
func XmlRemoveProp(cur XmlAttrPtr) c.Int
//go:linkname XmlUnsetNsProp C.xmlUnsetNsProp
func XmlUnsetNsProp(node XmlNodePtr, ns XmlNsPtr, name *XmlChar) c.Int
//go:linkname XmlUnsetProp C.xmlUnsetProp
func XmlUnsetProp(node XmlNodePtr, name *XmlChar) c.Int
//go:linkname XmlBufferWriteCHAR C.xmlBufferWriteCHAR
func XmlBufferWriteCHAR(buf XmlBufferPtr, string *XmlChar)
//go:linkname XmlBufferWriteChar C.xmlBufferWriteChar
func XmlBufferWriteChar(buf XmlBufferPtr, string *int8)
//go:linkname XmlBufferWriteQuotedString C.xmlBufferWriteQuotedString
func XmlBufferWriteQuotedString(buf XmlBufferPtr, string *XmlChar)
//go:linkname XmlAttrSerializeTxtContent C.xmlAttrSerializeTxtContent
func XmlAttrSerializeTxtContent(buf XmlBufferPtr, doc XmlDocPtr, attr XmlAttrPtr, string *XmlChar)
//go:linkname XmlReconciliateNs C.xmlReconciliateNs
func XmlReconciliateNs(doc XmlDocPtr, tree XmlNodePtr) c.Int
//go:linkname XmlDocDumpFormatMemory C.xmlDocDumpFormatMemory
func XmlDocDumpFormatMemory(cur XmlDocPtr, mem **XmlChar, size *c.Int, format c.Int)
//go:linkname XmlDocDumpMemory C.xmlDocDumpMemory
func XmlDocDumpMemory(cur XmlDocPtr, mem **XmlChar, size *c.Int)
//go:linkname XmlDocDumpMemoryEnc C.xmlDocDumpMemoryEnc
func XmlDocDumpMemoryEnc(out_doc XmlDocPtr, doc_txt_ptr **XmlChar, doc_txt_len *c.Int, txt_encoding *int8)
//go:linkname XmlDocDumpFormatMemoryEnc C.xmlDocDumpFormatMemoryEnc
func XmlDocDumpFormatMemoryEnc(out_doc XmlDocPtr, doc_txt_ptr **XmlChar, doc_txt_len *c.Int, txt_encoding *int8, format c.Int)
//go:linkname XmlDocFormatDump C.xmlDocFormatDump
func XmlDocFormatDump(f *c.FILE, cur XmlDocPtr, format c.Int) c.Int
//go:linkname XmlDocDump C.xmlDocDump
func XmlDocDump(f *c.FILE, cur XmlDocPtr) c.Int
//go:linkname XmlElemDump C.xmlElemDump
func XmlElemDump(f *c.FILE, doc XmlDocPtr, cur XmlNodePtr)
//go:linkname XmlSaveFile C.xmlSaveFile
func XmlSaveFile(filename *int8, cur XmlDocPtr) c.Int
//go:linkname XmlSaveFormatFile C.xmlSaveFormatFile
func XmlSaveFormatFile(filename *int8, cur XmlDocPtr, format c.Int) c.Int
//go:linkname XmlBufNodeDump C.xmlBufNodeDump
func XmlBufNodeDump(buf XmlBufPtr, doc XmlDocPtr, cur XmlNodePtr, level c.Int, format c.Int) uintptr
//go:linkname XmlNodeDump C.xmlNodeDump
func XmlNodeDump(buf XmlBufferPtr, doc XmlDocPtr, cur XmlNodePtr, level c.Int, format c.Int) c.Int
//go:linkname XmlSaveFileTo C.xmlSaveFileTo
func XmlSaveFileTo(buf XmlOutputBufferPtr, cur XmlDocPtr, encoding *int8) c.Int
//go:linkname XmlSaveFormatFileTo C.xmlSaveFormatFileTo
func XmlSaveFormatFileTo(buf XmlOutputBufferPtr, cur XmlDocPtr, encoding *int8, format c.Int) c.Int
//go:linkname XmlNodeDumpOutput C.xmlNodeDumpOutput
func XmlNodeDumpOutput(buf XmlOutputBufferPtr, doc XmlDocPtr, cur XmlNodePtr, level c.Int, format c.Int, encoding *int8)
//go:linkname XmlSaveFormatFileEnc C.xmlSaveFormatFileEnc
func XmlSaveFormatFileEnc(filename *int8, cur XmlDocPtr, encoding *int8, format c.Int) c.Int
//go:linkname XmlSaveFileEnc C.xmlSaveFileEnc
func XmlSaveFileEnc(filename *int8, cur XmlDocPtr, encoding *int8) c.Int
// llgo:link (*XmlChar).XmlIsXHTML C.xmlIsXHTML
func (recv_ *XmlChar) XmlIsXHTML(publicID *XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlDoc).XmlGetDocCompressMode C.xmlGetDocCompressMode
func (recv_ *XmlDoc) XmlGetDocCompressMode() c.Int {
	return 0
}
//go:linkname XmlSetDocCompressMode C.xmlSetDocCompressMode
func XmlSetDocCompressMode(doc XmlDocPtr, mode c.Int)
//go:linkname XmlGetCompressMode C.xmlGetCompressMode
func XmlGetCompressMode() c.Int
//go:linkname XmlSetCompressMode C.xmlSetCompressMode
func XmlSetCompressMode(mode c.Int)
//go:linkname XmlDOMWrapNewCtxt C.xmlDOMWrapNewCtxt
func XmlDOMWrapNewCtxt() XmlDOMWrapCtxtPtr
//go:linkname XmlDOMWrapFreeCtxt C.xmlDOMWrapFreeCtxt
func XmlDOMWrapFreeCtxt(ctxt XmlDOMWrapCtxtPtr)
//go:linkname XmlDOMWrapReconcileNamespaces C.xmlDOMWrapReconcileNamespaces
func XmlDOMWrapReconcileNamespaces(ctxt XmlDOMWrapCtxtPtr, elem XmlNodePtr, options c.Int) c.Int
//go:linkname XmlDOMWrapAdoptNode C.xmlDOMWrapAdoptNode
func XmlDOMWrapAdoptNode(ctxt XmlDOMWrapCtxtPtr, sourceDoc XmlDocPtr, node XmlNodePtr, destDoc XmlDocPtr, destParent XmlNodePtr, options c.Int) c.Int
//go:linkname XmlDOMWrapRemoveNode C.xmlDOMWrapRemoveNode
func XmlDOMWrapRemoveNode(ctxt XmlDOMWrapCtxtPtr, doc XmlDocPtr, node XmlNodePtr, options c.Int) c.Int
//go:linkname XmlDOMWrapCloneNode C.xmlDOMWrapCloneNode
func XmlDOMWrapCloneNode(ctxt XmlDOMWrapCtxtPtr, sourceDoc XmlDocPtr, node XmlNodePtr, clonedNode *XmlNodePtr, destDoc XmlDocPtr, destParent XmlNodePtr, deep c.Int, options c.Int) c.Int
//go:linkname XmlChildElementCount C.xmlChildElementCount
func XmlChildElementCount(parent XmlNodePtr) c.Ulong
//go:linkname XmlNextElementSibling C.xmlNextElementSibling
func XmlNextElementSibling(node XmlNodePtr) XmlNodePtr
//go:linkname XmlFirstElementChild C.xmlFirstElementChild
func XmlFirstElementChild(parent XmlNodePtr) XmlNodePtr
//go:linkname XmlLastElementChild C.xmlLastElementChild
func XmlLastElementChild(parent XmlNodePtr) XmlNodePtr
//go:linkname XmlPreviousElementSibling C.xmlPreviousElementSibling
func XmlPreviousElementSibling(node XmlNodePtr) XmlNodePtr
//go:linkname XmlRegisterNodeDefault C.xmlRegisterNodeDefault
func XmlRegisterNodeDefault(func_ XmlRegisterNodeFunc) XmlRegisterNodeFunc
//go:linkname XmlDeregisterNodeDefault C.xmlDeregisterNodeDefault
func XmlDeregisterNodeDefault(func_ XmlDeregisterNodeFunc) XmlDeregisterNodeFunc
//go:linkname XmlThrDefRegisterNodeDefault C.xmlThrDefRegisterNodeDefault
func XmlThrDefRegisterNodeDefault(func_ XmlRegisterNodeFunc) XmlRegisterNodeFunc
//go:linkname XmlThrDefDeregisterNodeDefault C.xmlThrDefDeregisterNodeDefault
func XmlThrDefDeregisterNodeDefault(func_ XmlDeregisterNodeFunc) XmlDeregisterNodeFunc
// llgo:link XmlBufferAllocationScheme.XmlThrDefBufferAllocScheme C.xmlThrDefBufferAllocScheme
func (recv_ XmlBufferAllocationScheme) XmlThrDefBufferAllocScheme() XmlBufferAllocationScheme {
	return 0
}
//go:linkname XmlThrDefDefaultBufferSize C.xmlThrDefDefaultBufferSize
func XmlThrDefDefaultBufferSize(v c.Int) c.Int
