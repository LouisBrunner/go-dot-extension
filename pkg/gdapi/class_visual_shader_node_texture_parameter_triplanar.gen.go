// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTextureParameterTriplanar struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTextureParameterTriplanar(obj gdc.ObjectPtr) *VisualShaderNodeTextureParameterTriplanar {
  return &VisualShaderNodeTextureParameterTriplanar{
    obj: obj,
  }
}

func (me *VisualShaderNodeTextureParameterTriplanar) BaseClass() string {
  return "VisualShaderNodeTextureParameterTriplanar"
}
