// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForAudioServerList struct {
  fnSetBusCount gdc.MethodBindPtr
  fnGetBusCount gdc.MethodBindPtr
  fnRemoveBus gdc.MethodBindPtr
  fnAddBus gdc.MethodBindPtr
  fnMoveBus gdc.MethodBindPtr
  fnSetBusName gdc.MethodBindPtr
  fnGetBusName gdc.MethodBindPtr
  fnGetBusIndex gdc.MethodBindPtr
  fnGetBusChannels gdc.MethodBindPtr
  fnSetBusVolumeDb gdc.MethodBindPtr
  fnGetBusVolumeDb gdc.MethodBindPtr
  fnSetBusSend gdc.MethodBindPtr
  fnGetBusSend gdc.MethodBindPtr
  fnSetBusSolo gdc.MethodBindPtr
  fnIsBusSolo gdc.MethodBindPtr
  fnSetBusMute gdc.MethodBindPtr
  fnIsBusMute gdc.MethodBindPtr
  fnSetBusBypassEffects gdc.MethodBindPtr
  fnIsBusBypassingEffects gdc.MethodBindPtr
  fnAddBusEffect gdc.MethodBindPtr
  fnRemoveBusEffect gdc.MethodBindPtr
  fnGetBusEffectCount gdc.MethodBindPtr
  fnGetBusEffect gdc.MethodBindPtr
  fnGetBusEffectInstance gdc.MethodBindPtr
  fnSwapBusEffects gdc.MethodBindPtr
  fnSetBusEffectEnabled gdc.MethodBindPtr
  fnIsBusEffectEnabled gdc.MethodBindPtr
  fnGetBusPeakVolumeLeftDb gdc.MethodBindPtr
  fnGetBusPeakVolumeRightDb gdc.MethodBindPtr
  fnSetPlaybackSpeedScale gdc.MethodBindPtr
  fnGetPlaybackSpeedScale gdc.MethodBindPtr
  fnLock gdc.MethodBindPtr
  fnUnlock gdc.MethodBindPtr
  fnGetSpeakerMode gdc.MethodBindPtr
  fnGetMixRate gdc.MethodBindPtr
  fnGetOutputDeviceList gdc.MethodBindPtr
  fnGetOutputDevice gdc.MethodBindPtr
  fnSetOutputDevice gdc.MethodBindPtr
  fnGetTimeToNextMix gdc.MethodBindPtr
  fnGetTimeSinceLastMix gdc.MethodBindPtr
  fnGetOutputLatency gdc.MethodBindPtr
  fnGetInputDeviceList gdc.MethodBindPtr
  fnGetInputDevice gdc.MethodBindPtr
  fnSetInputDevice gdc.MethodBindPtr
  fnSetBusLayout gdc.MethodBindPtr
  fnGenerateBusLayout gdc.MethodBindPtr
  fnSetEnableTaggingUsedAudioStreams gdc.MethodBindPtr
}

var ptrsForAudioServer ptrsForAudioServerList

func initAudioServerPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioServer")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_bus_count")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetBusCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_bus_count")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("remove_bus")
    defer methodName.Destroy()
    ptrsForAudioServer.fnRemoveBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("add_bus")
    defer methodName.Destroy()
    ptrsForAudioServer.fnAddBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1025054187))
  }
  {
    methodName := StringNameFromStr("move_bus")
    defer methodName.Destroy()
    ptrsForAudioServer.fnMoveBus = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("set_bus_name")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetBusName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
  }
  {
    methodName := StringNameFromStr("get_bus_name")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
  }
  {
    methodName := StringNameFromStr("get_bus_index")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2458036349))
  }
  {
    methodName := StringNameFromStr("get_bus_channels")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusChannels = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("set_bus_volume_db")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetBusVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("get_bus_volume_db")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusVolumeDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("set_bus_send")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetBusSend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780747571))
  }
  {
    methodName := StringNameFromStr("get_bus_send")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusSend = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
  }
  {
    methodName := StringNameFromStr("set_bus_solo")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetBusSolo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("is_bus_solo")
    defer methodName.Destroy()
    ptrsForAudioServer.fnIsBusSolo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_bus_mute")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetBusMute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("is_bus_mute")
    defer methodName.Destroy()
    ptrsForAudioServer.fnIsBusMute = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_bus_bypass_effects")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetBusBypassEffects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("is_bus_bypassing_effects")
    defer methodName.Destroy()
    ptrsForAudioServer.fnIsBusBypassingEffects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("add_bus_effect")
    defer methodName.Destroy()
    ptrsForAudioServer.fnAddBusEffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4068819785))
  }
  {
    methodName := StringNameFromStr("remove_bus_effect")
    defer methodName.Destroy()
    ptrsForAudioServer.fnRemoveBusEffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("get_bus_effect_count")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusEffectCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3744713108))
  }
  {
    methodName := StringNameFromStr("get_bus_effect")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusEffect = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 726064442))
  }
  {
    methodName := StringNameFromStr("get_bus_effect_instance")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusEffectInstance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1829771234))
  }
  {
    methodName := StringNameFromStr("swap_bus_effects")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSwapBusEffects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1649997291))
  }
  {
    methodName := StringNameFromStr("set_bus_effect_enabled")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetBusEffectEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1383440665))
  }
  {
    methodName := StringNameFromStr("is_bus_effect_enabled")
    defer methodName.Destroy()
    ptrsForAudioServer.fnIsBusEffectEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2522259332))
  }
  {
    methodName := StringNameFromStr("get_bus_peak_volume_left_db")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusPeakVolumeLeftDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
  }
  {
    methodName := StringNameFromStr("get_bus_peak_volume_right_db")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetBusPeakVolumeRightDb = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3085491603))
  }
  {
    methodName := StringNameFromStr("set_playback_speed_scale")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetPlaybackSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_playback_speed_scale")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetPlaybackSpeedScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("lock")
    defer methodName.Destroy()
    ptrsForAudioServer.fnLock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("unlock")
    defer methodName.Destroy()
    ptrsForAudioServer.fnUnlock = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
  }
  {
    methodName := StringNameFromStr("get_speaker_mode")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetSpeakerMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2549190337))
  }
  {
    methodName := StringNameFromStr("get_mix_rate")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetMixRate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_output_device_list")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetOutputDeviceList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
  }
  {
    methodName := StringNameFromStr("get_output_device")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetOutputDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
  }
  {
    methodName := StringNameFromStr("set_output_device")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetOutputDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("get_time_to_next_mix")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetTimeToNextMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_time_since_last_mix")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetTimeSinceLastMix = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_output_latency")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetOutputLatency = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("get_input_device_list")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetInputDeviceList = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2981934095))
  }
  {
    methodName := StringNameFromStr("get_input_device")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGetInputDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2841200299))
  }
  {
    methodName := StringNameFromStr("set_input_device")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetInputDevice = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
  }
  {
    methodName := StringNameFromStr("set_bus_layout")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetBusLayout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3319058824))
  }
  {
    methodName := StringNameFromStr("generate_bus_layout")
    defer methodName.Destroy()
    ptrsForAudioServer.fnGenerateBusLayout = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3769973890))
  }
  {
    methodName := StringNameFromStr("set_enable_tagging_used_audio_streams")
    defer methodName.Destroy()
    ptrsForAudioServer.fnSetEnableTaggingUsedAudioStreams = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
}

type AudioServer struct {
  Object
}

func (me *AudioServer) BaseClass() string {
  return "AudioServer"
}

