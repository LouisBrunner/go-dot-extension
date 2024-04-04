// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeParticleRingEmitter struct {
  VisualShaderNodeParticleEmitter
}

func (me *VisualShaderNodeParticleRingEmitter) BaseClass() string {
  return "VisualShaderNodeParticleRingEmitter"
}

func NewVisualShaderNodeParticleRingEmitter() *VisualShaderNodeParticleRingEmitter {
  str := StringNameFromStr("VisualShaderNodeParticleRingEmitter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeParticleRingEmitter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeParticleRingEmitter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleRingEmitter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleRingEmitter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
