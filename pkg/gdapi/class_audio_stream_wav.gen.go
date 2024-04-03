// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioStreamWAV struct {
  AudioStream
}

func (me *AudioStreamWAV) BaseClass() string {
  return "AudioStreamWAV"
}

func NewAudioStreamWAV() *AudioStreamWAV {
  str := StringNameFromStr("AudioStreamWAV") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioStreamWAV{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
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
  cargs := []gdc.ConstTypePtr{}
  var ret AudioStreamWAVFormat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
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
  cargs := []gdc.ConstTypePtr{}
  var ret AudioStreamWAVLoopMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AudioStreamWAV) SetLoopBegin(loop_begin int64, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_begin), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamWAV) GetLoopBegin() int64 {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loop_begin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamWAV) SetLoopEnd(loop_end int64, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_loop_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&loop_end), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamWAV) GetLoopEnd() int64 {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_loop_end")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamWAV) SetMixRate(mix_rate int64, )  {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mix_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mix_rate), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioStreamWAV) GetMixRate() int64 {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mix_rate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
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
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioStreamWAV) SaveToWav(path String, ) Error {
  classNameV := StringNameFromStr("AudioStreamWAV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("save_to_wav")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 166001499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(path.AsCTypePtr()), }
  var ret Error

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
