// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectBandLimitFilter struct {
  AudioEffectFilter
}

func (me *AudioEffectBandLimitFilter) BaseClass() string {
  return "AudioEffectBandLimitFilter"
}



// Enums

func (me *AudioEffectBandLimitFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectBandLimitFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectBandLimitFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
