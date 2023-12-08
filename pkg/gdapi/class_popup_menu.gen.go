// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PopupMenu struct {
  obj gdc.ObjectPtr
}

func createPopupMenu(obj gdc.ObjectPtr) *PopupMenu {
  return &PopupMenu{
    obj: obj,
  }
}

func (me *PopupMenu) BaseClass() string {
  return "PopupMenu"
}
