// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVaryingGetter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVaryingGetter(obj gdc.ObjectPtr) *VisualShaderNodeVaryingGetter {
  return &VisualShaderNodeVaryingGetter{
    obj: obj,
  }
}

func (me *VisualShaderNodeVaryingGetter) BaseClass() string {
  return "VisualShaderNodeVaryingGetter"
}
