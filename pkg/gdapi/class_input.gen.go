// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Input struct {
  obj gdc.ObjectPtr
}

func createInput(obj gdc.ObjectPtr) *Input {
  return &Input{
    obj: obj,
  }
}

func (me *Input) BaseClass() string {
  return "Input"
}
