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



// Enums

type VisualShaderNodeParticleRandomnessOpType int
const (
  VisualShaderNodeParticleRandomnessOpTypeOpTypeScalar VisualShaderNodeParticleRandomnessOpType = 0
  VisualShaderNodeParticleRandomnessOpTypeOpTypeVector2D VisualShaderNodeParticleRandomnessOpType = 1
  VisualShaderNodeParticleRandomnessOpTypeOpTypeVector3D VisualShaderNodeParticleRandomnessOpType = 2
  VisualShaderNodeParticleRandomnessOpTypeOpTypeVector4D VisualShaderNodeParticleRandomnessOpType = 3
  VisualShaderNodeParticleRandomnessOpTypeOpTypeMax VisualShaderNodeParticleRandomnessOpType = 4
)

func (me *VisualShaderNodeParticleRandomness) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleRandomness) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleRandomness) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeParticleRandomness) SetOpType(type_ VisualShaderNodeParticleRandomnessOpType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeParticleRandomness) GetOpType()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
