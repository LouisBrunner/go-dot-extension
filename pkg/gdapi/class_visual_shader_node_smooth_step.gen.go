// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeSmoothStep struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeSmoothStep(obj gdc.ObjectPtr) *VisualShaderNodeSmoothStep {
  return &VisualShaderNodeSmoothStep{
    obj: obj,
  }
}

func (me *VisualShaderNodeSmoothStep) BaseClass() string {
  return "VisualShaderNodeSmoothStep"
}