package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const UEOF c.Int = 65535

type UFILE struct {
	Unused [8]uint8
}
type UFileDirection c.Int

const (
	UFileDirectionUREAD      UFileDirection = 1
	UFileDirectionUWRITE     UFileDirection = 2
	UFileDirectionUREADWRITE UFileDirection = 3
)
