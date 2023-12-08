// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlayer struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamPlayer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamPlayer) BaseClass() string {
  return "AudioStreamPlayer"
}

type AudioStreamPlayerMixTarget int
const (
  AudioStreamPlayerMixTargetMixTargetStereo AudioStreamPlayerMixTarget = 0
  AudioStreamPlayerMixTargetMixTargetSurround AudioStreamPlayerMixTarget = 1
  AudioStreamPlayerMixTargetMixTargetCenter AudioStreamPlayerMixTarget = 2
)

func  (me *AudioStreamPlayer) SetStream(stream AudioStream, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) GetStream() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) SetVolumeDb(volume_db float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) GetVolumeDb() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) SetPitchScale(pitch_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) GetPitchScale() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) Play(from_position float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) Seek(to_position float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) Stop() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) IsPlaying() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) GetPlaybackPosition() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) SetBus(bus StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) GetBus() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) SetAutoplay(enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) IsAutoplayEnabled() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) SetMixTarget(mix_target AudioStreamPlayerMixTarget, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) GetMixTarget() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) SetStreamPaused(pause bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) GetStreamPaused() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) SetMaxPolyphony(max_polyphony int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) GetMaxPolyphony() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) HasStreamPlayback() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamPlayer) GetStreamPlayback() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
