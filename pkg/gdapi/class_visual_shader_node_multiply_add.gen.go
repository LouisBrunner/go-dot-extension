// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeMultiplyAdd struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeMultiplyAdd(obj gdc.ObjectPtr) *VisualShaderNodeMultiplyAdd {
  return &VisualShaderNodeMultiplyAdd{
    obj: obj,
  }
}

func (me *VisualShaderNodeMultiplyAdd) BaseClass() string {
  return "VisualShaderNodeMultiplyAdd"
}
