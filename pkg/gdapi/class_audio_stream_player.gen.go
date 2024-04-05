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

type ptrsForAudioStreamPlayerList struct {
  fnSetStream gdc.MethodBindPtr
  fnGetStream gdc.MethodBindPtr
  fnSetVolumeDb gdc.MethodBindPtr
  fnGetVolumeDb gdc.MethodBindPtr
  fnSetPitchScale gdc.MethodBindPtr
  fnGetPitchScale gdc.MethodBindPtr
  fnPlay gdc.MethodBindPtr
  fnSeek gdc.MethodBindPtr
  fnStop gdc.MethodBindPtr
  fnIsPlaying gdc.MethodBindPtr
  fnGetPlaybackPosition gdc.MethodBindPtr
  fnSetBus gdc.MethodBindPtr
  fnGetBus gdc.MethodBindPtr
  fnSetAutoplay gdc.MethodBindPtr
  fnIsAutoplayEnabled gdc.MethodBindPtr
  fnSetMixTarget gdc.MethodBindPtr
  fnGetMixTarget gdc.MethodBindPtr
  fnSetStreamPaused gdc.MethodBindPtr
  fnGetStreamPaused gdc.MethodBindPtr
  fnSetMaxPolyphony gdc.MethodBindPtr
  fnGetMaxPolyphony gdc.MethodBindPtr
  fnHasStreamPlayback gdc.MethodBindPtr
  fnGetStreamPlayback gdc.MethodBindPtr
}

var ptrsForAudioStreamPlayer ptrsForAudioStreamPlayerList

func initAudioStreamPlayerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioStreamPlayer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_stream")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnSetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2210767741))
  }
  {
    methodName := StringNameFromStr("get_stream")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnGetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 160907539))
  }
  {
    methodName := StringNameFromStr("set_volume_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnSetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volume_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnGetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_pitch_scale")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnSetPitchScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_pitch_scale")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnGetPitchScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("play")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnPlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1958160172))
  }
  {
    methodName := StringNameFromStr("seek")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnSeek = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("stop")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("is_playing")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnIsPlaying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_playback_position")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnGetPlaybackPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_bus")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnSetBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_bus")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnGetBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_autoplay")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnSetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_autoplay_enabled")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnIsAutoplayEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_mix_target")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnSetMixTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2300306138))
  }
  {
    methodName := StringNameFromStr("get_mix_target")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnGetMixTarget = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 172807476))
  }
  {
    methodName := StringNameFromStr("set_stream_paused")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnSetStreamPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_stream_paused")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnGetStreamPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_max_polyphony")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnSetMaxPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_polyphony")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnGetMaxPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("has_stream_playback")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnHasStreamPlayback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("get_stream_playback")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer.fnGetStreamPlayback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 210135309))
  }
}

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
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnSetStream), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetStream() AudioStream {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStream()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnGetStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamPlayer) SetVolumeDb(volume_db float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&volume_db) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnSetVolumeDb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetVolumeDb() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnGetVolumeDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) SetPitchScale(pitch_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pitch_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnSetPitchScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetPitchScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnGetPitchScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) Play(from_position float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnPlay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) Seek(to_position float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnSeek), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) Stop()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) IsPlaying() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnIsPlaying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) GetPlaybackPosition() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnGetPlaybackPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) SetBus(bus StringName, )  {
  cargs := []gdc.ConstTypePtr{bus.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnSetBus), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetBus() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnGetBus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamPlayer) SetAutoplay(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnSetAutoplay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) IsAutoplayEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnIsAutoplayEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) SetMixTarget(mix_target AudioStreamPlayerMixTarget, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix_target) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnSetMixTarget), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetMixTarget() AudioStreamPlayerMixTarget {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioStreamPlayerMixTarget

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnGetMixTarget), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AudioStreamPlayer) SetStreamPaused(pause bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pause) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnSetStreamPaused), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetStreamPaused() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnGetStreamPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) SetMaxPolyphony(max_polyphony int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_polyphony) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnSetMaxPolyphony), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer) GetMaxPolyphony() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnGetMaxPolyphony), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) HasStreamPlayback() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnHasStreamPlayback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer) GetStreamPlayback() AudioStreamPlayback {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStreamPlayback()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer.fnGetStreamPlayback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
