// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_buffer")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 354904730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(buffer.AsCTypePtr()), }
  ret := NewAudioStreamOggVorbis()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  AudioStreamOggVorbisLoadFromFile(path String, ) AudioStreamOggVorbis {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("load_from_file")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 797568536) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  ret := NewAudioStreamOggVorbis()

  giface.ObjectMethodBindPtrcall(methodPtr, nil, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamOggVorbis) SetPacketSequence(packet_sequence OggPacketSequence, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_packet_sequence")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 438882457) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(packet_sequence.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetPacketSequence() OggPacketSequence {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_packet_sequence")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2801636033) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewOggPacketSequence()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AudioStreamOggVorbis) SetLoop(enable bool, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) HasLoop() bool {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("has_loop")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamOggVorbis) SetLoopOffset(seconds float64, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetLoopOffset() float64 {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loop_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamOggVorbis) SetBpm(bpm float64, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bpm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bpm), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetBpm() float64 {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bpm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamOggVorbis) SetBeatCount(count int64, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_beat_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetBeatCount() int64 {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_beat_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamOggVorbis) SetBarBeats(count int64, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bar_beats")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamOggVorbis) GetBarBeats() int64 {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bar_beats")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
