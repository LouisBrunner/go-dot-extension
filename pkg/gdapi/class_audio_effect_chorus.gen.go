// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectChorus struct {
  obj gdc.ObjectPtr
}

func createAudioEffectChorus(obj gdc.ObjectPtr) *AudioEffectChorus {
  return &AudioEffectChorus{
    obj: obj,
  }
}

func (me *AudioEffectChorus) BaseClass() string {
  return "AudioEffectChorus"
}
