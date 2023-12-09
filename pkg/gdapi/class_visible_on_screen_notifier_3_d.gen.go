// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenNotifier3D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenNotifier3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenNotifier3D) BaseClass() string {
  return "VisibleOnScreenNotifier3D"
}



// Enums

func (me *VisibleOnScreenNotifier3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenNotifier3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisibleOnScreenNotifier3D) SetAabb(rect AABB, )  {
  panic("TODO: implement")
}

func  (me *VisibleOnScreenNotifier3D) IsOnScreen()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
