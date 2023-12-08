// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Vector4 struct {
  obj gdc.ObjectPtr
}

func (me *Vector4) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Vector4) BaseClass() string {
  return "Vector4"
}

// TODO: needed?
// const (
//   Vector4AxisX = 0
//   Vector4AxisY = 1
//   Vector4AxisZ = 2
//   Vector4AxisW = 3
// // )

var (
  Vector4Zero = "Vector4(0, 0, 0, 0)" // TODO: construct correctly
  Vector4One = "Vector4(1, 1, 1, 1)" // TODO: construct correctly
  Vector4Inf = "Vector4(inf, inf, inf, inf)" // TODO: construct correctly
)

type Vector4Axis int
const (
  Vector4AxisAxisX Vector4Axis = 0
  Vector4AxisAxisY Vector4Axis = 1
  Vector4AxisAxisZ Vector4Axis = 2
  Vector4AxisAxisW Vector4Axis = 3
)

func  (me *Vector4) MinAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) MaxAxisIndex() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Length() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) LengthSquared() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Abs() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Sign() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Floor() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Ceil() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Round() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Lerp(to Vector4, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) CubicInterpolate(b Vector4, pre_a Vector4, post_b Vector4, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) CubicInterpolateInTime(b Vector4, pre_a Vector4, post_b Vector4, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Posmod(mod float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Posmodv(modv Vector4, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Snapped(step Vector4, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Clamp(min Vector4, max Vector4, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Normalized() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) IsNormalized() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) DirectionTo(to Vector4, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) DistanceTo(to Vector4, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) DistanceSquaredTo(to Vector4, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Dot(with Vector4, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) Inverse() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) IsEqualApprox(to Vector4, ) { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) IsZeroApprox() { // TODO: return value
  // TODO: implement
}

func  (me *Vector4) IsFinite() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
