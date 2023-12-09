// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectBandLimitFilter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectBandLimitFilter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

// TODO: properties (class)

// TODO: signals (class)
