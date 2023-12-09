// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamOggVorbis struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamOggVorbis) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamOggVorbis) BaseClass() string {
  return "AudioStreamOggVorbis"
}



// Enums

func (me *AudioStreamOggVorbis) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamOggVorbis) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamOggVorbis) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamOggVorbis) SetPacketSequence(packet_sequence OggPacketSequence, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) GetPacketSequence()  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) SetLoop(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) HasLoop()  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) SetLoopOffset(seconds float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) GetLoopOffset()  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) SetBpm(bpm float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) GetBpm()  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) SetBeatCount(count int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) GetBeatCount()  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) SetBarBeats(count int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamOggVorbis) GetBarBeats()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
