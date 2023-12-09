// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventGesture struct {
  obj gdc.ObjectPtr
}

func (me *InputEventGesture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventGesture) BaseClass() string {
  return "InputEventGesture"
}



// Enums

func (me *InputEventGesture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventGesture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventGesture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventGesture) SetPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *InputEventGesture) GetPosition()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