func NewAudioServer() *AudioServer {
  str := StringNameFromStr("AudioServer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioServer{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *AudioServer) SetBusCount(amount int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetBusCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) GetBusCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) RemoveBus(index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnRemoveBus), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) AddBus(at_position int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&at_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnAddBus), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) MoveBus(index int64, to_index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&index) , gdc.ConstTypePtr(&to_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnMoveBus), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) SetBusName(bus_idx int64, name String, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetBusName), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) GetBusName(bus_idx int64, ) String {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&bus_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioServer) GetBusIndex(bus_name StringName, ) int64 {
  cargs := []gdc.ConstTypePtr{bus_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) GetBusChannels(bus_idx int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&bus_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusChannels), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) SetBusVolumeDb(bus_idx int64, volume_db float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&volume_db) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetBusVolumeDb), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) GetBusVolumeDb(bus_idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&bus_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusVolumeDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) SetBusSend(bus_idx int64, send StringName, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , send.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetBusSend), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) GetBusSend(bus_idx int64, ) StringName {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&bus_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusSend), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioServer) SetBusSolo(bus_idx int64, enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetBusSolo), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) IsBusSolo(bus_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&bus_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnIsBusSolo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) SetBusMute(bus_idx int64, enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetBusMute), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) IsBusMute(bus_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&bus_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnIsBusMute), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) SetBusBypassEffects(bus_idx int64, enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetBusBypassEffects), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) IsBusBypassingEffects(bus_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&bus_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnIsBusBypassingEffects), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) AddBusEffect(bus_idx int64, effect AudioEffect, at_position int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , effect.AsCTypePtr(), gdc.ConstTypePtr(&at_position) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnAddBusEffect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) RemoveBusEffect(bus_idx int64, effect_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&effect_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnRemoveBusEffect), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) GetBusEffectCount(bus_idx int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&bus_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusEffectCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) GetBusEffect(bus_idx int64, effect_idx int64, ) AudioEffect {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&effect_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioEffect()
  pinner.Pin(&bus_idx)
  pinner.Pin(&effect_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusEffect), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioServer) GetBusEffectInstance(bus_idx int64, effect_idx int64, channel int64, ) AudioEffectInstance {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&effect_idx) , gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioEffectInstance()
  pinner.Pin(&bus_idx)
  pinner.Pin(&effect_idx)
  pinner.Pin(&channel)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusEffectInstance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioServer) SwapBusEffects(bus_idx int64, effect_idx int64, by_effect_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&effect_idx) , gdc.ConstTypePtr(&by_effect_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSwapBusEffects), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) SetBusEffectEnabled(bus_idx int64, effect_idx int64, enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&effect_idx) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetBusEffectEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) IsBusEffectEnabled(bus_idx int64, effect_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&effect_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&bus_idx)
  pinner.Pin(&effect_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnIsBusEffectEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) GetBusPeakVolumeLeftDb(bus_idx int64, channel int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&bus_idx)
  pinner.Pin(&channel)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusPeakVolumeLeftDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) GetBusPeakVolumeRightDb(bus_idx int64, channel int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bus_idx) , gdc.ConstTypePtr(&channel) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&bus_idx)
  pinner.Pin(&channel)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetBusPeakVolumeRightDb), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) SetPlaybackSpeedScale(scale float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetPlaybackSpeedScale), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) GetPlaybackSpeedScale() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetPlaybackSpeedScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) Lock()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnLock), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) Unlock()  {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnUnlock), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) GetSpeakerMode() AudioServerSpeakerMode {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AudioServerSpeakerMode

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetSpeakerMode), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AudioServer) GetMixRate() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetMixRate), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) GetOutputDeviceList() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetOutputDeviceList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioServer) GetOutputDevice() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetOutputDevice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioServer) SetOutputDevice(name String, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetOutputDevice), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) GetTimeToNextMix() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetTimeToNextMix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) GetTimeSinceLastMix() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetTimeSinceLastMix), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) GetOutputLatency() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetOutputLatency), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioServer) GetInputDeviceList() PackedStringArray {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedStringArray()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetInputDeviceList), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioServer) GetInputDevice() String {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGetInputDevice), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioServer) SetInputDevice(name String, )  {
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetInputDevice), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) SetBusLayout(bus_layout AudioBusLayout, )  {
  cargs := []gdc.ConstTypePtr{bus_layout.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetBusLayout), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioServer) GenerateBusLayout() AudioBusLayout {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioBusLayout()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnGenerateBusLayout), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioServer) SetEnableTaggingUsedAudioStreams(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioServer.fnSetEnableTaggingUsedAudioStreams), me.obj, unsafe.SliceData(cargs), nil)

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
