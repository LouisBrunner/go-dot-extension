// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventMIDI struct {
  obj gdc.ObjectPtr
}

func createInputEventMIDI(obj gdc.ObjectPtr) *InputEventMIDI {
  return &InputEventMIDI{
    obj: obj,
  }
}

func (me *InputEventMIDI) BaseClass() string {
  return "InputEventMIDI"
}
