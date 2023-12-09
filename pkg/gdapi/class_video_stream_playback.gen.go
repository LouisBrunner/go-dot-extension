// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VideoStreamPlayback struct {
  obj gdc.ObjectPtr
}

func (me *VideoStreamPlayback) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VideoStreamPlayback) BaseClass() string {
  return "VideoStreamPlayback"
}



// Enums

func (me *VideoStreamPlayback) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VideoStreamPlayback) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VideoStreamPlayback) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VideoStreamPlayback) XStop()  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XPlay()  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XIsPlaying()  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XSetPaused(paused bool, )  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XIsPaused()  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XGetLength()  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XGetPlaybackPosition()  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XSeek(time float32, )  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XSetAudioTrack(idx int, )  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XGetTexture()  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XUpdate(delta float32, )  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XGetChannels()  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) XGetMixRate()  {
  panic("TODO: implement")
}

func  (me *VideoStreamPlayback) MixAudio(num_frames int, buffer PackedFloat32Array, offset int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
