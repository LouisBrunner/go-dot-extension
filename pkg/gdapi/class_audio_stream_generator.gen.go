// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamGenerator struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamGenerator) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamGenerator) BaseClass() string {
  return "AudioStreamGenerator"
}



// Enums

func (me *AudioStreamGenerator) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamGenerator) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamGenerator) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamGenerator) SetMixRate(hz float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamGenerator) GetMixRate()  {
  panic("TODO: implement")
}

func  (me *AudioStreamGenerator) SetBufferLength(seconds float32, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamGenerator) GetBufferLength()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
