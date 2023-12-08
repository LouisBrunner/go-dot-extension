// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Window struct {
  obj gdc.ObjectPtr
}

func createWindow(obj gdc.ObjectPtr) *Window {
  return &Window{
    obj: obj,
  }
}

func (me *Window) BaseClass() string {
  return "Window"
}
