// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StringName struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewStringName() StringName {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeStringName))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return StringName{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewStringNameFromStringName(from StringName, ) StringName {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeStringName))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return StringName{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewStringNameFromString(from String, ) StringName {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeStringName))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return StringName{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *StringName) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeStringName)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *StringName) Type() gdc.VariantType {
  return gdc.VariantTypeStringName
}

func (me *StringName) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *StringName) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *StringName) CasecmpTo(to String, ) int {
  panic("TODO: implement")
}

func  (me *StringName) NocasecmpTo(to String, ) int {
  panic("TODO: implement")
}

func  (me *StringName) NaturalcasecmpTo(to String, ) int {
  panic("TODO: implement")
}

func  (me *StringName) NaturalnocasecmpTo(to String, ) int {
  panic("TODO: implement")
}

func  (me *StringName) Length() int {
  panic("TODO: implement")
}

func  (me *StringName) Substr(from int, len int, ) String {
  panic("TODO: implement")
}

func  (me *StringName) GetSlice(delimiter String, slice int, ) String {
  panic("TODO: implement")
}

func  (me *StringName) GetSlicec(delimiter int, slice int, ) String {
  panic("TODO: implement")
}

func  (me *StringName) GetSliceCount(delimiter String, ) int {
  panic("TODO: implement")
}

func  (me *StringName) Find(what String, from int, ) int {
  panic("TODO: implement")
}

func  (me *StringName) Count(what String, from int, to int, ) int {
  panic("TODO: implement")
}

func  (me *StringName) Countn(what String, from int, to int, ) int {
  panic("TODO: implement")
}

func  (me *StringName) Findn(what String, from int, ) int {
  panic("TODO: implement")
}

func  (me *StringName) Rfind(what String, from int, ) int {
  panic("TODO: implement")
}

func  (me *StringName) Rfindn(what String, from int, ) int {
  panic("TODO: implement")
}

func  (me *StringName) Match(expr String, ) bool {
  panic("TODO: implement")
}

func  (me *StringName) Matchn(expr String, ) bool {
  panic("TODO: implement")
}

func  (me *StringName) BeginsWith(text String, ) bool {
  panic("TODO: implement")
}

func  (me *StringName) EndsWith(text String, ) bool {
  panic("TODO: implement")
}

func  (me *StringName) IsSubsequenceOf(text String, ) bool {
  panic("TODO: implement")
}

func  (me *StringName) IsSubsequenceOfn(text String, ) bool {
  panic("TODO: implement")
}

func  (me *StringName) Bigrams() PackedStringArray {
  panic("TODO: implement")
}

func  (me *StringName) Similarity(text String, ) float32 {
  panic("TODO: implement")
}

func  (me *StringName) Format(values Variant, placeholder String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Replace(what String, forwhat String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Replacen(what String, forwhat String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Repeat(count int, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Insert(position int, what String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Erase(position int, chars int, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Capitalize() String {
  panic("TODO: implement")
}

func  (me *StringName) ToCamelCase() String {
  panic("TODO: implement")
}

func  (me *StringName) ToPascalCase() String {
  panic("TODO: implement")
}

func  (me *StringName) ToSnakeCase() String {
  panic("TODO: implement")
}

func  (me *StringName) Split(delimiter String, allow_empty bool, maxsplit int, ) PackedStringArray {
  panic("TODO: implement")
}

func  (me *StringName) Rsplit(delimiter String, allow_empty bool, maxsplit int, ) PackedStringArray {
  panic("TODO: implement")
}

func  (me *StringName) SplitFloats(delimiter String, allow_empty bool, ) PackedFloat64Array {
  panic("TODO: implement")
}

func  (me *StringName) Join(parts PackedStringArray, ) String {
  panic("TODO: implement")
}

func  (me *StringName) ToUpper() String {
  panic("TODO: implement")
}

func  (me *StringName) ToLower() String {
  panic("TODO: implement")
}

func  (me *StringName) Left(length int, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Right(length int, ) String {
  panic("TODO: implement")
}

func  (me *StringName) StripEdges(left bool, right bool, ) String {
  panic("TODO: implement")
}

func  (me *StringName) StripEscapes() String {
  panic("TODO: implement")
}

func  (me *StringName) Lstrip(chars String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Rstrip(chars String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) GetExtension() String {
  panic("TODO: implement")
}

func  (me *StringName) GetBasename() String {
  panic("TODO: implement")
}

func  (me *StringName) PathJoin(file String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) UnicodeAt(at int, ) int {
  panic("TODO: implement")
}

func  (me *StringName) Indent(prefix String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Dedent() String {
  panic("TODO: implement")
}

func  (me *StringName) Md5Text() String {
  panic("TODO: implement")
}

func  (me *StringName) Sha1Text() String {
  panic("TODO: implement")
}

func  (me *StringName) Sha256Text() String {
  panic("TODO: implement")
}

func  (me *StringName) Md5Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *StringName) Sha1Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *StringName) Sha256Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *StringName) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *StringName) Contains(what String, ) bool {
  panic("TODO: implement")
}

func  (me *StringName) IsAbsolutePath() bool {
  panic("TODO: implement")
}

func  (me *StringName) IsRelativePath() bool {
  panic("TODO: implement")
}

func  (me *StringName) SimplifyPath() String {
  panic("TODO: implement")
}

func  (me *StringName) GetBaseDir() String {
  panic("TODO: implement")
}

func  (me *StringName) GetFile() String {
  panic("TODO: implement")
}

func  (me *StringName) XmlEscape(escape_quotes bool, ) String {
  panic("TODO: implement")
}

func  (me *StringName) XmlUnescape() String {
  panic("TODO: implement")
}

func  (me *StringName) UriEncode() String {
  panic("TODO: implement")
}

func  (me *StringName) UriDecode() String {
  panic("TODO: implement")
}

func  (me *StringName) CEscape() String {
  panic("TODO: implement")
}

func  (me *StringName) CUnescape() String {
  panic("TODO: implement")
}

func  (me *StringName) JsonEscape() String {
  panic("TODO: implement")
}

func  (me *StringName) ValidateNodeName() String {
  panic("TODO: implement")
}

func  (me *StringName) ValidateFilename() String {
  panic("TODO: implement")
}

func  (me *StringName) IsValidIdentifier() bool {
  panic("TODO: implement")
}

func  (me *StringName) IsValidInt() bool {
  panic("TODO: implement")
}

func  (me *StringName) IsValidFloat() bool {
  panic("TODO: implement")
}

func  (me *StringName) IsValidHexNumber(with_prefix bool, ) bool {
  panic("TODO: implement")
}

func  (me *StringName) IsValidHtmlColor() bool {
  panic("TODO: implement")
}

func  (me *StringName) IsValidIpAddress() bool {
  panic("TODO: implement")
}

func  (me *StringName) IsValidFilename() bool {
  panic("TODO: implement")
}

func  (me *StringName) ToInt() int {
  panic("TODO: implement")
}

func  (me *StringName) ToFloat() float32 {
  panic("TODO: implement")
}

func  (me *StringName) HexToInt() int {
  panic("TODO: implement")
}

func  (me *StringName) BinToInt() int {
  panic("TODO: implement")
}

func  (me *StringName) Lpad(min_length int, character String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) Rpad(min_length int, character String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) PadDecimals(digits int, ) String {
  panic("TODO: implement")
}

func  (me *StringName) PadZeros(digits int, ) String {
  panic("TODO: implement")
}

func  (me *StringName) TrimPrefix(prefix String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) TrimSuffix(suffix String, ) String {
  panic("TODO: implement")
}

func  (me *StringName) ToAsciiBuffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *StringName) ToUtf8Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *StringName) ToUtf16Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *StringName) ToUtf32Buffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *StringName) HexDecode() PackedByteArray {
  panic("TODO: implement")
}

func  (me *StringName) ToWcharBuffer() PackedByteArray {
  panic("TODO: implement")
}

func  (me *StringName) Hash() int {
  panic("TODO: implement")
}

// Operators

func (me *StringName) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *StringName) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *StringName) ModuloVariant(right Variant) String {
  panic("TODO: implement")
}

