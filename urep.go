package libxml_2_0

import "unsafe"

type UReplaceable unsafe.Pointer
/**
 * A set of function pointers that transliterators use to manipulate a
 * UReplaceable.  The caller should supply the required functions to
 * manipulate their text appropriately.  Related to the C++ class
 * Replaceable.
 * @stable ICU 2.0
 */

type UReplaceableCallbacks struct {
	Length   unsafe.Pointer
	CharAt   unsafe.Pointer
	Char32At unsafe.Pointer
	Replace  unsafe.Pointer
	Extract  unsafe.Pointer
	Copy     unsafe.Pointer
}
