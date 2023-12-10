// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectStereoEnhance struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectStereoEnhance) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectStereoEnhance) BaseClass() string {
  return "AudioEffectStereoEnhance"
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

func  (me *AudioEffectStereoEnhance) SetPanPullout(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pan_pullout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectStereoEnhance) GetPanPullout() float32 {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pan_pullout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectStereoEnhance) SetTimePullout(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_time_pullout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectStereoEnhance) GetTimePullout() float32 {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_pullout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectStereoEnhance) SetSurround(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_surround")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectStereoEnhance) GetSurround() float32 {
  classNameV := StringNameFromStr("AudioEffectStereoEnhance")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_surround")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AudioEffectStereoEnhance) GetPropPanPullout() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectStereoEnhance) SetPropPanPullout(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectStereoEnhance) GetPropTimePulloutMs() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectStereoEnhance) SetPropTimePulloutMs(value float32) {
  panic("TODO: implement")
}

func (me *AudioEffectStereoEnhance) GetPropSurround() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectStereoEnhance) SetPropSurround(value float32) {
  panic("TODO: implement")
}