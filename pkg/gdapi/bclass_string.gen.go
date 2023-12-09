// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type String struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewString() String {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeString))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeString, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return String{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewStringFromString(from String, ) String {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeString))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeString, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return String{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewStringFromStringName(from StringName, ) String {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeString))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeString, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return String{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewStringFromNodePath(from NodePath, ) String {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeString))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeString, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return String{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *String) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeString)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *String) Type() gdc.VariantType {
  return gdc.VariantTypeString
}

func (me *String) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *String) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *String) CasecmpTo(to String, ) int {
  panic("TODO: implement")
}

func  (me *String) NocasecmpTo(to String, ) int {
  panic("TODO: implement")
}

func  (me *String) NaturalcasecmpTo(to String, ) int {
  panic("TODO: implement")
}

func  (me *String) NaturalnocasecmpTo(to String, ) int {
  panic("TODO: implement")
}

func  (me *String) Length() int {
  panic("TODO: implement")
}

func  (me *String) Substr(from int, len int, ) String {
  panic("TODO: implement")
}

func  (me *String) GetSlice(delimiter String, slice int, ) String {
  panic("TODO: implement")
}

func  (me *String) GetSlicec(delimiter int, slice int, ) String {
  panic("TODO: implement")
}

func  (me *String) GetSliceCount(delimiter String, ) int {
  panic("TODO: implement")
}

func  (me *String) Find(what String, from int, ) int {
  panic("TODO: implement")
}

func  (me *String) Count(what String, from int, to int, ) int {
  panic("TODO: implement")
}

func  (me *String) Countn(what String, from int, to int, ) int {
  panic("TODO: implement")
}

func  (me *String) Findn(what String, from int, ) int {
  panic("TODO: implement")
}

func  (me *String) Rfind(what String, from int, ) int {
  panic("TODO: implement")
}

func  (me *String) Rfindn(what String, from int, ) int {
  panic("TODO: implement")
}

func  (me *String) Match(expr String, ) bool {
  panic("TODO: implement")
}

func  (me *String) Matchn(expr String, ) bool {
  panic("TODO: implement")
}

func  (me *String) BeginsWith(text String, ) bool {
  panic("TODO: implement")
}

func  (me *String) EndsWith(text String, ) bool {
  panic("TODO: implement")
}

func  (me *String) IsSubsequenceOf(text String, ) bool {
  panic("TODO: implement")
}

func  (me *String) IsSubsequenceOfn(text String, ) bool {
  panic("TODO: implement")
}

func  (me *String) Bigrams() PackedStringArray {
  panic("TODO: implement")
}

func  (me *String) Similarity(text String, ) float32 {
  panic("TODO: implement")
}

func  (me *String) Format(values Variant, placeholder String, ) String {
  panic("TODO: implement")
}

func  (me *String) Replace(what String, forwhat String, ) String {
  panic("TODO: implement")
}

func  (me *String) Replacen(what String, forwhat String, ) String {
  panic("TODO: implement")
}

func  (me *String) Repeat(count int, ) String {
  panic("TODO: implement")
}

func  (me *String) Insert(position int, what String, ) String {
  panic("TODO: implement")
}

func  (me *String) Erase(position int, chars int, ) String {
  panic("TODO: implement")
}

func  (me *String) Capitalize() String {
  panic("TODO: implement")
}

func  (me *String) ToCamelCase() String {
  panic("TODO: implement")
}

func  (me *String) ToPascalCase() String {
  panic("TODO: implement")
}

func  (me *String) ToSnakeCase() String {
  panic("TODO: implement")
}

func  (me *String) Split(delimiter String, allow_empty bool, maxsplit int, ) PackedStringArray {
  panic("TODO: implement")
}

func  (me *String) Rsplit(delimiter String, allow_empty bool, maxsplit int, ) PackedStringArray {
  panic("TODO: implement")
}

func  (me *String) SplitFloats(delimiter String, allow_empty bool, ) PackedFloat64Array {
  panic("TODO: implement")
}

func  (me *String) Join(parts PackedStringArray, ) String {
  panic("TODO: implement")
}

func  (me *String) ToUpper() String {
  panic("TODO: implement")
}

func  (me *String) ToLower() String {
  panic("TODO: implement")
}

func  (me *String) Left(length int, ) String {
  panic("TODO: implement")
}

func  (me *String) Right(length int, ) String {
  panic("TODO: implement")
}

func  (me *String) StripEdges(left bool, right bool, ) String {
  panic("TODO: implement")
}

func  (me *String) StripEscapes() String {
  panic("TODO: implement")
}

func  (me *String) Lstrip(chars String, ) String {
  panic("TODO: implement")
}

