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



// Enums

type AudioStreamRandomizerPlaybackMode int
const (
  AudioStreamRandomizerPlaybackModePlaybackRandomNoRepeats AudioStreamRandomizerPlaybackMode = 0
  AudioStreamRandomizerPlaybackModePlaybackRandom AudioStreamRandomizerPlaybackMode = 1
  AudioStreamRandomizerPlaybackModePlaybackSequential AudioStreamRandomizerPlaybackMode = 2
)

func (me *AudioStreamRandomizer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamRandomizer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamRandomizer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamRandomizer) AddStream(index int, stream AudioStream, weight float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) MoveStream(index_from int, index_to int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) RemoveStream(index int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) SetStream(index int, stream AudioStream, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) GetStream(index int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) SetStreamProbabilityWeight(index int, weight float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) GetStreamProbabilityWeight(index int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) SetStreamsCount(count int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) GetStreamsCount()  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) SetRandomPitch(scale float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) GetRandomPitch()  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) SetRandomVolumeOffsetDb(db_offset float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) GetRandomVolumeOffsetDb()  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) SetPlaybackMode(mode AudioStreamRandomizerPlaybackMode, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamRandomizer) GetPlaybackMode()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
