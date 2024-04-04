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

type AudioEffectLowPassFilter struct {
  AudioEffectFilter
}

func (me *AudioEffectLowPassFilter) BaseClass() string {
  return "AudioEffectLowPassFilter"
}

func NewAudioEffectLowPassFilter() *AudioEffectLowPassFilter {
  str := StringNameFromStr("AudioEffectLowPassFilter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectLowPassFilter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectLowPassFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectLowPassFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectLowPassFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
