// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector4 struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Vector4Zero = "Vector4(0, 0, 0, 0)" // TODO: construct correctly
  Vector4One = "Vector4(1, 1, 1, 1)" // TODO: construct correctly
  Vector4Inf = "Vector4(inf, inf, inf, inf)" // TODO: construct correctly
)

// Enums

type Vector4Axis int
const (
  Vector4AxisAxisX Vector4Axis = 0
  Vector4AxisAxisY Vector4Axis = 1
  Vector4AxisAxisZ Vector4Axis = 2
  Vector4AxisAxisW Vector4Axis = 3
)

// Constructors

func NewVector4() Vector4 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Vector4{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4FromVector4(from Vector4, ) Vector4 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector4{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4FromVector4I(from Vector4i, ) Vector4 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector4{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector4FromFloat32Float32Float32Float32(x float32, y float32, z float32, w float32, ) Vector4 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector4))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector4, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), gdc.ConstTypePtr(&w), }))
  return Vector4{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Vector4) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Vector4) Type() gdc.VariantType {
  return gdc.VariantTypeVector4
}

func (me *Vector4) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Vector4) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Vector4) MinAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector4) MaxAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector4) Length() float32 {
  panic("TODO: implement")
}

func  (me *Vector4) LengthSquared() float32 {
  panic("TODO: implement")
}

func  (me *Vector4) Abs() Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Sign() Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Floor() Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Ceil() Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Round() Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Lerp(to Vector4, weight float32, ) Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) CubicInterpolate(b Vector4, pre_a Vector4, post_b Vector4, weight float32, ) Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) CubicInterpolateInTime(b Vector4, pre_a Vector4, post_b Vector4, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Posmod(mod float32, ) Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Posmodv(modv Vector4, ) Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Snapped(step Vector4, ) Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Clamp(min Vector4, max Vector4, ) Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) Normalized() Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) IsNormalized() bool {
  panic("TODO: implement")
}

func  (me *Vector4) DirectionTo(to Vector4, ) Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) DistanceTo(to Vector4, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector4) DistanceSquaredTo(to Vector4, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector4) Dot(with Vector4, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector4) Inverse() Vector4 {
  panic("TODO: implement")
}

func  (me *Vector4) IsEqualApprox(to Vector4, ) bool {
  panic("TODO: implement")
}

func  (me *Vector4) IsZeroApprox() bool {
  panic("TODO: implement")
}

func  (me *Vector4) IsFinite() bool {
  panic("TODO: implement")
}

// Operators

func (me *Vector4) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector4) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector4) UnaryMinus() Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) UnaryPlus() Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) Not() bool {
  panic("TODO: implement")
}

func (me *Vector4) MultiplyInt(right int) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) DivideInt(right int) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) MultiplyFloat32(right float32) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) DivideFloat32(right float32) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) EqualsVector4(right Vector4) bool {
  panic("TODO: implement")
}

func (me *Vector4) NotEqualsVector4(right Vector4) bool {
  panic("TODO: implement")
}

func (me *Vector4) LessThanVector4(right Vector4) bool {
  panic("TODO: implement")
}

func (me *Vector4) LessThanOrEqualsVector4(right Vector4) bool {
  panic("TODO: implement")
}

func (me *Vector4) GreaterThanVector4(right Vector4) bool {
  panic("TODO: implement")
}

func (me *Vector4) GreaterThanOrEqualsVector4(right Vector4) bool {
  panic("TODO: implement")
}

func (me *Vector4) AddVector4(right Vector4) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) SubtractVector4(right Vector4) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) MultiplyVector4(right Vector4) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) DivideVector4(right Vector4) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) MultiplyProjection(right Projection) Vector4 {
  panic("TODO: implement")
}

func (me *Vector4) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Vector4) InArray(right Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
