// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectCompressor struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectCompressor) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectCompressor) BaseClass() string {
  return "AudioEffectCompressor"
}



// Enums

func (me *AudioEffectCompressor) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectCompressor) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectCompressor) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectCompressor) SetThreshold(threshold float32, )  {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectCompressor) GetThreshold() float32 {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_threshold")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectCompressor) SetRatio(ratio float32, )  {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ratio), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectCompressor) GetRatio() float32 {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectCompressor) SetGain(gain float32, )  {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&gain), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectCompressor) GetGain() float32 {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectCompressor) SetAttackUs(attack_us float32, )  {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_attack_us")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&attack_us), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectCompressor) GetAttackUs() float32 {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_attack_us")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectCompressor) SetReleaseMs(release_ms float32, )  {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_release_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&release_ms), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectCompressor) GetReleaseMs() float32 {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_release_ms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectCompressor) SetMix(mix float32, )  {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectCompressor) GetMix() float32 {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectCompressor) SetSidechain(sidechain StringName, )  {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_sidechain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(sidechain.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectCompressor) GetSidechain() StringName {
  classNameV := StringNameFromStr("AudioEffectCompressor")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_sidechain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
