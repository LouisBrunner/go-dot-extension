// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  AudioEffectDistortionModeClip AudioEffectDistortionMode = 0
  AudioEffectDistortionModeAtan AudioEffectDistortionMode = 1
  AudioEffectDistortionModeLofi AudioEffectDistortionMode = 2
  AudioEffectDistortionModeOverdrive AudioEffectDistortionMode = 3
  AudioEffectDistortionModeWaveshape AudioEffectDistortionMode = 4
)
