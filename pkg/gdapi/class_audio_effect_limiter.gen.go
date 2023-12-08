// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectLimiter struct {
  obj gdc.ObjectPtr
}

func createAudioEffectLimiter(obj gdc.ObjectPtr) *AudioEffectLimiter {
  return &AudioEffectLimiter{
    obj: obj,
  }
}

func (me *AudioEffectLimiter) BaseClass() string {
  return "AudioEffectLimiter"
}
