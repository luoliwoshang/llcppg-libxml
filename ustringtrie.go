package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UStringTrieResult c.Int

const (
	UStringTrieResultUSTRINGTRIENOMATCH           UStringTrieResult = 0
	UStringTrieResultUSTRINGTRIENOVALUE           UStringTrieResult = 1
	UStringTrieResultUSTRINGTRIEFINALVALUE        UStringTrieResult = 2
	UStringTrieResultUSTRINGTRIEINTERMEDIATEVALUE UStringTrieResult = 3
)
