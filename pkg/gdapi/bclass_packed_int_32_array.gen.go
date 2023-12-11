// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedInt32Array struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewPackedInt32Array() PackedInt32Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedInt32Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedInt32Array, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return PackedInt32Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedInt32ArrayFromPackedInt32Array(from PackedInt32Array, ) PackedInt32Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedInt32Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedInt32Array, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedInt32Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedInt32ArrayFromArray(from Array, ) PackedInt32Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedInt32Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedInt32Array, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedInt32Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *PackedInt32Array) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedInt32Array)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *PackedInt32Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedInt32Array
}

func (me *PackedInt32Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *PackedInt32Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *PackedInt32Array) Size() int {
  name := StringNameFromStr("size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Set(index int, value int, )  {
  name := StringNameFromStr("set")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 3638975848) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) PushBack(value int, ) bool {
  name := StringNameFromStr("push_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 694024632) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Append(value int, ) bool {
  name := StringNameFromStr("append")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 694024632) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) AppendArray(array PackedInt32Array, )  {
  name := StringNameFromStr("append_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 1087733270) // FIXME: should cache?

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) RemoveAt(index int, )  {
  name := StringNameFromStr("remove_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 2823966027) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Insert(at_index int, value int, ) int {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 1487112728) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&at_index), gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Fill(value int, )  {
  name := StringNameFromStr("fill")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 2823966027) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Resize(new_size int, ) int {
  name := StringNameFromStr("resize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 848867239) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&new_size), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Clear()  {
  name := StringNameFromStr("clear")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Has(value int, ) bool {
  name := StringNameFromStr("has")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 931488181) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Reverse()  {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Slice(begin int, end int, ) PackedInt32Array {
  name := StringNameFromStr("slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 1216021098) // FIXME: should cache?

  var ret PackedInt32Array
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin), gdc.ConstTypePtr(&end), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) ToByteArray() PackedByteArray {
  name := StringNameFromStr("to_byte_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Sort()  {
  name := StringNameFromStr("sort")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedInt32Array) Bsearch(value int, before bool, ) int {
  name := StringNameFromStr("bsearch")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 3380005890) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), gdc.ConstTypePtr(&before), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Duplicate() PackedInt32Array {
  name := StringNameFromStr("duplicate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 1997843129) // FIXME: should cache?

  var ret PackedInt32Array
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Find(value int, from int, ) int {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 2984303840) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Rfind(value int, from int, ) int {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 2984303840) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedInt32Array) Count(value int, ) int {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedInt32Array, name.AsCPtr(), 4103005248) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *PackedInt32Array) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt32Array) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt32Array) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt32Array) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt32Array) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt32Array) EqualPackedInt32Array(right PackedInt32Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt32Array) NotEqualPackedInt32Array(right PackedInt32Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedInt32Array) AddPackedInt32Array(right PackedInt32Array) PackedInt32Array {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret PackedInt32Array
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
