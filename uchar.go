package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const UUNICODEVERSION string = "16.0"
const UCHARMINVALUE c.Int = 0
const UCHARMAXVALUE c.Int = 1114111

type UProperty c.Int

const (
	UPropertyUCHARALPHABETIC                   UProperty = 0
	UPropertyUCHARBINARYSTART                  UProperty = 0
	UPropertyUCHARASCIIHEXDIGIT                UProperty = 1
	UPropertyUCHARBIDICONTROL                  UProperty = 2
	UPropertyUCHARBIDIMIRRORED                 UProperty = 3
	UPropertyUCHARDASH                         UProperty = 4
	UPropertyUCHARDEFAULTIGNORABLECODEPOINT    UProperty = 5
	UPropertyUCHARDEPRECATED                   UProperty = 6
	UPropertyUCHARDIACRITIC                    UProperty = 7
	UPropertyUCHAREXTENDER                     UProperty = 8
	UPropertyUCHARFULLCOMPOSITIONEXCLUSION     UProperty = 9
	UPropertyUCHARGRAPHEMEBASE                 UProperty = 10
	UPropertyUCHARGRAPHEMEEXTEND               UProperty = 11
	UPropertyUCHARGRAPHEMELINK                 UProperty = 12
	UPropertyUCHARHEXDIGIT                     UProperty = 13
	UPropertyUCHARHYPHEN                       UProperty = 14
	UPropertyUCHARIDCONTINUE                   UProperty = 15
	UPropertyUCHARIDSTART                      UProperty = 16
	UPropertyUCHARIDEOGRAPHIC                  UProperty = 17
	UPropertyUCHARIDSBINARYOPERATOR            UProperty = 18
	UPropertyUCHARIDSTRINARYOPERATOR           UProperty = 19
	UPropertyUCHARJOINCONTROL                  UProperty = 20
	UPropertyUCHARLOGICALORDEREXCEPTION        UProperty = 21
	UPropertyUCHARLOWERCASE                    UProperty = 22
	UPropertyUCHARMATH                         UProperty = 23
	UPropertyUCHARNONCHARACTERCODEPOINT        UProperty = 24
	UPropertyUCHARQUOTATIONMARK                UProperty = 25
	UPropertyUCHARRADICAL                      UProperty = 26
	UPropertyUCHARSOFTDOTTED                   UProperty = 27
	UPropertyUCHARTERMINALPUNCTUATION          UProperty = 28
	UPropertyUCHARUNIFIEDIDEOGRAPH             UProperty = 29
	UPropertyUCHARUPPERCASE                    UProperty = 30
	UPropertyUCHARWHITESPACE                   UProperty = 31
	UPropertyUCHARXIDCONTINUE                  UProperty = 32
	UPropertyUCHARXIDSTART                     UProperty = 33
	UPropertyUCHARCASESENSITIVE                UProperty = 34
	UPropertyUCHARSTERM                        UProperty = 35
	UPropertyUCHARVARIATIONSELECTOR            UProperty = 36
	UPropertyUCHARNFDINERT                     UProperty = 37
	UPropertyUCHARNFKDINERT                    UProperty = 38
	UPropertyUCHARNFCINERT                     UProperty = 39
	UPropertyUCHARNFKCINERT                    UProperty = 40
	UPropertyUCHARSEGMENTSTARTER               UProperty = 41
	UPropertyUCHARPATTERNSYNTAX                UProperty = 42
	UPropertyUCHARPATTERNWHITESPACE            UProperty = 43
	UPropertyUCHARPOSIXALNUM                   UProperty = 44
	UPropertyUCHARPOSIXBLANK                   UProperty = 45
	UPropertyUCHARPOSIXGRAPH                   UProperty = 46
	UPropertyUCHARPOSIXPRINT                   UProperty = 47
	UPropertyUCHARPOSIXXDIGIT                  UProperty = 48
	UPropertyUCHARCASED                        UProperty = 49
	UPropertyUCHARCASEIGNORABLE                UProperty = 50
	UPropertyUCHARCHANGESWHENLOWERCASED        UProperty = 51
	UPropertyUCHARCHANGESWHENUPPERCASED        UProperty = 52
	UPropertyUCHARCHANGESWHENTITLECASED        UProperty = 53
	UPropertyUCHARCHANGESWHENCASEFOLDED        UProperty = 54
	UPropertyUCHARCHANGESWHENCASEMAPPED        UProperty = 55
	UPropertyUCHARCHANGESWHENNFKCCASEFOLDED    UProperty = 56
	UPropertyUCHAREMOJI                        UProperty = 57
	UPropertyUCHAREMOJIPRESENTATION            UProperty = 58
	UPropertyUCHAREMOJIMODIFIER                UProperty = 59
	UPropertyUCHAREMOJIMODIFIERBASE            UProperty = 60
	UPropertyUCHAREMOJICOMPONENT               UProperty = 61
	UPropertyUCHARREGIONALINDICATOR            UProperty = 62
	UPropertyUCHARPREPENDEDCONCATENATIONMARK   UProperty = 63
	UPropertyUCHAREXTENDEDPICTOGRAPHIC         UProperty = 64
	UPropertyUCHARBASICEMOJI                   UProperty = 65
	UPropertyUCHAREMOJIKEYCAPSEQUENCE          UProperty = 66
	UPropertyUCHARRGIEMOJIMODIFIERSEQUENCE     UProperty = 67
	UPropertyUCHARRGIEMOJIFLAGSEQUENCE         UProperty = 68
	UPropertyUCHARRGIEMOJITAGSEQUENCE          UProperty = 69
	UPropertyUCHARRGIEMOJIZWJSEQUENCE          UProperty = 70
	UPropertyUCHARRGIEMOJI                     UProperty = 71
	UPropertyUCHARIDSUNARYOPERATOR             UProperty = 72
	UPropertyUCHARIDCOMPATMATHSTART            UProperty = 73
	UPropertyUCHARIDCOMPATMATHCONTINUE         UProperty = 74
	UPropertyUCHARMODIFIERCOMBININGMARK        UProperty = 75
	UPropertyUCHARBINARYLIMIT                  UProperty = 76
	UPropertyUCHARBIDICLASS                    UProperty = 4096
	UPropertyUCHARINTSTART                     UProperty = 4096
	UPropertyUCHARBLOCK                        UProperty = 4097
	UPropertyUCHARCANONICALCOMBININGCLASS      UProperty = 4098
	UPropertyUCHARDECOMPOSITIONTYPE            UProperty = 4099
	UPropertyUCHAREASTASIANWIDTH               UProperty = 4100
	UPropertyUCHARGENERALCATEGORY              UProperty = 4101
	UPropertyUCHARJOININGGROUP                 UProperty = 4102
	UPropertyUCHARJOININGTYPE                  UProperty = 4103
	UPropertyUCHARLINEBREAK                    UProperty = 4104
	UPropertyUCHARNUMERICTYPE                  UProperty = 4105
	UPropertyUCHARSCRIPT                       UProperty = 4106
	UPropertyUCHARHANGULSYLLABLETYPE           UProperty = 4107
	UPropertyUCHARNFDQUICKCHECK                UProperty = 4108
	UPropertyUCHARNFKDQUICKCHECK               UProperty = 4109
	UPropertyUCHARNFCQUICKCHECK                UProperty = 4110
	UPropertyUCHARNFKCQUICKCHECK               UProperty = 4111
	UPropertyUCHARLEADCANONICALCOMBININGCLASS  UProperty = 4112
	UPropertyUCHARTRAILCANONICALCOMBININGCLASS UProperty = 4113
	UPropertyUCHARGRAPHEMECLUSTERBREAK         UProperty = 4114
	UPropertyUCHARSENTENCEBREAK                UProperty = 4115
	UPropertyUCHARWORDBREAK                    UProperty = 4116
	UPropertyUCHARBIDIPAIREDBRACKETTYPE        UProperty = 4117
	UPropertyUCHARINDICPOSITIONALCATEGORY      UProperty = 4118
	UPropertyUCHARINDICSYLLABICCATEGORY        UProperty = 4119
	UPropertyUCHARVERTICALORIENTATION          UProperty = 4120
	UPropertyUCHARIDENTIFIERSTATUS             UProperty = 4121
	UPropertyUCHARINDICCONJUNCTBREAK           UProperty = 4122
	UPropertyUCHARINTLIMIT                     UProperty = 4123
	UPropertyUCHARGENERALCATEGORYMASK          UProperty = 8192
	UPropertyUCHARMASKSTART                    UProperty = 8192
	UPropertyUCHARMASKLIMIT                    UProperty = 8193
	UPropertyUCHARNUMERICVALUE                 UProperty = 12288
	UPropertyUCHARDOUBLESTART                  UProperty = 12288
	UPropertyUCHARDOUBLELIMIT                  UProperty = 12289
	UPropertyUCHARAGE                          UProperty = 16384
	UPropertyUCHARSTRINGSTART                  UProperty = 16384
	UPropertyUCHARBIDIMIRRORINGGLYPH           UProperty = 16385
	UPropertyUCHARCASEFOLDING                  UProperty = 16386
	UPropertyUCHARISOCOMMENT                   UProperty = 16387
	UPropertyUCHARLOWERCASEMAPPING             UProperty = 16388
	UPropertyUCHARNAME                         UProperty = 16389
	UPropertyUCHARSIMPLECASEFOLDING            UProperty = 16390
	UPropertyUCHARSIMPLELOWERCASEMAPPING       UProperty = 16391
	UPropertyUCHARSIMPLETITLECASEMAPPING       UProperty = 16392
	UPropertyUCHARSIMPLEUPPERCASEMAPPING       UProperty = 16393
	UPropertyUCHARTITLECASEMAPPING             UProperty = 16394
	UPropertyUCHARUNICODE1NAME                 UProperty = 16395
	UPropertyUCHARUPPERCASEMAPPING             UProperty = 16396
	UPropertyUCHARBIDIPAIREDBRACKET            UProperty = 16397
	UPropertyUCHARSTRINGLIMIT                  UProperty = 16398
	UPropertyUCHARSCRIPTEXTENSIONS             UProperty = 28672
	UPropertyUCHAROTHERPROPERTYSTART           UProperty = 28672
	UPropertyUCHARIDENTIFIERTYPE               UProperty = 28673
	UPropertyUCHAROTHERPROPERTYLIMIT           UProperty = 28674
	UPropertyUCHARINVALIDCODE                  UProperty = -1
)

