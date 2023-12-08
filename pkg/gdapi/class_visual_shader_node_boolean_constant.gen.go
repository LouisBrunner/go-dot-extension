// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeBooleanConstant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeBooleanConstant(obj gdc.ObjectPtr) *VisualShaderNodeBooleanConstant {
  return &VisualShaderNodeBooleanConstant{
    obj: obj,
  }
}

func (me *VisualShaderNodeBooleanConstant) BaseClass() string {
  return "VisualShaderNodeBooleanConstant"
}
