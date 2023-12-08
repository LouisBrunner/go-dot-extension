// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  AudioStreamWAVFormatFormat8Bits AudioStreamWAVFormat = 0
  AudioStreamWAVFormatFormat16Bits AudioStreamWAVFormat = 1
  AudioStreamWAVFormatFormatImaAdpcm AudioStreamWAVFormat = 2
)

type AudioStreamWAVLoopMode int
const (
  AudioStreamWAVLoopModeLoopDisabled AudioStreamWAVLoopMode = 0
  AudioStreamWAVLoopModeLoopForward AudioStreamWAVLoopMode = 1
  AudioStreamWAVLoopModeLoopPingpong AudioStreamWAVLoopMode = 2
  AudioStreamWAVLoopModeLoopBackward AudioStreamWAVLoopMode = 3
)

func  (me *AudioStreamWAV) SetData(data PackedByteArray, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) GetData() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) SetFormat(format AudioStreamWAVFormat, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) GetFormat() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) SetLoopMode(loop_mode AudioStreamWAVLoopMode, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) GetLoopMode() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) SetLoopBegin(loop_begin int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) GetLoopBegin() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) SetLoopEnd(loop_end int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) GetLoopEnd() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) SetMixRate(mix_rate int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) GetMixRate() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) SetStereo(stereo bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) IsStereo() { // TODO: return value
  // TODO: implement
}

func  (me *AudioStreamWAV) SaveToWav(path String, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
