// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleRandomness struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleRandomness(obj gdc.ObjectPtr) *VisualShaderNodeParticleRandomness {
  return &VisualShaderNodeParticleRandomness{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleRandomness) BaseClass() string {
  return "VisualShaderNodeParticleRandomness"
}
