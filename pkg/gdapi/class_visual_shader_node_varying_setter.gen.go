// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVaryingSetter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVaryingSetter(obj gdc.ObjectPtr) *VisualShaderNodeVaryingSetter {
  return &VisualShaderNodeVaryingSetter{
    obj: obj,
  }
}

func (me *VisualShaderNodeVaryingSetter) BaseClass() string {
  return "VisualShaderNodeVaryingSetter"
}
