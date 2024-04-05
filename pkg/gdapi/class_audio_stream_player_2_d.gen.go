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

type ptrsForAudioStreamPlayer2DList struct {
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
  fnSetMaxDistance gdc.MethodBindPtr
  fnGetMaxDistance gdc.MethodBindPtr
  fnSetAttenuation gdc.MethodBindPtr
  fnGetAttenuation gdc.MethodBindPtr
  fnSetAreaMask gdc.MethodBindPtr
  fnGetAreaMask gdc.MethodBindPtr
  fnSetStreamPaused gdc.MethodBindPtr
  fnGetStreamPaused gdc.MethodBindPtr
  fnSetMaxPolyphony gdc.MethodBindPtr
  fnGetMaxPolyphony gdc.MethodBindPtr
  fnSetPanningStrength gdc.MethodBindPtr
  fnGetPanningStrength gdc.MethodBindPtr
  fnHasStreamPlayback gdc.MethodBindPtr
  fnGetStreamPlayback gdc.MethodBindPtr
}

var ptrsForAudioStreamPlayer2D ptrsForAudioStreamPlayer2DList

func initAudioStreamPlayer2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioStreamPlayer2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_stream")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2210767741))
  }
  {
    methodName := StringNameFromStr("get_stream")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 160907539))
  }
  {
    methodName := StringNameFromStr("set_volume_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volume_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_pitch_scale")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetPitchScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_pitch_scale")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetPitchScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("play")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnPlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1958160172))
  }
  {
    methodName := StringNameFromStr("seek")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSeek = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("stop")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("is_playing")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnIsPlaying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_playback_position")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetPlaybackPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_bus")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_bus")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_autoplay")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_autoplay_enabled")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnIsAutoplayEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_max_distance")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_max_distance")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_attenuation")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetAttenuation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_attenuation")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetAttenuation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_area_mask")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetAreaMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_area_mask")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetAreaMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_stream_paused")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetStreamPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_stream_paused")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetStreamPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_max_polyphony")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetMaxPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_polyphony")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetMaxPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_panning_strength")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnSetPanningStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_panning_strength")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetPanningStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("has_stream_playback")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnHasStreamPlayback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("get_stream_playback")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer2D.fnGetStreamPlayback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 210135309))
  }
}

type AudioStreamPlayer2D struct {
  Node2D
}

func (me *AudioStreamPlayer2D) BaseClass() string {
  return "AudioStreamPlayer2D"
}

func NewAudioStreamPlayer2D() *AudioStreamPlayer2D {
  str := StringNameFromStr("AudioStreamPlayer2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamPlayer2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioStreamPlayer2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPlayer2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlayer2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamPlayer2D) SetStream(stream AudioStream, )  {
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetStream), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetStream() AudioStream {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStream()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamPlayer2D) SetVolumeDb(volume_db float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&volume_db) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetVolumeDb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetVolumeDb() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetVolumeDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) SetPitchScale(pitch_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pitch_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetPitchScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetPitchScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetPitchScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) Play(from_position float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnPlay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) Seek(to_position float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSeek), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) Stop()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) IsPlaying() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnIsPlaying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) GetPlaybackPosition() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetPlaybackPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) SetBus(bus StringName, )  {
  cargs := []gdc.ConstTypePtr{bus.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetBus), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetBus() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetBus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamPlayer2D) SetAutoplay(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetAutoplay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) IsAutoplayEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnIsAutoplayEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) SetMaxDistance(pixels float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pixels) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetMaxDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetMaxDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetMaxDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) SetAttenuation(curve float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&curve) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetAttenuation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetAttenuation() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetAttenuation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) SetAreaMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetAreaMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetAreaMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetAreaMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) SetStreamPaused(pause bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pause) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetStreamPaused), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetStreamPaused() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetStreamPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) SetMaxPolyphony(max_polyphony int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_polyphony) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetMaxPolyphony), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetMaxPolyphony() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetMaxPolyphony), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) SetPanningStrength(panning_strength float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&panning_strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnSetPanningStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer2D) GetPanningStrength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetPanningStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) HasStreamPlayback() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnHasStreamPlayback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer2D) GetStreamPlayback() AudioStreamPlayback {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStreamPlayback()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer2D.fnGetStreamPlayback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AudioStreamPlayer2DFinishedSignalFn func()

func (me *AudioStreamPlayer2D) ConnectFinished(subs SignalSubscribers, fn AudioStreamPlayer2DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AudioStreamPlayer2D) DisconnectFinished(subs SignalSubscribers, fn AudioStreamPlayer2DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