func  (me *String) Rstrip(chars String, ) String {
  panic("TODO: implement")
}

func  (me *String) GetExtension() String {
  panic("TODO: implement")
}

func  (me *String) GetBasename() String {
  panic("TODO: implement")
}

func  (me *String) PathJoin(file String, ) String {
  panic("TODO: implement")
}

func  (me *String) UnicodeAt(at int, ) int {
  panic("TODO: implement")
}

func  (me *String) Indent(prefix String, ) String {
  panic("TODO: implement")
}

func  (me *String) Dedent() String {
  panic("TODO: implement")
}

func  (me *String) Hash() int {
  panic("TODO: implement")
}

func  (me *String) Md5Text() String {
  panic("TODO: implement")
}

func  (me *String) Sha1Text() String {
  panic("TODO: implement")
}

func  (me *String) Sha256Text() String {
  panic("TODO: implement")
}

func  (me *String) Md5Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *String) Sha1Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *String) Sha256Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *String) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *String) Contains(what String, ) bool {
  panic("TODO: implement")
}

func  (me *String) IsAbsolutePath() bool {
  panic("TODO: implement")
}

func  (me *String) IsRelativePath() bool {
  panic("TODO: implement")
}

func  (me *String) SimplifyPath() String {
  panic("TODO: implement")
}

func  (me *String) GetBaseDir() String {
  panic("TODO: implement")
}

func  (me *String) GetFile() String {
  panic("TODO: implement")
}

func  (me *String) XmlEscape(escape_quotes bool, ) String {
  panic("TODO: implement")
}

func  (me *String) XmlUnescape() String {
  panic("TODO: implement")
}

func  (me *String) UriEncode() String {
  panic("TODO: implement")
}

func  (me *String) UriDecode() String {
  panic("TODO: implement")
}

func  (me *String) CEscape() String {
  panic("TODO: implement")
}

func  (me *String) CUnescape() String {
  panic("TODO: implement")
}

func  (me *String) JsonEscape() String {
  panic("TODO: implement")
}

func  (me *String) ValidateNodeName() String {
  panic("TODO: implement")
}

func  (me *String) ValidateFilename() String {
  panic("TODO: implement")
}

func  (me *String) IsValidIdentifier() bool {
  panic("TODO: implement")
}

func  (me *String) IsValidInt() bool {
  panic("TODO: implement")
}

func  (me *String) IsValidFloat() bool {
  panic("TODO: implement")
}

func  (me *String) IsValidHexNumber(with_prefix bool, ) bool {
  panic("TODO: implement")
}

func  (me *String) IsValidHtmlColor() bool {
  panic("TODO: implement")
}

func  (me *String) IsValidIpAddress() bool {
  panic("TODO: implement")
}

func  (me *String) IsValidFilename() bool {
  panic("TODO: implement")
}

func  (me *String) ToInt() int {
  panic("TODO: implement")
}

func  (me *String) ToFloat() float32 {
  panic("TODO: implement")
}

func  (me *String) HexToInt() int {
  panic("TODO: implement")
}

func  (me *String) BinToInt() int {
  panic("TODO: implement")
}

func  (me *String) Lpad(min_length int, character String, ) String {
  panic("TODO: implement")
}

func  (me *String) Rpad(min_length int, character String, ) String {
  panic("TODO: implement")
}

func  (me *String) PadDecimals(digits int, ) String {
  panic("TODO: implement")
}

func  (me *String) PadZeros(digits int, ) String {
  panic("TODO: implement")
}

func  (me *String) TrimPrefix(prefix String, ) String {
  panic("TODO: implement")
}

func  (me *String) TrimSuffix(suffix String, ) String {
  panic("TODO: implement")
}

func  (me *String) ToAsciiBuffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *String) ToUtf8Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *String) ToUtf16Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *String) ToUtf32Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *String) HexDecode() PackedByteArray {
  panic("TODO: implement")
}

func  (me *String) ToWcharBuffer() PackedByteArray {
  panic("TODO: implement")
}

func  StringNumScientific(number float32, ) String {
  panic("TODO: implement")
}

func  StringNum(number float32, decimals int, ) String {
  panic("TODO: implement")
}

func  StringNumInt64(number int, base int, capitalize_hex bool, ) String {
  panic("TODO: implement")
}

func  StringNumUint64(number int, base int, capitalize_hex bool, ) String {
  panic("TODO: implement")
}

func  StringChr(char int, ) String {
  panic("TODO: implement")
}

func  StringHumanizeSize(size int, ) String {
  panic("TODO: implement")
}

// Operators

func (me *String) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *String) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *String) ModuloVariant(right Variant) String {
  panic("TODO: implement")
}