type UCharCategory c.Int

const (
	UCharCategoryUUNASSIGNED           UCharCategory = 0
	UCharCategoryUGENERALOTHERTYPES    UCharCategory = 0
	UCharCategoryUUPPERCASELETTER      UCharCategory = 1
	UCharCategoryULOWERCASELETTER      UCharCategory = 2
	UCharCategoryUTITLECASELETTER      UCharCategory = 3
	UCharCategoryUMODIFIERLETTER       UCharCategory = 4
	UCharCategoryUOTHERLETTER          UCharCategory = 5
	UCharCategoryUNONSPACINGMARK       UCharCategory = 6
	UCharCategoryUENCLOSINGMARK        UCharCategory = 7
	UCharCategoryUCOMBININGSPACINGMARK UCharCategory = 8
	UCharCategoryUDECIMALDIGITNUMBER   UCharCategory = 9
	UCharCategoryULETTERNUMBER         UCharCategory = 10
	UCharCategoryUOTHERNUMBER          UCharCategory = 11
	UCharCategoryUSPACESEPARATOR       UCharCategory = 12
	UCharCategoryULINESEPARATOR        UCharCategory = 13
	UCharCategoryUPARAGRAPHSEPARATOR   UCharCategory = 14
	UCharCategoryUCONTROLCHAR          UCharCategory = 15
	UCharCategoryUFORMATCHAR           UCharCategory = 16
	UCharCategoryUPRIVATEUSECHAR       UCharCategory = 17
	UCharCategoryUSURROGATE            UCharCategory = 18
	UCharCategoryUDASHPUNCTUATION      UCharCategory = 19
	UCharCategoryUSTARTPUNCTUATION     UCharCategory = 20
	UCharCategoryUENDPUNCTUATION       UCharCategory = 21
	UCharCategoryUCONNECTORPUNCTUATION UCharCategory = 22
	UCharCategoryUOTHERPUNCTUATION     UCharCategory = 23
	UCharCategoryUMATHSYMBOL           UCharCategory = 24
	UCharCategoryUCURRENCYSYMBOL       UCharCategory = 25
	UCharCategoryUMODIFIERSYMBOL       UCharCategory = 26
	UCharCategoryUOTHERSYMBOL          UCharCategory = 27
	UCharCategoryUINITIALPUNCTUATION   UCharCategory = 28
	UCharCategoryUFINALPUNCTUATION     UCharCategory = 29
	UCharCategoryUCHARCATEGORYCOUNT    UCharCategory = 30
)

type UCharDirection c.Int

const (
	UCharDirectionULEFTTORIGHT              UCharDirection = 0
	UCharDirectionURIGHTTOLEFT              UCharDirection = 1
	UCharDirectionUEUROPEANNUMBER           UCharDirection = 2
	UCharDirectionUEUROPEANNUMBERSEPARATOR  UCharDirection = 3
	UCharDirectionUEUROPEANNUMBERTERMINATOR UCharDirection = 4
	UCharDirectionUARABICNUMBER             UCharDirection = 5
	UCharDirectionUCOMMONNUMBERSEPARATOR    UCharDirection = 6
	UCharDirectionUBLOCKSEPARATOR           UCharDirection = 7
	UCharDirectionUSEGMENTSEPARATOR         UCharDirection = 8
	UCharDirectionUWHITESPACENEUTRAL        UCharDirection = 9
	UCharDirectionUOTHERNEUTRAL             UCharDirection = 10
	UCharDirectionULEFTTORIGHTEMBEDDING     UCharDirection = 11
	UCharDirectionULEFTTORIGHTOVERRIDE      UCharDirection = 12
	UCharDirectionURIGHTTOLEFTARABIC        UCharDirection = 13
	UCharDirectionURIGHTTOLEFTEMBEDDING     UCharDirection = 14
	UCharDirectionURIGHTTOLEFTOVERRIDE      UCharDirection = 15
	UCharDirectionUPOPDIRECTIONALFORMAT     UCharDirection = 16
	UCharDirectionUDIRNONSPACINGMARK        UCharDirection = 17
	UCharDirectionUBOUNDARYNEUTRAL          UCharDirection = 18
	UCharDirectionUFIRSTSTRONGISOLATE       UCharDirection = 19
	UCharDirectionULEFTTORIGHTISOLATE       UCharDirection = 20
	UCharDirectionURIGHTTOLEFTISOLATE       UCharDirection = 21
	UCharDirectionUPOPDIRECTIONALISOLATE    UCharDirection = 22
	UCharDirectionUCHARDIRECTIONCOUNT       UCharDirection = 23
)

type UBidiPairedBracketType c.Int

const (
	UBidiPairedBracketTypeUBPTNONE  UBidiPairedBracketType = 0
	UBidiPairedBracketTypeUBPTOPEN  UBidiPairedBracketType = 1
	UBidiPairedBracketTypeUBPTCLOSE UBidiPairedBracketType = 2
	UBidiPairedBracketTypeUBPTCOUNT UBidiPairedBracketType = 3
)

type UBlockCode c.Int

