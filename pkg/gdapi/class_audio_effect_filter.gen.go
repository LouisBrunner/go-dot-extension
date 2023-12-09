// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectFilter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectFilter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectFilter) BaseClass() string {
  return "AudioEffectFilter"
}



// Enums

type AudioEffectFilterFilterDB int
const (
  AudioEffectFilterFilterDBFilter6Db AudioEffectFilterFilterDB = 0
  AudioEffectFilterFilterDBFilter12Db AudioEffectFilterFilterDB = 1
  AudioEffectFilterFilterDBFilter18Db AudioEffectFilterFilterDB = 2
  AudioEffectFilterFilterDBFilter24Db AudioEffectFilterFilterDB = 3
)

func (me *AudioEffectFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectFilter) SetCutoff(freq float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectFilter) GetCutoff()  {
  panic("TODO: implement")
}

func  (me *AudioEffectFilter) SetResonance(amount float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectFilter) GetResonance()  {
  panic("TODO: implement")
}

func  (me *AudioEffectFilter) SetGain(amount float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectFilter) GetGain()  {
  panic("TODO: implement")
}

func  (me *AudioEffectFilter) SetDb(amount AudioEffectFilterFilterDB, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectFilter) GetDb()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
