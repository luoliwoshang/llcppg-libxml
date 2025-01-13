package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	"unsafe"
)

type UText struct {
	Magic               uint32
	Flags               int32
	ProviderProperties  int32
	SizeOfStruct        int32
	ChunkNativeLimit    int64
	ExtraSize           int32
	NativeIndexingLimit int32
	ChunkNativeStart    int64
	ChunkOffset         int32
	ChunkLength         int32
	ChunkContents       *UChar
	PFuncs              *UTextFuncs
	PExtra              unsafe.Pointer
	Context             unsafe.Pointer
	P                   unsafe.Pointer
	Q                   unsafe.Pointer
	R                   unsafe.Pointer
	PrivP               unsafe.Pointer
	A                   int64
	B                   int32
	C                   int32
	PrivA               int64
	PrivB               int32
	PrivC               int32
}

const (
	UTEXTPROVIDERLENGTHISEXPENSIVE c.Int = 1
	UTEXTPROVIDERSTABLECHUNKS      c.Int = 2
	UTEXTPROVIDERWRITABLE          c.Int = 3
	UTEXTPROVIDERHASMETADATA       c.Int = 4
	UTEXTPROVIDEROWNSTEXT          c.Int = 5
)
// llgo:type C
type UTextClone func(*UText, *UText, UBool, *UErrorCode) *UText
// llgo:type C
type UTextNativeLength func(*UText) int64
// llgo:type C
type UTextAccess func(*UText, int64, UBool) UBool
// llgo:type C
type UTextExtract func(*UText, int64, int64, *UChar, int32, *UErrorCode) int32
// llgo:type C
type UTextReplace func(*UText, int64, int64, *UChar, int32, *UErrorCode) int32
// llgo:type C
type UTextCopy func(*UText, int64, int64, int64, UBool, *UErrorCode)
// llgo:type C
type UTextMapOffsetToNative func(*UText) int64
// llgo:type C
type UTextMapNativeIndexToUTF16 func(*UText, int64) int32
// llgo:type C
type UTextClose func(*UText)
/**
  *   (public)  Function dispatch table for UText.
  *             Conceptually very much like a C++ Virtual Function Table.
  *             This struct defines the organization of the table.
  *             Each text provider implementation must provide an
  *              actual table that is initialized with the appropriate functions
  *              for the type of text being handled.
  *   @stable ICU 3.6
  */
type UTextFuncs struct {
	TableSize             int32
	Reserved1             int32
	Reserved2             int32
	Reserved3             int32
	Clone                 *unsafe.Pointer
	NativeLength          *unsafe.Pointer
	Access                *unsafe.Pointer
	Extract               *unsafe.Pointer
	Replace               *unsafe.Pointer
	Copy                  *unsafe.Pointer
	MapOffsetToNative     *unsafe.Pointer
	MapNativeIndexToUTF16 *unsafe.Pointer
	Close                 *unsafe.Pointer
	Spare1                *unsafe.Pointer
	Spare2                *unsafe.Pointer
	Spare3                *unsafe.Pointer
}

const UTEXTMAGIC c.Int = 878368812
