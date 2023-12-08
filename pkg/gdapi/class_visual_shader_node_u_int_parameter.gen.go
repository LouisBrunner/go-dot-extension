// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeUIntParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeUIntParameter(obj gdc.ObjectPtr) *VisualShaderNodeUIntParameter {
  return &VisualShaderNodeUIntParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeUIntParameter) BaseClass() string {
  return "VisualShaderNodeUIntParameter"
}
