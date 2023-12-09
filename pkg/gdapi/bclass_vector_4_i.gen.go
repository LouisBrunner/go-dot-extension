// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector4i struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Vector4iZero = "Vector4i(0, 0, 0, 0)" // TODO: construct correctly
  Vector4iOne = "Vector4i(1, 1, 1, 1)" // TODO: construct correctly
)

// Enums

type Vector4iAxis int
const (
  Vector4iAxisAxisX Vector4iAxis = 0
  Vector4iAxisAxisY Vector4iAxis = 1
  Vector4iAxisAxisZ Vector4iAxis = 2
  Vector4iAxisAxisW Vector4iAxis = 3
)

// Constructors

func NewVector4i() Vector4i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Vector4i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4iFromVector4I(from Vector4i, ) Vector4i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector4i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4iFromVector4(from Vector4, ) Vector4i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector4i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4iFromIntIntIntInt(x int, y int, z int, w int, ) Vector4i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4I, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
  return Vector4i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Vector4i) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Vector4i) Type() gdc.VariantType {
  return gdc.VariantTypeVector4I
}

func (me *Vector4i) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Vector4i) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Vector4i) MinAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector4i) MaxAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector4i) Length() float32 {
  panic("TODO: implement")
}

func  (me *Vector4i) LengthSquared() int {
  panic("TODO: implement")
}

func  (me *Vector4i) Sign() Vector4i {
  panic("TODO: implement")
}

func  (me *Vector4i) Abs() Vector4i {
  panic("TODO: implement")
}

func  (me *Vector4i) Clamp(min Vector4i, max Vector4i, ) Vector4i {
  panic("TODO: implement")
}

func  (me *Vector4i) Snapped(step Vector4i, ) Vector4i {
  panic("TODO: implement")
}

// Operators

func (me *Vector4i) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector4i) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector4i) UnaryMinus() Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) UnaryPlus() Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) Not() bool {
  panic("TODO: implement")
}

func (me *Vector4i) MultiplyInt(right int) Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) DivideInt(right int) Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) ModuloInt(right int) Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) MultiplyFloat32(right float32) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4i) DivideFloat32(right float32) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4i) EqualsVector4I(right Vector4i) bool {
  panic("TODO: implement")
}

func (me *Vector4i) NotEqualsVector4I(right Vector4i) bool {
  panic("TODO: implement")
}

func (me *Vector4i) LessThanVector4I(right Vector4i) bool {
  panic("TODO: implement")
}

func (me *Vector4i) LessThanOrEqualsVector4I(right Vector4i) bool {
  panic("TODO: implement")
}

func (me *Vector4i) GreaterThanVector4I(right Vector4i) bool {
  panic("TODO: implement")
}

func (me *Vector4i) GreaterThanOrEqualsVector4I(right Vector4i) bool {
  panic("TODO: implement")
}

func (me *Vector4i) AddVector4I(right Vector4i) Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) SubtractVector4I(right Vector4i) Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) MultiplyVector4I(right Vector4i) Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) DivideVector4I(right Vector4i) Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) ModuloVector4I(right Vector4i) Vector4i {
  panic("TODO: implement")
}

func (me *Vector4i) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Vector4i) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
