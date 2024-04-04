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

type AudioEffectEQ21 struct {
  AudioEffectEQ
}

func (me *AudioEffectEQ21) BaseClass() string {
  return "AudioEffectEQ21"
}

func NewAudioEffectEQ21() *AudioEffectEQ21 {
  str := StringNameFromStr("AudioEffectEQ21") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectEQ21{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectEQ21) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectEQ21) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ21) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
