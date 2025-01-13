package libxml_2_0

import _ "unsafe"
/**
 * \file
 * \brief C API: Encoding/charset encoding selector
 *
 * A converter selector is built with a set of encoding/charset names
 * and given an input string returns the set of names of the
 * corresponding converters which can convert the string.
 *
 * A converter selector can be serialized into a buffer and reopened
 * from the serialized form.
 */

type UConverterSelector struct {
	Unused [8]uint8
}
