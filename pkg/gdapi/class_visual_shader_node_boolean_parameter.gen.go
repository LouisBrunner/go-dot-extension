// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeBooleanParameter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeBooleanParameter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeBooleanParameter) BaseClass() string {
  return "VisualShaderNodeBooleanParameter"
}



// Enums

func (me *VisualShaderNodeBooleanParameter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeBooleanParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeBooleanParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeBooleanParameter) SetDefaultValueEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeBooleanParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeBooleanParameter) IsDefaultValueEnabled() bool {
  classNameV := StringNameFromStr("VisualShaderNodeBooleanParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeBooleanParameter) SetDefaultValue(value bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeBooleanParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeBooleanParameter) GetDefaultValue() bool {
  classNameV := StringNameFromStr("VisualShaderNodeBooleanParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeBooleanParameter) GetPropDefaultValueEnabled() bool {
  panic("TODO: implement")
}

func (me *VisualShaderNodeBooleanParameter) SetPropDefaultValueEnabled(value bool) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeBooleanParameter) GetPropDefaultValue() bool {
  panic("TODO: implement")
}

func (me *VisualShaderNodeBooleanParameter) SetPropDefaultValue(value bool) {
  panic("TODO: implement")
}