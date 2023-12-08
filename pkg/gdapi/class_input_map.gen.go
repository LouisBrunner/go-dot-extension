// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type InputMap struct {
  obj gdc.ObjectPtr
}

func createInputMap(obj gdc.ObjectPtr) *InputMap {
  return &InputMap{
    obj: obj,
  }
}

func (me *InputMap) BaseClass() string {
  return "InputMap"
}
