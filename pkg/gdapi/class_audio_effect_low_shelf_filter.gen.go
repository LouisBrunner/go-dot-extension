// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectLowShelfFilter struct {
  AudioEffectFilter
}

func (me *AudioEffectLowShelfFilter) BaseClass() string {
  return "AudioEffectLowShelfFilter"
}



// Enums

func (me *AudioEffectLowShelfFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectLowShelfFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectLowShelfFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
