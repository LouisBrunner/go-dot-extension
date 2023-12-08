// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeProximityFade struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeProximityFade(obj gdc.ObjectPtr) *VisualShaderNodeProximityFade {
  return &VisualShaderNodeProximityFade{
    obj: obj,
  }
}

func (me *VisualShaderNodeProximityFade) BaseClass() string {
  return "VisualShaderNodeProximityFade"
}
