// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventPanGesture struct {
  obj gdc.ObjectPtr
}

func createInputEventPanGesture(obj gdc.ObjectPtr) *InputEventPanGesture {
  return &InputEventPanGesture{
    obj: obj,
  }
}

func (me *InputEventPanGesture) BaseClass() string {
  return "InputEventPanGesture"
}
