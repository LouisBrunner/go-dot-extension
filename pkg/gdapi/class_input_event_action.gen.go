// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventAction struct {
  obj gdc.ObjectPtr
}

func (me *InputEventAction) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEventAction) BaseClass() string {
  return "InputEventAction"
}



// Enums

func (me *InputEventAction) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEventAction) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEventAction) SetAction(action StringName, )  {
  panic("TODO: implement")
}

func  (me *InputEventAction) GetAction()  {
  panic("TODO: implement")
}

func  (me *InputEventAction) SetPressed(pressed bool, )  {
  panic("TODO: implement")
}

func  (me *InputEventAction) SetStrength(strength float32, )  {
  panic("TODO: implement")
}

func  (me *InputEventAction) GetStrength()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
