// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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
  AudioStreamPlayerMixTargetStereo AudioStreamPlayerMixTarget = 0
  AudioStreamPlayerMixTargetSurround AudioStreamPlayerMixTarget = 1
  AudioStreamPlayerMixTargetCenter AudioStreamPlayerMixTarget = 2
)
