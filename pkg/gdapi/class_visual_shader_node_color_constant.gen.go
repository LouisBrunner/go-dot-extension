// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeColorConstant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeColorConstant(obj gdc.ObjectPtr) *VisualShaderNodeColorConstant {
  return &VisualShaderNodeColorConstant{
    obj: obj,
  }
}

func (me *VisualShaderNodeColorConstant) BaseClass() string {
  return "VisualShaderNodeColorConstant"
}
