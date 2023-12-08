// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventJoypadMotion struct {
  obj gdc.ObjectPtr
}

func createInputEventJoypadMotion(obj gdc.ObjectPtr) *InputEventJoypadMotion {
  return &InputEventJoypadMotion{
    obj: obj,
  }
}

func (me *InputEventJoypadMotion) BaseClass() string {
  return "InputEventJoypadMotion"
}
