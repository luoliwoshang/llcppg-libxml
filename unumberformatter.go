package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UNumberRoundingPriority c.Int

const (
	UNumberRoundingPriorityUNUMROUNDINGPRIORITYRELAXED UNumberRoundingPriority = 0
	UNumberRoundingPriorityUNUMROUNDINGPRIORITYSTRICT  UNumberRoundingPriority = 1
)

type UNumberUnitWidth c.Int

const (
	UNumberUnitWidthUNUMUNITWIDTHNARROW   UNumberUnitWidth = 0
	UNumberUnitWidthUNUMUNITWIDTHSHORT    UNumberUnitWidth = 1
	UNumberUnitWidthUNUMUNITWIDTHFULLNAME UNumberUnitWidth = 2
	UNumberUnitWidthUNUMUNITWIDTHISOCODE  UNumberUnitWidth = 3
	UNumberUnitWidthUNUMUNITWIDTHFORMAL   UNumberUnitWidth = 4
	UNumberUnitWidthUNUMUNITWIDTHVARIANT  UNumberUnitWidth = 5
	UNumberUnitWidthUNUMUNITWIDTHHIDDEN   UNumberUnitWidth = 6
	UNumberUnitWidthUNUMUNITWIDTHCOUNT    UNumberUnitWidth = 7
)

type UNumberSignDisplay c.Int

const (
	UNumberSignDisplayUNUMSIGNAUTO                 UNumberSignDisplay = 0
	UNumberSignDisplayUNUMSIGNALWAYS               UNumberSignDisplay = 1
	UNumberSignDisplayUNUMSIGNNEVER                UNumberSignDisplay = 2
	UNumberSignDisplayUNUMSIGNACCOUNTING           UNumberSignDisplay = 3
	UNumberSignDisplayUNUMSIGNACCOUNTINGALWAYS     UNumberSignDisplay = 4
	UNumberSignDisplayUNUMSIGNEXCEPTZERO           UNumberSignDisplay = 5
	UNumberSignDisplayUNUMSIGNACCOUNTINGEXCEPTZERO UNumberSignDisplay = 6
	UNumberSignDisplayUNUMSIGNNEGATIVE             UNumberSignDisplay = 7
	UNumberSignDisplayUNUMSIGNACCOUNTINGNEGATIVE   UNumberSignDisplay = 8
	UNumberSignDisplayUNUMSIGNCOUNT                UNumberSignDisplay = 9
)

type UNumberDecimalSeparatorDisplay c.Int

const (
	UNumberDecimalSeparatorDisplayUNUMDECIMALSEPARATORAUTO   UNumberDecimalSeparatorDisplay = 0
	UNumberDecimalSeparatorDisplayUNUMDECIMALSEPARATORALWAYS UNumberDecimalSeparatorDisplay = 1
	UNumberDecimalSeparatorDisplayUNUMDECIMALSEPARATORCOUNT  UNumberDecimalSeparatorDisplay = 2
)

type UNumberTrailingZeroDisplay c.Int

const (
	UNumberTrailingZeroDisplayUNUMTRAILINGZEROAUTO        UNumberTrailingZeroDisplay = 0
	UNumberTrailingZeroDisplayUNUMTRAILINGZEROHIDEIFWHOLE UNumberTrailingZeroDisplay = 1
)

type UNumberFormatter struct {
	Unused [8]uint8
}
