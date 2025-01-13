package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type UNumberFormat unsafe.Pointer
type UNumberFormatStyle c.Int

const (
	UNumberFormatStyleUNUMPATTERNDECIMAL      UNumberFormatStyle = 0
	UNumberFormatStyleUNUMDECIMAL             UNumberFormatStyle = 1
	UNumberFormatStyleUNUMCURRENCY            UNumberFormatStyle = 2
	UNumberFormatStyleUNUMPERCENT             UNumberFormatStyle = 3
	UNumberFormatStyleUNUMSCIENTIFIC          UNumberFormatStyle = 4
	UNumberFormatStyleUNUMSPELLOUT            UNumberFormatStyle = 5
	UNumberFormatStyleUNUMORDINAL             UNumberFormatStyle = 6
	UNumberFormatStyleUNUMDURATION            UNumberFormatStyle = 7
	UNumberFormatStyleUNUMNUMBERINGSYSTEM     UNumberFormatStyle = 8
	UNumberFormatStyleUNUMPATTERNRULEBASED    UNumberFormatStyle = 9
	UNumberFormatStyleUNUMCURRENCYISO         UNumberFormatStyle = 10
	UNumberFormatStyleUNUMCURRENCYPLURAL      UNumberFormatStyle = 11
	UNumberFormatStyleUNUMCURRENCYACCOUNTING  UNumberFormatStyle = 12
	UNumberFormatStyleUNUMCASHCURRENCY        UNumberFormatStyle = 13
	UNumberFormatStyleUNUMDECIMALCOMPACTSHORT UNumberFormatStyle = 14
	UNumberFormatStyleUNUMDECIMALCOMPACTLONG  UNumberFormatStyle = 15
	UNumberFormatStyleUNUMCURRENCYSTANDARD    UNumberFormatStyle = 16
	UNumberFormatStyleUNUMFORMATSTYLECOUNT    UNumberFormatStyle = 17
	UNumberFormatStyleUNUMDEFAULT             UNumberFormatStyle = 1
	UNumberFormatStyleUNUMIGNORE              UNumberFormatStyle = 0
)

type UNumberFormatPadPosition c.Int

const (
	UNumberFormatPadPositionUNUMPADBEFOREPREFIX UNumberFormatPadPosition = 0
	UNumberFormatPadPositionUNUMPADAFTERPREFIX  UNumberFormatPadPosition = 1
	UNumberFormatPadPositionUNUMPADBEFORESUFFIX UNumberFormatPadPosition = 2
	UNumberFormatPadPositionUNUMPADAFTERSUFFIX  UNumberFormatPadPosition = 3
)

type UNumberCompactStyle c.Int

const (
	UNumberCompactStyleUNUMSHORT UNumberCompactStyle = 0
	UNumberCompactStyleUNUMLONG  UNumberCompactStyle = 1
)

type UCurrencySpacing c.Int

const (
	UCurrencySpacingUNUMCURRENCYMATCH            UCurrencySpacing = 0
	UCurrencySpacingUNUMCURRENCYSURROUNDINGMATCH UCurrencySpacing = 1
	UCurrencySpacingUNUMCURRENCYINSERT           UCurrencySpacing = 2
	UCurrencySpacingUNUMCURRENCYSPACINGCOUNT     UCurrencySpacing = 3
)

type UNumberFormatFields c.Int

const (
	UNumberFormatFieldsUNUMINTEGERFIELD           UNumberFormatFields = 0
	UNumberFormatFieldsUNUMFRACTIONFIELD          UNumberFormatFields = 1
	UNumberFormatFieldsUNUMDECIMALSEPARATORFIELD  UNumberFormatFields = 2
	UNumberFormatFieldsUNUMEXPONENTSYMBOLFIELD    UNumberFormatFields = 3
	UNumberFormatFieldsUNUMEXPONENTSIGNFIELD      UNumberFormatFields = 4
	UNumberFormatFieldsUNUMEXPONENTFIELD          UNumberFormatFields = 5
	UNumberFormatFieldsUNUMGROUPINGSEPARATORFIELD UNumberFormatFields = 6
	UNumberFormatFieldsUNUMCURRENCYFIELD          UNumberFormatFields = 7
	UNumberFormatFieldsUNUMPERCENTFIELD           UNumberFormatFields = 8
	UNumberFormatFieldsUNUMPERMILLFIELD           UNumberFormatFields = 9
	UNumberFormatFieldsUNUMSIGNFIELD              UNumberFormatFields = 10
	UNumberFormatFieldsUNUMMEASUREUNITFIELD       UNumberFormatFields = 11
	UNumberFormatFieldsUNUMCOMPACTFIELD           UNumberFormatFields = 12
	UNumberFormatFieldsUNUMAPPROXIMATELYSIGNFIELD UNumberFormatFields = 13
	UNumberFormatFieldsUNUMFIELDCOUNT             UNumberFormatFields = 14
)

