package libxml_2_0

import _ "unsafe"
/**
 * \file
 * \brief C API: Charset Detection API
 *
 * This API provides a facility for detecting the
 * charset or encoding of character data in an unknown text format.
 * The input data can be from an array of bytes.
 * <p>
 * Character set detection is at best an imprecise operation.  The detection
 * process will attempt to identify the charset that best matches the characteristics
 * of the byte data, but the process is partly statistical in nature, and
 * the results can not be guaranteed to always be correct.
 * <p>
 * For best accuracy in charset detection, the input data should be primarily
 * in a single language, and a minimum of a few hundred bytes worth of plain text
 * in the language are needed.  The detection process will attempt to
 * ignore html or xml style markup that could otherwise obscure the content.
 * <p>
 * An alternative to the ICU Charset Detector is the
 * Compact Encoding Detector, https://github.com/google/compact_enc_det.
 * It often gives more accurate results, especially with short input samples.
 */

type UCharsetDetector struct {
	Unused [8]uint8
}

type UCharsetMatch struct {
	Unused [8]uint8
}
