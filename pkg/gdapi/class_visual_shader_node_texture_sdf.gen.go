// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTextureSDF struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTextureSDF(obj gdc.ObjectPtr) *VisualShaderNodeTextureSDF {
  return &VisualShaderNodeTextureSDF{
    obj: obj,
  }
}

func (me *VisualShaderNodeTextureSDF) BaseClass() string {
  return "VisualShaderNodeTextureSDF"
}
