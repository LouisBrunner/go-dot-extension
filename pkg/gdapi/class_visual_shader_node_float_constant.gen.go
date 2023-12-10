// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeFloatConstant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeFloatConstant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeFloatConstant) BaseClass() string {
  return "VisualShaderNodeFloatConstant"
}



// Enums

func (me *VisualShaderNodeFloatConstant) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeFloatConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeFloatConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeFloatConstant) SetConstant(constant float32, )  {
  classNameV := StringNameFromStr("VisualShaderNodeFloatConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&constant), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeFloatConstant) GetConstant() float32 {
  classNameV := StringNameFromStr("VisualShaderNodeFloatConstant")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeFloatConstant) GetPropConstant() float32 {
  panic("TODO: implement")
}

func (me *VisualShaderNodeFloatConstant) SetPropConstant(value float32) {
  panic("TODO: implement")
}