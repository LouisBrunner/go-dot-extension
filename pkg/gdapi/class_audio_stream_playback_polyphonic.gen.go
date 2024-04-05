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

type ptrsForAudioStreamPlaybackPolyphonicList struct {
  fnPlayStream gdc.MethodBindPtr
  fnSetStreamVolume gdc.MethodBindPtr
  fnSetStreamPitchScale gdc.MethodBindPtr
  fnIsStreamPlaying gdc.MethodBindPtr
  fnStopStream gdc.MethodBindPtr
}

var ptrsForAudioStreamPlaybackPolyphonic ptrsForAudioStreamPlaybackPolyphonicList

func initAudioStreamPlaybackPolyphonicPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioStreamPlaybackPolyphonic")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("play_stream")
    defer methodName.Destroy()
    ptrsForAudioStreamPlaybackPolyphonic.fnPlayStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 604492179))
  }
  {
    methodName := StringNameFromStr("set_stream_volume")
    defer methodName.Destroy()
    ptrsForAudioStreamPlaybackPolyphonic.fnSetStreamVolume = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("set_stream_pitch_scale")
    defer methodName.Destroy()
    ptrsForAudioStreamPlaybackPolyphonic.fnSetStreamPitchScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("is_stream_playing")
    defer methodName.Destroy()
    ptrsForAudioStreamPlaybackPolyphonic.fnIsStreamPlaying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("stop_stream")
    defer methodName.Destroy()
    ptrsForAudioStreamPlaybackPolyphonic.fnStopStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
}

type AudioStreamPlaybackPolyphonic struct {
  AudioStreamPlayback
}

func (me *AudioStreamPlaybackPolyphonic) BaseClass() string {
  return "AudioStreamPlaybackPolyphonic"
}

func NewAudioStreamPlaybackPolyphonic() *AudioStreamPlaybackPolyphonic {
  str := StringNameFromStr("AudioStreamPlaybackPolyphonic") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamPlaybackPolyphonic{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  AudioStreamPlaybackPolyphonicInvalidId = "-1" // TODO: construct correctly
)

// Enums

func (me *AudioStreamPlaybackPolyphonic) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPlaybackPolyphonic) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaybackPolyphonic) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamPlaybackPolyphonic) PlayStream(stream AudioStream, from_offset float64, volume_db float64, pitch_scale float64, ) int64 {
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), gdc.ConstTypePtr(&from_offset) , gdc.ConstTypePtr(&volume_db) , gdc.ConstTypePtr(&pitch_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&from_offset)
  pinner.Pin(&volume_db)
  pinner.Pin(&pitch_scale)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaybackPolyphonic.fnPlayStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlaybackPolyphonic) SetStreamVolume(stream int64, volume_db float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream) , gdc.ConstTypePtr(&volume_db) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaybackPolyphonic.fnSetStreamVolume), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlaybackPolyphonic) SetStreamPitchScale(stream int64, pitch_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream) , gdc.ConstTypePtr(&pitch_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaybackPolyphonic.fnSetStreamPitchScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlaybackPolyphonic) IsStreamPlaying(stream int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&stream)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaybackPolyphonic.fnIsStreamPlaying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlaybackPolyphonic) StopStream(stream int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlaybackPolyphonic.fnStopStream), me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
