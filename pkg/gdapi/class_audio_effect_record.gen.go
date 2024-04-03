// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectRecord struct {
  AudioEffect
}

func (me *AudioEffectRecord) BaseClass() string {
  return "AudioEffectRecord"
}

func NewAudioEffectRecord() *AudioEffectRecord {
  str := StringNameFromStr("AudioEffectRecord") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AudioEffectRecord{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *AudioEffectRecord) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectRecord) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectRecord) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AudioEffectRecord) SetRecordingActive(record bool, )  {
  classNameV := StringNameFromStr("AudioEffectRecord")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_recording_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&record), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectRecord) IsRecordingActive() bool {
  classNameV := StringNameFromStr("AudioEffectRecord")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_recording_active")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AudioEffectRecord) SetFormat(format AudioStreamWAVFormat, )  {
  classNameV := StringNameFromStr("AudioEffectRecord")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 60648488) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&format), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AudioEffectRecord) GetFormat() AudioStreamWAVFormat {
  classNameV := StringNameFromStr("AudioEffectRecord")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_format")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3151724922) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  var ret AudioStreamWAVFormat

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AudioEffectRecord) GetRecording() AudioStreamWAV {
  classNameV := StringNameFromStr("AudioEffectRecord")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_recording")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2964110865) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewAudioStreamWAV()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
