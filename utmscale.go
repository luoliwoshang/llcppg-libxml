package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UDateTimeScale c.Int

const (
	UDateTimeScaleUDTSJAVATIME             UDateTimeScale = 0
	UDateTimeScaleUDTSUNIXTIME             UDateTimeScale = 1
	UDateTimeScaleUDTSICU4CTIME            UDateTimeScale = 2
	UDateTimeScaleUDTSWINDOWSFILETIME      UDateTimeScale = 3
	UDateTimeScaleUDTSDOTNETDATETIME       UDateTimeScale = 4
	UDateTimeScaleUDTSMACOLDTIME           UDateTimeScale = 5
	UDateTimeScaleUDTSMACTIME              UDateTimeScale = 6
	UDateTimeScaleUDTSEXCELTIME            UDateTimeScale = 7
	UDateTimeScaleUDTSDB2TIME              UDateTimeScale = 8
	UDateTimeScaleUDTSUNIXMICROSECONDSTIME UDateTimeScale = 9
	UDateTimeScaleUDTSMAXSCALE             UDateTimeScale = 10
)

type UTimeScaleValue c.Int

const (
	UTimeScaleValueUTSVUNITSVALUE             UTimeScaleValue = 0
	UTimeScaleValueUTSVEPOCHOFFSETVALUE       UTimeScaleValue = 1
	UTimeScaleValueUTSVFROMMINVALUE           UTimeScaleValue = 2
	UTimeScaleValueUTSVFROMMAXVALUE           UTimeScaleValue = 3
	UTimeScaleValueUTSVTOMINVALUE             UTimeScaleValue = 4
	UTimeScaleValueUTSVTOMAXVALUE             UTimeScaleValue = 5
	UTimeScaleValueUTSVEPOCHOFFSETPLUS1VALUE  UTimeScaleValue = 6
	UTimeScaleValueUTSVEPOCHOFFSETMINUS1VALUE UTimeScaleValue = 7
	UTimeScaleValueUTSVUNITSROUNDVALUE        UTimeScaleValue = 8
	UTimeScaleValueUTSVMINROUNDVALUE          UTimeScaleValue = 9
	UTimeScaleValueUTSVMAXROUNDVALUE          UTimeScaleValue = 10
	UTimeScaleValueUTSVMAXSCALEVALUE          UTimeScaleValue = 11
)
