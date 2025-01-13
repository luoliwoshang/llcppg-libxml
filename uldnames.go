package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UDialectHandling c.Int

const (
	UDialectHandlingULDNSTANDARDNAMES UDialectHandling = 0
	UDialectHandlingULDNDIALECTNAMES  UDialectHandling = 1
)
/**
 * Opaque C service object type for the locale display names API
 * @stable ICU 4.4
 */

type ULocaleDisplayNames struct {
	Unused [8]uint8
}
