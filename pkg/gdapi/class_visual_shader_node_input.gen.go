// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeInput struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeInput(obj gdc.ObjectPtr) *VisualShaderNodeInput {
  return &VisualShaderNodeInput{
    obj: obj,
  }
}

func (me *VisualShaderNodeInput) BaseClass() string {
  return "VisualShaderNodeInput"
}
