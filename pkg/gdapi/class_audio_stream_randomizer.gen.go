// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3197802065) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(stream.AsCTypePtr()), gdc.ConstTypePtr(&weight), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamRandomizer) MoveStream(index_from int, index_to int, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index_from), gdc.ConstTypePtr(&index_to), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamRandomizer) RemoveStream(index int, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamRandomizer) SetStream(index int, stream AudioStream, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 111075094) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(stream.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamRandomizer) GetStream(index int, ) AudioStream {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2739380747) // FIXME: should cache?
  var ret AudioStream
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamRandomizer) SetStreamProbabilityWeight(index int, weight float32, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream_probability_weight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&weight), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamRandomizer) GetStreamProbabilityWeight(index int, ) float32 {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream_probability_weight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamRandomizer) SetStreamsCount(count int, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_streams_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamRandomizer) GetStreamsCount() int {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_streams_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamRandomizer) SetRandomPitch(scale float32, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_random_pitch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamRandomizer) GetRandomPitch() float32 {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_random_pitch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamRandomizer) SetRandomVolumeOffsetDb(db_offset float32, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_random_volume_offset_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&db_offset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamRandomizer) GetRandomVolumeOffsetDb() float32 {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_random_volume_offset_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamRandomizer) SetPlaybackMode(mode AudioStreamRandomizerPlaybackMode, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_playback_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3950967023) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamRandomizer) GetPlaybackMode() AudioStreamRandomizerPlaybackMode {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playback_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3943055077) // FIXME: should cache?
  var ret AudioStreamRandomizerPlaybackMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
