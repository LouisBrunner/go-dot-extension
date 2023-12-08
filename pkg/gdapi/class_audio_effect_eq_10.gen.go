// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectEQ10 struct {
  obj gdc.ObjectPtr
}

func createAudioEffectEQ10(obj gdc.ObjectPtr) *AudioEffectEQ10 {
  return &AudioEffectEQ10{
    obj: obj,
  }
}

func (me *AudioEffectEQ10) BaseClass() string {
  return "AudioEffectEQ10"
}
