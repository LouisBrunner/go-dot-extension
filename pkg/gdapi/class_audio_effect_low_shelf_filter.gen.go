// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectLowShelfFilter struct {
  obj gdc.ObjectPtr
}

func createAudioEffectLowShelfFilter(obj gdc.ObjectPtr) *AudioEffectLowShelfFilter {
  return &AudioEffectLowShelfFilter{
    obj: obj,
  }
}

func (me *AudioEffectLowShelfFilter) BaseClass() string {
  return "AudioEffectLowShelfFilter"
}
