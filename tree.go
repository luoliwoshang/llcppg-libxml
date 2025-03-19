package libxml2

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const BASE_BUFFER_SIZE = 4096
const DOCB_DOCUMENT_NODE = 21

type X_xmlParserInputBuffer struct {
	Context       unsafe.Pointer
	Readcallback  unsafe.Pointer
	Closecallback unsafe.Pointer
	Encoder       CharEncodingHandlerPtr
	Buffer        BufPtr
	Raw           BufPtr
	Compressed    c.Int
	Error         c.Int
	Rawconsumed   c.Ulong
}
type ParserInputBuffer X_xmlParserInputBuffer
type ParserInputBufferPtr *ParserInputBuffer

type X_xmlOutputBuffer struct {
	Context       unsafe.Pointer
	Writecallback unsafe.Pointer
	Closecallback unsafe.Pointer
	Encoder       CharEncodingHandlerPtr
	Buffer        BufPtr
	Conv          BufPtr
	Written       c.Int
	Error         c.Int
}
type OutputBuffer X_xmlOutputBuffer
type OutputBufferPtr *OutputBuffer

type X_xmlParserInput struct {
	Buf            ParserInputBufferPtr
	Filename       *int8
	Directory      *int8
	Base           *Char
	Cur            *Char
	End            *Char
	Length         c.Int
	Line           c.Int
	Col            c.Int
	Consumed       c.Ulong
	Free           unsafe.Pointer
	Encoding       *Char
	Version        *Char
	Flags          c.Int
	Id             c.Int
	ParentConsumed c.Ulong
	Entity         EntityPtr
}
type ParserInput X_xmlParserInput
type ParserInputPtr *ParserInput

type X_xmlParserCtxt struct {
	Sax               *X_xmlSAXHandler
	UserData          unsafe.Pointer
	MyDoc             DocPtr
	WellFormed        c.Int
	ReplaceEntities   c.Int
	Version           *Char
	Encoding          *Char
	Standalone        c.Int
	Html              c.Int
	Input             ParserInputPtr
	InputNr           c.Int
	InputMax          c.Int
	InputTab          *ParserInputPtr
	Node              NodePtr
	NodeNr            c.Int
	NodeMax           c.Int
	NodeTab           *NodePtr
	RecordInfo        c.Int
	NodeSeq           ParserNodeInfoSeq
	ErrNo             c.Int
	HasExternalSubset c.Int
	HasPErefs         c.Int
	External          c.Int
	Valid             c.Int
	Validate          c.Int
	Vctxt             ValidCtxt
	Instate           ParserInputState
	Token             c.Int
	Directory         *int8
	Name              *Char
	NameNr            c.Int
	NameMax           c.Int
	NameTab           **Char
	NbChars           c.Long
	CheckIndex        c.Long
	KeepBlanks        c.Int
	DisableSAX        c.Int
	InSubset          c.Int
	IntSubName        *Char
	ExtSubURI         *Char
	ExtSubSystem      *Char
	Space             *c.Int
	SpaceNr           c.Int
	SpaceMax          c.Int
	SpaceTab          *c.Int
	Depth             c.Int
	Entity            ParserInputPtr
	Charset           c.Int
	Nodelen           c.Int
	Nodemem           c.Int
	Pedantic          c.Int
	X_private         unsafe.Pointer
	Loadsubset        c.Int
	Linenumbers       c.Int
	Catalogs          unsafe.Pointer
	Recovery          c.Int
	Progressive       c.Int
	Dict              DictPtr
	Atts              **Char
	Maxatts           c.Int
	Docdict           c.Int
	StrXml            *Char
	StrXmlns          *Char
	StrXmlNs          *Char
	Sax2              c.Int
	NsNr              c.Int
	NsMax             c.Int
	NsTab             **Char
	Attallocs         *c.Uint
	PushTab           *StartTag
	AttsDefault       HashTablePtr
	AttsSpecial       HashTablePtr
	NsWellFormed      c.Int
	Options           c.Int
	DictNames         c.Int
	FreeElemsNr       c.Int
	FreeElems         NodePtr
	FreeAttrsNr       c.Int
	FreeAttrs         AttrPtr
	LastError         Error
	ParseMode         ParserMode
	Nbentities        c.Ulong
	Sizeentities      c.Ulong
	NodeInfo          *ParserNodeInfo
	NodeInfoNr        c.Int
	NodeInfoMax       c.Int
	NodeInfoTab       *ParserNodeInfo
	InputId           c.Int
	Sizeentcopy       c.Ulong
	EndCheckState     c.Int
	NbErrors          uint16
	NbWarnings        uint16
	MaxAmpl           c.Uint
	Nsdb              *ParserNsData
	AttrHashMax       c.Uint
	AttrHash          *AttrHashBucket
	ErrorHandler      unsafe.Pointer
	ErrorCtxt         unsafe.Pointer
}
type ParserCtxt X_xmlParserCtxt
type ParserCtxtPtr *ParserCtxt

type X_xmlSAXLocator struct {
	GetPublicId     unsafe.Pointer
	GetSystemId     unsafe.Pointer
	GetLineNumber   unsafe.Pointer
	GetColumnNumber unsafe.Pointer
}
type SAXLocator X_xmlSAXLocator
type SAXLocatorPtr *SAXLocator

type X_xmlSAXHandler struct {
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
	X_private             unsafe.Pointer
	StartElementNs        unsafe.Pointer
	EndElementNs          unsafe.Pointer
	Serror                unsafe.Pointer
}
type SAXHandler X_xmlSAXHandler
type SAXHandlerPtr *SAXHandler

