// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformConstant struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTransformConstant(obj gdc.ObjectPtr) *VisualShaderNodeTransformConstant {
  return &VisualShaderNodeTransformConstant{
    obj: obj,
  }
}

func (me *VisualShaderNodeTransformConstant) BaseClass() string {
  return "VisualShaderNodeTransformConstant"
}
