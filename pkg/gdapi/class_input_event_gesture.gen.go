// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventGesture struct {
  obj gdc.ObjectPtr
}

func createInputEventGesture(obj gdc.ObjectPtr) *InputEventGesture {
  return &InputEventGesture{
    obj: obj,
  }
}

func (me *InputEventGesture) BaseClass() string {
  return "InputEventGesture"
}
