// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventMagnifyGesture struct {
  obj gdc.ObjectPtr
}

func createInputEventMagnifyGesture(obj gdc.ObjectPtr) *InputEventMagnifyGesture {
  return &InputEventMagnifyGesture{
    obj: obj,
  }
}

func (me *InputEventMagnifyGesture) BaseClass() string {
  return "InputEventMagnifyGesture"
}
