// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectStereoEnhance struct {
  obj gdc.ObjectPtr
}

func createAudioEffectStereoEnhance(obj gdc.ObjectPtr) *AudioEffectStereoEnhance {
  return &AudioEffectStereoEnhance{
    obj: obj,
  }
}

func (me *AudioEffectStereoEnhance) BaseClass() string {
  return "AudioEffectStereoEnhance"
}
