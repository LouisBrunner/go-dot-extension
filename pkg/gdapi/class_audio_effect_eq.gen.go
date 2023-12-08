// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectEQ struct {
  obj gdc.ObjectPtr
}

func createAudioEffectEQ(obj gdc.ObjectPtr) *AudioEffectEQ {
  return &AudioEffectEQ{
    obj: obj,
  }
}

func (me *AudioEffectEQ) BaseClass() string {
  return "AudioEffectEQ"
}
