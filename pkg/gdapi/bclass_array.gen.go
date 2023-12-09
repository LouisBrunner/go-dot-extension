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

func  (me *Array) Size() int {
  panic("TODO: implement")
}

func  (me *Array) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *Array) Clear()  {
  panic("TODO: implement")
}

func  (me *Array) Hash() int {
  panic("TODO: implement")
}

func  (me *Array) Assign(array Array, )  {
  panic("TODO: implement")
}

func  (me *Array) PushBack(value Variant, )  {
  panic("TODO: implement")
}

func  (me *Array) PushFront(value Variant, )  {
  panic("TODO: implement")
}

func  (me *Array) Append(value Variant, )  {
  panic("TODO: implement")
}

func  (me *Array) AppendArray(array Array, )  {
  panic("TODO: implement")
}

func  (me *Array) Resize(size int, ) int {
  panic("TODO: implement")
}

func  (me *Array) Insert(position int, value Variant, ) int {
  panic("TODO: implement")
}

func  (me *Array) RemoveAt(position int, )  {
  panic("TODO: implement")
}

func  (me *Array) Fill(value Variant, )  {
  panic("TODO: implement")
}

func  (me *Array) Erase(value Variant, )  {
  panic("TODO: implement")
}

func  (me *Array) Front() Variant {
  panic("TODO: implement")
}

func  (me *Array) Back() Variant {
  panic("TODO: implement")
}

func  (me *Array) PickRandom() Variant {
  panic("TODO: implement")
}

func  (me *Array) Find(what Variant, from int, ) int {
  panic("TODO: implement")
}

func  (me *Array) Rfind(what Variant, from int, ) int {
  panic("TODO: implement")
}

func  (me *Array) Count(value Variant, ) int {
  panic("TODO: implement")
}

func  (me *Array) Has(value Variant, ) bool {
  panic("TODO: implement")
}

func  (me *Array) PopBack() Variant {
  panic("TODO: implement")
}

func  (me *Array) PopFront() Variant {
  panic("TODO: implement")
}

func  (me *Array) PopAt(position int, ) Variant {
  panic("TODO: implement")
}

func  (me *Array) Sort()  {
  panic("TODO: implement")
}

func  (me *Array) SortCustom(func_ Callable, )  {
  panic("TODO: implement")
}

func  (me *Array) Shuffle()  {
  panic("TODO: implement")
}

func  (me *Array) Bsearch(value Variant, before bool, ) int {
  panic("TODO: implement")
}

func  (me *Array) BsearchCustom(value Variant, func_ Callable, before bool, ) int {
  panic("TODO: implement")
}

func  (me *Array) Reverse()  {
  panic("TODO: implement")
}

func  (me *Array) Duplicate(deep bool, ) Array {
  panic("TODO: implement")
}

func  (me *Array) Slice(begin int, end int, step int, deep bool, ) Array {
  panic("TODO: implement")
}

func  (me *Array) Filter(method Callable, ) Array {
  panic("TODO: implement")
}

func  (me *Array) Map(method Callable, ) Array {
  panic("TODO: implement")
}

func  (me *Array) Reduce(method Callable, accum Variant, ) Variant {
  panic("TODO: implement")
}

func  (me *Array) Any(method Callable, ) bool {
  panic("TODO: implement")
}

func  (me *Array) All(method Callable, ) bool {
  panic("TODO: implement")
}

func  (me *Array) Max() Variant {
  panic("TODO: implement")
}

func  (me *Array) Min() Variant {
  panic("TODO: implement")
}

func  (me *Array) IsTyped() bool {
  panic("TODO: implement")
}

func  (me *Array) IsSameTyped(array Array, ) bool {
  panic("TODO: implement")
}

func  (me *Array) GetTypedBuiltin() int {
  panic("TODO: implement")
}

func  (me *Array) GetTypedClassName() StringName {
  panic("TODO: implement")
}

func  (me *Array) GetTypedScript() Variant {
  panic("TODO: implement")
}

func  (me *Array) MakeReadOnly()  {
  panic("TODO: implement")
}

func  (me *Array) IsReadOnly() bool {
  panic("TODO: implement")
}

// Operators

func (me *Array) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Array) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Array) Not() bool {
  panic("TODO: implement")
}

func (me *Array) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Array) EqualsArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Array) NotEqualsArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Array) LessThanArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Array) LessThanOrEqualsArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Array) GreaterThanArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Array) GreaterThanOrEqualsArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Array) AddArray(right Array) Array {
  panic("TODO: implement")
}

func (me *Array) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
