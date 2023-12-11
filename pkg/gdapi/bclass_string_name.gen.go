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

func (me *StringName) CasecmpTo(to String, ) int {
  name := StringNameFromStr("casecmp_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) NocasecmpTo(to String, ) int {
  name := StringNameFromStr("nocasecmp_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) NaturalcasecmpTo(to String, ) int {
  name := StringNameFromStr("naturalcasecmp_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) NaturalnocasecmpTo(to String, ) int {
  name := StringNameFromStr("naturalnocasecmp_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Length() int {
  name := StringNameFromStr("length")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Substr(from int, len_ int, ) String {
  name := StringNameFromStr("substr")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 787537301) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&len_), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) GetSlice(delimiter String, slice int, ) String {
  name := StringNameFromStr("get_slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3535100402) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), gdc.ConstTypePtr(&slice), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) GetSlicec(delimiter int, slice int, ) String {
  name := StringNameFromStr("get_slicec")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 787537301) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delimiter), gdc.ConstTypePtr(&slice), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) GetSliceCount(delimiter String, ) int {
  name := StringNameFromStr("get_slice_count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Find(what String, from int, ) int {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1760645412) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Count(what String, from int, to int, ) int {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2343087891) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Countn(what String, from int, to int, ) int {
  name := StringNameFromStr("countn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2343087891) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), gdc.ConstTypePtr(&from), gdc.ConstTypePtr(&to), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Findn(what String, from int, ) int {
  name := StringNameFromStr("findn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1760645412) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Rfind(what String, from int, ) int {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1760645412) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Rfindn(what String, from int, ) int {
  name := StringNameFromStr("rfindn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1760645412) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Match(expr String, ) bool {
  name := StringNameFromStr("match")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{expr.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Matchn(expr String, ) bool {
  name := StringNameFromStr("matchn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{expr.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) BeginsWith(text String, ) bool {
  name := StringNameFromStr("begins_with")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) EndsWith(text String, ) bool {
  name := StringNameFromStr("ends_with")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsSubsequenceOf(text String, ) bool {
  name := StringNameFromStr("is_subsequence_of")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsSubsequenceOfn(text String, ) bool {
  name := StringNameFromStr("is_subsequence_ofn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Bigrams() PackedStringArray {
  name := StringNameFromStr("bigrams")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 747180633) // FIXME: should cache?

  var ret PackedStringArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Similarity(text String, ) float32 {
  name := StringNameFromStr("similarity")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2697460964) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Format(values Variant, placeholder String, ) String {
  name := StringNameFromStr("format")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3212199029) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{values.AsCTypePtr(), placeholder.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Replace(what String, forwhat String, ) String {
  name := StringNameFromStr("replace")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1340436205) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), forwhat.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Replacen(what String, forwhat String, ) String {
  name := StringNameFromStr("replacen")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1340436205) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), forwhat.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Repeat(count int, ) String {
  name := StringNameFromStr("repeat")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Insert(position int, what String, ) String {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 248737229) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), what.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Erase(position int, chars int, ) String {
  name := StringNameFromStr("erase")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 787537301) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), gdc.ConstTypePtr(&chars), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Capitalize() String {
  name := StringNameFromStr("capitalize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToCamelCase() String {
  name := StringNameFromStr("to_camel_case")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToPascalCase() String {
  name := StringNameFromStr("to_pascal_case")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToSnakeCase() String {
  name := StringNameFromStr("to_snake_case")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Split(delimiter String, allow_empty bool, maxsplit int, ) PackedStringArray {
  name := StringNameFromStr("split")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1252735785) // FIXME: should cache?

  var ret PackedStringArray
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), gdc.ConstTypePtr(&allow_empty), gdc.ConstTypePtr(&maxsplit), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Rsplit(delimiter String, allow_empty bool, maxsplit int, ) PackedStringArray {
  name := StringNameFromStr("rsplit")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1252735785) // FIXME: should cache?

  var ret PackedStringArray
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), gdc.ConstTypePtr(&allow_empty), gdc.ConstTypePtr(&maxsplit), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) SplitFloats(delimiter String, allow_empty bool, ) PackedFloat64Array {
  name := StringNameFromStr("split_floats")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2092079095) // FIXME: should cache?

  var ret PackedFloat64Array
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), gdc.ConstTypePtr(&allow_empty), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Join(parts PackedStringArray, ) String {
  name := StringNameFromStr("join")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3595973238) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{parts.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToUpper() String {
  name := StringNameFromStr("to_upper")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToLower() String {
  name := StringNameFromStr("to_lower")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Left(length int, ) String {
  name := StringNameFromStr("left")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Right(length int, ) String {
  name := StringNameFromStr("right")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) StripEdges(left bool, right bool, ) String {
  name := StringNameFromStr("strip_edges")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 907855311) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&left), gdc.ConstTypePtr(&right), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) StripEscapes() String {
  name := StringNameFromStr("strip_escapes")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Lstrip(chars String, ) String {
  name := StringNameFromStr("lstrip")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{chars.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Rstrip(chars String, ) String {
  name := StringNameFromStr("rstrip")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{chars.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) GetExtension() String {
  name := StringNameFromStr("get_extension")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) GetBasename() String {
  name := StringNameFromStr("get_basename")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) PathJoin(file String, ) String {
  name := StringNameFromStr("path_join")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{file.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) UnicodeAt(at int, ) int {
  name := StringNameFromStr("unicode_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&at), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Indent(prefix String, ) String {
  name := StringNameFromStr("indent")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{prefix.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Dedent() String {
  name := StringNameFromStr("dedent")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Md5Text() String {
  name := StringNameFromStr("md5_text")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Sha1Text() String {
  name := StringNameFromStr("sha1_text")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Sha256Text() String {
  name := StringNameFromStr("sha256_text")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Md5Buffer() PackedByteArray {
  name := StringNameFromStr("md5_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Sha1Buffer() PackedByteArray {
  name := StringNameFromStr("sha1_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Sha256Buffer() PackedByteArray {
  name := StringNameFromStr("sha256_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Contains(what String, ) bool {
  name := StringNameFromStr("contains")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsAbsolutePath() bool {
  name := StringNameFromStr("is_absolute_path")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsRelativePath() bool {
  name := StringNameFromStr("is_relative_path")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) SimplifyPath() String {
  name := StringNameFromStr("simplify_path")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) GetBaseDir() String {
  name := StringNameFromStr("get_base_dir")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) GetFile() String {
  name := StringNameFromStr("get_file")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) XmlEscape(escape_quotes bool, ) String {
  name := StringNameFromStr("xml_escape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3429816538) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&escape_quotes), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) XmlUnescape() String {
  name := StringNameFromStr("xml_unescape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) UriEncode() String {
  name := StringNameFromStr("uri_encode")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) UriDecode() String {
  name := StringNameFromStr("uri_decode")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) CEscape() String {
  name := StringNameFromStr("c_escape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) CUnescape() String {
  name := StringNameFromStr("c_unescape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) JsonEscape() String {
  name := StringNameFromStr("json_escape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ValidateNodeName() String {
  name := StringNameFromStr("validate_node_name")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ValidateFilename() String {
  name := StringNameFromStr("validate_filename")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsValidIdentifier() bool {
  name := StringNameFromStr("is_valid_identifier")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsValidInt() bool {
  name := StringNameFromStr("is_valid_int")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsValidFloat() bool {
  name := StringNameFromStr("is_valid_float")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsValidHexNumber(with_prefix bool, ) bool {
  name := StringNameFromStr("is_valid_hex_number")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 593672999) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&with_prefix), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsValidHtmlColor() bool {
  name := StringNameFromStr("is_valid_html_color")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsValidIpAddress() bool {
  name := StringNameFromStr("is_valid_ip_address")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) IsValidFilename() bool {
  name := StringNameFromStr("is_valid_filename")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToInt() int {
  name := StringNameFromStr("to_int")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToFloat() float32 {
  name := StringNameFromStr("to_float")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 466405837) // FIXME: should cache?

  var ret float32
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) HexToInt() int {
  name := StringNameFromStr("hex_to_int")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) BinToInt() int {
  name := StringNameFromStr("bin_to_int")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Lpad(min_length int, character String, ) String {
  name := StringNameFromStr("lpad")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 248737229) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_length), character.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Rpad(min_length int, character String, ) String {
  name := StringNameFromStr("rpad")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 248737229) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_length), character.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) PadDecimals(digits int, ) String {
  name := StringNameFromStr("pad_decimals")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&digits), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) PadZeros(digits int, ) String {
  name := StringNameFromStr("pad_zeros")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&digits), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) TrimPrefix(prefix String, ) String {
  name := StringNameFromStr("trim_prefix")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{prefix.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) TrimSuffix(suffix String, ) String {
  name := StringNameFromStr("trim_suffix")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  var ret String
  args := []gdc.ConstTypePtr{suffix.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToAsciiBuffer() PackedByteArray {
  name := StringNameFromStr("to_ascii_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToUtf8Buffer() PackedByteArray {
  name := StringNameFromStr("to_utf8_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToUtf16Buffer() PackedByteArray {
  name := StringNameFromStr("to_utf16_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToUtf32Buffer() PackedByteArray {
  name := StringNameFromStr("to_utf32_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) HexDecode() PackedByteArray {
  name := StringNameFromStr("hex_decode")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) ToWcharBuffer() PackedByteArray {
  name := StringNameFromStr("to_wchar_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *StringName) Hash() int {
  name := StringNameFromStr("hash")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *StringName) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleVariant(right Variant) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleBool(right Bool) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleInt(right Int) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleFloat32(right Float) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) EqualString(right String) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) NotEqualString(right String) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) AddString(right String) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleString(right String) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) InString(right String) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleVector2(right Vector2) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleVector2I(right Vector2i) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleRect2(right Rect2) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleRect2I(right Rect2i) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleVector3(right Vector3) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleVector3I(right Vector3i) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleTransform2D(right Transform2D) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleVector4(right Vector4) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleVector4I(right Vector4i) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePlane(right Plane) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleQuaternion(right Quaternion) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleAABB(right AABB) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleBasis(right Basis) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleTransform3D(right Transform3D) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleProjection(right Projection) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleColor(right Color) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) EqualStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) NotEqualStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) LessStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) LessEqualStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) GreaterStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) GreaterEqualStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) AddStringName(right StringName) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleStringName(right StringName) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) InStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleNodePath(right NodePath) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleObject(right Object) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) InObject(right Object) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleCallable(right Callable) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleSignal(right Signal) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleDictionary(right Dictionary) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModuleArray(right Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePackedByteArray(right PackedByteArray) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePackedInt32Array(right PackedInt32Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePackedInt64Array(right PackedInt64Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePackedFloat32Array(right PackedFloat32Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePackedFloat64Array(right PackedFloat64Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePackedStringArray(right PackedStringArray) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) InPackedStringArray(right PackedStringArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePackedVector2Array(right PackedVector2Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePackedVector3Array(right PackedVector3Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *StringName) ModulePackedColorArray(right PackedColorArray) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  var ret String
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
