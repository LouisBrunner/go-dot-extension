// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectInstance struct {
  obj gdc.ObjectPtr
}

func createAudioEffectInstance(obj gdc.ObjectPtr) *AudioEffectInstance {
  return &AudioEffectInstance{
    obj: obj,
  }
}

func (me *AudioEffectInstance) BaseClass() string {
  return "AudioEffectInstance"
}
