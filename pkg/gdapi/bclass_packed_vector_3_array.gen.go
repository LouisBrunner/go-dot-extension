// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedVector3Array struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewPackedVector3Array() PackedVector3Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedVector3Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedVector3Array, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return PackedVector3Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedVector3ArrayFromPackedVector3Array(from PackedVector3Array, ) PackedVector3Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedVector3Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedVector3Array, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedVector3Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewPackedVector3ArrayFromArray(from Array, ) PackedVector3Array {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizePackedVector3Array))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypePackedVector3Array, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return PackedVector3Array{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *PackedVector3Array) Destroy() {
  if me.ptr == nil {
    return
  }
  dtr := me.iface.VariantGetPtrDestructor(gdc.VariantTypePackedVector3Array)
	me.iface.CallPtrDestructor(dtr, gdc.TypePtr(me.ptr))
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *PackedVector3Array) Type() gdc.VariantType {
  return gdc.VariantTypePackedVector3Array
}

func (me *PackedVector3Array) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *PackedVector3Array) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *PackedVector3Array) Size() int {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) IsEmpty() bool {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Set(index int, value Vector3, )  {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) PushBack(value Vector3, ) bool {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Append(value Vector3, ) bool {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) AppendArray(array PackedVector3Array, )  {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) RemoveAt(index int, )  {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Insert(at_index int, value Vector3, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Fill(value Vector3, )  {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Resize(new_size int, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Clear()  {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Has(value Vector3, ) bool {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Reverse()  {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Slice(begin int, end int, ) PackedVector3Array {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) ToByteArray() PackedByteArray {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Sort()  {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Bsearch(value Vector3, before bool, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Duplicate() PackedVector3Array {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Find(value Vector3, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Rfind(value Vector3, from int, ) int {
  panic("TODO: implement")
}

func  (me *PackedVector3Array) Count(value Vector3, ) int {
  panic("TODO: implement")
}

// Operators

func (me *PackedVector3Array) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedVector3Array) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *PackedVector3Array) Not() bool {
  panic("TODO: implement")
}

func (me *PackedVector3Array) MultiplyTransform3D(right Transform3D) PackedVector3Array {
  panic("TODO: implement")
}

func (me *PackedVector3Array) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *PackedVector3Array) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *PackedVector3Array) EqualsPackedVector3Array(right PackedVector3Array) bool {
  panic("TODO: implement")
}

func (me *PackedVector3Array) NotEqualsPackedVector3Array(right PackedVector3Array) bool {
  panic("TODO: implement")
}

func (me *PackedVector3Array) AddPackedVector3Array(right PackedVector3Array) PackedVector3Array {
  panic("TODO: implement")
}

// TODO: members (bclass)
