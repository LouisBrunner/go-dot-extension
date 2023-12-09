// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventJoypadMotion struct {
  obj gdc.ObjectPtr
}

func (me *InputEventJoypadMotion) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventJoypadMotion) BaseClass() string {
  return "InputEventJoypadMotion"
}



// Enums

func (me *InputEventJoypadMotion) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventJoypadMotion) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventJoypadMotion) SetAxis(axis JoyAxis, )  {
  panic("TODO: implement")
}

func  (me *InputEventJoypadMotion) GetAxis()  {
  panic("TODO: implement")
}

func  (me *InputEventJoypadMotion) SetAxisValue(axis_value float32, )  {
  panic("TODO: implement")
}

func  (me *InputEventJoypadMotion) GetAxisValue()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
