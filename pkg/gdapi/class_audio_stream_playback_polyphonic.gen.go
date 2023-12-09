// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlaybackPolyphonic struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamPlaybackPolyphonic) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamPlaybackPolyphonic) BaseClass() string {
  return "AudioStreamPlaybackPolyphonic"
}



// Constants

var (
  AudioStreamPlaybackPolyphonicInvalidId = "-1" // TODO: construct correctly
)

// Enums

func (me *AudioStreamPlaybackPolyphonic) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaybackPolyphonic) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamPlaybackPolyphonic) PlayStream(stream AudioStream, from_offset float32, volume_db float32, pitch_scale float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlaybackPolyphonic) SetStreamVolume(stream int, volume_db float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlaybackPolyphonic) SetStreamPitchScale(stream int, pitch_scale float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlaybackPolyphonic) IsStreamPlaying(stream int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlaybackPolyphonic) StopStream(stream int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
