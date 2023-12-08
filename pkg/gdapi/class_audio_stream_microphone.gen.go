// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamMicrophone struct {
  obj gdc.ObjectPtr
}

func createAudioStreamMicrophone(obj gdc.ObjectPtr) *AudioStreamMicrophone {
  return &AudioStreamMicrophone{
    obj: obj,
  }
}

func (me *AudioStreamMicrophone) BaseClass() string {
  return "AudioStreamMicrophone"
}
