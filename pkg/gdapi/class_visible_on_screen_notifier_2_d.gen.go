// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisibleOnScreenNotifier2D struct {
  obj gdc.ObjectPtr
}

func (me *VisibleOnScreenNotifier2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisibleOnScreenNotifier2D) BaseClass() string {
  return "VisibleOnScreenNotifier2D"
}



// Enums

func (me *VisibleOnScreenNotifier2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisibleOnScreenNotifier2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisibleOnScreenNotifier2D) SetRect(rect Rect2, )  {
  panic("TODO: implement")
}

func  (me *VisibleOnScreenNotifier2D) GetRect()  {
  panic("TODO: implement")
}

func  (me *VisibleOnScreenNotifier2D) IsOnScreen()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
