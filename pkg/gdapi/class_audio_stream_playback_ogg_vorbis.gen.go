// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlaybackOggVorbis struct {
  obj gdc.ObjectPtr
}

func createAudioStreamPlaybackOggVorbis(obj gdc.ObjectPtr) *AudioStreamPlaybackOggVorbis {
  return &AudioStreamPlaybackOggVorbis{
    obj: obj,
  }
}

func (me *AudioStreamPlaybackOggVorbis) BaseClass() string {
  return "AudioStreamPlaybackOggVorbis"
}
