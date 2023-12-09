// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedInt64Array struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewPackedInt64Array() PackedInt64Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedInt64Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedInt64Array, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return PackedInt64Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedInt64ArrayFromPackedInt64Array(from PackedInt64Array, ) PackedInt64Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedInt64Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedInt64Array, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedInt64Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedInt64ArrayFromArray(from Array, ) PackedInt64Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedInt64Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedInt64Array, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedInt64Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *PackedInt64Array) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedInt64Array)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *PackedInt64Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedInt64Array
}

func (me *PackedInt64Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *PackedInt64Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *PackedInt64Array) Size() int {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Set(index int, value int, )  {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) PushBack(value int, ) bool {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Append(value int, ) bool {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) AppendArray(array PackedInt64Array, )  {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) RemoveAt(index int, )  {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Insert(at_index int, value int, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Fill(value int, )  {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Resize(new_size int, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Clear()  {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Has(value int, ) bool {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Reverse()  {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Slice(begin int, end int, ) PackedInt64Array {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) ToByteArray() PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Sort()  {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Bsearch(value int, before bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Duplicate() PackedInt64Array {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Find(value int, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Rfind(value int, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedInt64Array) Count(value int, ) int {
  panic("TODO: implement")
}

// Operators

func (me *PackedInt64Array) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedInt64Array) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedInt64Array) Not() bool {
  panic("TODO: implement")
}

func (me *PackedInt64Array) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *PackedInt64Array) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *PackedInt64Array) EqualsPackedInt64Array(right PackedInt64Array) bool {
  panic("TODO: implement")
}

func (me *PackedInt64Array) NotEqualsPackedInt64Array(right PackedInt64Array) bool {
  panic("TODO: implement")
}

func (me *PackedInt64Array) AddPackedInt64Array(right PackedInt64Array) PackedInt64Array {
  panic("TODO: implement")
}

// TODO: members (bclass)
