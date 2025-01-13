package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UNumberRangeCollapse c.Int

const (
	UNumberRangeCollapseUNUMRANGECOLLAPSEAUTO UNumberRangeCollapse = 0
	UNumberRangeCollapseUNUMRANGECOLLAPSENONE UNumberRangeCollapse = 1
	UNumberRangeCollapseUNUMRANGECOLLAPSEUNIT UNumberRangeCollapse = 2
	UNumberRangeCollapseUNUMRANGECOLLAPSEALL  UNumberRangeCollapse = 3
)

type UNumberRangeIdentityFallback c.Int

const (
	UNumberRangeIdentityFallbackUNUMIDENTITYFALLBACKSINGLEVALUE                UNumberRangeIdentityFallback = 0
	UNumberRangeIdentityFallbackUNUMIDENTITYFALLBACKAPPROXIMATELYORSINGLEVALUE UNumberRangeIdentityFallback = 1
	UNumberRangeIdentityFallbackUNUMIDENTITYFALLBACKAPPROXIMATELY              UNumberRangeIdentityFallback = 2
	UNumberRangeIdentityFallbackUNUMIDENTITYFALLBACKRANGE                      UNumberRangeIdentityFallback = 3
)

type UNumberRangeIdentityResult c.Int

const (
	UNumberRangeIdentityResultUNUMIDENTITYRESULTEQUALBEFOREROUNDING UNumberRangeIdentityResult = 0
	UNumberRangeIdentityResultUNUMIDENTITYRESULTEQUALAFTERROUNDING  UNumberRangeIdentityResult = 1
	UNumberRangeIdentityResultUNUMIDENTITYRESULTNOTEQUAL            UNumberRangeIdentityResult = 2
	UNumberRangeIdentityResultUNUMIDENTITYRESULTCOUNT               UNumberRangeIdentityResult = 3
)

type UNumberRangeFormatter struct {
	Unused [8]uint8
}
