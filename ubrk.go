package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const UBRKSAFECLONEBUFFERSIZE c.Int = 1

type UBreakIterator struct {
	Unused [8]uint8
}
type UBreakIteratorType c.Int

const (
	UBreakIteratorTypeUBRKCHARACTER UBreakIteratorType = 0
	UBreakIteratorTypeUBRKWORD      UBreakIteratorType = 1
	UBreakIteratorTypeUBRKLINE      UBreakIteratorType = 2
	UBreakIteratorTypeUBRKSENTENCE  UBreakIteratorType = 3
	UBreakIteratorTypeUBRKTITLE     UBreakIteratorType = 4
	UBreakIteratorTypeUBRKCOUNT     UBreakIteratorType = 5
)

type UWordBreak c.Int

const (
	UWordBreakUBRKWORDNONE        UWordBreak = 0
	UWordBreakUBRKWORDNONELIMIT   UWordBreak = 100
	UWordBreakUBRKWORDNUMBER      UWordBreak = 100
	UWordBreakUBRKWORDNUMBERLIMIT UWordBreak = 200
	UWordBreakUBRKWORDLETTER      UWordBreak = 200
	UWordBreakUBRKWORDLETTERLIMIT UWordBreak = 300
	UWordBreakUBRKWORDKANA        UWordBreak = 300
	UWordBreakUBRKWORDKANALIMIT   UWordBreak = 400
	UWordBreakUBRKWORDIDEO        UWordBreak = 400
	UWordBreakUBRKWORDIDEOLIMIT   UWordBreak = 500
)

type ULineBreakTag c.Int

const (
	ULineBreakTagUBRKLINESOFT      ULineBreakTag = 0
	ULineBreakTagUBRKLINESOFTLIMIT ULineBreakTag = 100
	ULineBreakTagUBRKLINEHARD      ULineBreakTag = 100
	ULineBreakTagUBRKLINEHARDLIMIT ULineBreakTag = 200
)

type USentenceBreakTag c.Int

const (
	USentenceBreakTagUBRKSENTENCETERM      USentenceBreakTag = 0
	USentenceBreakTagUBRKSENTENCETERMLIMIT USentenceBreakTag = 100
	USentenceBreakTagUBRKSENTENCESEP       USentenceBreakTag = 100
	USentenceBreakTagUBRKSENTENCESEPLIMIT  USentenceBreakTag = 200
)
