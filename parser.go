package libxml_2_0

import (
	"unsafe"

	"github.com/goplus/llgo/c"
)

const XMLDEFAULTVERSION string = "1.0"
const XMLDETECTIDS c.Int = 2
const XMLCOMPLETEATTRS c.Int = 4
const XMLSKIPIDS c.Int = 8
const XMLSAX2MAGIC c.Long = 3740122799

// llgo:type C
type XmlParserInputDeallocate func(*XmlChar)

type X_XmlParserNodeInfo struct {
	Node      *X_XmlNode
	BeginPos  c.Ulong
	BeginLine c.Ulong
	EndPos    c.Ulong
	EndLine   c.Ulong
}
type XmlParserNodeInfo X_XmlParserNodeInfo
type XmlParserNodeInfoPtr *XmlParserNodeInfo

type X_XmlParserNodeInfoSeq struct {
	Maximum c.Ulong
	Length  c.Ulong
	Buffer  *XmlParserNodeInfo
}
type XmlParserNodeInfoSeq X_XmlParserNodeInfoSeq
type XmlParserNodeInfoSeqPtr *XmlParserNodeInfoSeq
type XmlParserInputState c.Int

const (
	XmlParserInputStateXMLPARSEREOF            XmlParserInputState = -1
	XmlParserInputStateXMLPARSERSTART          XmlParserInputState = 0
	XmlParserInputStateXMLPARSERMISC           XmlParserInputState = 1
	XmlParserInputStateXMLPARSERPI             XmlParserInputState = 2
	XmlParserInputStateXMLPARSERDTD            XmlParserInputState = 3
	XmlParserInputStateXMLPARSERPROLOG         XmlParserInputState = 4
	XmlParserInputStateXMLPARSERCOMMENT        XmlParserInputState = 5
	XmlParserInputStateXMLPARSERSTARTTAG       XmlParserInputState = 6
	XmlParserInputStateXMLPARSERCONTENT        XmlParserInputState = 7
	XmlParserInputStateXMLPARSERCDATASECTION   XmlParserInputState = 8
	XmlParserInputStateXMLPARSERENDTAG         XmlParserInputState = 9
	XmlParserInputStateXMLPARSERENTITYDECL     XmlParserInputState = 10
	XmlParserInputStateXMLPARSERENTITYVALUE    XmlParserInputState = 11
	XmlParserInputStateXMLPARSERATTRIBUTEVALUE XmlParserInputState = 12
	XmlParserInputStateXMLPARSERSYSTEMLITERAL  XmlParserInputState = 13
	XmlParserInputStateXMLPARSEREPILOG         XmlParserInputState = 14
	XmlParserInputStateXMLPARSERIGNORE         XmlParserInputState = 15
	XmlParserInputStateXMLPARSERPUBLICLITERAL  XmlParserInputState = 16
	XmlParserInputStateXMLPARSERXMLDECL        XmlParserInputState = 17
)

type XmlParserMode c.Int

const (
	XmlParserModeXMLPARSEUNKNOWN XmlParserMode = 0
	XmlParserModeXMLPARSEDOM     XmlParserMode = 1
	XmlParserModeXMLPARSESAX     XmlParserMode = 2
	XmlParserModeXMLPARSEPUSHDOM XmlParserMode = 3
	XmlParserModeXMLPARSEPUSHSAX XmlParserMode = 4
	XmlParserModeXMLPARSEREADER  XmlParserMode = 5
)

type X_XmlStartTag struct {
	Unused [8]uint8
}
type XmlStartTag X_XmlStartTag

type X_XmlParserNsData struct {
	Unused [8]uint8
}
type XmlParserNsData X_XmlParserNsData

type X_XmlAttrHashBucket struct {
	Unused [8]uint8
}
type XmlAttrHashBucket X_XmlAttrHashBucket

// llgo:type C
type ResolveEntitySAXFunc func(unsafe.Pointer, *XmlChar, *XmlChar) XmlParserInputPtr