type X_xmlEntity struct {
	X_private    unsafe.Pointer
	Type         ElementType
	Name         *Char
	Children     *X_xmlNode
	Last         *X_xmlNode
	Parent       *X_xmlDtd
	Next         *X_xmlNode
	Prev         *X_xmlNode
	Doc          *X_xmlDoc
	Orig         *Char
	Content      *Char
	Length       c.Int
	Etype        EntityType
	ExternalID   *Char
	SystemID     *Char
	Nexte        *X_xmlEntity
	URI          *Char
	Owner        c.Int
	Flags        c.Int
	ExpandedSize c.Ulong
}
type Entity X_xmlEntity
type EntityPtr *Entity
type BufferAllocationScheme c.Int

const (
	BUFFERALLOCDOUBLEIT  BufferAllocationScheme = 0
	BUFFERALLOCEXACT     BufferAllocationScheme = 1
	BUFFERALLOCIMMUTABLE BufferAllocationScheme = 2
	BUFFERALLOCIO        BufferAllocationScheme = 3
	BUFFERALLOCHYBRID    BufferAllocationScheme = 4
	BUFFERALLOCBOUNDED   BufferAllocationScheme = 5
)

type X_xmlBuffer struct {
	Content   *Char
	Use       c.Uint
	Size      c.Uint
	Alloc     BufferAllocationScheme
	ContentIO *Char
}
type Buffer X_xmlBuffer
type BufferPtr *Buffer

type X_xmlBuf struct {
	Unused [8]uint8
}
type Buf X_xmlBuf
type BufPtr *Buf

/*
 * A few public routines for xmlBuf. As those are expected to be used
 * mostly internally the bulk of the routines are internal in buf.h
 */
// llgo:link (*Buf).BufContent C.xmlBufContent
func (recv_ *Buf) BufContent() *Char {
	return nil
}

//go:linkname BufEnd C.xmlBufEnd
func BufEnd(buf BufPtr) *Char

//go:linkname BufUse C.xmlBufUse
func BufUse(buf BufPtr) uintptr

//go:linkname BufShrink C.xmlBufShrink
func BufShrink(buf BufPtr, len uintptr) uintptr

type ElementType c.Int

const (
	ELEMENTNODE      ElementType = 1
	ATTRIBUTENODE    ElementType = 2
	TEXTNODE         ElementType = 3
	CDATASECTIONNODE ElementType = 4
	ENTITYREFNODE    ElementType = 5
	ENTITYNODE       ElementType = 6
	PINODE           ElementType = 7
	COMMENTNODE      ElementType = 8
	DOCUMENTNODE     ElementType = 9
	DOCUMENTTYPENODE ElementType = 10
	DOCUMENTFRAGNODE ElementType = 11
	NOTATIONNODE     ElementType = 12
	HTMLDOCUMENTNODE ElementType = 13
	DTDNODE          ElementType = 14
	ELEMENTDECL      ElementType = 15
	ATTRIBUTEDECL    ElementType = 16
	ENTITYDECL       ElementType = 17
	NAMESPACEDECL    ElementType = 18
	XINCLUDESTART    ElementType = 19
	XINCLUDEEND      ElementType = 20
)

type X_xmlNotation struct {
	Name     *Char
	PublicID *Char
	SystemID *Char
}
type Notation X_xmlNotation
type NotationPtr *Notation
type AttributeType c.Int

const (
	ATTRIBUTECDATA       AttributeType = 1
	ATTRIBUTEID          AttributeType = 2
	ATTRIBUTEIDREF       AttributeType = 3
	ATTRIBUTEIDREFS      AttributeType = 4
	ATTRIBUTEENTITY      AttributeType = 5
	ATTRIBUTEENTITIES    AttributeType = 6
	ATTRIBUTENMTOKEN     AttributeType = 7
	ATTRIBUTENMTOKENS    AttributeType = 8
	ATTRIBUTEENUMERATION AttributeType = 9
	ATTRIBUTENOTATION    AttributeType = 10
)

type AttributeDefault c.Int

const (
	ATTRIBUTENONE     AttributeDefault = 1
	ATTRIBUTEREQUIRED AttributeDefault = 2
	ATTRIBUTEIMPLIED  AttributeDefault = 3
	ATTRIBUTEFIXED    AttributeDefault = 4
)

type X_xmlEnumeration struct {
	Next *X_xmlEnumeration
	Name *Char
}
type Enumeration X_xmlEnumeration
type EnumerationPtr *Enumeration

type X_xmlAttribute struct {
	X_private    unsafe.Pointer
	Type         ElementType
	Name         *Char
	Children     *X_xmlNode
	Last         *X_xmlNode
	Parent       *X_xmlDtd
	Next         *X_xmlNode
	Prev         *X_xmlNode
	Doc          *X_xmlDoc
	Nexth        *X_xmlAttribute
	Atype        AttributeType
	Def          AttributeDefault
	DefaultValue *Char
	Tree         EnumerationPtr
	Prefix       *Char
	Elem         *Char
}
type Attribute X_xmlAttribute
type AttributePtr *Attribute

type X_xmlNode struct {
	X_private  unsafe.Pointer
	Type       ElementType
	Name       *Char
	Children   *X_xmlNode
	Last       *X_xmlNode
	Parent     *X_xmlNode
	Next       *X_xmlNode
	Prev       *X_xmlNode
	Doc        *X_xmlDoc
	Ns         *Ns
	Content    *Char
	Properties *X_xmlAttr
	NsDef      *Ns
	Psvi       unsafe.Pointer
	Line       uint16
	Extra      uint16
}

