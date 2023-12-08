// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector3 struct {
  obj gdc.ObjectPtr
}

func (me *Vector3) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Vector3) BaseClass() string {
  return "Vector3"
}

// TODO: needed?
// const (
//   Vector3AxisX = 0
//   Vector3AxisY = 1
//   Vector3AxisZ = 2
// // )

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

type Vector3Axis int
const (
  Vector3AxisAxisX Vector3Axis = 0
  Vector3AxisAxisY Vector3Axis = 1
  Vector3AxisAxisZ Vector3Axis = 2
)

func  (me *Vector3) MinAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) MaxAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) AngleTo(to Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) SignedAngleTo(to Vector3, axis Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) DirectionTo(to Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) DistanceTo(to Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) DistanceSquaredTo(to Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Length() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) LengthSquared() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) LimitLength(length float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Normalized() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) IsNormalized() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) IsEqualApprox(to Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) IsZeroApprox() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) IsFinite() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Inverse() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Clamp(min Vector3, max Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Snapped(step Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Rotated(axis Vector3, angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Lerp(to Vector3, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Slerp(to Vector3, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) CubicInterpolate(b Vector3, pre_a Vector3, post_b Vector3, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) CubicInterpolateInTime(b Vector3, pre_a Vector3, post_b Vector3, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) BezierInterpolate(control_1 Vector3, control_2 Vector3, end Vector3, t float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) BezierDerivative(control_1 Vector3, control_2 Vector3, end Vector3, t float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) MoveToward(to Vector3, delta float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Dot(with Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Cross(with Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Outer(with Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Abs() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Floor() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Ceil() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Round() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Posmod(mod float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Posmodv(modv Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Project(b Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Slide(n Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Bounce(n Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Reflect(n Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) Sign() { // TODO: return value
  // TODO: implement
}

func  (me *Vector3) OctahedronEncode() { // TODO: return value
  // TODO: implement
}

func  Vector3OctahedronDecode(uv Vector2, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
