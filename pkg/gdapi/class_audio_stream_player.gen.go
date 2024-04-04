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

type AudioStreamPlayer struct {
  Node
}

func (me *AudioStreamPlayer) BaseClass() string {
  return "AudioStreamPlayer"
}

func NewAudioStreamPlayer() *AudioStreamPlayer {
  str := StringNameFromStr("AudioStreamPlayer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamPlayer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AudioStreamPlayerMixTarget int
const (
  AudioStreamPlayerMixTargetMixTargetStereo AudioStreamPlayerMixTarget = 0
  AudioStreamPlayerMixTargetMixTargetSurround AudioStreamPlayerMixTarget = 1
  AudioStreamPlayerMixTargetMixTargetCenter AudioStreamPlayerMixTarget = 2
)

func (me *AudioStreamPlayer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPlayer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlayer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamPlayer) SetStream(stream AudioStream, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2210767741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetStream() AudioStream {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 160907539) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStream()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamPlayer) SetVolumeDb(volume_db float64, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&volume_db) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetVolumeDb() float64 {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) SetPitchScale(pitch_scale float64, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pitch_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pitch_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetPitchScale() float64 {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pitch_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) Play(from_position float64, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1958160172) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) Seek(to_position float64, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("seek")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) Stop()  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) IsPlaying() bool {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) GetPlaybackPosition() float64 {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playback_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) SetBus(bus StringName, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{bus.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetBus() StringName {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamPlayer) SetAutoplay(enable bool, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) IsAutoplayEnabled() bool {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_autoplay_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) SetMixTarget(mix_target AudioStreamPlayerMixTarget, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mix_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2300306138) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix_target) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetMixTarget() AudioStreamPlayerMixTarget {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mix_target")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 172807476) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioStreamPlayerMixTarget

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AudioStreamPlayer) SetStreamPaused(pause bool, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pause) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetStreamPaused() bool {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) SetMaxPolyphony(max_polyphony int64, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_polyphony) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetMaxPolyphony() int64 {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) HasStreamPlayback() bool {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_stream_playback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) GetStreamPlayback() AudioStreamPlayback {
  classNameV := StringNameFromStr("AudioStreamPlayer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream_playback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 210135309) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStreamPlayback()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AudioStreamPlayerFinishedSignalFn func()

func (me *AudioStreamPlayer) ConnectFinished(subs SignalSubscribers, fn AudioStreamPlayerFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AudioStreamPlayer) DisconnectFinished(subs SignalSubscribers, fn AudioStreamPlayerFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