const (
	UBlockCodeUBLOCKNOBLOCK                                     UBlockCode = 0
	UBlockCodeUBLOCKBASICLATIN                                  UBlockCode = 1
	UBlockCodeUBLOCKLATIN1SUPPLEMENT                            UBlockCode = 2
	UBlockCodeUBLOCKLATINEXTENDEDA                              UBlockCode = 3
	UBlockCodeUBLOCKLATINEXTENDEDB                              UBlockCode = 4
	UBlockCodeUBLOCKIPAEXTENSIONS                               UBlockCode = 5
	UBlockCodeUBLOCKSPACINGMODIFIERLETTERS                      UBlockCode = 6
	UBlockCodeUBLOCKCOMBININGDIACRITICALMARKS                   UBlockCode = 7
	UBlockCodeUBLOCKGREEK                                       UBlockCode = 8
	UBlockCodeUBLOCKCYRILLIC                                    UBlockCode = 9
	UBlockCodeUBLOCKARMENIAN                                    UBlockCode = 10
	UBlockCodeUBLOCKHEBREW                                      UBlockCode = 11
	UBlockCodeUBLOCKARABIC                                      UBlockCode = 12
	UBlockCodeUBLOCKSYRIAC                                      UBlockCode = 13
	UBlockCodeUBLOCKTHAANA                                      UBlockCode = 14
	UBlockCodeUBLOCKDEVANAGARI                                  UBlockCode = 15
	UBlockCodeUBLOCKBENGALI                                     UBlockCode = 16
	UBlockCodeUBLOCKGURMUKHI                                    UBlockCode = 17
	UBlockCodeUBLOCKGUJARATI                                    UBlockCode = 18
	UBlockCodeUBLOCKORIYA                                       UBlockCode = 19
	UBlockCodeUBLOCKTAMIL                                       UBlockCode = 20
	UBlockCodeUBLOCKTELUGU                                      UBlockCode = 21
	UBlockCodeUBLOCKKANNADA                                     UBlockCode = 22
	UBlockCodeUBLOCKMALAYALAM                                   UBlockCode = 23
	UBlockCodeUBLOCKSINHALA                                     UBlockCode = 24
	UBlockCodeUBLOCKTHAI                                        UBlockCode = 25
	UBlockCodeUBLOCKLAO                                         UBlockCode = 26
	UBlockCodeUBLOCKTIBETAN                                     UBlockCode = 27
	UBlockCodeUBLOCKMYANMAR                                     UBlockCode = 28
	UBlockCodeUBLOCKGEORGIAN                                    UBlockCode = 29
	UBlockCodeUBLOCKHANGULJAMO                                  UBlockCode = 30
	UBlockCodeUBLOCKETHIOPIC                                    UBlockCode = 31
	UBlockCodeUBLOCKCHEROKEE                                    UBlockCode = 32
	UBlockCodeUBLOCKUNIFIEDCANADIANABORIGINALSYLLABICS          UBlockCode = 33
	UBlockCodeUBLOCKOGHAM                                       UBlockCode = 34
	UBlockCodeUBLOCKRUNIC                                       UBlockCode = 35
	UBlockCodeUBLOCKKHMER                                       UBlockCode = 36
	UBlockCodeUBLOCKMONGOLIAN                                   UBlockCode = 37
	UBlockCodeUBLOCKLATINEXTENDEDADDITIONAL                     UBlockCode = 38
	UBlockCodeUBLOCKGREEKEXTENDED                               UBlockCode = 39
	UBlockCodeUBLOCKGENERALPUNCTUATION                          UBlockCode = 40
	UBlockCodeUBLOCKSUPERSCRIPTSANDSUBSCRIPTS                   UBlockCode = 41
	UBlockCodeUBLOCKCURRENCYSYMBOLS                             UBlockCode = 42
	UBlockCodeUBLOCKCOMBININGMARKSFORSYMBOLS                    UBlockCode = 43
	UBlockCodeUBLOCKLETTERLIKESYMBOLS                           UBlockCode = 44
	UBlockCodeUBLOCKNUMBERFORMS                                 UBlockCode = 45
	UBlockCodeUBLOCKARROWS                                      UBlockCode = 46
	UBlockCodeUBLOCKMATHEMATICALOPERATORS                       UBlockCode = 47
	UBlockCodeUBLOCKMISCELLANEOUSTECHNICAL                      UBlockCode = 48
	UBlockCodeUBLOCKCONTROLPICTURES                             UBlockCode = 49
	UBlockCodeUBLOCKOPTICALCHARACTERRECOGNITION                 UBlockCode = 50
	UBlockCodeUBLOCKENCLOSEDALPHANUMERICS                       UBlockCode = 51
	UBlockCodeUBLOCKBOXDRAWING                                  UBlockCode = 52
	UBlockCodeUBLOCKBLOCKELEMENTS                               UBlockCode = 53
	UBlockCodeUBLOCKGEOMETRICSHAPES                             UBlockCode = 54
	UBlockCodeUBLOCKMISCELLANEOUSSYMBOLS                        UBlockCode = 55
	UBlockCodeUBLOCKDINGBATS                                    UBlockCode = 56
	UBlockCodeUBLOCKBRAILLEPATTERNS                             UBlockCode = 57
	UBlockCodeUBLOCKCJKRADICALSSUPPLEMENT                       UBlockCode = 58
	UBlockCodeUBLOCKKANGXIRADICALS                              UBlockCode = 59
	UBlockCodeUBLOCKIDEOGRAPHICDESCRIPTIONCHARACTERS            UBlockCode = 60
	UBlockCodeUBLOCKCJKSYMBOLSANDPUNCTUATION                    UBlockCode = 61
	UBlockCodeUBLOCKHIRAGANA                                    UBlockCode = 62
	UBlockCodeUBLOCKKATAKANA                                    UBlockCode = 63
	UBlockCodeUBLOCKBOPOMOFO                                    UBlockCode = 64
	UBlockCodeUBLOCKHANGULCOMPATIBILITYJAMO                     UBlockCode = 65
	UBlockCodeUBLOCKKANBUN                                      UBlockCode = 66
	UBlockCodeUBLOCKBOPOMOFOEXTENDED                            UBlockCode = 67
	UBlockCodeUBLOCKENCLOSEDCJKLETTERSANDMONTHS                 UBlockCode = 68
	UBlockCodeUBLOCKCJKCOMPATIBILITY                            UBlockCode = 69
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHSEXTENSIONA              UBlockCode = 70
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHS                        UBlockCode = 71
	UBlockCodeUBLOCKYISYLLABLES                                 UBlockCode = 72
	UBlockCodeUBLOCKYIRADICALS                                  UBlockCode = 73
	UBlockCodeUBLOCKHANGULSYLLABLES                             UBlockCode = 74
	UBlockCodeUBLOCKHIGHSURROGATES                              UBlockCode = 75
	UBlockCodeUBLOCKHIGHPRIVATEUSESURROGATES                    UBlockCode = 76
	UBlockCodeUBLOCKLOWSURROGATES                               UBlockCode = 77
	UBlockCodeUBLOCKPRIVATEUSEAREA                              UBlockCode = 78
	UBlockCodeUBLOCKPRIVATEUSE                                  UBlockCode = 78
	UBlockCodeUBLOCKCJKCOMPATIBILITYIDEOGRAPHS                  UBlockCode = 79
	UBlockCodeUBLOCKALPHABETICPRESENTATIONFORMS                 UBlockCode = 80
	UBlockCodeUBLOCKARABICPRESENTATIONFORMSA                    UBlockCode = 81
	UBlockCodeUBLOCKCOMBININGHALFMARKS                          UBlockCode = 82
	UBlockCodeUBLOCKCJKCOMPATIBILITYFORMS                       UBlockCode = 83
	UBlockCodeUBLOCKSMALLFORMVARIANTS                           UBlockCode = 84
	UBlockCodeUBLOCKARABICPRESENTATIONFORMSB                    UBlockCode = 85
	UBlockCodeUBLOCKSPECIALS                                    UBlockCode = 86
	UBlockCodeUBLOCKHALFWIDTHANDFULLWIDTHFORMS                  UBlockCode = 87
	UBlockCodeUBLOCKOLDITALIC                                   UBlockCode = 88
	UBlockCodeUBLOCKGOTHIC                                      UBlockCode = 89
	UBlockCodeUBLOCKDESERET                                     UBlockCode = 90
	UBlockCodeUBLOCKBYZANTINEMUSICALSYMBOLS                     UBlockCode = 91
	UBlockCodeUBLOCKMUSICALSYMBOLS                              UBlockCode = 92
	UBlockCodeUBLOCKMATHEMATICALALPHANUMERICSYMBOLS             UBlockCode = 93
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHSEXTENSIONB              UBlockCode = 94
	UBlockCodeUBLOCKCJKCOMPATIBILITYIDEOGRAPHSSUPPLEMENT        UBlockCode = 95
	UBlockCodeUBLOCKTAGS                                        UBlockCode = 96
	UBlockCodeUBLOCKCYRILLICSUPPLEMENT                          UBlockCode = 97
	UBlockCodeUBLOCKCYRILLICSUPPLEMENTARY                       UBlockCode = 97
	UBlockCodeUBLOCKTAGALOG                                     UBlockCode = 98
	UBlockCodeUBLOCKHANUNOO                                     UBlockCode = 99
	UBlockCodeUBLOCKBUHID                                       UBlockCode = 100
	UBlockCodeUBLOCKTAGBANWA                                    UBlockCode = 101
	UBlockCodeUBLOCKMISCELLANEOUSMATHEMATICALSYMBOLSA           UBlockCode = 102
	UBlockCodeUBLOCKSUPPLEMENTALARROWSA                         UBlockCode = 103
	UBlockCodeUBLOCKSUPPLEMENTALARROWSB                         UBlockCode = 104
	UBlockCodeUBLOCKMISCELLANEOUSMATHEMATICALSYMBOLSB           UBlockCode = 105
	UBlockCodeUBLOCKSUPPLEMENTALMATHEMATICALOPERATORS           UBlockCode = 106
	UBlockCodeUBLOCKKATAKANAPHONETICEXTENSIONS                  UBlockCode = 107
	UBlockCodeUBLOCKVARIATIONSELECTORS                          UBlockCode = 108
	UBlockCodeUBLOCKSUPPLEMENTARYPRIVATEUSEAREAA                UBlockCode = 109
	UBlockCodeUBLOCKSUPPLEMENTARYPRIVATEUSEAREAB                UBlockCode = 110
	UBlockCodeUBLOCKLIMBU                                       UBlockCode = 111
	UBlockCodeUBLOCKTAILE                                       UBlockCode = 112
	UBlockCodeUBLOCKKHMERSYMBOLS                                UBlockCode = 113
	UBlockCodeUBLOCKPHONETICEXTENSIONS                          UBlockCode = 114
	UBlockCodeUBLOCKMISCELLANEOUSSYMBOLSANDARROWS               UBlockCode = 115
	UBlockCodeUBLOCKYIJINGHEXAGRAMSYMBOLS                       UBlockCode = 116
	UBlockCodeUBLOCKLINEARBSYLLABARY                            UBlockCode = 117
	UBlockCodeUBLOCKLINEARBIDEOGRAMS                            UBlockCode = 118
	UBlockCodeUBLOCKAEGEANNUMBERS                               UBlockCode = 119
	UBlockCodeUBLOCKUGARITIC                                    UBlockCode = 120
	UBlockCodeUBLOCKSHAVIAN                                     UBlockCode = 121
	UBlockCodeUBLOCKOSMANYA                                     UBlockCode = 122
	UBlockCodeUBLOCKCYPRIOTSYLLABARY                            UBlockCode = 123
	UBlockCodeUBLOCKTAIXUANJINGSYMBOLS                          UBlockCode = 124
	UBlockCodeUBLOCKVARIATIONSELECTORSSUPPLEMENT                UBlockCode = 125
	UBlockCodeUBLOCKANCIENTGREEKMUSICALNOTATION                 UBlockCode = 126
	UBlockCodeUBLOCKANCIENTGREEKNUMBERS                         UBlockCode = 127
	UBlockCodeUBLOCKARABICSUPPLEMENT                            UBlockCode = 128
	UBlockCodeUBLOCKBUGINESE                                    UBlockCode = 129
	UBlockCodeUBLOCKCJKSTROKES                                  UBlockCode = 130
	UBlockCodeUBLOCKCOMBININGDIACRITICALMARKSSUPPLEMENT         UBlockCode = 131
	UBlockCodeUBLOCKCOPTIC                                      UBlockCode = 132
	UBlockCodeUBLOCKETHIOPICEXTENDED                            UBlockCode = 133
	UBlockCodeUBLOCKETHIOPICSUPPLEMENT                          UBlockCode = 134
	UBlockCodeUBLOCKGEORGIANSUPPLEMENT                          UBlockCode = 135
	UBlockCodeUBLOCKGLAGOLITIC                                  UBlockCode = 136
	UBlockCodeUBLOCKKHAROSHTHI                                  UBlockCode = 137
	UBlockCodeUBLOCKMODIFIERTONELETTERS                         UBlockCode = 138
	UBlockCodeUBLOCKNEWTAILUE                                   UBlockCode = 139
	UBlockCodeUBLOCKOLDPERSIAN                                  UBlockCode = 140
	UBlockCodeUBLOCKPHONETICEXTENSIONSSUPPLEMENT                UBlockCode = 141
	UBlockCodeUBLOCKSUPPLEMENTALPUNCTUATION                     UBlockCode = 142
	UBlockCodeUBLOCKSYLOTINAGRI                                 UBlockCode = 143
	UBlockCodeUBLOCKTIFINAGH                                    UBlockCode = 144
	UBlockCodeUBLOCKVERTICALFORMS                               UBlockCode = 145
	UBlockCodeUBLOCKNKO                                         UBlockCode = 146
	UBlockCodeUBLOCKBALINESE                                    UBlockCode = 147
	UBlockCodeUBLOCKLATINEXTENDEDC                              UBlockCode = 148
	UBlockCodeUBLOCKLATINEXTENDEDD                              UBlockCode = 149
	UBlockCodeUBLOCKPHAGSPA                                     UBlockCode = 150
	UBlockCodeUBLOCKPHOENICIAN                                  UBlockCode = 151
	UBlockCodeUBLOCKCUNEIFORM                                   UBlockCode = 152
	UBlockCodeUBLOCKCUNEIFORMNUMBERSANDPUNCTUATION              UBlockCode = 153
	UBlockCodeUBLOCKCOUNTINGRODNUMERALS                         UBlockCode = 154
	UBlockCodeUBLOCKSUNDANESE                                   UBlockCode = 155
	UBlockCodeUBLOCKLEPCHA                                      UBlockCode = 156
	UBlockCodeUBLOCKOLCHIKI                                     UBlockCode = 157
	UBlockCodeUBLOCKCYRILLICEXTENDEDA                           UBlockCode = 158
	UBlockCodeUBLOCKVAI                                         UBlockCode = 159
	UBlockCodeUBLOCKCYRILLICEXTENDEDB                           UBlockCode = 160
	UBlockCodeUBLOCKSAURASHTRA                                  UBlockCode = 161
	UBlockCodeUBLOCKKAYAHLI                                     UBlockCode = 162
	UBlockCodeUBLOCKREJANG                                      UBlockCode = 163
	UBlockCodeUBLOCKCHAM                                        UBlockCode = 164
	UBlockCodeUBLOCKANCIENTSYMBOLS                              UBlockCode = 165
	UBlockCodeUBLOCKPHAISTOSDISC                                UBlockCode = 166
	UBlockCodeUBLOCKLYCIAN                                      UBlockCode = 167
	UBlockCodeUBLOCKCARIAN                                      UBlockCode = 168
	UBlockCodeUBLOCKLYDIAN                                      UBlockCode = 169
	UBlockCodeUBLOCKMAHJONGTILES                                UBlockCode = 170
	UBlockCodeUBLOCKDOMINOTILES                                 UBlockCode = 171
	UBlockCodeUBLOCKSAMARITAN                                   UBlockCode = 172
	UBlockCodeUBLOCKUNIFIEDCANADIANABORIGINALSYLLABICSEXTENDED  UBlockCode = 173
	UBlockCodeUBLOCKTAITHAM                                     UBlockCode = 174
	UBlockCodeUBLOCKVEDICEXTENSIONS                             UBlockCode = 175
	UBlockCodeUBLOCKLISU                                        UBlockCode = 176
	UBlockCodeUBLOCKBAMUM                                       UBlockCode = 177
	UBlockCodeUBLOCKCOMMONINDICNUMBERFORMS                      UBlockCode = 178
	UBlockCodeUBLOCKDEVANAGARIEXTENDED                          UBlockCode = 179
	UBlockCodeUBLOCKHANGULJAMOEXTENDEDA                         UBlockCode = 180
	UBlockCodeUBLOCKJAVANESE                                    UBlockCode = 181
	UBlockCodeUBLOCKMYANMAREXTENDEDA                            UBlockCode = 182
	UBlockCodeUBLOCKTAIVIET                                     UBlockCode = 183
	UBlockCodeUBLOCKMEETEIMAYEK                                 UBlockCode = 184
	UBlockCodeUBLOCKHANGULJAMOEXTENDEDB                         UBlockCode = 185
	UBlockCodeUBLOCKIMPERIALARAMAIC                             UBlockCode = 186
	UBlockCodeUBLOCKOLDSOUTHARABIAN                             UBlockCode = 187
	UBlockCodeUBLOCKAVESTAN                                     UBlockCode = 188
	UBlockCodeUBLOCKINSCRIPTIONALPARTHIAN                       UBlockCode = 189
	UBlockCodeUBLOCKINSCRIPTIONALPAHLAVI                        UBlockCode = 190
	UBlockCodeUBLOCKOLDTURKIC                                   UBlockCode = 191
	UBlockCodeUBLOCKRUMINUMERALSYMBOLS                          UBlockCode = 192
	UBlockCodeUBLOCKKAITHI                                      UBlockCode = 193
	UBlockCodeUBLOCKEGYPTIANHIEROGLYPHS                         UBlockCode = 194
	UBlockCodeUBLOCKENCLOSEDALPHANUMERICSUPPLEMENT              UBlockCode = 195
	UBlockCodeUBLOCKENCLOSEDIDEOGRAPHICSUPPLEMENT               UBlockCode = 196
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHSEXTENSIONC              UBlockCode = 197
	UBlockCodeUBLOCKMANDAIC                                     UBlockCode = 198
	UBlockCodeUBLOCKBATAK                                       UBlockCode = 199
	UBlockCodeUBLOCKETHIOPICEXTENDEDA                           UBlockCode = 200
	UBlockCodeUBLOCKBRAHMI                                      UBlockCode = 201
	UBlockCodeUBLOCKBAMUMSUPPLEMENT                             UBlockCode = 202
	UBlockCodeUBLOCKKANASUPPLEMENT                              UBlockCode = 203
	UBlockCodeUBLOCKPLAYINGCARDS                                UBlockCode = 204
	UBlockCodeUBLOCKMISCELLANEOUSSYMBOLSANDPICTOGRAPHS          UBlockCode = 205
	UBlockCodeUBLOCKEMOTICONS                                   UBlockCode = 206
	UBlockCodeUBLOCKTRANSPORTANDMAPSYMBOLS                      UBlockCode = 207
	UBlockCodeUBLOCKALCHEMICALSYMBOLS                           UBlockCode = 208
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHSEXTENSIOND              UBlockCode = 209
	UBlockCodeUBLOCKARABICEXTENDEDA                             UBlockCode = 210
	UBlockCodeUBLOCKARABICMATHEMATICALALPHABETICSYMBOLS         UBlockCode = 211
	UBlockCodeUBLOCKCHAKMA                                      UBlockCode = 212
	UBlockCodeUBLOCKMEETEIMAYEKEXTENSIONS                       UBlockCode = 213
	UBlockCodeUBLOCKMEROITICCURSIVE                             UBlockCode = 214
	UBlockCodeUBLOCKMEROITICHIEROGLYPHS                         UBlockCode = 215
	UBlockCodeUBLOCKMIAO                                        UBlockCode = 216
	UBlockCodeUBLOCKSHARADA                                     UBlockCode = 217
	UBlockCodeUBLOCKSORASOMPENG                                 UBlockCode = 218
	UBlockCodeUBLOCKSUNDANESESUPPLEMENT                         UBlockCode = 219
	UBlockCodeUBLOCKTAKRI                                       UBlockCode = 220
	UBlockCodeUBLOCKBASSAVAH                                    UBlockCode = 221
	UBlockCodeUBLOCKCAUCASIANALBANIAN                           UBlockCode = 222
	UBlockCodeUBLOCKCOPTICEPACTNUMBERS                          UBlockCode = 223
	UBlockCodeUBLOCKCOMBININGDIACRITICALMARKSEXTENDED           UBlockCode = 224
	UBlockCodeUBLOCKDUPLOYAN                                    UBlockCode = 225
	UBlockCodeUBLOCKELBASAN                                     UBlockCode = 226
	UBlockCodeUBLOCKGEOMETRICSHAPESEXTENDED                     UBlockCode = 227
	UBlockCodeUBLOCKGRANTHA                                     UBlockCode = 228
	UBlockCodeUBLOCKKHOJKI                                      UBlockCode = 229
	UBlockCodeUBLOCKKHUDAWADI                                   UBlockCode = 230
	UBlockCodeUBLOCKLATINEXTENDEDE                              UBlockCode = 231
	UBlockCodeUBLOCKLINEARA                                     UBlockCode = 232
	UBlockCodeUBLOCKMAHAJANI                                    UBlockCode = 233
	UBlockCodeUBLOCKMANICHAEAN                                  UBlockCode = 234
	UBlockCodeUBLOCKMENDEKIKAKUI                                UBlockCode = 235
	UBlockCodeUBLOCKMODI                                        UBlockCode = 236
	UBlockCodeUBLOCKMRO                                         UBlockCode = 237
	UBlockCodeUBLOCKMYANMAREXTENDEDB                            UBlockCode = 238
	UBlockCodeUBLOCKNABATAEAN                                   UBlockCode = 239
	UBlockCodeUBLOCKOLDNORTHARABIAN                             UBlockCode = 240
	UBlockCodeUBLOCKOLDPERMIC                                   UBlockCode = 241
	UBlockCodeUBLOCKORNAMENTALDINGBATS                          UBlockCode = 242
	UBlockCodeUBLOCKPAHAWHHMONG                                 UBlockCode = 243
	UBlockCodeUBLOCKPALMYRENE                                   UBlockCode = 244
	UBlockCodeUBLOCKPAUCINHAU                                   UBlockCode = 245
	UBlockCodeUBLOCKPSALTERPAHLAVI                              UBlockCode = 246
	UBlockCodeUBLOCKSHORTHANDFORMATCONTROLS                     UBlockCode = 247
	UBlockCodeUBLOCKSIDDHAM                                     UBlockCode = 248
	UBlockCodeUBLOCKSINHALAARCHAICNUMBERS                       UBlockCode = 249
	UBlockCodeUBLOCKSUPPLEMENTALARROWSC                         UBlockCode = 250
	UBlockCodeUBLOCKTIRHUTA                                     UBlockCode = 251
	UBlockCodeUBLOCKWARANGCITI                                  UBlockCode = 252
	UBlockCodeUBLOCKAHOM                                        UBlockCode = 253
	UBlockCodeUBLOCKANATOLIANHIEROGLYPHS                        UBlockCode = 254
	UBlockCodeUBLOCKCHEROKEESUPPLEMENT                          UBlockCode = 255
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHSEXTENSIONE              UBlockCode = 256
	UBlockCodeUBLOCKEARLYDYNASTICCUNEIFORM                      UBlockCode = 257
	UBlockCodeUBLOCKHATRAN                                      UBlockCode = 258
	UBlockCodeUBLOCKMULTANI                                     UBlockCode = 259
	UBlockCodeUBLOCKOLDHUNGARIAN                                UBlockCode = 260
	UBlockCodeUBLOCKSUPPLEMENTALSYMBOLSANDPICTOGRAPHS           UBlockCode = 261
	UBlockCodeUBLOCKSUTTONSIGNWRITING                           UBlockCode = 262
	UBlockCodeUBLOCKADLAM                                       UBlockCode = 263
	UBlockCodeUBLOCKBHAIKSUKI                                   UBlockCode = 264
	UBlockCodeUBLOCKCYRILLICEXTENDEDC                           UBlockCode = 265
	UBlockCodeUBLOCKGLAGOLITICSUPPLEMENT                        UBlockCode = 266
	UBlockCodeUBLOCKIDEOGRAPHICSYMBOLSANDPUNCTUATION            UBlockCode = 267
	UBlockCodeUBLOCKMARCHEN                                     UBlockCode = 268
	UBlockCodeUBLOCKMONGOLIANSUPPLEMENT                         UBlockCode = 269
	UBlockCodeUBLOCKNEWA                                        UBlockCode = 270
	UBlockCodeUBLOCKOSAGE                                       UBlockCode = 271
	UBlockCodeUBLOCKTANGUT                                      UBlockCode = 272
	UBlockCodeUBLOCKTANGUTCOMPONENTS                            UBlockCode = 273
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHSEXTENSIONF              UBlockCode = 274
	UBlockCodeUBLOCKKANAEXTENDEDA                               UBlockCode = 275
	UBlockCodeUBLOCKMASARAMGONDI                                UBlockCode = 276
	UBlockCodeUBLOCKNUSHU                                       UBlockCode = 277
	UBlockCodeUBLOCKSOYOMBO                                     UBlockCode = 278
	UBlockCodeUBLOCKSYRIACSUPPLEMENT                            UBlockCode = 279
	UBlockCodeUBLOCKZANABAZARSQUARE                             UBlockCode = 280
	UBlockCodeUBLOCKCHESSSYMBOLS                                UBlockCode = 281
	UBlockCodeUBLOCKDOGRA                                       UBlockCode = 282
	UBlockCodeUBLOCKGEORGIANEXTENDED                            UBlockCode = 283
	UBlockCodeUBLOCKGUNJALAGONDI                                UBlockCode = 284
	UBlockCodeUBLOCKHANIFIROHINGYA                              UBlockCode = 285
	UBlockCodeUBLOCKINDICSIYAQNUMBERS                           UBlockCode = 286
	UBlockCodeUBLOCKMAKASAR                                     UBlockCode = 287
	UBlockCodeUBLOCKMAYANNUMERALS                               UBlockCode = 288
	UBlockCodeUBLOCKMEDEFAIDRIN                                 UBlockCode = 289
	UBlockCodeUBLOCKOLDSOGDIAN                                  UBlockCode = 290
	UBlockCodeUBLOCKSOGDIAN                                     UBlockCode = 291
	UBlockCodeUBLOCKEGYPTIANHIEROGLYPHFORMATCONTROLS            UBlockCode = 292
	UBlockCodeUBLOCKELYMAIC                                     UBlockCode = 293
	UBlockCodeUBLOCKNANDINAGARI                                 UBlockCode = 294
	UBlockCodeUBLOCKNYIAKENGPUACHUEHMONG                        UBlockCode = 295
	UBlockCodeUBLOCKOTTOMANSIYAQNUMBERS                         UBlockCode = 296
	UBlockCodeUBLOCKSMALLKANAEXTENSION                          UBlockCode = 297
	UBlockCodeUBLOCKSYMBOLSANDPICTOGRAPHSEXTENDEDA              UBlockCode = 298
	UBlockCodeUBLOCKTAMILSUPPLEMENT                             UBlockCode = 299
	UBlockCodeUBLOCKWANCHO                                      UBlockCode = 300
	UBlockCodeUBLOCKCHORASMIAN                                  UBlockCode = 301
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHSEXTENSIONG              UBlockCode = 302
	UBlockCodeUBLOCKDIVESAKURU                                  UBlockCode = 303
	UBlockCodeUBLOCKKHITANSMALLSCRIPT                           UBlockCode = 304
	UBlockCodeUBLOCKLISUSUPPLEMENT                              UBlockCode = 305
	UBlockCodeUBLOCKSYMBOLSFORLEGACYCOMPUTING                   UBlockCode = 306
	UBlockCodeUBLOCKTANGUTSUPPLEMENT                            UBlockCode = 307
	UBlockCodeUBLOCKYEZIDI                                      UBlockCode = 308
	UBlockCodeUBLOCKARABICEXTENDEDB                             UBlockCode = 309
	UBlockCodeUBLOCKCYPROMINOAN                                 UBlockCode = 310
	UBlockCodeUBLOCKETHIOPICEXTENDEDB                           UBlockCode = 311
	UBlockCodeUBLOCKKANAEXTENDEDB                               UBlockCode = 312
	UBlockCodeUBLOCKLATINEXTENDEDF                              UBlockCode = 313
	UBlockCodeUBLOCKLATINEXTENDEDG                              UBlockCode = 314
	UBlockCodeUBLOCKOLDUYGHUR                                   UBlockCode = 315
	UBlockCodeUBLOCKTANGSA                                      UBlockCode = 316
	UBlockCodeUBLOCKTOTO                                        UBlockCode = 317
	UBlockCodeUBLOCKUNIFIEDCANADIANABORIGINALSYLLABICSEXTENDEDA UBlockCode = 318
	UBlockCodeUBLOCKVITHKUQI                                    UBlockCode = 319
	UBlockCodeUBLOCKZNAMENNYMUSICALNOTATION                     UBlockCode = 320
	UBlockCodeUBLOCKARABICEXTENDEDC                             UBlockCode = 321
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHSEXTENSIONH              UBlockCode = 322
	UBlockCodeUBLOCKCYRILLICEXTENDEDD                           UBlockCode = 323
	UBlockCodeUBLOCKDEVANAGARIEXTENDEDA                         UBlockCode = 324
	UBlockCodeUBLOCKKAKTOVIKNUMERALS                            UBlockCode = 325
	UBlockCodeUBLOCKKAWI                                        UBlockCode = 326
	UBlockCodeUBLOCKNAGMUNDARI                                  UBlockCode = 327
	UBlockCodeUBLOCKCJKUNIFIEDIDEOGRAPHSEXTENSIONI              UBlockCode = 328
	UBlockCodeUBLOCKEGYPTIANHIEROGLYPHSEXTENDEDA                UBlockCode = 329
	UBlockCodeUBLOCKGARAY                                       UBlockCode = 330
	UBlockCodeUBLOCKGURUNGKHEMA                                 UBlockCode = 331
	UBlockCodeUBLOCKKIRATRAI                                    UBlockCode = 332
	UBlockCodeUBLOCKMYANMAREXTENDEDC                            UBlockCode = 333
	UBlockCodeUBLOCKOLONAL                                      UBlockCode = 334
	UBlockCodeUBLOCKSUNUWAR                                     UBlockCode = 335
	UBlockCodeUBLOCKSYMBOLSFORLEGACYCOMPUTINGSUPPLEMENT         UBlockCode = 336
	UBlockCodeUBLOCKTODHRI                                      UBlockCode = 337
	UBlockCodeUBLOCKTULUTIGALARI                                UBlockCode = 338
	UBlockCodeUBLOCKCOUNT                                       UBlockCode = 339
	UBlockCodeUBLOCKINVALIDCODE                                 UBlockCode = -1
)

