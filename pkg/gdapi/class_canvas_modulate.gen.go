// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasModulate struct {
  obj gdc.ObjectPtr
}

func createCanvasModulate(obj gdc.ObjectPtr) *CanvasModulate {
  return &CanvasModulate{
    obj: obj,
  }
}

func (me *CanvasModulate) BaseClass() string {
  return "CanvasModulate"
}