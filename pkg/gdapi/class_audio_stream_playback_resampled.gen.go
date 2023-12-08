// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlaybackResampled struct {
  obj gdc.ObjectPtr
}

func createAudioStreamPlaybackResampled(obj gdc.ObjectPtr) *AudioStreamPlaybackResampled {
  return &AudioStreamPlaybackResampled{
    obj: obj,
  }
}

func (me *AudioStreamPlaybackResampled) BaseClass() string {
  return "AudioStreamPlaybackResampled"
}
