package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type USimpleNumberSign c.Int

const (
	USimpleNumberSignUNUMSIMPLENUMBERPLUSSIGN  USimpleNumberSign = 0
	USimpleNumberSignUNUMSIMPLENUMBERNOSIGN    USimpleNumberSign = 1
	USimpleNumberSignUNUMSIMPLENUMBERMINUSSIGN USimpleNumberSign = 2
)

type USimpleNumber struct {
	Unused [8]uint8
}

type USimpleNumberFormatter struct {
	Unused [8]uint8
}
