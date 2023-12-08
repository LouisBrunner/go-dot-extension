// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeLinearSceneDepth struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeLinearSceneDepth(obj gdc.ObjectPtr) *VisualShaderNodeLinearSceneDepth {
  return &VisualShaderNodeLinearSceneDepth{
    obj: obj,
  }
}

func (me *VisualShaderNodeLinearSceneDepth) BaseClass() string {
  return "VisualShaderNodeLinearSceneDepth"
}
