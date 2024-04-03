// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectChorus struct {
  AudioEffect
}

func (me *AudioEffectChorus) BaseClass() string {
  return "AudioEffectChorus"
}

func NewAudioEffectChorus() *AudioEffectChorus {
  str := StringNameFromStr("AudioEffectChorus") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectChorus{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectChorus) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectChorus) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectChorus) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectChorus) SetVoiceCount(voices int64, )  {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_voice_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voices), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectChorus) GetVoiceCount() int64 {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_voice_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectChorus) SetVoiceDelayMs(voice_idx int64, delay_ms float64, )  {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_voice_delay_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&delay_ms), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectChorus) GetVoiceDelayMs(voice_idx int64, ) float64 {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_voice_delay_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectChorus) SetVoiceRateHz(voice_idx int64, rate_hz float64, )  {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_voice_rate_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&rate_hz), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectChorus) GetVoiceRateHz(voice_idx int64, ) float64 {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_voice_rate_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectChorus) SetVoiceDepthMs(voice_idx int64, depth_ms float64, )  {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_voice_depth_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&depth_ms), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectChorus) GetVoiceDepthMs(voice_idx int64, ) float64 {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_voice_depth_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectChorus) SetVoiceLevelDb(voice_idx int64, level_db float64, )  {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_voice_level_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&level_db), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectChorus) GetVoiceLevelDb(voice_idx int64, ) float64 {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_voice_level_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectChorus) SetVoiceCutoffHz(voice_idx int64, cutoff_hz float64, )  {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_voice_cutoff_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&cutoff_hz), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectChorus) GetVoiceCutoffHz(voice_idx int64, ) float64 {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_voice_cutoff_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectChorus) SetVoicePan(voice_idx int64, pan float64, )  {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_voice_pan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), gdc.ConstTypePtr(&pan), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectChorus) GetVoicePan(voice_idx int64, ) float64 {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_voice_pan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voice_idx), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectChorus) SetWet(amount float64, )  {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_wet")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectChorus) GetWet() float64 {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_wet")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectChorus) SetDry(amount float64, )  {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dry")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectChorus) GetDry() float64 {
  classNameV := StringNameFromStr("AudioEffectChorus")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dry")
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
