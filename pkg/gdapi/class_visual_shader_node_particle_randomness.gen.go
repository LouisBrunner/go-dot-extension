// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeParticleRandomness struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeParticleRandomness) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeParticleRandomness) BaseClass() string {
  return "VisualShaderNodeParticleRandomness"
}

type VisualShaderNodeParticleRandomnessOpType int
const (
  VisualShaderNodeParticleRandomnessOpTypeScalar VisualShaderNodeParticleRandomnessOpType = 0
  VisualShaderNodeParticleRandomnessOpTypeVector2D VisualShaderNodeParticleRandomnessOpType = 1
  VisualShaderNodeParticleRandomnessOpTypeVector3D VisualShaderNodeParticleRandomnessOpType = 2
  VisualShaderNodeParticleRandomnessOpTypeVector4D VisualShaderNodeParticleRandomnessOpType = 3
  VisualShaderNodeParticleRandomnessOpTypeMax VisualShaderNodeParticleRandomnessOpType = 4
)
