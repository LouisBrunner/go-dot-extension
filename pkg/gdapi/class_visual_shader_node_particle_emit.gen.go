// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  VisualShaderNodeParticleEmitEmitFlagPosition VisualShaderNodeParticleEmitEmitFlags = 1
  VisualShaderNodeParticleEmitEmitFlagRotScale VisualShaderNodeParticleEmitEmitFlags = 2
  VisualShaderNodeParticleEmitEmitFlagVelocity VisualShaderNodeParticleEmitEmitFlags = 4
  VisualShaderNodeParticleEmitEmitFlagColor VisualShaderNodeParticleEmitEmitFlags = 8
  VisualShaderNodeParticleEmitEmitFlagCustom VisualShaderNodeParticleEmitEmitFlags = 16
)
