// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffectAmplify struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffectAmplify) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffectAmplify) BaseClass() string {
  return "AudioEffectAmplify"
}



// Enums

func (me *AudioEffectAmplify) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffectAmplify) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffectAmplify) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffectAmplify) SetVolumeDb(volume float32, )  {
  panic("TODO: implement")
}

func  (me *AudioEffectAmplify) GetVolumeDb()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
