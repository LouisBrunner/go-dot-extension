// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type XROrigin3D struct {
  obj gdc.ObjectPtr
}

func (me *XROrigin3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XROrigin3D) BaseClass() string {
  return "XROrigin3D"
}



// Enums

func (me *XROrigin3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XROrigin3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *XROrigin3D) SetWorldScale(world_scale float32, )  {
  panic("TODO: implement")
}

func  (me *XROrigin3D) GetWorldScale()  {
  panic("TODO: implement")
}

func  (me *XROrigin3D) SetCurrent(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *XROrigin3D) IsCurrent()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
