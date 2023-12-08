// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeSDFToScreenUV struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeSDFToScreenUV(obj gdc.ObjectPtr) *VisualShaderNodeSDFToScreenUV {
  return &VisualShaderNodeSDFToScreenUV{
    obj: obj,
  }
}

func (me *VisualShaderNodeSDFToScreenUV) BaseClass() string {
  return "VisualShaderNodeSDFToScreenUV"
}
