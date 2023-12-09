// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventMouse struct {
  obj gdc.ObjectPtr
}

func (me *InputEventMouse) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventMouse) BaseClass() string {
  return "InputEventMouse"
}



// Enums

func (me *InputEventMouse) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventMouse) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventMouse) SetButtonMask(button_mask MouseButtonMask, )  {
  panic("TODO: implement")
}

func  (me *InputEventMouse) GetButtonMask()  {
  panic("TODO: implement")
}

func  (me *InputEventMouse) SetPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *InputEventMouse) GetPosition()  {
  panic("TODO: implement")
}

func  (me *InputEventMouse) SetGlobalPosition(global_position Vector2, )  {
  panic("TODO: implement")
}

func  (me *InputEventMouse) GetGlobalPosition()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