type UEastAsianWidth c.Int

const (
	UEastAsianWidthUEANEUTRAL   UEastAsianWidth = 0
	UEastAsianWidthUEAAMBIGUOUS UEastAsianWidth = 1
	UEastAsianWidthUEAHALFWIDTH UEastAsianWidth = 2
	UEastAsianWidthUEAFULLWIDTH UEastAsianWidth = 3
	UEastAsianWidthUEANARROW    UEastAsianWidth = 4
	UEastAsianWidthUEAWIDE      UEastAsianWidth = 5
	UEastAsianWidthUEACOUNT     UEastAsianWidth = 6
)

type UCharNameChoice c.Int

const (
	UCharNameChoiceUUNICODECHARNAME     UCharNameChoice = 0
	UCharNameChoiceUUNICODE10CHARNAME   UCharNameChoice = 1
	UCharNameChoiceUEXTENDEDCHARNAME    UCharNameChoice = 2
	UCharNameChoiceUCHARNAMEALIAS       UCharNameChoice = 3
	UCharNameChoiceUCHARNAMECHOICECOUNT UCharNameChoice = 4
)

type UPropertyNameChoice c.Int

const (
	UPropertyNameChoiceUSHORTPROPERTYNAME       UPropertyNameChoice = 0
	UPropertyNameChoiceULONGPROPERTYNAME        UPropertyNameChoice = 1
	UPropertyNameChoiceUPROPERTYNAMECHOICECOUNT UPropertyNameChoice = 2
)

