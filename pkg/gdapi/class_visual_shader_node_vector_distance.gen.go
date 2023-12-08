// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVectorDistance struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVectorDistance(obj gdc.ObjectPtr) *VisualShaderNodeVectorDistance {
  return &VisualShaderNodeVectorDistance{
    obj: obj,
  }
}

func (me *VisualShaderNodeVectorDistance) BaseClass() string {
  return "VisualShaderNodeVectorDistance"
}
