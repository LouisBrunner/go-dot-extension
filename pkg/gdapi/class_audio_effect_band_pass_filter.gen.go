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

type AudioEffectBandPassFilter struct {
  AudioEffectFilter
}

func (me *AudioEffectBandPassFilter) BaseClass() string {
  return "AudioEffectBandPassFilter"
}

func NewAudioEffectBandPassFilter() *AudioEffectBandPassFilter {
  str := StringNameFromStr("AudioEffectBandPassFilter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectBandPassFilter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectBandPassFilter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectBandPassFilter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectBandPassFilter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
