// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

// TODO: properties (class)

// TODO: signals (class)
