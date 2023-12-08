// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeSample3D struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeSample3D(obj gdc.ObjectPtr) *VisualShaderNodeSample3D {
  return &VisualShaderNodeSample3D{
    obj: obj,
  }
}

func (me *VisualShaderNodeSample3D) BaseClass() string {
  return "VisualShaderNodeSample3D"
}