type UNumberFormatMinimumGroupingDigits c.Int

const (
	UNumberFormatMinimumGroupingDigitsUNUMMINIMUMGROUPINGDIGITSAUTO UNumberFormatMinimumGroupingDigits = -2
	UNumberFormatMinimumGroupingDigitsUNUMMINIMUMGROUPINGDIGITSMIN2 UNumberFormatMinimumGroupingDigits = -3
)

type UNumberFormatAttributeValue c.Int

const (
	UNumberFormatAttributeValueUNUMNO    UNumberFormatAttributeValue = 0
	UNumberFormatAttributeValueUNUMYES   UNumberFormatAttributeValue = 1
	UNumberFormatAttributeValueUNUMMAYBE UNumberFormatAttributeValue = 2
)

type UNumberFormatAttribute c.Int

const (
	UNumberFormatAttributeUNUMPARSEINTONLY                  UNumberFormatAttribute = 0
	UNumberFormatAttributeUNUMGROUPINGUSED                  UNumberFormatAttribute = 1
	UNumberFormatAttributeUNUMDECIMALALWAYSSHOWN            UNumberFormatAttribute = 2
	UNumberFormatAttributeUNUMMAXINTEGERDIGITS              UNumberFormatAttribute = 3
	UNumberFormatAttributeUNUMMININTEGERDIGITS              UNumberFormatAttribute = 4
	UNumberFormatAttributeUNUMINTEGERDIGITS                 UNumberFormatAttribute = 5
	UNumberFormatAttributeUNUMMAXFRACTIONDIGITS             UNumberFormatAttribute = 6
	UNumberFormatAttributeUNUMMINFRACTIONDIGITS             UNumberFormatAttribute = 7
	UNumberFormatAttributeUNUMFRACTIONDIGITS                UNumberFormatAttribute = 8
	UNumberFormatAttributeUNUMMULTIPLIER                    UNumberFormatAttribute = 9
	UNumberFormatAttributeUNUMGROUPINGSIZE                  UNumberFormatAttribute = 10
	UNumberFormatAttributeUNUMROUNDINGMODE                  UNumberFormatAttribute = 11
	UNumberFormatAttributeUNUMROUNDINGINCREMENT             UNumberFormatAttribute = 12
	UNumberFormatAttributeUNUMFORMATWIDTH                   UNumberFormatAttribute = 13
	UNumberFormatAttributeUNUMPADDINGPOSITION               UNumberFormatAttribute = 14
	UNumberFormatAttributeUNUMSECONDARYGROUPINGSIZE         UNumberFormatAttribute = 15
	UNumberFormatAttributeUNUMSIGNIFICANTDIGITSUSED         UNumberFormatAttribute = 16
	UNumberFormatAttributeUNUMMINSIGNIFICANTDIGITS          UNumberFormatAttribute = 17
	UNumberFormatAttributeUNUMMAXSIGNIFICANTDIGITS          UNumberFormatAttribute = 18
	UNumberFormatAttributeUNUMLENIENTPARSE                  UNumberFormatAttribute = 19
	UNumberFormatAttributeUNUMPARSEALLINPUT                 UNumberFormatAttribute = 20
	UNumberFormatAttributeUNUMSCALE                         UNumberFormatAttribute = 21
	UNumberFormatAttributeUNUMMINIMUMGROUPINGDIGITS         UNumberFormatAttribute = 22
	UNumberFormatAttributeUNUMCURRENCYUSAGE                 UNumberFormatAttribute = 23
	UNumberFormatAttributeUNUMMAXNONBOOLEANATTRIBUTE        UNumberFormatAttribute = 4095
	UNumberFormatAttributeUNUMFORMATFAILIFMORETHANMAXDIGITS UNumberFormatAttribute = 4096
	UNumberFormatAttributeUNUMPARSENOEXPONENT               UNumberFormatAttribute = 4097
	UNumberFormatAttributeUNUMPARSEDECIMALMARKREQUIRED      UNumberFormatAttribute = 4098
	UNumberFormatAttributeUNUMPARSECASESENSITIVE            UNumberFormatAttribute = 4099
	UNumberFormatAttributeUNUMSIGNALWAYSSHOWN               UNumberFormatAttribute = 4100
	UNumberFormatAttributeUNUMLIMITBOOLEANATTRIBUTE         UNumberFormatAttribute = 4101
)

