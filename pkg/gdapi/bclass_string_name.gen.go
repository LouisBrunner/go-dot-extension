// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "fmt"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StringName struct {
  data   *[classSizeStringName]byte
  iface  gdc.Interface
  pinner runtime.Pinner
}

// Enums

// Constructors
func newStringName() *StringName {
  me := &StringName{
    data:   new([classSizeStringName]byte),
    iface:  giface,
  }
  me.pinner.Pin(me)
  me.pinner.Pin(me.AsTypePtr())
  return me
}

func NewStringName() *StringName {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newStringName()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 0) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{}))
  return me
}

func NewStringNameFromStringName(from StringName, ) *StringName {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newStringName()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 1) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

func NewStringNameFromString(from String, ) *StringName {
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  me := newStringName()
  ctr := me.iface.VariantGetPtrConstructor(gdc.VariantTypeStringName, 2) // FIXME: should cache?
  me.iface.CallPtrConstructor(ctr, me.asUninitialized(), unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return me
}

// Destructor
func (me *StringName) Destroy() {
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeStringName)
	me.iface.CallPtrDestructor(dtr, me.AsTypePtr())
  me.pinner.Unpin()
}

// Variant support
func (me *Variant) AsStringName() (*StringName, error) {
	if me.Type() != gdc.VariantTypeStringName {
		return nil, fmt.Errorf("variant is not a StringName")
	}
  bclass := newStringName()
	fn := me.iface.GetVariantToTypeConstructor(me.Type())
	me.iface.CallTypeFromVariantConstructorFunc(fn, bclass.asUninitialized(), me.AsPtr())
	return bclass, nil
}

func (me *StringName) AsVariant() *Variant {
  va := newVariant()
  va.inner = me
  fn := me.iface.GetVariantFromTypeConstructor(me.Type())
  me.iface.CallVariantFromTypeConstructorFunc(fn, va.asUninitialized(), me.AsTypePtr())
  return va
}

// Pointers
func StringNameFromPtr(ptr gdc.ConstTypePtr) *StringName {
  me := newStringName()
  dataFromPtr(me.data[:], ptr)
  return me
}

func (me *StringName) Type() gdc.VariantType {
  return gdc.VariantTypeStringName
}

func (me *StringName) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(unsafe.Pointer(me.data))
}

func (me *StringName) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.AsTypePtr())
}

func (me *StringName) asUninitialized() gdc.UninitializedTypePtr {
  return gdc.UninitializedTypePtr(me.AsTypePtr())
}

// Methods

