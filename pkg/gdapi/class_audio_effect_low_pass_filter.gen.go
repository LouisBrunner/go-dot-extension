// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectLowPassFilter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectLowPassFilter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectLowPassFilter) BaseClass() string {
  return "AudioEffectLowPassFilter"
}



// Enums

func (me *AudioEffectLowPassFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectLowPassFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectLowPassFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties