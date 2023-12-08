// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFresnel struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeFresnel(obj gdc.ObjectPtr) *VisualShaderNodeFresnel {
  return &VisualShaderNodeFresnel{
    obj: obj,
  }
}

func (me *VisualShaderNodeFresnel) BaseClass() string {
  return "VisualShaderNodeFresnel"
}
