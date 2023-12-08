// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventJoypadButton struct {
  obj gdc.ObjectPtr
}

func createInputEventJoypadButton(obj gdc.ObjectPtr) *InputEventJoypadButton {
  return &InputEventJoypadButton{
    obj: obj,
  }
}

func (me *InputEventJoypadButton) BaseClass() string {
  return "InputEventJoypadButton"
}
