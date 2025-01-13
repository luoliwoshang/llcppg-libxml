package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)
/** Forward declaration of the ULocaleData structure. @stable ICU 3.6 */

type ULocaleData struct {
	Unused [8]uint8
}
type ULocaleDataExemplarSetType c.Int

const (
	ULocaleDataExemplarSetTypeULOCDATAESSTANDARD    ULocaleDataExemplarSetType = 0
	ULocaleDataExemplarSetTypeULOCDATAESAUXILIARY   ULocaleDataExemplarSetType = 1
	ULocaleDataExemplarSetTypeULOCDATAESINDEX       ULocaleDataExemplarSetType = 2
	ULocaleDataExemplarSetTypeULOCDATAESPUNCTUATION ULocaleDataExemplarSetType = 3
	ULocaleDataExemplarSetTypeULOCDATAESCOUNT       ULocaleDataExemplarSetType = 4
)

type ULocaleDataDelimiterType c.Int

const (
	ULocaleDataDelimiterTypeULOCDATAQUOTATIONSTART    ULocaleDataDelimiterType = 0
	ULocaleDataDelimiterTypeULOCDATAQUOTATIONEND      ULocaleDataDelimiterType = 1
	ULocaleDataDelimiterTypeULOCDATAALTQUOTATIONSTART ULocaleDataDelimiterType = 2
	ULocaleDataDelimiterTypeULOCDATAALTQUOTATIONEND   ULocaleDataDelimiterType = 3
	ULocaleDataDelimiterTypeULOCDATADELIMITERCOUNT    ULocaleDataDelimiterType = 4
)

type UMeasurementSystem c.Int

const (
	UMeasurementSystemUMSSI    UMeasurementSystem = 0
	UMeasurementSystemUMSUS    UMeasurementSystem = 1
	UMeasurementSystemUMSUK    UMeasurementSystem = 2
	UMeasurementSystemUMSLIMIT UMeasurementSystem = 3
)
