// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventScreenTouch struct {
  obj gdc.ObjectPtr
}

func createInputEventScreenTouch(obj gdc.ObjectPtr) *InputEventScreenTouch {
  return &InputEventScreenTouch{
    obj: obj,
  }
}

func (me *InputEventScreenTouch) BaseClass() string {
  return "InputEventScreenTouch"
}
