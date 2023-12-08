// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeStep struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeStep(obj gdc.ObjectPtr) *VisualShaderNodeStep {
  return &VisualShaderNodeStep{
    obj: obj,
  }
}

func (me *VisualShaderNodeStep) BaseClass() string {
  return "VisualShaderNodeStep"
}
