// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Line2D struct {
  obj gdc.ObjectPtr
}

func (me *Line2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Line2D) BaseClass() string {
  return "Line2D"
}

type Line2DLineJointMode int
const (
  Line2DLineJointModeLineJointSharp Line2DLineJointMode = 0
  Line2DLineJointModeLineJointBevel Line2DLineJointMode = 1
  Line2DLineJointModeLineJointRound Line2DLineJointMode = 2
)

type Line2DLineCapMode int
const (
  Line2DLineCapModeLineCapNone Line2DLineCapMode = 0
  Line2DLineCapModeLineCapBox Line2DLineCapMode = 1
  Line2DLineCapModeLineCapRound Line2DLineCapMode = 2
)

type Line2DLineTextureMode int
const (
  Line2DLineTextureModeLineTextureNone Line2DLineTextureMode = 0
  Line2DLineTextureModeLineTextureTile Line2DLineTextureMode = 1
  Line2DLineTextureModeLineTextureStretch Line2DLineTextureMode = 2
)

func  (me *Line2D) SetPoints(points PackedVector2Array, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetPoints() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetPointPosition(index int, position Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetPointPosition(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetPointCount() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) AddPoint(position Vector2, index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) RemovePoint(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) ClearPoints() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetWidth(width float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetWidth() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetCurve(curve Curve, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetCurve() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetDefaultColor(color Color, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetDefaultColor() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetGradient(color Gradient, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetGradient() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetTexture(texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetTexture() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetTextureMode(mode Line2DLineTextureMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetTextureMode() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetJointMode(mode Line2DLineJointMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetJointMode() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetBeginCapMode(mode Line2DLineCapMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetBeginCapMode() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetEndCapMode(mode Line2DLineCapMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetEndCapMode() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetSharpLimit(limit float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetSharpLimit() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetRoundPrecision(precision int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetRoundPrecision() { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) SetAntialiased(antialiased bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Line2D) GetAntialiased() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
