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



// Enums

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

func (me *AudioStreamWAV) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamWAV) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamWAV) SetData(data PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) GetData()  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) SetFormat(format AudioStreamWAVFormat, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) GetFormat()  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) SetLoopMode(loop_mode AudioStreamWAVLoopMode, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) GetLoopMode()  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) SetLoopBegin(loop_begin int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) GetLoopBegin()  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) SetLoopEnd(loop_end int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) GetLoopEnd()  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) SetMixRate(mix_rate int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) GetMixRate()  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) SetStereo(stereo bool, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) IsStereo()  {
  panic("TODO: implement")
}

func  (me *AudioStreamWAV) SaveToWav(path String, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
