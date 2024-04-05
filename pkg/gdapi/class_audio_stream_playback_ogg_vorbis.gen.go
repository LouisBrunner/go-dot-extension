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

type ptrsForAudioStreamPlaybackOggVorbisList struct {
}

var ptrsForAudioStreamPlaybackOggVorbis ptrsForAudioStreamPlaybackOggVorbisList

func initAudioStreamPlaybackOggVorbisPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioStreamPlaybackOggVorbis")
  defer className.Destroy()
}

type AudioStreamPlaybackOggVorbis struct {
  AudioStreamPlaybackResampled
}

func (me *AudioStreamPlaybackOggVorbis) BaseClass() string {
  return "AudioStreamPlaybackOggVorbis"
}

func NewAudioStreamPlaybackOggVorbis() *AudioStreamPlaybackOggVorbis {
  str := StringNameFromStr("AudioStreamPlaybackOggVorbis") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamPlaybackOggVorbis{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
