// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeRandomRange struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeRandomRange(obj gdc.ObjectPtr) *VisualShaderNodeRandomRange {
  return &VisualShaderNodeRandomRange{
    obj: obj,
  }
}

func (me *VisualShaderNodeRandomRange) BaseClass() string {
  return "VisualShaderNodeRandomRange"
}
