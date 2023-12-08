// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeCubemap struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeCubemap(obj gdc.ObjectPtr) *VisualShaderNodeCubemap {
  return &VisualShaderNodeCubemap{
    obj: obj,
  }
}

func (me *VisualShaderNodeCubemap) BaseClass() string {
  return "VisualShaderNodeCubemap"
}
