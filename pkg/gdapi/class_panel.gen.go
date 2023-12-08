// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Panel struct {
  obj gdc.ObjectPtr
}

func createPanel(obj gdc.ObjectPtr) *Panel {
  return &Panel{
    obj: obj,
  }
}

func (me *Panel) BaseClass() string {
  return "Panel"
}
