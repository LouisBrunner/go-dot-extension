// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeCurveXYZTexture struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeCurveXYZTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeCurveXYZTexture) BaseClass() string {
  return "VisualShaderNodeCurveXYZTexture"
}



// Enums

func (me *VisualShaderNodeCurveXYZTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCurveXYZTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCurveXYZTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeCurveXYZTexture) SetTexture(texture CurveXYZTexture, )  {
  classNameV := StringNameFromStr("VisualShaderNodeCurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 8031783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeCurveXYZTexture) GetTexture() CurveXYZTexture {
  classNameV := StringNameFromStr("VisualShaderNodeCurveXYZTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1950275015) // FIXME: should cache?
  var ret CurveXYZTexture
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeCurveXYZTexture) GetPropTexture() CurveXYZTexture {
  panic("TODO: implement")
}

func (me *VisualShaderNodeCurveXYZTexture) SetPropTexture(value CurveXYZTexture) {
  panic("TODO: implement")
}