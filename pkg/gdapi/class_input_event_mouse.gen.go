// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventMouse struct {
  obj gdc.ObjectPtr
}

func createInputEventMouse(obj gdc.ObjectPtr) *InputEventMouse {
  return &InputEventMouse{
    obj: obj,
  }
}

func (me *InputEventMouse) BaseClass() string {
  return "InputEventMouse"
}
