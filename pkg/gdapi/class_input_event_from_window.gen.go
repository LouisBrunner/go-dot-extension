// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventFromWindow struct {
  obj gdc.ObjectPtr
}

func createInputEventFromWindow(obj gdc.ObjectPtr) *InputEventFromWindow {
  return &InputEventFromWindow{
    obj: obj,
  }
}

func (me *InputEventFromWindow) BaseClass() string {
  return "InputEventFromWindow"
}
