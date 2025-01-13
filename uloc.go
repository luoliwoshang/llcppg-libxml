package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const ULOCCHINESE string = "zh"
const ULOCENGLISH string = "en"
const ULOCFRENCH string = "fr"
const ULOCGERMAN string = "de"
const ULOCITALIAN string = "it"
const ULOCJAPANESE string = "ja"
const ULOCKOREAN string = "ko"
const ULOCSIMPLIFIEDCHINESE string = "zh_CN"
const ULOCTRADITIONALCHINESE string = "zh_TW"
const ULOCCANADA string = "en_CA"
const ULOCCANADAFRENCH string = "fr_CA"
const ULOCCHINA string = "zh_CN"
const ULOCPRC string = "zh_CN"
const ULOCFRANCE string = "fr_FR"
const ULOCGERMANY string = "de_DE"
const ULOCITALY string = "it_IT"
const ULOCJAPAN string = "ja_JP"
const ULOCKOREA string = "ko_KR"
const ULOCTAIWAN string = "zh_TW"
const ULOCUK string = "en_GB"
const ULOCUS string = "en_US"
const ULOCLANGCAPACITY c.Int = 12
const ULOCCOUNTRYCAPACITY c.Int = 4
const ULOCFULLNAMECAPACITY c.Int = 157
const ULOCSCRIPTCAPACITY c.Int = 6
const ULOCKEYWORDSCAPACITY c.Int = 96
const ULOCKEYWORDANDVALUESCAPACITY c.Int = 100
const ULOCKEYWORDSEPARATOR string = "@"
const ULOCKEYWORDSEPARATORUNICODE c.Int = 64
const ULOCKEYWORDASSIGN string = "="
const ULOCKEYWORDASSIGNUNICODE c.Int = 61
const ULOCKEYWORDITEMSEPARATOR string = ";"
const ULOCKEYWORDITEMSEPARATORUNICODE c.Int = 59

type ULocDataLocaleType c.Int

const (
	ULocDataLocaleTypeULOCACTUALLOCALE        ULocDataLocaleType = 0
	ULocDataLocaleTypeULOCVALIDLOCALE         ULocDataLocaleType = 1
	ULocDataLocaleTypeULOCREQUESTEDLOCALE     ULocDataLocaleType = 2
	ULocDataLocaleTypeULOCDATALOCALETYPELIMIT ULocDataLocaleType = 3
)

type ULocAvailableType c.Int

const (
	ULocAvailableTypeULOCAVAILABLEDEFAULT           ULocAvailableType = 0
	ULocAvailableTypeULOCAVAILABLEONLYLEGACYALIASES ULocAvailableType = 1
	ULocAvailableTypeULOCAVAILABLEWITHLEGACYALIASES ULocAvailableType = 2
	ULocAvailableTypeULOCAVAILABLECOUNT             ULocAvailableType = 3
)

type ULayoutType c.Int

const (
	ULayoutTypeULOCLAYOUTLTR     ULayoutType = 0
	ULayoutTypeULOCLAYOUTRTL     ULayoutType = 1
	ULayoutTypeULOCLAYOUTTTB     ULayoutType = 2
	ULayoutTypeULOCLAYOUTBTT     ULayoutType = 3
	ULayoutTypeULOCLAYOUTUNKNOWN ULayoutType = 4
)

type UAcceptResult c.Int

const (
	UAcceptResultULOCACCEPTFAILED   UAcceptResult = 0
	UAcceptResultULOCACCEPTVALID    UAcceptResult = 1
	UAcceptResultULOCACCEPTFALLBACK UAcceptResult = 2
)
