package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UNormalization2Mode c.Int

const (
	UNormalization2ModeUNORM2COMPOSE           UNormalization2Mode = 0
	UNormalization2ModeUNORM2DECOMPOSE         UNormalization2Mode = 1
	UNormalization2ModeUNORM2FCD               UNormalization2Mode = 2
	UNormalization2ModeUNORM2COMPOSECONTIGUOUS UNormalization2Mode = 3
)

type UNormalizationCheckResult c.Int

const (
	UNormalizationCheckResultUNORMNO    UNormalizationCheckResult = 0
	UNormalizationCheckResultUNORMYES   UNormalizationCheckResult = 1
	UNormalizationCheckResultUNORMMAYBE UNormalizationCheckResult = 2
)
/**
 * Opaque C service object type for the new normalization API.
 * @stable ICU 4.4
 */

type UNormalizer2 struct {
	Unused [8]uint8
}
