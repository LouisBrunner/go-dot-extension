// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamMP3 struct {
  obj gdc.ObjectPtr
}

func createAudioStreamMP3(obj gdc.ObjectPtr) *AudioStreamMP3 {
  return &AudioStreamMP3{
    obj: obj,
  }
}

func (me *AudioStreamMP3) BaseClass() string {
  return "AudioStreamMP3"
}
