// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectSpectrumAnalyzer struct {
  obj gdc.ObjectPtr
}

func createAudioEffectSpectrumAnalyzer(obj gdc.ObjectPtr) *AudioEffectSpectrumAnalyzer {
  return &AudioEffectSpectrumAnalyzer{
    obj: obj,
  }
}

func (me *AudioEffectSpectrumAnalyzer) BaseClass() string {
  return "AudioEffectSpectrumAnalyzer"
}
