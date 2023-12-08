// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFloatOp struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeFloatOp(obj gdc.ObjectPtr) *VisualShaderNodeFloatOp {
  return &VisualShaderNodeFloatOp{
    obj: obj,
  }
}

func (me *VisualShaderNodeFloatOp) BaseClass() string {
  return "VisualShaderNodeFloatOp"
}
