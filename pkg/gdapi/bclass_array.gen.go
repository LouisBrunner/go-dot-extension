// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Array struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewArray() Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromArray(from Array, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromArrayIntStringNameVariant(base Array, type_ int, class_name StringName, script Variant, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{base.AsCTypePtr(), gdc.ConstTypePtr(&type_), class_name.AsCTypePtr(), script.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromPackedByteArray(from PackedByteArray, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromPackedInt32Array(from PackedInt32Array, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 4) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromPackedInt64Array(from PackedInt64Array, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 5) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromPackedFloat32Array(from PackedFloat32Array, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 6) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromPackedFloat64Array(from PackedFloat64Array, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 7) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromPackedStringArray(from PackedStringArray, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 8) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromPackedVector2Array(from PackedVector2Array, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 9) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromPackedVector3Array(from PackedVector3Array, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 10) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewArrayFromPackedColorArray(from PackedColorArray, ) Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeArray, 11) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Array) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypeArray)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Array) Type() gdc.VariantType {
  return gdc.VariantTypeArray
}

func (me *Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func (me *Array) Size() int {
  name := StringNameFromStr("size")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) IsEmpty() bool {
  name := StringNameFromStr("is_empty")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Clear()  {
  name := StringNameFromStr("clear")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Hash() int {
  name := StringNameFromStr("hash")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Assign(array Array, )  {
  name := StringNameFromStr("assign")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 2307260970) // FIXME: should cache?

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) PushBack(value Variant, )  {
  name := StringNameFromStr("push_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3316032543) // FIXME: should cache?

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) PushFront(value Variant, )  {
  name := StringNameFromStr("push_front")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3316032543) // FIXME: should cache?

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Append(value Variant, )  {
  name := StringNameFromStr("append")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3316032543) // FIXME: should cache?

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) AppendArray(array Array, )  {
  name := StringNameFromStr("append_array")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 2307260970) // FIXME: should cache?

  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Resize(size int, ) int {
  name := StringNameFromStr("resize")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 848867239) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Insert(position int, value Variant, ) int {
  name := StringNameFromStr("insert")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3176316662) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) RemoveAt(position int, )  {
  name := StringNameFromStr("remove_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 2823966027) // FIXME: should cache?

  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Fill(value Variant, )  {
  name := StringNameFromStr("fill")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3316032543) // FIXME: should cache?

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Erase(value Variant, )  {
  name := StringNameFromStr("erase")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3316032543) // FIXME: should cache?

  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Front() Variant {
  name := StringNameFromStr("front")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1460142086) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Back() Variant {
  name := StringNameFromStr("back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1460142086) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) PickRandom() Variant {
  name := StringNameFromStr("pick_random")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1460142086) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Find(what Variant, from int, ) int {
  name := StringNameFromStr("find")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 2336346817) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Rfind(what Variant, from int, ) int {
  name := StringNameFromStr("rfind")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 2336346817) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{what.AsCTypePtr(), gdc.ConstTypePtr(&from), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Count(value Variant, ) int {
  name := StringNameFromStr("count")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1481661226) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Has(value Variant, ) bool {
  name := StringNameFromStr("has")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3680194679) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) PopBack() Variant {
  name := StringNameFromStr("pop_back")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1321915136) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) PopFront() Variant {
  name := StringNameFromStr("pop_front")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1321915136) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) PopAt(position int, ) Variant {
  name := StringNameFromStr("pop_at")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3518259424) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&position), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Sort()  {
  name := StringNameFromStr("sort")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) SortCustom(func_ Callable, )  {
  name := StringNameFromStr("sort_custom")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3470848906) // FIXME: should cache?

  args := []gdc.ConstTypePtr{func_.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Shuffle()  {
  name := StringNameFromStr("shuffle")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Bsearch(value Variant, before bool, ) int {
  name := StringNameFromStr("bsearch")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3372222236) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), gdc.ConstTypePtr(&before), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) BsearchCustom(value Variant, func_ Callable, before bool, ) int {
  name := StringNameFromStr("bsearch_custom")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 161317131) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{value.AsCTypePtr(), func_.AsCTypePtr(), gdc.ConstTypePtr(&before), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Reverse()  {
  name := StringNameFromStr("reverse")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) Duplicate(deep bool, ) Array {
  name := StringNameFromStr("duplicate")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 636440122) // FIXME: should cache?

  var ret Array
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&deep), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Slice(begin int, end int, step int, deep bool, ) Array {
  name := StringNameFromStr("slice")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1393718243) // FIXME: should cache?

  var ret Array
  args := []gdc.ConstTypePtr{gdc.ConstTypePtr(&begin), gdc.ConstTypePtr(&end), gdc.ConstTypePtr(&step), gdc.ConstTypePtr(&deep), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Filter(method Callable, ) Array {
  name := StringNameFromStr("filter")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 4075186556) // FIXME: should cache?

  var ret Array
  args := []gdc.ConstTypePtr{method.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Map(method Callable, ) Array {
  name := StringNameFromStr("map")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 4075186556) // FIXME: should cache?

  var ret Array
  args := []gdc.ConstTypePtr{method.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Reduce(method Callable, accum Variant, ) Variant {
  name := StringNameFromStr("reduce")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 4272450342) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{method.AsCTypePtr(), accum.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Any(method Callable, ) bool {
  name := StringNameFromStr("any")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 4129521963) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{method.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) All(method Callable, ) bool {
  name := StringNameFromStr("all")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 4129521963) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{method.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Max() Variant {
  name := StringNameFromStr("max")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1460142086) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) Min() Variant {
  name := StringNameFromStr("min")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1460142086) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) IsTyped() bool {
  name := StringNameFromStr("is_typed")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) IsSameTyped(array Array, ) bool {
  name := StringNameFromStr("is_same_typed")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 2988181878) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{array.AsCTypePtr(), }

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) GetTypedBuiltin() int {
  name := StringNameFromStr("get_typed_builtin")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3173160232) // FIXME: should cache?

  var ret int
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) GetTypedClassName() StringName {
  name := StringNameFromStr("get_typed_class_name")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1825232092) // FIXME: should cache?

  var ret StringName
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) GetTypedScript() Variant {
  name := StringNameFromStr("get_typed_script")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 1460142086) // FIXME: should cache?

  var ret Variant
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

func (me *Array) MakeReadOnly()  {
  name := StringNameFromStr("make_read_only")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3218959716) // FIXME: should cache?

  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), nil, len(args))
}