func (me *String) Not() bool {
  panic("TODO: implement")
}

func (me *String) ModuloBool(right bool) String {
  panic("TODO: implement")
}

func (me *String) ModuloInt(right int) String {
  panic("TODO: implement")
}

func (me *String) ModuloFloat32(right float32) String {
  panic("TODO: implement")
}

func (me *String) EqualsString(right String) bool {
  panic("TODO: implement")
}

func (me *String) NotEqualsString(right String) bool {
  panic("TODO: implement")
}

func (me *String) LessThanString(right String) bool {
  panic("TODO: implement")
}

func (me *String) LessThanOrEqualsString(right String) bool {
  panic("TODO: implement")
}

func (me *String) GreaterThanString(right String) bool {
  panic("TODO: implement")
}

func (me *String) GreaterThanOrEqualsString(right String) bool {
  panic("TODO: implement")
}

func (me *String) AddString(right String) String {
  panic("TODO: implement")
}

func (me *String) ModuloString(right String) String {
  panic("TODO: implement")
}

func (me *String) InString(right String) bool {
  panic("TODO: implement")
}

func (me *String) ModuloVector2(right Vector2) String {
  panic("TODO: implement")
}

func (me *String) ModuloVector2I(right Vector2i) String {
  panic("TODO: implement")
}

func (me *String) ModuloRect2(right Rect2) String {
  panic("TODO: implement")
}

func (me *String) ModuloRect2I(right Rect2i) String {
  panic("TODO: implement")
}

func (me *String) ModuloVector3(right Vector3) String {
  panic("TODO: implement")
}

func (me *String) ModuloVector3I(right Vector3i) String {
  panic("TODO: implement")
}

func (me *String) ModuloTransform2D(right Transform2D) String {
  panic("TODO: implement")
}

func (me *String) ModuloVector4(right Vector4) String {
  panic("TODO: implement")
}

func (me *String) ModuloVector4I(right Vector4i) String {
  panic("TODO: implement")
}

func (me *String) ModuloPlane(right Plane) String {
  panic("TODO: implement")
}

func (me *String) ModuloQuaternion(right Quaternion) String {
  panic("TODO: implement")
}

func (me *String) ModuloAABB(right AABB) String {
  panic("TODO: implement")
}

func (me *String) ModuloBasis(right Basis) String {
  panic("TODO: implement")
}

func (me *String) ModuloTransform3D(right Transform3D) String {
  panic("TODO: implement")
}

func (me *String) ModuloProjection(right Projection) String {
  panic("TODO: implement")
}

func (me *String) ModuloColor(right Color) String {
  panic("TODO: implement")
}

func (me *String) EqualsStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *String) NotEqualsStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *String) AddStringName(right StringName) String {
  panic("TODO: implement")
}

func (me *String) ModuloStringName(right StringName) String {
  panic("TODO: implement")
}

func (me *String) InStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *String) ModuloNodePath(right NodePath) String {
  panic("TODO: implement")
}

func (me *String) ModuloObject(right Object) String {
  panic("TODO: implement")
}

func (me *String) InObject(right Object) bool {
  panic("TODO: implement")
}

func (me *String) ModuloCallable(right Callable) String {
  panic("TODO: implement")
}

func (me *String) ModuloSignal(right Signal) String {
  panic("TODO: implement")
}

func (me *String) ModuloDictionary(right Dictionary) String {
  panic("TODO: implement")
}

func (me *String) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *String) ModuloArray(right Array) String {
  panic("TODO: implement")
}

func (me *String) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *String) ModuloPackedByteArray(right PackedByteArray) String {
  panic("TODO: implement")
}

func (me *String) ModuloPackedInt32Array(right PackedInt32Array) String {
  panic("TODO: implement")
}

func (me *String) ModuloPackedInt64Array(right PackedInt64Array) String {
  panic("TODO: implement")
}

func (me *String) ModuloPackedFloat32Array(right PackedFloat32Array) String {
  panic("TODO: implement")
}

func (me *String) ModuloPackedFloat64Array(right PackedFloat64Array) String {
  panic("TODO: implement")
}

func (me *String) ModuloPackedStringArray(right PackedStringArray) String {
  panic("TODO: implement")
}

func (me *String) InPackedStringArray(right PackedStringArray) bool {
  panic("TODO: implement")
}

func (me *String) ModuloPackedVector2Array(right PackedVector2Array) String {
  panic("TODO: implement")
}

func (me *String) ModuloPackedVector3Array(right PackedVector3Array) String {
  panic("TODO: implement")
}

func (me *String) ModuloPackedColorArray(right PackedColorArray) String {
  panic("TODO: implement")
}

// TODO: members (bclass)
