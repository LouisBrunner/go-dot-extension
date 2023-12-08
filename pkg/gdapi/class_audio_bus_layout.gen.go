// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioBusLayout struct {
  obj gdc.ObjectPtr
}

func createAudioBusLayout(obj gdc.ObjectPtr) *AudioBusLayout {
  return &AudioBusLayout{
    obj: obj,
  }
}

func (me *AudioBusLayout) BaseClass() string {
  return "AudioBusLayout"
}
