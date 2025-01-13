package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const UCNVMAXCONVERTERNAMELENGTH c.Int = 60
const UCNVSI c.Int = 15
const UCNVSO c.Int = 14
const UCNVOPTIONSEPCHAR string = ","
const UCNVOPTIONSEPSTRING string = ","
const UCNVVALUESEPCHAR string = "="
const UCNVVALUESEPSTRING string = "="
const UCNVLOCALEOPTIONSTRING string = ",locale="
const UCNVVERSIONOPTIONSTRING string = ",version="
const UCNVSWAPLFNLOPTIONSTRING string = ",swaplfnl"
const UCNVSAFECLONEBUFFERSIZE c.Int = 1024

type USet struct {
	Unused [8]uint8
}
type UConverterType c.Int

const (
	UConverterTypeUCNVUNSUPPORTEDCONVERTER            UConverterType = -1
	UConverterTypeUCNVSBCS                            UConverterType = 0
	UConverterTypeUCNVDBCS                            UConverterType = 1
	UConverterTypeUCNVMBCS                            UConverterType = 2
	UConverterTypeUCNVLATIN1                          UConverterType = 3
	UConverterTypeUCNVUTF8                            UConverterType = 4
	UConverterTypeUCNVUTF16BigEndian                  UConverterType = 5
	UConverterTypeUCNVUTF16LittleEndian               UConverterType = 6
	UConverterTypeUCNVUTF32BigEndian                  UConverterType = 7
	UConverterTypeUCNVUTF32LittleEndian               UConverterType = 8
	UConverterTypeUCNVEBCDICSTATEFUL                  UConverterType = 9
	UConverterTypeUCNVISO2022                         UConverterType = 10
	UConverterTypeUCNVLMBCS1                          UConverterType = 11
	UConverterTypeUCNVLMBCS2                          UConverterType = 12
	UConverterTypeUCNVLMBCS3                          UConverterType = 13
	UConverterTypeUCNVLMBCS4                          UConverterType = 14
	UConverterTypeUCNVLMBCS5                          UConverterType = 15
	UConverterTypeUCNVLMBCS6                          UConverterType = 16
	UConverterTypeUCNVLMBCS8                          UConverterType = 17
	UConverterTypeUCNVLMBCS11                         UConverterType = 18
	UConverterTypeUCNVLMBCS16                         UConverterType = 19
	UConverterTypeUCNVLMBCS17                         UConverterType = 20
	UConverterTypeUCNVLMBCS18                         UConverterType = 21
	UConverterTypeUCNVLMBCS19                         UConverterType = 22
	UConverterTypeUCNVLMBCSLAST                       UConverterType = 22
	UConverterTypeUCNVHZ                              UConverterType = 23
	UConverterTypeUCNVSCSU                            UConverterType = 24
	UConverterTypeUCNVISCII                           UConverterType = 25
	UConverterTypeUCNVUSASCII                         UConverterType = 26
	UConverterTypeUCNVUTF7                            UConverterType = 27
	UConverterTypeUCNVBOCU1                           UConverterType = 28
	UConverterTypeUCNVUTF16                           UConverterType = 29
	UConverterTypeUCNVUTF32                           UConverterType = 30
	UConverterTypeUCNVCESU8                           UConverterType = 31
	UConverterTypeUCNVIMAPMAILBOX                     UConverterType = 32
	UConverterTypeUCNVCOMPOUNDTEXT                    UConverterType = 33
	UConverterTypeUCNVNUMBEROFSUPPORTEDCONVERTERTYPES UConverterType = 34
)

type UConverterPlatform c.Int

