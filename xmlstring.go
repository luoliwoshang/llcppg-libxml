package libxml_2_0

import (
	"github.com/goplus/llgo/c"
	_ "unsafe"
)

type XmlChar int8
// llgo:link (*XmlChar).XmlStrdup C.xmlStrdup
func (recv_ *XmlChar) XmlStrdup() *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlStrndup C.xmlStrndup
func (recv_ *XmlChar) XmlStrndup(len c.Int) *XmlChar {
	return nil
}
//go:linkname XmlCharStrndup C.xmlCharStrndup
func XmlCharStrndup(cur *int8, len c.Int) *XmlChar
//go:linkname XmlCharStrdup C.xmlCharStrdup
func XmlCharStrdup(cur *int8) *XmlChar
// llgo:link (*XmlChar).XmlStrsub C.xmlStrsub
func (recv_ *XmlChar) XmlStrsub(start c.Int, len c.Int) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlStrchr C.xmlStrchr
func (recv_ *XmlChar) XmlStrchr(val XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlStrstr C.xmlStrstr
func (recv_ *XmlChar) XmlStrstr(val *XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlStrcasestr C.xmlStrcasestr
func (recv_ *XmlChar) XmlStrcasestr(val *XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlStrcmp C.xmlStrcmp
func (recv_ *XmlChar) XmlStrcmp(str2 *XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlStrncmp C.xmlStrncmp
func (recv_ *XmlChar) XmlStrncmp(str2 *XmlChar, len c.Int) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlStrcasecmp C.xmlStrcasecmp
func (recv_ *XmlChar) XmlStrcasecmp(str2 *XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlStrncasecmp C.xmlStrncasecmp
func (recv_ *XmlChar) XmlStrncasecmp(str2 *XmlChar, len c.Int) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlStrEqual C.xmlStrEqual
func (recv_ *XmlChar) XmlStrEqual(str2 *XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlStrQEqual C.xmlStrQEqual
func (recv_ *XmlChar) XmlStrQEqual(name *XmlChar, str *XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlStrlen C.xmlStrlen
func (recv_ *XmlChar) XmlStrlen() c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlStrcat C.xmlStrcat
func (recv_ *XmlChar) XmlStrcat(add *XmlChar) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlStrncat C.xmlStrncat
func (recv_ *XmlChar) XmlStrncat(add *XmlChar, len c.Int) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlStrncatNew C.xmlStrncatNew
func (recv_ *XmlChar) XmlStrncatNew(str2 *XmlChar, len c.Int) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlStrPrintf C.xmlStrPrintf
func (recv_ *XmlChar) XmlStrPrintf(len c.Int, msg *int8, __llgo_va_list ...interface{}) c.Int {
	return 0
}
//go:linkname XmlGetUTF8Char C.xmlGetUTF8Char
func XmlGetUTF8Char(utf *int8, len *c.Int) c.Int
//go:linkname XmlCheckUTF8 C.xmlCheckUTF8
func XmlCheckUTF8(utf *int8) c.Int
// llgo:link (*XmlChar).XmlUTF8Strsize C.xmlUTF8Strsize
func (recv_ *XmlChar) XmlUTF8Strsize(len c.Int) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlUTF8Strndup C.xmlUTF8Strndup
func (recv_ *XmlChar) XmlUTF8Strndup(len c.Int) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlUTF8Strpos C.xmlUTF8Strpos
func (recv_ *XmlChar) XmlUTF8Strpos(pos c.Int) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlUTF8Strloc C.xmlUTF8Strloc
func (recv_ *XmlChar) XmlUTF8Strloc(utfchar *XmlChar) c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlUTF8Strsub C.xmlUTF8Strsub
func (recv_ *XmlChar) XmlUTF8Strsub(start c.Int, len c.Int) *XmlChar {
	return nil
}
// llgo:link (*XmlChar).XmlUTF8Strlen C.xmlUTF8Strlen
func (recv_ *XmlChar) XmlUTF8Strlen() c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlUTF8Size C.xmlUTF8Size
func (recv_ *XmlChar) XmlUTF8Size() c.Int {
	return 0
}
// llgo:link (*XmlChar).XmlUTF8Charcmp C.xmlUTF8Charcmp
func (recv_ *XmlChar) XmlUTF8Charcmp(utf2 *XmlChar) c.Int {
	return 0
}
