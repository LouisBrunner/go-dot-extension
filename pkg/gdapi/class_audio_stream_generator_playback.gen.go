// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamGeneratorPlayback struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamGeneratorPlayback) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamGeneratorPlayback) BaseClass() string {
  return "AudioStreamGeneratorPlayback"
}



// Enums

func (me *AudioStreamGeneratorPlayback) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamGeneratorPlayback) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamGeneratorPlayback) PushFrame(frame Vector2, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamGeneratorPlayback) CanPushBuffer(amount int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamGeneratorPlayback) PushBuffer(frames PackedVector2Array, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamGeneratorPlayback) GetFramesAvailable()  {
  panic("TODO: implement")
}

func  (me *AudioStreamGeneratorPlayback) GetSkips()  {
  panic("TODO: implement")
}

func  (me *AudioStreamGeneratorPlayback) ClearBuffer()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
