// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type HScrollBar struct {
  obj gdc.ObjectPtr
}

func createHScrollBar(obj gdc.ObjectPtr) *HScrollBar {
  return &HScrollBar{
    obj: obj,
  }
}

func (me *HScrollBar) BaseClass() string {
  return "HScrollBar"
}
