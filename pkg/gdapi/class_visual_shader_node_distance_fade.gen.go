// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeDistanceFade struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeDistanceFade(obj gdc.ObjectPtr) *VisualShaderNodeDistanceFade {
  return &VisualShaderNodeDistanceFade{
    obj: obj,
  }
}

func (me *VisualShaderNodeDistanceFade) BaseClass() string {
  return "VisualShaderNodeDistanceFade"
}
