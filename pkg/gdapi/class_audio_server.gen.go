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

type AudioServerSpeakerMode int
const (
  AudioServerSpeakerModeSpeakerModeStereo AudioServerSpeakerMode = 0
  AudioServerSpeakerModeSpeakerSurround31 AudioServerSpeakerMode = 1
  AudioServerSpeakerModeSpeakerSurround51 AudioServerSpeakerMode = 2
  AudioServerSpeakerModeSpeakerSurround71 AudioServerSpeakerMode = 3
)

func  (me *AudioServer) SetBusCount(amount int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusCount() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) RemoveBus(index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) AddBus(at_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) MoveBus(index int, to_index int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetBusName(bus_idx int, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusName(bus_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusIndex(bus_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusChannels(bus_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetBusVolumeDb(bus_idx int, volume_db float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusVolumeDb(bus_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetBusSend(bus_idx int, send StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusSend(bus_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetBusSolo(bus_idx int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) IsBusSolo(bus_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetBusMute(bus_idx int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) IsBusMute(bus_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetBusBypassEffects(bus_idx int, enable bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) IsBusBypassingEffects(bus_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) AddBusEffect(bus_idx int, effect AudioEffect, at_position int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) RemoveBusEffect(bus_idx int, effect_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusEffectCount(bus_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusEffect(bus_idx int, effect_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusEffectInstance(bus_idx int, effect_idx int, channel int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SwapBusEffects(bus_idx int, effect_idx int, by_effect_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetBusEffectEnabled(bus_idx int, effect_idx int, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) IsBusEffectEnabled(bus_idx int, effect_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusPeakVolumeLeftDb(bus_idx int, channel int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetBusPeakVolumeRightDb(bus_idx int, channel int, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetPlaybackSpeedScale(scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetPlaybackSpeedScale() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) Lock() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) Unlock() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetSpeakerMode() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetMixRate() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetOutputDeviceList() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetOutputDevice() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetOutputDevice(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetTimeToNextMix() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetTimeSinceLastMix() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetOutputLatency() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetInputDeviceList() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GetInputDevice() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetInputDevice(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetBusLayout(bus_layout AudioBusLayout, ) { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) GenerateBusLayout() { // TODO: return value
  // TODO: implement
}

func  (me *AudioServer) SetEnableTaggingUsedAudioStreams(enable bool, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
