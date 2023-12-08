// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTexture2DArrayParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTexture2DArrayParameter(obj gdc.ObjectPtr) *VisualShaderNodeTexture2DArrayParameter {
  return &VisualShaderNodeTexture2DArrayParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeTexture2DArrayParameter) BaseClass() string {
  return "VisualShaderNodeTexture2DArrayParameter"
}
