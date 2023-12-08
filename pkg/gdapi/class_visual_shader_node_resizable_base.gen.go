// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeResizableBase struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeResizableBase(obj gdc.ObjectPtr) *VisualShaderNodeResizableBase {
  return &VisualShaderNodeResizableBase{
    obj: obj,
  }
}

func (me *VisualShaderNodeResizableBase) BaseClass() string {
  return "VisualShaderNodeResizableBase"
}
