package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const UPARSECONTEXTLEN c.Int = 16
/**
 * A UParseError struct is used to returned detailed information about
 * parsing errors.  It is used by ICU parsing engines that parse long
 * rules, patterns, or programs, where the text being parsed is long
 * enough that more information than a UErrorCode is needed to
 * localize the error.
 *
 * <p>The line, offset, and context fields are optional; parsing
 * engines may choose not to use to use them.
 *
 * <p>The preContext and postContext strings include some part of the
 * context surrounding the error.  If the source text is "let for=7"
 * and "for" is the error (e.g., because it is a reserved word), then
 * some examples of what a parser might produce are the following:
 *
 * <pre>
 * preContext   postContext
 * ""           ""            The parser does not support context
 * "let "       "=7"          Pre- and post-context only
 * "let "       "for=7"       Pre- and post-context and error text
 * ""           "for"         Error text only
 * </pre>
 *
 * <p>Examples of engines which use UParseError (or may use it in the
 * future) are Transliterator, RuleBasedBreakIterator, and
 * RegexPattern.
 *
 * @stable ICU 2.0
 */

type UParseError struct {
	Line        int32
	Offset      int32
	PreContext  [16]UChar
	PostContext [16]UChar
}
