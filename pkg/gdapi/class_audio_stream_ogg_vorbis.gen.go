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

type ptrsForAudioStreamOggVorbisList struct {
  fnLoadFromBuffer gdc.MethodBindPtr
  fnLoadFromFile gdc.MethodBindPtr
  fnSetPacketSequence gdc.MethodBindPtr
  fnGetPacketSequence gdc.MethodBindPtr
  fnSetLoop gdc.MethodBindPtr
  fnHasLoop gdc.MethodBindPtr
  fnSetLoopOffset gdc.MethodBindPtr
  fnGetLoopOffset gdc.MethodBindPtr
  fnSetBpm gdc.MethodBindPtr
  fnGetBpm gdc.MethodBindPtr
  fnSetBeatCount gdc.MethodBindPtr
  fnGetBeatCount gdc.MethodBindPtr
  fnSetBarBeats gdc.MethodBindPtr
  fnGetBarBeats gdc.MethodBindPtr
}

var ptrsForAudioStreamOggVorbis ptrsForAudioStreamOggVorbisList

func initAudioStreamOggVorbisPtrs(iface gdc.Interface) {

  className := StringNameFromStr("AudioStreamOggVorbis")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("load_from_buffer")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnLoadFromBuffer = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 354904730))
  }
  {
    methodName := StringNameFromStr("load_from_file")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnLoadFromFile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 797568536))
  }
  {
    methodName := StringNameFromStr("set_packet_sequence")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnSetPacketSequence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 438882457))
  }
  {
    methodName := StringNameFromStr("get_packet_sequence")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnGetPacketSequence = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2801636033))
  }
  {
    methodName := StringNameFromStr("set_loop")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnSetLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("has_loop")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnHasLoop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_loop_offset")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnSetLoopOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_loop_offset")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnGetLoopOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_bpm")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnSetBpm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_bpm")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnGetBpm = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_beat_count")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnSetBeatCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_beat_count")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnGetBeatCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_bar_beats")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnSetBarBeats = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_bar_beats")
    defer methodName.Destroy()
    ptrsForAudioStreamOggVorbis.fnGetBarBeats = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

type AudioStreamOggVorbis struct {
  AudioStream
}

func (me *AudioStreamOggVorbis) BaseClass() string {
  return "AudioStreamOggVorbis"
}

func NewAudioStreamOggVorbis() *AudioStreamOggVorbis {
  str := StringNameFromStr("AudioStreamOggVorbis") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamOggVorbis{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioStreamOggVorbis) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamOggVorbis) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamOggVorbis) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  AudioStreamOggVorbisLoadFromBuffer(buffer PackedByteArray, ) AudioStreamOggVorbis {
  cargs := []gdc.ConstTypePtr{buffer.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStreamOggVorbis()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnLoadFromBuffer), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  AudioStreamOggVorbisLoadFromFile(path String, ) AudioStreamOggVorbis {
  cargs := []gdc.ConstTypePtr{path.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAudioStreamOggVorbis()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnLoadFromFile), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamOggVorbis) SetPacketSequence(packet_sequence OggPacketSequence, )  {
  cargs := []gdc.ConstTypePtr{packet_sequence.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnSetPacketSequence), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetPacketSequence() OggPacketSequence {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewOggPacketSequence()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnGetPacketSequence), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamOggVorbis) SetLoop(enable bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnSetLoop), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) HasLoop() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnHasLoop), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamOggVorbis) SetLoopOffset(seconds float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnSetLoopOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetLoopOffset() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnGetLoopOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamOggVorbis) SetBpm(bpm float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bpm) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnSetBpm), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetBpm() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnGetBpm), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamOggVorbis) SetBeatCount(count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnSetBeatCount), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetBeatCount() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnGetBeatCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamOggVorbis) SetBarBeats(count int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnSetBarBeats), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetBarBeats() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForAudioStreamOggVorbis.fnGetBarBeats), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
