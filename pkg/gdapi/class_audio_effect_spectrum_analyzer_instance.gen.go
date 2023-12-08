// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectSpectrumAnalyzerInstance struct {
  obj gdc.ObjectPtr
}

func createAudioEffectSpectrumAnalyzerInstance(obj gdc.ObjectPtr) *AudioEffectSpectrumAnalyzerInstance {
  return &AudioEffectSpectrumAnalyzerInstance{
    obj: obj,
  }
}

func (me *AudioEffectSpectrumAnalyzerInstance) BaseClass() string {
  return "AudioEffectSpectrumAnalyzerInstance"
}
