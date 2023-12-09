// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CurveTexture struct {
  obj gdc.ObjectPtr
}

func (me *CurveTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CurveTexture) BaseClass() string {
  return "CurveTexture"
}



// Enums

type CurveTextureTextureMode int
const (
  CurveTextureTextureModeTextureModeRgb CurveTextureTextureMode = 0
  CurveTextureTextureModeTextureModeRed CurveTextureTextureMode = 1
)

func (me *CurveTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CurveTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *CurveTexture) SetWidth(width int, )  {
  panic("TODO: implement")
}

func  (me *CurveTexture) SetCurve(curve Curve, )  {
  panic("TODO: implement")
}

func  (me *CurveTexture) GetCurve()  {
  panic("TODO: implement")
}

func  (me *CurveTexture) SetTextureMode(texture_mode CurveTextureTextureMode, )  {
  panic("TODO: implement")
}

func  (me *CurveTexture) GetTextureMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
