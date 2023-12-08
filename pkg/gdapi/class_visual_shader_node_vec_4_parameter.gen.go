// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVec4Parameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVec4Parameter(obj gdc.ObjectPtr) *VisualShaderNodeVec4Parameter {
  return &VisualShaderNodeVec4Parameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeVec4Parameter) BaseClass() string {
  return "VisualShaderNodeVec4Parameter"
}
