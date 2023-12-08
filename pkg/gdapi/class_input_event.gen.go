// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEvent struct {
  obj gdc.ObjectPtr
}

func createInputEvent(obj gdc.ObjectPtr) *InputEvent {
  return &InputEvent{
    obj: obj,
  }
}

func (me *InputEvent) BaseClass() string {
  return "InputEvent"
}
