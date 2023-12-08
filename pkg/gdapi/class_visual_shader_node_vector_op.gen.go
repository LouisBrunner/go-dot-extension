// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeVectorOp struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeVectorOp(obj gdc.ObjectPtr) *VisualShaderNodeVectorOp {
  return &VisualShaderNodeVectorOp{
    obj: obj,
  }
}

func (me *VisualShaderNodeVectorOp) BaseClass() string {
  return "VisualShaderNodeVectorOp"
}
