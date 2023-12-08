// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type AudioEffectPitchShiftFFTSize int
const (
  AudioEffectPitchShiftFftSize256 AudioEffectPitchShiftFFTSize = 0
  AudioEffectPitchShiftFftSize512 AudioEffectPitchShiftFFTSize = 1
  AudioEffectPitchShiftFftSize1024 AudioEffectPitchShiftFFTSize = 2
  AudioEffectPitchShiftFftSize2048 AudioEffectPitchShiftFFTSize = 3
  AudioEffectPitchShiftFftSize4096 AudioEffectPitchShiftFFTSize = 4
  AudioEffectPitchShiftFftSizeMax AudioEffectPitchShiftFFTSize = 5
)