type UDecompositionType c.Int

const (
	UDecompositionTypeUDTNONE      UDecompositionType = 0
	UDecompositionTypeUDTCANONICAL UDecompositionType = 1
	UDecompositionTypeUDTCOMPAT    UDecompositionType = 2
	UDecompositionTypeUDTCIRCLE    UDecompositionType = 3
	UDecompositionTypeUDTFINAL     UDecompositionType = 4
	UDecompositionTypeUDTFONT      UDecompositionType = 5
	UDecompositionTypeUDTFRACTION  UDecompositionType = 6
	UDecompositionTypeUDTINITIAL   UDecompositionType = 7
	UDecompositionTypeUDTISOLATED  UDecompositionType = 8
	UDecompositionTypeUDTMEDIAL    UDecompositionType = 9
	UDecompositionTypeUDTNARROW    UDecompositionType = 10
	UDecompositionTypeUDTNOBREAK   UDecompositionType = 11
	UDecompositionTypeUDTSMALL     UDecompositionType = 12
	UDecompositionTypeUDTSQUARE    UDecompositionType = 13
	UDecompositionTypeUDTSUB       UDecompositionType = 14
	UDecompositionTypeUDTSUPER     UDecompositionType = 15
	UDecompositionTypeUDTVERTICAL  UDecompositionType = 16
	UDecompositionTypeUDTWIDE      UDecompositionType = 17
	UDecompositionTypeUDTCOUNT     UDecompositionType = 18
)