type X_xmlDtd struct {
	X_private  unsafe.Pointer
	Type       ElementType
	Name       *Char
	Children   *X_xmlNode
	Last       *X_xmlNode
	Parent     *X_xmlDoc
	Next       *X_xmlNode
	Prev       *X_xmlNode
	Doc        *X_xmlDoc
	Notations  unsafe.Pointer
	Elements   unsafe.Pointer
	Attributes unsafe.Pointer
	Entities   unsafe.Pointer
	ExternalID *Char
	SystemID   *Char
	Pentities  unsafe.Pointer
}

type X_xmlDoc struct {
	X_private   unsafe.Pointer
	Type        ElementType
	Name        *int8
	Children    *X_xmlNode
	Last        *X_xmlNode
	Parent      *X_xmlNode
	Next        *X_xmlNode
	Prev        *X_xmlNode
	Doc         *X_xmlDoc
	Compression c.Int
	Standalone  c.Int
	IntSubset   *X_xmlDtd
	ExtSubset   *X_xmlDtd
	OldNs       *X_xmlNs
	Version     *Char
	Encoding    *Char
	Ids         unsafe.Pointer
	Refs        unsafe.Pointer
	URL         *Char
	Charset     c.Int
	Dict        *X_xmlDict
	Psvi        unsafe.Pointer
	ParseFlags  c.Int
	Properties  c.Int
}
type ElementContentType c.Int

const (
	ELEMENTCONTENTPCDATA  ElementContentType = 1
	ELEMENTCONTENTELEMENT ElementContentType = 2
	ELEMENTCONTENTSEQ     ElementContentType = 3
	ELEMENTCONTENTOR      ElementContentType = 4
)

type ElementContentOccur c.Int

const (
	ELEMENTCONTENTONCE ElementContentOccur = 1
	ELEMENTCONTENTOPT  ElementContentOccur = 2
	ELEMENTCONTENTMULT ElementContentOccur = 3
	ELEMENTCONTENTPLUS ElementContentOccur = 4
)

type X_xmlElementContent struct {
	Type   ElementContentType
	Ocur   ElementContentOccur
	Name   *Char
	C1     *X_xmlElementContent
	C2     *X_xmlElementContent
	Parent *X_xmlElementContent
	Prefix *Char
}
type ElementContent X_xmlElementContent
type ElementContentPtr *ElementContent
type ElementTypeVal c.Int

const (
	ELEMENTTYPEUNDEFINED ElementTypeVal = 0
	ELEMENTTYPEEMPTY     ElementTypeVal = 1
	ELEMENTTYPEANY       ElementTypeVal = 2
	ELEMENTTYPEMIXED     ElementTypeVal = 3
	ELEMENTTYPEELEMENT   ElementTypeVal = 4
)

type X_xmlElement struct {
	X_private  unsafe.Pointer
	Type       ElementType
	Name       *Char
	Children   *X_xmlNode
	Last       *X_xmlNode
	Parent     *X_xmlDtd
	Next       *X_xmlNode
	Prev       *X_xmlNode
	Doc        *X_xmlDoc
	Etype      ElementTypeVal
	Content    ElementContentPtr
	Attributes AttributePtr
	Prefix     *Char
	ContModel  RegexpPtr
}
type Element X_xmlElement
type ElementPtr *Element
type NsType ElementType

type X_xmlNs struct {
	Next      *X_xmlNs
	Type      NsType
	Href      *Char
	Prefix    *Char
	X_private unsafe.Pointer
	Context   *X_xmlDoc
}
type Ns X_xmlNs
type NsPtr *Ns
type Dtd X_xmlDtd
type DtdPtr *Dtd

type X_xmlAttr struct {
	X_private unsafe.Pointer
	Type      ElementType
	Name      *Char
	Children  *X_xmlNode
	Last      *X_xmlNode
	Parent    *X_xmlNode
	Next      *X_xmlAttr
	Prev      *X_xmlAttr
	Doc       *X_xmlDoc
	Ns        *Ns
	Atype     AttributeType
	Psvi      unsafe.Pointer
	Id        *X_xmlID
}
type Attr X_xmlAttr
type AttrPtr *Attr

type X_xmlID struct {
	Next   *X_xmlID
	Value  *Char
	Attr   AttrPtr
	Name   *Char
	Lineno c.Int
	Doc    *X_xmlDoc
}
type ID X_xmlID
type IDPtr *ID

type X_xmlRef struct {
	Next   *X_xmlRef
	Value  *Char
	Attr   AttrPtr
	Name   *Char
	Lineno c.Int
}
type Ref X_xmlRef
type RefPtr *Ref
type Node X_xmlNode
type NodePtr *Node
type DocProperties c.Int

const (
	DOCWELLFORMED DocProperties = 1
	DOCNSVALID    DocProperties = 2
	DOCOLD10      DocProperties = 4
	DOCDTDVALID   DocProperties = 8
	DOCXINCLUDE   DocProperties = 16
	DOCUSERBUILT  DocProperties = 32
	DOCINTERNAL   DocProperties = 64
	DOCHTML       DocProperties = 128
)

type Doc X_xmlDoc
type DocPtr *Doc

type X_xmlDict struct {
	Unused [8]uint8
}

type X_xmlDOMWrapCtxt struct {
	X_private        unsafe.Pointer
	Type             c.Int
	NamespaceMap     unsafe.Pointer
	GetNsForNodeFunc unsafe.Pointer
}
type DOMWrapCtxt X_xmlDOMWrapCtxt
type DOMWrapCtxtPtr *DOMWrapCtxt

