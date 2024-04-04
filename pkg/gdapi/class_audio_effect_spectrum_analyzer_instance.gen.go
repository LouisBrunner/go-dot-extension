// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type AudioEffectSpectrumAnalyzerInstance struct {
  AudioEffectInstance
}

func (me *AudioEffectSpectrumAnalyzerInstance) BaseClass() string {
  return "AudioEffectSpectrumAnalyzerInstance"
}

func NewAudioEffectSpectrumAnalyzerInstance() *AudioEffectSpectrumAnalyzerInstance {
  str := StringNameFromStr("AudioEffectSpectrumAnalyzerInstance") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectSpectrumAnalyzerInstance{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *AudioEffectSpectrumAnalyzerInstance) GetMagnitudeForFrequencyRange(from_hz float64, to_hz float64, mode AudioEffectSpectrumAnalyzerInstanceMagnitudeMode, ) Vector2 {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzerInstance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_magnitude_for_frequency_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797993915) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_hz) , gdc.ConstTypePtr(&to_hz) , gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&from_hz)
  pinner.Pin(&to_hz)
  pinner.Pin(&mode)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
