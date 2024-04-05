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

type ptrsForAudioStreamPlayer3DList struct {
  fnSetStream gdc.MethodBindPtr
  fnGetStream gdc.MethodBindPtr
  fnSetVolumeDb gdc.MethodBindPtr
  fnGetVolumeDb gdc.MethodBindPtr
  fnSetUnitSize gdc.MethodBindPtr
  fnGetUnitSize gdc.MethodBindPtr
  fnSetMaxDb gdc.MethodBindPtr
  fnGetMaxDb gdc.MethodBindPtr
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
  fnSetAreaMask gdc.MethodBindPtr
  fnGetAreaMask gdc.MethodBindPtr
  fnSetEmissionAngle gdc.MethodBindPtr
  fnGetEmissionAngle gdc.MethodBindPtr
  fnSetEmissionAngleEnabled gdc.MethodBindPtr
  fnIsEmissionAngleEnabled gdc.MethodBindPtr
  fnSetEmissionAngleFilterAttenuationDb gdc.MethodBindPtr
  fnGetEmissionAngleFilterAttenuationDb gdc.MethodBindPtr
  fnSetAttenuationFilterCutoffHz gdc.MethodBindPtr
  fnGetAttenuationFilterCutoffHz gdc.MethodBindPtr
  fnSetAttenuationFilterDb gdc.MethodBindPtr
  fnGetAttenuationFilterDb gdc.MethodBindPtr
  fnSetAttenuationModel gdc.MethodBindPtr
  fnGetAttenuationModel gdc.MethodBindPtr
  fnSetDopplerTracking gdc.MethodBindPtr
  fnGetDopplerTracking gdc.MethodBindPtr
  fnSetStreamPaused gdc.MethodBindPtr
  fnGetStreamPaused gdc.MethodBindPtr
  fnSetMaxPolyphony gdc.MethodBindPtr
  fnGetMaxPolyphony gdc.MethodBindPtr
  fnSetPanningStrength gdc.MethodBindPtr
  fnGetPanningStrength gdc.MethodBindPtr
  fnHasStreamPlayback gdc.MethodBindPtr
  fnGetStreamPlayback gdc.MethodBindPtr
}

var ptrsForAudioStreamPlayer3D ptrsForAudioStreamPlayer3DList

func initAudioStreamPlayer3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioStreamPlayer3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_stream")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2210767741))
  }
  {
    methodName := StringNameFromStr("get_stream")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetStream = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 160907539))
  }
  {
    methodName := StringNameFromStr("set_volume_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_volume_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_unit_size")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetUnitSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_unit_size")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetUnitSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_max_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetMaxDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_max_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetMaxDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_pitch_scale")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetPitchScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_pitch_scale")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetPitchScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("play")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnPlay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1958160172))
  }
  {
    methodName := StringNameFromStr("seek")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSeek = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("stop")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("is_playing")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnIsPlaying = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("get_playback_position")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetPlaybackPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 191475506))
  }
  {
    methodName := StringNameFromStr("set_bus")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
  }
  {
    methodName := StringNameFromStr("get_bus")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
  }
  {
    methodName := StringNameFromStr("set_autoplay")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetAutoplay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_autoplay_enabled")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnIsAutoplayEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("set_max_distance")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_max_distance")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetMaxDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_area_mask")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetAreaMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_area_mask")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetAreaMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_emission_angle")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetEmissionAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_emission_angle")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetEmissionAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_emission_angle_enabled")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetEmissionAngleEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_emission_angle_enabled")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnIsEmissionAngleEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_emission_angle_filter_attenuation_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetEmissionAngleFilterAttenuationDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_emission_angle_filter_attenuation_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetEmissionAngleFilterAttenuationDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_attenuation_filter_cutoff_hz")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetAttenuationFilterCutoffHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_attenuation_filter_cutoff_hz")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetAttenuationFilterCutoffHz = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_attenuation_filter_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetAttenuationFilterDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_attenuation_filter_db")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetAttenuationFilterDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_attenuation_model")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetAttenuationModel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2988086229))
  }
  {
    methodName := StringNameFromStr("get_attenuation_model")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetAttenuationModel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3035106060))
  }
  {
    methodName := StringNameFromStr("set_doppler_tracking")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetDopplerTracking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3968161450))
  }
  {
    methodName := StringNameFromStr("get_doppler_tracking")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetDopplerTracking = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1702418664))
  }
  {
    methodName := StringNameFromStr("set_stream_paused")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetStreamPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_stream_paused")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetStreamPaused = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_max_polyphony")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetMaxPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_max_polyphony")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetMaxPolyphony = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_panning_strength")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnSetPanningStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_panning_strength")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetPanningStrength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("has_stream_playback")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnHasStreamPlayback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("get_stream_playback")
    defer methodName.Destroy()
    ptrsForAudioStreamPlayer3D.fnGetStreamPlayback = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 210135309))
  }
}

type AudioStreamPlayer3D struct {
  Node3D
}

