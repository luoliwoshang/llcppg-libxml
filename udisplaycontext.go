package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type UDisplayContextType c.Int

const (
	UDisplayContextTypeUDISPCTXTYPEDIALECTHANDLING    UDisplayContextType = 0
	UDisplayContextTypeUDISPCTXTYPECAPITALIZATION     UDisplayContextType = 1
	UDisplayContextTypeUDISPCTXTYPEDISPLAYLENGTH      UDisplayContextType = 2
	UDisplayContextTypeUDISPCTXTYPESUBSTITUTEHANDLING UDisplayContextType = 3
)

type UDisplayContext c.Int

const (
	UDisplayContextUDISPCTXSTANDARDNAMES                        UDisplayContext = 0
	UDisplayContextUDISPCTXDIALECTNAMES                         UDisplayContext = 1
	UDisplayContextUDISPCTXCAPITALIZATIONNONE                   UDisplayContext = 256
	UDisplayContextUDISPCTXCAPITALIZATIONFORMIDDLEOFSENTENCE    UDisplayContext = 257
	UDisplayContextUDISPCTXCAPITALIZATIONFORBEGINNINGOFSENTENCE UDisplayContext = 258
	UDisplayContextUDISPCTXCAPITALIZATIONFORUILISTORMENU        UDisplayContext = 259
	UDisplayContextUDISPCTXCAPITALIZATIONFORSTANDALONE          UDisplayContext = 260
	UDisplayContextUDISPCTXLENGTHFULL                           UDisplayContext = 512
	UDisplayContextUDISPCTXLENGTHSHORT                          UDisplayContext = 513
	UDisplayContextUDISPCTXSUBSTITUTE                           UDisplayContext = 768
	UDisplayContextUDISPCTXNOSUBSTITUTE                         UDisplayContext = 769
)
