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



// Enums

type AudioEffectSpectrumAnalyzerInstanceMagnitudeMode int
const (
  AudioEffectSpectrumAnalyzerInstanceMagnitudeModeMagnitudeAverage AudioEffectSpectrumAnalyzerInstanceMagnitudeMode = 0
  AudioEffectSpectrumAnalyzerInstanceMagnitudeModeMagnitudeMax AudioEffectSpectrumAnalyzerInstanceMagnitudeMode = 1
)

func (me *AudioEffectSpectrumAnalyzerInstance) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectSpectrumAnalyzerInstance) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectSpectrumAnalyzerInstance) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectSpectrumAnalyzerInstance) GetMagnitudeForFrequencyRange(from_hz float32, to_hz float32, mode AudioEffectSpectrumAnalyzerInstanceMagnitudeMode, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
