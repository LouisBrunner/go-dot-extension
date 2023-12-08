// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VScrollBar struct {
  obj gdc.ObjectPtr
}

func createVScrollBar(obj gdc.ObjectPtr) *VScrollBar {
  return &VScrollBar{
    obj: obj,
  }
}

func (me *VScrollBar) BaseClass() string {
  return "VScrollBar"
}
