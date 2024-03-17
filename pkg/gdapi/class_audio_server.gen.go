// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioServer struct {
  Object
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

func (me *AudioServer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioServer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioServer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioServer) SetBusCount(amount int, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) GetBusCount() int {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) RemoveBus(index int, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) AddBus(at_position int, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1025054187) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&at_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) MoveBus(index int, to_index int, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_bus")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index), gdc.ConstTypePtr(&to_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) SetBusName(bus_idx int, name String, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) GetBusName(bus_idx int, ) String {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetBusIndex(bus_name StringName, ) int {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2458036349) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bus_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetBusChannels(bus_idx int, ) int {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_channels")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SetBusVolumeDb(bus_idx int, volume_db float32, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&volume_db), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) GetBusVolumeDb(bus_idx int, ) float32 {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_volume_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SetBusSend(bus_idx int, send StringName, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus_send")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(send.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) GetBusSend(bus_idx int, ) StringName {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_send")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SetBusSolo(bus_idx int, enable bool, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus_solo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) IsBusSolo(bus_idx int, ) bool {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_bus_solo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SetBusMute(bus_idx int, enable bool, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus_mute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) IsBusMute(bus_idx int, ) bool {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_bus_mute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SetBusBypassEffects(bus_idx int, enable bool, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus_bypass_effects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) IsBusBypassingEffects(bus_idx int, ) bool {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_bus_bypassing_effects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) AddBusEffect(bus_idx int, effect AudioEffect, at_position int, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_bus_effect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4068819785) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(effect.AsCTypePtr()), gdc.ConstTypePtr(&at_position), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) RemoveBusEffect(bus_idx int, effect_idx int, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("remove_bus_effect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&effect_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) GetBusEffectCount(bus_idx int, ) int {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_effect_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3744713108) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetBusEffect(bus_idx int, effect_idx int, ) AudioEffect {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_effect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 726064442) // FIXME: should cache?
  var ret AudioEffect
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&effect_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetBusEffectInstance(bus_idx int, effect_idx int, channel int, ) AudioEffectInstance {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_effect_instance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1829771234) // FIXME: should cache?
  var ret AudioEffectInstance
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&effect_idx), gdc.ConstTypePtr(&channel), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SwapBusEffects(bus_idx int, effect_idx int, by_effect_idx int, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("swap_bus_effects")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1649997291) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&effect_idx), gdc.ConstTypePtr(&by_effect_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) SetBusEffectEnabled(bus_idx int, effect_idx int, enabled bool, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus_effect_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1383440665) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&effect_idx), gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) IsBusEffectEnabled(bus_idx int, effect_idx int, ) bool {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_bus_effect_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2522259332) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&effect_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetBusPeakVolumeLeftDb(bus_idx int, channel int, ) float32 {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_peak_volume_left_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&channel), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetBusPeakVolumeRightDb(bus_idx int, channel int, ) float32 {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bus_peak_volume_right_db")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3085491603) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx), gdc.ConstTypePtr(&channel), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SetPlaybackSpeedScale(scale float32, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_playback_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) GetPlaybackSpeedScale() float32 {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_playback_speed_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) Lock()  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("lock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) Unlock()  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unlock")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) GetSpeakerMode() AudioServerSpeakerMode {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_speaker_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2549190337) // FIXME: should cache?
  var ret AudioServerSpeakerMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetMixRate() float32 {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mix_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetOutputDeviceList() PackedStringArray {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_device_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetOutputDevice() String {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SetOutputDevice(name String, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_output_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) GetTimeToNextMix() float32 {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_to_next_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetTimeSinceLastMix() float32 {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_time_since_last_mix")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetOutputLatency() float32 {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_output_latency")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetInputDeviceList() PackedStringArray {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_device_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2981934095) // FIXME: should cache?
  var ret PackedStringArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) GetInputDevice() String {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2841200299) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SetInputDevice(name String, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_device")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) SetBusLayout(bus_layout AudioBusLayout, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bus_layout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3319058824) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bus_layout.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioServer) GenerateBusLayout() AudioBusLayout {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("generate_bus_layout")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3769973890) // FIXME: should cache?
  var ret AudioBusLayout
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioServer) SetEnableTaggingUsedAudioStreams(enable bool, )  {
  classNameV := StringNameFromStr("AudioServer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_tagging_used_audio_streams")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AudioServerBusLayoutChangedSignalFn func()

func (me *AudioServer) ConnectBusLayoutChanged(subs SignalSubscribers, fn AudioServerBusLayoutChangedSignalFn) {
  sig := StringNameFromStr("bus_layout_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AudioServer) DisconnectBusLayoutChanged(subs SignalSubscribers, fn AudioServerBusLayoutChangedSignalFn) {
  sig := StringNameFromStr("bus_layout_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type AudioServerBusRenamedSignalFn func(bus_index int, old_name StringName, new_name StringName, )

func (me *AudioServer) ConnectBusRenamed(subs SignalSubscribers, fn AudioServerBusRenamedSignalFn) {
  sig := StringNameFromStr("bus_renamed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AudioServer) DisconnectBusRenamed(subs SignalSubscribers, fn AudioServerBusRenamedSignalFn) {
  sig := StringNameFromStr("bus_renamed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