func (me *AudioStreamPlayer3D) BaseClass() string {
  return "AudioStreamPlayer3D"
}

func NewAudioStreamPlayer3D() *AudioStreamPlayer3D {
  str := StringNameFromStr("AudioStreamPlayer3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamPlayer3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AudioStreamPlayer3DAttenuationModel int
const (
  AudioStreamPlayer3DAttenuationModelAttenuationInverseDistance AudioStreamPlayer3DAttenuationModel = 0
  AudioStreamPlayer3DAttenuationModelAttenuationInverseSquareDistance AudioStreamPlayer3DAttenuationModel = 1
  AudioStreamPlayer3DAttenuationModelAttenuationLogarithmic AudioStreamPlayer3DAttenuationModel = 2
  AudioStreamPlayer3DAttenuationModelAttenuationDisabled AudioStreamPlayer3DAttenuationModel = 3
)

type AudioStreamPlayer3DDopplerTracking int
const (
  AudioStreamPlayer3DDopplerTrackingDopplerTrackingDisabled AudioStreamPlayer3DDopplerTracking = 0
  AudioStreamPlayer3DDopplerTrackingDopplerTrackingIdleStep AudioStreamPlayer3DDopplerTracking = 1
  AudioStreamPlayer3DDopplerTrackingDopplerTrackingPhysicsStep AudioStreamPlayer3DDopplerTracking = 2
)

func (me *AudioStreamPlayer3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPlayer3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlayer3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamPlayer3D) SetStream(stream AudioStream, )  {
  cargs := []gdc.ConstTypePtr{stream.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetStream), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetStream() AudioStream {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStream()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetStream), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamPlayer3D) SetVolumeDb(volume_db float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&volume_db) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetVolumeDb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetVolumeDb() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetVolumeDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetUnitSize(unit_size float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unit_size) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetUnitSize), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetUnitSize() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetUnitSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetMaxDb(max_db float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_db) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetMaxDb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetMaxDb() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetMaxDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetPitchScale(pitch_scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pitch_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetPitchScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetPitchScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetPitchScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) Play(from_position float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnPlay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) Seek(to_position float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSeek), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) Stop()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) IsPlaying() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnIsPlaying), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) GetPlaybackPosition() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetPlaybackPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetBus(bus StringName, )  {
  cargs := []gdc.ConstTypePtr{bus.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetBus), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetBus() StringName {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetBus), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamPlayer3D) SetAutoplay(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetAutoplay), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) IsAutoplayEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnIsAutoplayEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetMaxDistance(meters float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&meters) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetMaxDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetMaxDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetMaxDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetAreaMask(mask int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetAreaMask), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetAreaMask() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetAreaMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetEmissionAngle(degrees float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetEmissionAngle), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetEmissionAngle() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetEmissionAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetEmissionAngleEnabled(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetEmissionAngleEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) IsEmissionAngleEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnIsEmissionAngleEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetEmissionAngleFilterAttenuationDb(db float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&db) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetEmissionAngleFilterAttenuationDb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetEmissionAngleFilterAttenuationDb() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetEmissionAngleFilterAttenuationDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetAttenuationFilterCutoffHz(degrees float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetAttenuationFilterCutoffHz), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetAttenuationFilterCutoffHz() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetAttenuationFilterCutoffHz), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetAttenuationFilterDb(db float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&db) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetAttenuationFilterDb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetAttenuationFilterDb() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetAttenuationFilterDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetAttenuationModel(model AudioStreamPlayer3DAttenuationModel, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&model) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetAttenuationModel), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetAttenuationModel() AudioStreamPlayer3DAttenuationModel {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioStreamPlayer3DAttenuationModel

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetAttenuationModel), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AudioStreamPlayer3D) SetDopplerTracking(mode AudioStreamPlayer3DDopplerTracking, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetDopplerTracking), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetDopplerTracking() AudioStreamPlayer3DDopplerTracking {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioStreamPlayer3DDopplerTracking

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetDopplerTracking), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AudioStreamPlayer3D) SetStreamPaused(pause bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pause) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetStreamPaused), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetStreamPaused() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetStreamPaused), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetMaxPolyphony(max_polyphony int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_polyphony) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetMaxPolyphony), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetMaxPolyphony() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetMaxPolyphony), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) SetPanningStrength(panning_strength float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&panning_strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnSetPanningStrength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamPlayer3D) GetPanningStrength() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetPanningStrength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) HasStreamPlayback() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnHasStreamPlayback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamPlayer3D) GetStreamPlayback() AudioStreamPlayback {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStreamPlayback()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamPlayer3D.fnGetStreamPlayback), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AudioStreamPlayer3DFinishedSignalFn func()

func (me *AudioStreamPlayer3D) ConnectFinished(subs SignalSubscribers, fn AudioStreamPlayer3DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AudioStreamPlayer3D) DisconnectFinished(subs SignalSubscribers, fn AudioStreamPlayer3DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
