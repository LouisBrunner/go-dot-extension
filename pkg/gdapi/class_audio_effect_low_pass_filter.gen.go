// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectLowPassFilter struct {
  obj gdc.ObjectPtr
}

func createAudioEffectLowPassFilter(obj gdc.ObjectPtr) *AudioEffectLowPassFilter {
  return &AudioEffectLowPassFilter{
    obj: obj,
  }
}

func (me *AudioEffectLowPassFilter) BaseClass() string {
  return "AudioEffectLowPassFilter"
}
