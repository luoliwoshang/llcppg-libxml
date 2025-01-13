package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

const UCALUNKNOWNZONEID string = "Etc/Unknown"

type UCalendar unsafe.Pointer
type UCalendarType c.Int

const (
	UCalendarTypeUCALTRADITIONAL UCalendarType = 0
	UCalendarTypeUCALDEFAULT     UCalendarType = 0
	UCalendarTypeUCALGREGORIAN   UCalendarType = 1
)

type UCalendarDateFields c.Int

const (
	UCalendarDateFieldsUCALERA               UCalendarDateFields = 0
	UCalendarDateFieldsUCALYEAR              UCalendarDateFields = 1
	UCalendarDateFieldsUCALMONTH             UCalendarDateFields = 2
	UCalendarDateFieldsUCALWEEKOFYEAR        UCalendarDateFields = 3
	UCalendarDateFieldsUCALWEEKOFMONTH       UCalendarDateFields = 4
	UCalendarDateFieldsUCALDATE              UCalendarDateFields = 5
	UCalendarDateFieldsUCALDAYOFYEAR         UCalendarDateFields = 6
	UCalendarDateFieldsUCALDAYOFWEEK         UCalendarDateFields = 7
	UCalendarDateFieldsUCALDAYOFWEEKINMONTH  UCalendarDateFields = 8
	UCalendarDateFieldsUCALAMPM              UCalendarDateFields = 9
	UCalendarDateFieldsUCALHOUR              UCalendarDateFields = 10
	UCalendarDateFieldsUCALHOUROFDAY         UCalendarDateFields = 11
	UCalendarDateFieldsUCALMINUTE            UCalendarDateFields = 12
	UCalendarDateFieldsUCALSECOND            UCalendarDateFields = 13
	UCalendarDateFieldsUCALMILLISECOND       UCalendarDateFields = 14
	UCalendarDateFieldsUCALZONEOFFSET        UCalendarDateFields = 15
	UCalendarDateFieldsUCALDSTOFFSET         UCalendarDateFields = 16
	UCalendarDateFieldsUCALYEARWOY           UCalendarDateFields = 17
	UCalendarDateFieldsUCALDOWLOCAL          UCalendarDateFields = 18
	UCalendarDateFieldsUCALEXTENDEDYEAR      UCalendarDateFields = 19
	UCalendarDateFieldsUCALJULIANDAY         UCalendarDateFields = 20
	UCalendarDateFieldsUCALMILLISECONDSINDAY UCalendarDateFields = 21
	UCalendarDateFieldsUCALISLEAPMONTH       UCalendarDateFields = 22
	UCalendarDateFieldsUCALORDINALMONTH      UCalendarDateFields = 23
	UCalendarDateFieldsUCALFIELDCOUNT        UCalendarDateFields = 24
	UCalendarDateFieldsUCALDAYOFMONTH        UCalendarDateFields = 5
)

type UCalendarDaysOfWeek c.Int

const (
	UCalendarDaysOfWeekUCALSUNDAY    UCalendarDaysOfWeek = 1
	UCalendarDaysOfWeekUCALMONDAY    UCalendarDaysOfWeek = 2
	UCalendarDaysOfWeekUCALTUESDAY   UCalendarDaysOfWeek = 3
	UCalendarDaysOfWeekUCALWEDNESDAY UCalendarDaysOfWeek = 4
	UCalendarDaysOfWeekUCALTHURSDAY  UCalendarDaysOfWeek = 5
	UCalendarDaysOfWeekUCALFRIDAY    UCalendarDaysOfWeek = 6
	UCalendarDaysOfWeekUCALSATURDAY  UCalendarDaysOfWeek = 7
)

type UCalendarMonths c.Int

const (
	UCalendarMonthsUCALJANUARY    UCalendarMonths = 0
	UCalendarMonthsUCALFEBRUARY   UCalendarMonths = 1
	UCalendarMonthsUCALMARCH      UCalendarMonths = 2
	UCalendarMonthsUCALAPRIL      UCalendarMonths = 3
	UCalendarMonthsUCALMAY        UCalendarMonths = 4
	UCalendarMonthsUCALJUNE       UCalendarMonths = 5
	UCalendarMonthsUCALJULY       UCalendarMonths = 6
	UCalendarMonthsUCALAUGUST     UCalendarMonths = 7
	UCalendarMonthsUCALSEPTEMBER  UCalendarMonths = 8
	UCalendarMonthsUCALOCTOBER    UCalendarMonths = 9
	UCalendarMonthsUCALNOVEMBER   UCalendarMonths = 10
	UCalendarMonthsUCALDECEMBER   UCalendarMonths = 11
	UCalendarMonthsUCALUNDECIMBER UCalendarMonths = 12
)

