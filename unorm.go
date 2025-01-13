package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const UNORMCOMPARENORMOPTIONSSHIFT c.Int = 20

type UNormalizationMode c.Int

const (
	UNormalizationModeUNORMNONE      UNormalizationMode = 1
	UNormalizationModeUNORMNFD       UNormalizationMode = 2
	UNormalizationModeUNORMNFKD      UNormalizationMode = 3
	UNormalizationModeUNORMNFC       UNormalizationMode = 4
	UNormalizationModeUNORMDEFAULT   UNormalizationMode = 4
	UNormalizationModeUNORMNFKC      UNormalizationMode = 5
	UNormalizationModeUNORMFCD       UNormalizationMode = 6
	UNormalizationModeUNORMMODECOUNT UNormalizationMode = 7
)
const UNORMUNICODE32 c.Int = 32
