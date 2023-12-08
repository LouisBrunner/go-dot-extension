// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeFaceForward struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeFaceForward(obj gdc.ObjectPtr) *VisualShaderNodeFaceForward {
  return &VisualShaderNodeFaceForward{
    obj: obj,
  }
}

func (me *VisualShaderNodeFaceForward) BaseClass() string {
  return "VisualShaderNodeFaceForward"
}
