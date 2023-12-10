// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectLimiter struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectLimiter) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectLimiter) BaseClass() string {
  return "AudioEffectLimiter"
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

func  (me *AudioEffectLimiter) SetCeilingDb(ceiling float32, )  {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ceiling_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&ceiling), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectLimiter) GetCeilingDb() float32 {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ceiling_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectLimiter) SetThresholdDb(threshold float32, )  {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_threshold_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&threshold), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectLimiter) GetThresholdDb() float32 {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_threshold_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectLimiter) SetSoftClipDb(soft_clip float32, )  {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_soft_clip_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&soft_clip), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectLimiter) GetSoftClipDb() float32 {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_soft_clip_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectLimiter) SetSoftClipRatio(soft_clip float32, )  {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_soft_clip_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&soft_clip), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectLimiter) GetSoftClipRatio() float32 {
  classNameV := StringNameFromStr("AudioEffectLimiter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_soft_clip_ratio")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AudioEffectLimiter) GetPropCeilingDb() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectLimiter) SetPropCeilingDb(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectLimiter) GetPropThresholdDb() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectLimiter) SetPropThresholdDb(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectLimiter) GetPropSoftClipDb() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectLimiter) SetPropSoftClipDb(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectLimiter) GetPropSoftClipRatio() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectLimiter) SetPropSoftClipRatio(value float32) {
  panic("TODO: implement")
}