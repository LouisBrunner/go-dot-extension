// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioStream struct {
  Resource
}

func (me *AudioStream) BaseClass() string {
  return "AudioStream"
}

func NewAudioStream() *AudioStream {
  str := StringNameFromStr("AudioStream") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStream{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioStream) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStream) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStream) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStream) GetLength() float64 {
  classNameV := StringNameFromStr("AudioStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStream) IsMonophonic() bool {
  classNameV := StringNameFromStr("AudioStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_monophonic")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStream) InstantiatePlayback() AudioStreamPlayback {
  classNameV := StringNameFromStr("AudioStream")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("instantiate_playback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 210135309) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewAudioStreamPlayback()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
