// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectPitchShift struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectPitchShift) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectPitchShift) BaseClass() string {
  return "AudioEffectPitchShift"
}



// Enums

type AudioEffectPitchShiftFFTSize int
const (
  AudioEffectPitchShiftFFTSizeFftSize256 AudioEffectPitchShiftFFTSize = 0
  AudioEffectPitchShiftFFTSizeFftSize512 AudioEffectPitchShiftFFTSize = 1
  AudioEffectPitchShiftFFTSizeFftSize1024 AudioEffectPitchShiftFFTSize = 2
  AudioEffectPitchShiftFFTSizeFftSize2048 AudioEffectPitchShiftFFTSize = 3
  AudioEffectPitchShiftFFTSizeFftSize4096 AudioEffectPitchShiftFFTSize = 4
  AudioEffectPitchShiftFFTSizeFftSizeMax AudioEffectPitchShiftFFTSize = 5
)

func (me *AudioEffectPitchShift) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectPitchShift) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectPitchShift) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectPitchShift) SetPitchScale(rate float32, )  {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pitch_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rate), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectPitchShift) GetPitchScale() float32 {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pitch_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectPitchShift) SetOversampling(amount int, )  {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectPitchShift) GetOversampling() int {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectPitchShift) SetFftSize(size AudioEffectPitchShiftFFTSize, )  {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fft_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323518741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectPitchShift) GetFftSize() AudioEffectPitchShiftFFTSize {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fft_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2361246789) // FIXME: should cache?
  var ret AudioEffectPitchShiftFFTSize
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
