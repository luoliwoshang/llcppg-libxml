package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const UCNVSUBSTOPONILLEGAL string = "i"
const UCNVSKIPSTOPONILLEGAL string = "i"
const UCNVESCAPEJAVA string = "J"
const UCNVESCAPEC string = "C"
const UCNVESCAPEXMLDEC string = "D"
const UCNVESCAPEXMLHEX string = "X"
const UCNVESCAPEUNICODE string = "U"
const UCNVESCAPECSS2 string = "S"
/** Forward declaring the UConverter structure. @stable ICU 2.0 */

type UConverter struct {
	Unused [8]uint8
}
type UConverterCallbackReason c.Int

const (
	UConverterCallbackReasonUCNVUNASSIGNED UConverterCallbackReason = 0
	UConverterCallbackReasonUCNVILLEGAL    UConverterCallbackReason = 1
	UConverterCallbackReasonUCNVIRREGULAR  UConverterCallbackReason = 2
	UConverterCallbackReasonUCNVRESET      UConverterCallbackReason = 3
	UConverterCallbackReasonUCNVCLOSE      UConverterCallbackReason = 4
	UConverterCallbackReasonUCNVCLONE      UConverterCallbackReason = 5
)
/**
 * The structure for the fromUnicode callback function parameter.
 * @stable ICU 2.0
 */

type UConverterFromUnicodeArgs struct {
	Size        uint16
	Flush       UBool
	Converter   *UConverter
	Source      *UChar
	SourceLimit *UChar
	Target      *int8
	TargetLimit *int8
	Offsets     *int32
}
/**
 * The structure for the toUnicode callback function parameter.
 * @stable ICU 2.0
 */

type UConverterToUnicodeArgs struct {
	Size        uint16
	Flush       UBool
	Converter   *UConverter
	Source      *int8
	SourceLimit *int8
	Target      *UChar
	TargetLimit *UChar
	Offsets     *int32
}
/**
 * DO NOT CALL THIS FUNCTION DIRECTLY!
 * This From Unicode callback STOPS at the ILLEGAL_SEQUENCE,
 * returning the error code back to the caller immediately.
 *
 * @param context Pointer to the callback's private data
 * @param fromUArgs Information about the conversion in progress
 * @param codeUnits Points to 'length' UChars of the concerned Unicode sequence
 * @param length Size (in bytes) of the concerned codepage sequence
 * @param codePoint Single UChar32 (UTF-32) containing the concerend Unicode codepoint.
 * @param reason Defines the reason the callback was invoked
 * @param err This should always be set to a failure status prior to calling.
 * @stable ICU 2.0
 */
//go:linkname UCNVFROMUCALLBACKSTOP76 C.UCNV_FROM_U_CALLBACK_STOP_76
func UCNVFROMUCALLBACKSTOP76(context unsafe.Pointer, fromUArgs *UConverterFromUnicodeArgs, codeUnits *UChar, length int32, codePoint UChar32, reason UConverterCallbackReason, err *UErrorCode)
/**
 * DO NOT CALL THIS FUNCTION DIRECTLY!
 * This To Unicode callback STOPS at the ILLEGAL_SEQUENCE,
 * returning the error code back to the caller immediately.
 *
 * @param context Pointer to the callback's private data
 * @param toUArgs Information about the conversion in progress
 * @param codeUnits Points to 'length' bytes of the concerned codepage sequence
 * @param length Size (in bytes) of the concerned codepage sequence
 * @param reason Defines the reason the callback was invoked
 * @param err This should always be set to a failure status prior to calling.
 * @stable ICU 2.0
 */
//go:linkname UCNVTOUCALLBACKSTOP76 C.UCNV_TO_U_CALLBACK_STOP_76
func UCNVTOUCALLBACKSTOP76(context unsafe.Pointer, toUArgs *UConverterToUnicodeArgs, codeUnits *int8, length int32, reason UConverterCallbackReason, err *UErrorCode)
