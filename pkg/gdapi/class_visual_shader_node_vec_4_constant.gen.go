// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVec4Constant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVec4Constant(obj gdc.ObjectPtr) *VisualShaderNodeVec4Constant {
  return &VisualShaderNodeVec4Constant{
    obj: obj,
  }
}

func (me *VisualShaderNodeVec4Constant) BaseClass() string {
  return "VisualShaderNodeVec4Constant"
}
