// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIf struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeIf(obj gdc.ObjectPtr) *VisualShaderNodeIf {
  return &VisualShaderNodeIf{
    obj: obj,
  }
}

func (me *VisualShaderNodeIf) BaseClass() string {
  return "VisualShaderNodeIf"
}
