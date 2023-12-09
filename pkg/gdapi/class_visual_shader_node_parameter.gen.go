// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *VisualShaderNodeParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeParameter) SetParameterName(name String, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeParameter) GetParameterName()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeParameter) SetQualifier(qualifier VisualShaderNodeParameterQualifier, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeParameter) GetQualifier()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
