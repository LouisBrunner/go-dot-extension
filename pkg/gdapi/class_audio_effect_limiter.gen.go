// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectLimiter struct {
  AudioEffect
}

func (me *AudioEffectLimiter) BaseClass() string {
  return "AudioEffectLimiter"
}

func NewAudioEffectLimiter() *AudioEffectLimiter {
  str := StringNameFromStr("AudioEffectLimiter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectLimiter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectLimiter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectLimiter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectLimiter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectLimiter) SetCeilingDb(ceiling float64, )  {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ceiling_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ceiling), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectLimiter) GetCeilingDb() float64 {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ceiling_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectLimiter) SetThresholdDb(threshold float64, )  {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_threshold_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectLimiter) GetThresholdDb() float64 {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_threshold_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectLimiter) SetSoftClipDb(soft_clip float64, )  {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_soft_clip_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&soft_clip), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectLimiter) GetSoftClipDb() float64 {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_soft_clip_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectLimiter) SetSoftClipRatio(soft_clip float64, )  {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_soft_clip_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&soft_clip), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectLimiter) GetSoftClipRatio() float64 {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_soft_clip_ratio")
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
