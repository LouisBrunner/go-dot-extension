// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamGeneratorPlayback struct {
  obj gdc.ObjectPtr
}

func createAudioStreamGeneratorPlayback(obj gdc.ObjectPtr) *AudioStreamGeneratorPlayback {
  return &AudioStreamGeneratorPlayback{
    obj: obj,
  }
}

func (me *AudioStreamGeneratorPlayback) BaseClass() string {
  return "AudioStreamGeneratorPlayback"
}
