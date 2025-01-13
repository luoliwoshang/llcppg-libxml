package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UGender c.Int

const (
	UGenderUGENDERMALE   UGender = 0
	UGenderUGENDERFEMALE UGender = 1
	UGenderUGENDEROTHER  UGender = 2
)

type UGenderInfo struct {
	Unused [8]uint8
}
