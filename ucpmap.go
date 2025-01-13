package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type UCPMap struct {
	Unused [8]uint8
}
type UCPMapRangeOption c.Int

const (
	UCPMapRangeOptionUCPMAPRANGENORMAL              UCPMapRangeOption = 0
	UCPMapRangeOptionUCPMAPRANGEFIXEDLEADSURROGATES UCPMapRangeOption = 1
	UCPMapRangeOptionUCPMAPRANGEFIXEDALLSURROGATES  UCPMapRangeOption = 2
)
// llgo:type C
type UCPMapValueFilter func(unsafe.Pointer, uint32) uint32
