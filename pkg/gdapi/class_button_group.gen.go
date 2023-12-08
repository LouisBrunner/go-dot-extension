// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ButtonGroup struct {
  obj gdc.ObjectPtr
}

func createButtonGroup(obj gdc.ObjectPtr) *ButtonGroup {
  return &ButtonGroup{
    obj: obj,
  }
}

func (me *ButtonGroup) BaseClass() string {
  return "ButtonGroup"
}
