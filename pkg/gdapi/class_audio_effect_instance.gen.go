// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectInstance struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectInstance) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectInstance) BaseClass() string {
  return "AudioEffectInstance"
}



// Enums

func (me *AudioEffectInstance) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectInstance) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectInstance) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
