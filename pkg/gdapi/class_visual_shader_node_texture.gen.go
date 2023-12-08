// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTexture struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTexture(obj gdc.ObjectPtr) *VisualShaderNodeTexture {
  return &VisualShaderNodeTexture{
    obj: obj,
  }
}

func (me *VisualShaderNodeTexture) BaseClass() string {
  return "VisualShaderNodeTexture"
}
