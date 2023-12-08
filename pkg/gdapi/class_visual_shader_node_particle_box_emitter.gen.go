// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleBoxEmitter struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleBoxEmitter(obj gdc.ObjectPtr) *VisualShaderNodeParticleBoxEmitter {
  return &VisualShaderNodeParticleBoxEmitter{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleBoxEmitter) BaseClass() string {
  return "VisualShaderNodeParticleBoxEmitter"
}