func (me *Array) IsReadOnly() bool {
  name := StringNameFromStr("is_read_only")
  defer name.Destroy()
  methodPtr := giface.VariantGetPtrBuiltinMethod(gdc.VariantTypeArray, name.AsCPtr(), 3918633141) // FIXME: should cache?

  var ret bool
  args := []gdc.ConstTypePtr{}

  giface.CallPtrBuiltInMethod(methodPtr, me.AsTypePtr(), unsafe.SliceData(args), gdc.TypePtr(&ret), len(args))
  return ret
}

// Operators

func (me *Array) EqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) NotEqualVariant(right Variant) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) Not() bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNot, me.Type(), gdc.VariantTypeNil) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), nil, gdc.TypePtr(&ret))
  return ret
}

func (me *Array) InDictionary(right Dictionary) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) EqualArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) NotEqualArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpNotEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) LessArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLess, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) LessEqualArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpLessEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) GreaterArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreater, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) GreaterEqualArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpGreaterEqual, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) AddArray(right Array) Array {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpAdd, me.Type(), right.Type()) // FIXME: cache
  var ret Array
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

func (me *Array) InArray(right Array) bool {
  op := me.iface.VariantGetPtrOperatorEvaluator(gdc.VariantOpIn, me.Type(), right.Type()) // FIXME: cache
  var ret bool
  me.iface.CallPtrOperatorEvaluator(op, me.AsCTypePtr(), right.AsCTypePtr(), gdc.TypePtr(&ret))
  return ret
}

// Members
