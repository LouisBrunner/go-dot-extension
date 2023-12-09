// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectEQ struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectEQ) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectEQ) BaseClass() string {
  return "AudioEffectEQ"
}



// Enums

func (me *AudioEffectEQ) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectEQ) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectEQ) SetBandGainDb(band_idx int, volume_db float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectEQ) GetBandGainDb(band_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectEQ) GetBandCount()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
