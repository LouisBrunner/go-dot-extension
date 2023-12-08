// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleRingEmitter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleRingEmitter(obj gdc.ObjectPtr) *VisualShaderNodeParticleRingEmitter {
  return &VisualShaderNodeParticleRingEmitter{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleRingEmitter) BaseClass() string {
  return "VisualShaderNodeParticleRingEmitter"
}
