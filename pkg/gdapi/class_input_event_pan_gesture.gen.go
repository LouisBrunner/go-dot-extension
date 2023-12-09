// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventPanGesture struct {
  obj gdc.ObjectPtr
}

func (me *InputEventPanGesture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventPanGesture) BaseClass() string {
  return "InputEventPanGesture"
}



// Enums

func (me *InputEventPanGesture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventPanGesture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventPanGesture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventPanGesture) SetDelta(delta Vector2, )  {
  panic("TODO: implement")
}

func  (me *InputEventPanGesture) GetDelta()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