const (
	UConverterPlatformUCNVUNKNOWN UConverterPlatform = -1
	UConverterPlatformUCNVIBM     UConverterPlatform = 0
)
// llgo:type C
type UConverterToUCallback func(unsafe.Pointer, *UConverterToUnicodeArgs, *int8, int32, UConverterCallbackReason, *UErrorCode)
// llgo:type C
type UConverterFromUCallback func(unsafe.Pointer, *UConverterFromUnicodeArgs, *UChar, int32, UChar32, UConverterCallbackReason, *UErrorCode)
/**
 * Creates a UConverter object with the name of a coded character set specified as a C string.
 * The actual name will be resolved with the alias file
 * using a case-insensitive string comparison that ignores
 * leading zeroes and all non-alphanumeric characters.
 * E.g., the names "UTF8", "utf-8", "u*T@f08" and "Utf 8" are all equivalent.
 * (See also ucnv_compareNames().)
 * If <code>NULL</code> is passed for the converter name, it will create one with the
 * getDefaultName return value.
 *
 * <p>A converter name for ICU 1.5 and above may contain options
 * like a locale specification to control the specific behavior of
 * the newly instantiated converter.
 * The meaning of the options depends on the particular converter.
 * If an option is not defined for or recognized by a given converter, then it is ignored.</p>
 *
 * <p>Options are appended to the converter name string, with a
 * <code>UCNV_OPTION_SEP_CHAR</code> between the name and the first option and
 * also between adjacent options.</p>
 *
 * <p>If the alias is ambiguous, then the preferred converter is used
 * and the status is set to U_AMBIGUOUS_ALIAS_WARNING.</p>
 *
 * <p>The conversion behavior and names can vary between platforms. ICU may
 * convert some characters differently from other platforms. Details on this topic
 * are in the <a href="https://unicode-org.github.io/icu/userguide/conversion/">User
 * Guide</a>. Aliases starting with a "cp" prefix have no specific meaning
 * other than its an alias starting with the letters "cp". Please do not
 * associate any meaning to these aliases.</p>
 *
 * @param converterName Name of the coded character set table.
 *          This may have options appended to the string.
 *          IANA alias character set names, IBM CCSIDs starting with "ibm-",
 *          Windows codepage numbers starting with "windows-" are frequently
 *          used for this parameter. See ucnv_getAvailableName and
 *          ucnv_getAlias for a complete list that is available.
 *          If this parameter is NULL, the default converter will be used.
 * @param err outgoing error status <TT>U_MEMORY_ALLOCATION_ERROR, U_FILE_ACCESS_ERROR</TT>
 * @return the created Unicode converter object, or <TT>NULL</TT> if an error occurred
 * @see ucnv_openU
 * @see ucnv_openCCSID
 * @see ucnv_getAvailableName
 * @see ucnv_getAlias
 * @see ucnv_getDefaultName
 * @see ucnv_close
 * @see ucnv_compareNames
 * @stable ICU 2.0
 */
//go:linkname UcnvOpen76 C.ucnv_open_76
func UcnvOpen76(converterName *int8, err *UErrorCode) *UConverter
/**
 * Deletes the unicode converter and releases resources associated
 * with just this instance.
 * Does not free up shared converter tables.
 *
 * @param converter the converter object to be deleted
 * @see ucnv_open
 * @see ucnv_openU
 * @see ucnv_openCCSID
 * @stable ICU 2.0
 */
// llgo:link (*UConverter).UcnvClose76 C.ucnv_close_76
func (recv_ *UConverter) UcnvClose76() {
}

type UConverterUnicodeSet c.Int

const (
	UConverterUnicodeSetUCNVROUNDTRIPSET            UConverterUnicodeSet = 0
	UConverterUnicodeSetUCNVROUNDTRIPANDFALLBACKSET UConverterUnicodeSet = 1
	UConverterUnicodeSetUCNVSETCOUNT                UConverterUnicodeSet = 2
)
/**
 * Changes the callback function used by the converter when
 * an illegal or invalid sequence is found.
 * Context pointers are always owned by the caller.
 * Predefined actions and contexts can be found in the ucnv_err.h header.
 *
 * @param converter the unicode converter
 * @param newAction the new callback function
 * @param newContext the new toUnicode callback context pointer. This can be NULL.
 * @param oldAction fillin: returns the old callback function pointer. This can be NULL.
 * @param oldContext fillin: returns the old callback's private void* context. This can be NULL.
 * @param err The error code status
 * @see ucnv_getToUCallBack
 * @stable ICU 2.0
 */
