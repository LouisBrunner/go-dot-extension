// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Quaternion struct {
  obj gdc.ObjectPtr
}

func (me *Quaternion) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Quaternion) BaseClass() string {
  return "Quaternion"
}

// TODO: needed?
// const (
// // )

var (
  QuaternionIdentity = "Quaternion(0, 0, 0, 1)" // TODO: construct correctly
)

func  (me *Quaternion) Length() { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) LengthSquared() { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) Normalized() { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) IsNormalized() { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) IsEqualApprox(to Quaternion, ) { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) IsFinite() { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) Inverse() { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) Log() { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) Exp() { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) AngleTo(to Quaternion, ) { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) Dot(with Quaternion, ) { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) Slerp(to Quaternion, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) Slerpni(to Quaternion, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) SphericalCubicInterpolate(b Quaternion, pre_a Quaternion, post_b Quaternion, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) SphericalCubicInterpolateInTime(b Quaternion, pre_a Quaternion, post_b Quaternion, weight float32, b_t float32, pre_a_t float32, post_b_t float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) GetEuler(order int, ) { // TODO: return value
  // TODO: implement
}

func  QuaternionFromEuler(euler Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) GetAxis() { // TODO: return value
  // TODO: implement
}

func  (me *Quaternion) GetAngle() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
