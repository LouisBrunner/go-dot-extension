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

type AudioEffectPitchShift struct {
  AudioEffect
}

func (me *AudioEffectPitchShift) BaseClass() string {
  return "AudioEffectPitchShift"
}

func NewAudioEffectPitchShift() *AudioEffectPitchShift {
  str := StringNameFromStr("AudioEffectPitchShift") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectPitchShift{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *AudioEffectPitchShift) SetPitchScale(rate float64, )  {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pitch_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rate) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectPitchShift) GetPitchScale() float64 {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pitch_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectPitchShift) SetOversampling(amount int64, )  {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectPitchShift) GetOversampling() int64 {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_oversampling")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectPitchShift) SetFftSize(size AudioEffectPitchShiftFFTSize, )  {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fft_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2323518741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectPitchShift) GetFftSize() AudioEffectPitchShiftFFTSize {
  classNameV := StringNameFromStr("AudioEffectPitchShift")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fft_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2361246789) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioEffectPitchShiftFFTSize

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
