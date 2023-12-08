// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTransformCompose struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeTransformCompose(obj gdc.ObjectPtr) *VisualShaderNodeTransformCompose {
  return &VisualShaderNodeTransformCompose{
    obj: obj,
  }
}

func (me *VisualShaderNodeTransformCompose) BaseClass() string {
  return "VisualShaderNodeTransformCompose"
}
