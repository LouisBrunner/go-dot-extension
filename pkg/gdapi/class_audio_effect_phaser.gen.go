// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectPhaser struct {
  obj gdc.ObjectPtr
}

func createAudioEffectPhaser(obj gdc.ObjectPtr) *AudioEffectPhaser {
  return &AudioEffectPhaser{
    obj: obj,
  }
}

func (me *AudioEffectPhaser) BaseClass() string {
  return "AudioEffectPhaser"
}
