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

type AudioStreamMicrophone struct {
  AudioStream
}

func (me *AudioStreamMicrophone) BaseClass() string {
  return "AudioStreamMicrophone"
}

func NewAudioStreamMicrophone() *AudioStreamMicrophone {
  str := StringNameFromStr("AudioStreamMicrophone") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamMicrophone{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioStreamMicrophone) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamMicrophone) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamMicrophone) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
