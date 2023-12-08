// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectReverb struct {
  obj gdc.ObjectPtr
}

func createAudioEffectReverb(obj gdc.ObjectPtr) *AudioEffectReverb {
  return &AudioEffectReverb{
    obj: obj,
  }
}

func (me *AudioEffectReverb) BaseClass() string {
  return "AudioEffectReverb"
}
