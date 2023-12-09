// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector3 struct {
  iface gdc.Interface
  ptr gdc.TypePtr
}

// Constants

var (
  Vector3Zero = "Vector3(0, 0, 0)" // TODO: construct correctly
  Vector3One = "Vector3(1, 1, 1)" // TODO: construct correctly
  Vector3Inf = "Vector3(inf, inf, inf)" // TODO: construct correctly
  Vector3Left = "Vector3(-1, 0, 0)" // TODO: construct correctly
  Vector3Right = "Vector3(1, 0, 0)" // TODO: construct correctly
  Vector3Up = "Vector3(0, 1, 0)" // TODO: construct correctly
  Vector3Down = "Vector3(0, -1, 0)" // TODO: construct correctly
  Vector3Forward = "Vector3(0, 0, -1)" // TODO: construct correctly
  Vector3Back = "Vector3(0, 0, 1)" // TODO: construct correctly
  Vector3ModelLeft = "Vector3(1, 0, 0)" // TODO: construct correctly
  Vector3ModelRight = "Vector3(-1, 0, 0)" // TODO: construct correctly
  Vector3ModelTop = "Vector3(0, 1, 0)" // TODO: construct correctly
  Vector3ModelBottom = "Vector3(0, -1, 0)" // TODO: construct correctly
  Vector3ModelFront = "Vector3(0, 0, 1)" // TODO: construct correctly
  Vector3ModelRear = "Vector3(0, 0, -1)" // TODO: construct correctly
)

// Enums

type Vector3Axis int
const (
  Vector3AxisAxisX Vector3Axis = 0
  Vector3AxisAxisY Vector3Axis = 1
  Vector3AxisAxisZ Vector3Axis = 2
)

// Constructors

func NewVector3() Vector3 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector3))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector3, 0) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{}))
  return Vector3{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector3FromVector3(from Vector3, ) Vector3 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector3))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector3, 1) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector3{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector3FromVector3I(from Vector3i, ) Vector3 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector3))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector3, 2) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{from.AsCTypePtr(), }))
  return Vector3{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

func NewVector3FromFloat32Float32Float32(x float32, y float32, z float32, ) Vector3 {
  ptr := (gdc.UninitializedTypePtr)(cmalloc(classSizeVector3))
  ctr := giface.VariantGetPtrConstructor(gdc.VariantTypeVector3, 3) // FIXME: should cache?
  giface.CallPtrConstructor(ctr, ptr, unsafe.SliceData([]gdc.ConstTypePtr{gdc.ConstTypePtr(&x), gdc.ConstTypePtr(&y), gdc.ConstTypePtr(&z), }))
  return Vector3{
    iface: giface,
    ptr: gdc.TypePtr(ptr),
  }
}

// Destructor
func (me *Vector3) Destroy() {
  if me.ptr == nil {
    return
  }
	cfree(unsafe.Pointer(me.ptr))
  me.ptr = nil
}

func (me *Vector3) Type() gdc.VariantType {
  return gdc.VariantTypeVector3
}

func (me *Vector3) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.ptr)
}

func (me *Vector3) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.ptr)
}

// Methods

func  (me *Vector3) MinAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector3) MaxAxisIndex() int {
  panic("TODO: implement")
}

func  (me *Vector3) AngleTo(to Vector3, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector3) SignedAngleTo(to Vector3, axis Vector3, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector3) DirectionTo(to Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) DistanceTo(to Vector3, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector3) DistanceSquaredTo(to Vector3, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector3) Length() float32 {
  panic("TODO: implement")
}

func  (me *Vector3) LengthSquared() float32 {
  panic("TODO: implement")
}

func  (me *Vector3) LimitLength(length float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Normalized() Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) IsNormalized() bool {
  panic("TODO: implement")
}

func  (me *Vector3) IsEqualApprox(to Vector3, ) bool {
  panic("TODO: implement")
}

func  (me *Vector3) IsZeroApprox() bool {
  panic("TODO: implement")
}

func  (me *Vector3) IsFinite() bool {
  panic("TODO: implement")
}

func  (me *Vector3) Inverse() Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Clamp(min Vector3, max Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Snapped(step Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Rotated(axis Vector3, angle float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Lerp(to Vector3, weight float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Slerp(to Vector3, weight float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) CubicInterpolate(b Vector3, pre_a Vector3, post_b Vector3, weight float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) CubicInterpolateInTime(b Vector3, pre_a Vector3, post_b Vector3, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) BezierInterpolate(control_1 Vector3, control_2 Vector3, end Vector3, t float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) BezierDerivative(control_1 Vector3, control_2 Vector3, end Vector3, t float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) MoveToward(to Vector3, delta float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Dot(with Vector3, ) float32 {
  panic("TODO: implement")
}

func  (me *Vector3) Cross(with Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Outer(with Vector3, ) Basis {
  panic("TODO: implement")
}

func  (me *Vector3) Abs() Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Floor() Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Ceil() Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Round() Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Posmod(mod float32, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Posmodv(modv Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Project(b Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Slide(n Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Bounce(n Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Reflect(n Vector3, ) Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) Sign() Vector3 {
  panic("TODO: implement")
}

func  (me *Vector3) OctahedronEncode() Vector2 {
  panic("TODO: implement")
}

func  Vector3OctahedronDecode(uv Vector2, ) Vector3 {
  panic("TODO: implement")
}

// Operators

func (me *Vector3) EqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector3) NotEqualsVariant(right Variant) bool {
  panic("TODO: implement")
}

func (me *Vector3) UnaryMinus() Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) UnaryPlus() Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) Not() bool {
  panic("TODO: implement")
}

func (me *Vector3) MultiplyInt(right int) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) DivideInt(right int) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) MultiplyFloat32(right float32) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) DivideFloat32(right float32) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) EqualsVector3(right Vector3) bool {
  panic("TODO: implement")
}

func (me *Vector3) NotEqualsVector3(right Vector3) bool {
  panic("TODO: implement")
}

func (me *Vector3) LessThanVector3(right Vector3) bool {
  panic("TODO: implement")
}

func (me *Vector3) LessThanOrEqualsVector3(right Vector3) bool {
  panic("TODO: implement")
}

func (me *Vector3) GreaterThanVector3(right Vector3) bool {
  panic("TODO: implement")
}

func (me *Vector3) GreaterThanOrEqualsVector3(right Vector3) bool {
  panic("TODO: implement")
}

func (me *Vector3) AddVector3(right Vector3) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) SubtractVector3(right Vector3) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) MultiplyVector3(right Vector3) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) DivideVector3(right Vector3) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) MultiplyQuaternion(right Quaternion) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) MultiplyBasis(right Basis) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) MultiplyTransform3D(right Transform3D) Vector3 {
  panic("TODO: implement")
}

func (me *Vector3) InDictionary(right Dictionary) bool {
  panic("TODO: implement")
}

func (me *Vector3) InArray(right Array) bool {
  panic("TODO: implement")
}

func (me *Vector3) InPackedVector3Array(right PackedVector3Array) bool {
  panic("TODO: implement")
}

// TODO: members (bclass)
