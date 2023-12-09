// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *VisualShaderNodeBooleanParameter) IsDefaultValueEnabled()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeBooleanParameter) SetDefaultValue(value bool, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeBooleanParameter) GetDefaultValue()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
