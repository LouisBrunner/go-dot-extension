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

func  (me *InputEvent) SetDevice(device int, ) { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) GetDevice() { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) IsAction(action StringName, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) IsActionPressed(action StringName, allow_echo bool, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) IsActionReleased(action StringName, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) GetActionStrength(action StringName, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) IsCanceled() { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) IsPressed() { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) IsReleased() { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) IsEcho() { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) AsText() { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) IsMatch(event InputEvent, exact_match bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) IsActionType() { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) Accumulate(with_event InputEvent, ) { // TODO: return value
  // TODO: implement
}

func  (me *InputEvent) XformedBy(xform Transform2D, local_ofs Vector2, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
