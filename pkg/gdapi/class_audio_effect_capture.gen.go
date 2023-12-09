// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectCapture struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectCapture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectCapture) BaseClass() string {
  return "AudioEffectCapture"
}



// Enums

func (me *AudioEffectCapture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectCapture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectCapture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectCapture) CanGetBuffer(frames int, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectCapture) GetBuffer(frames int, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectCapture) ClearBuffer()  {
  panic("TODO: implement")
}

func  (me *AudioEffectCapture) SetBufferLength(buffer_length_seconds float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectCapture) GetBufferLength()  {
  panic("TODO: implement")
}

func  (me *AudioEffectCapture) GetFramesAvailable()  {
  panic("TODO: implement")
}

func  (me *AudioEffectCapture) GetDiscardedFrames()  {
  panic("TODO: implement")
}

func  (me *AudioEffectCapture) GetBufferLengthFrames()  {
  panic("TODO: implement")
}

func  (me *AudioEffectCapture) GetPushedFrames()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
