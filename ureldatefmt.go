package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UDateRelativeDateTimeFormatterStyle c.Int

const (
	UDateRelativeDateTimeFormatterStyleUDATSTYLELONG   UDateRelativeDateTimeFormatterStyle = 0
	UDateRelativeDateTimeFormatterStyleUDATSTYLESHORT  UDateRelativeDateTimeFormatterStyle = 1
	UDateRelativeDateTimeFormatterStyleUDATSTYLENARROW UDateRelativeDateTimeFormatterStyle = 2
	UDateRelativeDateTimeFormatterStyleUDATSTYLECOUNT  UDateRelativeDateTimeFormatterStyle = 3
)

type URelativeDateTimeUnit c.Int

const (
	URelativeDateTimeUnitUDATRELUNITYEAR      URelativeDateTimeUnit = 0
	URelativeDateTimeUnitUDATRELUNITQUARTER   URelativeDateTimeUnit = 1
	URelativeDateTimeUnitUDATRELUNITMONTH     URelativeDateTimeUnit = 2
	URelativeDateTimeUnitUDATRELUNITWEEK      URelativeDateTimeUnit = 3
	URelativeDateTimeUnitUDATRELUNITDAY       URelativeDateTimeUnit = 4
	URelativeDateTimeUnitUDATRELUNITHOUR      URelativeDateTimeUnit = 5
	URelativeDateTimeUnitUDATRELUNITMINUTE    URelativeDateTimeUnit = 6
	URelativeDateTimeUnitUDATRELUNITSECOND    URelativeDateTimeUnit = 7
	URelativeDateTimeUnitUDATRELUNITSUNDAY    URelativeDateTimeUnit = 8
	URelativeDateTimeUnitUDATRELUNITMONDAY    URelativeDateTimeUnit = 9
	URelativeDateTimeUnitUDATRELUNITTUESDAY   URelativeDateTimeUnit = 10
	URelativeDateTimeUnitUDATRELUNITWEDNESDAY URelativeDateTimeUnit = 11
	URelativeDateTimeUnitUDATRELUNITTHURSDAY  URelativeDateTimeUnit = 12
	URelativeDateTimeUnitUDATRELUNITFRIDAY    URelativeDateTimeUnit = 13
	URelativeDateTimeUnitUDATRELUNITSATURDAY  URelativeDateTimeUnit = 14
	URelativeDateTimeUnitUDATRELUNITCOUNT     URelativeDateTimeUnit = 15
)

type URelativeDateTimeFormatterField c.Int

const (
	URelativeDateTimeFormatterFieldUDATRELLITERALFIELD URelativeDateTimeFormatterField = 0
	URelativeDateTimeFormatterFieldUDATRELNUMERICFIELD URelativeDateTimeFormatterField = 1
)
/**
 * Opaque URelativeDateTimeFormatter object for use in C programs.
 * @stable ICU 57
 */

type URelativeDateTimeFormatter struct {
	Unused [8]uint8
}

type UFormattedRelativeDateTime struct {
	Unused [8]uint8
}
