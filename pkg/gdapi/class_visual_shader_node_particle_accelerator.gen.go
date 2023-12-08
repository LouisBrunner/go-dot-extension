// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeParticleAcceleratorModeModeLinear VisualShaderNodeParticleAcceleratorMode = 0
  VisualShaderNodeParticleAcceleratorModeModeRadial VisualShaderNodeParticleAcceleratorMode = 1
  VisualShaderNodeParticleAcceleratorModeModeTangential VisualShaderNodeParticleAcceleratorMode = 2
  VisualShaderNodeParticleAcceleratorModeModeMax VisualShaderNodeParticleAcceleratorMode = 3
)

func  (me *VisualShaderNodeParticleAccelerator) SetMode(mode VisualShaderNodeParticleAcceleratorMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeParticleAccelerator) GetMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
