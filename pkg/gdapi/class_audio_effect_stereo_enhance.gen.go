// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectStereoEnhance struct {
  AudioEffect
}

func (me *AudioEffectStereoEnhance) BaseClass() string {
  return "AudioEffectStereoEnhance"
}

func NewAudioEffectStereoEnhance() *AudioEffectStereoEnhance {
  str := StringNameFromStr("AudioEffectStereoEnhance") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectStereoEnhance{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectStereoEnhance) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectStereoEnhance) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectStereoEnhance) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectStereoEnhance) SetPanPullout(amount float64, )  {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pan_pullout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectStereoEnhance) GetPanPullout() float64 {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pan_pullout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectStereoEnhance) SetTimePullout(amount float64, )  {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_time_pullout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectStereoEnhance) GetTimePullout() float64 {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_pullout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectStereoEnhance) SetSurround(amount float64, )  {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_surround")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectStereoEnhance) GetSurround() float64 {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surround")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
