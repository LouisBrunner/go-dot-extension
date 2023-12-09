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
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) GetStream()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) SetVolumeDb(volume_db float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) GetVolumeDb()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) SetPitchScale(pitch_scale float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) GetPitchScale()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) Play(from_position float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) Seek(to_position float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) Stop()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) IsPlaying()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) GetPlaybackPosition()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) SetBus(bus StringName, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) GetBus()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) SetAutoplay(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) IsAutoplayEnabled()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) SetMixTarget(mix_target AudioStreamPlayerMixTarget, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) GetMixTarget()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) SetStreamPaused(pause bool, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) GetStreamPaused()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) SetMaxPolyphony(max_polyphony int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) GetMaxPolyphony()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) HasStreamPlayback()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlayer) GetStreamPlayback()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
