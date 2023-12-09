// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEvent struct {
  obj gdc.ObjectPtr
}

func (me *InputEvent) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *InputEvent) BaseClass() string {
  return "InputEvent"
}



// Enums

func (me *InputEvent) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *InputEvent) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *InputEvent) SetDevice(device int, )  {
  panic("TODO: implement")
}

func  (me *InputEvent) GetDevice()  {
  panic("TODO: implement")
}

func  (me *InputEvent) IsAction(action StringName, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *InputEvent) IsActionPressed(action StringName, allow_echo bool, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *InputEvent) IsActionReleased(action StringName, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *InputEvent) GetActionStrength(action StringName, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *InputEvent) IsCanceled()  {
  panic("TODO: implement")
}

func  (me *InputEvent) IsPressed()  {
  panic("TODO: implement")
}

func  (me *InputEvent) IsReleased()  {
  panic("TODO: implement")
}

func  (me *InputEvent) IsEcho()  {
  panic("TODO: implement")
}

func  (me *InputEvent) AsText()  {
  panic("TODO: implement")
}

func  (me *InputEvent) IsMatch(event InputEvent, exact_match bool, )  {
  panic("TODO: implement")
}

func  (me *InputEvent) IsActionType()  {
  panic("TODO: implement")
}

func  (me *InputEvent) Accumulate(with_event InputEvent, )  {
  panic("TODO: implement")
}

func  (me *InputEvent) XformedBy(xform Transform2D, local_ofs Vector2, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
