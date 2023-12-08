// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeGroupBase struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeGroupBase(obj gdc.ObjectPtr) *VisualShaderNodeGroupBase {
  return &VisualShaderNodeGroupBase{
    obj: obj,
  }
}

func (me *VisualShaderNodeGroupBase) BaseClass() string {
  return "VisualShaderNodeGroupBase"
}
