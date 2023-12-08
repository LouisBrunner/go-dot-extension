// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeMix struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeMix(obj gdc.ObjectPtr) *VisualShaderNodeMix {
  return &VisualShaderNodeMix{
    obj: obj,
  }
}

func (me *VisualShaderNodeMix) BaseClass() string {
  return "VisualShaderNodeMix"
}
