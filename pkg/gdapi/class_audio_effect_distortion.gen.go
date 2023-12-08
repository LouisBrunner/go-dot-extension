// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectDistortion struct {
  obj gdc.ObjectPtr
}

func createAudioEffectDistortion(obj gdc.ObjectPtr) *AudioEffectDistortion {
  return &AudioEffectDistortion{
    obj: obj,
  }
}

func (me *AudioEffectDistortion) BaseClass() string {
  return "AudioEffectDistortion"
}