// llgo:type C
type DOMWrapAcquireNsFunction func(DOMWrapCtxtPtr, NodePtr, *Char, *Char) NsPtr

// llgo:type C
type RegisterNodeFunc func(NodePtr)

// llgo:type C
type DeregisterNodeFunc func(NodePtr)

//go:linkname X__xmlBufferAllocScheme C.__xmlBufferAllocScheme
func X__xmlBufferAllocScheme() *BufferAllocationScheme

//go:linkname X__xmlDefaultBufferSize C.__xmlDefaultBufferSize
func X__xmlDefaultBufferSize() *c.Int

//go:linkname X__xmlRegisterNodeDefaultValue C.__xmlRegisterNodeDefaultValue
func X__xmlRegisterNodeDefaultValue() RegisterNodeFunc

//go:linkname X__xmlDeregisterNodeDefaultValue C.__xmlDeregisterNodeDefaultValue
func X__xmlDeregisterNodeDefaultValue() DeregisterNodeFunc

/*
 * Some helper functions
 */
// llgo:link (*Char).ValidateNCName C.xmlValidateNCName
func (recv_ *Char) ValidateNCName(space c.Int) c.Int {
	return 0
}

// llgo:link (*Char).ValidateQName C.xmlValidateQName
func (recv_ *Char) ValidateQName(space c.Int) c.Int {
	return 0
}

// llgo:link (*Char).ValidateName C.xmlValidateName
func (recv_ *Char) ValidateName(space c.Int) c.Int {
	return 0
}

// llgo:link (*Char).ValidateNMToken C.xmlValidateNMToken
func (recv_ *Char) ValidateNMToken(space c.Int) c.Int {
	return 0
}

// llgo:link (*Char).BuildQName C.xmlBuildQName
func (recv_ *Char) BuildQName(prefix *Char, memory *Char, len c.Int) *Char {
	return nil
}

// llgo:link (*Char).SplitQName2 C.xmlSplitQName2
func (recv_ *Char) SplitQName2(prefix **Char) *Char {
	return nil
}

// llgo:link (*Char).SplitQName3 C.xmlSplitQName3
func (recv_ *Char) SplitQName3(len *c.Int) *Char {
	return nil
}

/*
 * Handling Buffers, the old ones see @xmlBuf for the new ones.
 */
// llgo:link BufferAllocationScheme.SetBufferAllocationScheme C.xmlSetBufferAllocationScheme
func (recv_ BufferAllocationScheme) SetBufferAllocationScheme() {
}

//go:linkname GetBufferAllocationScheme C.xmlGetBufferAllocationScheme
func GetBufferAllocationScheme() BufferAllocationScheme

//go:linkname BufferCreate C.xmlBufferCreate
func BufferCreate() BufferPtr

//go:linkname BufferCreateSize C.xmlBufferCreateSize
func BufferCreateSize(size uintptr) BufferPtr

//go:linkname BufferCreateStatic C.xmlBufferCreateStatic
func BufferCreateStatic(mem unsafe.Pointer, size uintptr) BufferPtr

//go:linkname BufferResize C.xmlBufferResize
func BufferResize(buf BufferPtr, size c.Uint) c.Int

//go:linkname BufferFree C.xmlBufferFree
func BufferFree(buf BufferPtr)

//go:linkname BufferDump C.xmlBufferDump
func BufferDump(file *c.FILE, buf BufferPtr) c.Int

//go:linkname BufferAdd C.xmlBufferAdd
func BufferAdd(buf BufferPtr, str *Char, len c.Int) c.Int

//go:linkname BufferAddHead C.xmlBufferAddHead
func BufferAddHead(buf BufferPtr, str *Char, len c.Int) c.Int

//go:linkname BufferCat C.xmlBufferCat
func BufferCat(buf BufferPtr, str *Char) c.Int

//go:linkname BufferCCat C.xmlBufferCCat
func BufferCCat(buf BufferPtr, str *int8) c.Int

//go:linkname BufferShrink C.xmlBufferShrink
func BufferShrink(buf BufferPtr, len c.Uint) c.Int

//go:linkname BufferGrow C.xmlBufferGrow
func BufferGrow(buf BufferPtr, len c.Uint) c.Int

//go:linkname BufferEmpty C.xmlBufferEmpty
func BufferEmpty(buf BufferPtr)

// llgo:link (*Buffer).BufferContent C.xmlBufferContent
func (recv_ *Buffer) BufferContent() *Char {
	return nil
}

//go:linkname BufferDetach C.xmlBufferDetach
func BufferDetach(buf BufferPtr) *Char

//go:linkname BufferSetAllocationScheme C.xmlBufferSetAllocationScheme
func BufferSetAllocationScheme(buf BufferPtr, scheme BufferAllocationScheme)

// llgo:link (*Buffer).BufferLength C.xmlBufferLength
func (recv_ *Buffer) BufferLength() c.Int {
	return 0
}

/*
 * Creating/freeing new structures.
 */
//go:linkname CreateIntSubset C.xmlCreateIntSubset
func CreateIntSubset(doc DocPtr, name *Char, ExternalID *Char, SystemID *Char) DtdPtr

//go:linkname NewDtd C.xmlNewDtd
func NewDtd(doc DocPtr, name *Char, ExternalID *Char, SystemID *Char) DtdPtr

// llgo:link (*Doc).GetIntSubset C.xmlGetIntSubset
func (recv_ *Doc) GetIntSubset() DtdPtr {
	return nil
}

//go:linkname FreeDtd C.xmlFreeDtd
func FreeDtd(cur DtdPtr)

