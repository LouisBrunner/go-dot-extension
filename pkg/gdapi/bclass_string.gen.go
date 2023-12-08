// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type String struct {
  obj gdc.ObjectPtr
}

func (me *String) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *String) BaseClass() string {
  return "String"
}

func  (me *String) CasecmpTo(to String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) NocasecmpTo(to String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) NaturalcasecmpTo(to String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) NaturalnocasecmpTo(to String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Length() { // TODO: return value
  // TODO: implement
}

func  (me *String) Substr(from int, len int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) GetSlice(delimiter String, slice int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) GetSlicec(delimiter int, slice int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) GetSliceCount(delimiter String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Find(what String, from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Count(what String, from int, to int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Countn(what String, from int, to int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Findn(what String, from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Rfind(what String, from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Rfindn(what String, from int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Match(expr String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Matchn(expr String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) BeginsWith(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) EndsWith(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) IsSubsequenceOf(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) IsSubsequenceOfn(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Bigrams() { // TODO: return value
  // TODO: implement
}

func  (me *String) Similarity(text String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Format(values Variant, placeholder String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Replace(what String, forwhat String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Replacen(what String, forwhat String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Repeat(count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Insert(position int, what String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Erase(position int, chars int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Capitalize() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToCamelCase() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToPascalCase() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToSnakeCase() { // TODO: return value
  // TODO: implement
}

func  (me *String) Split(delimiter String, allow_empty bool, maxsplit int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Rsplit(delimiter String, allow_empty bool, maxsplit int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) SplitFloats(delimiter String, allow_empty bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Join(parts PackedStringArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) ToUpper() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToLower() { // TODO: return value
  // TODO: implement
}

func  (me *String) Left(length int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Right(length int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) StripEdges(left bool, right bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) StripEscapes() { // TODO: return value
  // TODO: implement
}

func  (me *String) Lstrip(chars String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Rstrip(chars String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) GetExtension() { // TODO: return value
  // TODO: implement
}

func  (me *String) GetBasename() { // TODO: return value
  // TODO: implement
}

func  (me *String) PathJoin(file String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) UnicodeAt(at int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Indent(prefix String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Dedent() { // TODO: return value
  // TODO: implement
}

func  (me *String) Hash() { // TODO: return value
  // TODO: implement
}

func  (me *String) Md5Text() { // TODO: return value
  // TODO: implement
}

func  (me *String) Sha1Text() { // TODO: return value
  // TODO: implement
}

func  (me *String) Sha256Text() { // TODO: return value
  // TODO: implement
}

func  (me *String) Md5Buffer() { // TODO: return value
  // TODO: implement
}

func  (me *String) Sha1Buffer() { // TODO: return value
  // TODO: implement
}

func  (me *String) Sha256Buffer() { // TODO: return value
  // TODO: implement
}

func  (me *String) IsEmpty() { // TODO: return value
  // TODO: implement
}

func  (me *String) Contains(what String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) IsAbsolutePath() { // TODO: return value
  // TODO: implement
}

func  (me *String) IsRelativePath() { // TODO: return value
  // TODO: implement
}

func  (me *String) SimplifyPath() { // TODO: return value
  // TODO: implement
}

func  (me *String) GetBaseDir() { // TODO: return value
  // TODO: implement
}

func  (me *String) GetFile() { // TODO: return value
  // TODO: implement
}

func  (me *String) XmlEscape(escape_quotes bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) XmlUnescape() { // TODO: return value
  // TODO: implement
}

func  (me *String) UriEncode() { // TODO: return value
  // TODO: implement
}

func  (me *String) UriDecode() { // TODO: return value
  // TODO: implement
}

func  (me *String) CEscape() { // TODO: return value
  // TODO: implement
}

func  (me *String) CUnescape() { // TODO: return value
  // TODO: implement
}

func  (me *String) JsonEscape() { // TODO: return value
  // TODO: implement
}

func  (me *String) ValidateNodeName() { // TODO: return value
  // TODO: implement
}

func  (me *String) ValidateFilename() { // TODO: return value
  // TODO: implement
}

func  (me *String) IsValidIdentifier() { // TODO: return value
  // TODO: implement
}

func  (me *String) IsValidInt() { // TODO: return value
  // TODO: implement
}

func  (me *String) IsValidFloat() { // TODO: return value
  // TODO: implement
}

func  (me *String) IsValidHexNumber(with_prefix bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) IsValidHtmlColor() { // TODO: return value
  // TODO: implement
}

func  (me *String) IsValidIpAddress() { // TODO: return value
  // TODO: implement
}

func  (me *String) IsValidFilename() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToInt() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToFloat() { // TODO: return value
  // TODO: implement
}

func  (me *String) HexToInt() { // TODO: return value
  // TODO: implement
}

func  (me *String) BinToInt() { // TODO: return value
  // TODO: implement
}

func  (me *String) Lpad(min_length int, character String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) Rpad(min_length int, character String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) PadDecimals(digits int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) PadZeros(digits int, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) TrimPrefix(prefix String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) TrimSuffix(suffix String, ) { // TODO: return value
  // TODO: implement
}

func  (me *String) ToAsciiBuffer() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToUtf8Buffer() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToUtf16Buffer() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToUtf32Buffer() { // TODO: return value
  // TODO: implement
}

func  (me *String) HexDecode() { // TODO: return value
  // TODO: implement
}

func  (me *String) ToWcharBuffer() { // TODO: return value
  // TODO: implement
}

func  StringNumScientific(number float32, ) { // TODO: return value
  // TODO: implement
}

func  StringNum(number float32, decimals int, ) { // TODO: return value
  // TODO: implement
}

func  StringNumInt64(number int, base int, capitalize_hex bool, ) { // TODO: return value
  // TODO: implement
}

func  StringNumUint64(number int, base int, capitalize_hex bool, ) { // TODO: return value
  // TODO: implement
}

func  StringChr(char int, ) { // TODO: return value
  // TODO: implement
}

func  StringHumanizeSize(size int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
