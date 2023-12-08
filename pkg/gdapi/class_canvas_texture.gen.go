// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CanvasTexture struct {
  obj gdc.ObjectPtr
}

func createCanvasTexture(obj gdc.ObjectPtr) *CanvasTexture {
  return &CanvasTexture{
    obj: obj,
  }
}

func (me *CanvasTexture) BaseClass() string {
  return "CanvasTexture"
}