// llgo:type C
type InternalSubsetSAXFunc func(unsafe.Pointer, *XmlChar, *XmlChar, *XmlChar)

// llgo:type C
type ExternalSubsetSAXFunc func(unsafe.Pointer, *XmlChar, *XmlChar, *XmlChar)

// llgo:type C
type GetEntitySAXFunc func(unsafe.Pointer, *XmlChar) XmlEntityPtr

// llgo:type C
type GetParameterEntitySAXFunc func(unsafe.Pointer, *XmlChar) XmlEntityPtr

// llgo:type C
type EntityDeclSAXFunc func(unsafe.Pointer, *XmlChar, c.Int, *XmlChar, *XmlChar, *XmlChar)

// llgo:type C
type NotationDeclSAXFunc func(unsafe.Pointer, *XmlChar, *XmlChar, *XmlChar)

// llgo:type C
type AttributeDeclSAXFunc func(unsafe.Pointer, *XmlChar, *XmlChar, c.Int, c.Int, *XmlChar, XmlEnumerationPtr)

// llgo:type C
type ElementDeclSAXFunc func(unsafe.Pointer, *XmlChar, c.Int, XmlElementContentPtr)

// llgo:type C
type UnparsedEntityDeclSAXFunc func(unsafe.Pointer, *XmlChar, *XmlChar, *XmlChar, *XmlChar)

// llgo:type C
type SetDocumentLocatorSAXFunc func(unsafe.Pointer, XmlSAXLocatorPtr)

// llgo:type C
type StartDocumentSAXFunc func(unsafe.Pointer)

// llgo:type C
type EndDocumentSAXFunc func(unsafe.Pointer)

// llgo:type C
type StartElementSAXFunc func(unsafe.Pointer, *XmlChar, **XmlChar)

// llgo:type C
type EndElementSAXFunc func(unsafe.Pointer, *XmlChar)

// llgo:type C
type AttributeSAXFunc func(unsafe.Pointer, *XmlChar, *XmlChar)

// llgo:type C
type ReferenceSAXFunc func(unsafe.Pointer, *XmlChar)

// llgo:type C
type CharactersSAXFunc func(unsafe.Pointer, *XmlChar, c.Int)

// llgo:type C
type IgnorableWhitespaceSAXFunc func(unsafe.Pointer, *XmlChar, c.Int)

// llgo:type C
type ProcessingInstructionSAXFunc func(unsafe.Pointer, *XmlChar, *XmlChar)

// llgo:type C
type CommentSAXFunc func(unsafe.Pointer, *XmlChar)

// llgo:type C
type CdataBlockSAXFunc func(unsafe.Pointer, *XmlChar, c.Int)

// llgo:type C
type WarningSAXFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type ErrorSAXFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type FatalErrorSAXFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})

// llgo:type C
type IsStandaloneSAXFunc func(unsafe.Pointer) c.Int

// llgo:type C
type HasInternalSubsetSAXFunc func(unsafe.Pointer) c.Int

// llgo:type C
type HasExternalSubsetSAXFunc func(unsafe.Pointer) c.Int

// llgo:type C
type StartElementNsSAX2Func func(unsafe.Pointer, *XmlChar, *XmlChar, *XmlChar, c.Int, **XmlChar, c.Int, c.Int, **XmlChar)

// llgo:type C
type EndElementNsSAX2Func func(unsafe.Pointer, *XmlChar, *XmlChar, *XmlChar)

type X_XmlSAXHandlerV1 struct {
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
}
type XmlSAXHandlerV1 X_XmlSAXHandlerV1
type XmlSAXHandlerV1Ptr *XmlSAXHandlerV1

// llgo:type C
type XmlExternalEntityLoader func(*int8, *int8, XmlParserCtxtPtr) XmlParserInputPtr

//go:linkname X__XmlParserVersion C.__xmlParserVersion
func X__XmlParserVersion() **int8

//go:linkname X__OldXMLWDcompatibility C.__oldXMLWDcompatibility
func X__OldXMLWDcompatibility() *c.Int

