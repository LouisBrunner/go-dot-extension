// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  AudioStreamPlayer3DAttenuationInverseDistance AudioStreamPlayer3DAttenuationModel = 0
  AudioStreamPlayer3DAttenuationInverseSquareDistance AudioStreamPlayer3DAttenuationModel = 1
  AudioStreamPlayer3DAttenuationLogarithmic AudioStreamPlayer3DAttenuationModel = 2
  AudioStreamPlayer3DAttenuationDisabled AudioStreamPlayer3DAttenuationModel = 3
)

type AudioStreamPlayer3DDopplerTracking int
const (
  AudioStreamPlayer3DDopplerTrackingDisabled AudioStreamPlayer3DDopplerTracking = 0
  AudioStreamPlayer3DDopplerTrackingIdleStep AudioStreamPlayer3DDopplerTracking = 1
  AudioStreamPlayer3DDopplerTrackingPhysicsStep AudioStreamPlayer3DDopplerTracking = 2
)
