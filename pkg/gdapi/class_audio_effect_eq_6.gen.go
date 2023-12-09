// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectEQ6 struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectEQ6) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectEQ6) BaseClass() string {
  return "AudioEffectEQ6"
}



// Enums

func (me *AudioEffectEQ6) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectEQ6) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ6) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
