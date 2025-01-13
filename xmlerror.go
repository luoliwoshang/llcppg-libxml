package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type XmlErrorLevel c.Int

const (
	XmlErrorLevelXMLERRNONE    XmlErrorLevel = 0
	XmlErrorLevelXMLERRWARNING XmlErrorLevel = 1
	XmlErrorLevelXMLERRERROR   XmlErrorLevel = 2
	XmlErrorLevelXMLERRFATAL   XmlErrorLevel = 3
)

type XmlErrorDomain c.Int

const (
	XmlErrorDomainXMLFROMNONE        XmlErrorDomain = 0
	XmlErrorDomainXMLFROMPARSER      XmlErrorDomain = 1
	XmlErrorDomainXMLFROMTREE        XmlErrorDomain = 2
	XmlErrorDomainXMLFROMNAMESPACE   XmlErrorDomain = 3
	XmlErrorDomainXMLFROMDTD         XmlErrorDomain = 4
	XmlErrorDomainXMLFROMHTML        XmlErrorDomain = 5
	XmlErrorDomainXMLFROMMEMORY      XmlErrorDomain = 6
	XmlErrorDomainXMLFROMOUTPUT      XmlErrorDomain = 7
	XmlErrorDomainXMLFROMIO          XmlErrorDomain = 8
	XmlErrorDomainXMLFROMFTP         XmlErrorDomain = 9
	XmlErrorDomainXMLFROMHTTP        XmlErrorDomain = 10
	XmlErrorDomainXMLFROMXINCLUDE    XmlErrorDomain = 11
	XmlErrorDomainXMLFROMXPATH       XmlErrorDomain = 12
	XmlErrorDomainXMLFROMXPOINTER    XmlErrorDomain = 13
	XmlErrorDomainXMLFROMREGEXP      XmlErrorDomain = 14
	XmlErrorDomainXMLFROMDATATYPE    XmlErrorDomain = 15
	XmlErrorDomainXMLFROMSCHEMASP    XmlErrorDomain = 16
	XmlErrorDomainXMLFROMSCHEMASV    XmlErrorDomain = 17
	XmlErrorDomainXMLFROMRELAXNGP    XmlErrorDomain = 18
	XmlErrorDomainXMLFROMRELAXNGV    XmlErrorDomain = 19
	XmlErrorDomainXMLFROMCATALOG     XmlErrorDomain = 20
	XmlErrorDomainXMLFROMC14N        XmlErrorDomain = 21
	XmlErrorDomainXMLFROMXSLT        XmlErrorDomain = 22
	XmlErrorDomainXMLFROMVALID       XmlErrorDomain = 23
	XmlErrorDomainXMLFROMCHECK       XmlErrorDomain = 24
	XmlErrorDomainXMLFROMWRITER      XmlErrorDomain = 25
	XmlErrorDomainXMLFROMMODULE      XmlErrorDomain = 26
	XmlErrorDomainXMLFROMI18N        XmlErrorDomain = 27
	XmlErrorDomainXMLFROMSCHEMATRONV XmlErrorDomain = 28
	XmlErrorDomainXMLFROMBUFFER      XmlErrorDomain = 29
	XmlErrorDomainXMLFROMURI         XmlErrorDomain = 30
)

type X_XmlError struct {
	Domain  c.Int
	Code    c.Int
	Message *int8
	Level   XmlErrorLevel
	File    *int8
	Line    c.Int
	Str1    *int8
	Str2    *int8
	Str3    *int8
	Int1    c.Int
	Int2    c.Int
	Ctxt    unsafe.Pointer
	Node    unsafe.Pointer
}
type XmlError X_XmlError
type XmlErrorPtr *XmlError
type XmlParserErrors c.Int

