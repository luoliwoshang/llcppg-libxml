package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UFieldCategory c.Int

const (
	UFieldCategoryUFIELDCATEGORYUNDEFINED        UFieldCategory = 0
	UFieldCategoryUFIELDCATEGORYDATE             UFieldCategory = 1
	UFieldCategoryUFIELDCATEGORYNUMBER           UFieldCategory = 2
	UFieldCategoryUFIELDCATEGORYLIST             UFieldCategory = 3
	UFieldCategoryUFIELDCATEGORYRELATIVEDATETIME UFieldCategory = 4
	UFieldCategoryUFIELDCATEGORYDATEINTERVAL     UFieldCategory = 5
	UFieldCategoryUFIELDCATEGORYCOUNT            UFieldCategory = 6
	UFieldCategoryUFIELDCATEGORYLISTSPAN         UFieldCategory = 4099
	UFieldCategoryUFIELDCATEGORYDATEINTERVALSPAN UFieldCategory = 4101
	UFieldCategoryUFIELDCATEGORYNUMBERRANGESPAN  UFieldCategory = 4098
)

type UConstrainedFieldPosition struct {
	Unused [8]uint8
}

type UFormattedValue struct {
	Unused [8]uint8
}
