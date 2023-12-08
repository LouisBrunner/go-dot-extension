// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeIs struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeIs(obj gdc.ObjectPtr) *VisualShaderNodeIs {
  return &VisualShaderNodeIs{
    obj: obj,
  }
}

func (me *VisualShaderNodeIs) BaseClass() string {
  return "VisualShaderNodeIs"
}
