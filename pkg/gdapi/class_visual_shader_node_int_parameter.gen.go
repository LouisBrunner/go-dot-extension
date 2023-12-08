// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIntParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeIntParameter(obj gdc.ObjectPtr) *VisualShaderNodeIntParameter {
  return &VisualShaderNodeIntParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeIntParameter) BaseClass() string {
  return "VisualShaderNodeIntParameter"
}
