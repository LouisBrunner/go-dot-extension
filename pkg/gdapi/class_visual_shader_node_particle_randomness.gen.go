// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeParticleRandomnessOpTypeOpTypeScalar VisualShaderNodeParticleRandomnessOpType = 0
  VisualShaderNodeParticleRandomnessOpTypeOpTypeVector2D VisualShaderNodeParticleRandomnessOpType = 1
  VisualShaderNodeParticleRandomnessOpTypeOpTypeVector3D VisualShaderNodeParticleRandomnessOpType = 2
  VisualShaderNodeParticleRandomnessOpTypeOpTypeVector4D VisualShaderNodeParticleRandomnessOpType = 3
  VisualShaderNodeParticleRandomnessOpTypeOpTypeMax VisualShaderNodeParticleRandomnessOpType = 4
)

func  (me *VisualShaderNodeParticleRandomness) SetOpType(type_ VisualShaderNodeParticleRandomnessOpType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeParticleRandomness) GetOpType() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
