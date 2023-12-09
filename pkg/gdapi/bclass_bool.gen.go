// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Bool struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Enums

// Constructors

func NewBool() Bool {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBool))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBool, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Bool{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBoolFromBool(from bool, ) Bool {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBool))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBool, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return Bool{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBoolFromInt(from int, ) Bool {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBool))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBool, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return Bool{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewBoolFromFloat32(from float32, ) Bool {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeBool))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeBool, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&from), }))
  return Bool{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Bool) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Bool) Type() gdc.VariantType {
  return gdc.VariantTypeBool
}

func (me *Bool) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Bool) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

// Operators

func (me *Bool) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Bool) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Bool) AndVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Bool) OrVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Bool) XorVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Bool) Not() bool {
  panic("TODO: implement")
}

func (me *Bool) EqualsBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Bool) NotEqualsBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Bool) LessThanBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Bool) GreaterThanBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Bool) AndBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Bool) OrBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Bool) XorBool(right bool) bool {
  panic("TODO: implement")
}

func (me *Bool) AndInt(right int) bool {
  panic("TODO: implement")
}

func (me *Bool) OrInt(right int) bool {
  panic("TODO: implement")
}

func (me *Bool) XorInt(right int) bool {
  panic("TODO: implement")
}

func (me *Bool) AndFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Bool) OrFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Bool) XorFloat32(right float32) bool {
  panic("TODO: implement")
}

func (me *Bool) AndObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Bool) OrObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Bool) XorObject(right Object) bool {
  panic("TODO: implement")
}

func (me *Bool) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Bool) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
