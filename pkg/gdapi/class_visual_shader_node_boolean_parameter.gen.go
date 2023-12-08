// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeBooleanParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeBooleanParameter(obj gdc.ObjectPtr) *VisualShaderNodeBooleanParameter {
  return &VisualShaderNodeBooleanParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeBooleanParameter) BaseClass() string {
  return "VisualShaderNodeBooleanParameter"
}
