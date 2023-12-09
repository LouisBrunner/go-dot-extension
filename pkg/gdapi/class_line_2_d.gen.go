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



// Enums

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

func (me *Line2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Line2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Line2D) SetPoints(points PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetPoints()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetPointPosition(index int, position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetPointPosition(index int, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetPointCount()  {
  panic("TODO: implement")
}

func  (me *Line2D) AddPoint(position Vector2, index int, )  {
  panic("TODO: implement")
}

func  (me *Line2D) RemovePoint(index int, )  {
  panic("TODO: implement")
}

func  (me *Line2D) ClearPoints()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetWidth(width float32, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetWidth()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetCurve(curve Curve, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetCurve()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetDefaultColor(color Color, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetDefaultColor()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetGradient(color Gradient, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetGradient()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetTexture(texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetTexture()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetTextureMode(mode Line2DLineTextureMode, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetTextureMode()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetJointMode(mode Line2DLineJointMode, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetJointMode()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetBeginCapMode(mode Line2DLineCapMode, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetBeginCapMode()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetEndCapMode(mode Line2DLineCapMode, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetEndCapMode()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetSharpLimit(limit float32, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetSharpLimit()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetRoundPrecision(precision int, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetRoundPrecision()  {
  panic("TODO: implement")
}

func  (me *Line2D) SetAntialiased(antialiased bool, )  {
  panic("TODO: implement")
}

func  (me *Line2D) GetAntialiased()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
