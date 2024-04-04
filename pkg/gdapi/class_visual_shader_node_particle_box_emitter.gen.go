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

type VisualShaderNodeParticleBoxEmitter struct {
  VisualShaderNodeParticleEmitter
}

func (me *VisualShaderNodeParticleBoxEmitter) BaseClass() string {
  return "VisualShaderNodeParticleBoxEmitter"
}

func NewVisualShaderNodeParticleBoxEmitter() *VisualShaderNodeParticleBoxEmitter {
  str := StringNameFromStr("VisualShaderNodeParticleBoxEmitter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeParticleBoxEmitter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeParticleBoxEmitter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleBoxEmitter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleBoxEmitter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
