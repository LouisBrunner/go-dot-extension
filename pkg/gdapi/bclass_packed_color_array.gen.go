// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedColorArray struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewPackedColorArray() PackedColorArray {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedColorArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedColorArray, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return PackedColorArray{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedColorArrayFromPackedColorArray(from PackedColorArray, ) PackedColorArray {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedColorArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedColorArray, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedColorArray{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedColorArrayFromArray(from Array, ) PackedColorArray {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedColorArray))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedColorArray, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedColorArray{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *PackedColorArray) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedColorArray)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *PackedColorArray) Type() gdc.VariantType {
  return gdc.VariantTypePackedColorArray
}

func (me *PackedColorArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *PackedColorArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *PackedColorArray) Size() int {
  panic("TODO: implement")
}

func  (me *PackedColorArray) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Set(index int, value Color, )  {
  panic("TODO: implement")
}

func  (me *PackedColorArray) PushBack(value Color, ) bool {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Append(value Color, ) bool {
  panic("TODO: implement")
}

func  (me *PackedColorArray) AppendArray(array PackedColorArray, )  {
  panic("TODO: implement")
}

func  (me *PackedColorArray) RemoveAt(index int, )  {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Insert(at_index int, value Color, ) int {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Fill(value Color, )  {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Resize(new_size int, ) int {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Clear()  {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Has(value Color, ) bool {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Reverse()  {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Slice(begin int, end int, ) PackedColorArray {
  panic("TODO: implement")
}

func  (me *PackedColorArray) ToByteArray() PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Sort()  {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Bsearch(value Color, before bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Duplicate() PackedColorArray {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Find(value Color, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Rfind(value Color, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedColorArray) Count(value Color, ) int {
  panic("TODO: implement")
}

// Operators

func (me *PackedColorArray) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedColorArray) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedColorArray) Not() bool {
  panic("TODO: implement")
}

func (me *PackedColorArray) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *PackedColorArray) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *PackedColorArray) EqualsPackedColorArray(right PackedColorArray) bool {
  panic("TODO: implement")
}

func (me *PackedColorArray) NotEqualsPackedColorArray(right PackedColorArray) bool {
  panic("TODO: implement")
}

func (me *PackedColorArray) AddPackedColorArray(right PackedColorArray) PackedColorArray {
  panic("TODO: implement")
}

// TODO: members (bclass)
