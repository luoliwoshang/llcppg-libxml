package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UBiDiOrder c.Int

const (
	UBiDiOrderUBIDILOGICAL UBiDiOrder = 0
	UBiDiOrderUBIDIVISUAL  UBiDiOrder = 1
)

type UBiDiMirroring c.Int

const (
	UBiDiMirroringUBIDIMIRRORINGOFF UBiDiMirroring = 0
	UBiDiMirroringUBIDIMIRRORINGON  UBiDiMirroring = 1
)

type UBiDiTransform struct {
	Unused [8]uint8
}
