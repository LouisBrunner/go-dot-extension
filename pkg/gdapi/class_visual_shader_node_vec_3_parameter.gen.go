// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVec3Parameter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVec3Parameter(obj gdc.ObjectPtr) *VisualShaderNodeVec3Parameter {
  return &VisualShaderNodeVec3Parameter{
    obj: obj,
  }
}

func (me *VisualShaderNodeVec3Parameter) BaseClass() string {
  return "VisualShaderNodeVec3Parameter"
}
