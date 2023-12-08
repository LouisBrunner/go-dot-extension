// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeUIntConstant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeUIntConstant(obj gdc.ObjectPtr) *VisualShaderNodeUIntConstant {
  return &VisualShaderNodeUIntConstant{
    obj: obj,
  }
}

func (me *VisualShaderNodeUIntConstant) BaseClass() string {
  return "VisualShaderNodeUIntConstant"
}
