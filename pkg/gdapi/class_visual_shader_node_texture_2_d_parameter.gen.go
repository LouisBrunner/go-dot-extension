// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTexture2DParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTexture2DParameter(obj gdc.ObjectPtr) *VisualShaderNodeTexture2DParameter {
  return &VisualShaderNodeTexture2DParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeTexture2DParameter) BaseClass() string {
  return "VisualShaderNodeTexture2DParameter"
}
