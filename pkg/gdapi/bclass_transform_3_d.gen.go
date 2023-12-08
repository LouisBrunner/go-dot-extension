// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Transform3D struct {
  obj gdc.ObjectPtr
}

func (me *Transform3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Transform3D) BaseClass() string {
  return "Transform3D"
}

// TODO: needed?
// const (
// // )

var (
  Transform3DIdentity = "Transform3D(1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipX = "Transform3D(-1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipY = "Transform3D(1, 0, 0, 0, -1, 0, 0, 0, 1, 0, 0, 0)" // TODO: construct correctly
  Transform3DFlipZ = "Transform3D(1, 0, 0, 0, 1, 0, 0, 0, -1, 0, 0, 0)" // TODO: construct correctly
)

func  (me *Transform3D) Inverse() { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) AffineInverse() { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) Orthonormalized() { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) Rotated(axis Vector3, angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) RotatedLocal(axis Vector3, angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) Scaled(scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) ScaledLocal(scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) Translated(offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) TranslatedLocal(offset Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) LookingAt(target Vector3, up Vector3, use_model_front bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) InterpolateWith(xform Transform3D, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) IsEqualApprox(xform Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform3D) IsFinite() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
