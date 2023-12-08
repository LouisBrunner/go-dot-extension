// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStream struct {
  obj gdc.ObjectPtr
}

func createAudioStream(obj gdc.ObjectPtr) *AudioStream {
  return &AudioStream{
    obj: obj,
  }
}

func (me *AudioStream) BaseClass() string {
  return "AudioStream"
}