//go:linkname X__XmlParserDebugEntities C.__xmlParserDebugEntities
func X__XmlParserDebugEntities() *c.Int

//go:linkname X__XmlDefaultSAXLocator C.__xmlDefaultSAXLocator
func X__XmlDefaultSAXLocator() *XmlSAXLocator

//go:linkname X__XmlDefaultSAXHandler C.__xmlDefaultSAXHandler
func X__XmlDefaultSAXHandler() *XmlSAXHandlerV1

//go:linkname X__XmlDoValidityCheckingDefaultValue C.__xmlDoValidityCheckingDefaultValue
func X__XmlDoValidityCheckingDefaultValue() *c.Int

//go:linkname X__XmlGetWarningsDefaultValue C.__xmlGetWarningsDefaultValue
func X__XmlGetWarningsDefaultValue() *c.Int

//go:linkname X__XmlKeepBlanksDefaultValue C.__xmlKeepBlanksDefaultValue
func X__XmlKeepBlanksDefaultValue() *c.Int

//go:linkname X__XmlLineNumbersDefaultValue C.__xmlLineNumbersDefaultValue
func X__XmlLineNumbersDefaultValue() *c.Int

//go:linkname X__XmlLoadExtDtdDefaultValue C.__xmlLoadExtDtdDefaultValue
func X__XmlLoadExtDtdDefaultValue() *c.Int

//go:linkname X__XmlPedanticParserDefaultValue C.__xmlPedanticParserDefaultValue
func X__XmlPedanticParserDefaultValue() *c.Int

//go:linkname X__XmlSubstituteEntitiesDefaultValue C.__xmlSubstituteEntitiesDefaultValue
func X__XmlSubstituteEntitiesDefaultValue() *c.Int

//go:linkname X__XmlIndentTreeOutput C.__xmlIndentTreeOutput
func X__XmlIndentTreeOutput() *c.Int

//go:linkname X__XmlTreeIndentString C.__xmlTreeIndentString
func X__XmlTreeIndentString() **int8

//go:linkname X__XmlSaveNoEmptyTags C.__xmlSaveNoEmptyTags
func X__XmlSaveNoEmptyTags() *c.Int

/** DOC_ENABLE */
//go:linkname XmlInitParser C.xmlInitParser
func XmlInitParser()

//go:linkname XmlCleanupParser C.xmlCleanupParser
func XmlCleanupParser()

//go:linkname XmlInitGlobals C.xmlInitGlobals
func XmlInitGlobals()

//go:linkname XmlCleanupGlobals C.xmlCleanupGlobals
func XmlCleanupGlobals()

//go:linkname XmlParserInputRead C.xmlParserInputRead
func XmlParserInputRead(in XmlParserInputPtr, len c.Int) c.Int

//go:linkname XmlParserInputGrow C.xmlParserInputGrow
func XmlParserInputGrow(in XmlParserInputPtr, len c.Int) c.Int

// llgo:link (*XmlChar).XmlParseDoc C.xmlParseDoc
func (recv_ *XmlChar) XmlParseDoc() XmlDocPtr {
	return nil
}

//go:linkname XmlParseFile C.xmlParseFile
func XmlParseFile(filename *int8) XmlDocPtr

//go:linkname XmlParseMemory C.xmlParseMemory
func XmlParseMemory(buffer *int8, size c.Int) XmlDocPtr

//go:linkname XmlSubstituteEntitiesDefault C.xmlSubstituteEntitiesDefault
func XmlSubstituteEntitiesDefault(val c.Int) c.Int

//go:linkname XmlThrDefSubstituteEntitiesDefaultValue C.xmlThrDefSubstituteEntitiesDefaultValue
func XmlThrDefSubstituteEntitiesDefaultValue(v c.Int) c.Int

//go:linkname XmlKeepBlanksDefault C.xmlKeepBlanksDefault
func XmlKeepBlanksDefault(val c.Int) c.Int

