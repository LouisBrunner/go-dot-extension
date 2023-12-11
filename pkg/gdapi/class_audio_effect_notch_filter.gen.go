// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectNotchFilter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectNotchFilter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectNotchFilter) BaseClass() string {
  return "AudioEffectNotchFilter"
}



// Enums

func (me *AudioEffectNotchFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectNotchFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectNotchFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
