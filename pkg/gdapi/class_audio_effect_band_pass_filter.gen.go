// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectBandPassFilter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectBandPassFilter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectBandPassFilter) BaseClass() string {
  return "AudioEffectBandPassFilter"
}



// Enums

func (me *AudioEffectBandPassFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectBandPassFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectBandPassFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties