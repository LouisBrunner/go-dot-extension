// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFloatFunc struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeFloatFunc(obj gdc.ObjectPtr) *VisualShaderNodeFloatFunc {
  return &VisualShaderNodeFloatFunc{
    obj: obj,
  }
}

func (me *VisualShaderNodeFloatFunc) BaseClass() string {
  return "VisualShaderNodeFloatFunc"
}
