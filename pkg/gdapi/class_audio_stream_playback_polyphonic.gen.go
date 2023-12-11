// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func  (me *AudioStreamPlaybackPolyphonic) PlayStream(stream AudioStream, from_offset float32, volume_db float32, pitch_scale float32, ) int {
  classNameV := StringNameFromStr("AudioStreamPlaybackPolyphonic")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3792189967) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(stream.AsCTypePtr()), gdc.ConstTypePtr(&from_offset), gdc.ConstTypePtr(&volume_db), gdc.ConstTypePtr(&pitch_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlaybackPolyphonic) SetStreamVolume(stream int, volume_db float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlaybackPolyphonic")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream_volume")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream), gdc.ConstTypePtr(&volume_db), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlaybackPolyphonic) SetStreamPitchScale(stream int, pitch_scale float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlaybackPolyphonic")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream_pitch_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream), gdc.ConstTypePtr(&pitch_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlaybackPolyphonic) IsStreamPlaying(stream int, ) bool {
  classNameV := StringNameFromStr("AudioStreamPlaybackPolyphonic")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_stream_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlaybackPolyphonic) StopStream(stream int, )  {
  classNameV := StringNameFromStr("AudioStreamPlaybackPolyphonic")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stream), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
