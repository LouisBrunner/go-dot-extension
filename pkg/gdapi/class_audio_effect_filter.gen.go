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

type AudioEffectFilter struct {
  AudioEffect
}

func (me *AudioEffectFilter) BaseClass() string {
  return "AudioEffectFilter"
}

func NewAudioEffectFilter() *AudioEffectFilter {
  str := StringNameFromStr("AudioEffectFilter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectFilter{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *AudioEffectFilter) SetCutoff(freq float64, )  {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cutoff")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&freq) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectFilter) GetCutoff() float64 {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cutoff")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectFilter) SetResonance(amount float64, )  {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_resonance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectFilter) GetResonance() float64 {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_resonance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectFilter) SetGain(amount float64, )  {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectFilter) GetGain() float64 {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gain")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectFilter) SetDb(amount AudioEffectFilterFilterDB, )  {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 771740901) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectFilter) GetDb() AudioEffectFilterFilterDB {
  classNameV := StringNameFromStr("AudioEffectFilter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3981721890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioEffectFilterFilterDB

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
