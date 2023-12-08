// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioListener2D struct {
  obj gdc.ObjectPtr
}

func createAudioListener2D(obj gdc.ObjectPtr) *AudioListener2D {
  return &AudioListener2D{
    obj: obj,
  }
}

func (me *AudioListener2D) BaseClass() string {
  return "AudioListener2D"
}
