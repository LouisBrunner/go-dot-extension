// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioServer struct {
  obj gdc.ObjectPtr
}

func (me *AudioServer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioServer) BaseClass() string {
  return "AudioServer"
}



// Enums

type AudioServerSpeakerMode int
const (
  AudioServerSpeakerModeSpeakerModeStereo AudioServerSpeakerMode = 0
  AudioServerSpeakerModeSpeakerSurround31 AudioServerSpeakerMode = 1
  AudioServerSpeakerModeSpeakerSurround51 AudioServerSpeakerMode = 2
  AudioServerSpeakerModeSpeakerSurround71 AudioServerSpeakerMode = 3
)

func (me *AudioServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioServer) SetBusCount(amount int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusCount()  {
  panic("TODO: implement")
}

func  (me *AudioServer) RemoveBus(index int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) AddBus(at_position int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) MoveBus(index int, to_index int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetBusName(bus_idx int, name String, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusName(bus_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusIndex(bus_name StringName, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusChannels(bus_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetBusVolumeDb(bus_idx int, volume_db float32, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusVolumeDb(bus_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetBusSend(bus_idx int, send StringName, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusSend(bus_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetBusSolo(bus_idx int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) IsBusSolo(bus_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetBusMute(bus_idx int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) IsBusMute(bus_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetBusBypassEffects(bus_idx int, enable bool, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) IsBusBypassingEffects(bus_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) AddBusEffect(bus_idx int, effect AudioEffect, at_position int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) RemoveBusEffect(bus_idx int, effect_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusEffectCount(bus_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusEffect(bus_idx int, effect_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusEffectInstance(bus_idx int, effect_idx int, channel int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SwapBusEffects(bus_idx int, effect_idx int, by_effect_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetBusEffectEnabled(bus_idx int, effect_idx int, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) IsBusEffectEnabled(bus_idx int, effect_idx int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusPeakVolumeLeftDb(bus_idx int, channel int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetBusPeakVolumeRightDb(bus_idx int, channel int, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetPlaybackSpeedScale(scale float32, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetPlaybackSpeedScale()  {
  panic("TODO: implement")
}

func  (me *AudioServer) Lock()  {
  panic("TODO: implement")
}

func  (me *AudioServer) Unlock()  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetSpeakerMode()  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetMixRate()  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetOutputDeviceList()  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetOutputDevice()  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetOutputDevice(name String, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetTimeToNextMix()  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetTimeSinceLastMix()  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetOutputLatency()  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetInputDeviceList()  {
  panic("TODO: implement")
}

func  (me *AudioServer) GetInputDevice()  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetInputDevice(name String, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetBusLayout(bus_layout AudioBusLayout, )  {
  panic("TODO: implement")
}

func  (me *AudioServer) GenerateBusLayout()  {
  panic("TODO: implement")
}

func  (me *AudioServer) SetEnableTaggingUsedAudioStreams(enable bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
