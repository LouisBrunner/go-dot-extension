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

type AudioEffectPanner struct {
  AudioEffect
}

func (me *AudioEffectPanner) BaseClass() string {
  return "AudioEffectPanner"
}

func NewAudioEffectPanner() *AudioEffectPanner {
  str := StringNameFromStr("AudioEffectPanner") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectPanner{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectPanner) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectPanner) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectPanner) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectPanner) SetPan(cpanume float64, )  {
  classNameV := StringNameFromStr("AudioEffectPanner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pan")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cpanume) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectPanner) GetPan() float64 {
  classNameV := StringNameFromStr("AudioEffectPanner")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pan")
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
