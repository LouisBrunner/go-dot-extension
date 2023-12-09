// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventJoypadButton struct {
  obj gdc.ObjectPtr
}

func (me *InputEventJoypadButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventJoypadButton) BaseClass() string {
  return "InputEventJoypadButton"
}



// Enums

func (me *InputEventJoypadButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventJoypadButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventJoypadButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventJoypadButton) SetButtonIndex(button_index JoyButton, )  {
  panic("TODO: implement")
}

func  (me *InputEventJoypadButton) GetButtonIndex()  {
  panic("TODO: implement")
}

func  (me *InputEventJoypadButton) SetPressure(pressure float32, )  {
  panic("TODO: implement")
}

func  (me *InputEventJoypadButton) GetPressure()  {
  panic("TODO: implement")
}

func  (me *InputEventJoypadButton) SetPressed(pressed bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