type UJoiningType c.Int

const (
	UJoiningTypeUJTNONJOINING   UJoiningType = 0
	UJoiningTypeUJTJOINCAUSING  UJoiningType = 1
	UJoiningTypeUJTDUALJOINING  UJoiningType = 2
	UJoiningTypeUJTLEFTJOINING  UJoiningType = 3
	UJoiningTypeUJTRIGHTJOINING UJoiningType = 4
	UJoiningTypeUJTTRANSPARENT  UJoiningType = 5
	UJoiningTypeUJTCOUNT        UJoiningType = 6
)

type UJoiningGroup c.Int

const (
	UJoiningGroupUJGNOJOININGGROUP        UJoiningGroup = 0
	UJoiningGroupUJGAIN                   UJoiningGroup = 1
	UJoiningGroupUJGALAPH                 UJoiningGroup = 2
	UJoiningGroupUJGALEF                  UJoiningGroup = 3
	UJoiningGroupUJGBEH                   UJoiningGroup = 4
	UJoiningGroupUJGBETH                  UJoiningGroup = 5
	UJoiningGroupUJGDAL                   UJoiningGroup = 6
	UJoiningGroupUJGDALATHRISH            UJoiningGroup = 7
	UJoiningGroupUJGE                     UJoiningGroup = 8
	UJoiningGroupUJGFEH                   UJoiningGroup = 9
	UJoiningGroupUJGFINALSEMKATH          UJoiningGroup = 10
	UJoiningGroupUJGGAF                   UJoiningGroup = 11
	UJoiningGroupUJGGAMAL                 UJoiningGroup = 12
	UJoiningGroupUJGHAH                   UJoiningGroup = 13
	UJoiningGroupUJGTEHMARBUTAGOAL        UJoiningGroup = 14
	UJoiningGroupUJGHAMZAONHEHGOAL        UJoiningGroup = 14
	UJoiningGroupUJGHE                    UJoiningGroup = 15
	UJoiningGroupUJGHEH                   UJoiningGroup = 16
	UJoiningGroupUJGHEHGOAL               UJoiningGroup = 17
	UJoiningGroupUJGHETH                  UJoiningGroup = 18
	UJoiningGroupUJGKAF                   UJoiningGroup = 19
	UJoiningGroupUJGKAPH                  UJoiningGroup = 20
	UJoiningGroupUJGKNOTTEDHEH            UJoiningGroup = 21
	UJoiningGroupUJGLAM                   UJoiningGroup = 22
	UJoiningGroupUJGLAMADH                UJoiningGroup = 23
	UJoiningGroupUJGMEEM                  UJoiningGroup = 24
	UJoiningGroupUJGMIM                   UJoiningGroup = 25
	UJoiningGroupUJGNOON                  UJoiningGroup = 26
	UJoiningGroupUJGNUN                   UJoiningGroup = 27
	UJoiningGroupUJGPE                    UJoiningGroup = 28
	UJoiningGroupUJGQAF                   UJoiningGroup = 29
	UJoiningGroupUJGQAPH                  UJoiningGroup = 30
	UJoiningGroupUJGREH                   UJoiningGroup = 31
	UJoiningGroupUJGREVERSEDPE            UJoiningGroup = 32
	UJoiningGroupUJGSAD                   UJoiningGroup = 33
	UJoiningGroupUJGSADHE                 UJoiningGroup = 34
	UJoiningGroupUJGSEEN                  UJoiningGroup = 35
	UJoiningGroupUJGSEMKATH               UJoiningGroup = 36
	UJoiningGroupUJGSHIN                  UJoiningGroup = 37
	UJoiningGroupUJGSWASHKAF              UJoiningGroup = 38
	UJoiningGroupUJGSYRIACWAW             UJoiningGroup = 39
	UJoiningGroupUJGTAH                   UJoiningGroup = 40
	UJoiningGroupUJGTAW                   UJoiningGroup = 41
	UJoiningGroupUJGTEHMARBUTA            UJoiningGroup = 42
	UJoiningGroupUJGTETH                  UJoiningGroup = 43
	UJoiningGroupUJGWAW                   UJoiningGroup = 44
	UJoiningGroupUJGYEH                   UJoiningGroup = 45
	UJoiningGroupUJGYEHBARREE             UJoiningGroup = 46
	UJoiningGroupUJGYEHWITHTAIL           UJoiningGroup = 47
	UJoiningGroupUJGYUDH                  UJoiningGroup = 48
	UJoiningGroupUJGYUDHHE                UJoiningGroup = 49
	UJoiningGroupUJGZAIN                  UJoiningGroup = 50
	UJoiningGroupUJGFE                    UJoiningGroup = 51
	UJoiningGroupUJGKHAPH                 UJoiningGroup = 52
	UJoiningGroupUJGZHAIN                 UJoiningGroup = 53
	UJoiningGroupUJGBURUSHASKIYEHBARREE   UJoiningGroup = 54
	UJoiningGroupUJGFARSIYEH              UJoiningGroup = 55
	UJoiningGroupUJGNYA                   UJoiningGroup = 56
	UJoiningGroupUJGROHINGYAYEH           UJoiningGroup = 57
	UJoiningGroupUJGMANICHAEANALEPH       UJoiningGroup = 58
	UJoiningGroupUJGMANICHAEANAYIN        UJoiningGroup = 59
	UJoiningGroupUJGMANICHAEANBETH        UJoiningGroup = 60
	UJoiningGroupUJGMANICHAEANDALETH      UJoiningGroup = 61
	UJoiningGroupUJGMANICHAEANDHAMEDH     UJoiningGroup = 62
	UJoiningGroupUJGMANICHAEANFIVE        UJoiningGroup = 63
	UJoiningGroupUJGMANICHAEANGIMEL       UJoiningGroup = 64
	UJoiningGroupUJGMANICHAEANHETH        UJoiningGroup = 65
	UJoiningGroupUJGMANICHAEANHUNDRED     UJoiningGroup = 66
	UJoiningGroupUJGMANICHAEANKAPH        UJoiningGroup = 67
	UJoiningGroupUJGMANICHAEANLAMEDH      UJoiningGroup = 68
	UJoiningGroupUJGMANICHAEANMEM         UJoiningGroup = 69
	UJoiningGroupUJGMANICHAEANNUN         UJoiningGroup = 70
	UJoiningGroupUJGMANICHAEANONE         UJoiningGroup = 71
	UJoiningGroupUJGMANICHAEANPE          UJoiningGroup = 72
	UJoiningGroupUJGMANICHAEANQOPH        UJoiningGroup = 73
	UJoiningGroupUJGMANICHAEANRESH        UJoiningGroup = 74
	UJoiningGroupUJGMANICHAEANSADHE       UJoiningGroup = 75
	UJoiningGroupUJGMANICHAEANSAMEKH      UJoiningGroup = 76
	UJoiningGroupUJGMANICHAEANTAW         UJoiningGroup = 77
	UJoiningGroupUJGMANICHAEANTEN         UJoiningGroup = 78
	UJoiningGroupUJGMANICHAEANTETH        UJoiningGroup = 79
	UJoiningGroupUJGMANICHAEANTHAMEDH     UJoiningGroup = 80
	UJoiningGroupUJGMANICHAEANTWENTY      UJoiningGroup = 81
	UJoiningGroupUJGMANICHAEANWAW         UJoiningGroup = 82
	UJoiningGroupUJGMANICHAEANYODH        UJoiningGroup = 83
	UJoiningGroupUJGMANICHAEANZAYIN       UJoiningGroup = 84
	UJoiningGroupUJGSTRAIGHTWAW           UJoiningGroup = 85
	UJoiningGroupUJGAFRICANFEH            UJoiningGroup = 86
	UJoiningGroupUJGAFRICANNOON           UJoiningGroup = 87
	UJoiningGroupUJGAFRICANQAF            UJoiningGroup = 88
	UJoiningGroupUJGMALAYALAMBHA          UJoiningGroup = 89
	UJoiningGroupUJGMALAYALAMJA           UJoiningGroup = 90
	UJoiningGroupUJGMALAYALAMLLA          UJoiningGroup = 91
	UJoiningGroupUJGMALAYALAMLLLA         UJoiningGroup = 92
	UJoiningGroupUJGMALAYALAMNGA          UJoiningGroup = 93
	UJoiningGroupUJGMALAYALAMNNA          UJoiningGroup = 94
	UJoiningGroupUJGMALAYALAMNNNA         UJoiningGroup = 95
	UJoiningGroupUJGMALAYALAMNYA          UJoiningGroup = 96
	UJoiningGroupUJGMALAYALAMRA           UJoiningGroup = 97
	UJoiningGroupUJGMALAYALAMSSA          UJoiningGroup = 98
	UJoiningGroupUJGMALAYALAMTTA          UJoiningGroup = 99
	UJoiningGroupUJGHANIFIROHINGYAKINNAYA UJoiningGroup = 100
	UJoiningGroupUJGHANIFIROHINGYAPA      UJoiningGroup = 101
	UJoiningGroupUJGTHINYEH               UJoiningGroup = 102
	UJoiningGroupUJGVERTICALTAIL          UJoiningGroup = 103
	UJoiningGroupUJGKASHMIRIYEH           UJoiningGroup = 104
	UJoiningGroupUJGCOUNT                 UJoiningGroup = 105
)

