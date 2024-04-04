// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type AudioStreamPolyphonic struct {
  AudioStream
}

func (me *AudioStreamPolyphonic) BaseClass() string {
  return "AudioStreamPolyphonic"
}

func NewAudioStreamPolyphonic() *AudioStreamPolyphonic {
  str := StringNameFromStr("AudioStreamPolyphonic") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamPolyphonic{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioStreamPolyphonic) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPolyphonic) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPolyphonic) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamPolyphonic) SetPolyphony(voices int64, )  {
  classNameV := StringNameFromStr("AudioStreamPolyphonic")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&voices) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPolyphonic) GetPolyphony() int64 {
  classNameV := StringNameFromStr("AudioStreamPolyphonic")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