//go:linkname NewGlobalNs C.xmlNewGlobalNs
func NewGlobalNs(doc DocPtr, href *Char, prefix *Char) NsPtr

//go:linkname NewNs C.xmlNewNs
func NewNs(node NodePtr, href *Char, prefix *Char) NsPtr

//go:linkname FreeNs C.xmlFreeNs
func FreeNs(cur NsPtr)

//go:linkname FreeNsList C.xmlFreeNsList
func FreeNsList(cur NsPtr)

// llgo:link (*Char).NewDoc C.xmlNewDoc
func (recv_ *Char) NewDoc() DocPtr {
	return nil
}

//go:linkname FreeDoc C.xmlFreeDoc
func FreeDoc(cur DocPtr)

//go:linkname NewDocProp C.xmlNewDocProp
func NewDocProp(doc DocPtr, name *Char, value *Char) AttrPtr

//go:linkname NewProp C.xmlNewProp
func NewProp(node NodePtr, name *Char, value *Char) AttrPtr

//go:linkname NewNsProp C.xmlNewNsProp
func NewNsProp(node NodePtr, ns NsPtr, name *Char, value *Char) AttrPtr

//go:linkname NewNsPropEatName C.xmlNewNsPropEatName
func NewNsPropEatName(node NodePtr, ns NsPtr, name *Char, value *Char) AttrPtr

//go:linkname FreePropList C.xmlFreePropList
func FreePropList(cur AttrPtr)

//go:linkname FreeProp C.xmlFreeProp
func FreeProp(cur AttrPtr)

//go:linkname CopyProp C.xmlCopyProp
func CopyProp(target NodePtr, cur AttrPtr) AttrPtr

//go:linkname CopyPropList C.xmlCopyPropList
func CopyPropList(target NodePtr, cur AttrPtr) AttrPtr

//go:linkname CopyDtd C.xmlCopyDtd
func CopyDtd(dtd DtdPtr) DtdPtr

//go:linkname CopyDoc C.xmlCopyDoc
func CopyDoc(doc DocPtr, recursive c.Int) DocPtr

/*
 * Creating new nodes.
 */
//go:linkname NewDocNode C.xmlNewDocNode
func NewDocNode(doc DocPtr, ns NsPtr, name *Char, content *Char) NodePtr

//go:linkname NewDocNodeEatName C.xmlNewDocNodeEatName
func NewDocNodeEatName(doc DocPtr, ns NsPtr, name *Char, content *Char) NodePtr

//go:linkname NewNode C.xmlNewNode
func NewNode(ns NsPtr, name *Char) NodePtr

//go:linkname NewNodeEatName C.xmlNewNodeEatName
func NewNodeEatName(ns NsPtr, name *Char) NodePtr

//go:linkname NewChild C.xmlNewChild
func NewChild(parent NodePtr, ns NsPtr, name *Char, content *Char) NodePtr

// llgo:link (*Doc).NewDocText C.xmlNewDocText
func (recv_ *Doc) NewDocText(content *Char) NodePtr {
	return nil
}

// llgo:link (*Char).NewText C.xmlNewText
func (recv_ *Char) NewText() NodePtr {
	return nil
}

//go:linkname NewDocPI C.xmlNewDocPI
func NewDocPI(doc DocPtr, name *Char, content *Char) NodePtr

// llgo:link (*Char).NewPI C.xmlNewPI
func (recv_ *Char) NewPI(content *Char) NodePtr {
	return nil
}

//go:linkname NewDocTextLen C.xmlNewDocTextLen
func NewDocTextLen(doc DocPtr, content *Char, len c.Int) NodePtr

// llgo:link (*Char).NewTextLen C.xmlNewTextLen
func (recv_ *Char) NewTextLen(len c.Int) NodePtr {
	return nil
}

//go:linkname NewDocComment C.xmlNewDocComment
func NewDocComment(doc DocPtr, content *Char) NodePtr

// llgo:link (*Char).NewComment C.xmlNewComment
func (recv_ *Char) NewComment() NodePtr {
	return nil
}

//go:linkname NewCDataBlock C.xmlNewCDataBlock
func NewCDataBlock(doc DocPtr, content *Char, len c.Int) NodePtr

//go:linkname NewCharRef C.xmlNewCharRef
func NewCharRef(doc DocPtr, name *Char) NodePtr

// llgo:link (*Doc).NewReference C.xmlNewReference
func (recv_ *Doc) NewReference(name *Char) NodePtr {
	return nil
}

//go:linkname CopyNode C.xmlCopyNode
func CopyNode(node NodePtr, recursive c.Int) NodePtr

//go:linkname DocCopyNode C.xmlDocCopyNode
func DocCopyNode(node NodePtr, doc DocPtr, recursive c.Int) NodePtr

//go:linkname DocCopyNodeList C.xmlDocCopyNodeList
func DocCopyNodeList(doc DocPtr, node NodePtr) NodePtr

//go:linkname CopyNodeList C.xmlCopyNodeList
func CopyNodeList(node NodePtr) NodePtr

//go:linkname NewTextChild C.xmlNewTextChild
func NewTextChild(parent NodePtr, ns NsPtr, name *Char, content *Char) NodePtr

//go:linkname NewDocRawNode C.xmlNewDocRawNode
func NewDocRawNode(doc DocPtr, ns NsPtr, name *Char, content *Char) NodePtr

//go:linkname NewDocFragment C.xmlNewDocFragment
func NewDocFragment(doc DocPtr) NodePtr

/*
 * Navigating.
 */
