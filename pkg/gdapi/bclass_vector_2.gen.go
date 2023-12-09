// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector2 struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Vector2Zero = "Vector2(0, 0)" // TODO: construct correctly
  Vector2One = "Vector2(1, 1)" // TODO: construct correctly
  Vector2Inf = "Vector2(inf, inf)" // TODO: construct correctly
  Vector2Left = "Vector2(-1, 0)" // TODO: construct correctly
  Vector2Right = "Vector2(1, 0)" // TODO: construct correctly
  Vector2Up = "Vector2(0, -1)" // TODO: construct correctly
  Vector2Down = "Vector2(0, 1)" // TODO: construct correctly
)

// Enums

type Vector2Axis int
const (
  Vector2AxisAxisX Vector2Axis = 0
  Vector2AxisAxisY Vector2Axis = 1
)

// Constructors

func NewVector2() Vector2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Vector2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector2FromVector2(from Vector2, ) Vector2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector2FromVector2I(from Vector2i, ) Vector2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector2FromFloat32Float32(x float32, y float32, ) Vector2 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector2))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector2, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), }))
  return Vector2{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Vector2) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Vector2) Type() gdc.VariantType {
  return gdc.VariantTypeVector2
}

func (me *Vector2) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Vector2) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Vector2) Angle() float32 {
  panic("TODO: implement")
}

func  (me *Vector2) AngleTo(to Vector2, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector2) AngleToPoint(to Vector2, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector2) DirectionTo(to Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) DistanceTo(to Vector2, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector2) DistanceSquaredTo(to Vector2, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector2) Length() float32 {
  panic("TODO: implement")
}

func  (me *Vector2) LengthSquared() float32 {
  panic("TODO: implement")
}

func  (me *Vector2) LimitLength(length float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Normalized() Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) IsNormalized() bool {
  panic("TODO: implement")
}

func  (me *Vector2) IsEqualApprox(to Vector2, ) bool {
  panic("TODO: implement")
}

func  (me *Vector2) IsZeroApprox() bool {
  panic("TODO: implement")
}

func  (me *Vector2) IsFinite() bool {
  panic("TODO: implement")
}

func  (me *Vector2) Posmod(mod float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Posmodv(modv Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Project(b Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Lerp(to Vector2, weight float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Slerp(to Vector2, weight float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) CubicInterpolate(b Vector2, pre_a Vector2, post_b Vector2, weight float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) CubicInterpolateInTime(b Vector2, pre_a Vector2, post_b Vector2, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) BezierInterpolate(control_1 Vector2, control_2 Vector2, end Vector2, t float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) BezierDerivative(control_1 Vector2, control_2 Vector2, end Vector2, t float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) MaxAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector2) MinAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector2) MoveToward(to Vector2, delta float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Rotated(angle float32, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Orthogonal() Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Floor() Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Ceil() Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Round() Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Aspect() float32 {
  panic("TODO: implement")
}

func  (me *Vector2) Dot(with Vector2, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector2) Slide(n Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Bounce(n Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Reflect(n Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Cross(with Vector2, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector2) Abs() Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Sign() Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Clamp(min Vector2, max Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  (me *Vector2) Snapped(step Vector2, ) Vector2 {
  panic("TODO: implement")
}

func  Vector2FromAngle(angle float32, ) Vector2 {
  panic("TODO: implement")
}

// Operators

func (me *Vector2) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector2) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector2) UnaryMinus() Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) UnaryPlus() Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) Not() bool {
  panic("TODO: implement")
}

func (me *Vector2) MultiplyInt(right int) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) DivideInt(right int) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) MultiplyFloat32(right float32) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) DivideFloat32(right float32) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) EqualsVector2(right Vector2) bool {
  panic("TODO: implement")
}

func (me *Vector2) NotEqualsVector2(right Vector2) bool {
  panic("TODO: implement")
}

func (me *Vector2) LessThanVector2(right Vector2) bool {
  panic("TODO: implement")
}

func (me *Vector2) LessThanOrEqualsVector2(right Vector2) bool {
  panic("TODO: implement")
}

func (me *Vector2) GreaterThanVector2(right Vector2) bool {
  panic("TODO: implement")
}

func (me *Vector2) GreaterThanOrEqualsVector2(right Vector2) bool {
  panic("TODO: implement")
}

func (me *Vector2) AddVector2(right Vector2) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) SubtractVector2(right Vector2) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) MultiplyVector2(right Vector2) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) DivideVector2(right Vector2) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) MultiplyTransform2D(right Transform2D) Vector2 {
  panic("TODO: implement")
}

func (me *Vector2) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Vector2) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Vector2) InPackedVector2Array(right PackedVector2Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