type UCalendarAMPMs c.Int

const (
	UCalendarAMPMsUCALAM UCalendarAMPMs = 0
	UCalendarAMPMsUCALPM UCalendarAMPMs = 1
)

type USystemTimeZoneType c.Int

const (
	USystemTimeZoneTypeUCALZONETYPEANY               USystemTimeZoneType = 0
	USystemTimeZoneTypeUCALZONETYPECANONICAL         USystemTimeZoneType = 1
	USystemTimeZoneTypeUCALZONETYPECANONICALLOCATION USystemTimeZoneType = 2
)

type UCalendarDisplayNameType c.Int

const (
	UCalendarDisplayNameTypeUCALSTANDARD      UCalendarDisplayNameType = 0
	UCalendarDisplayNameTypeUCALSHORTSTANDARD UCalendarDisplayNameType = 1
	UCalendarDisplayNameTypeUCALDST           UCalendarDisplayNameType = 2
	UCalendarDisplayNameTypeUCALSHORTDST      UCalendarDisplayNameType = 3
)

type UCalendarAttribute c.Int

const (
	UCalendarAttributeUCALLENIENT                UCalendarAttribute = 0
	UCalendarAttributeUCALFIRSTDAYOFWEEK         UCalendarAttribute = 1
	UCalendarAttributeUCALMINIMALDAYSINFIRSTWEEK UCalendarAttribute = 2
	UCalendarAttributeUCALREPEATEDWALLTIME       UCalendarAttribute = 3
	UCalendarAttributeUCALSKIPPEDWALLTIME        UCalendarAttribute = 4
)

type UCalendarWallTimeOption c.Int

const (
	UCalendarWallTimeOptionUCALWALLTIMELAST      UCalendarWallTimeOption = 0
	UCalendarWallTimeOptionUCALWALLTIMEFIRST     UCalendarWallTimeOption = 1
	UCalendarWallTimeOptionUCALWALLTIMENEXTVALID UCalendarWallTimeOption = 2
)

type UCalendarLimitType c.Int

const (
	UCalendarLimitTypeUCALMINIMUM         UCalendarLimitType = 0
	UCalendarLimitTypeUCALMAXIMUM         UCalendarLimitType = 1
	UCalendarLimitTypeUCALGREATESTMINIMUM UCalendarLimitType = 2
	UCalendarLimitTypeUCALLEASTMAXIMUM    UCalendarLimitType = 3
	UCalendarLimitTypeUCALACTUALMINIMUM   UCalendarLimitType = 4
	UCalendarLimitTypeUCALACTUALMAXIMUM   UCalendarLimitType = 5
)

type UCalendarWeekdayType c.Int

const (
	UCalendarWeekdayTypeUCALWEEKDAY      UCalendarWeekdayType = 0
	UCalendarWeekdayTypeUCALWEEKEND      UCalendarWeekdayType = 1
	UCalendarWeekdayTypeUCALWEEKENDONSET UCalendarWeekdayType = 2
	UCalendarWeekdayTypeUCALWEEKENDCEASE UCalendarWeekdayType = 3
)

type UTimeZoneTransitionType c.Int

const (
	UTimeZoneTransitionTypeUCALTZTRANSITIONNEXT              UTimeZoneTransitionType = 0
	UTimeZoneTransitionTypeUCALTZTRANSITIONNEXTINCLUSIVE     UTimeZoneTransitionType = 1
	UTimeZoneTransitionTypeUCALTZTRANSITIONPREVIOUS          UTimeZoneTransitionType = 2
	UTimeZoneTransitionTypeUCALTZTRANSITIONPREVIOUSINCLUSIVE UTimeZoneTransitionType = 3
)

type UTimeZoneLocalOption c.Int

const (
	UTimeZoneLocalOptionUCALTZLOCALFORMER         UTimeZoneLocalOption = 4
	UTimeZoneLocalOptionUCALTZLOCALLATTER         UTimeZoneLocalOption = 12
	UTimeZoneLocalOptionUCALTZLOCALSTANDARDFORMER UTimeZoneLocalOption = 5
	UTimeZoneLocalOptionUCALTZLOCALSTANDARDLATTER UTimeZoneLocalOption = 13
	UTimeZoneLocalOptionUCALTZLOCALDAYLIGHTFORMER UTimeZoneLocalOption = 7
	UTimeZoneLocalOptionUCALTZLOCALDAYLIGHTLATTER UTimeZoneLocalOption = 15
)
