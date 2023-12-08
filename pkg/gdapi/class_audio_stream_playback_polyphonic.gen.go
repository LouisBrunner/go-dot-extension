// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlaybackPolyphonic struct {
  obj gdc.ObjectPtr
}

func createAudioStreamPlaybackPolyphonic(obj gdc.ObjectPtr) *AudioStreamPlaybackPolyphonic {
  return &AudioStreamPlaybackPolyphonic{
    obj: obj,
  }
}

func (me *AudioStreamPlaybackPolyphonic) BaseClass() string {
  return "AudioStreamPlaybackPolyphonic"
}
