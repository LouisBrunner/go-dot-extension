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

type AudioEffectHighPassFilter struct {
  AudioEffectFilter
}

func (me *AudioEffectHighPassFilter) BaseClass() string {
  return "AudioEffectHighPassFilter"
}

func NewAudioEffectHighPassFilter() *AudioEffectHighPassFilter {
  str := StringNameFromStr("AudioEffectHighPassFilter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectHighPassFilter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectHighPassFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectHighPassFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectHighPassFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
