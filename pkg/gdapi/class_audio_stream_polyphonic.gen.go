// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPolyphonic struct {
  obj gdc.ObjectPtr
}

func createAudioStreamPolyphonic(obj gdc.ObjectPtr) *AudioStreamPolyphonic {
  return &AudioStreamPolyphonic{
    obj: obj,
  }
}

func (me *AudioStreamPolyphonic) BaseClass() string {
  return "AudioStreamPolyphonic"
}