//go:linkname XmlThrDefKeepBlanksDefaultValue C.xmlThrDefKeepBlanksDefaultValue
func XmlThrDefKeepBlanksDefaultValue(v c.Int) c.Int

//go:linkname XmlStopParser C.xmlStopParser
func XmlStopParser(ctxt XmlParserCtxtPtr)

//go:linkname XmlPedanticParserDefault C.xmlPedanticParserDefault
func XmlPedanticParserDefault(val c.Int) c.Int

//go:linkname XmlThrDefPedanticParserDefaultValue C.xmlThrDefPedanticParserDefaultValue
func XmlThrDefPedanticParserDefaultValue(v c.Int) c.Int

//go:linkname XmlLineNumbersDefault C.xmlLineNumbersDefault
func XmlLineNumbersDefault(val c.Int) c.Int

//go:linkname XmlThrDefLineNumbersDefaultValue C.xmlThrDefLineNumbersDefaultValue
func XmlThrDefLineNumbersDefaultValue(v c.Int) c.Int

//go:linkname XmlThrDefDoValidityCheckingDefaultValue C.xmlThrDefDoValidityCheckingDefaultValue
func XmlThrDefDoValidityCheckingDefaultValue(v c.Int) c.Int

//go:linkname XmlThrDefGetWarningsDefaultValue C.xmlThrDefGetWarningsDefaultValue
func XmlThrDefGetWarningsDefaultValue(v c.Int) c.Int

//go:linkname XmlThrDefLoadExtDtdDefaultValue C.xmlThrDefLoadExtDtdDefaultValue
func XmlThrDefLoadExtDtdDefaultValue(v c.Int) c.Int

//go:linkname XmlThrDefParserDebugEntities C.xmlThrDefParserDebugEntities
func XmlThrDefParserDebugEntities(v c.Int) c.Int

// llgo:link (*XmlChar).XmlRecoverDoc C.xmlRecoverDoc
func (recv_ *XmlChar) XmlRecoverDoc() XmlDocPtr {
	return nil
}

//go:linkname XmlRecoverMemory C.xmlRecoverMemory
func XmlRecoverMemory(buffer *int8, size c.Int) XmlDocPtr

//go:linkname XmlRecoverFile C.xmlRecoverFile
func XmlRecoverFile(filename *int8) XmlDocPtr

//go:linkname XmlParseDocument C.xmlParseDocument
func XmlParseDocument(ctxt XmlParserCtxtPtr) c.Int

//go:linkname XmlParseExtParsedEnt C.xmlParseExtParsedEnt
func XmlParseExtParsedEnt(ctxt XmlParserCtxtPtr) c.Int

//go:linkname XmlSAXUserParseFile C.xmlSAXUserParseFile
func XmlSAXUserParseFile(sax XmlSAXHandlerPtr, user_data unsafe.Pointer, filename *int8) c.Int

//go:linkname XmlSAXUserParseMemory C.xmlSAXUserParseMemory
func XmlSAXUserParseMemory(sax XmlSAXHandlerPtr, user_data unsafe.Pointer, buffer *int8, size c.Int) c.Int

//go:linkname XmlSAXParseDoc C.xmlSAXParseDoc
func XmlSAXParseDoc(sax XmlSAXHandlerPtr, cur *XmlChar, recovery c.Int) XmlDocPtr

//go:linkname XmlSAXParseMemory C.xmlSAXParseMemory
func XmlSAXParseMemory(sax XmlSAXHandlerPtr, buffer *int8, size c.Int, recovery c.Int) XmlDocPtr

//go:linkname XmlSAXParseMemoryWithData C.xmlSAXParseMemoryWithData
func XmlSAXParseMemoryWithData(sax XmlSAXHandlerPtr, buffer *int8, size c.Int, recovery c.Int, data unsafe.Pointer) XmlDocPtr

//go:linkname XmlSAXParseFile C.xmlSAXParseFile
func XmlSAXParseFile(sax XmlSAXHandlerPtr, filename *int8, recovery c.Int) XmlDocPtr

