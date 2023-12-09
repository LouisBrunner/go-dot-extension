// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedFloat32Array struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewPackedFloat32Array() PackedFloat32Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedFloat32Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat32Array, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return PackedFloat32Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedFloat32ArrayFromPackedFloat32Array(from PackedFloat32Array, ) PackedFloat32Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedFloat32Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat32Array, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedFloat32Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedFloat32ArrayFromArray(from Array, ) PackedFloat32Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedFloat32Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedFloat32Array, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedFloat32Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *PackedFloat32Array) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedFloat32Array)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *PackedFloat32Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedFloat32Array
}

func (me *PackedFloat32Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *PackedFloat32Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *PackedFloat32Array) Size() int {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Set(index int, value float32, )  {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) PushBack(value float32, ) bool {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Append(value float32, ) bool {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) AppendArray(array PackedFloat32Array, )  {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) RemoveAt(index int, )  {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Insert(at_index int, value float32, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Fill(value float32, )  {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Resize(new_size int, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Clear()  {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Has(value float32, ) bool {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Reverse()  {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Slice(begin int, end int, ) PackedFloat32Array {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) ToByteArray() PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Sort()  {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Bsearch(value float32, before bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Duplicate() PackedFloat32Array {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Find(value float32, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Rfind(value float32, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat32Array) Count(value float32, ) int {
  panic("TODO: implement")
}

// Operators

func (me *PackedFloat32Array) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedFloat32Array) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedFloat32Array) Not() bool {
  panic("TODO: implement")
}

func (me *PackedFloat32Array) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *PackedFloat32Array) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *PackedFloat32Array) EqualsPackedFloat32Array(right PackedFloat32Array) bool {
  panic("TODO: implement")
}

func (me *PackedFloat32Array) NotEqualsPackedFloat32Array(right PackedFloat32Array) bool {
  panic("TODO: implement")
}

func (me *PackedFloat32Array) AddPackedFloat32Array(right PackedFloat32Array) PackedFloat32Array {
  panic("TODO: implement")
}

// TODO: members (bclass)
