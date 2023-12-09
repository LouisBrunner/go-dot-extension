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

func  (me *PackedFloat64Array) Size() int {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Set(index int, value float32, )  {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) PushBack(value float32, ) bool {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Append(value float32, ) bool {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) AppendArray(array PackedFloat64Array, )  {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) RemoveAt(index int, )  {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Insert(at_index int, value float32, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Fill(value float32, )  {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Resize(new_size int, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Clear()  {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Has(value float32, ) bool {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Reverse()  {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Slice(begin int, end int, ) PackedFloat64Array {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) ToByteArray() PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Sort()  {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Bsearch(value float32, before bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Duplicate() PackedFloat64Array {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Find(value float32, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Rfind(value float32, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedFloat64Array) Count(value float32, ) int {
  panic("TODO: implement")
}

// Operators

func (me *PackedFloat64Array) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedFloat64Array) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedFloat64Array) Not() bool {
  panic("TODO: implement")
}

func (me *PackedFloat64Array) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *PackedFloat64Array) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *PackedFloat64Array) EqualsPackedFloat64Array(right PackedFloat64Array) bool {
  panic("TODO: implement")
}

func (me *PackedFloat64Array) NotEqualsPackedFloat64Array(right PackedFloat64Array) bool {
  panic("TODO: implement")
}

func (me *PackedFloat64Array) AddPackedFloat64Array(right PackedFloat64Array) PackedFloat64Array {
  panic("TODO: implement")
}

// TODO: members (bclass)
