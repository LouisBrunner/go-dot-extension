// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectNotchFilter struct {
  obj gdc.ObjectPtr
}

func createAudioEffectNotchFilter(obj gdc.ObjectPtr) *AudioEffectNotchFilter {
  return &AudioEffectNotchFilter{
    obj: obj,
  }
}

func (me *AudioEffectNotchFilter) BaseClass() string {
  return "AudioEffectNotchFilter"
}
