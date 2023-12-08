// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVectorBase struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVectorBase(obj gdc.ObjectPtr) *VisualShaderNodeVectorBase {
  return &VisualShaderNodeVectorBase{
    obj: obj,
  }
}

func (me *VisualShaderNodeVectorBase) BaseClass() string {
  return "VisualShaderNodeVectorBase"
}
