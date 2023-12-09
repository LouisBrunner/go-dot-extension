// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventMagnifyGesture struct {
  obj gdc.ObjectPtr
}

func (me *InputEventMagnifyGesture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventMagnifyGesture) BaseClass() string {
  return "InputEventMagnifyGesture"
}



// Enums

func (me *InputEventMagnifyGesture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventMagnifyGesture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventMagnifyGesture) SetFactor(factor float32, )  {
  panic("TODO: implement")
}

func  (me *InputEventMagnifyGesture) GetFactor()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
