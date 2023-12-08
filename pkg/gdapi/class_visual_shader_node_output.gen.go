// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeOutput struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeOutput(obj gdc.ObjectPtr) *VisualShaderNodeOutput {
  return &VisualShaderNodeOutput{
    obj: obj,
  }
}

func (me *VisualShaderNodeOutput) BaseClass() string {
  return "VisualShaderNodeOutput"
}
