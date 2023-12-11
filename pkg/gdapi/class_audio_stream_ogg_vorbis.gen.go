// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioStreamOggVorbis struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamOggVorbis) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamOggVorbis) BaseClass() string {
  return "AudioStreamOggVorbis"
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
  var ret OggPacketSequence
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamOggVorbis) SetLoopOffset(seconds float32, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&seconds), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamOggVorbis) GetLoopOffset() float32 {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loop_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamOggVorbis) SetBpm(bpm float32, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bpm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bpm), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamOggVorbis) GetBpm() float32 {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bpm")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamOggVorbis) SetBeatCount(count int, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_beat_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamOggVorbis) GetBeatCount() int {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_beat_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamOggVorbis) SetBarBeats(count int, )  {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bar_beats")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamOggVorbis) GetBarBeats() int {
  classNameV := StringNameFromStr("AudioStreamOggVorbis")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bar_beats")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
