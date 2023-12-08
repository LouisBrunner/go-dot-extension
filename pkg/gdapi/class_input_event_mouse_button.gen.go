// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventMouseButton struct {
  obj gdc.ObjectPtr
}

func createInputEventMouseButton(obj gdc.ObjectPtr) *InputEventMouseButton {
  return &InputEventMouseButton{
    obj: obj,
  }
}

func (me *InputEventMouseButton) BaseClass() string {
  return "InputEventMouseButton"
}
