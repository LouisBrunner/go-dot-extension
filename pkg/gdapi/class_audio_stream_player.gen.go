// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlayer struct {
  obj gdc.ObjectPtr
}

func createAudioStreamPlayer(obj gdc.ObjectPtr) *AudioStreamPlayer {
  return &AudioStreamPlayer{
    obj: obj,
  }
}

func (me *AudioStreamPlayer) BaseClass() string {
  return "AudioStreamPlayer"
}
