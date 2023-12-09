// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStream struct {
  obj gdc.ObjectPtr
}

func (me *AudioStream) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStream) BaseClass() string {
  return "AudioStream"
}



// Enums

func (me *AudioStream) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStream) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStream) XInstantiatePlayback()  {
  panic("TODO: implement")
}

func  (me *AudioStream) XGetStreamName()  {
  panic("TODO: implement")
}

func  (me *AudioStream) XGetLength()  {
  panic("TODO: implement")
}

func  (me *AudioStream) XIsMonophonic()  {
  panic("TODO: implement")
}

func  (me *AudioStream) XGetBpm()  {
  panic("TODO: implement")
}

func  (me *AudioStream) XGetBeatCount()  {
  panic("TODO: implement")
}

func  (me *AudioStream) GetLength()  {
  panic("TODO: implement")
}

func  (me *AudioStream) IsMonophonic()  {
  panic("TODO: implement")
}

func  (me *AudioStream) InstantiatePlayback()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
