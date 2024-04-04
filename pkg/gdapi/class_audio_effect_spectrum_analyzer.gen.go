// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type AudioEffectSpectrumAnalyzer struct {
  AudioEffect
}

func (me *AudioEffectSpectrumAnalyzer) BaseClass() string {
  return "AudioEffectSpectrumAnalyzer"
}

func NewAudioEffectSpectrumAnalyzer() *AudioEffectSpectrumAnalyzer {
  str := StringNameFromStr("AudioEffectSpectrumAnalyzer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectSpectrumAnalyzer{}
  obj.SetBaseObject(objPtr)
  return obj
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

func (me *AudioEffectSpectrumAnalyzer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectSpectrumAnalyzer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectSpectrumAnalyzer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectSpectrumAnalyzer) SetBufferLength(seconds float64, )  {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffer_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectSpectrumAnalyzer) GetBufferLength() float64 {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectSpectrumAnalyzer) SetTapBackPos(seconds float64, )  {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap_back_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectSpectrumAnalyzer) GetTapBackPos() float64 {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tap_back_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectSpectrumAnalyzer) SetFftSize(size AudioEffectSpectrumAnalyzerFFTSize, )  {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fft_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1202879215) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectSpectrumAnalyzer) GetFftSize() AudioEffectSpectrumAnalyzerFFTSize {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fft_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3925405343) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioEffectSpectrumAnalyzerFFTSize

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
