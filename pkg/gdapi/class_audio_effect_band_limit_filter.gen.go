// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectBandLimitFilter struct {
  obj gdc.ObjectPtr
}

func createAudioEffectBandLimitFilter(obj gdc.ObjectPtr) *AudioEffectBandLimitFilter {
  return &AudioEffectBandLimitFilter{
    obj: obj,
  }
}

func (me *AudioEffectBandLimitFilter) BaseClass() string {
  return "AudioEffectBandLimitFilter"
}
