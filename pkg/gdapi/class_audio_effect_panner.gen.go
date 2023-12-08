// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectPanner struct {
  obj gdc.ObjectPtr
}

func createAudioEffectPanner(obj gdc.ObjectPtr) *AudioEffectPanner {
  return &AudioEffectPanner{
    obj: obj,
  }
}

func (me *AudioEffectPanner) BaseClass() string {
  return "AudioEffectPanner"
}
