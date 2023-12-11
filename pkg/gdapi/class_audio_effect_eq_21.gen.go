// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectEQ21 struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectEQ21) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectEQ21) BaseClass() string {
  return "AudioEffectEQ21"
}



// Enums

func (me *AudioEffectEQ21) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectEQ21) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ21) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
