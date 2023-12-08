// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioListener3D struct {
  obj gdc.ObjectPtr
}

func createAudioListener3D(obj gdc.ObjectPtr) *AudioListener3D {
  return &AudioListener3D{
    obj: obj,
  }
}

func (me *AudioListener3D) BaseClass() string {
  return "AudioListener3D"
}
