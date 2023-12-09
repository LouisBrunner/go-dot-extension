// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectLowShelfFilter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectLowShelfFilter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

// TODO: properties (class)

// TODO: signals (class)
