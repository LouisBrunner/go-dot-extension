// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Popup struct {
  obj gdc.ObjectPtr
}

func createPopup(obj gdc.ObjectPtr) *Popup {
  return &Popup{
    obj: obj,
  }
}

func (me *Popup) BaseClass() string {
  return "Popup"
}
