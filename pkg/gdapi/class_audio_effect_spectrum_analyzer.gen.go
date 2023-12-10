// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *AudioEffectSpectrumAnalyzer) SetBufferLength(seconds float32, )  {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffer_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectSpectrumAnalyzer) GetBufferLength() float32 {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_buffer_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectSpectrumAnalyzer) SetTapBackPos(seconds float32, )  {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tap_back_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectSpectrumAnalyzer) GetTapBackPos() float32 {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tap_back_pos")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectSpectrumAnalyzer) SetFftSize(size AudioEffectSpectrumAnalyzerFFTSize, )  {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fft_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1202879215) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectSpectrumAnalyzer) GetFftSize() AudioEffectSpectrumAnalyzerFFTSize {
  classNameV := StringNameFromStr("AudioEffectSpectrumAnalyzer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fft_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3925405343) // FIXME: should cache?
  var ret AudioEffectSpectrumAnalyzerFFTSize
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AudioEffectSpectrumAnalyzer) GetPropBufferLength() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectSpectrumAnalyzer) SetPropBufferLength(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectSpectrumAnalyzer) GetPropTapBackPos() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectSpectrumAnalyzer) SetPropTapBackPos(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectSpectrumAnalyzer) GetPropFftSize() int {
  panic("TODO: implement")
}

func (me *AudioEffectSpectrumAnalyzer) SetPropFftSize(value int) {
  panic("TODO: implement")
}