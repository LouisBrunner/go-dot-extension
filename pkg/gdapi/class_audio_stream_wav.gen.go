// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamWAV struct {
  obj gdc.ObjectPtr
}

func createAudioStreamWAV(obj gdc.ObjectPtr) *AudioStreamWAV {
  return &AudioStreamWAV{
    obj: obj,
  }
}

func (me *AudioStreamWAV) BaseClass() string {
  return "AudioStreamWAV"
}
