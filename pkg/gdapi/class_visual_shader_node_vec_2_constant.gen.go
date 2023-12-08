// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVec2Constant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVec2Constant(obj gdc.ObjectPtr) *VisualShaderNodeVec2Constant {
  return &VisualShaderNodeVec2Constant{
    obj: obj,
  }
}

func (me *VisualShaderNodeVec2Constant) BaseClass() string {
  return "VisualShaderNodeVec2Constant"
}
