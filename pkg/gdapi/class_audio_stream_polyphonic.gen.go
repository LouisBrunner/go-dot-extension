// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioStreamPolyphonic struct {
  obj gdc.ObjectPtr
}

func (me *AudioStreamPolyphonic) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioStreamPolyphonic) BaseClass() string {
  return "AudioStreamPolyphonic"
}



// Enums

func (me *AudioStreamPolyphonic) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioStreamPolyphonic) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioStreamPolyphonic) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioStreamPolyphonic) SetPolyphony(voices int, )  {
  panic("TODO: implement")
}

func  (me *AudioStreamPolyphonic) GetPolyphony()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
