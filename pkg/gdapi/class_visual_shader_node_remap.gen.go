// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeRemap struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeRemap(obj gdc.ObjectPtr) *VisualShaderNodeRemap {
  return &VisualShaderNodeRemap{
    obj: obj,
  }
}

func (me *VisualShaderNodeRemap) BaseClass() string {
  return "VisualShaderNodeRemap"
}
