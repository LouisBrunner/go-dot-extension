// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamGenerator struct {
  obj gdc.ObjectPtr
}

func createAudioStreamGenerator(obj gdc.ObjectPtr) *AudioStreamGenerator {
  return &AudioStreamGenerator{
    obj: obj,
  }
}

func (me *AudioStreamGenerator) BaseClass() string {
  return "AudioStreamGenerator"
}
