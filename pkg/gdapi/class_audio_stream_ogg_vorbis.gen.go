// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamOggVorbis struct {
  obj gdc.ObjectPtr
}

func createAudioStreamOggVorbis(obj gdc.ObjectPtr) *AudioStreamOggVorbis {
  return &AudioStreamOggVorbis{
    obj: obj,
  }
}

func (me *AudioStreamOggVorbis) BaseClass() string {
  return "AudioStreamOggVorbis"
}
