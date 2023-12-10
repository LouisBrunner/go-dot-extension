// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AudioEffectRecord struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectRecord) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectRecord) BaseClass() string {
  return "AudioEffectRecord"
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret AudioStreamWAVFormat
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AudioEffectRecord) GetRecording() AudioStreamWAV {
  classNameV := StringNameFromStr("AudioEffectRecord")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_recording")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2964110865) // FIXME: should cache?
  var ret AudioStreamWAV
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AudioEffectRecord) GetPropFormat() int {
  panic("TODO: implement")
}

func (me *AudioEffectRecord) SetPropFormat(value int) {
  panic("TODO: implement")
}