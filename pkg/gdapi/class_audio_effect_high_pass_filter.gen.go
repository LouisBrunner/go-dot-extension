// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectHighPassFilter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectHighPassFilter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectHighPassFilter) BaseClass() string {
  return "AudioEffectHighPassFilter"
}



// Enums

func (me *AudioEffectHighPassFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectHighPassFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectHighPassFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties