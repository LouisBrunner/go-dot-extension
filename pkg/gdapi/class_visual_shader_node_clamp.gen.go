// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeClamp struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeClamp(obj gdc.ObjectPtr) *VisualShaderNodeClamp {
  return &VisualShaderNodeClamp{
    obj: obj,
  }
}

func (me *VisualShaderNodeClamp) BaseClass() string {
  return "VisualShaderNodeClamp"
}
