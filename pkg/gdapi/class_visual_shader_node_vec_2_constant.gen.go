// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeVec2Constant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeVec2Constant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeVec2Constant) BaseClass() string {
  return "VisualShaderNodeVec2Constant"
}



// Enums

func (me *VisualShaderNodeVec2Constant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVec2Constant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec2Constant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeVec2Constant) SetConstant(constant Vector2, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVec2Constant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(constant.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeVec2Constant) GetConstant() Vector2 {
  classNameV := StringNameFromStr("VisualShaderNodeVec2Constant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeVec2Constant) GetPropConstant() Vector2 {
  panic("TODO: implement")
}

func (me *VisualShaderNodeVec2Constant) SetPropConstant(value Vector2) {
  panic("TODO: implement")
}