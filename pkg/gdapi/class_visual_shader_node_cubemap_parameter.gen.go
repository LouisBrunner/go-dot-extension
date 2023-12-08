// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeCubemapParameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeCubemapParameter(obj gdc.ObjectPtr) *VisualShaderNodeCubemapParameter {
  return &VisualShaderNodeCubemapParameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeCubemapParameter) BaseClass() string {
  return "VisualShaderNodeCubemapParameter"
}
