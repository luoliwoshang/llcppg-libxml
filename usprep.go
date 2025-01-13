package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

const USPREPDEFAULT c.Int = 0
const USPREPALLOWUNASSIGNED c.Int = 1

type UStringPrepProfile struct {
	Unused [8]uint8
}
type UStringPrepProfileType c.Int

const (
	UStringPrepProfileTypeUSPREPRFC3491NAMEPREP            UStringPrepProfileType = 0
	UStringPrepProfileTypeUSPREPRFC3530NFS4CSPREP          UStringPrepProfileType = 1
	UStringPrepProfileTypeUSPREPRFC3530NFS4CSPREPCI        UStringPrepProfileType = 2
	UStringPrepProfileTypeUSPREPRFC3530NFS4CISPREP         UStringPrepProfileType = 3
	UStringPrepProfileTypeUSPREPRFC3530NFS4MIXEDPREPPREFIX UStringPrepProfileType = 4
	UStringPrepProfileTypeUSPREPRFC3530NFS4MIXEDPREPSUFFIX UStringPrepProfileType = 5
	UStringPrepProfileTypeUSPREPRFC3722ISCSI               UStringPrepProfileType = 6
	UStringPrepProfileTypeUSPREPRFC3920NODEPREP            UStringPrepProfileType = 7
	UStringPrepProfileTypeUSPREPRFC3920RESOURCEPREP        UStringPrepProfileType = 8
	UStringPrepProfileTypeUSPREPRFC4011MIB                 UStringPrepProfileType = 9
	UStringPrepProfileTypeUSPREPRFC4013SASLPREP            UStringPrepProfileType = 10
	UStringPrepProfileTypeUSPREPRFC4505TRACE               UStringPrepProfileType = 11
	UStringPrepProfileTypeUSPREPRFC4518LDAP                UStringPrepProfileType = 12
	UStringPrepProfileTypeUSPREPRFC4518LDAPCI              UStringPrepProfileType = 13
)