// llgo:link (*Node).GetLineNo C.xmlGetLineNo
func (recv_ *Node) GetLineNo() c.Long {
	return 0
}

// llgo:link (*Node).GetNodePath C.xmlGetNodePath
func (recv_ *Node) GetNodePath() *Char {
	return nil
}

// llgo:link (*Doc).DocGetRootElement C.xmlDocGetRootElement
func (recv_ *Doc) DocGetRootElement() NodePtr {
	return nil
}

// llgo:link (*Node).GetLastChild C.xmlGetLastChild
func (recv_ *Node) GetLastChild() NodePtr {
	return nil
}

// llgo:link (*Node).NodeIsText C.xmlNodeIsText
func (recv_ *Node) NodeIsText() c.Int {
	return 0
}

// llgo:link (*Node).IsBlankNode C.xmlIsBlankNode
func (recv_ *Node) IsBlankNode() c.Int {
	return 0
}

/*
 * Changing the structure.
 */
//go:linkname DocSetRootElement C.xmlDocSetRootElement
func DocSetRootElement(doc DocPtr, root NodePtr) NodePtr

//go:linkname NodeSetName C.xmlNodeSetName
func NodeSetName(cur NodePtr, name *Char)

//go:linkname AddChild C.xmlAddChild
func AddChild(parent NodePtr, cur NodePtr) NodePtr

//go:linkname AddChildList C.xmlAddChildList
func AddChildList(parent NodePtr, cur NodePtr) NodePtr

//go:linkname ReplaceNode C.xmlReplaceNode
func ReplaceNode(old NodePtr, cur NodePtr) NodePtr

//go:linkname AddPrevSibling C.xmlAddPrevSibling
func AddPrevSibling(cur NodePtr, elem NodePtr) NodePtr

//go:linkname AddSibling C.xmlAddSibling
func AddSibling(cur NodePtr, elem NodePtr) NodePtr

//go:linkname AddNextSibling C.xmlAddNextSibling
func AddNextSibling(cur NodePtr, elem NodePtr) NodePtr

//go:linkname UnlinkNode C.xmlUnlinkNode
func UnlinkNode(cur NodePtr)

//go:linkname TextMerge C.xmlTextMerge
func TextMerge(first NodePtr, second NodePtr) NodePtr

//go:linkname TextConcat C.xmlTextConcat
func TextConcat(node NodePtr, content *Char, len c.Int) c.Int

//go:linkname FreeNodeList C.xmlFreeNodeList
func FreeNodeList(cur NodePtr)

//go:linkname FreeNode C.xmlFreeNode
func FreeNode(cur NodePtr)

//go:linkname SetTreeDoc C.xmlSetTreeDoc
func SetTreeDoc(tree NodePtr, doc DocPtr) c.Int

//go:linkname SetListDoc C.xmlSetListDoc
func SetListDoc(list NodePtr, doc DocPtr) c.Int

/*
 * Namespaces.
 */
//go:linkname SearchNs C.xmlSearchNs
func SearchNs(doc DocPtr, node NodePtr, nameSpace *Char) NsPtr

//go:linkname SearchNsByHref C.xmlSearchNsByHref
func SearchNsByHref(doc DocPtr, node NodePtr, href *Char) NsPtr

// llgo:link (*Doc).GetNsListSafe C.xmlGetNsListSafe
func (recv_ *Doc) GetNsListSafe(node *Node, out **NsPtr) c.Int {
	return 0
}

// llgo:link (*Doc).GetNsList C.xmlGetNsList
func (recv_ *Doc) GetNsList(node *Node) *NsPtr {
	return nil
}

//go:linkname SetNs C.xmlSetNs
func SetNs(node NodePtr, ns NsPtr)

//go:linkname CopyNamespace C.xmlCopyNamespace
func CopyNamespace(cur NsPtr) NsPtr

//go:linkname CopyNamespaceList C.xmlCopyNamespaceList
func CopyNamespaceList(cur NsPtr) NsPtr

/*
 * Changing the content.
 */
//go:linkname SetProp C.xmlSetProp
func SetProp(node NodePtr, name *Char, value *Char) AttrPtr

//go:linkname SetNsProp C.xmlSetNsProp
func SetNsProp(node NodePtr, ns NsPtr, name *Char, value *Char) AttrPtr

// llgo:link (*Node).NodeGetAttrValue C.xmlNodeGetAttrValue
func (recv_ *Node) NodeGetAttrValue(name *Char, nsUri *Char, out **Char) c.Int {
	return 0
}

// llgo:link (*Node).GetNoNsProp C.xmlGetNoNsProp
func (recv_ *Node) GetNoNsProp(name *Char) *Char {
	return nil
}

// llgo:link (*Node).GetProp C.xmlGetProp
func (recv_ *Node) GetProp(name *Char) *Char {
	return nil
}

// llgo:link (*Node).HasProp C.xmlHasProp
func (recv_ *Node) HasProp(name *Char) AttrPtr {
	return nil
}

// llgo:link (*Node).HasNsProp C.xmlHasNsProp
func (recv_ *Node) HasNsProp(name *Char, nameSpace *Char) AttrPtr {
	return nil
}

// llgo:link (*Node).GetNsProp C.xmlGetNsProp
func (recv_ *Node) GetNsProp(name *Char, nameSpace *Char) *Char {
	return nil
}

// llgo:link (*Doc).StringGetNodeList C.xmlStringGetNodeList
func (recv_ *Doc) StringGetNodeList(value *Char) NodePtr {
	return nil
}

