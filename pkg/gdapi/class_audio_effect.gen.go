// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffect struct {
  obj gdc.ObjectPtr
}

func createAudioEffect(obj gdc.ObjectPtr) *AudioEffect {
  return &AudioEffect{
    obj: obj,
  }
}

func (me *AudioEffect) BaseClass() string {
  return "AudioEffect"
}
