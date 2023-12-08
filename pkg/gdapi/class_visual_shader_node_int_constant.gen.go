// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIntConstant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeIntConstant(obj gdc.ObjectPtr) *VisualShaderNodeIntConstant {
  return &VisualShaderNodeIntConstant{
    obj: obj,
  }
}

func (me *VisualShaderNodeIntConstant) BaseClass() string {
  return "VisualShaderNodeIntConstant"
}