// llgo:link (*Doc).StringLenGetNodeList C.xmlStringLenGetNodeList
func (recv_ *Doc) StringLenGetNodeList(value *Char, len c.Int) NodePtr {
	return nil
}

//go:linkname NodeListGetString C.xmlNodeListGetString
func NodeListGetString(doc DocPtr, list *Node, inLine c.Int) *Char

// llgo:link (*Doc).NodeListGetRawString C.xmlNodeListGetRawString
func (recv_ *Doc) NodeListGetRawString(list *Node, inLine c.Int) *Char {
	return nil
}

//go:linkname NodeSetContent C.xmlNodeSetContent
func NodeSetContent(cur NodePtr, content *Char) c.Int

//go:linkname NodeSetContentLen C.xmlNodeSetContentLen
func NodeSetContentLen(cur NodePtr, content *Char, len c.Int) c.Int

//go:linkname NodeAddContent C.xmlNodeAddContent
func NodeAddContent(cur NodePtr, content *Char) c.Int

//go:linkname NodeAddContentLen C.xmlNodeAddContentLen
func NodeAddContentLen(cur NodePtr, content *Char, len c.Int) c.Int

// llgo:link (*Node).NodeGetContent C.xmlNodeGetContent
func (recv_ *Node) NodeGetContent() *Char {
	return nil
}

//go:linkname NodeBufGetContent C.xmlNodeBufGetContent
func NodeBufGetContent(buffer BufferPtr, cur *Node) c.Int

//go:linkname BufGetNodeContent C.xmlBufGetNodeContent
func BufGetNodeContent(buf BufPtr, cur *Node) c.Int

// llgo:link (*Node).NodeGetLang C.xmlNodeGetLang
func (recv_ *Node) NodeGetLang() *Char {
	return nil
}

// llgo:link (*Node).NodeGetSpacePreserve C.xmlNodeGetSpacePreserve
func (recv_ *Node) NodeGetSpacePreserve() c.Int {
	return 0
}

//go:linkname NodeSetLang C.xmlNodeSetLang
func NodeSetLang(cur NodePtr, lang *Char) c.Int

//go:linkname NodeSetSpacePreserve C.xmlNodeSetSpacePreserve
func NodeSetSpacePreserve(cur NodePtr, val c.Int) c.Int

// llgo:link (*Doc).NodeGetBaseSafe C.xmlNodeGetBaseSafe
func (recv_ *Doc) NodeGetBaseSafe(cur *Node, baseOut **Char) c.Int {
	return 0
}

// llgo:link (*Doc).NodeGetBase C.xmlNodeGetBase
func (recv_ *Doc) NodeGetBase(cur *Node) *Char {
	return nil
}

//go:linkname NodeSetBase C.xmlNodeSetBase
func NodeSetBase(cur NodePtr, uri *Char) c.Int

/*
 * Removing content.
 */
//go:linkname RemoveProp C.xmlRemoveProp
func RemoveProp(cur AttrPtr) c.Int

//go:linkname UnsetNsProp C.xmlUnsetNsProp
func UnsetNsProp(node NodePtr, ns NsPtr, name *Char) c.Int

//go:linkname UnsetProp C.xmlUnsetProp
func UnsetProp(node NodePtr, name *Char) c.Int

/*
 * Internal, don't use.
 */
//go:linkname BufferWriteCHAR C.xmlBufferWriteCHAR
func BufferWriteCHAR(buf BufferPtr, string *Char)

//go:linkname BufferWriteChar C.xmlBufferWriteChar
func BufferWriteChar(buf BufferPtr, string *int8)

//go:linkname BufferWriteQuotedString C.xmlBufferWriteQuotedString
func BufferWriteQuotedString(buf BufferPtr, string *Char)

//go:linkname AttrSerializeTxtContent C.xmlAttrSerializeTxtContent
func AttrSerializeTxtContent(buf BufferPtr, doc DocPtr, attr AttrPtr, string *Char)

/*
 * Namespace handling.
 */
//go:linkname ReconciliateNs C.xmlReconciliateNs
func ReconciliateNs(doc DocPtr, tree NodePtr) c.Int

/*
 * Saving.
 */
//go:linkname DocDumpFormatMemory C.xmlDocDumpFormatMemory
func DocDumpFormatMemory(cur DocPtr, mem **Char, size *c.Int, format c.Int)

//go:linkname DocDumpMemory C.xmlDocDumpMemory
func DocDumpMemory(cur DocPtr, mem **Char, size *c.Int)

//go:linkname DocDumpMemoryEnc C.xmlDocDumpMemoryEnc
func DocDumpMemoryEnc(out_doc DocPtr, doc_txt_ptr **Char, doc_txt_len *c.Int, txt_encoding *int8)

//go:linkname DocDumpFormatMemoryEnc C.xmlDocDumpFormatMemoryEnc
func DocDumpFormatMemoryEnc(out_doc DocPtr, doc_txt_ptr **Char, doc_txt_len *c.Int, txt_encoding *int8, format c.Int)

//go:linkname DocFormatDump C.xmlDocFormatDump
func DocFormatDump(f *c.FILE, cur DocPtr, format c.Int) c.Int

//go:linkname DocDump C.xmlDocDump
func DocDump(f *c.FILE, cur DocPtr) c.Int

//go:linkname ElemDump C.xmlElemDump
func ElemDump(f *c.FILE, doc DocPtr, cur NodePtr)

//go:linkname SaveFile C.xmlSaveFile
func SaveFile(filename *int8, cur DocPtr) c.Int

//go:linkname SaveFormatFile C.xmlSaveFormatFile
func SaveFormatFile(filename *int8, cur DocPtr, format c.Int) c.Int

