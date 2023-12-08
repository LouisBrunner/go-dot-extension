// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamRandomizer struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamRandomizer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamRandomizer) BaseClass() string {
  return "AudioStreamRandomizer"
}

type AudioStreamRandomizerPlaybackMode int
const (
  AudioStreamRandomizerPlaybackRandomNoRepeats AudioStreamRandomizerPlaybackMode = 0
  AudioStreamRandomizerPlaybackRandom AudioStreamRandomizerPlaybackMode = 1
  AudioStreamRandomizerPlaybackSequential AudioStreamRandomizerPlaybackMode = 2
)
