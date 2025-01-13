package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type URegionType c.Int

const (
	URegionTypeURGNUNKNOWN      URegionType = 0
	URegionTypeURGNTERRITORY    URegionType = 1
	URegionTypeURGNWORLD        URegionType = 2
	URegionTypeURGNCONTINENT    URegionType = 3
	URegionTypeURGNSUBCONTINENT URegionType = 4
	URegionTypeURGNGROUPING     URegionType = 5
	URegionTypeURGNDEPRECATED   URegionType = 6
	URegionTypeURGNLIMIT        URegionType = 7
)
/**
 * Opaque URegion object for use in C programs.
 * @stable ICU 52
 */

type URegion struct {
	Unused [8]uint8
}