func (me *StringName) CasecmpTo(to String, ) int64 {
  name := StringNameFromStr("casecmp_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) NocasecmpTo(to String, ) int64 {
  name := StringNameFromStr("nocasecmp_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) NaturalcasecmpTo(to String, ) int64 {
  name := StringNameFromStr("naturalcasecmp_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) NaturalnocasecmpTo(to String, ) int64 {
  name := StringNameFromStr("naturalnocasecmp_to")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{to.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Length() int64 {
  name := StringNameFromStr("length")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Substr(from int64, len_ int64, ) String {
  name := StringNameFromStr("substr")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 787537301) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(from)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(len_)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) GetSlice(delimiter String, slice int64, ) String {
  name := StringNameFromStr("get_slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3535100402) // FIXME: should cache?

  ret := NewString()


  varg1 := NewIntFromInt(slice)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) GetSlicec(delimiter int64, slice int64, ) String {
  name := StringNameFromStr("get_slicec")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 787537301) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(delimiter)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(slice)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) GetSliceCount(delimiter String, ) int64 {
  name := StringNameFromStr("get_slice_count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2920860731) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Find(what String, from int64, ) int64 {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1760645412) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Count(what String, from int64, to int64, ) int64 {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2343087891) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  varg2 := NewIntFromInt(to)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Countn(what String, from int64, to int64, ) int64 {
  name := StringNameFromStr("countn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2343087891) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  varg2 := NewIntFromInt(to)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Findn(what String, from int64, ) int64 {
  name := StringNameFromStr("findn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1760645412) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Rfind(what String, from int64, ) int64 {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1760645412) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Rfindn(what String, from int64, ) int64 {
  name := StringNameFromStr("rfindn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1760645412) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()

  varg1 := NewIntFromInt(from)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Match(expr String, ) bool {
  name := StringNameFromStr("match")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{expr.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Matchn(expr String, ) bool {
  name := StringNameFromStr("matchn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{expr.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) BeginsWith(text String, ) bool {
  name := StringNameFromStr("begins_with")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) EndsWith(text String, ) bool {
  name := StringNameFromStr("ends_with")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsSubsequenceOf(text String, ) bool {
  name := StringNameFromStr("is_subsequence_of")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsSubsequenceOfn(text String, ) bool {
  name := StringNameFromStr("is_subsequence_ofn")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Bigrams() PackedStringArray {
  name := StringNameFromStr("bigrams")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 747180633) // FIXME: should cache?

  ret := NewPackedStringArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Similarity(text String, ) float64 {
  name := StringNameFromStr("similarity")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2697460964) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{text.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Format(values Variant, placeholder String, ) String {
  name := StringNameFromStr("format")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3212199029) // FIXME: should cache?

  ret := NewString()



  args := []gdc.ConstTypePtr{values.AsCTypePtr(), placeholder.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Replace(what String, forwhat String, ) String {
  name := StringNameFromStr("replace")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1340436205) // FIXME: should cache?

  ret := NewString()



  args := []gdc.ConstTypePtr{what.AsCTypePtr(), forwhat.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Replacen(what String, forwhat String, ) String {
  name := StringNameFromStr("replacen")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1340436205) // FIXME: should cache?

  ret := NewString()



  args := []gdc.ConstTypePtr{what.AsCTypePtr(), forwhat.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Repeat(count int64, ) String {
  name := StringNameFromStr("repeat")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(count)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Reverse() String {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Insert(position int64, what String, ) String {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 248737229) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(position)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), what.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Erase(position int64, chars int64, ) String {
  name := StringNameFromStr("erase")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 787537301) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(position)
  defer varg0.Destroy()
  varg1 := NewIntFromInt(chars)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Capitalize() String {
  name := StringNameFromStr("capitalize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToCamelCase() String {
  name := StringNameFromStr("to_camel_case")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToPascalCase() String {
  name := StringNameFromStr("to_pascal_case")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToSnakeCase() String {
  name := StringNameFromStr("to_snake_case")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Split(delimiter String, allow_empty bool, maxsplit int64, ) PackedStringArray {
  name := StringNameFromStr("split")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1252735785) // FIXME: should cache?

  ret := NewPackedStringArray()


  varg1 := NewBoolFromBool(allow_empty)
  defer varg1.Destroy()
  varg2 := NewIntFromInt(maxsplit)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Rsplit(delimiter String, allow_empty bool, maxsplit int64, ) PackedStringArray {
  name := StringNameFromStr("rsplit")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 1252735785) // FIXME: should cache?

  ret := NewPackedStringArray()


  varg1 := NewBoolFromBool(allow_empty)
  defer varg1.Destroy()
  varg2 := NewIntFromInt(maxsplit)
  defer varg2.Destroy()
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), varg2.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) SplitFloats(delimiter String, allow_empty bool, ) PackedFloat64Array {
  name := StringNameFromStr("split_floats")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2092079095) // FIXME: should cache?

  ret := NewPackedFloat64Array()


  varg1 := NewBoolFromBool(allow_empty)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{delimiter.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Join(parts PackedStringArray, ) String {
  name := StringNameFromStr("join")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3595973238) // FIXME: should cache?

  ret := NewString()


  args := []gdc.ConstTypePtr{parts.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToUpper() String {
  name := StringNameFromStr("to_upper")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToLower() String {
  name := StringNameFromStr("to_lower")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Left(length int64, ) String {
  name := StringNameFromStr("left")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(length)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Right(length int64, ) String {
  name := StringNameFromStr("right")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(length)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) StripEdges(left bool, right bool, ) String {
  name := StringNameFromStr("strip_edges")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 907855311) // FIXME: should cache?

  ret := NewString()

  varg0 := NewBoolFromBool(left)
  defer varg0.Destroy()
  varg1 := NewBoolFromBool(right)
  defer varg1.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), varg1.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) StripEscapes() String {
  name := StringNameFromStr("strip_escapes")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Lstrip(chars String, ) String {
  name := StringNameFromStr("lstrip")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  ret := NewString()


  args := []gdc.ConstTypePtr{chars.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Rstrip(chars String, ) String {
  name := StringNameFromStr("rstrip")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  ret := NewString()


  args := []gdc.ConstTypePtr{chars.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) GetExtension() String {
  name := StringNameFromStr("get_extension")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) GetBasename() String {
  name := StringNameFromStr("get_basename")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) PathJoin(file String, ) String {
  name := StringNameFromStr("path_join")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  ret := NewString()


  args := []gdc.ConstTypePtr{file.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) UnicodeAt(at int64, ) int64 {
  name := StringNameFromStr("unicode_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 4103005248) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  varg0 := NewIntFromInt(at)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Indent(prefix String, ) String {
  name := StringNameFromStr("indent")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  ret := NewString()


  args := []gdc.ConstTypePtr{prefix.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Dedent() String {
  name := StringNameFromStr("dedent")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Md5Text() String {
  name := StringNameFromStr("md5_text")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Sha1Text() String {
  name := StringNameFromStr("sha1_text")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Sha256Text() String {
  name := StringNameFromStr("sha256_text")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Md5Buffer() PackedByteArray {
  name := StringNameFromStr("md5_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Sha1Buffer() PackedByteArray {
  name := StringNameFromStr("sha1_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Sha256Buffer() PackedByteArray {
  name := StringNameFromStr("sha256_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Contains(what String, ) bool {
  name := StringNameFromStr("contains")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2566493496) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()

  args := []gdc.ConstTypePtr{what.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsAbsolutePath() bool {
  name := StringNameFromStr("is_absolute_path")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsRelativePath() bool {
  name := StringNameFromStr("is_relative_path")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) SimplifyPath() String {
  name := StringNameFromStr("simplify_path")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) GetBaseDir() String {
  name := StringNameFromStr("get_base_dir")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) GetFile() String {
  name := StringNameFromStr("get_file")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) XmlEscape(escape_quotes bool, ) String {
  name := StringNameFromStr("xml_escape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3429816538) // FIXME: should cache?

  ret := NewString()

  varg0 := NewBoolFromBool(escape_quotes)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) XmlUnescape() String {
  name := StringNameFromStr("xml_unescape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) UriEncode() String {
  name := StringNameFromStr("uri_encode")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) UriDecode() String {
  name := StringNameFromStr("uri_decode")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) CEscape() String {
  name := StringNameFromStr("c_escape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) CUnescape() String {
  name := StringNameFromStr("c_unescape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) JsonEscape() String {
  name := StringNameFromStr("json_escape")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ValidateNodeName() String {
  name := StringNameFromStr("validate_node_name")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ValidateFilename() String {
  name := StringNameFromStr("validate_filename")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3942272618) // FIXME: should cache?

  ret := NewString()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) IsValidIdentifier() bool {
  name := StringNameFromStr("is_valid_identifier")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsValidInt() bool {
  name := StringNameFromStr("is_valid_int")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsValidFloat() bool {
  name := StringNameFromStr("is_valid_float")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsValidHexNumber(with_prefix bool, ) bool {
  name := StringNameFromStr("is_valid_hex_number")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 593672999) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  varg0 := NewBoolFromBool(with_prefix)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsValidHtmlColor() bool {
  name := StringNameFromStr("is_valid_html_color")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsValidIpAddress() bool {
  name := StringNameFromStr("is_valid_ip_address")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) IsValidFilename() bool {
  name := StringNameFromStr("is_valid_filename")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3918633141) // FIXME: should cache?

  ret := NewBool()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) ToInt() int64 {
  name := StringNameFromStr("to_int")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) ToFloat() float64 {
  name := StringNameFromStr("to_float")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 466405837) // FIXME: should cache?

  ret := NewFloat()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) HexToInt() int64 {
  name := StringNameFromStr("hex_to_int")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) BinToInt() int64 {
  name := StringNameFromStr("bin_to_int")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

func (me *StringName) Lpad(min_length int64, character String, ) String {
  name := StringNameFromStr("lpad")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 248737229) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(min_length)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), character.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Rpad(min_length int64, character String, ) String {
  name := StringNameFromStr("rpad")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 248737229) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(min_length)
  defer varg0.Destroy()

  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), character.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) PadDecimals(digits int64, ) String {
  name := StringNameFromStr("pad_decimals")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(digits)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) PadZeros(digits int64, ) String {
  name := StringNameFromStr("pad_zeros")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 2162347432) // FIXME: should cache?

  ret := NewString()

  varg0 := NewIntFromInt(digits)
  defer varg0.Destroy()
  args := []gdc.ConstTypePtr{varg0.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) TrimPrefix(prefix String, ) String {
  name := StringNameFromStr("trim_prefix")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  ret := NewString()


  args := []gdc.ConstTypePtr{prefix.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) TrimSuffix(suffix String, ) String {
  name := StringNameFromStr("trim_suffix")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3134094431) // FIXME: should cache?

  ret := NewString()


  args := []gdc.ConstTypePtr{suffix.AsCTypePtr(), }


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToAsciiBuffer() PackedByteArray {
  name := StringNameFromStr("to_ascii_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToUtf8Buffer() PackedByteArray {
  name := StringNameFromStr("to_utf8_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToUtf16Buffer() PackedByteArray {
  name := StringNameFromStr("to_utf16_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToUtf32Buffer() PackedByteArray {
  name := StringNameFromStr("to_utf32_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) HexDecode() PackedByteArray {
  name := StringNameFromStr("hex_decode")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) ToWcharBuffer() PackedByteArray {
  name := StringNameFromStr("to_wchar_buffer")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 247621236) // FIXME: should cache?

  ret := NewPackedByteArray()

  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return *ret
}

func (me *StringName) Hash() int64 {
  name := StringNameFromStr("hash")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeStringName, name.AsCPtr(), 3173160232) // FIXME: should cache?

  ret := NewInt()
  defer ret.Destroy()
  args := []gdc.ConstTypePtr{}


  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), ret.AsTypePtr(), len(args))
  return ret.Get()
}

// Operators

func (me *StringName) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) ModuleVariant(right Variant) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) ModuleBool(rightArg bool) String {
  right := NewBoolFromBool(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleInt(rightArg int64) String {
  right := NewIntFromInt(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleFloat32(rightArg float64) String {
  right := NewFloatFromFloat32(rightArg)
  defer right.Destroy()

  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) EqualString(right String) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) NotEqualString(right String) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) AddString(right String) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleString(right String) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) InString(right String) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) ModuleVector2(right Vector2) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleVector2I(right Vector2i) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleRect2(right Rect2) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleRect2I(right Rect2i) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleVector3(right Vector3) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleVector3I(right Vector3i) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleTransform2D(right Transform2D) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleVector4(right Vector4) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleVector4I(right Vector4i) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModulePlane(right Plane) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleQuaternion(right Quaternion) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleAABB(right AABB) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleBasis(right Basis) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleTransform3D(right Transform3D) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleProjection(right Projection) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleColor(right Color) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) EqualStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) NotEqualStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) LessStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) LessEqualStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) GreaterStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) GreaterEqualStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) AddStringName(right StringName) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleStringName(right StringName) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) InStringName(right StringName) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) ModuleNodePath(right NodePath) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleObject(right Object) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) InObject(right Object) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) ModuleCallable(right Callable) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleSignal(right Signal) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModuleDictionary(right Dictionary) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) ModuleArray(right Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) ModulePackedByteArray(right PackedByteArray) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModulePackedInt32Array(right PackedInt32Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModulePackedInt64Array(right PackedInt64Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModulePackedFloat32Array(right PackedFloat32Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModulePackedFloat64Array(right PackedFloat64Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModulePackedStringArray(right PackedStringArray) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) InPackedStringArray(right PackedStringArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  ret := NewBool()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return ret.Get()
}

func (me *StringName) ModulePackedVector2Array(right PackedVector2Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModulePackedVector3Array(right PackedVector3Array) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

func (me *StringName) ModulePackedColorArray(right PackedColorArray) String {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpModule, me.Type(), right.Type()) // FIXME: cache
  ret := NewString()
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), ret.AsTypePtr())
  return *ret
}

// Members
