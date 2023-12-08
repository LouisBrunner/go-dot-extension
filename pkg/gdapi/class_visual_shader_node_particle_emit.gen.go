// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleEmit struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleEmit(obj gdc.ObjectPtr) *VisualShaderNodeParticleEmit {
  return &VisualShaderNodeParticleEmit{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleEmit) BaseClass() string {
  return "VisualShaderNodeParticleEmit"
}
