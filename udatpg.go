package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type UDateTimePatternGenerator unsafe.Pointer
type UDateTimePatternField c.Int

const (
	UDateTimePatternFieldUDATPGERAFIELD              UDateTimePatternField = 0
	UDateTimePatternFieldUDATPGYEARFIELD             UDateTimePatternField = 1
	UDateTimePatternFieldUDATPGQUARTERFIELD          UDateTimePatternField = 2
	UDateTimePatternFieldUDATPGMONTHFIELD            UDateTimePatternField = 3
	UDateTimePatternFieldUDATPGWEEKOFYEARFIELD       UDateTimePatternField = 4
	UDateTimePatternFieldUDATPGWEEKOFMONTHFIELD      UDateTimePatternField = 5
	UDateTimePatternFieldUDATPGWEEKDAYFIELD          UDateTimePatternField = 6
	UDateTimePatternFieldUDATPGDAYOFYEARFIELD        UDateTimePatternField = 7
	UDateTimePatternFieldUDATPGDAYOFWEEKINMONTHFIELD UDateTimePatternField = 8
	UDateTimePatternFieldUDATPGDAYFIELD              UDateTimePatternField = 9
	UDateTimePatternFieldUDATPGDAYPERIODFIELD        UDateTimePatternField = 10
	UDateTimePatternFieldUDATPGHOURFIELD             UDateTimePatternField = 11
	UDateTimePatternFieldUDATPGMINUTEFIELD           UDateTimePatternField = 12
	UDateTimePatternFieldUDATPGSECONDFIELD           UDateTimePatternField = 13
	UDateTimePatternFieldUDATPGFRACTIONALSECONDFIELD UDateTimePatternField = 14
	UDateTimePatternFieldUDATPGZONEFIELD             UDateTimePatternField = 15
	UDateTimePatternFieldUDATPGFIELDCOUNT            UDateTimePatternField = 16
)

type UDateTimePGDisplayWidth c.Int

const (
	UDateTimePGDisplayWidthUDATPGWIDE        UDateTimePGDisplayWidth = 0
	UDateTimePGDisplayWidthUDATPGABBREVIATED UDateTimePGDisplayWidth = 1
	UDateTimePGDisplayWidthUDATPGNARROW      UDateTimePGDisplayWidth = 2
)

type UDateTimePatternMatchOptions c.Int

const (
	UDateTimePatternMatchOptionsUDATPGMATCHNOOPTIONS         UDateTimePatternMatchOptions = 0
	UDateTimePatternMatchOptionsUDATPGMATCHHOURFIELDLENGTH   UDateTimePatternMatchOptions = 2048
	UDateTimePatternMatchOptionsUDATPGMATCHMINUTEFIELDLENGTH UDateTimePatternMatchOptions = 4096
	UDateTimePatternMatchOptionsUDATPGMATCHSECONDFIELDLENGTH UDateTimePatternMatchOptions = 8192
	UDateTimePatternMatchOptionsUDATPGMATCHALLFIELDSLENGTH   UDateTimePatternMatchOptions = 65535
)

type UDateTimePatternConflict c.Int

const (
	UDateTimePatternConflictUDATPGNOCONFLICT    UDateTimePatternConflict = 0
	UDateTimePatternConflictUDATPGBASECONFLICT  UDateTimePatternConflict = 1
	UDateTimePatternConflictUDATPGCONFLICT      UDateTimePatternConflict = 2
	UDateTimePatternConflictUDATPGCONFLICTCOUNT UDateTimePatternConflict = 3
)
