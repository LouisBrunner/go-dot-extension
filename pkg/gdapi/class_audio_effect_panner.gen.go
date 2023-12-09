// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectPanner struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectPanner) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectPanner) BaseClass() string {
  return "AudioEffectPanner"
}



// Enums

func (me *AudioEffectPanner) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectPanner) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectPanner) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectPanner) SetPan(cpanume float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectPanner) GetPan()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
