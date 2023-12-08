// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PopupPanel struct {
  obj gdc.ObjectPtr
}

func createPopupPanel(obj gdc.ObjectPtr) *PopupPanel {
  return &PopupPanel{
    obj: obj,
  }
}

func (me *PopupPanel) BaseClass() string {
  return "PopupPanel"
}
