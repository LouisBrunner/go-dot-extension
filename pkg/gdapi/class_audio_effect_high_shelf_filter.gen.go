// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectHighShelfFilter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectHighShelfFilter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectHighShelfFilter) BaseClass() string {
  return "AudioEffectHighShelfFilter"
}



// Enums

func (me *AudioEffectHighShelfFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectHighShelfFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectHighShelfFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties