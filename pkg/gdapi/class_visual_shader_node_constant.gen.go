// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeConstant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeConstant(obj gdc.ObjectPtr) *VisualShaderNodeConstant {
  return &VisualShaderNodeConstant{
    obj: obj,
  }
}

func (me *VisualShaderNodeConstant) BaseClass() string {
  return "VisualShaderNodeConstant"
}
