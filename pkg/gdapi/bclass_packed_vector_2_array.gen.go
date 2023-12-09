// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedVector2Array struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewPackedVector2Array() PackedVector2Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedVector2Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedVector2Array, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return PackedVector2Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedVector2ArrayFromPackedVector2Array(from PackedVector2Array, ) PackedVector2Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedVector2Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedVector2Array, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedVector2Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedVector2ArrayFromArray(from Array, ) PackedVector2Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedVector2Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedVector2Array, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedVector2Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *PackedVector2Array) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedVector2Array)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *PackedVector2Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedVector2Array
}

func (me *PackedVector2Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *PackedVector2Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *PackedVector2Array) Size() int {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Set(index int, value Vector2, )  {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) PushBack(value Vector2, ) bool {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Append(value Vector2, ) bool {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) AppendArray(array PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) RemoveAt(index int, )  {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Insert(at_index int, value Vector2, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Fill(value Vector2, )  {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Resize(new_size int, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Clear()  {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Has(value Vector2, ) bool {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Reverse()  {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Slice(begin int, end int, ) PackedVector2Array {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) ToByteArray() PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Sort()  {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Bsearch(value Vector2, before bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Duplicate() PackedVector2Array {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Find(value Vector2, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Rfind(value Vector2, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector2Array) Count(value Vector2, ) int {
  panic("TODO: implement")
}

// Operators

func (me *PackedVector2Array) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedVector2Array) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedVector2Array) Not() bool {
  panic("TODO: implement")
}

func (me *PackedVector2Array) MultiplyTransform2D(right Transform2D) PackedVector2Array {
  panic("TODO: implement")
}

func (me *PackedVector2Array) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *PackedVector2Array) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *PackedVector2Array) EqualsPackedVector2Array(right PackedVector2Array) bool {
  panic("TODO: implement")
}

func (me *PackedVector2Array) NotEqualsPackedVector2Array(right PackedVector2Array) bool {
  panic("TODO: implement")
}

func (me *PackedVector2Array) AddPackedVector2Array(right PackedVector2Array) PackedVector2Array {
  panic("TODO: implement")
}

// TODO: members (bclass)
