// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlayback struct {
  obj gdc.ObjectPtr
}

func createAudioStreamPlayback(obj gdc.ObjectPtr) *AudioStreamPlayback {
  return &AudioStreamPlayback{
    obj: obj,
  }
}

func (me *AudioStreamPlayback) BaseClass() string {
  return "AudioStreamPlayback"
}
