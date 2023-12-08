// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  VisualShaderNodeParameterQualNone VisualShaderNodeParameterQualifier = 0
  VisualShaderNodeParameterQualGlobal VisualShaderNodeParameterQualifier = 1
  VisualShaderNodeParameterQualInstance VisualShaderNodeParameterQualifier = 2
  VisualShaderNodeParameterQualMax VisualShaderNodeParameterQualifier = 3
)
