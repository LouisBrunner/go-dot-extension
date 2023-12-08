// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectCompressor struct {
  obj gdc.ObjectPtr
}

func createAudioEffectCompressor(obj gdc.ObjectPtr) *AudioEffectCompressor {
  return &AudioEffectCompressor{
    obj: obj,
  }
}

func (me *AudioEffectCompressor) BaseClass() string {
  return "AudioEffectCompressor"
}
