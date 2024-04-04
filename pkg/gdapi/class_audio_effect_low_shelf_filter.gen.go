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

type AudioEffectLowShelfFilter struct {
  AudioEffectFilter
}

func (me *AudioEffectLowShelfFilter) BaseClass() string {
  return "AudioEffectLowShelfFilter"
}

func NewAudioEffectLowShelfFilter() *AudioEffectLowShelfFilter {
  str := StringNameFromStr("AudioEffectLowShelfFilter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectLowShelfFilter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectLowShelfFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectLowShelfFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectLowShelfFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
