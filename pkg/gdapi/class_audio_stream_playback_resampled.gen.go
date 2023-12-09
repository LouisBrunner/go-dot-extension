// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPlaybackResampled struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamPlaybackResampled) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamPlaybackResampled) BaseClass() string {
  return "AudioStreamPlaybackResampled"
}



// Enums

func (me *AudioStreamPlaybackResampled) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPlaybackResampled) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPlaybackResampled) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamPlaybackResampled) XMixResampled(dst_buffer *AudioFrame, frame_count int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlaybackResampled) XGetStreamSamplingRate()  {
  panic("TODO: implement")
}

func  (me *AudioStreamPlaybackResampled) BeginResample()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
