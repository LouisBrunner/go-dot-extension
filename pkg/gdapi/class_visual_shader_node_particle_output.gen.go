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

type ptrsForVisualShaderNodeParticleOutputList struct {
}

var ptrsForVisualShaderNodeParticleOutput ptrsForVisualShaderNodeParticleOutputList

func initVisualShaderNodeParticleOutputPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeParticleOutput")
  defer className.Destroy()
}

type VisualShaderNodeParticleOutput struct {
  VisualShaderNodeOutput
}

func (me *VisualShaderNodeParticleOutput) BaseClass() string {
  return "VisualShaderNodeParticleOutput"
}

func NewVisualShaderNodeParticleOutput() *VisualShaderNodeParticleOutput {
  str := StringNameFromStr("VisualShaderNodeParticleOutput") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeParticleOutput{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeParticleOutput) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleOutput) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleOutput) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
