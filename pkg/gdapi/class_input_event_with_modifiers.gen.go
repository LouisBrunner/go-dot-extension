// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventWithModifiers struct {
  obj gdc.ObjectPtr
}

func createInputEventWithModifiers(obj gdc.ObjectPtr) *InputEventWithModifiers {
  return &InputEventWithModifiers{
    obj: obj,
  }
}

func (me *InputEventWithModifiers) BaseClass() string {
  return "InputEventWithModifiers"
}
