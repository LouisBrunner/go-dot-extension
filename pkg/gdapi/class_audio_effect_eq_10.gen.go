// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectEQ10 struct {
  AudioEffectEQ
}

func (me *AudioEffectEQ10) BaseClass() string {
  return "AudioEffectEQ10"
}

func NewAudioEffectEQ10() *AudioEffectEQ10 {
  str := StringNameFromStr("AudioEffectEQ10") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectEQ10{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectEQ10) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectEQ10) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ10) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
