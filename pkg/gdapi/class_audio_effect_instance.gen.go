// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports
  "unsafe"






  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectInstance struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectInstance) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectInstance) BaseClass() string {
  return "AudioEffectInstance"
}



// Enums

func (me *AudioEffectInstance) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectInstance) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectInstance) XProcess(src_buffer unsafe.Pointer, dst_buffer *AudioFrame, frame_count int, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectInstance) XProcessSilence()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
