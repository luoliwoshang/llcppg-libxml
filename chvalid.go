package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type X_XmlChSRange struct {
	Low  uint16
	High uint16
}
type XmlChSRange X_XmlChSRange
type XmlChSRangePtr *XmlChSRange

type X_XmlChLRange struct {
	Low  c.Uint
	High c.Uint
}
type XmlChLRange X_XmlChLRange
type XmlChLRangePtr *XmlChLRange

type X_XmlChRangeGroup struct {
	NbShortRange c.Int
	NbLongRange  c.Int
	ShortRange   *XmlChSRange
	LongRange    *XmlChLRange
}
type XmlChRangeGroup X_XmlChRangeGroup
type XmlChRangeGroupPtr *XmlChRangeGroup
/**
 * Range checking routine
 */
//go:linkname XmlCharInRange C.xmlCharInRange
func XmlCharInRange(val c.Uint, group *XmlChRangeGroup) c.Int
//go:linkname XmlIsBaseChar C.xmlIsBaseChar
func XmlIsBaseChar(ch c.Uint) c.Int
//go:linkname XmlIsBlank C.xmlIsBlank
func XmlIsBlank(ch c.Uint) c.Int
//go:linkname XmlIsChar C.xmlIsChar
func XmlIsChar(ch c.Uint) c.Int
//go:linkname XmlIsCombining C.xmlIsCombining
func XmlIsCombining(ch c.Uint) c.Int
//go:linkname XmlIsDigit C.xmlIsDigit
func XmlIsDigit(ch c.Uint) c.Int
//go:linkname XmlIsExtender C.xmlIsExtender
func XmlIsExtender(ch c.Uint) c.Int
//go:linkname XmlIsIdeographic C.xmlIsIdeographic
func XmlIsIdeographic(ch c.Uint) c.Int
//go:linkname XmlIsPubidChar C.xmlIsPubidChar
func XmlIsPubidChar(ch c.Uint) c.Int
