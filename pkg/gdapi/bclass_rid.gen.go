// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RID struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewRID() RID {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRID))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRID, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return RID{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewRIDFromRID(from RID, ) RID {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeRID))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeRID, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return RID{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *RID) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *RID) Type() gdc.VariantType {
  return gdc.VariantTypeRID
}

func (me *RID) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *RID) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *RID) IsValid() bool {
  panic("TODO: implement")
}

func  (me *RID) GetId() int {
  panic("TODO: implement")
}

// Operators

func (me *RID) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *RID) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *RID) Not() bool {
  panic("TODO: implement")
}

func (me *RID) EqualsRID(right RID) bool {
  panic("TODO: implement")
}

func (me *RID) NotEqualsRID(right RID) bool {
  panic("TODO: implement")
}

func (me *RID) LessThanRID(right RID) bool {
  panic("TODO: implement")
}

func (me *RID) LessThanOrEqualsRID(right RID) bool {
  panic("TODO: implement")
}

func (me *RID) GreaterThanRID(right RID) bool {
  panic("TODO: implement")
}

func (me *RID) GreaterThanOrEqualsRID(right RID) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
