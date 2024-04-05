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

type ptrsForVideoStreamPlaybackList struct {
  fnXStop gdc.MethodBindPtr
  fnXPlay gdc.MethodBindPtr
  fnXIsPlaying gdc.MethodBindPtr
  fnXSetPaused gdc.MethodBindPtr
  fnXIsPaused gdc.MethodBindPtr
  fnXGetLength gdc.MethodBindPtr
  fnXGetPlaybackPosition gdc.MethodBindPtr
  fnXSeek gdc.MethodBindPtr
  fnXSetAudioTrack gdc.MethodBindPtr
  fnXGetTexture gdc.MethodBindPtr
  fnXUpdate gdc.MethodBindPtr
  fnXGetChannels gdc.MethodBindPtr
  fnXGetMixRate gdc.MethodBindPtr
  fnMixAudio gdc.MethodBindPtr
}

var ptrsForVideoStreamPlayback ptrsForVideoStreamPlaybackList

func initVideoStreamPlaybackPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VideoStreamPlayback")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("mix_audio")
    defer methodName.Destroy()
    ptrsForVideoStreamPlayback.fnMixAudio = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 93876830))
  }
}

type VideoStreamPlayback struct {
  Resource
}

func (me *VideoStreamPlayback) BaseClass() string {
  return "VideoStreamPlayback"
}

func NewVideoStreamPlayback() *VideoStreamPlayback {
  str := StringNameFromStr("VideoStreamPlayback") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VideoStreamPlayback{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *VideoStreamPlayback) MixAudio(num_frames int64, buffer PackedFloat32Array, offset int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&num_frames) , buffer.AsCTypePtr(), gdc.ConstTypePtr(&offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&num_frames)
  pinner.Pin(&offset)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVideoStreamPlayback.fnMixAudio), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

// Signals
