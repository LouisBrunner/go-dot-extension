// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeCustom struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeCustom(obj gdc.ObjectPtr) *VisualShaderNodeCustom {
  return &VisualShaderNodeCustom{
    obj: obj,
  }
}

func (me *VisualShaderNodeCustom) BaseClass() string {
  return "VisualShaderNodeCustom"
}
