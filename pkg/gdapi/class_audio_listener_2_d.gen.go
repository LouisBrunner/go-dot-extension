// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioListener2D struct {
  obj gdc.ObjectPtr
}

func (me *AudioListener2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioListener2D) BaseClass() string {
  return "AudioListener2D"
}



// Enums

func (me *AudioListener2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioListener2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioListener2D) MakeCurrent()  {
  panic("TODO: implement")
}

func  (me *AudioListener2D) ClearCurrent()  {
  panic("TODO: implement")
}

func  (me *AudioListener2D) IsCurrent()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
