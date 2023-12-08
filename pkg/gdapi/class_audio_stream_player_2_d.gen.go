// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlayer2D struct {
  obj gdc.ObjectPtr
}

func createAudioStreamPlayer2D(obj gdc.ObjectPtr) *AudioStreamPlayer2D {
  return &AudioStreamPlayer2D{
    obj: obj,
  }
}

func (me *AudioStreamPlayer2D) BaseClass() string {
  return "AudioStreamPlayer2D"
}
