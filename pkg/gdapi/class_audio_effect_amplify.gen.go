// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectAmplify struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectAmplify) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectAmplify) BaseClass() string {
  return "AudioEffectAmplify"
}



// Enums

func (me *AudioEffectAmplify) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectAmplify) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectAmplify) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectAmplify) SetVolumeDb(volume float32, )  {
  classNameV := StringNameFromStr("AudioEffectAmplify")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&volume), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectAmplify) GetVolumeDb() float32 {
  classNameV := StringNameFromStr("AudioEffectAmplify")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AudioEffectAmplify) GetPropVolumeDb() float32 {
  panic("TODO: implement")
}

func (me *AudioEffectAmplify) SetPropVolumeDb(value float32) {
  panic("TODO: implement")
}