type UNumberFormatTextAttribute c.Int

const (
	UNumberFormatTextAttributeUNUMPOSITIVEPREFIX   UNumberFormatTextAttribute = 0
	UNumberFormatTextAttributeUNUMPOSITIVESUFFIX   UNumberFormatTextAttribute = 1
	UNumberFormatTextAttributeUNUMNEGATIVEPREFIX   UNumberFormatTextAttribute = 2
	UNumberFormatTextAttributeUNUMNEGATIVESUFFIX   UNumberFormatTextAttribute = 3
	UNumberFormatTextAttributeUNUMPADDINGCHARACTER UNumberFormatTextAttribute = 4
	UNumberFormatTextAttributeUNUMCURRENCYCODE     UNumberFormatTextAttribute = 5
	UNumberFormatTextAttributeUNUMDEFAULTRULESET   UNumberFormatTextAttribute = 6
	UNumberFormatTextAttributeUNUMPUBLICRULESETS   UNumberFormatTextAttribute = 7
)

type UNumberFormatSymbol c.Int

const (
	UNumberFormatSymbolUNUMDECIMALSEPARATORSYMBOL          UNumberFormatSymbol = 0
	UNumberFormatSymbolUNUMGROUPINGSEPARATORSYMBOL         UNumberFormatSymbol = 1
	UNumberFormatSymbolUNUMPATTERNSEPARATORSYMBOL          UNumberFormatSymbol = 2
	UNumberFormatSymbolUNUMPERCENTSYMBOL                   UNumberFormatSymbol = 3
	UNumberFormatSymbolUNUMZERODIGITSYMBOL                 UNumberFormatSymbol = 4
	UNumberFormatSymbolUNUMDIGITSYMBOL                     UNumberFormatSymbol = 5
	UNumberFormatSymbolUNUMMINUSSIGNSYMBOL                 UNumberFormatSymbol = 6
	UNumberFormatSymbolUNUMPLUSSIGNSYMBOL                  UNumberFormatSymbol = 7
	UNumberFormatSymbolUNUMCURRENCYSYMBOL                  UNumberFormatSymbol = 8
	UNumberFormatSymbolUNUMINTLCURRENCYSYMBOL              UNumberFormatSymbol = 9
	UNumberFormatSymbolUNUMMONETARYSEPARATORSYMBOL         UNumberFormatSymbol = 10
	UNumberFormatSymbolUNUMEXPONENTIALSYMBOL               UNumberFormatSymbol = 11
	UNumberFormatSymbolUNUMPERMILLSYMBOL                   UNumberFormatSymbol = 12
	UNumberFormatSymbolUNUMPADESCAPESYMBOL                 UNumberFormatSymbol = 13
	UNumberFormatSymbolUNUMINFINITYSYMBOL                  UNumberFormatSymbol = 14
	UNumberFormatSymbolUNUMNANSYMBOL                       UNumberFormatSymbol = 15
	UNumberFormatSymbolUNUMSIGNIFICANTDIGITSYMBOL          UNumberFormatSymbol = 16
	UNumberFormatSymbolUNUMMONETARYGROUPINGSEPARATORSYMBOL UNumberFormatSymbol = 17
	UNumberFormatSymbolUNUMONEDIGITSYMBOL                  UNumberFormatSymbol = 18
	UNumberFormatSymbolUNUMTWODIGITSYMBOL                  UNumberFormatSymbol = 19
	UNumberFormatSymbolUNUMTHREEDIGITSYMBOL                UNumberFormatSymbol = 20
	UNumberFormatSymbolUNUMFOURDIGITSYMBOL                 UNumberFormatSymbol = 21
	UNumberFormatSymbolUNUMFIVEDIGITSYMBOL                 UNumberFormatSymbol = 22
	UNumberFormatSymbolUNUMSIXDIGITSYMBOL                  UNumberFormatSymbol = 23
	UNumberFormatSymbolUNUMSEVENDIGITSYMBOL                UNumberFormatSymbol = 24
	UNumberFormatSymbolUNUMEIGHTDIGITSYMBOL                UNumberFormatSymbol = 25
	UNumberFormatSymbolUNUMNINEDIGITSYMBOL                 UNumberFormatSymbol = 26
	UNumberFormatSymbolUNUMEXPONENTMULTIPLICATIONSYMBOL    UNumberFormatSymbol = 27
	UNumberFormatSymbolUNUMAPPROXIMATELYSIGNSYMBOL         UNumberFormatSymbol = 28
	UNumberFormatSymbolUNUMFORMATSYMBOLCOUNT               UNumberFormatSymbol = 29
)
