// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioStreamPlaybackOggVorbis struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamPlaybackOggVorbis) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamPlaybackOggVorbis) BaseClass() string {
  return "AudioStreamPlaybackOggVorbis"
}



// Enums

func (me *AudioStreamPlaybackOggVorbis) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPlaybackOggVorbis) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaybackOggVorbis) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties