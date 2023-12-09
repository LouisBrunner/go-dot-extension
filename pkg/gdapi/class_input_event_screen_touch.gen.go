// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventScreenTouch struct {
  obj gdc.ObjectPtr
}

func (me *InputEventScreenTouch) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventScreenTouch) BaseClass() string {
  return "InputEventScreenTouch"
}



// Enums

func (me *InputEventScreenTouch) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventScreenTouch) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventScreenTouch) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventScreenTouch) SetIndex(index int, )  {
  panic("TODO: implement")
}

func  (me *InputEventScreenTouch) GetIndex()  {
  panic("TODO: implement")
}

func  (me *InputEventScreenTouch) SetPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *InputEventScreenTouch) GetPosition()  {
  panic("TODO: implement")
}

func  (me *InputEventScreenTouch) SetPressed(pressed bool, )  {
  panic("TODO: implement")
}

func  (me *InputEventScreenTouch) SetCanceled(canceled bool, )  {
  panic("TODO: implement")
}

func  (me *InputEventScreenTouch) SetDoubleTap(double_tap bool, )  {
  panic("TODO: implement")
}

func  (me *InputEventScreenTouch) IsDoubleTap()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
