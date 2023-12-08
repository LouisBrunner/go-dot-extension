// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasLayer struct {
  obj gdc.ObjectPtr
}

func createCanvasLayer(obj gdc.ObjectPtr) *CanvasLayer {
  return &CanvasLayer{
    obj: obj,
  }
}

func (me *CanvasLayer) BaseClass() string {
  return "CanvasLayer"
}
