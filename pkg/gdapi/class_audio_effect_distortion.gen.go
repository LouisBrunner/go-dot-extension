// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectDistortion struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectDistortion) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectDistortion) BaseClass() string {
  return "AudioEffectDistortion"
}

type AudioEffectDistortionMode int
const (
  AudioEffectDistortionModeModeClip AudioEffectDistortionMode = 0
  AudioEffectDistortionModeModeAtan AudioEffectDistortionMode = 1
  AudioEffectDistortionModeModeLofi AudioEffectDistortionMode = 2
  AudioEffectDistortionModeModeOverdrive AudioEffectDistortionMode = 3
  AudioEffectDistortionModeModeWaveshape AudioEffectDistortionMode = 4
)

func  (me *AudioEffectDistortion) SetMode(mode AudioEffectDistortionMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectDistortion) GetMode() { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectDistortion) SetPreGain(pre_gain float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectDistortion) GetPreGain() { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectDistortion) SetKeepHfHz(keep_hf_hz float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectDistortion) GetKeepHfHz() { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectDistortion) SetDrive(drive float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectDistortion) GetDrive() { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectDistortion) SetPostGain(post_gain float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioEffectDistortion) GetPostGain() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