type UGraphemeClusterBreak c.Int

const (
	UGraphemeClusterBreakUGCBOTHER             UGraphemeClusterBreak = 0
	UGraphemeClusterBreakUGCBCONTROL           UGraphemeClusterBreak = 1
	UGraphemeClusterBreakUGCBCR                UGraphemeClusterBreak = 2
	UGraphemeClusterBreakUGCBEXTEND            UGraphemeClusterBreak = 3
	UGraphemeClusterBreakUGCBL                 UGraphemeClusterBreak = 4
	UGraphemeClusterBreakUGCBLF                UGraphemeClusterBreak = 5
	UGraphemeClusterBreakUGCBLV                UGraphemeClusterBreak = 6
	UGraphemeClusterBreakUGCBLVT               UGraphemeClusterBreak = 7
	UGraphemeClusterBreakUGCBT                 UGraphemeClusterBreak = 8
	UGraphemeClusterBreakUGCBV                 UGraphemeClusterBreak = 9
	UGraphemeClusterBreakUGCBSPACINGMARK       UGraphemeClusterBreak = 10
	UGraphemeClusterBreakUGCBPREPEND           UGraphemeClusterBreak = 11
	UGraphemeClusterBreakUGCBREGIONALINDICATOR UGraphemeClusterBreak = 12
	UGraphemeClusterBreakUGCBEBASE             UGraphemeClusterBreak = 13
	UGraphemeClusterBreakUGCBEBASEGAZ          UGraphemeClusterBreak = 14
	UGraphemeClusterBreakUGCBEMODIFIER         UGraphemeClusterBreak = 15
	UGraphemeClusterBreakUGCBGLUEAFTERZWJ      UGraphemeClusterBreak = 16
	UGraphemeClusterBreakUGCBZWJ               UGraphemeClusterBreak = 17
	UGraphemeClusterBreakUGCBCOUNT             UGraphemeClusterBreak = 18
)

type UWordBreakValues c.Int

const (
	UWordBreakValuesUWBOTHER             UWordBreakValues = 0
	UWordBreakValuesUWBALETTER           UWordBreakValues = 1
	UWordBreakValuesUWBFORMAT            UWordBreakValues = 2
	UWordBreakValuesUWBKATAKANA          UWordBreakValues = 3
	UWordBreakValuesUWBMIDLETTER         UWordBreakValues = 4
	UWordBreakValuesUWBMIDNUM            UWordBreakValues = 5
	UWordBreakValuesUWBNUMERIC           UWordBreakValues = 6
	UWordBreakValuesUWBEXTENDNUMLET      UWordBreakValues = 7
	UWordBreakValuesUWBCR                UWordBreakValues = 8
	UWordBreakValuesUWBEXTEND            UWordBreakValues = 9
	UWordBreakValuesUWBLF                UWordBreakValues = 10
	UWordBreakValuesUWBMIDNUMLET         UWordBreakValues = 11
	UWordBreakValuesUWBNEWLINE           UWordBreakValues = 12
	UWordBreakValuesUWBREGIONALINDICATOR UWordBreakValues = 13
	UWordBreakValuesUWBHEBREWLETTER      UWordBreakValues = 14
	UWordBreakValuesUWBSINGLEQUOTE       UWordBreakValues = 15
	UWordBreakValuesUWBDOUBLEQUOTE       UWordBreakValues = 16
	UWordBreakValuesUWBEBASE             UWordBreakValues = 17
	UWordBreakValuesUWBEBASEGAZ          UWordBreakValues = 18
	UWordBreakValuesUWBEMODIFIER         UWordBreakValues = 19
	UWordBreakValuesUWBGLUEAFTERZWJ      UWordBreakValues = 20
	UWordBreakValuesUWBZWJ               UWordBreakValues = 21
	UWordBreakValuesUWBWSEGSPACE         UWordBreakValues = 22
	UWordBreakValuesUWBCOUNT             UWordBreakValues = 23
)

type USentenceBreak c.Int

const (
	USentenceBreakUSBOTHER     USentenceBreak = 0
	USentenceBreakUSBATERM     USentenceBreak = 1
	USentenceBreakUSBCLOSE     USentenceBreak = 2
	USentenceBreakUSBFORMAT    USentenceBreak = 3
	USentenceBreakUSBLOWER     USentenceBreak = 4
	USentenceBreakUSBNUMERIC   USentenceBreak = 5
	USentenceBreakUSBOLETTER   USentenceBreak = 6
	USentenceBreakUSBSEP       USentenceBreak = 7
	USentenceBreakUSBSP        USentenceBreak = 8
	USentenceBreakUSBSTERM     USentenceBreak = 9
	USentenceBreakUSBUPPER     USentenceBreak = 10
	USentenceBreakUSBCR        USentenceBreak = 11
	USentenceBreakUSBEXTEND    USentenceBreak = 12
	USentenceBreakUSBLF        USentenceBreak = 13
	USentenceBreakUSBSCONTINUE USentenceBreak = 14
	USentenceBreakUSBCOUNT     USentenceBreak = 15
)

type ULineBreak c.Int

const (
	ULineBreakULBUNKNOWN                    ULineBreak = 0
	ULineBreakULBAMBIGUOUS                  ULineBreak = 1
	ULineBreakULBALPHABETIC                 ULineBreak = 2
	ULineBreakULBBREAKBOTH                  ULineBreak = 3
	ULineBreakULBBREAKAFTER                 ULineBreak = 4
	ULineBreakULBBREAKBEFORE                ULineBreak = 5
	ULineBreakULBMANDATORYBREAK             ULineBreak = 6
	ULineBreakULBCONTINGENTBREAK            ULineBreak = 7
	ULineBreakULBCLOSEPUNCTUATION           ULineBreak = 8
	ULineBreakULBCOMBININGMARK              ULineBreak = 9
	ULineBreakULBCARRIAGERETURN             ULineBreak = 10
	ULineBreakULBEXCLAMATION                ULineBreak = 11
	ULineBreakULBGLUE                       ULineBreak = 12
	ULineBreakULBHYPHEN                     ULineBreak = 13
	ULineBreakULBIDEOGRAPHIC                ULineBreak = 14
	ULineBreakULBINSEPARABLE                ULineBreak = 15
	ULineBreakULBINSEPERABLE                ULineBreak = 15
	ULineBreakULBINFIXNUMERIC               ULineBreak = 16
	ULineBreakULBLINEFEED                   ULineBreak = 17
	ULineBreakULBNONSTARTER                 ULineBreak = 18
	ULineBreakULBNUMERIC                    ULineBreak = 19
	ULineBreakULBOPENPUNCTUATION            ULineBreak = 20
	ULineBreakULBPOSTFIXNUMERIC             ULineBreak = 21
	ULineBreakULBPREFIXNUMERIC              ULineBreak = 22
	ULineBreakULBQUOTATION                  ULineBreak = 23
	ULineBreakULBCOMPLEXCONTEXT             ULineBreak = 24
	ULineBreakULBSURROGATE                  ULineBreak = 25
	ULineBreakULBSPACE                      ULineBreak = 26
	ULineBreakULBBREAKSYMBOLS               ULineBreak = 27
	ULineBreakULBZWSPACE                    ULineBreak = 28
	ULineBreakULBNEXTLINE                   ULineBreak = 29
	ULineBreakULBWORDJOINER                 ULineBreak = 30
	ULineBreakULBH2                         ULineBreak = 31
	ULineBreakULBH3                         ULineBreak = 32
	ULineBreakULBJL                         ULineBreak = 33
	ULineBreakULBJT                         ULineBreak = 34
	ULineBreakULBJV                         ULineBreak = 35
	ULineBreakULBCLOSEPARENTHESIS           ULineBreak = 36
	ULineBreakULBCONDITIONALJAPANESESTARTER ULineBreak = 37
	ULineBreakULBHEBREWLETTER               ULineBreak = 38
	ULineBreakULBREGIONALINDICATOR          ULineBreak = 39
	ULineBreakULBEBASE                      ULineBreak = 40
	ULineBreakULBEMODIFIER                  ULineBreak = 41
	ULineBreakULBZWJ                        ULineBreak = 42
	ULineBreakULBAKSARA                     ULineBreak = 43
	ULineBreakULBAKSARAPREBASE              ULineBreak = 44
	ULineBreakULBAKSARASTART                ULineBreak = 45
	ULineBreakULBVIRAMAFINAL                ULineBreak = 46
	ULineBreakULBVIRAMA                     ULineBreak = 47
	ULineBreakULBCOUNT                      ULineBreak = 48
)

