// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectPitchShift struct {
  obj gdc.ObjectPtr
}

func createAudioEffectPitchShift(obj gdc.ObjectPtr) *AudioEffectPitchShift {
  return &AudioEffectPitchShift{
    obj: obj,
  }
}

func (me *AudioEffectPitchShift) BaseClass() string {
  return "AudioEffectPitchShift"
}
