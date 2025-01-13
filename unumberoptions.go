package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UNumberFormatRoundingMode c.Int

const (
	UNumberFormatRoundingModeUNUMROUNDCEILING     UNumberFormatRoundingMode = 0
	UNumberFormatRoundingModeUNUMROUNDFLOOR       UNumberFormatRoundingMode = 1
	UNumberFormatRoundingModeUNUMROUNDDOWN        UNumberFormatRoundingMode = 2
	UNumberFormatRoundingModeUNUMROUNDUP          UNumberFormatRoundingMode = 3
	UNumberFormatRoundingModeUNUMROUNDHALFEVEN    UNumberFormatRoundingMode = 4
	UNumberFormatRoundingModeUNUMFOUNDHALFEVEN    UNumberFormatRoundingMode = 4
	UNumberFormatRoundingModeUNUMROUNDHALFDOWN    UNumberFormatRoundingMode = 5
	UNumberFormatRoundingModeUNUMROUNDHALFUP      UNumberFormatRoundingMode = 6
	UNumberFormatRoundingModeUNUMROUNDUNNECESSARY UNumberFormatRoundingMode = 7
	UNumberFormatRoundingModeUNUMROUNDHALFODD     UNumberFormatRoundingMode = 8
	UNumberFormatRoundingModeUNUMROUNDHALFCEILING UNumberFormatRoundingMode = 9
	UNumberFormatRoundingModeUNUMROUNDHALFFLOOR   UNumberFormatRoundingMode = 10
)

type UNumberGroupingStrategy c.Int

const (
	UNumberGroupingStrategyUNUMGROUPINGOFF       UNumberGroupingStrategy = 0
	UNumberGroupingStrategyUNUMGROUPINGMIN2      UNumberGroupingStrategy = 1
	UNumberGroupingStrategyUNUMGROUPINGAUTO      UNumberGroupingStrategy = 2
	UNumberGroupingStrategyUNUMGROUPINGONALIGNED UNumberGroupingStrategy = 3
	UNumberGroupingStrategyUNUMGROUPINGTHOUSANDS UNumberGroupingStrategy = 4
	UNumberGroupingStrategyUNUMGROUPINGCOUNT     UNumberGroupingStrategy = 5
)
