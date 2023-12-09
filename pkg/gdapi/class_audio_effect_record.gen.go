// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *AudioEffectRecord) IsRecordingActive()  {
  panic("TODO: implement")
}

func  (me *AudioEffectRecord) SetFormat(format AudioStreamWAVFormat, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectRecord) GetFormat()  {
  panic("TODO: implement")
}

func  (me *AudioEffectRecord) GetRecording()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
