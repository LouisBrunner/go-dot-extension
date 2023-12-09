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

func  (me *PackedInt32Array) Size() int {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Set(index int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) PushBack(value int, ) bool {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Append(value int, ) bool {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) AppendArray(array PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) RemoveAt(index int, )  {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Insert(at_index int, value int, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Fill(value int, )  {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Resize(new_size int, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Clear()  {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Has(value int, ) bool {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Reverse()  {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Slice(begin int, end int, ) PackedInt32Array {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) ToByteArray() PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Sort()  {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Bsearch(value int, before bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Duplicate() PackedInt32Array {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Find(value int, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Rfind(value int, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt32Array) Count(value int, ) int {
  panic("TODO: implement")
}

// Operators

func (me *PackedInt32Array) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedInt32Array) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedInt32Array) Not() bool {
  panic("TODO: implement")
}

func (me *PackedInt32Array) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *PackedInt32Array) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *PackedInt32Array) EqualsPackedInt32Array(right PackedInt32Array) bool {
  panic("TODO: implement")
}

func (me *PackedInt32Array) NotEqualsPackedInt32Array(right PackedInt32Array) bool {
  panic("TODO: implement")
}

func (me *PackedInt32Array) AddPackedInt32Array(right PackedInt32Array) PackedInt32Array {
  panic("TODO: implement")
}

// TODO: members (bclass)
