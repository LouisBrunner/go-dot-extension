// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventScreenDrag struct {
  obj gdc.ObjectPtr
}

func createInputEventScreenDrag(obj gdc.ObjectPtr) *InputEventScreenDrag {
  return &InputEventScreenDrag{
    obj: obj,
  }
}

func (me *InputEventScreenDrag) BaseClass() string {
  return "InputEventScreenDrag"
}
