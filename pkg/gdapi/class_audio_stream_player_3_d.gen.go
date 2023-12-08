// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlayer3D struct {
  obj gdc.ObjectPtr
}

func createAudioStreamPlayer3D(obj gdc.ObjectPtr) *AudioStreamPlayer3D {
  return &AudioStreamPlayer3D{
    obj: obj,
  }
}

func (me *AudioStreamPlayer3D) BaseClass() string {
  return "AudioStreamPlayer3D"
}
