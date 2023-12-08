// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamWAV struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamWAV) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamWAV) BaseClass() string {
  return "AudioStreamWAV"
}

type AudioStreamWAVFormat int
const (
  AudioStreamWAVFormat8Bits AudioStreamWAVFormat = 0
  AudioStreamWAVFormat16Bits AudioStreamWAVFormat = 1
  AudioStreamWAVFormatImaAdpcm AudioStreamWAVFormat = 2
)

type AudioStreamWAVLoopMode int
const (
  AudioStreamWAVLoopDisabled AudioStreamWAVLoopMode = 0
  AudioStreamWAVLoopForward AudioStreamWAVLoopMode = 1
  AudioStreamWAVLoopPingpong AudioStreamWAVLoopMode = 2
  AudioStreamWAVLoopBackward AudioStreamWAVLoopMode = 3
)
