// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectSpectrumAnalyzerInstance struct {
  AudioEffectInstance
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

func  (me *AudioEffectSpectrumAnalyzerInstance) GetMagnitudeForFrequencyRange(from_hz float32, to_hz float32, mode AudioEffectSpectrumAnalyzerInstanceMagnitudeMode, ) Vector2 {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzerInstance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_magnitude_for_frequency_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797993915) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_hz), gdc.ConstTypePtr(&to_hz), gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
