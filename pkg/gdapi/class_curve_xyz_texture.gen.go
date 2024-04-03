// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CurveXYZTexture struct {
  Texture2D
}

func (me *CurveXYZTexture) BaseClass() string {
  return "CurveXYZTexture"
}

func NewCurveXYZTexture() *CurveXYZTexture {
  str := StringNameFromStr("CurveXYZTexture") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CurveXYZTexture{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CurveXYZTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CurveXYZTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CurveXYZTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CurveXYZTexture) SetWidth(width int64, )  {
  classNameV := StringNameFromStr("CurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_width")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&width), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CurveXYZTexture) SetCurveX(curve Curve, )  {
  classNameV := StringNameFromStr("CurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CurveXYZTexture) GetCurveX() Curve {
  classNameV := StringNameFromStr("CurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CurveXYZTexture) SetCurveY(curve Curve, )  {
  classNameV := StringNameFromStr("CurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CurveXYZTexture) GetCurveY() Curve {
  classNameV := StringNameFromStr("CurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CurveXYZTexture) SetCurveZ(curve Curve, )  {
  classNameV := StringNameFromStr("CurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve_z")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CurveXYZTexture) GetCurveZ() Curve {
  classNameV := StringNameFromStr("CurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve_z")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