//go:linkname BufNodeDump C.xmlBufNodeDump
func BufNodeDump(buf BufPtr, doc DocPtr, cur NodePtr, level c.Int, format c.Int) uintptr

//go:linkname NodeDump C.xmlNodeDump
func NodeDump(buf BufferPtr, doc DocPtr, cur NodePtr, level c.Int, format c.Int) c.Int

//go:linkname SaveFileTo C.xmlSaveFileTo
func SaveFileTo(buf OutputBufferPtr, cur DocPtr, encoding *int8) c.Int

//go:linkname SaveFormatFileTo C.xmlSaveFormatFileTo
func SaveFormatFileTo(buf OutputBufferPtr, cur DocPtr, encoding *int8, format c.Int) c.Int

//go:linkname NodeDumpOutput C.xmlNodeDumpOutput
func NodeDumpOutput(buf OutputBufferPtr, doc DocPtr, cur NodePtr, level c.Int, format c.Int, encoding *int8)

//go:linkname SaveFormatFileEnc C.xmlSaveFormatFileEnc
func SaveFormatFileEnc(filename *int8, cur DocPtr, encoding *int8, format c.Int) c.Int

//go:linkname SaveFileEnc C.xmlSaveFileEnc
func SaveFileEnc(filename *int8, cur DocPtr, encoding *int8) c.Int

/*
 * XHTML
 */
// llgo:link (*Char).IsXHTML C.xmlIsXHTML
func (recv_ *Char) IsXHTML(publicID *Char) c.Int {
	return 0
}

/*
 * Compression.
 */
// llgo:link (*Doc).GetDocCompressMode C.xmlGetDocCompressMode
func (recv_ *Doc) GetDocCompressMode() c.Int {
	return 0
}

//go:linkname SetDocCompressMode C.xmlSetDocCompressMode
func SetDocCompressMode(doc DocPtr, mode c.Int)

//go:linkname GetCompressMode C.xmlGetCompressMode
func GetCompressMode() c.Int

//go:linkname SetCompressMode C.xmlSetCompressMode
func SetCompressMode(mode c.Int)

/*
* DOM-wrapper helper functions.
 */
//go:linkname DOMWrapNewCtxt C.xmlDOMWrapNewCtxt
func DOMWrapNewCtxt() DOMWrapCtxtPtr

//go:linkname DOMWrapFreeCtxt C.xmlDOMWrapFreeCtxt
func DOMWrapFreeCtxt(ctxt DOMWrapCtxtPtr)

//go:linkname DOMWrapReconcileNamespaces C.xmlDOMWrapReconcileNamespaces
func DOMWrapReconcileNamespaces(ctxt DOMWrapCtxtPtr, elem NodePtr, options c.Int) c.Int

//go:linkname DOMWrapAdoptNode C.xmlDOMWrapAdoptNode
func DOMWrapAdoptNode(ctxt DOMWrapCtxtPtr, sourceDoc DocPtr, node NodePtr, destDoc DocPtr, destParent NodePtr, options c.Int) c.Int

//go:linkname DOMWrapRemoveNode C.xmlDOMWrapRemoveNode
func DOMWrapRemoveNode(ctxt DOMWrapCtxtPtr, doc DocPtr, node NodePtr, options c.Int) c.Int

//go:linkname DOMWrapCloneNode C.xmlDOMWrapCloneNode
func DOMWrapCloneNode(ctxt DOMWrapCtxtPtr, sourceDoc DocPtr, node NodePtr, clonedNode *NodePtr, destDoc DocPtr, destParent NodePtr, deep c.Int, options c.Int) c.Int

/*
 * 5 interfaces from DOM ElementTraversal, but different in entities
 * traversal.
 */
//go:linkname ChildElementCount C.xmlChildElementCount
func ChildElementCount(parent NodePtr) c.Ulong

//go:linkname NextElementSibling C.xmlNextElementSibling
func NextElementSibling(node NodePtr) NodePtr

//go:linkname FirstElementChild C.xmlFirstElementChild
func FirstElementChild(parent NodePtr) NodePtr

//go:linkname LastElementChild C.xmlLastElementChild
func LastElementChild(parent NodePtr) NodePtr

//go:linkname PreviousElementSibling C.xmlPreviousElementSibling
func PreviousElementSibling(node NodePtr) NodePtr

//go:linkname RegisterNodeDefault C.xmlRegisterNodeDefault
func RegisterNodeDefault(func_ RegisterNodeFunc) RegisterNodeFunc

//go:linkname DeregisterNodeDefault C.xmlDeregisterNodeDefault
func DeregisterNodeDefault(func_ DeregisterNodeFunc) DeregisterNodeFunc

//go:linkname ThrDefRegisterNodeDefault C.xmlThrDefRegisterNodeDefault
func ThrDefRegisterNodeDefault(func_ RegisterNodeFunc) RegisterNodeFunc

//go:linkname ThrDefDeregisterNodeDefault C.xmlThrDefDeregisterNodeDefault
func ThrDefDeregisterNodeDefault(func_ DeregisterNodeFunc) DeregisterNodeFunc

// llgo:link BufferAllocationScheme.ThrDefBufferAllocScheme C.xmlThrDefBufferAllocScheme
func (recv_ BufferAllocationScheme) ThrDefBufferAllocScheme() BufferAllocationScheme {
	return 0
}

//go:linkname ThrDefDefaultBufferSize C.xmlThrDefDefaultBufferSize
func ThrDefDefaultBufferSize(v c.Int) c.Int
