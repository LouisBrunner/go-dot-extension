// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleAccelerator struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParticleAccelerator) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParticleAccelerator) BaseClass() string {
  return "VisualShaderNodeParticleAccelerator"
}

type VisualShaderNodeParticleAcceleratorMode int
const (
  VisualShaderNodeParticleAcceleratorModeLinear VisualShaderNodeParticleAcceleratorMode = 0
  VisualShaderNodeParticleAcceleratorModeRadial VisualShaderNodeParticleAcceleratorMode = 1
  VisualShaderNodeParticleAcceleratorModeTangential VisualShaderNodeParticleAcceleratorMode = 2
  VisualShaderNodeParticleAcceleratorModeMax VisualShaderNodeParticleAcceleratorMode = 3
)
