// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Transform2D struct {
  obj gdc.ObjectPtr
}

func (me *Transform2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Transform2D) BaseClass() string {
  return "Transform2D"
}

// TODO: needed?
// const (
// // )

var (
  Transform2DIdentity = "Transform2D(1, 0, 0, 1, 0, 0)" // TODO: construct correctly
  Transform2DFlipX = "Transform2D(-1, 0, 0, 1, 0, 0)" // TODO: construct correctly
  Transform2DFlipY = "Transform2D(1, 0, 0, -1, 0, 0)" // TODO: construct correctly
)

func  (me *Transform2D) Inverse() { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) AffineInverse() { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) GetRotation() { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) GetOrigin() { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) GetScale() { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) GetSkew() { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) Orthonormalized() { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) Rotated(angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) RotatedLocal(angle float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) Scaled(scale Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) ScaledLocal(scale Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) Translated(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) TranslatedLocal(offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) Determinant() { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) BasisXform(v Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) BasisXformInv(v Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) InterpolateWith(xform Transform2D, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) IsEqualApprox(xform Transform2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) IsFinite() { // TODO: return value
  // TODO: implement
}

func  (me *Transform2D) LookingAt(target Vector2, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
