// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Basis struct {
  obj gdc.ObjectPtr
}

func (me *Basis) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Basis) BaseClass() string {
  return "Basis"
}

// TODO: needed?
// const (
// // )

var (
  BasisIdentity = "Basis(1, 0, 0, 0, 1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipX = "Basis(-1, 0, 0, 0, 1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipY = "Basis(1, 0, 0, 0, -1, 0, 0, 0, 1)" // TODO: construct correctly
  BasisFlipZ = "Basis(1, 0, 0, 0, 1, 0, 0, 0, -1)" // TODO: construct correctly
)

func  (me *Basis) Inverse() { // TODO: return value
  // TODO: implement
}

func  (me *Basis) Transposed() { // TODO: return value
  // TODO: implement
}

func  (me *Basis) Orthonormalized() { // TODO: return value
  // TODO: implement
}

func  (me *Basis) Determinant() { // TODO: return value
  // TODO: implement
}

func  (me *Basis) Rotated(axis Vector3, angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Basis) Scaled(scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Basis) GetScale() { // TODO: return value
  // TODO: implement
}

func  (me *Basis) GetEuler(order int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Basis) Tdotx(with Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Basis) Tdoty(with Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Basis) Tdotz(with Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Basis) Slerp(to Basis, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Basis) IsEqualApprox(b Basis, ) { // TODO: return value
  // TODO: implement
}

func  (me *Basis) IsFinite() { // TODO: return value
  // TODO: implement
}

func  (me *Basis) GetRotationQuaternion() { // TODO: return value
  // TODO: implement
}

func  BasisLookingAt(target Vector3, up Vector3, use_model_front bool, ) { // TODO: return value
  // TODO: implement
}

func  BasisFromScale(scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  BasisFromEuler(euler Vector3, order int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
