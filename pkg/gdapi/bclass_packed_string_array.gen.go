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

func  (me *PackedStringArray) Size() int {
  panic("TODO: implement")
}

func  (me *PackedStringArray) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Set(index int, value String, )  {
  panic("TODO: implement")
}

func  (me *PackedStringArray) PushBack(value String, ) bool {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Append(value String, ) bool {
  panic("TODO: implement")
}

func  (me *PackedStringArray) AppendArray(array PackedStringArray, )  {
  panic("TODO: implement")
}

func  (me *PackedStringArray) RemoveAt(index int, )  {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Insert(at_index int, value String, ) int {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Fill(value String, )  {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Resize(new_size int, ) int {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Clear()  {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Has(value String, ) bool {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Reverse()  {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Slice(begin int, end int, ) PackedStringArray {
  panic("TODO: implement")
}

func  (me *PackedStringArray) ToByteArray() PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Sort()  {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Bsearch(value String, before bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Duplicate() PackedStringArray {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Find(value String, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Rfind(value String, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedStringArray) Count(value String, ) int {
  panic("TODO: implement")
}

// Operators

func (me *PackedStringArray) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedStringArray) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedStringArray) Not() bool {
  panic("TODO: implement")
}

func (me *PackedStringArray) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *PackedStringArray) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *PackedStringArray) EqualsPackedStringArray(right PackedStringArray) bool {
  panic("TODO: implement")
}

func (me *PackedStringArray) NotEqualsPackedStringArray(right PackedStringArray) bool {
  panic("TODO: implement")
}

func (me *PackedStringArray) AddPackedStringArray(right PackedStringArray) PackedStringArray {
  panic("TODO: implement")
}

// TODO: members (bclass)
