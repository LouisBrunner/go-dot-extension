// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlayback struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamPlayback) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamPlayback) BaseClass() string {
  return "AudioStreamPlayback"
}



// Enums

func (me *AudioStreamPlayback) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlayback) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamPlayback) XStart(from_pos float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayback) XStop()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayback) XIsPlaying()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayback) XGetLoopCount()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayback) XGetPlaybackPosition()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayback) XSeek(position float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayback) XMix(buffer *AudioFrame, rate_scale float32, frames int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayback) XTagUsedStreams()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
