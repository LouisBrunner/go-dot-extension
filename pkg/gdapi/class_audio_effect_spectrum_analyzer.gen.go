// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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



// Enums

type AudioEffectSpectrumAnalyzerFFTSize int
const (
  AudioEffectSpectrumAnalyzerFFTSizeFftSize256 AudioEffectSpectrumAnalyzerFFTSize = 0
  AudioEffectSpectrumAnalyzerFFTSizeFftSize512 AudioEffectSpectrumAnalyzerFFTSize = 1
  AudioEffectSpectrumAnalyzerFFTSizeFftSize1024 AudioEffectSpectrumAnalyzerFFTSize = 2
  AudioEffectSpectrumAnalyzerFFTSizeFftSize2048 AudioEffectSpectrumAnalyzerFFTSize = 3
  AudioEffectSpectrumAnalyzerFFTSizeFftSize4096 AudioEffectSpectrumAnalyzerFFTSize = 4
  AudioEffectSpectrumAnalyzerFFTSizeFftSizeMax AudioEffectSpectrumAnalyzerFFTSize = 5
)

func (me *AudioEffectSpectrumAnalyzer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectSpectrumAnalyzer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectSpectrumAnalyzer) SetBufferLength(seconds float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectSpectrumAnalyzer) GetBufferLength()  {
  panic("TODO: implement")
}

func  (me *AudioEffectSpectrumAnalyzer) SetTapBackPos(seconds float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectSpectrumAnalyzer) GetTapBackPos()  {
  panic("TODO: implement")
}

func  (me *AudioEffectSpectrumAnalyzer) SetFftSize(size AudioEffectSpectrumAnalyzerFFTSize, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectSpectrumAnalyzer) GetFftSize()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
