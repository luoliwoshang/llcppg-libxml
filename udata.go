package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const UTREESEPARATOR string = "-"
const UTREESEPARATORSTRING string = "-"
const UTREEENTRYSEPCHAR string = "/"
const UTREEENTRYSEPSTRING string = "/"
const UICUDATAALIAS string = "ICUDATA"
/**
 * UDataInfo contains the properties about the requested data.
 * This is meta data.
 *
 * <p>This structure may grow in the future, indicated by the
 * <code>size</code> field.</p>
 *
 * <p>ICU data must be at least 8-aligned, and should be 16-aligned.
 * The UDataInfo struct begins 4 bytes after the start of the data item,
 * so it is 4-aligned.
 *
 * <p>The platform data property fields help determine if a data
 * file can be efficiently used on a given machine.
 * The particular fields are of importance only if the data
 * is affected by the properties - if there is integer data
 * with word sizes > 1 byte, char* text, or UChar* text.</p>
 *
 * <p>The implementation for the <code>udata_open[Choice]()</code>
 * functions may reject data based on the value in <code>isBigEndian</code>.
 * No other field is used by the <code>udata</code> API implementation.</p>
 *
 * <p>The <code>dataFormat</code> may be used to identify
 * the kind of data, e.g. a converter table.</p>
 *
 * <p>The <code>formatVersion</code> field should be used to
 * make sure that the format can be interpreted.
 * It may be a good idea to check only for the one or two highest
 * of the version elements to allow the data memory to
 * get more or somewhat rearranged contents, for as long
 * as the using code can still interpret the older contents.</p>
 *
 * <p>The <code>dataVersion</code> field is intended to be a
 * common place to store the source version of the data;
 * for data from the Unicode character database, this could
 * reflect the Unicode version.</p>
 *
 * @stable ICU 2.0
 */

type UDataInfo struct {
	Size          uint16
	ReservedWord  uint16
	IsBigEndian   uint8
	CharsetFamily uint8
	SizeofUChar   uint8
	ReservedByte  uint8
	DataFormat    [4]uint8
	FormatVersion [4]uint8
	DataVersion   [4]uint8
}

type UDataMemory struct {
	Unused [8]uint8
}
// llgo:type C
type UDataMemoryIsAcceptable func(unsafe.Pointer, *int8, *int8, *UDataInfo) UBool
type UDataFileAccess c.Int

const (
	UDataFileAccessUDATAFILESFIRST      UDataFileAccess = 0
	UDataFileAccessUDATADEFAULTACCESS   UDataFileAccess = 0
	UDataFileAccessUDATAONLYPACKAGES    UDataFileAccess = 1
	UDataFileAccessUDATAPACKAGESFIRST   UDataFileAccess = 2
	UDataFileAccessUDATANOFILES         UDataFileAccess = 3
	UDataFileAccessUDATAFILEACCESSCOUNT UDataFileAccess = 4
)
