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

func  (me *AudioStreamPlayer3D) SetStream(stream AudioStream, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetStream() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetVolumeDb(volume_db float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetVolumeDb() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetUnitSize(unit_size float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetUnitSize() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetMaxDb(max_db float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetMaxDb() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetPitchScale(pitch_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetPitchScale() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) Play(from_position float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) Seek(to_position float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) Stop() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) IsPlaying() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetPlaybackPosition() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetBus(bus StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetBus() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetAutoplay(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) IsAutoplayEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetMaxDistance(meters float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetMaxDistance() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetAreaMask(mask int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetAreaMask() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetEmissionAngle(degrees float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetEmissionAngle() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetEmissionAngleEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) IsEmissionAngleEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetEmissionAngleFilterAttenuationDb(db float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetEmissionAngleFilterAttenuationDb() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetAttenuationFilterCutoffHz(degrees float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetAttenuationFilterCutoffHz() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetAttenuationFilterDb(db float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetAttenuationFilterDb() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetAttenuationModel(model AudioStreamPlayer3DAttenuationModel, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetAttenuationModel() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetDopplerTracking(mode AudioStreamPlayer3DDopplerTracking, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetDopplerTracking() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetStreamPaused(pause bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetStreamPaused() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetMaxPolyphony(max_polyphony int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetMaxPolyphony() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) SetPanningStrength(panning_strength float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetPanningStrength() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) HasStreamPlayback() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer3D) GetStreamPlayback() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
