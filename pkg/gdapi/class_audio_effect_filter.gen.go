// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectFilter struct {
  obj gdc.ObjectPtr
}

func createAudioEffectFilter(obj gdc.ObjectPtr) *AudioEffectFilter {
  return &AudioEffectFilter{
    obj: obj,
  }
}

func (me *AudioEffectFilter) BaseClass() string {
  return "AudioEffectFilter"
}
