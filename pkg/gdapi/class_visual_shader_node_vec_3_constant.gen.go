// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVec3Constant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVec3Constant(obj gdc.ObjectPtr) *VisualShaderNodeVec3Constant {
  return &VisualShaderNodeVec3Constant{
    obj: obj,
  }
}

func (me *VisualShaderNodeVec3Constant) BaseClass() string {
  return "VisualShaderNodeVec3Constant"
}
