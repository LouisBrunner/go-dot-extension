// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleEmitter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleEmitter(obj gdc.ObjectPtr) *VisualShaderNodeParticleEmitter {
  return &VisualShaderNodeParticleEmitter{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleEmitter) BaseClass() string {
  return "VisualShaderNodeParticleEmitter"
}