const (
	XmlParserErrorsXMLERROK                                   XmlParserErrors = 0
	XmlParserErrorsXMLERRINTERNALERROR                        XmlParserErrors = 1
	XmlParserErrorsXMLERRNOMEMORY                             XmlParserErrors = 2
	XmlParserErrorsXMLERRDOCUMENTSTART                        XmlParserErrors = 3
	XmlParserErrorsXMLERRDOCUMENTEMPTY                        XmlParserErrors = 4
	XmlParserErrorsXMLERRDOCUMENTEND                          XmlParserErrors = 5
	XmlParserErrorsXMLERRINVALIDHEXCHARREF                    XmlParserErrors = 6
	XmlParserErrorsXMLERRINVALIDDECCHARREF                    XmlParserErrors = 7
	XmlParserErrorsXMLERRINVALIDCHARREF                       XmlParserErrors = 8
	XmlParserErrorsXMLERRINVALIDCHAR                          XmlParserErrors = 9
	XmlParserErrorsXMLERRCHARREFATEOF                         XmlParserErrors = 10
	XmlParserErrorsXMLERRCHARREFINPROLOG                      XmlParserErrors = 11
	XmlParserErrorsXMLERRCHARREFINEPILOG                      XmlParserErrors = 12
	XmlParserErrorsXMLERRCHARREFINDTD                         XmlParserErrors = 13
	XmlParserErrorsXMLERRENTITYREFATEOF                       XmlParserErrors = 14
	XmlParserErrorsXMLERRENTITYREFINPROLOG                    XmlParserErrors = 15
	XmlParserErrorsXMLERRENTITYREFINEPILOG                    XmlParserErrors = 16
	XmlParserErrorsXMLERRENTITYREFINDTD                       XmlParserErrors = 17
	XmlParserErrorsXMLERRPEREFATEOF                           XmlParserErrors = 18
	XmlParserErrorsXMLERRPEREFINPROLOG                        XmlParserErrors = 19
	XmlParserErrorsXMLERRPEREFINEPILOG                        XmlParserErrors = 20
	XmlParserErrorsXMLERRPEREFININTSUBSET                     XmlParserErrors = 21
	XmlParserErrorsXMLERRENTITYREFNONAME                      XmlParserErrors = 22
	XmlParserErrorsXMLERRENTITYREFSEMICOLMISSING              XmlParserErrors = 23
	XmlParserErrorsXMLERRPEREFNONAME                          XmlParserErrors = 24
	XmlParserErrorsXMLERRPEREFSEMICOLMISSING                  XmlParserErrors = 25
	XmlParserErrorsXMLERRUNDECLAREDENTITY                     XmlParserErrors = 26
	XmlParserErrorsXMLWARUNDECLAREDENTITY                     XmlParserErrors = 27
	XmlParserErrorsXMLERRUNPARSEDENTITY                       XmlParserErrors = 28
	XmlParserErrorsXMLERRENTITYISEXTERNAL                     XmlParserErrors = 29
	XmlParserErrorsXMLERRENTITYISPARAMETER                    XmlParserErrors = 30
	XmlParserErrorsXMLERRUNKNOWNENCODING                      XmlParserErrors = 31
	XmlParserErrorsXMLERRUNSUPPORTEDENCODING                  XmlParserErrors = 32
	XmlParserErrorsXMLERRSTRINGNOTSTARTED                     XmlParserErrors = 33
	XmlParserErrorsXMLERRSTRINGNOTCLOSED                      XmlParserErrors = 34
	XmlParserErrorsXMLERRNSDECLERROR                          XmlParserErrors = 35
	XmlParserErrorsXMLERRENTITYNOTSTARTED                     XmlParserErrors = 36
	XmlParserErrorsXMLERRENTITYNOTFINISHED                    XmlParserErrors = 37
	XmlParserErrorsXMLERRLTINATTRIBUTE                        XmlParserErrors = 38
	XmlParserErrorsXMLERRATTRIBUTENOTSTARTED                  XmlParserErrors = 39
	XmlParserErrorsXMLERRATTRIBUTENOTFINISHED                 XmlParserErrors = 40
	XmlParserErrorsXMLERRATTRIBUTEWITHOUTVALUE                XmlParserErrors = 41
	XmlParserErrorsXMLERRATTRIBUTEREDEFINED                   XmlParserErrors = 42
	XmlParserErrorsXMLERRLITERALNOTSTARTED                    XmlParserErrors = 43
	XmlParserErrorsXMLERRLITERALNOTFINISHED                   XmlParserErrors = 44
	XmlParserErrorsXMLERRCOMMENTNOTFINISHED                   XmlParserErrors = 45
	XmlParserErrorsXMLERRPINOTSTARTED                         XmlParserErrors = 46
	XmlParserErrorsXMLERRPINOTFINISHED                        XmlParserErrors = 47
	XmlParserErrorsXMLERRNOTATIONNOTSTARTED                   XmlParserErrors = 48
	XmlParserErrorsXMLERRNOTATIONNOTFINISHED                  XmlParserErrors = 49
	XmlParserErrorsXMLERRATTLISTNOTSTARTED                    XmlParserErrors = 50
	XmlParserErrorsXMLERRATTLISTNOTFINISHED                   XmlParserErrors = 51
	XmlParserErrorsXMLERRMIXEDNOTSTARTED                      XmlParserErrors = 52
	XmlParserErrorsXMLERRMIXEDNOTFINISHED                     XmlParserErrors = 53
	XmlParserErrorsXMLERRELEMCONTENTNOTSTARTED                XmlParserErrors = 54
	XmlParserErrorsXMLERRELEMCONTENTNOTFINISHED               XmlParserErrors = 55
	XmlParserErrorsXMLERRXMLDECLNOTSTARTED                    XmlParserErrors = 56
	XmlParserErrorsXMLERRXMLDECLNOTFINISHED                   XmlParserErrors = 57
	XmlParserErrorsXMLERRCONDSECNOTSTARTED                    XmlParserErrors = 58
	XmlParserErrorsXMLERRCONDSECNOTFINISHED                   XmlParserErrors = 59
	XmlParserErrorsXMLERREXTSUBSETNOTFINISHED                 XmlParserErrors = 60
	XmlParserErrorsXMLERRDOCTYPENOTFINISHED                   XmlParserErrors = 61
	XmlParserErrorsXMLERRMISPLACEDCDATAEND                    XmlParserErrors = 62
	XmlParserErrorsXMLERRCDATANOTFINISHED                     XmlParserErrors = 63
	XmlParserErrorsXMLERRRESERVEDXMLNAME                      XmlParserErrors = 64
	XmlParserErrorsXMLERRSPACEREQUIRED                        XmlParserErrors = 65
	XmlParserErrorsXMLERRSEPARATORREQUIRED                    XmlParserErrors = 66
	XmlParserErrorsXMLERRNMTOKENREQUIRED                      XmlParserErrors = 67
	XmlParserErrorsXMLERRNAMEREQUIRED                         XmlParserErrors = 68
	XmlParserErrorsXMLERRPCDATAREQUIRED                       XmlParserErrors = 69
	XmlParserErrorsXMLERRURIREQUIRED                          XmlParserErrors = 70
	XmlParserErrorsXMLERRPUBIDREQUIRED                        XmlParserErrors = 71
	XmlParserErrorsXMLERRLTREQUIRED                           XmlParserErrors = 72
	XmlParserErrorsXMLERRGTREQUIRED                           XmlParserErrors = 73
	XmlParserErrorsXMLERRLTSLASHREQUIRED                      XmlParserErrors = 74
	XmlParserErrorsXMLERREQUALREQUIRED                        XmlParserErrors = 75
	XmlParserErrorsXMLERRTAGNAMEMISMATCH                      XmlParserErrors = 76
	XmlParserErrorsXMLERRTAGNOTFINISHED                       XmlParserErrors = 77
	XmlParserErrorsXMLERRSTANDALONEVALUE                      XmlParserErrors = 78
	XmlParserErrorsXMLERRENCODINGNAME                         XmlParserErrors = 79
	XmlParserErrorsXMLERRHYPHENINCOMMENT                      XmlParserErrors = 80
	XmlParserErrorsXMLERRINVALIDENCODING                      XmlParserErrors = 81
	XmlParserErrorsXMLERREXTENTITYSTANDALONE                  XmlParserErrors = 82
	XmlParserErrorsXMLERRCONDSECINVALID                       XmlParserErrors = 83
	XmlParserErrorsXMLERRVALUEREQUIRED                        XmlParserErrors = 84
	XmlParserErrorsXMLERRNOTWELLBALANCED                      XmlParserErrors = 85
	XmlParserErrorsXMLERREXTRACONTENT                         XmlParserErrors = 86
	XmlParserErrorsXMLERRENTITYCHARERROR                      XmlParserErrors = 87
	XmlParserErrorsXMLERRENTITYPEINTERNAL                     XmlParserErrors = 88
	XmlParserErrorsXMLERRENTITYLOOP                           XmlParserErrors = 89
	XmlParserErrorsXMLERRENTITYBOUNDARY                       XmlParserErrors = 90
	XmlParserErrorsXMLERRINVALIDURI                           XmlParserErrors = 91
	XmlParserErrorsXMLERRURIFRAGMENT                          XmlParserErrors = 92
	XmlParserErrorsXMLWARCATALOGPI                            XmlParserErrors = 93
	XmlParserErrorsXMLERRNODTD                                XmlParserErrors = 94
	XmlParserErrorsXMLERRCONDSECINVALIDKEYWORD                XmlParserErrors = 95
	XmlParserErrorsXMLERRVERSIONMISSING                       XmlParserErrors = 96
	XmlParserErrorsXMLWARUNKNOWNVERSION                       XmlParserErrors = 97
	XmlParserErrorsXMLWARLANGVALUE                            XmlParserErrors = 98
	XmlParserErrorsXMLWARNSURI                                XmlParserErrors = 99
	XmlParserErrorsXMLWARNSURIRELATIVE                        XmlParserErrors = 100
	XmlParserErrorsXMLERRMISSINGENCODING                      XmlParserErrors = 101
	XmlParserErrorsXMLWARSPACEVALUE                           XmlParserErrors = 102
	XmlParserErrorsXMLERRNOTSTANDALONE                        XmlParserErrors = 103
	XmlParserErrorsXMLERRENTITYPROCESSING                     XmlParserErrors = 104
	XmlParserErrorsXMLERRNOTATIONPROCESSING                   XmlParserErrors = 105
	XmlParserErrorsXMLWARNSCOLUMN                             XmlParserErrors = 106
	XmlParserErrorsXMLWARENTITYREDEFINED                      XmlParserErrors = 107
	XmlParserErrorsXMLERRUNKNOWNVERSION                       XmlParserErrors = 108
	XmlParserErrorsXMLERRVERSIONMISMATCH                      XmlParserErrors = 109
	XmlParserErrorsXMLERRNAMETOOLONG                          XmlParserErrors = 110
	XmlParserErrorsXMLERRUSERSTOP                             XmlParserErrors = 111
	XmlParserErrorsXMLERRCOMMENTABRUPTLYENDED                 XmlParserErrors = 112
	XmlParserErrorsXMLWARENCODINGMISMATCH                     XmlParserErrors = 113
	XmlParserErrorsXMLERRRESOURCELIMIT                        XmlParserErrors = 114
	XmlParserErrorsXMLERRARGUMENT                             XmlParserErrors = 115
	XmlParserErrorsXMLERRSYSTEM                               XmlParserErrors = 116
	XmlParserErrorsXMLERRREDECLPREDEFENTITY                   XmlParserErrors = 117
	XmlParserErrorsXMLERRINTSUBSETNOTFINISHED                 XmlParserErrors = 118
	XmlParserErrorsXMLNSERRXMLNAMESPACE                       XmlParserErrors = 200
	XmlParserErrorsXMLNSERRUNDEFINEDNAMESPACE                 XmlParserErrors = 201
	XmlParserErrorsXMLNSERRQNAME                              XmlParserErrors = 202
	XmlParserErrorsXMLNSERRATTRIBUTEREDEFINED                 XmlParserErrors = 203
	XmlParserErrorsXMLNSERREMPTY                              XmlParserErrors = 204
	XmlParserErrorsXMLNSERRCOLON                              XmlParserErrors = 205
	XmlParserErrorsXMLDTDATTRIBUTEDEFAULT                     XmlParserErrors = 500
	XmlParserErrorsXMLDTDATTRIBUTEREDEFINED                   XmlParserErrors = 501
	XmlParserErrorsXMLDTDATTRIBUTEVALUE                       XmlParserErrors = 502
	XmlParserErrorsXMLDTDCONTENTERROR                         XmlParserErrors = 503
	XmlParserErrorsXMLDTDCONTENTMODEL                         XmlParserErrors = 504
	XmlParserErrorsXMLDTDCONTENTNOTDETERMINIST                XmlParserErrors = 505
	XmlParserErrorsXMLDTDDIFFERENTPREFIX                      XmlParserErrors = 506
	XmlParserErrorsXMLDTDELEMDEFAULTNAMESPACE                 XmlParserErrors = 507
	XmlParserErrorsXMLDTDELEMNAMESPACE                        XmlParserErrors = 508
	XmlParserErrorsXMLDTDELEMREDEFINED                        XmlParserErrors = 509
	XmlParserErrorsXMLDTDEMPTYNOTATION                        XmlParserErrors = 510
	XmlParserErrorsXMLDTDENTITYTYPE                           XmlParserErrors = 511
	XmlParserErrorsXMLDTDIDFIXED                              XmlParserErrors = 512
	XmlParserErrorsXMLDTDIDREDEFINED                          XmlParserErrors = 513
	XmlParserErrorsXMLDTDIDSUBSET                             XmlParserErrors = 514
	XmlParserErrorsXMLDTDINVALIDCHILD                         XmlParserErrors = 515
	XmlParserErrorsXMLDTDINVALIDDEFAULT                       XmlParserErrors = 516
	XmlParserErrorsXMLDTDLOADERROR                            XmlParserErrors = 517
	XmlParserErrorsXMLDTDMISSINGATTRIBUTE                     XmlParserErrors = 518
	XmlParserErrorsXMLDTDMIXEDCORRUPT                         XmlParserErrors = 519
	XmlParserErrorsXMLDTDMULTIPLEID                           XmlParserErrors = 520
	XmlParserErrorsXMLDTDNODOC                                XmlParserErrors = 521
	XmlParserErrorsXMLDTDNODTD                                XmlParserErrors = 522
	XmlParserErrorsXMLDTDNOELEMNAME                           XmlParserErrors = 523
	XmlParserErrorsXMLDTDNOPREFIX                             XmlParserErrors = 524
	XmlParserErrorsXMLDTDNOROOT                               XmlParserErrors = 525
	XmlParserErrorsXMLDTDNOTATIONREDEFINED                    XmlParserErrors = 526
	XmlParserErrorsXMLDTDNOTATIONVALUE                        XmlParserErrors = 527
	XmlParserErrorsXMLDTDNOTEMPTY                             XmlParserErrors = 528
	XmlParserErrorsXMLDTDNOTPCDATA                            XmlParserErrors = 529
	XmlParserErrorsXMLDTDNOTSTANDALONE                        XmlParserErrors = 530
	XmlParserErrorsXMLDTDROOTNAME                             XmlParserErrors = 531
	XmlParserErrorsXMLDTDSTANDALONEWHITESPACE                 XmlParserErrors = 532
	XmlParserErrorsXMLDTDUNKNOWNATTRIBUTE                     XmlParserErrors = 533
	XmlParserErrorsXMLDTDUNKNOWNELEM                          XmlParserErrors = 534
	XmlParserErrorsXMLDTDUNKNOWNENTITY                        XmlParserErrors = 535
	XmlParserErrorsXMLDTDUNKNOWNID                            XmlParserErrors = 536
	XmlParserErrorsXMLDTDUNKNOWNNOTATION                      XmlParserErrors = 537
	XmlParserErrorsXMLDTDSTANDALONEDEFAULTED                  XmlParserErrors = 538
	XmlParserErrorsXMLDTDXMLIDVALUE                           XmlParserErrors = 539
	XmlParserErrorsXMLDTDXMLIDTYPE                            XmlParserErrors = 540
	XmlParserErrorsXMLDTDDUPTOKEN                             XmlParserErrors = 541
	XmlParserErrorsXMLHTMLSTRUCUREERROR                       XmlParserErrors = 800
	XmlParserErrorsXMLHTMLUNKNOWNTAG                          XmlParserErrors = 801
	XmlParserErrorsXMLHTMLINCORRECTLYOPENEDCOMMENT            XmlParserErrors = 802
	XmlParserErrorsXMLRNGPANYNAMEATTRANCESTOR                 XmlParserErrors = 1000
	XmlParserErrorsXMLRNGPATTRCONFLICT                        XmlParserErrors = 1001
	XmlParserErrorsXMLRNGPATTRIBUTECHILDREN                   XmlParserErrors = 1002
	XmlParserErrorsXMLRNGPATTRIBUTECONTENT                    XmlParserErrors = 1003
	XmlParserErrorsXMLRNGPATTRIBUTEEMPTY                      XmlParserErrors = 1004
	XmlParserErrorsXMLRNGPATTRIBUTENOOP                       XmlParserErrors = 1005
	XmlParserErrorsXMLRNGPCHOICECONTENT                       XmlParserErrors = 1006
	XmlParserErrorsXMLRNGPCHOICEEMPTY                         XmlParserErrors = 1007
	XmlParserErrorsXMLRNGPCREATEFAILURE                       XmlParserErrors = 1008
	XmlParserErrorsXMLRNGPDATACONTENT                         XmlParserErrors = 1009
	XmlParserErrorsXMLRNGPDEFCHOICEANDINTERLEAVE              XmlParserErrors = 1010
	XmlParserErrorsXMLRNGPDEFINECREATEFAILED                  XmlParserErrors = 1011
	XmlParserErrorsXMLRNGPDEFINEEMPTY                         XmlParserErrors = 1012
	XmlParserErrorsXMLRNGPDEFINEMISSING                       XmlParserErrors = 1013
	XmlParserErrorsXMLRNGPDEFINENAMEMISSING                   XmlParserErrors = 1014
	XmlParserErrorsXMLRNGPELEMCONTENTEMPTY                    XmlParserErrors = 1015
	XmlParserErrorsXMLRNGPELEMCONTENTERROR                    XmlParserErrors = 1016
	XmlParserErrorsXMLRNGPELEMENTEMPTY                        XmlParserErrors = 1017
	XmlParserErrorsXMLRNGPELEMENTCONTENT                      XmlParserErrors = 1018
	XmlParserErrorsXMLRNGPELEMENTNAME                         XmlParserErrors = 1019
	XmlParserErrorsXMLRNGPELEMENTNOCONTENT                    XmlParserErrors = 1020
	XmlParserErrorsXMLRNGPELEMTEXTCONFLICT                    XmlParserErrors = 1021
	XmlParserErrorsXMLRNGPEMPTY                               XmlParserErrors = 1022
	XmlParserErrorsXMLRNGPEMPTYCONSTRUCT                      XmlParserErrors = 1023
	XmlParserErrorsXMLRNGPEMPTYCONTENT                        XmlParserErrors = 1024
	XmlParserErrorsXMLRNGPEMPTYNOTEMPTY                       XmlParserErrors = 1025
	XmlParserErrorsXMLRNGPERRORTYPELIB                        XmlParserErrors = 1026
	XmlParserErrorsXMLRNGPEXCEPTEMPTY                         XmlParserErrors = 1027
	XmlParserErrorsXMLRNGPEXCEPTMISSING                       XmlParserErrors = 1028
	XmlParserErrorsXMLRNGPEXCEPTMULTIPLE                      XmlParserErrors = 1029
	XmlParserErrorsXMLRNGPEXCEPTNOCONTENT                     XmlParserErrors = 1030
	XmlParserErrorsXMLRNGPEXTERNALREFEMTPY                    XmlParserErrors = 1031
	XmlParserErrorsXMLRNGPEXTERNALREFFAILURE                  XmlParserErrors = 1032
	XmlParserErrorsXMLRNGPEXTERNALREFRECURSE                  XmlParserErrors = 1033
	XmlParserErrorsXMLRNGPFORBIDDENATTRIBUTE                  XmlParserErrors = 1034
	XmlParserErrorsXMLRNGPFOREIGNELEMENT                      XmlParserErrors = 1035
	XmlParserErrorsXMLRNGPGRAMMARCONTENT                      XmlParserErrors = 1036
	XmlParserErrorsXMLRNGPGRAMMAREMPTY                        XmlParserErrors = 1037
	XmlParserErrorsXMLRNGPGRAMMARMISSING                      XmlParserErrors = 1038
	XmlParserErrorsXMLRNGPGRAMMARNOSTART                      XmlParserErrors = 1039
	XmlParserErrorsXMLRNGPGROUPATTRCONFLICT                   XmlParserErrors = 1040
	XmlParserErrorsXMLRNGPHREFERROR                           XmlParserErrors = 1041
	XmlParserErrorsXMLRNGPINCLUDEEMPTY                        XmlParserErrors = 1042
	XmlParserErrorsXMLRNGPINCLUDEFAILURE                      XmlParserErrors = 1043
	XmlParserErrorsXMLRNGPINCLUDERECURSE                      XmlParserErrors = 1044
	XmlParserErrorsXMLRNGPINTERLEAVEADD                       XmlParserErrors = 1045
	XmlParserErrorsXMLRNGPINTERLEAVECREATEFAILED              XmlParserErrors = 1046
	XmlParserErrorsXMLRNGPINTERLEAVEEMPTY                     XmlParserErrors = 1047
	XmlParserErrorsXMLRNGPINTERLEAVENOCONTENT                 XmlParserErrors = 1048
	XmlParserErrorsXMLRNGPINVALIDDEFINENAME                   XmlParserErrors = 1049
	XmlParserErrorsXMLRNGPINVALIDURI                          XmlParserErrors = 1050
	XmlParserErrorsXMLRNGPINVALIDVALUE                        XmlParserErrors = 1051
	XmlParserErrorsXMLRNGPMISSINGHREF                         XmlParserErrors = 1052
	XmlParserErrorsXMLRNGPNAMEMISSING                         XmlParserErrors = 1053
	XmlParserErrorsXMLRNGPNEEDCOMBINE                         XmlParserErrors = 1054
	XmlParserErrorsXMLRNGPNOTALLOWEDNOTEMPTY                  XmlParserErrors = 1055
	XmlParserErrorsXMLRNGPNSNAMEATTRANCESTOR                  XmlParserErrors = 1056
	XmlParserErrorsXMLRNGPNSNAMENONS                          XmlParserErrors = 1057
	XmlParserErrorsXMLRNGPPARAMFORBIDDEN                      XmlParserErrors = 1058
	XmlParserErrorsXMLRNGPPARAMNAMEMISSING                    XmlParserErrors = 1059
	XmlParserErrorsXMLRNGPPARENTREFCREATEFAILED               XmlParserErrors = 1060
	XmlParserErrorsXMLRNGPPARENTREFNAMEINVALID                XmlParserErrors = 1061
	XmlParserErrorsXMLRNGPPARENTREFNONAME                     XmlParserErrors = 1062
	XmlParserErrorsXMLRNGPPARENTREFNOPARENT                   XmlParserErrors = 1063
	XmlParserErrorsXMLRNGPPARENTREFNOTEMPTY                   XmlParserErrors = 1064
	XmlParserErrorsXMLRNGPPARSEERROR                          XmlParserErrors = 1065
	XmlParserErrorsXMLRNGPPATANYNAMEEXCEPTANYNAME             XmlParserErrors = 1066
	XmlParserErrorsXMLRNGPPATATTRATTR                         XmlParserErrors = 1067
	XmlParserErrorsXMLRNGPPATATTRELEM                         XmlParserErrors = 1068
	XmlParserErrorsXMLRNGPPATDATAEXCEPTATTR                   XmlParserErrors = 1069
	XmlParserErrorsXMLRNGPPATDATAEXCEPTELEM                   XmlParserErrors = 1070
	XmlParserErrorsXMLRNGPPATDATAEXCEPTEMPTY                  XmlParserErrors = 1071
	XmlParserErrorsXMLRNGPPATDATAEXCEPTGROUP                  XmlParserErrors = 1072
	XmlParserErrorsXMLRNGPPATDATAEXCEPTINTERLEAVE             XmlParserErrors = 1073
	XmlParserErrorsXMLRNGPPATDATAEXCEPTLIST                   XmlParserErrors = 1074
	XmlParserErrorsXMLRNGPPATDATAEXCEPTONEMORE                XmlParserErrors = 1075
	XmlParserErrorsXMLRNGPPATDATAEXCEPTREF                    XmlParserErrors = 1076
	XmlParserErrorsXMLRNGPPATDATAEXCEPTTEXT                   XmlParserErrors = 1077
	XmlParserErrorsXMLRNGPPATLISTATTR                         XmlParserErrors = 1078
	XmlParserErrorsXMLRNGPPATLISTELEM                         XmlParserErrors = 1079
	XmlParserErrorsXMLRNGPPATLISTINTERLEAVE                   XmlParserErrors = 1080
	XmlParserErrorsXMLRNGPPATLISTLIST                         XmlParserErrors = 1081
	XmlParserErrorsXMLRNGPPATLISTREF                          XmlParserErrors = 1082
	XmlParserErrorsXMLRNGPPATLISTTEXT                         XmlParserErrors = 1083
	XmlParserErrorsXMLRNGPPATNSNAMEEXCEPTANYNAME              XmlParserErrors = 1084
	XmlParserErrorsXMLRNGPPATNSNAMEEXCEPTNSNAME               XmlParserErrors = 1085
	XmlParserErrorsXMLRNGPPATONEMOREGROUPATTR                 XmlParserErrors = 1086
	XmlParserErrorsXMLRNGPPATONEMOREINTERLEAVEATTR            XmlParserErrors = 1087
	XmlParserErrorsXMLRNGPPATSTARTATTR                        XmlParserErrors = 1088
	XmlParserErrorsXMLRNGPPATSTARTDATA                        XmlParserErrors = 1089
	XmlParserErrorsXMLRNGPPATSTARTEMPTY                       XmlParserErrors = 1090
	XmlParserErrorsXMLRNGPPATSTARTGROUP                       XmlParserErrors = 1091
	XmlParserErrorsXMLRNGPPATSTARTINTERLEAVE                  XmlParserErrors = 1092
	XmlParserErrorsXMLRNGPPATSTARTLIST                        XmlParserErrors = 1093
	XmlParserErrorsXMLRNGPPATSTARTONEMORE                     XmlParserErrors = 1094
	XmlParserErrorsXMLRNGPPATSTARTTEXT                        XmlParserErrors = 1095
	XmlParserErrorsXMLRNGPPATSTARTVALUE                       XmlParserErrors = 1096
	XmlParserErrorsXMLRNGPPREFIXUNDEFINED                     XmlParserErrors = 1097
	XmlParserErrorsXMLRNGPREFCREATEFAILED                     XmlParserErrors = 1098
	XmlParserErrorsXMLRNGPREFCYCLE                            XmlParserErrors = 1099
	XmlParserErrorsXMLRNGPREFNAMEINVALID                      XmlParserErrors = 1100
	XmlParserErrorsXMLRNGPREFNODEF                            XmlParserErrors = 1101
	XmlParserErrorsXMLRNGPREFNONAME                           XmlParserErrors = 1102
	XmlParserErrorsXMLRNGPREFNOTEMPTY                         XmlParserErrors = 1103
	XmlParserErrorsXMLRNGPSTARTCHOICEANDINTERLEAVE            XmlParserErrors = 1104
	XmlParserErrorsXMLRNGPSTARTCONTENT                        XmlParserErrors = 1105
	XmlParserErrorsXMLRNGPSTARTEMPTY                          XmlParserErrors = 1106
	XmlParserErrorsXMLRNGPSTARTMISSING                        XmlParserErrors = 1107
	XmlParserErrorsXMLRNGPTEXTEXPECTED                        XmlParserErrors = 1108
	XmlParserErrorsXMLRNGPTEXTHASCHILD                        XmlParserErrors = 1109
	XmlParserErrorsXMLRNGPTYPEMISSING                         XmlParserErrors = 1110
	XmlParserErrorsXMLRNGPTYPENOTFOUND                        XmlParserErrors = 1111
	XmlParserErrorsXMLRNGPTYPEVALUE                           XmlParserErrors = 1112
	XmlParserErrorsXMLRNGPUNKNOWNATTRIBUTE                    XmlParserErrors = 1113
	XmlParserErrorsXMLRNGPUNKNOWNCOMBINE                      XmlParserErrors = 1114
	XmlParserErrorsXMLRNGPUNKNOWNCONSTRUCT                    XmlParserErrors = 1115
	XmlParserErrorsXMLRNGPUNKNOWNTYPELIB                      XmlParserErrors = 1116
	XmlParserErrorsXMLRNGPURIFRAGMENT                         XmlParserErrors = 1117
	XmlParserErrorsXMLRNGPURINOTABSOLUTE                      XmlParserErrors = 1118
	XmlParserErrorsXMLRNGPVALUEEMPTY                          XmlParserErrors = 1119
	XmlParserErrorsXMLRNGPVALUENOCONTENT                      XmlParserErrors = 1120
	XmlParserErrorsXMLRNGPXMLNSNAME                           XmlParserErrors = 1121
	XmlParserErrorsXMLRNGPXMLNS                               XmlParserErrors = 1122
	XmlParserErrorsXMLXPATHEXPRESSIONOK                       XmlParserErrors = 1200
	XmlParserErrorsXMLXPATHNUMBERERROR                        XmlParserErrors = 1201
	XmlParserErrorsXMLXPATHUNFINISHEDLITERALERROR             XmlParserErrors = 1202
	XmlParserErrorsXMLXPATHSTARTLITERALERROR                  XmlParserErrors = 1203
	XmlParserErrorsXMLXPATHVARIABLEREFERROR                   XmlParserErrors = 1204
	XmlParserErrorsXMLXPATHUNDEFVARIABLEERROR                 XmlParserErrors = 1205
	XmlParserErrorsXMLXPATHINVALIDPREDICATEERROR              XmlParserErrors = 1206
	XmlParserErrorsXMLXPATHEXPRERROR                          XmlParserErrors = 1207
	XmlParserErrorsXMLXPATHUNCLOSEDERROR                      XmlParserErrors = 1208
	XmlParserErrorsXMLXPATHUNKNOWNFUNCERROR                   XmlParserErrors = 1209
	XmlParserErrorsXMLXPATHINVALIDOPERAND                     XmlParserErrors = 1210
	XmlParserErrorsXMLXPATHINVALIDTYPE                        XmlParserErrors = 1211
	XmlParserErrorsXMLXPATHINVALIDARITY                       XmlParserErrors = 1212
	XmlParserErrorsXMLXPATHINVALIDCTXTSIZE                    XmlParserErrors = 1213
	XmlParserErrorsXMLXPATHINVALIDCTXTPOSITION                XmlParserErrors = 1214
	XmlParserErrorsXMLXPATHMEMORYERROR                        XmlParserErrors = 1215
	XmlParserErrorsXMLXPTRSYNTAXERROR                         XmlParserErrors = 1216
	XmlParserErrorsXMLXPTRRESOURCEERROR                       XmlParserErrors = 1217
	XmlParserErrorsXMLXPTRSUBRESOURCEERROR                    XmlParserErrors = 1218
	XmlParserErrorsXMLXPATHUNDEFPREFIXERROR                   XmlParserErrors = 1219
	XmlParserErrorsXMLXPATHENCODINGERROR                      XmlParserErrors = 1220
	XmlParserErrorsXMLXPATHINVALIDCHARERROR                   XmlParserErrors = 1221
	XmlParserErrorsXMLTREEINVALIDHEX                          XmlParserErrors = 1300
	XmlParserErrorsXMLTREEINVALIDDEC                          XmlParserErrors = 1301
	XmlParserErrorsXMLTREEUNTERMINATEDENTITY                  XmlParserErrors = 1302
	XmlParserErrorsXMLTREENOTUTF8                             XmlParserErrors = 1303
	XmlParserErrorsXMLSAVENOTUTF8                             XmlParserErrors = 1400
	XmlParserErrorsXMLSAVECHARINVALID                         XmlParserErrors = 1401
	XmlParserErrorsXMLSAVENODOCTYPE                           XmlParserErrors = 1402
	XmlParserErrorsXMLSAVEUNKNOWNENCODING                     XmlParserErrors = 1403
	XmlParserErrorsXMLREGEXPCOMPILEERROR                      XmlParserErrors = 1450
	XmlParserErrorsXMLIOUNKNOWN                               XmlParserErrors = 1500
	XmlParserErrorsXMLIOEACCES                                XmlParserErrors = 1501
	XmlParserErrorsXMLIOEAGAIN                                XmlParserErrors = 1502
	XmlParserErrorsXMLIOEBADF                                 XmlParserErrors = 1503
	XmlParserErrorsXMLIOEBADMSG                               XmlParserErrors = 1504
	XmlParserErrorsXMLIOEBUSY                                 XmlParserErrors = 1505
	XmlParserErrorsXMLIOECANCELED                             XmlParserErrors = 1506
	XmlParserErrorsXMLIOECHILD                                XmlParserErrors = 1507
	XmlParserErrorsXMLIOEDEADLK                               XmlParserErrors = 1508
	XmlParserErrorsXMLIOEDOM                                  XmlParserErrors = 1509
	XmlParserErrorsXMLIOEEXIST                                XmlParserErrors = 1510
	XmlParserErrorsXMLIOEFAULT                                XmlParserErrors = 1511
	XmlParserErrorsXMLIOEFBIG                                 XmlParserErrors = 1512
	XmlParserErrorsXMLIOEINPROGRESS                           XmlParserErrors = 1513
	XmlParserErrorsXMLIOEINTR                                 XmlParserErrors = 1514
	XmlParserErrorsXMLIOEINVAL                                XmlParserErrors = 1515
	XmlParserErrorsXMLIOEIO                                   XmlParserErrors = 1516
	XmlParserErrorsXMLIOEISDIR                                XmlParserErrors = 1517
	XmlParserErrorsXMLIOEMFILE                                XmlParserErrors = 1518
	XmlParserErrorsXMLIOEMLINK                                XmlParserErrors = 1519
	XmlParserErrorsXMLIOEMSGSIZE                              XmlParserErrors = 1520
	XmlParserErrorsXMLIOENAMETOOLONG                          XmlParserErrors = 1521
	XmlParserErrorsXMLIOENFILE                                XmlParserErrors = 1522
	XmlParserErrorsXMLIOENODEV                                XmlParserErrors = 1523
	XmlParserErrorsXMLIOENOENT                                XmlParserErrors = 1524
	XmlParserErrorsXMLIOENOEXEC                               XmlParserErrors = 1525
	XmlParserErrorsXMLIOENOLCK                                XmlParserErrors = 1526
	XmlParserErrorsXMLIOENOMEM                                XmlParserErrors = 1527
	XmlParserErrorsXMLIOENOSPC                                XmlParserErrors = 1528
	XmlParserErrorsXMLIOENOSYS                                XmlParserErrors = 1529
	XmlParserErrorsXMLIOENOTDIR                               XmlParserErrors = 1530
	XmlParserErrorsXMLIOENOTEMPTY                             XmlParserErrors = 1531
	XmlParserErrorsXMLIOENOTSUP                               XmlParserErrors = 1532
	XmlParserErrorsXMLIOENOTTY                                XmlParserErrors = 1533
	XmlParserErrorsXMLIOENXIO                                 XmlParserErrors = 1534
	XmlParserErrorsXMLIOEPERM                                 XmlParserErrors = 1535
	XmlParserErrorsXMLIOEPIPE                                 XmlParserErrors = 1536
	XmlParserErrorsXMLIOERANGE                                XmlParserErrors = 1537
	XmlParserErrorsXMLIOEROFS                                 XmlParserErrors = 1538
	XmlParserErrorsXMLIOESPIPE                                XmlParserErrors = 1539
	XmlParserErrorsXMLIOESRCH                                 XmlParserErrors = 1540
	XmlParserErrorsXMLIOETIMEDOUT                             XmlParserErrors = 1541
	XmlParserErrorsXMLIOEXDEV                                 XmlParserErrors = 1542
	XmlParserErrorsXMLIONETWORKATTEMPT                        XmlParserErrors = 1543
	XmlParserErrorsXMLIOENCODER                               XmlParserErrors = 1544
	XmlParserErrorsXMLIOFLUSH                                 XmlParserErrors = 1545
	XmlParserErrorsXMLIOWRITE                                 XmlParserErrors = 1546
	XmlParserErrorsXMLIONOINPUT                               XmlParserErrors = 1547
	XmlParserErrorsXMLIOBUFFERFULL                            XmlParserErrors = 1548
	XmlParserErrorsXMLIOLOADERROR                             XmlParserErrors = 1549
	XmlParserErrorsXMLIOENOTSOCK                              XmlParserErrors = 1550
	XmlParserErrorsXMLIOEISCONN                               XmlParserErrors = 1551
	XmlParserErrorsXMLIOECONNREFUSED                          XmlParserErrors = 1552
	XmlParserErrorsXMLIOENETUNREACH                           XmlParserErrors = 1553
	XmlParserErrorsXMLIOEADDRINUSE                            XmlParserErrors = 1554
	XmlParserErrorsXMLIOEALREADY                              XmlParserErrors = 1555
	XmlParserErrorsXMLIOEAFNOSUPPORT                          XmlParserErrors = 1556
	XmlParserErrorsXMLIOUNSUPPORTEDPROTOCOL                   XmlParserErrors = 1557
	XmlParserErrorsXMLXINCLUDERECURSION                       XmlParserErrors = 1600
	XmlParserErrorsXMLXINCLUDEPARSEVALUE                      XmlParserErrors = 1601
	XmlParserErrorsXMLXINCLUDEENTITYDEFMISMATCH               XmlParserErrors = 1602
	XmlParserErrorsXMLXINCLUDENOHREF                          XmlParserErrors = 1603
	XmlParserErrorsXMLXINCLUDENOFALLBACK                      XmlParserErrors = 1604
	XmlParserErrorsXMLXINCLUDEHREFURI                         XmlParserErrors = 1605
	XmlParserErrorsXMLXINCLUDETEXTFRAGMENT                    XmlParserErrors = 1606
	XmlParserErrorsXMLXINCLUDETEXTDOCUMENT                    XmlParserErrors = 1607
	XmlParserErrorsXMLXINCLUDEINVALIDCHAR                     XmlParserErrors = 1608
	XmlParserErrorsXMLXINCLUDEBUILDFAILED                     XmlParserErrors = 1609
	XmlParserErrorsXMLXINCLUDEUNKNOWNENCODING                 XmlParserErrors = 1610
	XmlParserErrorsXMLXINCLUDEMULTIPLEROOT                    XmlParserErrors = 1611
	XmlParserErrorsXMLXINCLUDEXPTRFAILED                      XmlParserErrors = 1612
	XmlParserErrorsXMLXINCLUDEXPTRRESULT                      XmlParserErrors = 1613
	XmlParserErrorsXMLXINCLUDEINCLUDEININCLUDE                XmlParserErrors = 1614
	XmlParserErrorsXMLXINCLUDEFALLBACKSININCLUDE              XmlParserErrors = 1615
	XmlParserErrorsXMLXINCLUDEFALLBACKNOTININCLUDE            XmlParserErrors = 1616
	XmlParserErrorsXMLXINCLUDEDEPRECATEDNS                    XmlParserErrors = 1617
	XmlParserErrorsXMLXINCLUDEFRAGMENTID                      XmlParserErrors = 1618
	XmlParserErrorsXMLCATALOGMISSINGATTR                      XmlParserErrors = 1650
	XmlParserErrorsXMLCATALOGENTRYBROKEN                      XmlParserErrors = 1651
	XmlParserErrorsXMLCATALOGPREFERVALUE                      XmlParserErrors = 1652
	XmlParserErrorsXMLCATALOGNOTCATALOG                       XmlParserErrors = 1653
	XmlParserErrorsXMLCATALOGRECURSION                        XmlParserErrors = 1654
	XmlParserErrorsXMLSCHEMAPPREFIXUNDEFINED                  XmlParserErrors = 1700
	XmlParserErrorsXMLSCHEMAPATTRFORMDEFAULTVALUE             XmlParserErrors = 1701
	XmlParserErrorsXMLSCHEMAPATTRGRPNONAMENOREF               XmlParserErrors = 1702
	XmlParserErrorsXMLSCHEMAPATTRNONAMENOREF                  XmlParserErrors = 1703
	XmlParserErrorsXMLSCHEMAPCOMPLEXTYPENONAMENOREF           XmlParserErrors = 1704
	XmlParserErrorsXMLSCHEMAPELEMFORMDEFAULTVALUE             XmlParserErrors = 1705
	XmlParserErrorsXMLSCHEMAPELEMNONAMENOREF                  XmlParserErrors = 1706
	XmlParserErrorsXMLSCHEMAPEXTENSIONNOBASE                  XmlParserErrors = 1707
	XmlParserErrorsXMLSCHEMAPFACETNOVALUE                     XmlParserErrors = 1708
	XmlParserErrorsXMLSCHEMAPFAILEDBUILDIMPORT                XmlParserErrors = 1709
	XmlParserErrorsXMLSCHEMAPGROUPNONAMENOREF                 XmlParserErrors = 1710
	XmlParserErrorsXMLSCHEMAPIMPORTNAMESPACENOTURI            XmlParserErrors = 1711
	XmlParserErrorsXMLSCHEMAPIMPORTREDEFINENSNAME             XmlParserErrors = 1712
	XmlParserErrorsXMLSCHEMAPIMPORTSCHEMANOTURI               XmlParserErrors = 1713
	XmlParserErrorsXMLSCHEMAPINVALIDBOOLEAN                   XmlParserErrors = 1714
	XmlParserErrorsXMLSCHEMAPINVALIDENUM                      XmlParserErrors = 1715
	XmlParserErrorsXMLSCHEMAPINVALIDFACET                     XmlParserErrors = 1716
	XmlParserErrorsXMLSCHEMAPINVALIDFACETVALUE                XmlParserErrors = 1717
	XmlParserErrorsXMLSCHEMAPINVALIDMAXOCCURS                 XmlParserErrors = 1718
	XmlParserErrorsXMLSCHEMAPINVALIDMINOCCURS                 XmlParserErrors = 1719
	XmlParserErrorsXMLSCHEMAPINVALIDREFANDSUBTYPE             XmlParserErrors = 1720
	XmlParserErrorsXMLSCHEMAPINVALIDWHITESPACE                XmlParserErrors = 1721
	XmlParserErrorsXMLSCHEMAPNOATTRNOREF                      XmlParserErrors = 1722
	XmlParserErrorsXMLSCHEMAPNOTATIONNONAME                   XmlParserErrors = 1723
	XmlParserErrorsXMLSCHEMAPNOTYPENOREF                      XmlParserErrors = 1724
	XmlParserErrorsXMLSCHEMAPREFANDSUBTYPE                    XmlParserErrors = 1725
	XmlParserErrorsXMLSCHEMAPRESTRICTIONNONAMENOREF           XmlParserErrors = 1726
	XmlParserErrorsXMLSCHEMAPSIMPLETYPENONAME                 XmlParserErrors = 1727
	XmlParserErrorsXMLSCHEMAPTYPEANDSUBTYPE                   XmlParserErrors = 1728
	XmlParserErrorsXMLSCHEMAPUNKNOWNALLCHILD                  XmlParserErrors = 1729
	XmlParserErrorsXMLSCHEMAPUNKNOWNANYATTRIBUTECHILD         XmlParserErrors = 1730
	XmlParserErrorsXMLSCHEMAPUNKNOWNATTRCHILD                 XmlParserErrors = 1731
	XmlParserErrorsXMLSCHEMAPUNKNOWNATTRGRPCHILD              XmlParserErrors = 1732
	XmlParserErrorsXMLSCHEMAPUNKNOWNATTRIBUTEGROUP            XmlParserErrors = 1733
	XmlParserErrorsXMLSCHEMAPUNKNOWNBASETYPE                  XmlParserErrors = 1734
	XmlParserErrorsXMLSCHEMAPUNKNOWNCHOICECHILD               XmlParserErrors = 1735
	XmlParserErrorsXMLSCHEMAPUNKNOWNCOMPLEXCONTENTCHILD       XmlParserErrors = 1736
	XmlParserErrorsXMLSCHEMAPUNKNOWNCOMPLEXTYPECHILD          XmlParserErrors = 1737
	XmlParserErrorsXMLSCHEMAPUNKNOWNELEMCHILD                 XmlParserErrors = 1738
	XmlParserErrorsXMLSCHEMAPUNKNOWNEXTENSIONCHILD            XmlParserErrors = 1739
	XmlParserErrorsXMLSCHEMAPUNKNOWNFACETCHILD                XmlParserErrors = 1740
	XmlParserErrorsXMLSCHEMAPUNKNOWNFACETTYPE                 XmlParserErrors = 1741
	XmlParserErrorsXMLSCHEMAPUNKNOWNGROUPCHILD                XmlParserErrors = 1742
	XmlParserErrorsXMLSCHEMAPUNKNOWNIMPORTCHILD               XmlParserErrors = 1743
	XmlParserErrorsXMLSCHEMAPUNKNOWNLISTCHILD                 XmlParserErrors = 1744
	XmlParserErrorsXMLSCHEMAPUNKNOWNNOTATIONCHILD             XmlParserErrors = 1745
	XmlParserErrorsXMLSCHEMAPUNKNOWNPROCESSCONTENTCHILD       XmlParserErrors = 1746
	XmlParserErrorsXMLSCHEMAPUNKNOWNREF                       XmlParserErrors = 1747
	XmlParserErrorsXMLSCHEMAPUNKNOWNRESTRICTIONCHILD          XmlParserErrors = 1748
	XmlParserErrorsXMLSCHEMAPUNKNOWNSCHEMASCHILD              XmlParserErrors = 1749
	XmlParserErrorsXMLSCHEMAPUNKNOWNSEQUENCECHILD             XmlParserErrors = 1750
	XmlParserErrorsXMLSCHEMAPUNKNOWNSIMPLECONTENTCHILD        XmlParserErrors = 1751
	XmlParserErrorsXMLSCHEMAPUNKNOWNSIMPLETYPECHILD           XmlParserErrors = 1752
	XmlParserErrorsXMLSCHEMAPUNKNOWNTYPE                      XmlParserErrors = 1753
	XmlParserErrorsXMLSCHEMAPUNKNOWNUNIONCHILD                XmlParserErrors = 1754
	XmlParserErrorsXMLSCHEMAPELEMDEFAULTFIXED                 XmlParserErrors = 1755
	XmlParserErrorsXMLSCHEMAPREGEXPINVALID                    XmlParserErrors = 1756
	XmlParserErrorsXMLSCHEMAPFAILEDLOAD                       XmlParserErrors = 1757
	XmlParserErrorsXMLSCHEMAPNOTHINGTOPARSE                   XmlParserErrors = 1758
	XmlParserErrorsXMLSCHEMAPNOROOT                           XmlParserErrors = 1759
	XmlParserErrorsXMLSCHEMAPREDEFINEDGROUP                   XmlParserErrors = 1760
	XmlParserErrorsXMLSCHEMAPREDEFINEDTYPE                    XmlParserErrors = 1761
	XmlParserErrorsXMLSCHEMAPREDEFINEDELEMENT                 XmlParserErrors = 1762
	XmlParserErrorsXMLSCHEMAPREDEFINEDATTRGROUP               XmlParserErrors = 1763
	XmlParserErrorsXMLSCHEMAPREDEFINEDATTR                    XmlParserErrors = 1764
	XmlParserErrorsXMLSCHEMAPREDEFINEDNOTATION                XmlParserErrors = 1765
	XmlParserErrorsXMLSCHEMAPFAILEDPARSE                      XmlParserErrors = 1766
	XmlParserErrorsXMLSCHEMAPUNKNOWNPREFIX                    XmlParserErrors = 1767
	XmlParserErrorsXMLSCHEMAPDEFANDPREFIX                     XmlParserErrors = 1768
	XmlParserErrorsXMLSCHEMAPUNKNOWNINCLUDECHILD              XmlParserErrors = 1769
	XmlParserErrorsXMLSCHEMAPINCLUDESCHEMANOTURI              XmlParserErrors = 1770
	XmlParserErrorsXMLSCHEMAPINCLUDESCHEMANOURI               XmlParserErrors = 1771
	XmlParserErrorsXMLSCHEMAPNOTSCHEMA                        XmlParserErrors = 1772
	XmlParserErrorsXMLSCHEMAPUNKNOWNMEMBERTYPE                XmlParserErrors = 1773
	XmlParserErrorsXMLSCHEMAPINVALIDATTRUSE                   XmlParserErrors = 1774
	XmlParserErrorsXMLSCHEMAPRECURSIVE                        XmlParserErrors = 1775
	XmlParserErrorsXMLSCHEMAPSUPERNUMEROUSLISTITEMTYPE        XmlParserErrors = 1776
	XmlParserErrorsXMLSCHEMAPINVALIDATTRCOMBINATION           XmlParserErrors = 1777
	XmlParserErrorsXMLSCHEMAPINVALIDATTRINLINECOMBINATION     XmlParserErrors = 1778
	XmlParserErrorsXMLSCHEMAPMISSINGSIMPLETYPECHILD           XmlParserErrors = 1779
	XmlParserErrorsXMLSCHEMAPINVALIDATTRNAME                  XmlParserErrors = 1780
	XmlParserErrorsXMLSCHEMAPREFANDCONTENT                    XmlParserErrors = 1781
	XmlParserErrorsXMLSCHEMAPCTPROPSCORRECT1                  XmlParserErrors = 1782
	XmlParserErrorsXMLSCHEMAPCTPROPSCORRECT2                  XmlParserErrors = 1783
	XmlParserErrorsXMLSCHEMAPCTPROPSCORRECT3                  XmlParserErrors = 1784
	XmlParserErrorsXMLSCHEMAPCTPROPSCORRECT4                  XmlParserErrors = 1785
	XmlParserErrorsXMLSCHEMAPCTPROPSCORRECT5                  XmlParserErrors = 1786
	XmlParserErrorsXMLSCHEMAPDERIVATIONOKRESTRICTION1         XmlParserErrors = 1787
	XmlParserErrorsXMLSCHEMAPDERIVATIONOKRESTRICTION211       XmlParserErrors = 1788
	XmlParserErrorsXMLSCHEMAPDERIVATIONOKRESTRICTION212       XmlParserErrors = 1789
	XmlParserErrorsXMLSCHEMAPDERIVATIONOKRESTRICTION22        XmlParserErrors = 1790
	XmlParserErrorsXMLSCHEMAPDERIVATIONOKRESTRICTION3         XmlParserErrors = 1791
	XmlParserErrorsXMLSCHEMAPWILDCARDINVALIDNSMEMBER          XmlParserErrors = 1792
	XmlParserErrorsXMLSCHEMAPINTERSECTIONNOTEXPRESSIBLE       XmlParserErrors = 1793
	XmlParserErrorsXMLSCHEMAPUNIONNOTEXPRESSIBLE              XmlParserErrors = 1794
	XmlParserErrorsXMLSCHEMAPSRCIMPORT31                      XmlParserErrors = 1795
	XmlParserErrorsXMLSCHEMAPSRCIMPORT32                      XmlParserErrors = 1796
	XmlParserErrorsXMLSCHEMAPDERIVATIONOKRESTRICTION41        XmlParserErrors = 1797
	XmlParserErrorsXMLSCHEMAPDERIVATIONOKRESTRICTION42        XmlParserErrors = 1798
	XmlParserErrorsXMLSCHEMAPDERIVATIONOKRESTRICTION43        XmlParserErrors = 1799
	XmlParserErrorsXMLSCHEMAPCOSCTEXTENDS13                   XmlParserErrors = 1800
	XmlParserErrorsXMLSCHEMAVNOROOT                           XmlParserErrors = 1801
	XmlParserErrorsXMLSCHEMAVUNDECLAREDELEM                   XmlParserErrors = 1802
	XmlParserErrorsXMLSCHEMAVNOTTOPLEVEL                      XmlParserErrors = 1803
	XmlParserErrorsXMLSCHEMAVMISSING                          XmlParserErrors = 1804
	XmlParserErrorsXMLSCHEMAVWRONGELEM                        XmlParserErrors = 1805
	XmlParserErrorsXMLSCHEMAVNOTYPE                           XmlParserErrors = 1806
	XmlParserErrorsXMLSCHEMAVNOROLLBACK                       XmlParserErrors = 1807
	XmlParserErrorsXMLSCHEMAVISABSTRACT                       XmlParserErrors = 1808
	XmlParserErrorsXMLSCHEMAVNOTEMPTY                         XmlParserErrors = 1809
	XmlParserErrorsXMLSCHEMAVELEMCONT                         XmlParserErrors = 1810
	XmlParserErrorsXMLSCHEMAVHAVEDEFAULT                      XmlParserErrors = 1811
	XmlParserErrorsXMLSCHEMAVNOTNILLABLE                      XmlParserErrors = 1812
	XmlParserErrorsXMLSCHEMAVEXTRACONTENT                     XmlParserErrors = 1813
	XmlParserErrorsXMLSCHEMAVINVALIDATTR                      XmlParserErrors = 1814
	XmlParserErrorsXMLSCHEMAVINVALIDELEM                      XmlParserErrors = 1815
	XmlParserErrorsXMLSCHEMAVNOTDETERMINIST                   XmlParserErrors = 1816
	XmlParserErrorsXMLSCHEMAVCONSTRUCT                        XmlParserErrors = 1817
	XmlParserErrorsXMLSCHEMAVINTERNAL                         XmlParserErrors = 1818
	XmlParserErrorsXMLSCHEMAVNOTSIMPLE                        XmlParserErrors = 1819
	XmlParserErrorsXMLSCHEMAVATTRUNKNOWN                      XmlParserErrors = 1820
	XmlParserErrorsXMLSCHEMAVATTRINVALID                      XmlParserErrors = 1821
	XmlParserErrorsXMLSCHEMAVVALUE                            XmlParserErrors = 1822
	XmlParserErrorsXMLSCHEMAVFACET                            XmlParserErrors = 1823
	XmlParserErrorsXMLSCHEMAVCVCDATATYPEVALID121              XmlParserErrors = 1824
	XmlParserErrorsXMLSCHEMAVCVCDATATYPEVALID122              XmlParserErrors = 1825
	XmlParserErrorsXMLSCHEMAVCVCDATATYPEVALID123              XmlParserErrors = 1826
	XmlParserErrorsXMLSCHEMAVCVCTYPE311                       XmlParserErrors = 1827
	XmlParserErrorsXMLSCHEMAVCVCTYPE312                       XmlParserErrors = 1828
	XmlParserErrorsXMLSCHEMAVCVCFACETVALID                    XmlParserErrors = 1829
	XmlParserErrorsXMLSCHEMAVCVCLENGTHVALID                   XmlParserErrors = 1830
	XmlParserErrorsXMLSCHEMAVCVCMINLENGTHVALID                XmlParserErrors = 1831
	XmlParserErrorsXMLSCHEMAVCVCMAXLENGTHVALID                XmlParserErrors = 1832
	XmlParserErrorsXMLSCHEMAVCVCMININCLUSIVEVALID             XmlParserErrors = 1833
	XmlParserErrorsXMLSCHEMAVCVCMAXINCLUSIVEVALID             XmlParserErrors = 1834
	XmlParserErrorsXMLSCHEMAVCVCMINEXCLUSIVEVALID             XmlParserErrors = 1835
	XmlParserErrorsXMLSCHEMAVCVCMAXEXCLUSIVEVALID             XmlParserErrors = 1836
	XmlParserErrorsXMLSCHEMAVCVCTOTALDIGITSVALID              XmlParserErrors = 1837
	XmlParserErrorsXMLSCHEMAVCVCFRACTIONDIGITSVALID           XmlParserErrors = 1838
	XmlParserErrorsXMLSCHEMAVCVCPATTERNVALID                  XmlParserErrors = 1839
	XmlParserErrorsXMLSCHEMAVCVCENUMERATIONVALID              XmlParserErrors = 1840
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE21                 XmlParserErrors = 1841
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE22                 XmlParserErrors = 1842
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE23                 XmlParserErrors = 1843
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE24                 XmlParserErrors = 1844
	XmlParserErrorsXMLSCHEMAVCVCELT1                          XmlParserErrors = 1845
	XmlParserErrorsXMLSCHEMAVCVCELT2                          XmlParserErrors = 1846
	XmlParserErrorsXMLSCHEMAVCVCELT31                         XmlParserErrors = 1847
	XmlParserErrorsXMLSCHEMAVCVCELT321                        XmlParserErrors = 1848
	XmlParserErrorsXMLSCHEMAVCVCELT322                        XmlParserErrors = 1849
	XmlParserErrorsXMLSCHEMAVCVCELT41                         XmlParserErrors = 1850
	XmlParserErrorsXMLSCHEMAVCVCELT42                         XmlParserErrors = 1851
	XmlParserErrorsXMLSCHEMAVCVCELT43                         XmlParserErrors = 1852
	XmlParserErrorsXMLSCHEMAVCVCELT511                        XmlParserErrors = 1853
	XmlParserErrorsXMLSCHEMAVCVCELT512                        XmlParserErrors = 1854
	XmlParserErrorsXMLSCHEMAVCVCELT521                        XmlParserErrors = 1855
	XmlParserErrorsXMLSCHEMAVCVCELT5221                       XmlParserErrors = 1856
	XmlParserErrorsXMLSCHEMAVCVCELT52221                      XmlParserErrors = 1857
	XmlParserErrorsXMLSCHEMAVCVCELT52222                      XmlParserErrors = 1858
	XmlParserErrorsXMLSCHEMAVCVCELT6                          XmlParserErrors = 1859
	XmlParserErrorsXMLSCHEMAVCVCELT7                          XmlParserErrors = 1860
	XmlParserErrorsXMLSCHEMAVCVCATTRIBUTE1                    XmlParserErrors = 1861
	XmlParserErrorsXMLSCHEMAVCVCATTRIBUTE2                    XmlParserErrors = 1862
	XmlParserErrorsXMLSCHEMAVCVCATTRIBUTE3                    XmlParserErrors = 1863
	XmlParserErrorsXMLSCHEMAVCVCATTRIBUTE4                    XmlParserErrors = 1864
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE31                 XmlParserErrors = 1865
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE321                XmlParserErrors = 1866
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE322                XmlParserErrors = 1867
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE4                  XmlParserErrors = 1868
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE51                 XmlParserErrors = 1869
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE52                 XmlParserErrors = 1870
	XmlParserErrorsXMLSCHEMAVELEMENTCONTENT                   XmlParserErrors = 1871
	XmlParserErrorsXMLSCHEMAVDOCUMENTELEMENTMISSING           XmlParserErrors = 1872
	XmlParserErrorsXMLSCHEMAVCVCCOMPLEXTYPE1                  XmlParserErrors = 1873
	XmlParserErrorsXMLSCHEMAVCVCAU                            XmlParserErrors = 1874
	XmlParserErrorsXMLSCHEMAVCVCTYPE1                         XmlParserErrors = 1875
	XmlParserErrorsXMLSCHEMAVCVCTYPE2                         XmlParserErrors = 1876
	XmlParserErrorsXMLSCHEMAVCVCIDC                           XmlParserErrors = 1877
	XmlParserErrorsXMLSCHEMAVCVCWILDCARD                      XmlParserErrors = 1878
	XmlParserErrorsXMLSCHEMAVMISC                             XmlParserErrors = 1879
	XmlParserErrorsXMLXPTRUNKNOWNSCHEME                       XmlParserErrors = 1900
	XmlParserErrorsXMLXPTRCHILDSEQSTART                       XmlParserErrors = 1901
	XmlParserErrorsXMLXPTREVALFAILED                          XmlParserErrors = 1902
	XmlParserErrorsXMLXPTREXTRAOBJECTS                        XmlParserErrors = 1903
	XmlParserErrorsXMLC14NCREATECTXT                          XmlParserErrors = 1950
	XmlParserErrorsXMLC14NREQUIRESUTF8                        XmlParserErrors = 1951
	XmlParserErrorsXMLC14NCREATESTACK                         XmlParserErrors = 1952
	XmlParserErrorsXMLC14NINVALIDNODE                         XmlParserErrors = 1953
	XmlParserErrorsXMLC14NUNKNOWNODE                          XmlParserErrors = 1954
	XmlParserErrorsXMLC14NRELATIVENAMESPACE                   XmlParserErrors = 1955
	XmlParserErrorsXMLFTPPASVANSWER                           XmlParserErrors = 2000
	XmlParserErrorsXMLFTPEPSVANSWER                           XmlParserErrors = 2001
	XmlParserErrorsXMLFTPACCNT                                XmlParserErrors = 2002
	XmlParserErrorsXMLFTPURLSYNTAX                            XmlParserErrors = 2003
	XmlParserErrorsXMLHTTPURLSYNTAX                           XmlParserErrors = 2020
	XmlParserErrorsXMLHTTPUSEIP                               XmlParserErrors = 2021
	XmlParserErrorsXMLHTTPUNKNOWNHOST                         XmlParserErrors = 2022
	XmlParserErrorsXMLSCHEMAPSRCSIMPLETYPE1                   XmlParserErrors = 3000
	XmlParserErrorsXMLSCHEMAPSRCSIMPLETYPE2                   XmlParserErrors = 3001
	XmlParserErrorsXMLSCHEMAPSRCSIMPLETYPE3                   XmlParserErrors = 3002
	XmlParserErrorsXMLSCHEMAPSRCSIMPLETYPE4                   XmlParserErrors = 3003
	XmlParserErrorsXMLSCHEMAPSRCRESOLVE                       XmlParserErrors = 3004
	XmlParserErrorsXMLSCHEMAPSRCRESTRICTIONBASEORSIMPLETYPE   XmlParserErrors = 3005
	XmlParserErrorsXMLSCHEMAPSRCLISTITEMTYPEORSIMPLETYPE      XmlParserErrors = 3006
	XmlParserErrorsXMLSCHEMAPSRCUNIONMEMBERTYPESORSIMPLETYPES XmlParserErrors = 3007
	XmlParserErrorsXMLSCHEMAPSTPROPSCORRECT1                  XmlParserErrors = 3008
	XmlParserErrorsXMLSCHEMAPSTPROPSCORRECT2                  XmlParserErrors = 3009
	XmlParserErrorsXMLSCHEMAPSTPROPSCORRECT3                  XmlParserErrors = 3010
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS11                 XmlParserErrors = 3011
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS12                 XmlParserErrors = 3012
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS131                XmlParserErrors = 3013
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS132                XmlParserErrors = 3014
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS21                 XmlParserErrors = 3015
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS2311               XmlParserErrors = 3016
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS2312               XmlParserErrors = 3017
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS2321               XmlParserErrors = 3018
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS2322               XmlParserErrors = 3019
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS2323               XmlParserErrors = 3020
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS2324               XmlParserErrors = 3021
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS2325               XmlParserErrors = 3022
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS31                 XmlParserErrors = 3023
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS331                XmlParserErrors = 3024
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS3312               XmlParserErrors = 3025
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS3322               XmlParserErrors = 3026
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS3321               XmlParserErrors = 3027
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS3323               XmlParserErrors = 3028
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS3324               XmlParserErrors = 3029
	XmlParserErrorsXMLSCHEMAPCOSSTRESTRICTS3325               XmlParserErrors = 3030
	XmlParserErrorsXMLSCHEMAPCOSSTDERIVEDOK21                 XmlParserErrors = 3031
	XmlParserErrorsXMLSCHEMAPCOSSTDERIVEDOK22                 XmlParserErrors = 3032
	XmlParserErrorsXMLSCHEMAPS4SELEMNOTALLOWED                XmlParserErrors = 3033
	XmlParserErrorsXMLSCHEMAPS4SELEMMISSING                   XmlParserErrors = 3034
	XmlParserErrorsXMLSCHEMAPS4SATTRNOTALLOWED                XmlParserErrors = 3035
	XmlParserErrorsXMLSCHEMAPS4SATTRMISSING                   XmlParserErrors = 3036
	XmlParserErrorsXMLSCHEMAPS4SATTRINVALIDVALUE              XmlParserErrors = 3037
	XmlParserErrorsXMLSCHEMAPSRCELEMENT1                      XmlParserErrors = 3038
	XmlParserErrorsXMLSCHEMAPSRCELEMENT21                     XmlParserErrors = 3039
	XmlParserErrorsXMLSCHEMAPSRCELEMENT22                     XmlParserErrors = 3040
	XmlParserErrorsXMLSCHEMAPSRCELEMENT3                      XmlParserErrors = 3041
	XmlParserErrorsXMLSCHEMAPPPROPSCORRECT1                   XmlParserErrors = 3042
	XmlParserErrorsXMLSCHEMAPPPROPSCORRECT21                  XmlParserErrors = 3043
	XmlParserErrorsXMLSCHEMAPPPROPSCORRECT22                  XmlParserErrors = 3044
	XmlParserErrorsXMLSCHEMAPEPROPSCORRECT2                   XmlParserErrors = 3045
	XmlParserErrorsXMLSCHEMAPEPROPSCORRECT3                   XmlParserErrors = 3046
	XmlParserErrorsXMLSCHEMAPEPROPSCORRECT4                   XmlParserErrors = 3047
	XmlParserErrorsXMLSCHEMAPEPROPSCORRECT5                   XmlParserErrors = 3048
	XmlParserErrorsXMLSCHEMAPEPROPSCORRECT6                   XmlParserErrors = 3049
	XmlParserErrorsXMLSCHEMAPSRCINCLUDE                       XmlParserErrors = 3050
	XmlParserErrorsXMLSCHEMAPSRCATTRIBUTE1                    XmlParserErrors = 3051
	XmlParserErrorsXMLSCHEMAPSRCATTRIBUTE2                    XmlParserErrors = 3052
	XmlParserErrorsXMLSCHEMAPSRCATTRIBUTE31                   XmlParserErrors = 3053
	XmlParserErrorsXMLSCHEMAPSRCATTRIBUTE32                   XmlParserErrors = 3054
	XmlParserErrorsXMLSCHEMAPSRCATTRIBUTE4                    XmlParserErrors = 3055
	XmlParserErrorsXMLSCHEMAPNOXMLNS                          XmlParserErrors = 3056
	XmlParserErrorsXMLSCHEMAPNOXSI                            XmlParserErrors = 3057
	XmlParserErrorsXMLSCHEMAPCOSVALIDDEFAULT1                 XmlParserErrors = 3058
	XmlParserErrorsXMLSCHEMAPCOSVALIDDEFAULT21                XmlParserErrors = 3059
	XmlParserErrorsXMLSCHEMAPCOSVALIDDEFAULT221               XmlParserErrors = 3060
	XmlParserErrorsXMLSCHEMAPCOSVALIDDEFAULT222               XmlParserErrors = 3061
	XmlParserErrorsXMLSCHEMAPCVCSIMPLETYPE                    XmlParserErrors = 3062
	XmlParserErrorsXMLSCHEMAPCOSCTEXTENDS11                   XmlParserErrors = 3063
	XmlParserErrorsXMLSCHEMAPSRCIMPORT11                      XmlParserErrors = 3064
	XmlParserErrorsXMLSCHEMAPSRCIMPORT12                      XmlParserErrors = 3065
	XmlParserErrorsXMLSCHEMAPSRCIMPORT2                       XmlParserErrors = 3066
	XmlParserErrorsXMLSCHEMAPSRCIMPORT21                      XmlParserErrors = 3067
	XmlParserErrorsXMLSCHEMAPSRCIMPORT22                      XmlParserErrors = 3068
	XmlParserErrorsXMLSCHEMAPINTERNAL                         XmlParserErrors = 3069
	XmlParserErrorsXMLSCHEMAPNOTDETERMINISTIC                 XmlParserErrors = 3070
	XmlParserErrorsXMLSCHEMAPSRCATTRIBUTEGROUP1               XmlParserErrors = 3071
	XmlParserErrorsXMLSCHEMAPSRCATTRIBUTEGROUP2               XmlParserErrors = 3072
	XmlParserErrorsXMLSCHEMAPSRCATTRIBUTEGROUP3               XmlParserErrors = 3073
	XmlParserErrorsXMLSCHEMAPMGPROPSCORRECT1                  XmlParserErrors = 3074
	XmlParserErrorsXMLSCHEMAPMGPROPSCORRECT2                  XmlParserErrors = 3075
	XmlParserErrorsXMLSCHEMAPSRCCT1                           XmlParserErrors = 3076
	XmlParserErrorsXMLSCHEMAPDERIVATIONOKRESTRICTION213       XmlParserErrors = 3077
	XmlParserErrorsXMLSCHEMAPAUPROPSCORRECT2                  XmlParserErrors = 3078
	XmlParserErrorsXMLSCHEMAPAPROPSCORRECT2                   XmlParserErrors = 3079
	XmlParserErrorsXMLSCHEMAPCPROPSCORRECT                    XmlParserErrors = 3080
	XmlParserErrorsXMLSCHEMAPSRCREDEFINE                      XmlParserErrors = 3081
	XmlParserErrorsXMLSCHEMAPSRCIMPORT                        XmlParserErrors = 3082
	XmlParserErrorsXMLSCHEMAPWARNSKIPSCHEMA                   XmlParserErrors = 3083
	XmlParserErrorsXMLSCHEMAPWARNUNLOCATEDSCHEMA              XmlParserErrors = 3084
	XmlParserErrorsXMLSCHEMAPWARNATTRREDECLPROH               XmlParserErrors = 3085
	XmlParserErrorsXMLSCHEMAPWARNATTRPOINTLESSPROH            XmlParserErrors = 3086
	XmlParserErrorsXMLSCHEMAPAGPROPSCORRECT                   XmlParserErrors = 3087
	XmlParserErrorsXMLSCHEMAPCOSCTEXTENDS12                   XmlParserErrors = 3088
	XmlParserErrorsXMLSCHEMAPAUPROPSCORRECT                   XmlParserErrors = 3089
	XmlParserErrorsXMLSCHEMAPAPROPSCORRECT3                   XmlParserErrors = 3090
	XmlParserErrorsXMLSCHEMAPCOSALLLIMITED                    XmlParserErrors = 3091
	XmlParserErrorsXMLSCHEMATRONVASSERT                       XmlParserErrors = 4000
	XmlParserErrorsXMLSCHEMATRONVREPORT                       XmlParserErrors = 4001
	XmlParserErrorsXMLMODULEOPEN                              XmlParserErrors = 4900
	XmlParserErrorsXMLMODULECLOSE                             XmlParserErrors = 4901
	XmlParserErrorsXMLCHECKFOUNDELEMENT                       XmlParserErrors = 5000
	XmlParserErrorsXMLCHECKFOUNDATTRIBUTE                     XmlParserErrors = 5001
	XmlParserErrorsXMLCHECKFOUNDTEXT                          XmlParserErrors = 5002
	XmlParserErrorsXMLCHECKFOUNDCDATA                         XmlParserErrors = 5003
	XmlParserErrorsXMLCHECKFOUNDENTITYREF                     XmlParserErrors = 5004
	XmlParserErrorsXMLCHECKFOUNDENTITY                        XmlParserErrors = 5005
	XmlParserErrorsXMLCHECKFOUNDPI                            XmlParserErrors = 5006
	XmlParserErrorsXMLCHECKFOUNDCOMMENT                       XmlParserErrors = 5007
	XmlParserErrorsXMLCHECKFOUNDDOCTYPE                       XmlParserErrors = 5008
	XmlParserErrorsXMLCHECKFOUNDFRAGMENT                      XmlParserErrors = 5009
	XmlParserErrorsXMLCHECKFOUNDNOTATION                      XmlParserErrors = 5010
	XmlParserErrorsXMLCHECKUNKNOWNNODE                        XmlParserErrors = 5011
	XmlParserErrorsXMLCHECKENTITYTYPE                         XmlParserErrors = 5012
	XmlParserErrorsXMLCHECKNOPARENT                           XmlParserErrors = 5013
	XmlParserErrorsXMLCHECKNODOC                              XmlParserErrors = 5014
	XmlParserErrorsXMLCHECKNONAME                             XmlParserErrors = 5015
	XmlParserErrorsXMLCHECKNOELEM                             XmlParserErrors = 5016
	XmlParserErrorsXMLCHECKWRONGDOC                           XmlParserErrors = 5017
	XmlParserErrorsXMLCHECKNOPREV                             XmlParserErrors = 5018
	XmlParserErrorsXMLCHECKWRONGPREV                          XmlParserErrors = 5019
	XmlParserErrorsXMLCHECKNONEXT                             XmlParserErrors = 5020
	XmlParserErrorsXMLCHECKWRONGNEXT                          XmlParserErrors = 5021
	XmlParserErrorsXMLCHECKNOTDTD                             XmlParserErrors = 5022
	XmlParserErrorsXMLCHECKNOTATTR                            XmlParserErrors = 5023
	XmlParserErrorsXMLCHECKNOTATTRDECL                        XmlParserErrors = 5024
	XmlParserErrorsXMLCHECKNOTELEMDECL                        XmlParserErrors = 5025
	XmlParserErrorsXMLCHECKNOTENTITYDECL                      XmlParserErrors = 5026
	XmlParserErrorsXMLCHECKNOTNSDECL                          XmlParserErrors = 5027
	XmlParserErrorsXMLCHECKNOHREF                             XmlParserErrors = 5028
	XmlParserErrorsXMLCHECKWRONGPARENT                        XmlParserErrors = 5029
	XmlParserErrorsXMLCHECKNSSCOPE                            XmlParserErrors = 5030
	XmlParserErrorsXMLCHECKNSANCESTOR                         XmlParserErrors = 5031
	XmlParserErrorsXMLCHECKNOTUTF8                            XmlParserErrors = 5032
	XmlParserErrorsXMLCHECKNODICT                             XmlParserErrors = 5033
	XmlParserErrorsXMLCHECKNOTNCNAME                          XmlParserErrors = 5034
	XmlParserErrorsXMLCHECKOUTSIDEDICT                        XmlParserErrors = 5035
	XmlParserErrorsXMLCHECKWRONGNAME                          XmlParserErrors = 5036
	XmlParserErrorsXMLCHECKNAMENOTNULL                        XmlParserErrors = 5037
	XmlParserErrorsXMLI18NNONAME                              XmlParserErrors = 6000
	XmlParserErrorsXMLI18NNOHANDLER                           XmlParserErrors = 6001
	XmlParserErrorsXMLI18NEXCESSHANDLER                       XmlParserErrors = 6002
	XmlParserErrorsXMLI18NCONVFAILED                          XmlParserErrors = 6003
	XmlParserErrorsXMLI18NNOOUTPUT                            XmlParserErrors = 6004
	XmlParserErrorsXMLBUFOVERFLOW                             XmlParserErrors = 7000
)
// llgo:type C
type XmlGenericErrorFunc func(__llgo_arg_0 unsafe.Pointer, __llgo_arg_1 *int8, __llgo_va_list ...interface{})
// llgo:type C
type XmlStructuredErrorFunc func(unsafe.Pointer, *XmlError)
//go:linkname X__XmlLastError C.__xmlLastError
func X__XmlLastError() *XmlError
//go:linkname X__XmlGenericError C.__xmlGenericError
func X__XmlGenericError() XmlGenericErrorFunc
//go:linkname X__XmlGenericErrorContext C.__xmlGenericErrorContext
func X__XmlGenericErrorContext() *unsafe.Pointer
//go:linkname X__XmlStructuredError C.__xmlStructuredError
func X__XmlStructuredError() XmlStructuredErrorFunc
//go:linkname X__XmlStructuredErrorContext C.__xmlStructuredErrorContext
func X__XmlStructuredErrorContext() *unsafe.Pointer
/** DOC_ENABLE */
//go:linkname XmlSetGenericErrorFunc C.xmlSetGenericErrorFunc
func XmlSetGenericErrorFunc(ctx unsafe.Pointer, handler XmlGenericErrorFunc)
//go:linkname XmlThrDefSetGenericErrorFunc C.xmlThrDefSetGenericErrorFunc
func XmlThrDefSetGenericErrorFunc(ctx unsafe.Pointer, handler XmlGenericErrorFunc)
//go:linkname InitGenericErrorDefaultFunc C.initGenericErrorDefaultFunc
func InitGenericErrorDefaultFunc(handler XmlGenericErrorFunc)
//go:linkname XmlSetStructuredErrorFunc C.xmlSetStructuredErrorFunc
func XmlSetStructuredErrorFunc(ctx unsafe.Pointer, handler XmlStructuredErrorFunc)
//go:linkname XmlThrDefSetStructuredErrorFunc C.xmlThrDefSetStructuredErrorFunc
func XmlThrDefSetStructuredErrorFunc(ctx unsafe.Pointer, handler XmlStructuredErrorFunc)
//go:linkname XmlParserError C.xmlParserError
func XmlParserError(ctx unsafe.Pointer, msg *int8, __llgo_va_list ...interface{})
//go:linkname XmlParserWarning C.xmlParserWarning
func XmlParserWarning(ctx unsafe.Pointer, msg *int8, __llgo_va_list ...interface{})
//go:linkname XmlParserValidityError C.xmlParserValidityError
func XmlParserValidityError(ctx unsafe.Pointer, msg *int8, __llgo_va_list ...interface{})
//go:linkname XmlParserValidityWarning C.xmlParserValidityWarning
func XmlParserValidityWarning(ctx unsafe.Pointer, msg *int8, __llgo_va_list ...interface{})
/** DOC_ENABLE */
// llgo:link (*X_XmlParserInput).XmlParserPrintFileInfo C.xmlParserPrintFileInfo
func (recv_ *X_XmlParserInput) XmlParserPrintFileInfo() {
}
// llgo:link (*X_XmlParserInput).XmlParserPrintFileContext C.xmlParserPrintFileContext
func (recv_ *X_XmlParserInput) XmlParserPrintFileContext() {
}
// llgo:link (*XmlError).XmlFormatError C.xmlFormatError
func (recv_ *XmlError) XmlFormatError(channel XmlGenericErrorFunc, data unsafe.Pointer) {
}
//go:linkname XmlGetLastError C.xmlGetLastError
func XmlGetLastError() *XmlError
//go:linkname XmlResetLastError C.xmlResetLastError
func XmlResetLastError()
//go:linkname XmlCtxtGetLastError C.xmlCtxtGetLastError
func XmlCtxtGetLastError(ctx unsafe.Pointer) *XmlError
//go:linkname XmlCtxtResetLastError C.xmlCtxtResetLastError
func XmlCtxtResetLastError(ctx unsafe.Pointer)
//go:linkname XmlResetError C.xmlResetError
func XmlResetError(err XmlErrorPtr)
// llgo:link (*XmlError).XmlCopyError C.xmlCopyError
func (recv_ *XmlError) XmlCopyError(to XmlErrorPtr) c.Int {
	return 0
}
