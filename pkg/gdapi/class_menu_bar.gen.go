// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MenuBar struct {
  obj gdc.ObjectPtr
}

func createMenuBar(obj gdc.ObjectPtr) *MenuBar {
  return &MenuBar{
    obj: obj,
  }
}

func (me *MenuBar) BaseClass() string {
  return "MenuBar"
}
