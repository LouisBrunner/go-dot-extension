// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

// TODO: properties (class)

// TODO: signals (class)
