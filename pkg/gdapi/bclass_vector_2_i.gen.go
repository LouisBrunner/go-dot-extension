// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector2i struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Vector2iZero = "Vector2i(0, 0)" // TODO: construct correctly
  Vector2iOne = "Vector2i(1, 1)" // TODO: construct correctly
  Vector2iLeft = "Vector2i(-1, 0)" // TODO: construct correctly
  Vector2iRight = "Vector2i(1, 0)" // TODO: construct correctly
  Vector2iUp = "Vector2i(0, -1)" // TODO: construct correctly
  Vector2iDown = "Vector2i(0, 1)" // TODO: construct correctly
)

// Enums

type Vector2iAxis int
const (
  Vector2iAxisAxisX Vector2iAxis = 0
  Vector2iAxisAxisY Vector2iAxis = 1
)

// Constructors

func NewVector2i() Vector2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector2I, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Vector2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector2iFromVector2I(from Vector2i, ) Vector2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector2I, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector2iFromVector2(from Vector2, ) Vector2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector2I, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector2iFromIntInt(x int, y int, ) Vector2i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector2i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector2I, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), }))
  return Vector2i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Vector2i) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Vector2i) Type() gdc.VariantType {
  return gdc.VariantTypeVector2I
}

func (me *Vector2i) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Vector2i) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Vector2i) Aspect() float32 {
  panic("TODO: implement")
}

func  (me *Vector2i) MaxAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector2i) MinAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector2i) Length() float32 {
  panic("TODO: implement")
}

func  (me *Vector2i) LengthSquared() int {
  panic("TODO: implement")
}

func  (me *Vector2i) Sign() Vector2i {
  panic("TODO: implement")
}

func  (me *Vector2i) Abs() Vector2i {
  panic("TODO: implement")
}

func  (me *Vector2i) Clamp(min Vector2i, max Vector2i, ) Vector2i {
  panic("TODO: implement")
}

func  (me *Vector2i) Snapped(step Vector2i, ) Vector2i {
  panic("TODO: implement")
}

// Operators

func (me *Vector2i) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector2i) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector2i) UnaryMinus() Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) UnaryPlus() Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) Not() bool {
  panic("TODO: implement")
}

func (me *Vector2i) MultiplyInt(right int) Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) DivideInt(right int) Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) ModuloInt(right int) Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) MultiplyFloat32(right float32) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2i) DivideFloat32(right float32) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2i) EqualsVector2I(right Vector2i) bool {
  panic("TODO: implement")
}

func (me *Vector2i) NotEqualsVector2I(right Vector2i) bool {
  panic("TODO: implement")
}

func (me *Vector2i) LessThanVector2I(right Vector2i) bool {
  panic("TODO: implement")
}

func (me *Vector2i) LessThanOrEqualsVector2I(right Vector2i) bool {
  panic("TODO: implement")
}

func (me *Vector2i) GreaterThanVector2I(right Vector2i) bool {
  panic("TODO: implement")
}

func (me *Vector2i) GreaterThanOrEqualsVector2I(right Vector2i) bool {
  panic("TODO: implement")
}

func (me *Vector2i) AddVector2I(right Vector2i) Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) SubtractVector2I(right Vector2i) Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) MultiplyVector2I(right Vector2i) Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) DivideVector2I(right Vector2i) Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) ModuloVector2I(right Vector2i) Vector2i {
  panic("TODO: implement")
}

func (me *Vector2i) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Vector2i) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
