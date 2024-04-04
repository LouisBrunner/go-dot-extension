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

type AudioStreamGenerator struct {
  AudioStream
}

func (me *AudioStreamGenerator) BaseClass() string {
  return "AudioStreamGenerator"
}

func NewAudioStreamGenerator() *AudioStreamGenerator {
  str := StringNameFromStr("AudioStreamGenerator") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamGenerator{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioStreamGenerator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamGenerator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamGenerator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamGenerator) SetMixRate(hz float64, )  {
  classNameV := StringNameFromStr("AudioStreamGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mix_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&hz) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamGenerator) GetMixRate() float64 {
  classNameV := StringNameFromStr("AudioStreamGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mix_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamGenerator) SetBufferLength(seconds float64, )  {
  classNameV := StringNameFromStr("AudioStreamGenerator")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_buffer_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamGenerator) GetBufferLength() float64 {
  classNameV := StringNameFromStr("AudioStreamGenerator")
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
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
