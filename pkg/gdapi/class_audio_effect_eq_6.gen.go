// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectEQ6 struct {
  AudioEffectEQ
}

func (me *AudioEffectEQ6) BaseClass() string {
  return "AudioEffectEQ6"
}

func NewAudioEffectEQ6() *AudioEffectEQ6 {
  str := StringNameFromStr("AudioEffectEQ6") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectEQ6{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectEQ6) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectEQ6) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectEQ6) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
