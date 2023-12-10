// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeTransformConstant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTransformConstant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTransformConstant) BaseClass() string {
  return "VisualShaderNodeTransformConstant"
}



// Enums

func (me *VisualShaderNodeTransformConstant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTransformConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTransformConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeTransformConstant) SetConstant(constant Transform3D, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTransformConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2952846383) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(constant.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTransformConstant) GetConstant() Transform3D {
  classNameV := StringNameFromStr("VisualShaderNodeTransformConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeTransformConstant) GetPropConstant() Transform3D {
  panic("TODO: implement")
}

func (me *VisualShaderNodeTransformConstant) SetPropConstant(value Transform3D) {
  panic("TODO: implement")
}