// llgo:link (*UConverter).UcnvSetToUCallBack76 C.ucnv_setToUCallBack_76
func (recv_ *UConverter) UcnvSetToUCallBack76(newAction UConverterToUCallback, newContext unsafe.Pointer, oldAction UConverterToUCallback, oldContext *unsafe.Pointer, err *UErrorCode) {
}
/**
 * Changes the current callback function used by the converter when
 * an illegal or invalid sequence is found.
 * Context pointers are always owned by the caller.
 * Predefined actions and contexts can be found in the ucnv_err.h header.
 *
 * @param converter the unicode converter
 * @param newAction the new callback function
 * @param newContext the new fromUnicode callback context pointer. This can be NULL.
 * @param oldAction fillin: returns the old callback function pointer. This can be NULL.
 * @param oldContext fillin: returns the old callback's private void* context. This can be NULL.
 * @param err The error code status
 * @see ucnv_getFromUCallBack
 * @stable ICU 2.0
 */
// llgo:link (*UConverter).UcnvSetFromUCallBack76 C.ucnv_setFromUCallBack_76
func (recv_ *UConverter) UcnvSetFromUCallBack76(newAction UConverterFromUCallback, newContext unsafe.Pointer, oldAction UConverterFromUCallback, oldContext *unsafe.Pointer, err *UErrorCode) {
}
/**
 * Convert from one external charset to another using two existing UConverters.
 * Internally, two conversions - ucnv_toUnicode() and ucnv_fromUnicode() -
 * are used, "pivoting" through 16-bit Unicode.
 *
 * Important: For streaming conversion (multiple function calls for successive
 * parts of a text stream), the caller must provide a pivot buffer explicitly,
 * and must preserve the pivot buffer and associated pointers from one
 * call to another. (The buffer may be moved if its contents and the relative
 * pointer positions are preserved.)
 *
 * There is a similar function, ucnv_convert(),
 * which has the following limitations:
 * - it takes charset names, not converter objects, so that
 *   - two converters are opened for each call
 *   - only single-string conversion is possible, not streaming operation
 * - it does not provide enough information to find out,
 *   in case of failure, whether the toUnicode or
 *   the fromUnicode conversion failed
 *
 * By contrast, ucnv_convertEx()
 * - takes UConverter parameters instead of charset names
 * - fully exposes the pivot buffer for streaming conversion and complete error handling
 *
 * ucnv_convertEx() also provides further convenience:
 * - an option to reset the converters at the beginning
 *   (if reset==true, see parameters;
 *    also sets *pivotTarget=*pivotSource=pivotStart)
 * - allow NUL-terminated input
 *   (only a single NUL byte, will not work for charsets with multi-byte NULs)
 *   (if sourceLimit==NULL, see parameters)
 * - terminate with a NUL on output
 *   (only a single NUL byte, not useful for charsets with multi-byte NULs),
 *   or set U_STRING_NOT_TERMINATED_WARNING if the output exactly fills
 *   the target buffer
 * - the pivot buffer can be provided internally;
 *   possible only for whole-string conversion, not streaming conversion;
 *   in this case, the caller will not be able to get details about where an
 *   error occurred
 *   (if pivotStart==NULL, see below)
 *
 * The function returns when one of the following is true:
 * - the entire source text has been converted successfully to the target buffer
 * - a target buffer overflow occurred (U_BUFFER_OVERFLOW_ERROR)
 * - a conversion error occurred
 *   (other U_FAILURE(), see description of pErrorCode)
 *
 * Limitation compared to the direct use of
 * ucnv_fromUnicode() and ucnv_toUnicode():
 * ucnv_convertEx() does not provide offset information.
 *
 * Limitation compared to ucnv_fromUChars() and ucnv_toUChars():
 * ucnv_convertEx() does not support preflighting directly.
 *
 * Sample code for converting a single string from
 * one external charset to UTF-8, ignoring the location of errors:
 *
 * \code
 * int32_t
 * myToUTF8(UConverter *cnv,
 *          const char *s, int32_t length,
 *          char *u8, int32_t capacity,
 *          UErrorCode *pErrorCode) {
 *     UConverter *utf8Cnv;
 *     char *target;
 *
 *     if(U_FAILURE(*pErrorCode)) {
 *         return 0;
 *     }
 *
 *     utf8Cnv=myGetCachedUTF8Converter(pErrorCode);
 *     if(U_FAILURE(*pErrorCode)) {
 *         return 0;
 *     }
 *
 *     if(length<0) {
 *         length=strlen(s);
 *     }
 *     target=u8;
 *     ucnv_convertEx(utf8Cnv, cnv,
 *                    &target, u8+capacity,
 *                    &s, s+length,
 *                    NULL, NULL, NULL, NULL,
 *                    true, true,
 *                    pErrorCode);
 *
 *     myReleaseCachedUTF8Converter(utf8Cnv);
 *
 *     // return the output string length, but without preflighting
 *     return (int32_t)(target-u8);
 * }
 * \endcode
 *
 * @param targetCnv     Output converter, used to convert from the UTF-16 pivot
 *                      to the target using ucnv_fromUnicode().
 * @param sourceCnv     Input converter, used to convert from the source to
 *                      the UTF-16 pivot using ucnv_toUnicode().
 * @param target        I/O parameter, same as for ucnv_fromUChars().
 *                      Input: *target points to the beginning of the target buffer.
 *                      Output: *target points to the first unit after the last char written.
 * @param targetLimit   Pointer to the first unit after the target buffer.
 * @param source        I/O parameter, same as for ucnv_toUChars().
 *                      Input: *source points to the beginning of the source buffer.
 *                      Output: *source points to the first unit after the last char read.
 * @param sourceLimit   Pointer to the first unit after the source buffer.
 * @param pivotStart    Pointer to the UTF-16 pivot buffer. If pivotStart==NULL,
 *                      then an internal buffer is used and the other pivot
 *                      arguments are ignored and can be NULL as well.
 * @param pivotSource   I/O parameter, same as source in ucnv_fromUChars() for
 *                      conversion from the pivot buffer to the target buffer.
 * @param pivotTarget   I/O parameter, same as target in ucnv_toUChars() for
 *                      conversion from the source buffer to the pivot buffer.
 *                      It must be pivotStart<=*pivotSource<=*pivotTarget<=pivotLimit
 *                      and pivotStart<pivotLimit (unless pivotStart==NULL).
 * @param pivotLimit    Pointer to the first unit after the pivot buffer.
 * @param reset         If true, then ucnv_resetToUnicode(sourceCnv) and
 *                      ucnv_resetFromUnicode(targetCnv) are called, and the
 *                      pivot pointers are reset (*pivotTarget=*pivotSource=pivotStart).
 * @param flush         If true, indicates the end of the input.
 *                      Passed directly to ucnv_toUnicode(), and carried over to
 *                      ucnv_fromUnicode() when the source is empty as well.
 * @param pErrorCode    ICU error code in/out parameter.
 *                      Must fulfill U_SUCCESS before the function call.
 *                      U_BUFFER_OVERFLOW_ERROR always refers to the target buffer
 *                      because overflows into the pivot buffer are handled internally.
 *                      Other conversion errors are from the source-to-pivot
 *                      conversion if *pivotSource==pivotStart, otherwise from
 *                      the pivot-to-target conversion.
 *
 * @see ucnv_convert
 * @see ucnv_fromAlgorithmic
 * @see ucnv_toAlgorithmic
 * @see ucnv_fromUnicode
 * @see ucnv_toUnicode
 * @see ucnv_fromUChars
 * @see ucnv_toUChars
 * @stable ICU 2.6
 */
// llgo:link (*UConverter).UcnvConvertEx76 C.ucnv_convertEx_76
func (recv_ *UConverter) UcnvConvertEx76(sourceCnv *UConverter, target **int8, targetLimit *int8, source **int8, sourceLimit *int8, pivotStart *UChar, pivotSource **UChar, pivotTarget **UChar, pivotLimit *UChar, reset UBool, flush UBool, pErrorCode *UErrorCode) {
}
