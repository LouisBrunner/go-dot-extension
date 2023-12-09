// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector3i struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Vector3iZero = "Vector3i(0, 0, 0)" // TODO: construct correctly
  Vector3iOne = "Vector3i(1, 1, 1)" // TODO: construct correctly
  Vector3iLeft = "Vector3i(-1, 0, 0)" // TODO: construct correctly
  Vector3iRight = "Vector3i(1, 0, 0)" // TODO: construct correctly
  Vector3iUp = "Vector3i(0, 1, 0)" // TODO: construct correctly
  Vector3iDown = "Vector3i(0, -1, 0)" // TODO: construct correctly
  Vector3iForward = "Vector3i(0, 0, -1)" // TODO: construct correctly
  Vector3iBack = "Vector3i(0, 0, 1)" // TODO: construct correctly
)

// Enums

type Vector3iAxis int
const (
  Vector3iAxisAxisX Vector3iAxis = 0
  Vector3iAxisAxisY Vector3iAxis = 1
  Vector3iAxisAxisZ Vector3iAxis = 2
)

// Constructors

func NewVector3i() Vector3i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector3i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Vector3i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector3iFromVector3I(from Vector3i, ) Vector3i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector3i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector3i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector3iFromVector3(from Vector3, ) Vector3i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector3i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector3i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector3iFromIntIntInt(x int, y int, z int, ) Vector3i {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector3i))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector3I, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), }))
  return Vector3i{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Vector3i) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Vector3i) Type() gdc.VariantType {
  return gdc.VariantTypeVector3I
}

func (me *Vector3i) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Vector3i) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Vector3i) MinAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector3i) MaxAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector3i) Length() float32 {
  panic("TODO: implement")
}

func  (me *Vector3i) LengthSquared() int {
  panic("TODO: implement")
}

func  (me *Vector3i) Sign() Vector3i {
  panic("TODO: implement")
}

func  (me *Vector3i) Abs() Vector3i {
  panic("TODO: implement")
}

func  (me *Vector3i) Clamp(min Vector3i, max Vector3i, ) Vector3i {
  panic("TODO: implement")
}

func  (me *Vector3i) Snapped(step Vector3i, ) Vector3i {
  panic("TODO: implement")
}

// Operators

func (me *Vector3i) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector3i) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector3i) UnaryMinus() Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) UnaryPlus() Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) Not() bool {
  panic("TODO: implement")
}

func (me *Vector3i) MultiplyInt(right int) Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) DivideInt(right int) Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) ModuloInt(right int) Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) MultiplyFloat32(right float32) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3i) DivideFloat32(right float32) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3i) EqualsVector3I(right Vector3i) bool {
  panic("TODO: implement")
}

func (me *Vector3i) NotEqualsVector3I(right Vector3i) bool {
  panic("TODO: implement")
}

func (me *Vector3i) LessThanVector3I(right Vector3i) bool {
  panic("TODO: implement")
}

func (me *Vector3i) LessThanOrEqualsVector3I(right Vector3i) bool {
  panic("TODO: implement")
}

func (me *Vector3i) GreaterThanVector3I(right Vector3i) bool {
  panic("TODO: implement")
}

func (me *Vector3i) GreaterThanOrEqualsVector3I(right Vector3i) bool {
  panic("TODO: implement")
}

func (me *Vector3i) AddVector3I(right Vector3i) Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) SubtractVector3I(right Vector3i) Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) MultiplyVector3I(right Vector3i) Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) DivideVector3I(right Vector3i) Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) ModuloVector3I(right Vector3i) Vector3i {
  panic("TODO: implement")
}

func (me *Vector3i) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Vector3i) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
