// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeColorOp struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeColorOp(obj gdc.ObjectPtr) *VisualShaderNodeColorOp {
  return &VisualShaderNodeColorOp{
    obj: obj,
  }
}

func (me *VisualShaderNodeColorOp) BaseClass() string {
  return "VisualShaderNodeColorOp"
}
