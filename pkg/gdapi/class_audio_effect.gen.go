// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioEffect struct {
  obj gdc.ObjectPtr
}

func (me *AudioEffect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioEffect) BaseClass() string {
  return "AudioEffect"
}



// Enums

func (me *AudioEffect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AudioEffect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioEffect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioEffect) XInstantiate()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
