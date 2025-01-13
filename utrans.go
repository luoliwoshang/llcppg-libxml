package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type UTransliterator unsafe.Pointer
type UTransDirection c.Int

const (
	UTransDirectionUTRANSFORWARD UTransDirection = 0
	UTransDirectionUTRANSREVERSE UTransDirection = 1
)
/**
 * Position structure for utrans_transIncremental() incremental
 * transliteration.  This structure defines two substrings of the text
 * being transliterated.  The first region, [contextStart,
 * contextLimit), defines what characters the transliterator will read
 * as context.  The second region, [start, limit), defines what
 * characters will actually be transliterated.  The second region
 * should be a subset of the first.
 *
 * <p>After a transliteration operation, some of the indices in this
 * structure will be modified.  See the field descriptions for
 * details.
 *
 * <p>contextStart <= start <= limit <= contextLimit
 *
 * <p>Note: All index values in this structure must be at code point
 * boundaries.  That is, none of them may occur between two code units
 * of a surrogate pair.  If any index does split a surrogate pair,
 * results are unspecified.
 *
 * @stable ICU 2.0
 */

type UTransPosition struct {
	ContextStart int32
	ContextLimit int32
	Start        int32
	Limit        int32
}
