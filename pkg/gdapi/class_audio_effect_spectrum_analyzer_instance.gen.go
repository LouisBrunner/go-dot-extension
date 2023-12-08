// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectSpectrumAnalyzerInstance struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectSpectrumAnalyzerInstance) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectSpectrumAnalyzerInstance) BaseClass() string {
  return "AudioEffectSpectrumAnalyzerInstance"
}

type AudioEffectSpectrumAnalyzerInstanceMagnitudeMode int
const (
  AudioEffectSpectrumAnalyzerInstanceMagnitudeModeMagnitudeAverage AudioEffectSpectrumAnalyzerInstanceMagnitudeMode = 0
  AudioEffectSpectrumAnalyzerInstanceMagnitudeModeMagnitudeMax AudioEffectSpectrumAnalyzerInstanceMagnitudeMode = 1
)

func  (me *AudioEffectSpectrumAnalyzerInstance) GetMagnitudeForFrequencyRange(from_hz float32, to_hz float32, mode AudioEffectSpectrumAnalyzerInstanceMagnitudeMode, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
