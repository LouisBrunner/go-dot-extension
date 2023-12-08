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

type VisualShaderNodeParameterQualifier int
const (
  VisualShaderNodeParameterQualifierQualNone VisualShaderNodeParameterQualifier = 0
  VisualShaderNodeParameterQualifierQualGlobal VisualShaderNodeParameterQualifier = 1
  VisualShaderNodeParameterQualifierQualInstance VisualShaderNodeParameterQualifier = 2
  VisualShaderNodeParameterQualifierQualMax VisualShaderNodeParameterQualifier = 3
)

func  (me *VisualShaderNodeParameter) SetParameterName(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeParameter) GetParameterName() { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeParameter) SetQualifier(qualifier VisualShaderNodeParameterQualifier, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeParameter) GetQualifier() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
