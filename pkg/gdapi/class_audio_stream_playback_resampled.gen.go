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

type AudioStreamPlaybackResampled struct {
  AudioStreamPlayback
}

func (me *AudioStreamPlaybackResampled) BaseClass() string {
  return "AudioStreamPlaybackResampled"
}

func NewAudioStreamPlaybackResampled() *AudioStreamPlaybackResampled {
  str := StringNameFromStr("AudioStreamPlaybackResampled") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamPlaybackResampled{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioStreamPlaybackResampled) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPlaybackResampled) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaybackResampled) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamPlaybackResampled) BeginResample()  {
  classNameV := StringNameFromStr("AudioStreamPlaybackResampled")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("begin_resample")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
