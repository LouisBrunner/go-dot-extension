// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedStringArray struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewPackedStringArray() PackedStringArray {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedStringArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedStringArray, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return PackedStringArray{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedStringArrayFromPackedStringArray(from PackedStringArray, ) PackedStringArray {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedStringArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedStringArray, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedStringArray{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedStringArrayFromArray(from Array, ) PackedStringArray {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedStringArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedStringArray, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedStringArray{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *PackedStringArray) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedStringArray)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *PackedStringArray) Type() gdc.VariantType {
  return gdc.VariantTypePackedStringArray
}

func (me *PackedStringArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *PackedStringArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *PackedStringArray) Size() int {
  name := StringNameFromStr("size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Set(index int, value String, )  {
  name := StringNameFromStr("set")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 725585539) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) PushBack(value String, ) bool {
  name := StringNameFromStr("push_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 816187996) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Append(value String, ) bool {
  name := StringNameFromStr("append")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 816187996) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) AppendArray(array PackedStringArray, )  {
  name := StringNameFromStr("append_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 1120103966) // FIXME: should cache?

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) RemoveAt(index int, )  {
  name := StringNameFromStr("remove_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 2823966027) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Insert(at_index int, value String, ) int {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 2432393153) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&at_index), value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Fill(value String, )  {
  name := StringNameFromStr("fill")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 3174917410) // FIXME: should cache?

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Resize(new_size int, ) int {
  name := StringNameFromStr("resize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 848867239) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&new_size), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Clear()  {
  name := StringNameFromStr("clear")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Has(value String, ) bool {
  name := StringNameFromStr("has")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 2566493496) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Reverse()  {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Slice(begin int, end int, ) PackedStringArray {
  name := StringNameFromStr("slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 2094601407) // FIXME: should cache?

  var ret PackedStringArray
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin), gdc.ConstTypePtr(&end), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) ToByteArray() PackedByteArray {
  name := StringNameFromStr("to_byte_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Sort()  {
  name := StringNameFromStr("sort")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedStringArray) Bsearch(value String, before bool, ) int {
  name := StringNameFromStr("bsearch")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 328976671) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(&before), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Duplicate() PackedStringArray {
  name := StringNameFromStr("duplicate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 2991231410) // FIXME: should cache?

  var ret PackedStringArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Find(value String, from int, ) int {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 1760645412) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Rfind(value String, from int, ) int {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 1760645412) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedStringArray) Count(value String, ) int {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedStringArray, name.AsCPtr(), 2920860731) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *PackedStringArray) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedStringArray) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedStringArray) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *PackedStringArray) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedStringArray) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedStringArray) EqualPackedStringArray(right PackedStringArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedStringArray) NotEqualPackedStringArray(right PackedStringArray) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedStringArray) AddPackedStringArray(right PackedStringArray) PackedStringArray {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret PackedStringArray
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
