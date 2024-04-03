// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CurveTexture struct {
  Texture2D
}

func (me *CurveTexture) BaseClass() string {
  return "CurveTexture"
}

func NewCurveTexture() *CurveTexture {
  str := StringNameFromStr("CurveTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CurveTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type CurveTextureTextureMode int
const (
  CurveTextureTextureModeTextureModeRgb CurveTextureTextureMode = 0
  CurveTextureTextureModeTextureModeRed CurveTextureTextureMode = 1
)

func (me *CurveTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CurveTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CurveTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CurveTexture) SetWidth(width int64, )  {
  classNameV := StringNameFromStr("CurveTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CurveTexture) SetCurve(curve Curve, )  {
  classNameV := StringNameFromStr("CurveTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CurveTexture) GetCurve() Curve {
  classNameV := StringNameFromStr("CurveTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CurveTexture) SetTextureMode(texture_mode CurveTextureTextureMode, )  {
  classNameV := StringNameFromStr("CurveTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321955367) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&texture_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CurveTexture) GetTextureMode() CurveTextureTextureMode {
  classNameV := StringNameFromStr("CurveTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 715756376) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret CurveTextureTextureMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
