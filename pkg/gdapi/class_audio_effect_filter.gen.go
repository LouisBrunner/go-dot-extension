// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectFilter struct {
  AudioEffect
}

func (me *AudioEffectFilter) BaseClass() string {
  return "AudioEffectFilter"
}



// Enums

type AudioEffectFilterFilterDB int
const (
  AudioEffectFilterFilterDBFilter6Db AudioEffectFilterFilterDB = 0
  AudioEffectFilterFilterDBFilter12Db AudioEffectFilterFilterDB = 1
  AudioEffectFilterFilterDBFilter18Db AudioEffectFilterFilterDB = 2
  AudioEffectFilterFilterDBFilter24Db AudioEffectFilterFilterDB = 3
)

func (me *AudioEffectFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectFilter) SetCutoff(freq float32, )  {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cutoff")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freq), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectFilter) GetCutoff() float32 {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cutoff")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectFilter) SetResonance(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resonance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectFilter) GetResonance() float32 {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resonance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectFilter) SetGain(amount float32, )  {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectFilter) GetGain() float32 {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectFilter) SetDb(amount AudioEffectFilterFilterDB, )  {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 771740901) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioEffectFilter) GetDb() AudioEffectFilterFilterDB {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3981721890) // FIXME: should cache?
  var ret AudioEffectFilterFilterDB
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
