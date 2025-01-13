package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const (
	UIDNADEFAULT                  c.Int = 48
	UIDNAALLOWUNASSIGNED          c.Int = 1
	UIDNAUSESTD3RULES             c.Int = 2
	UIDNACHECKBIDI                c.Int = 4
	UIDNACHECKCONTEXTJ            c.Int = 8
	UIDNANONTRANSITIONALTOASCII   c.Int = 16
	UIDNANONTRANSITIONALTOUNICODE c.Int = 32
	UIDNACHECKCONTEXTO            c.Int = 64
)
/**
 * Opaque C service object type for the new IDNA API.
 * @stable ICU 4.6
 */

type UIDNA struct {
	Unused [8]uint8
}
/**
 * Output container for IDNA processing errors.
 * Initialize with UIDNA_INFO_INITIALIZER:
 * \code
 * UIDNAInfo info = UIDNA_INFO_INITIALIZER;
 * int32_t length = uidna_nameToASCII(..., &info, &errorCode);
 * if(U_SUCCESS(errorCode) && info.errors!=0) { ... }
 * \endcode
 * @stable ICU 4.6
 */

type UIDNAInfo struct {
	Size                    int16
	IsTransitionalDifferent UBool
	ReservedB3              UBool
	Errors                  uint32
	ReservedI2              int32
	ReservedI3              int32
}

const (
	UIDNAERROREMPTYLABEL           c.Int = 1
	UIDNAERRORLABELTOOLONG         c.Int = 2
	UIDNAERRORDOMAINNAMETOOLONG    c.Int = 4
	UIDNAERRORLEADINGHYPHEN        c.Int = 8
	UIDNAERRORTRAILINGHYPHEN       c.Int = 16
	UIDNAERRORHYPHEN34             c.Int = 32
	UIDNAERRORLEADINGCOMBININGMARK c.Int = 64
	UIDNAERRORDISALLOWED           c.Int = 128
	UIDNAERRORPUNYCODE             c.Int = 256
	UIDNAERRORLABELHASDOT          c.Int = 512
	UIDNAERRORINVALIDACELABEL      c.Int = 1024
	UIDNAERRORBIDI                 c.Int = 2048
	UIDNAERRORCONTEXTJ             c.Int = 4096
	UIDNAERRORCONTEXTOPUNCTUATION  c.Int = 8192
	UIDNAERRORCONTEXTODIGITS       c.Int = 16384
)
