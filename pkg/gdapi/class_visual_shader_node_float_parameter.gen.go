// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFloatParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeFloatParameter(obj gdc.ObjectPtr) *VisualShaderNodeFloatParameter {
  return &VisualShaderNodeFloatParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeFloatParameter) BaseClass() string {
  return "VisualShaderNodeFloatParameter"
}
