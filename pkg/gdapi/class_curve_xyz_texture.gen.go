// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CurveXYZTexture struct {
  obj gdc.ObjectPtr
}

func (me *CurveXYZTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CurveXYZTexture) BaseClass() string {
  return "CurveXYZTexture"
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

func  (me *CurveXYZTexture) SetWidth(width int, )  {
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
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *CurveXYZTexture) GetPropWidth() int {
  panic("TODO: implement")
}

func (me *CurveXYZTexture) SetPropWidth(value int) {
  panic("TODO: implement")
}

func (me *CurveXYZTexture) GetPropCurveX() Curve {
  panic("TODO: implement")
}

func (me *CurveXYZTexture) SetPropCurveX(value Curve) {
  panic("TODO: implement")
}

func (me *CurveXYZTexture) GetPropCurveY() Curve {
  panic("TODO: implement")
}

func (me *CurveXYZTexture) SetPropCurveY(value Curve) {
  panic("TODO: implement")
}

func (me *CurveXYZTexture) GetPropCurveZ() Curve {
  panic("TODO: implement")
}

func (me *CurveXYZTexture) SetPropCurveZ(value Curve) {
  panic("TODO: implement")
}