func (me *StringName) Not() bool {
  panic("TODO: implement")
}

func (me *StringName) ModuloBool(right bool) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloInt(right int) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloFloat32(right float32) String {
  panic("TODO: implement")
}

func (me *StringName) EqualsString(right String) bool {
  panic("TODO: implement")
}

func (me *StringName) NotEqualsString(right String) bool {
  panic("TODO: implement")
}

func (me *StringName) AddString(right String) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloString(right String) String {
  panic("TODO: implement")
}

func (me *StringName) InString(right String) bool {
  panic("TODO: implement")
}

func (me *StringName) ModuloVector2(right Vector2) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloVector2I(right Vector2i) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloRect2(right Rect2) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloRect2I(right Rect2i) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloVector3(right Vector3) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloVector3I(right Vector3i) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloTransform2D(right Transform2D) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloVector4(right Vector4) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloVector4I(right Vector4i) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloPlane(right Plane) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloQuaternion(right Quaternion) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloAABB(right AABB) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloBasis(right Basis) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloTransform3D(right Transform3D) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloProjection(right Projection) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloColor(right Color) String {
  panic("TODO: implement")
}

func (me *StringName) EqualsStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *StringName) NotEqualsStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *StringName) LessThanStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *StringName) LessThanOrEqualsStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *StringName) GreaterThanStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *StringName) GreaterThanOrEqualsStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *StringName) AddStringName(right StringName) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloStringName(right StringName) String {
  panic("TODO: implement")
}

func (me *StringName) InStringName(right StringName) bool {
  panic("TODO: implement")
}

func (me *StringName) ModuloNodePath(right NodePath) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloObject(right Object) String {
  panic("TODO: implement")
}

func (me *StringName) InObject(right Object) bool {
  panic("TODO: implement")
}

func (me *StringName) ModuloCallable(right Callable) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloSignal(right Signal) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloDictionary(right Dictionary) String {
  panic("TODO: implement")
}

func (me *StringName) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *StringName) ModuloArray(right Array) String {
  panic("TODO: implement")
}

func (me *StringName) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *StringName) ModuloPackedByteArray(right PackedByteArray) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloPackedInt32Array(right PackedInt32Array) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloPackedInt64Array(right PackedInt64Array) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloPackedFloat32Array(right PackedFloat32Array) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloPackedFloat64Array(right PackedFloat64Array) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloPackedStringArray(right PackedStringArray) String {
  panic("TODO: implement")
}

func (me *StringName) InPackedStringArray(right PackedStringArray) bool {
  panic("TODO: implement")
}

func (me *StringName) ModuloPackedVector2Array(right PackedVector2Array) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloPackedVector3Array(right PackedVector3Array) String {
  panic("TODO: implement")
}

func (me *StringName) ModuloPackedColorArray(right PackedColorArray) String {
  panic("TODO: implement")
}

// TODO: members (bclass)
