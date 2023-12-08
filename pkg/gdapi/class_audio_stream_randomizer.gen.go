// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamRandomizer struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamRandomizer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamRandomizer) BaseClass() string {
  return "AudioStreamRandomizer"
}

type AudioStreamRandomizerPlaybackMode int
const (
  AudioStreamRandomizerPlaybackModePlaybackRandomNoRepeats AudioStreamRandomizerPlaybackMode = 0
  AudioStreamRandomizerPlaybackModePlaybackRandom AudioStreamRandomizerPlaybackMode = 1
  AudioStreamRandomizerPlaybackModePlaybackSequential AudioStreamRandomizerPlaybackMode = 2
)

func  (me *AudioStreamRandomizer) AddStream(index int, stream AudioStream, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) MoveStream(index_from int, index_to int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) RemoveStream(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) SetStream(index int, stream AudioStream, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) GetStream(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) SetStreamProbabilityWeight(index int, weight float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) GetStreamProbabilityWeight(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) SetStreamsCount(count int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) GetStreamsCount() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) SetRandomPitch(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) GetRandomPitch() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) SetRandomVolumeOffsetDb(db_offset float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) GetRandomVolumeOffsetDb() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) SetPlaybackMode(mode AudioStreamRandomizerPlaybackMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamRandomizer) GetPlaybackMode() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
