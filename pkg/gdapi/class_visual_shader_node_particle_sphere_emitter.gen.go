// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForVisualShaderNodeParticleSphereEmitterList struct {
}

var ptrsForVisualShaderNodeParticleSphereEmitter ptrsForVisualShaderNodeParticleSphereEmitterList

func initVisualShaderNodeParticleSphereEmitterPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeParticleSphereEmitter")
  defer className.Destroy()
}

type VisualShaderNodeParticleSphereEmitter struct {
  VisualShaderNodeParticleEmitter
}

func (me *VisualShaderNodeParticleSphereEmitter) BaseClass() string {
  return "VisualShaderNodeParticleSphereEmitter"
}

func NewVisualShaderNodeParticleSphereEmitter() *VisualShaderNodeParticleSphereEmitter {
  str := StringNameFromStr("VisualShaderNodeParticleSphereEmitter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeParticleSphereEmitter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeParticleSphereEmitter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleSphereEmitter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleSphereEmitter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
