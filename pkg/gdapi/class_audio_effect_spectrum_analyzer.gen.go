// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectSpectrumAnalyzer struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectSpectrumAnalyzer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectSpectrumAnalyzer) BaseClass() string {
  return "AudioEffectSpectrumAnalyzer"
}

type AudioEffectSpectrumAnalyzerFFTSize int
const (
  AudioEffectSpectrumAnalyzerFftSize256 AudioEffectSpectrumAnalyzerFFTSize = 0
  AudioEffectSpectrumAnalyzerFftSize512 AudioEffectSpectrumAnalyzerFFTSize = 1
  AudioEffectSpectrumAnalyzerFftSize1024 AudioEffectSpectrumAnalyzerFFTSize = 2
  AudioEffectSpectrumAnalyzerFftSize2048 AudioEffectSpectrumAnalyzerFFTSize = 3
  AudioEffectSpectrumAnalyzerFftSize4096 AudioEffectSpectrumAnalyzerFFTSize = 4
  AudioEffectSpectrumAnalyzerFftSizeMax AudioEffectSpectrumAnalyzerFFTSize = 5
)
