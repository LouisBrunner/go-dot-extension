// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeParameterRef struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParameterRef) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParameterRef) BaseClass() string {
  return "VisualShaderNodeParameterRef"
}



// Enums

func (me *VisualShaderNodeParameterRef) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParameterRef) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParameterRef) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParameterRef) SetParameterName(name String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParameterRef")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parameter_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParameterRef) GetParameterName() String {
  classNameV := StringNameFromStr("VisualShaderNodeParameterRef")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parameter_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeParameterRef) GetPropParameterName() StringName {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParameterRef) SetPropParameterName(value StringName) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParameterRef) GetPropParamType() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParameterRef) SetPropParamType(value int) {
  panic("TODO: implement")
}