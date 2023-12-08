// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTransformParameter(obj gdc.ObjectPtr) *VisualShaderNodeTransformParameter {
  return &VisualShaderNodeTransformParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeTransformParameter) BaseClass() string {
  return "VisualShaderNodeTransformParameter"
}
