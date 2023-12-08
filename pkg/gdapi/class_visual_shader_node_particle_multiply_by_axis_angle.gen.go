// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleMultiplyByAxisAngle struct {
  obj gdc.ObjectPtr
}

func createVisualShaderNodeParticleMultiplyByAxisAngle(obj gdc.ObjectPtr) *VisualShaderNodeParticleMultiplyByAxisAngle {
  return &VisualShaderNodeParticleMultiplyByAxisAngle{
    obj: obj,
  }
}

func (me *VisualShaderNodeParticleMultiplyByAxisAngle) BaseClass() string {
  return "VisualShaderNodeParticleMultiplyByAxisAngle"
}
