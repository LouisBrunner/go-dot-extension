// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ScrollBar struct {
  obj gdc.ObjectPtr
}

func createScrollBar(obj gdc.ObjectPtr) *ScrollBar {
  return &ScrollBar{
    obj: obj,
  }
}

func (me *ScrollBar) BaseClass() string {
  return "ScrollBar"
}
