// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputMap struct {
  obj gdc.ObjectPtr
}

func (me *InputMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputMap) BaseClass() string {
  return "InputMap"
}



// Enums

func (me *InputMap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputMap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputMap) HasAction(action StringName, )  {
  panic("TODO: implement")
}

func  (me *InputMap) GetActions()  {
  panic("TODO: implement")
}

func  (me *InputMap) AddAction(action StringName, deadzone float32, )  {
  panic("TODO: implement")
}

func  (me *InputMap) EraseAction(action StringName, )  {
  panic("TODO: implement")
}

func  (me *InputMap) ActionSetDeadzone(action StringName, deadzone float32, )  {
  panic("TODO: implement")
}

func  (me *InputMap) ActionGetDeadzone(action StringName, )  {
  panic("TODO: implement")
}

func  (me *InputMap) ActionAddEvent(action StringName, event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *InputMap) ActionHasEvent(action StringName, event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *InputMap) ActionEraseEvent(action StringName, event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *InputMap) ActionEraseEvents(action StringName, )  {
  panic("TODO: implement")
}

func  (me *InputMap) ActionGetEvents(action StringName, )  {
  panic("TODO: implement")
}

func  (me *InputMap) EventIsAction(event InputEvent, action StringName, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *InputMap) LoadFromProjectSettings()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