//go:linkname XmlSAXParseFileWithData C.xmlSAXParseFileWithData
func XmlSAXParseFileWithData(sax XmlSAXHandlerPtr, filename *int8, recovery c.Int, data unsafe.Pointer) XmlDocPtr

//go:linkname XmlSAXParseEntity C.xmlSAXParseEntity
func XmlSAXParseEntity(sax XmlSAXHandlerPtr, filename *int8) XmlDocPtr

//go:linkname XmlParseEntity C.xmlParseEntity
func XmlParseEntity(filename *int8) XmlDocPtr

//go:linkname XmlSAXParseDTD C.xmlSAXParseDTD
func XmlSAXParseDTD(sax XmlSAXHandlerPtr, ExternalID *XmlChar, SystemID *XmlChar) XmlDtdPtr

// llgo:link (*XmlChar).XmlParseDTD C.xmlParseDTD
func (recv_ *XmlChar) XmlParseDTD(SystemID *XmlChar) XmlDtdPtr {
	return nil
}

//go:linkname XmlIOParseDTD C.xmlIOParseDTD
func XmlIOParseDTD(sax XmlSAXHandlerPtr, input XmlParserInputBufferPtr, enc XmlCharEncoding) XmlDtdPtr

//go:linkname XmlParseBalancedChunkMemory C.xmlParseBalancedChunkMemory
func XmlParseBalancedChunkMemory(doc XmlDocPtr, sax XmlSAXHandlerPtr, user_data unsafe.Pointer, depth c.Int, string *XmlChar, lst *XmlNodePtr) c.Int

//go:linkname XmlParseInNodeContext C.xmlParseInNodeContext
func XmlParseInNodeContext(node XmlNodePtr, data *int8, datalen c.Int, options c.Int, lst *XmlNodePtr) XmlParserErrors

//go:linkname XmlParseBalancedChunkMemoryRecover C.xmlParseBalancedChunkMemoryRecover
func XmlParseBalancedChunkMemoryRecover(doc XmlDocPtr, sax XmlSAXHandlerPtr, user_data unsafe.Pointer, depth c.Int, string *XmlChar, lst *XmlNodePtr, recover c.Int) c.Int

//go:linkname XmlParseExternalEntity C.xmlParseExternalEntity
func XmlParseExternalEntity(doc XmlDocPtr, sax XmlSAXHandlerPtr, user_data unsafe.Pointer, depth c.Int, URL *XmlChar, ID *XmlChar, lst *XmlNodePtr) c.Int

//go:linkname XmlParseCtxtExternalEntity C.xmlParseCtxtExternalEntity
func XmlParseCtxtExternalEntity(ctx XmlParserCtxtPtr, URL *XmlChar, ID *XmlChar, lst *XmlNodePtr) c.Int

//go:linkname XmlNewParserCtxt C.xmlNewParserCtxt
func XmlNewParserCtxt() XmlParserCtxtPtr

// llgo:link (*XmlSAXHandler).XmlNewSAXParserCtxt C.xmlNewSAXParserCtxt
func (recv_ *XmlSAXHandler) XmlNewSAXParserCtxt(userData unsafe.Pointer) XmlParserCtxtPtr {
	return nil
}

//go:linkname XmlInitParserCtxt C.xmlInitParserCtxt
func XmlInitParserCtxt(ctxt XmlParserCtxtPtr) c.Int

//go:linkname XmlClearParserCtxt C.xmlClearParserCtxt
func XmlClearParserCtxt(ctxt XmlParserCtxtPtr)

//go:linkname XmlFreeParserCtxt C.xmlFreeParserCtxt
func XmlFreeParserCtxt(ctxt XmlParserCtxtPtr)

//go:linkname XmlSetupParserForBuffer C.xmlSetupParserForBuffer
func XmlSetupParserForBuffer(ctxt XmlParserCtxtPtr, buffer *XmlChar, filename *int8)

