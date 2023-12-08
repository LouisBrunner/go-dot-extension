// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleOutput struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleOutput(obj gdc.ObjectPtr) *VisualShaderNodeParticleOutput {
  return &VisualShaderNodeParticleOutput{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleOutput) BaseClass() string {
  return "VisualShaderNodeParticleOutput"
}
