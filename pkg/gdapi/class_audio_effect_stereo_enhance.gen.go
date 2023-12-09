// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectStereoEnhance struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectStereoEnhance) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectStereoEnhance) BaseClass() string {
  return "AudioEffectStereoEnhance"
}



// Enums

func (me *AudioEffectStereoEnhance) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectStereoEnhance) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectStereoEnhance) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectStereoEnhance) SetPanPullout(amount float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectStereoEnhance) GetPanPullout()  {
  panic("TODO: implement")
}

func  (me *AudioEffectStereoEnhance) SetTimePullout(amount float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectStereoEnhance) GetTimePullout()  {
  panic("TODO: implement")
}

func  (me *AudioEffectStereoEnhance) SetSurround(amount float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectStereoEnhance) GetSurround()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
