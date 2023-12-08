// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeBillboard struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeBillboard(obj gdc.ObjectPtr) *VisualShaderNodeBillboard {
  return &VisualShaderNodeBillboard{
    obj: obj,
  }
}

func (me *VisualShaderNodeBillboard) BaseClass() string {
  return "VisualShaderNodeBillboard"
}
