// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AudioListener3D struct {
  obj gdc.ObjectPtr
}

func (me *AudioListener3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AudioListener3D) BaseClass() string {
  return "AudioListener3D"
}



// Enums

func (me *AudioListener3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AudioListener3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AudioListener3D) MakeCurrent()  {
  panic("TODO: implement")
}

func  (me *AudioListener3D) ClearCurrent()  {
  panic("TODO: implement")
}

func  (me *AudioListener3D) IsCurrent()  {
  panic("TODO: implement")
}

func  (me *AudioListener3D) GetListenerTransform()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