// llgo:link (*XmlChar).XmlCreateDocParserCtxt C.xmlCreateDocParserCtxt
func (recv_ *XmlChar) XmlCreateDocParserCtxt() XmlParserCtxtPtr {
	return nil
}

//go:linkname XmlGetFeaturesList C.xmlGetFeaturesList
func XmlGetFeaturesList(len *c.Int, result **int8) c.Int

//go:linkname XmlGetFeature C.xmlGetFeature
func XmlGetFeature(ctxt XmlParserCtxtPtr, name *int8, result unsafe.Pointer) c.Int

//go:linkname XmlSetFeature C.xmlSetFeature
func XmlSetFeature(ctxt XmlParserCtxtPtr, name *int8, value unsafe.Pointer) c.Int

//go:linkname XmlCreatePushParserCtxt C.xmlCreatePushParserCtxt
func XmlCreatePushParserCtxt(sax XmlSAXHandlerPtr, user_data unsafe.Pointer, chunk *int8, size c.Int, filename *int8) XmlParserCtxtPtr

//go:linkname XmlParseChunk C.xmlParseChunk
func XmlParseChunk(ctxt XmlParserCtxtPtr, chunk *int8, size c.Int, terminate c.Int) c.Int

//go:linkname XmlCreateIOParserCtxt C.xmlCreateIOParserCtxt
func XmlCreateIOParserCtxt(sax XmlSAXHandlerPtr, user_data unsafe.Pointer, ioread XmlInputReadCallback, ioclose XmlInputCloseCallback, ioctx unsafe.Pointer, enc XmlCharEncoding) XmlParserCtxtPtr

//go:linkname XmlNewIOInputStream C.xmlNewIOInputStream
func XmlNewIOInputStream(ctxt XmlParserCtxtPtr, input XmlParserInputBufferPtr, enc XmlCharEncoding) XmlParserInputPtr

//go:linkname XmlParserFindNodeInfo C.xmlParserFindNodeInfo
func XmlParserFindNodeInfo(ctxt XmlParserCtxtPtr, node XmlNodePtr) *XmlParserNodeInfo

//go:linkname XmlInitNodeInfoSeq C.xmlInitNodeInfoSeq
func XmlInitNodeInfoSeq(seq XmlParserNodeInfoSeqPtr)

//go:linkname XmlClearNodeInfoSeq C.xmlClearNodeInfoSeq
func XmlClearNodeInfoSeq(seq XmlParserNodeInfoSeqPtr)

//go:linkname XmlParserFindNodeInfoIndex C.xmlParserFindNodeInfoIndex
func XmlParserFindNodeInfoIndex(seq XmlParserNodeInfoSeqPtr, node XmlNodePtr) c.Ulong

//go:linkname XmlParserAddNodeInfo C.xmlParserAddNodeInfo
func XmlParserAddNodeInfo(ctxt XmlParserCtxtPtr, info XmlParserNodeInfoPtr)

//go:linkname XmlSetExternalEntityLoader C.xmlSetExternalEntityLoader
func XmlSetExternalEntityLoader(f XmlExternalEntityLoader)

//go:linkname XmlGetExternalEntityLoader C.xmlGetExternalEntityLoader
func XmlGetExternalEntityLoader() XmlExternalEntityLoader

//go:linkname XmlLoadExternalEntity C.xmlLoadExternalEntity
func XmlLoadExternalEntity(URL *int8, ID *int8, ctxt XmlParserCtxtPtr) XmlParserInputPtr

//go:linkname XmlByteConsumed C.xmlByteConsumed
func XmlByteConsumed(ctxt XmlParserCtxtPtr) c.Long

type XmlParserOption c.Int

