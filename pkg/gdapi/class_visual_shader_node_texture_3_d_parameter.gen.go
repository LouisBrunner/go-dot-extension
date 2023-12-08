// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTexture3DParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTexture3DParameter(obj gdc.ObjectPtr) *VisualShaderNodeTexture3DParameter {
  return &VisualShaderNodeTexture3DParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeTexture3DParameter) BaseClass() string {
  return "VisualShaderNodeTexture3DParameter"
}
