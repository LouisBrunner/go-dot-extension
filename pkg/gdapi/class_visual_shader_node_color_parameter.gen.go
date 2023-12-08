// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeColorParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeColorParameter(obj gdc.ObjectPtr) *VisualShaderNodeColorParameter {
  return &VisualShaderNodeColorParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeColorParameter) BaseClass() string {
  return "VisualShaderNodeColorParameter"
}