const (
	XmlParserOptionXMLPARSERECOVER    XmlParserOption = 1
	XmlParserOptionXMLPARSENOENT      XmlParserOption = 2
	XmlParserOptionXMLPARSEDTDLOAD    XmlParserOption = 4
	XmlParserOptionXMLPARSEDTDATTR    XmlParserOption = 8
	XmlParserOptionXMLPARSEDTDVALID   XmlParserOption = 16
	XmlParserOptionXMLPARSENOERROR    XmlParserOption = 32
	XmlParserOptionXMLPARSENOWARNING  XmlParserOption = 64
	XmlParserOptionXMLPARSEPEDANTIC   XmlParserOption = 128
	XmlParserOptionXMLPARSENOBLANKS   XmlParserOption = 256
	XmlParserOptionXMLPARSESAX1       XmlParserOption = 512
	XmlParserOptionXMLPARSEXINCLUDE   XmlParserOption = 1024
	XmlParserOptionXMLPARSENONET      XmlParserOption = 2048
	XmlParserOptionXMLPARSENODICT     XmlParserOption = 4096
	XmlParserOptionXMLPARSENSCLEAN    XmlParserOption = 8192
	XmlParserOptionXMLPARSENOCDATA    XmlParserOption = 16384
	XmlParserOptionXMLPARSENOXINCNODE XmlParserOption = 32768
	XmlParserOptionXMLPARSECOMPACT    XmlParserOption = 65536
	XmlParserOptionXMLPARSEOLD10      XmlParserOption = 131072
	XmlParserOptionXMLPARSENOBASEFIX  XmlParserOption = 262144
	XmlParserOptionXMLPARSEHUGE       XmlParserOption = 524288
	XmlParserOptionXMLPARSEOLDSAX     XmlParserOption = 1048576
	XmlParserOptionXMLPARSEIGNOREENC  XmlParserOption = 2097152
	XmlParserOptionXMLPARSEBIGLINES   XmlParserOption = 4194304
	XmlParserOptionXMLPARSENOXXE      XmlParserOption = 8388608
)

//go:linkname XmlCtxtReset C.xmlCtxtReset
func XmlCtxtReset(ctxt XmlParserCtxtPtr)

//go:linkname XmlCtxtResetPush C.xmlCtxtResetPush
func XmlCtxtResetPush(ctxt XmlParserCtxtPtr, chunk *int8, size c.Int, filename *int8, encoding *int8) c.Int

//go:linkname XmlCtxtSetOptions C.xmlCtxtSetOptions
func XmlCtxtSetOptions(ctxt XmlParserCtxtPtr, options c.Int) c.Int

//go:linkname XmlCtxtUseOptions C.xmlCtxtUseOptions
func XmlCtxtUseOptions(ctxt XmlParserCtxtPtr, options c.Int) c.Int

//go:linkname XmlCtxtSetErrorHandler C.xmlCtxtSetErrorHandler
func XmlCtxtSetErrorHandler(ctxt XmlParserCtxtPtr, handler XmlStructuredErrorFunc, data unsafe.Pointer)

//go:linkname XmlCtxtSetMaxAmplification C.xmlCtxtSetMaxAmplification
func XmlCtxtSetMaxAmplification(ctxt XmlParserCtxtPtr, maxAmpl c.Uint)

// llgo:link (*XmlChar).XmlReadDoc C.xmlReadDoc
func (recv_ *XmlChar) XmlReadDoc(URL *int8, encoding *int8, options c.Int) XmlDocPtr {
	return nil
}

//go:linkname XmlReadFile C.xmlReadFile
func XmlReadFile(URL *int8, encoding *int8, options c.Int) XmlDocPtr

//go:linkname XmlReadMemory C.xmlReadMemory
func XmlReadMemory(buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) XmlDocPtr

//go:linkname XmlReadFd C.xmlReadFd
func XmlReadFd(fd c.Int, URL *int8, encoding *int8, options c.Int) XmlDocPtr

//go:linkname XmlReadIO C.xmlReadIO
func XmlReadIO(ioread XmlInputReadCallback, ioclose XmlInputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) XmlDocPtr

//go:linkname XmlCtxtParseDocument C.xmlCtxtParseDocument
func XmlCtxtParseDocument(ctxt XmlParserCtxtPtr, input XmlParserInputPtr) XmlDocPtr