type UNumericType c.Int

const (
	UNumericTypeUNTNONE    UNumericType = 0
	UNumericTypeUNTDECIMAL UNumericType = 1
	UNumericTypeUNTDIGIT   UNumericType = 2
	UNumericTypeUNTNUMERIC UNumericType = 3
	UNumericTypeUNTCOUNT   UNumericType = 4
)

type UHangulSyllableType c.Int

const (
	UHangulSyllableTypeUHSTNOTAPPLICABLE UHangulSyllableType = 0
	UHangulSyllableTypeUHSTLEADINGJAMO   UHangulSyllableType = 1
	UHangulSyllableTypeUHSTVOWELJAMO     UHangulSyllableType = 2
	UHangulSyllableTypeUHSTTRAILINGJAMO  UHangulSyllableType = 3
	UHangulSyllableTypeUHSTLVSYLLABLE    UHangulSyllableType = 4
	UHangulSyllableTypeUHSTLVTSYLLABLE   UHangulSyllableType = 5
	UHangulSyllableTypeUHSTCOUNT         UHangulSyllableType = 6
)

type UIndicPositionalCategory c.Int

const (
	UIndicPositionalCategoryUINPCNA                   UIndicPositionalCategory = 0
	UIndicPositionalCategoryUINPCBOTTOM               UIndicPositionalCategory = 1
	UIndicPositionalCategoryUINPCBOTTOMANDLEFT        UIndicPositionalCategory = 2
	UIndicPositionalCategoryUINPCBOTTOMANDRIGHT       UIndicPositionalCategory = 3
	UIndicPositionalCategoryUINPCLEFT                 UIndicPositionalCategory = 4
	UIndicPositionalCategoryUINPCLEFTANDRIGHT         UIndicPositionalCategory = 5
	UIndicPositionalCategoryUINPCOVERSTRUCK           UIndicPositionalCategory = 6
	UIndicPositionalCategoryUINPCRIGHT                UIndicPositionalCategory = 7
	UIndicPositionalCategoryUINPCTOP                  UIndicPositionalCategory = 8
	UIndicPositionalCategoryUINPCTOPANDBOTTOM         UIndicPositionalCategory = 9
	UIndicPositionalCategoryUINPCTOPANDBOTTOMANDRIGHT UIndicPositionalCategory = 10
	UIndicPositionalCategoryUINPCTOPANDLEFT           UIndicPositionalCategory = 11
	UIndicPositionalCategoryUINPCTOPANDLEFTANDRIGHT   UIndicPositionalCategory = 12
	UIndicPositionalCategoryUINPCTOPANDRIGHT          UIndicPositionalCategory = 13
	UIndicPositionalCategoryUINPCVISUALORDERLEFT      UIndicPositionalCategory = 14
	UIndicPositionalCategoryUINPCTOPANDBOTTOMANDLEFT  UIndicPositionalCategory = 15
)

type UIndicSyllabicCategory c.Int

const (
	UIndicSyllabicCategoryUINSCOTHER                     UIndicSyllabicCategory = 0
	UIndicSyllabicCategoryUINSCAVAGRAHA                  UIndicSyllabicCategory = 1
	UIndicSyllabicCategoryUINSCBINDU                     UIndicSyllabicCategory = 2
	UIndicSyllabicCategoryUINSCBRAHMIJOININGNUMBER       UIndicSyllabicCategory = 3
	UIndicSyllabicCategoryUINSCCANTILLATIONMARK          UIndicSyllabicCategory = 4
	UIndicSyllabicCategoryUINSCCONSONANT                 UIndicSyllabicCategory = 5
	UIndicSyllabicCategoryUINSCCONSONANTDEAD             UIndicSyllabicCategory = 6
	UIndicSyllabicCategoryUINSCCONSONANTFINAL            UIndicSyllabicCategory = 7
	UIndicSyllabicCategoryUINSCCONSONANTHEADLETTER       UIndicSyllabicCategory = 8
	UIndicSyllabicCategoryUINSCCONSONANTINITIALPOSTFIXED UIndicSyllabicCategory = 9
	UIndicSyllabicCategoryUINSCCONSONANTKILLER           UIndicSyllabicCategory = 10
	UIndicSyllabicCategoryUINSCCONSONANTMEDIAL           UIndicSyllabicCategory = 11
	UIndicSyllabicCategoryUINSCCONSONANTPLACEHOLDER      UIndicSyllabicCategory = 12
	UIndicSyllabicCategoryUINSCCONSONANTPRECEDINGREPHA   UIndicSyllabicCategory = 13
	UIndicSyllabicCategoryUINSCCONSONANTPREFIXED         UIndicSyllabicCategory = 14
	UIndicSyllabicCategoryUINSCCONSONANTSUBJOINED        UIndicSyllabicCategory = 15
	UIndicSyllabicCategoryUINSCCONSONANTSUCCEEDINGREPHA  UIndicSyllabicCategory = 16
	UIndicSyllabicCategoryUINSCCONSONANTWITHSTACKER      UIndicSyllabicCategory = 17
	UIndicSyllabicCategoryUINSCGEMINATIONMARK            UIndicSyllabicCategory = 18
	UIndicSyllabicCategoryUINSCINVISIBLESTACKER          UIndicSyllabicCategory = 19
	UIndicSyllabicCategoryUINSCJOINER                    UIndicSyllabicCategory = 20
	UIndicSyllabicCategoryUINSCMODIFYINGLETTER           UIndicSyllabicCategory = 21
	UIndicSyllabicCategoryUINSCNONJOINER                 UIndicSyllabicCategory = 22
	UIndicSyllabicCategoryUINSCNUKTA                     UIndicSyllabicCategory = 23
	UIndicSyllabicCategoryUINSCNUMBER                    UIndicSyllabicCategory = 24
	UIndicSyllabicCategoryUINSCNUMBERJOINER              UIndicSyllabicCategory = 25
	UIndicSyllabicCategoryUINSCPUREKILLER                UIndicSyllabicCategory = 26
	UIndicSyllabicCategoryUINSCREGISTERSHIFTER           UIndicSyllabicCategory = 27
	UIndicSyllabicCategoryUINSCSYLLABLEMODIFIER          UIndicSyllabicCategory = 28
	UIndicSyllabicCategoryUINSCTONELETTER                UIndicSyllabicCategory = 29
	UIndicSyllabicCategoryUINSCTONEMARK                  UIndicSyllabicCategory = 30
	UIndicSyllabicCategoryUINSCVIRAMA                    UIndicSyllabicCategory = 31
	UIndicSyllabicCategoryUINSCVISARGA                   UIndicSyllabicCategory = 32
	UIndicSyllabicCategoryUINSCVOWEL                     UIndicSyllabicCategory = 33
	UIndicSyllabicCategoryUINSCVOWELDEPENDENT            UIndicSyllabicCategory = 34
	UIndicSyllabicCategoryUINSCVOWELINDEPENDENT          UIndicSyllabicCategory = 35
	UIndicSyllabicCategoryUINSCREORDERINGKILLER          UIndicSyllabicCategory = 36
)

type UIndicConjunctBreak c.Int

const (
	UIndicConjunctBreakUINCBNONE      UIndicConjunctBreak = 0
	UIndicConjunctBreakUINCBCONSONANT UIndicConjunctBreak = 1
	UIndicConjunctBreakUINCBEXTEND    UIndicConjunctBreak = 2
	UIndicConjunctBreakUINCBLINKER    UIndicConjunctBreak = 3
)

type UVerticalOrientation c.Int

const (
	UVerticalOrientationUVOROTATED            UVerticalOrientation = 0
	UVerticalOrientationUVOTRANSFORMEDROTATED UVerticalOrientation = 1
	UVerticalOrientationUVOTRANSFORMEDUPRIGHT UVerticalOrientation = 2
	UVerticalOrientationUVOUPRIGHT            UVerticalOrientation = 3
)

type UIdentifierStatus c.Int

const (
	UIdentifierStatusUIDSTATUSRESTRICTED UIdentifierStatus = 0
	UIdentifierStatusUIDSTATUSALLOWED    UIdentifierStatus = 1
)

type UIdentifierType c.Int

const (
	UIdentifierTypeUIDTYPENOTCHARACTER     UIdentifierType = 0
	UIdentifierTypeUIDTYPEDEPRECATED       UIdentifierType = 1
	UIdentifierTypeUIDTYPEDEFAULTIGNORABLE UIdentifierType = 2
	UIdentifierTypeUIDTYPENOTNFKC          UIdentifierType = 3
	UIdentifierTypeUIDTYPENOTXID           UIdentifierType = 4
	UIdentifierTypeUIDTYPEEXCLUSION        UIdentifierType = 5
	UIdentifierTypeUIDTYPEOBSOLETE         UIdentifierType = 6
	UIdentifierTypeUIDTYPETECHNICAL        UIdentifierType = 7
	UIdentifierTypeUIDTYPEUNCOMMONUSE      UIdentifierType = 8
	UIdentifierTypeUIDTYPELIMITEDUSE       UIdentifierType = 9
	UIdentifierTypeUIDTYPEINCLUSION        UIdentifierType = 10
	UIdentifierTypeUIDTYPERECOMMENDED      UIdentifierType = 11
)
// llgo:type C
type UCharEnumTypeRange func(unsafe.Pointer, UChar32, UChar32, UCharCategory) UBool
// llgo:type C
type UEnumCharNamesFn func(unsafe.Pointer, UChar32, UCharNameChoice, *int8, int32) UBool
