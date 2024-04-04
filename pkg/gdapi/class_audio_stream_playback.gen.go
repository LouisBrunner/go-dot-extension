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

type AudioStreamPlayback struct {
  RefCounted
}

func (me *AudioStreamPlayback) BaseClass() string {
  return "AudioStreamPlayback"
}

func NewAudioStreamPlayback() *AudioStreamPlayback {
  str := StringNameFromStr("AudioStreamPlayback") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamPlayback{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioStreamPlayback) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPlayback) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlayback) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
