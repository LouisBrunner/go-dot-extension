// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectBandPassFilter struct {
  obj gdc.ObjectPtr
}

func createAudioEffectBandPassFilter(obj gdc.ObjectPtr) *AudioEffectBandPassFilter {
  return &AudioEffectBandPassFilter{
    obj: obj,
  }
}

func (me *AudioEffectBandPassFilter) BaseClass() string {
  return "AudioEffectBandPassFilter"
}
