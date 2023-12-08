// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector2 struct {
  obj gdc.ObjectPtr
}

func (me *Vector2) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Vector2) BaseClass() string {
  return "Vector2"
}

// TODO: needed?
// const (
//   Vector2AxisX = 0
//   Vector2AxisY = 1
// // )

var (
  Vector2Zero = "Vector2(0, 0)" // TODO: construct correctly
  Vector2One = "Vector2(1, 1)" // TODO: construct correctly
  Vector2Inf = "Vector2(inf, inf)" // TODO: construct correctly
  Vector2Left = "Vector2(-1, 0)" // TODO: construct correctly
  Vector2Right = "Vector2(1, 0)" // TODO: construct correctly
  Vector2Up = "Vector2(0, -1)" // TODO: construct correctly
  Vector2Down = "Vector2(0, 1)" // TODO: construct correctly
)

type Vector2Axis int
const (
  Vector2AxisAxisX Vector2Axis = 0
  Vector2AxisAxisY Vector2Axis = 1
)

func  (me *Vector2) Angle() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) AngleTo(to Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) AngleToPoint(to Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) DirectionTo(to Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) DistanceTo(to Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) DistanceSquaredTo(to Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Length() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) LengthSquared() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) LimitLength(length float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Normalized() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) IsNormalized() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) IsEqualApprox(to Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) IsZeroApprox() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) IsFinite() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Posmod(mod float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Posmodv(modv Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Project(b Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Lerp(to Vector2, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Slerp(to Vector2, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) CubicInterpolate(b Vector2, pre_a Vector2, post_b Vector2, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) CubicInterpolateInTime(b Vector2, pre_a Vector2, post_b Vector2, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) BezierInterpolate(control_1 Vector2, control_2 Vector2, end Vector2, t float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) BezierDerivative(control_1 Vector2, control_2 Vector2, end Vector2, t float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) MaxAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) MinAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) MoveToward(to Vector2, delta float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Rotated(angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Orthogonal() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Floor() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Ceil() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Round() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Aspect() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Dot(with Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Slide(n Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Bounce(n Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Reflect(n Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Cross(with Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Abs() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Sign() { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Clamp(min Vector2, max Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector2) Snapped(step Vector2, ) { // TODO: return value
  // TODO: implement
}

func  Vector2FromAngle(angle float32, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
