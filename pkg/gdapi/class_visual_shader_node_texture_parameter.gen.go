// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTextureParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTextureParameter(obj gdc.ObjectPtr) *VisualShaderNodeTextureParameter {
  return &VisualShaderNodeTextureParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeTextureParameter) BaseClass() string {
  return "VisualShaderNodeTextureParameter"
}
