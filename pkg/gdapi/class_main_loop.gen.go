// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MainLoop struct {
  obj gdc.ObjectPtr
}

func createMainLoop(obj gdc.ObjectPtr) *MainLoop {
  return &MainLoop{
    obj: obj,
  }
}

func (me *MainLoop) BaseClass() string {
  return "MainLoop"
}
