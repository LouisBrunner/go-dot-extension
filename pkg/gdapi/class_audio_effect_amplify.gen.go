// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectAmplify struct {
  obj gdc.ObjectPtr
}

func createAudioEffectAmplify(obj gdc.ObjectPtr) *AudioEffectAmplify {
  return &AudioEffectAmplify{
    obj: obj,
  }
}

func (me *AudioEffectAmplify) BaseClass() string {
  return "AudioEffectAmplify"
}
