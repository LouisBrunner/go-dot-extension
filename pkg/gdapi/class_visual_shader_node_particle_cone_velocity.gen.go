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

type VisualShaderNodeParticleConeVelocity struct {
  VisualShaderNode
}

func (me *VisualShaderNodeParticleConeVelocity) BaseClass() string {
  return "VisualShaderNodeParticleConeVelocity"
}

func NewVisualShaderNodeParticleConeVelocity() *VisualShaderNodeParticleConeVelocity {
  str := StringNameFromStr("VisualShaderNodeParticleConeVelocity") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeParticleConeVelocity{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeParticleConeVelocity) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParticleConeVelocity) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParticleConeVelocity) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
