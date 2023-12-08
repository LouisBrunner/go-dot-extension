// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParameter(obj gdc.ObjectPtr) *VisualShaderNodeParameter {
  return &VisualShaderNodeParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeParameter) BaseClass() string {
  return "VisualShaderNodeParameter"
}
