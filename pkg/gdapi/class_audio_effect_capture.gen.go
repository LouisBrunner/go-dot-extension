// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectCapture struct {
  obj gdc.ObjectPtr
}

func createAudioEffectCapture(obj gdc.ObjectPtr) *AudioEffectCapture {
  return &AudioEffectCapture{
    obj: obj,
  }
}

func (me *AudioEffectCapture) BaseClass() string {
  return "AudioEffectCapture"
}
