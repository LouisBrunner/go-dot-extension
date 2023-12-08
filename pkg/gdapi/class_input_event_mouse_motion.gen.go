// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputEventMouseMotion struct {
  obj gdc.ObjectPtr
}

func createInputEventMouseMotion(obj gdc.ObjectPtr) *InputEventMouseMotion {
  return &InputEventMouseMotion{
    obj: obj,
  }
}

func (me *InputEventMouseMotion) BaseClass() string {
  return "InputEventMouseMotion"
}
