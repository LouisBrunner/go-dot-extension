// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFloatConstant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeFloatConstant(obj gdc.ObjectPtr) *VisualShaderNodeFloatConstant {
  return &VisualShaderNodeFloatConstant{
    obj: obj,
  }
}

func (me *VisualShaderNodeFloatConstant) BaseClass() string {
  return "VisualShaderNodeFloatConstant"
}
