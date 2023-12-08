// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamRandomizer struct {
  obj gdc.ObjectPtr
}

func createAudioStreamRandomizer(obj gdc.ObjectPtr) *AudioStreamRandomizer {
  return &AudioStreamRandomizer{
    obj: obj,
  }
}

func (me *AudioStreamRandomizer) BaseClass() string {
  return "AudioStreamRandomizer"
}
