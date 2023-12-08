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

type AudioEffectFilterFilterDB int
const (
  AudioEffectFilterFilterDBFilter6Db AudioEffectFilterFilterDB = 0
  AudioEffectFilterFilterDBFilter12Db AudioEffectFilterFilterDB = 1
  AudioEffectFilterFilterDBFilter18Db AudioEffectFilterFilterDB = 2
  AudioEffectFilterFilterDBFilter24Db AudioEffectFilterFilterDB = 3
)

func  (me *AudioEffectFilter) SetCutoff(freq float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectFilter) GetCutoff() { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectFilter) SetResonance(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectFilter) GetResonance() { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectFilter) SetGain(amount float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectFilter) GetGain() { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectFilter) SetDb(amount AudioEffectFilterFilterDB, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectFilter) GetDb() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
