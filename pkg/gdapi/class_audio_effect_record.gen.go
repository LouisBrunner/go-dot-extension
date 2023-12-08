// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectRecord struct {
  obj gdc.ObjectPtr
}

func createAudioEffectRecord(obj gdc.ObjectPtr) *AudioEffectRecord {
  return &AudioEffectRecord{
    obj: obj,
  }
}

func (me *AudioEffectRecord) BaseClass() string {
  return "AudioEffectRecord"
}
