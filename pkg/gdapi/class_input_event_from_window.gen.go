// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventFromWindow struct {
  obj gdc.ObjectPtr
}

func (me *InputEventFromWindow) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventFromWindow) BaseClass() string {
  return "InputEventFromWindow"
}



// Enums

func (me *InputEventFromWindow) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *InputEventFromWindow) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventFromWindow) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventFromWindow) SetWindowId(id int, )  {
  panic("TODO: implement")
}

func  (me *InputEventFromWindow) GetWindowId()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
