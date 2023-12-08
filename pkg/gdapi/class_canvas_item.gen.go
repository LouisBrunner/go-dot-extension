// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasItem struct {
  obj gdc.ObjectPtr
}

func createCanvasItem(obj gdc.ObjectPtr) *CanvasItem {
  return &CanvasItem{
    obj: obj,
  }
}

func (me *CanvasItem) BaseClass() string {
  return "CanvasItem"
}
