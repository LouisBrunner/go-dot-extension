// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeSwitch struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeSwitch(obj gdc.ObjectPtr) *VisualShaderNodeSwitch {
  return &VisualShaderNodeSwitch{
    obj: obj,
  }
}

func (me *VisualShaderNodeSwitch) BaseClass() string {
  return "VisualShaderNodeSwitch"
}
