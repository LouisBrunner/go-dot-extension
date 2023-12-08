// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasItemMaterial struct {
  obj gdc.ObjectPtr
}

func createCanvasItemMaterial(obj gdc.ObjectPtr) *CanvasItemMaterial {
  return &CanvasItemMaterial{
    obj: obj,
  }
}

func (me *CanvasItemMaterial) BaseClass() string {
  return "CanvasItemMaterial"
}
