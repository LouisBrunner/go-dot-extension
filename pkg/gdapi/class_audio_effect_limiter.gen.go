// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectLimiter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectLimiter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectLimiter) BaseClass() string {
  return "AudioEffectLimiter"
}



// Enums

func (me *AudioEffectLimiter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectLimiter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectLimiter) SetCeilingDb(ceiling float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectLimiter) GetCeilingDb()  {
  panic("TODO: implement")
}

func  (me *AudioEffectLimiter) SetThresholdDb(threshold float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectLimiter) GetThresholdDb()  {
  panic("TODO: implement")
}

func  (me *AudioEffectLimiter) SetSoftClipDb(soft_clip float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectLimiter) GetSoftClipDb()  {
  panic("TODO: implement")
}

func  (me *AudioEffectLimiter) SetSoftClipRatio(soft_clip float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectLimiter) GetSoftClipRatio()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
