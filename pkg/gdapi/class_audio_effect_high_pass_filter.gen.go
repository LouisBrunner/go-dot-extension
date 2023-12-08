// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectHighPassFilter struct {
  obj gdc.ObjectPtr
}

func createAudioEffectHighPassFilter(obj gdc.ObjectPtr) *AudioEffectHighPassFilter {
  return &AudioEffectHighPassFilter{
    obj: obj,
  }
}

func (me *AudioEffectHighPassFilter) BaseClass() string {
  return "AudioEffectHighPassFilter"
}
