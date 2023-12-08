// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectHighShelfFilter struct {
  obj gdc.ObjectPtr
}

func createAudioEffectHighShelfFilter(obj gdc.ObjectPtr) *AudioEffectHighShelfFilter {
  return &AudioEffectHighShelfFilter{
    obj: obj,
  }
}

func (me *AudioEffectHighShelfFilter) BaseClass() string {
  return "AudioEffectHighShelfFilter"
}
