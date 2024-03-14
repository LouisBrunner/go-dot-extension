// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioStreamPlayer3D struct {
  Node3D
}

func (me *AudioStreamPlayer3D) BaseClass() string {
  return "AudioStreamPlayer3D"
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
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2210767741) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(stream.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetStream() AudioStream {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 160907539) // FIXME: should cache?
  var ret AudioStream
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetVolumeDb(volume_db float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&volume_db), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetVolumeDb() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetUnitSize(unit_size float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_unit_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&unit_size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetUnitSize() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_unit_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetMaxDb(max_db float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_db), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetMaxDb() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetPitchScale(pitch_scale float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_pitch_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pitch_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetPitchScale() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_pitch_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) Play(from_position float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("play")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1958160172) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&from_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) Seek(to_position float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("seek")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&to_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) Stop()  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) IsPlaying() bool {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_playing")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) GetPlaybackPosition() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playback_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 191475506) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetBus(bus StringName, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bus.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetBus() StringName {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetAutoplay(enable bool, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_autoplay")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) IsAutoplayEnabled() bool {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_autoplay_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetMaxDistance(meters float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&meters), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetMaxDistance() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetAreaMask(mask int, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_area_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mask), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetAreaMask() int {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_area_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetEmissionAngle(degrees float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetEmissionAngle() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetEmissionAngleEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_angle_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) IsEmissionAngleEnabled() bool {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_emission_angle_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetEmissionAngleFilterAttenuationDb(db float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_emission_angle_filter_attenuation_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&db), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetEmissionAngleFilterAttenuationDb() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_emission_angle_filter_attenuation_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetAttenuationFilterCutoffHz(degrees float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_attenuation_filter_cutoff_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetAttenuationFilterCutoffHz() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_attenuation_filter_cutoff_hz")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetAttenuationFilterDb(db float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_attenuation_filter_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&db), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetAttenuationFilterDb() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_attenuation_filter_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetAttenuationModel(model AudioStreamPlayer3DAttenuationModel, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_attenuation_model")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2988086229) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&model), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetAttenuationModel() AudioStreamPlayer3DAttenuationModel {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_attenuation_model")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3035106060) // FIXME: should cache?
  var ret AudioStreamPlayer3DAttenuationModel
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetDopplerTracking(mode AudioStreamPlayer3DDopplerTracking, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_doppler_tracking")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3968161450) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetDopplerTracking() AudioStreamPlayer3DDopplerTracking {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_doppler_tracking")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1702418664) // FIXME: should cache?
  var ret AudioStreamPlayer3DDopplerTracking
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetStreamPaused(pause bool, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stream_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&pause), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetStreamPaused() bool {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream_paused")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetMaxPolyphony(max_polyphony int, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&max_polyphony), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetMaxPolyphony() int {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_max_polyphony")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) SetPanningStrength(panning_strength float32, )  {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_panning_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&panning_strength), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamPlayer3D) GetPanningStrength() float32 {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_panning_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) HasStreamPlayback() bool {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_stream_playback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamPlayer3D) GetStreamPlayback() AudioStreamPlayback {
  classNameV := StringNameFromStr("AudioStreamPlayer3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stream_playback")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 210135309) // FIXME: should cache?
  var ret AudioStreamPlayback
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AudioStreamPlayer3DFinishedSignalFn func()

func (me *AudioStreamPlayer3D) ConnectFinished(subs SignalSubscribers, fn AudioStreamPlayer3DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *AudioStreamPlayer3D) DisconnectFinished(subs SignalSubscribers, fn AudioStreamPlayer3DFinishedSignalFn) {
  sig := StringNameFromStr("finished")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
