// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParameterRef struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParameterRef(obj gdc.ObjectPtr) *VisualShaderNodeParameterRef {
  return &VisualShaderNodeParameterRef{
    obj: obj,
  }
}

func (me *VisualShaderNodeParameterRef) BaseClass() string {
  return "VisualShaderNodeParameterRef"
}
