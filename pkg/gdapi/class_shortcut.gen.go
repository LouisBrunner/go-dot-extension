// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shortcut struct {
  obj gdc.ObjectPtr
}

func createShortcut(obj gdc.ObjectPtr) *Shortcut {
  return &Shortcut{
    obj: obj,
  }
}

func (me *Shortcut) BaseClass() string {
  return "Shortcut"
}
