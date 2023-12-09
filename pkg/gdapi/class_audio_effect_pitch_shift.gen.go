// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectPitchShift struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectPitchShift) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectPitchShift) BaseClass() string {
  return "AudioEffectPitchShift"
}



// Enums

type AudioEffectPitchShiftFFTSize int
const (
  AudioEffectPitchShiftFFTSizeFftSize256 AudioEffectPitchShiftFFTSize = 0
  AudioEffectPitchShiftFFTSizeFftSize512 AudioEffectPitchShiftFFTSize = 1
  AudioEffectPitchShiftFFTSizeFftSize1024 AudioEffectPitchShiftFFTSize = 2
  AudioEffectPitchShiftFFTSizeFftSize2048 AudioEffectPitchShiftFFTSize = 3
  AudioEffectPitchShiftFFTSizeFftSize4096 AudioEffectPitchShiftFFTSize = 4
  AudioEffectPitchShiftFFTSizeFftSizeMax AudioEffectPitchShiftFFTSize = 5
)

func (me *AudioEffectPitchShift) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectPitchShift) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectPitchShift) SetPitchScale(rate float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectPitchShift) GetPitchScale()  {
  panic("TODO: implement")
}

func  (me *AudioEffectPitchShift) SetOversampling(amount int, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectPitchShift) GetOversampling()  {
  panic("TODO: implement")
}

func  (me *AudioEffectPitchShift) SetFftSize(size AudioEffectPitchShiftFFTSize, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectPitchShift) GetFftSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
