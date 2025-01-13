package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
/**
 * \file
 * \brief C API: Unicode Security and Spoofing Detection
 *
 * <p>
 * This class, based on <a href="http://unicode.org/reports/tr36">Unicode Technical Report #36</a> and
 * <a href="http://unicode.org/reports/tr39">Unicode Technical Standard #39</a>, has two main functions:
 *
 * <ol>
 * <li>Checking whether two strings are visually <em>confusable</em> with each other, such as "Harvest" and
 * &quot;&Eta;arvest&quot;, where the second string starts with the Greek capital letter Eta.</li>
 * <li>Checking whether an individual string is likely to be an attempt at confusing the reader (<em>spoof
 * detection</em>), such as "paypal" with some Latin characters substituted with Cyrillic look-alikes.</li>
 * </ol>
 *
 * <p>
 * Although originally designed as a method for flagging suspicious identifier strings such as URLs,
 * <code>USpoofChecker</code> has a number of other practical use cases, such as preventing attempts to evade bad-word
 * content filters.
 *
 * <p>
 * The functions of this class are exposed as C API, with a handful of syntactical conveniences for C++.
 *
 * <h2>Confusables</h2>
 *
 * <p>
 * The following example shows how to use <code>USpoofChecker</code> to check for confusability between two strings:
 *
 * \code{.c}
 * UErrorCode status = U_ZERO_ERROR;
 * UChar* str1 = (UChar*) u"Harvest";
 * UChar* str2 = (UChar*) u"\u0397arvest";  // with U+0397 GREEK CAPITAL LETTER ETA
 *
 * USpoofChecker* sc = uspoof_open(&status);
 * uspoof_setChecks(sc, USPOOF_CONFUSABLE, &status);
 *
 * int32_t bitmask = uspoof_areConfusable(sc, str1, -1, str2, -1, &status);
 * UBool result = bitmask != 0;
 * // areConfusable: 1 (status: U_ZERO_ERROR)
 * printf("areConfusable: %d (status: %s)\n", result, u_errorName(status));
 * uspoof_close(sc);
 * \endcode
 *
 * <p>
 * The call to {@link uspoof_open} creates a <code>USpoofChecker</code> object; the call to {@link uspoof_setChecks}
 * enables confusable checking and disables all other checks; the call to {@link uspoof_areConfusable} performs the
 * confusability test; and the following line extracts the result out of the return value. For best performance,
 * the instance should be created once (e.g., upon application startup), and the efficient
 * {@link uspoof_areConfusable} method can be used at runtime.
 *
 * If the paragraph direction used to display the strings is known, the bidi function should be used instead:
 *
 * \code{.c}
 * UErrorCode status = U_ZERO_ERROR;
 * // These strings look identical when rendered in a left-to-right context.
 * // They look distinct in a right-to-left context.
 * UChar* str1 = (UChar*) u"A1\u05D0";  // A1א
 * UChar* str2 = (UChar*) u"A\u05D01";  // Aא1
 *
 * USpoofChecker* sc = uspoof_open(&status);
 * uspoof_setChecks(sc, USPOOF_CONFUSABLE, &status);
 *
 * int32_t bitmask = uspoof_areBidiConfusable(sc, UBIDI_LTR, str1, -1, str2, -1, &status);
 * UBool result = bitmask != 0;
 * // areBidiConfusable: 1 (status: U_ZERO_ERROR)
 * printf("areBidiConfusable: %d (status: %s)\n", result, u_errorName(status));
 * uspoof_close(sc);
 * \endcode
 *
 * <p>
 * The type {@link LocalUSpoofCheckerPointer} is exposed for C++ programmers.  It will automatically call
 * {@link uspoof_close} when the object goes out of scope:
 *
 * \code{.cpp}
 * UErrorCode status = U_ZERO_ERROR;
 * LocalUSpoofCheckerPointer sc(uspoof_open(&status));
 * uspoof_setChecks(sc.getAlias(), USPOOF_CONFUSABLE, &status);
 * // ...
 * \endcode
 *
 * UTS 39 defines two strings to be <em>confusable</em> if they map to the same <em>skeleton string</em>. A skeleton can
 * be thought of as a "hash code". {@link uspoof_getSkeleton} computes the skeleton for a particular string, so
 * the following snippet is equivalent to the example above:
 *
 * \code{.c}
 * UErrorCode status = U_ZERO_ERROR;
 * UChar* str1 = (UChar*) u"Harvest";
 * UChar* str2 = (UChar*) u"\u0397arvest";  // with U+0397 GREEK CAPITAL LETTER ETA
 *
 * USpoofChecker* sc = uspoof_open(&status);
 * uspoof_setChecks(sc, USPOOF_CONFUSABLE, &status);
 *
 * // Get skeleton 1
 * int32_t skel1Len = uspoof_getSkeleton(sc, 0, str1, -1, NULL, 0, &status);
 * UChar* skel1 = (UChar*) malloc(++skel1Len * sizeof(UChar));
 * status = U_ZERO_ERROR;
 * uspoof_getSkeleton(sc, 0, str1, -1, skel1, skel1Len, &status);
 *
 * // Get skeleton 2
 * int32_t skel2Len = uspoof_getSkeleton(sc, 0, str2, -1, NULL, 0, &status);
 * UChar* skel2 = (UChar*) malloc(++skel2Len * sizeof(UChar));
 * status = U_ZERO_ERROR;
 * uspoof_getSkeleton(sc, 0, str2, -1, skel2, skel2Len, &status);
 *
 * // Are the skeletons the same?
 * UBool result = u_strcmp(skel1, skel2) == 0;
 * // areConfusable: 1 (status: U_ZERO_ERROR)
 * printf("areConfusable: %d (status: %s)\n", result, u_errorName(status));
 * uspoof_close(sc);
 * free(skel1);
 * free(skel2);
 * \endcode
 *
 * If you need to check if a string is confusable with any string in a dictionary of many strings, rather than calling
 * {@link uspoof_areConfusable} many times in a loop, {@link uspoof_getSkeleton} can be used instead, as shown below:
 *
 * \code{.c}
 * UErrorCode status = U_ZERO_ERROR;
 * #define DICTIONARY_LENGTH 2
 * UChar* dictionary[DICTIONARY_LENGTH] = { (UChar*) u"lorem", (UChar*) u"ipsum" };
 * UChar* skeletons[DICTIONARY_LENGTH];
 * UChar* str = (UChar*) u"1orern";
 *
 * // Setup:
 * USpoofChecker* sc = uspoof_open(&status);
 * uspoof_setChecks(sc, USPOOF_CONFUSABLE, &status);
 * for (size_t i=0; i<DICTIONARY_LENGTH; i++) {
 *     UChar* word = dictionary[i];
 *     int32_t len = uspoof_getSkeleton(sc, 0, word, -1, NULL, 0, &status);
 *     skeletons[i] = (UChar*) malloc(++len * sizeof(UChar));
 *     status = U_ZERO_ERROR;
 *     uspoof_getSkeleton(sc, 0, word, -1, skeletons[i], len, &status);
 * }
 *
 * // Live Check:
 * {
 *     int32_t len = uspoof_getSkeleton(sc, 0, str, -1, NULL, 0, &status);
 *     UChar* skel = (UChar*) malloc(++len * sizeof(UChar));
 *     status = U_ZERO_ERROR;
 *     uspoof_getSkeleton(sc, 0, str, -1, skel, len, &status);
 *     UBool result = false;
 *     for (size_t i=0; i<DICTIONARY_LENGTH; i++) {
 *         result = u_strcmp(skel, skeletons[i]) == 0;
 *         if (result == true) { break; }
 *     }
 *     // Has confusable in dictionary: 1 (status: U_ZERO_ERROR)
 *     printf("Has confusable in dictionary: %d (status: %s)\n", result, u_errorName(status));
 *     free(skel);
 * }
 *
 * for (size_t i=0; i<DICTIONARY_LENGTH; i++) {
 *     free(skeletons[i]);
 * }
 * uspoof_close(sc);
 * \endcode
 *
 * <b>Note:</b> Since the Unicode confusables mapping table is frequently updated, confusable skeletons are <em>not</em>
 * guaranteed to be the same between ICU releases. We therefore recommend that you always compute confusable skeletons
 * at runtime and do not rely on creating a permanent, or difficult to update, database of skeletons.
 *
 * <h2>Spoof Detection</h2>
 *
 * The following snippet shows a minimal example of using <code>USpoofChecker</code> to perform spoof detection on a
 * string:
 *
 * \code{.c}
 * UErrorCode status = U_ZERO_ERROR;
 * UChar* str = (UChar*) u"p\u0430ypal";  // with U+0430 CYRILLIC SMALL LETTER A
 *
 * // Get the default set of allowable characters:
 * USet* allowed = uset_openEmpty();
 * uset_addAll(allowed, uspoof_getRecommendedSet(&status));
 * uset_addAll(allowed, uspoof_getInclusionSet(&status));
 *
 * USpoofChecker* sc = uspoof_open(&status);
 * uspoof_setAllowedChars(sc, allowed, &status);
 * uspoof_setRestrictionLevel(sc, USPOOF_MODERATELY_RESTRICTIVE);
 *
 * int32_t bitmask = uspoof_check(sc, str, -1, NULL, &status);
 * UBool result = bitmask != 0;
 * // fails checks: 1 (status: U_ZERO_ERROR)
 * printf("fails checks: %d (status: %s)\n", result, u_errorName(status));
 * uspoof_close(sc);
 * uset_close(allowed);
 * \endcode
 *
 * As in the case for confusability checking, it is good practice to create one <code>USpoofChecker</code> instance at
 * startup, and call the cheaper {@link uspoof_check} online. We specify the set of
 * allowed characters to be those with type RECOMMENDED or INCLUSION, according to the recommendation in UTS 39.
 *
 * In addition to {@link uspoof_check}, the function {@link uspoof_checkUTF8} is exposed for UTF8-encoded char* strings,
 * and {@link uspoof_checkUnicodeString} is exposed for C++ programmers.
 *
 * If the {@link USPOOF_AUX_INFO} check is enabled, a limited amount of information on why a string failed the checks
 * is available in the returned bitmask.  For complete information, use the {@link uspoof_check2} class of functions
 * with a {@link USpoofCheckResult} parameter:
 *
 * \code{.c}
 * UErrorCode status = U_ZERO_ERROR;
 * UChar* str = (UChar*) u"p\u0430ypal";  // with U+0430 CYRILLIC SMALL LETTER A
 *
 * // Get the default set of allowable characters:
 * USet* allowed = uset_openEmpty();
 * uset_addAll(allowed, uspoof_getRecommendedSet(&status));
 * uset_addAll(allowed, uspoof_getInclusionSet(&status));
 *
 * USpoofChecker* sc = uspoof_open(&status);
 * uspoof_setAllowedChars(sc, allowed, &status);
 * uspoof_setRestrictionLevel(sc, USPOOF_MODERATELY_RESTRICTIVE);
 *
 * USpoofCheckResult* checkResult = uspoof_openCheckResult(&status);
 * int32_t bitmask = uspoof_check2(sc, str, -1, checkResult, &status);
 *
 * int32_t failures1 = bitmask;
 * int32_t failures2 = uspoof_getCheckResultChecks(checkResult, &status);
 * assert(failures1 == failures2);
 * // checks that failed: 0x00000010 (status: U_ZERO_ERROR)
 * printf("checks that failed: %#010x (status: %s)\n", failures1, u_errorName(status));
 *
 * // Cleanup:
 * uspoof_close(sc);
 * uset_close(allowed);
 * uspoof_closeCheckResult(checkResult);
 * \endcode
 *
 * C++ users can take advantage of a few syntactical conveniences.  The following snippet is functionally
 * equivalent to the one above:
 *
 * \code{.cpp}
 * UErrorCode status = U_ZERO_ERROR;
 * UnicodeString str((UChar*) u"p\u0430ypal");  // with U+0430 CYRILLIC SMALL LETTER A
 *
 * // Get the default set of allowable characters:
 * UnicodeSet allowed;
 * allowed.addAll(*uspoof_getRecommendedUnicodeSet(&status));
 * allowed.addAll(*uspoof_getInclusionUnicodeSet(&status));
 *
 * LocalUSpoofCheckerPointer sc(uspoof_open(&status));
 * uspoof_setAllowedChars(sc.getAlias(), allowed.toUSet(), &status);
 * uspoof_setRestrictionLevel(sc.getAlias(), USPOOF_MODERATELY_RESTRICTIVE);
 *
 * LocalUSpoofCheckResultPointer checkResult(uspoof_openCheckResult(&status));
 * int32_t bitmask = uspoof_check2UnicodeString(sc.getAlias(), str, checkResult.getAlias(), &status);
 *
 * int32_t failures1 = bitmask;
 * int32_t failures2 = uspoof_getCheckResultChecks(checkResult.getAlias(), &status);
 * assert(failures1 == failures2);
 * // checks that failed: 0x00000010 (status: U_ZERO_ERROR)
 * printf("checks that failed: %#010x (status: %s)\n", failures1, u_errorName(status));
 *
 * // Explicit cleanup not necessary.
 * \endcode
 *
 * The return value is a bitmask of the checks that failed. In this case, there was one check that failed:
 * {@link USPOOF_RESTRICTION_LEVEL}, corresponding to the fifth bit (16). The possible checks are:
 *
 * <ul>
 * <li><code>RESTRICTION_LEVEL</code>: flags strings that violate the
 * <a href="http://unicode.org/reports/tr39/#Restriction_Level_Detection">Restriction Level</a> test as specified in UTS
 * 39; in most cases, this means flagging strings that contain characters from multiple different scripts.</li>
 * <li><code>INVISIBLE</code>: flags strings that contain invisible characters, such as zero-width spaces, or character
 * sequences that are likely not to display, such as multiple occurrences of the same non-spacing mark.</li>
 * <li><code>CHAR_LIMIT</code>: flags strings that contain characters outside of a specified set of acceptable
 * characters. See {@link uspoof_setAllowedChars} and {@link uspoof_setAllowedLocales}.</li>
 * <li><code>MIXED_NUMBERS</code>: flags strings that contain digits from multiple different numbering systems.</li>
 * </ul>
 *
 * <p>
 * These checks can be enabled independently of each other. For example, if you were interested in checking for only the
 * INVISIBLE and MIXED_NUMBERS conditions, you could do:
 *
 * \code{.c}
 * UErrorCode status = U_ZERO_ERROR;
 * UChar* str = (UChar*) u"8\u09EA";  // 8 mixed with U+09EA BENGALI DIGIT FOUR
 *
 * USpoofChecker* sc = uspoof_open(&status);
 * uspoof_setChecks(sc, USPOOF_INVISIBLE | USPOOF_MIXED_NUMBERS, &status);
 *
 * int32_t bitmask = uspoof_check2(sc, str, -1, NULL, &status);
 * UBool result = bitmask != 0;
 * // fails checks: 1 (status: U_ZERO_ERROR)
 * printf("fails checks: %d (status: %s)\n", result, u_errorName(status));
 * uspoof_close(sc);
 * \endcode
 *
 * Here is an example in C++ showing how to compute the restriction level of a string:
 *
 * \code{.cpp}
 * UErrorCode status = U_ZERO_ERROR;
 * UnicodeString str((UChar*) u"p\u0430ypal");  // with U+0430 CYRILLIC SMALL LETTER A
 *
 * // Get the default set of allowable characters:
 * UnicodeSet allowed;
 * allowed.addAll(*uspoof_getRecommendedUnicodeSet(&status));
 * allowed.addAll(*uspoof_getInclusionUnicodeSet(&status));
 *
 * LocalUSpoofCheckerPointer sc(uspoof_open(&status));
 * uspoof_setAllowedChars(sc.getAlias(), allowed.toUSet(), &status);
 * uspoof_setRestrictionLevel(sc.getAlias(), USPOOF_MODERATELY_RESTRICTIVE);
 * uspoof_setChecks(sc.getAlias(), USPOOF_RESTRICTION_LEVEL | USPOOF_AUX_INFO, &status);
 *
 * LocalUSpoofCheckResultPointer checkResult(uspoof_openCheckResult(&status));
 * int32_t bitmask = uspoof_check2UnicodeString(sc.getAlias(), str, checkResult.getAlias(), &status);
 *
 * URestrictionLevel restrictionLevel = uspoof_getCheckResultRestrictionLevel(checkResult.getAlias(), &status);
 * // Since USPOOF_AUX_INFO was enabled, the restriction level is also available in the upper bits of the bitmask:
 * assert((restrictionLevel & bitmask) == restrictionLevel);
 * // Restriction level: 0x50000000 (status: U_ZERO_ERROR)
 * printf("Restriction level: %#010x (status: %s)\n", restrictionLevel, u_errorName(status));
 * \endcode
 *
 * The code '0x50000000' corresponds to the restriction level USPOOF_MINIMALLY_RESTRICTIVE.  Since
 * USPOOF_MINIMALLY_RESTRICTIVE is weaker than USPOOF_MODERATELY_RESTRICTIVE, the string fails the check.
 *
 * <b>Note:</b> The Restriction Level is the most powerful of the checks. The full logic is documented in
 * <a href="http://unicode.org/reports/tr39/#Restriction_Level_Detection">UTS 39</a>, but the basic idea is that strings
 * are restricted to contain characters from only a single script, <em>except</em> that most scripts are allowed to have
 * Latin characters interspersed. Although the default restriction level is <code>HIGHLY_RESTRICTIVE</code>, it is
 * recommended that users set their restriction level to <code>MODERATELY_RESTRICTIVE</code>, which allows Latin mixed
 * with all other scripts except Cyrillic, Greek, and Cherokee, with which it is often confusable. For more details on
 * the levels, see UTS 39 or {@link URestrictionLevel}. The Restriction Level test is aware of the set of
 * allowed characters set in {@link uspoof_setAllowedChars}. Note that characters which have script code
 * COMMON or INHERITED, such as numbers and punctuation, are ignored when computing whether a string has multiple
 * scripts.
 *
 * <h2>Advanced bidirectional usage</h2>
 * If the paragraph direction with which the identifiers will be displayed is not known, there are
 * multiple options for confusable detection depending on the circumstances.
 *
 * <p>
 * In some circumstances, the only concern is confusion between identifiers displayed with the same
 * paragraph direction.
 *
 * <p>
 * An example is the case where identifiers are usernames prefixed with the @ symbol.
 * That symbol will appear to the left in a left-to-right context, and to the right in a
 * right-to-left context, so that an identifier displayed in a left-to-right context can never be
 * confused with an identifier displayed in a right-to-left context:
 * <ul>
 * <li>
 * The usernames "A1א" (A one aleph) and "Aא1" (A aleph 1)
 * would be considered confusable, since they both appear as \@A1א in a left-to-right context, and the
 * usernames "אA_1" (aleph A underscore one) and "א1_A" (aleph one underscore A) would be considered
 * confusable, since they both appear as A_1א@ in a right-to-left context.
 * </li>
 * <li>
 * The username "Mark_" would not be considered confusable with the username "_Mark",
 * even though the latter would appear as Mark_@ in a right-to-left context, and the
 * former as \@Mark_ in a left-to-right context.
 * </li>
 * </ul>
 * <p>
 * In that case, the caller should check for both LTR-confusability and RTL-confusability:
 *
 * \code{.cpp}
 * bool confusableInEitherDirection =
 *     uspoof_areBidiConfusableUnicodeString(sc, UBIDI_LTR, id1, id2, &status) ||
 *     uspoof_areBidiConfusableUnicodeString(sc, UBIDI_RTL, id1, id2, &status);
 * \endcode
 *
 * If the bidiSkeleton is used, the LTR and RTL skeleta should be kept separately and compared, LTR
 * with LTR and RTL with RTL.
 *
 * <p>
 * In cases where confusability between the visual appearances of an identifier displayed in a
 * left-to-right context with another identifier displayed in a right-to-left context is a concern,
 * the LTR skeleton of one can be compared with the RTL skeleton of the other.  However, this
 * very broad definition of confusability may have unexpected results; for instance, it treats the
 * ASCII identifiers "Mark_" and "_Mark" as confusable.
 *
 * <h2>Additional Information</h2>
 *
 * A <code>USpoofChecker</code> instance may be used repeatedly to perform checks on any number of identifiers.
 *
 * <b>Thread Safety:</b> The test functions for checking a single identifier, or for testing whether
 * two identifiers are possible confusable, are thread safe. They may called concurrently, from multiple threads,
 * using the same USpoofChecker instance.
 *
 * More generally, the standard ICU thread safety rules apply: functions that take a const USpoofChecker parameter are
 * thread safe. Those that take a non-const USpoofChecker are not thread safe..
 *
 * @stable ICU 4.6
 */

