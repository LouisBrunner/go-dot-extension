// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventMouseButton struct {
  obj gdc.ObjectPtr
}

func (me *InputEventMouseButton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventMouseButton) BaseClass() string {
  return "InputEventMouseButton"
}



// Enums

func (me *InputEventMouseButton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventMouseButton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventMouseButton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventMouseButton) SetFactor(factor float32, )  {
  panic("TODO: implement")
}

func  (me *InputEventMouseButton) GetFactor()  {
  panic("TODO: implement")
}

func  (me *InputEventMouseButton) SetButtonIndex(button_index MouseButton, )  {
  panic("TODO: implement")
}

func  (me *InputEventMouseButton) GetButtonIndex()  {
  panic("TODO: implement")
}

func  (me *InputEventMouseButton) SetPressed(pressed bool, )  {
  panic("TODO: implement")
}

func  (me *InputEventMouseButton) SetCanceled(canceled bool, )  {
  panic("TODO: implement")
}

func  (me *InputEventMouseButton) SetDoubleClick(double_click bool, )  {
  panic("TODO: implement")
}

func  (me *InputEventMouseButton) IsDoubleClick()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
