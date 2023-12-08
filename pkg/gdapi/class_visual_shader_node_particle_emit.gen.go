// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleEmit struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParticleEmit) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParticleEmit) BaseClass() string {
  return "VisualShaderNodeParticleEmit"
}

type VisualShaderNodeParticleEmitEmitFlags int
const (
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagPosition VisualShaderNodeParticleEmitEmitFlags = 1
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagRotScale VisualShaderNodeParticleEmitEmitFlags = 2
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagVelocity VisualShaderNodeParticleEmitEmitFlags = 4
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagColor VisualShaderNodeParticleEmitEmitFlags = 8
  VisualShaderNodeParticleEmitEmitFlagsEmitFlagCustom VisualShaderNodeParticleEmitEmitFlags = 16
)

func  (me *VisualShaderNodeParticleEmit) SetFlags(flags VisualShaderNodeParticleEmitEmitFlags, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeParticleEmit) GetFlags() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
