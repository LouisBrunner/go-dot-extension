// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MenuButton struct {
  obj gdc.ObjectPtr
}

func createMenuButton(obj gdc.ObjectPtr) *MenuButton {
  return &MenuButton{
    obj: obj,
  }
}

func (me *MenuButton) BaseClass() string {
  return "MenuButton"
}
