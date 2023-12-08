// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleAccelerator struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleAccelerator(obj gdc.ObjectPtr) *VisualShaderNodeParticleAccelerator {
  return &VisualShaderNodeParticleAccelerator{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleAccelerator) BaseClass() string {
  return "VisualShaderNodeParticleAccelerator"
}