type USpoofChecker struct {
	Unused [8]uint8
}

type USpoofCheckResult struct {
	Unused [8]uint8
}
type USpoofChecks c.Int

const (
	USpoofChecksUSPOOFSINGLESCRIPTCONFUSABLE USpoofChecks = 1
	USpoofChecksUSPOOFMIXEDSCRIPTCONFUSABLE  USpoofChecks = 2
	USpoofChecksUSPOOFWHOLESCRIPTCONFUSABLE  USpoofChecks = 4
	USpoofChecksUSPOOFCONFUSABLE             USpoofChecks = 7
	USpoofChecksUSPOOFANYCASE                USpoofChecks = 8
	USpoofChecksUSPOOFRESTRICTIONLEVEL       USpoofChecks = 16
	USpoofChecksUSPOOFSINGLESCRIPT           USpoofChecks = 16
	USpoofChecksUSPOOFINVISIBLE              USpoofChecks = 32
	USpoofChecksUSPOOFCHARLIMIT              USpoofChecks = 64
	USpoofChecksUSPOOFMIXEDNUMBERS           USpoofChecks = 128
	USpoofChecksUSPOOFHIDDENOVERLAY          USpoofChecks = 256
	USpoofChecksUSPOOFALLCHECKS              USpoofChecks = 65535
	USpoofChecksUSPOOFAUXINFO                USpoofChecks = 1073741824
)

type URestrictionLevel c.Int

const (
	URestrictionLevelUSPOOFASCII                   URestrictionLevel = 268435456
	URestrictionLevelUSPOOFSINGLESCRIPTRESTRICTIVE URestrictionLevel = 536870912
	URestrictionLevelUSPOOFHIGHLYRESTRICTIVE       URestrictionLevel = 805306368
	URestrictionLevelUSPOOFMODERATELYRESTRICTIVE   URestrictionLevel = 1073741824
	URestrictionLevelUSPOOFMINIMALLYRESTRICTIVE    URestrictionLevel = 1342177280
	URestrictionLevelUSPOOFUNRESTRICTIVE           URestrictionLevel = 1610612736
	URestrictionLevelUSPOOFRESTRICTIONLEVELMASK    URestrictionLevel = 2130706432
	URestrictionLevelUSPOOFUNDEFINEDRESTRICTIVE    URestrictionLevel = -1
)
