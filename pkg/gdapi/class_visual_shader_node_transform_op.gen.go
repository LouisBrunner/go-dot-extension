// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformOp struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTransformOp(obj gdc.ObjectPtr) *VisualShaderNodeTransformOp {
  return &VisualShaderNodeTransformOp{
    obj: obj,
  }
}

func (me *VisualShaderNodeTransformOp) BaseClass() string {
  return "VisualShaderNodeTransformOp"
}
