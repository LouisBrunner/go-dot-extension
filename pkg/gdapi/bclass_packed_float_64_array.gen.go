// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedFloat64Array struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewPackedFloat64Array() PackedFloat64Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedFloat64Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat64Array, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return PackedFloat64Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedFloat64ArrayFromPackedFloat64Array(from PackedFloat64Array, ) PackedFloat64Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedFloat64Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat64Array, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedFloat64Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedFloat64ArrayFromArray(from Array, ) PackedFloat64Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedFloat64Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat64Array, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedFloat64Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *PackedFloat64Array) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedFloat64Array)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *PackedFloat64Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedFloat64Array
}

func (me *PackedFloat64Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *PackedFloat64Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *PackedFloat64Array) Size() int {
  name := StringNameFromStr("size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Set(index int, value float32, )  {
  name := StringNameFromStr("set")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1113000516) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) PushBack(value float32, ) bool {
  name := StringNameFromStr("push_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 4094791666) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Append(value float32, ) bool {
  name := StringNameFromStr("append")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 4094791666) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) AppendArray(array PackedFloat64Array, )  {
  name := StringNameFromStr("append_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 792078629) // FIXME: should cache?

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) RemoveAt(index int, )  {
  name := StringNameFromStr("remove_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 2823966027) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Insert(at_index int, value float32, ) int {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1379903876) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&at_index), gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Fill(value float32, )  {
  name := StringNameFromStr("fill")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 833936903) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Resize(new_size int, ) int {
  name := StringNameFromStr("resize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 848867239) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&new_size), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Clear()  {
  name := StringNameFromStr("clear")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Has(value float32, ) bool {
  name := StringNameFromStr("has")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1296369134) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Reverse()  {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Slice(begin int, end int, ) PackedFloat64Array {
  name := StringNameFromStr("slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 2192974324) // FIXME: should cache?

  var ret PackedFloat64Array
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin), gdc.ConstTypePtr(&end), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) ToByteArray() PackedByteArray {
  name := StringNameFromStr("to_byte_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 247621236) // FIXME: should cache?

  var ret PackedByteArray
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Sort()  {
  name := StringNameFromStr("sort")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *PackedFloat64Array) Bsearch(value float32, before bool, ) int {
  name := StringNameFromStr("bsearch")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1188816338) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), gdc.ConstTypePtr(&before), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Duplicate() PackedFloat64Array {
  name := StringNameFromStr("duplicate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 949266573) // FIXME: should cache?

  var ret PackedFloat64Array
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Find(value float32, from int, ) int {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1343150241) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Rfind(value float32, from int, ) int {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 1343150241) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *PackedFloat64Array) Count(value float32, ) int {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypePackedFloat64Array, name.AsCPtr(), 2859915090) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *PackedFloat64Array) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) EqualPackedFloat64Array(right PackedFloat64Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) NotEqualPackedFloat64Array(right PackedFloat64Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *PackedFloat64Array) AddPackedFloat64Array(right PackedFloat64Array) PackedFloat64Array {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret PackedFloat64Array
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
