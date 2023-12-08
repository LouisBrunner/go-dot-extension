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

// TODO: needed?
// const (
// // )

var (
  AudioStreamPlaybackPolyphonicInvalidId = "-1" // TODO: construct correctly
)

func  (me *AudioStreamPlaybackPolyphonic) PlayStream(stream AudioStream, from_offset float32, volume_db float32, pitch_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlaybackPolyphonic) SetStreamVolume(stream int, volume_db float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlaybackPolyphonic) SetStreamPitchScale(stream int, pitch_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlaybackPolyphonic) IsStreamPlaying(stream int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlaybackPolyphonic) StopStream(stream int, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
