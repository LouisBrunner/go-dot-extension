// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVectorLen struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVectorLen(obj gdc.ObjectPtr) *VisualShaderNodeVectorLen {
  return &VisualShaderNodeVectorLen{
    obj: obj,
  }
}

func (me *VisualShaderNodeVectorLen) BaseClass() string {
  return "VisualShaderNodeVectorLen"
}
