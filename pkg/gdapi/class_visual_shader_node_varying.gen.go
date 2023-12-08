// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVarying struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVarying(obj gdc.ObjectPtr) *VisualShaderNodeVarying {
  return &VisualShaderNodeVarying{
    obj: obj,
  }
}

func (me *VisualShaderNodeVarying) BaseClass() string {
  return "VisualShaderNodeVarying"
}
