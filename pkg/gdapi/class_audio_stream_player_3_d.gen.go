// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlayer3D struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamPlayer3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *AudioStreamPlayer3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlayer3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamPlayer3D) SetStream(stream AudioStream, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetStream()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetVolumeDb(volume_db float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetVolumeDb()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetUnitSize(unit_size float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetUnitSize()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetMaxDb(max_db float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetMaxDb()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetPitchScale(pitch_scale float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetPitchScale()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) Play(from_position float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) Seek(to_position float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) Stop()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) IsPlaying()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetPlaybackPosition()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetBus(bus StringName, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetBus()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetAutoplay(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) IsAutoplayEnabled()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetMaxDistance(meters float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetMaxDistance()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetAreaMask(mask int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetAreaMask()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetEmissionAngle(degrees float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetEmissionAngle()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetEmissionAngleEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) IsEmissionAngleEnabled()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetEmissionAngleFilterAttenuationDb(db float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetEmissionAngleFilterAttenuationDb()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetAttenuationFilterCutoffHz(degrees float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetAttenuationFilterCutoffHz()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetAttenuationFilterDb(db float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetAttenuationFilterDb()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetAttenuationModel(model AudioStreamPlayer3DAttenuationModel, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetAttenuationModel()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetDopplerTracking(mode AudioStreamPlayer3DDopplerTracking, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetDopplerTracking()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetStreamPaused(pause bool, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetStreamPaused()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetMaxPolyphony(max_polyphony int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetMaxPolyphony()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) SetPanningStrength(panning_strength float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetPanningStrength()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) HasStreamPlayback()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer3D) GetStreamPlayback()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
