// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTexture3D struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTexture3D(obj gdc.ObjectPtr) *VisualShaderNodeTexture3D {
  return &VisualShaderNodeTexture3D{
    obj: obj,
  }
}

func (me *VisualShaderNodeTexture3D) BaseClass() string {
  return "VisualShaderNodeTexture3D"
}
