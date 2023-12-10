// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeParameter struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParameter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParameter) BaseClass() string {
  return "VisualShaderNodeParameter"
}



// Enums

type VisualShaderNodeParameterQualifier int
const (
  VisualShaderNodeParameterQualifierQualNone VisualShaderNodeParameterQualifier = 0
  VisualShaderNodeParameterQualifierQualGlobal VisualShaderNodeParameterQualifier = 1
  VisualShaderNodeParameterQualifierQualInstance VisualShaderNodeParameterQualifier = 2
  VisualShaderNodeParameterQualifierQualMax VisualShaderNodeParameterQualifier = 3
)

func (me *VisualShaderNodeParameter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParameter) SetParameterName(name String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parameter_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParameter) GetParameterName() String {
  classNameV := StringNameFromStr("VisualShaderNodeParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parameter_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeParameter) SetQualifier(qualifier VisualShaderNodeParameterQualifier, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_qualifier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1276489447) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&qualifier), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeParameter) GetQualifier() VisualShaderNodeParameterQualifier {
  classNameV := StringNameFromStr("VisualShaderNodeParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_qualifier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3558406205) // FIXME: should cache?
  var ret VisualShaderNodeParameterQualifier
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeParameter) GetPropParameterName() StringName {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParameter) SetPropParameterName(value StringName) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParameter) GetPropQualifier() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeParameter) SetPropQualifier(value int) {
  panic("TODO: implement")
}