//go:linkname XmlCtxtReadDoc C.xmlCtxtReadDoc
func XmlCtxtReadDoc(ctxt XmlParserCtxtPtr, cur *XmlChar, URL *int8, encoding *int8, options c.Int) XmlDocPtr

//go:linkname XmlCtxtReadFile C.xmlCtxtReadFile
func XmlCtxtReadFile(ctxt XmlParserCtxtPtr, filename *int8, encoding *int8, options c.Int) XmlDocPtr

//go:linkname XmlCtxtReadMemory C.xmlCtxtReadMemory
func XmlCtxtReadMemory(ctxt XmlParserCtxtPtr, buffer *int8, size c.Int, URL *int8, encoding *int8, options c.Int) XmlDocPtr

//go:linkname XmlCtxtReadFd C.xmlCtxtReadFd
func XmlCtxtReadFd(ctxt XmlParserCtxtPtr, fd c.Int, URL *int8, encoding *int8, options c.Int) XmlDocPtr

//go:linkname XmlCtxtReadIO C.xmlCtxtReadIO
func XmlCtxtReadIO(ctxt XmlParserCtxtPtr, ioread XmlInputReadCallback, ioclose XmlInputCloseCallback, ioctx unsafe.Pointer, URL *int8, encoding *int8, options c.Int) XmlDocPtr

type XmlFeature c.Int

const (
	XmlFeatureXMLWITHTHREAD     XmlFeature = 1
	XmlFeatureXMLWITHTREE       XmlFeature = 2
	XmlFeatureXMLWITHOUTPUT     XmlFeature = 3
	XmlFeatureXMLWITHPUSH       XmlFeature = 4
	XmlFeatureXMLWITHREADER     XmlFeature = 5
	XmlFeatureXMLWITHPATTERN    XmlFeature = 6
	XmlFeatureXMLWITHWRITER     XmlFeature = 7
	XmlFeatureXMLWITHSAX1       XmlFeature = 8
	XmlFeatureXMLWITHFTP        XmlFeature = 9
	XmlFeatureXMLWITHHTTP       XmlFeature = 10
	XmlFeatureXMLWITHVALID      XmlFeature = 11
	XmlFeatureXMLWITHHTML       XmlFeature = 12
	XmlFeatureXMLWITHLEGACY     XmlFeature = 13
	XmlFeatureXMLWITHC14N       XmlFeature = 14
	XmlFeatureXMLWITHCATALOG    XmlFeature = 15
	XmlFeatureXMLWITHXPATH      XmlFeature = 16
	XmlFeatureXMLWITHXPTR       XmlFeature = 17
	XmlFeatureXMLWITHXINCLUDE   XmlFeature = 18
	XmlFeatureXMLWITHICONV      XmlFeature = 19
	XmlFeatureXMLWITHISO8859X   XmlFeature = 20
	XmlFeatureXMLWITHUNICODE    XmlFeature = 21
	XmlFeatureXMLWITHREGEXP     XmlFeature = 22
	XmlFeatureXMLWITHAUTOMATA   XmlFeature = 23
	XmlFeatureXMLWITHEXPR       XmlFeature = 24
	XmlFeatureXMLWITHSCHEMAS    XmlFeature = 25
	XmlFeatureXMLWITHSCHEMATRON XmlFeature = 26
	XmlFeatureXMLWITHMODULES    XmlFeature = 27
	XmlFeatureXMLWITHDEBUG      XmlFeature = 28
	XmlFeatureXMLWITHDEBUGMEM   XmlFeature = 29
	XmlFeatureXMLWITHDEBUGRUN   XmlFeature = 30
	XmlFeatureXMLWITHZLIB       XmlFeature = 31
	XmlFeatureXMLWITHICU        XmlFeature = 32
	XmlFeatureXMLWITHLZMA       XmlFeature = 33
	XmlFeatureXMLWITHNONE       XmlFeature = 99999
)

// llgo:link XmlFeature.XmlHasFeature C.xmlHasFeature
func (recv_ XmlFeature) XmlHasFeature() c.Int {
	return 0
}
