// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XRController3D struct {
  obj gdc.ObjectPtr
}

func (me *XRController3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XRController3D) BaseClass() string {
  return "XRController3D"
}



// Enums

func (me *XRController3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XRController3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *XRController3D) IsButtonPressed(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRController3D) GetInput(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRController3D) GetFloat(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRController3D) GetVector2(name StringName, )  {
  panic("TODO: implement")
}

func  (me *XRController3D) GetTrackerHand()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
