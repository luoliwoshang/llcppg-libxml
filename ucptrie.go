package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)
/** @internal */

type UCPTrieData struct {
	Ptr0 unsafe.Pointer
}
/**
 * Immutable Unicode code point trie structure.
 * Fast, reasonably compact, map from Unicode code points (U+0000..U+10FFFF) to integer values.
 * For details see https://icu.unicode.org/design/struct/utrie
 *
 * Do not access UCPTrie fields directly; use public functions and macros.
 * Functions are easy to use: They support all trie types and value widths.
 *
 * When performance is really important, macros provide faster access.
 * Most macros are specific to either "fast" or "small" tries, see UCPTrieType.
 * There are "fast" macros for special optimized use cases.
 *
 * The macros will return bogus values, or may crash, if used on the wrong type or value width.
 *
 * @see UMutableCPTrie
 * @stable ICU 63
 */

type UCPTrie struct {
	Index              *uint16
	Data               UCPTrieData
	IndexLength        int32
	DataLength         int32
	HighStart          UChar32
	Shifted12HighStart uint16
	Type               int8
	ValueWidth         int8
	Reserved32         uint32
	Reserved16         uint16
	Index3NullOffset   uint16
	DataNullOffset     int32
	NullValue          uint32
}
type UCPTrieType c.Int

const (
	UCPTrieTypeUCPTRIETYPEANY   UCPTrieType = -1
	UCPTrieTypeUCPTRIETYPEFAST  UCPTrieType = 0
	UCPTrieTypeUCPTRIETYPESMALL UCPTrieType = 1
)

type UCPTrieValueWidth c.Int

const (
	UCPTrieValueWidthUCPTRIEVALUEBITSANY UCPTrieValueWidth = -1
	UCPTrieValueWidthUCPTRIEVALUEBITS16  UCPTrieValueWidth = 0
	UCPTrieValueWidthUCPTRIEVALUEBITS32  UCPTrieValueWidth = 1
	UCPTrieValueWidthUCPTRIEVALUEBITS8   UCPTrieValueWidth = 2
)
const (
	UCPTRIEFASTSHIFT               c.Int = 6
	UCPTRIEFASTDATABLOCKLENGTH     c.Int = 64
	UCPTRIEFASTDATAMASK            c.Int = 63
	UCPTRIESMALLMAX                c.Int = 4095
	UCPTRIEERRORVALUENEGDATAOFFSET c.Int = 1
	UCPTRIEHIGHVALUENEGDATAOFFSET  c.Int = 2
)
