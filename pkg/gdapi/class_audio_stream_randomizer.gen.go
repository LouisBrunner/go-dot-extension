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

type AudioStreamRandomizer struct {
  AudioStream
}

func (me *AudioStreamRandomizer) BaseClass() string {
  return "AudioStreamRandomizer"
}

func NewAudioStreamRandomizer() *AudioStreamRandomizer {
  str := StringNameFromStr("AudioStreamRandomizer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamRandomizer{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *AudioStreamRandomizer) AddStream(index int64, stream AudioStream, weight float64, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1892018854) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , stream.AsCTypePtr(), gdc.ConstTypePtr(&weight) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamRandomizer) MoveStream(index_from int64, index_to int64, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index_from) , gdc.ConstTypePtr(&index_to) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamRandomizer) RemoveStream(index int64, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamRandomizer) SetStream(index int64, stream AudioStream, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 111075094) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , stream.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamRandomizer) GetStream(index int64, ) AudioStream {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2739380747) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStream()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamRandomizer) SetStreamProbabilityWeight(index int64, weight float64, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream_probability_weight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , gdc.ConstTypePtr(&weight) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamRandomizer) GetStreamProbabilityWeight(index int64, ) float64 {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream_probability_weight")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamRandomizer) SetStreamsCount(count int64, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_streams_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamRandomizer) GetStreamsCount() int64 {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_streams_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamRandomizer) SetRandomPitch(scale float64, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_random_pitch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamRandomizer) GetRandomPitch() float64 {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_random_pitch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamRandomizer) SetRandomVolumeOffsetDb(db_offset float64, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_random_volume_offset_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&db_offset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamRandomizer) GetRandomVolumeOffsetDb() float64 {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_random_volume_offset_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamRandomizer) SetPlaybackMode(mode AudioStreamRandomizerPlaybackMode, )  {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_playback_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3950967023) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamRandomizer) GetPlaybackMode() AudioStreamRandomizerPlaybackMode {
  classNameV := StringNameFromStr("AudioStreamRandomizer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playback_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3943055077) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioStreamRandomizerPlaybackMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
