// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *AudioStreamWAV) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamWAV) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamWAV) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioStreamWAV) SetData(data PackedByteArray, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2971499966) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(data.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamWAV) GetData() PackedByteArray {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_data")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2362200018) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamWAV) SetFormat(format AudioStreamWAVFormat, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 60648488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamWAV) GetFormat() AudioStreamWAVFormat {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3151724922) // FIXME: should cache?
  var ret AudioStreamWAVFormat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamWAV) SetLoopMode(loop_mode AudioStreamWAVLoopMode, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2444882972) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamWAV) GetLoopMode() AudioStreamWAVLoopMode {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loop_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 393560655) // FIXME: should cache?
  var ret AudioStreamWAVLoopMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamWAV) SetLoopBegin(loop_begin int, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_begin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamWAV) GetLoopBegin() int {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loop_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamWAV) SetLoopEnd(loop_end int, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_end), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamWAV) GetLoopEnd() int {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loop_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamWAV) SetMixRate(mix_rate int, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mix_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix_rate), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamWAV) GetMixRate() int {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mix_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamWAV) SetStereo(stereo bool, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stereo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stereo), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AudioStreamWAV) IsStereo() bool {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_stereo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioStreamWAV) SaveToWav(path String, ) Error {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_to_wav")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  var ret Error
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AudioStreamWAV) GetPropData() PackedByteArray {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) SetPropData(value PackedByteArray) {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) GetPropFormat() int {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) SetPropFormat(value int) {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) GetPropLoopMode() int {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) SetPropLoopMode(value int) {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) GetPropLoopBegin() int {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) SetPropLoopBegin(value int) {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) GetPropLoopEnd() int {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) SetPropLoopEnd(value int) {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) GetPropMixRate() int {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) SetPropMixRate(value int) {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) GetPropStereo() bool {
  panic("TODO: implement")
}

func (me *AudioStreamWAV) SetPropStereo(value bool) {
  panic("TODO: implement")
}