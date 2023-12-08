// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeCurveTexture struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeCurveTexture(obj gdc.ObjectPtr) *VisualShaderNodeCurveTexture {
  return &VisualShaderNodeCurveTexture{
    obj: obj,
  }
}

func (me *VisualShaderNodeCurveTexture) BaseClass() string {
  return "VisualShaderNodeCurveTexture"
}
