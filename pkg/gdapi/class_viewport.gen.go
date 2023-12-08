// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Viewport struct {
  obj gdc.ObjectPtr
}

func createViewport(obj gdc.ObjectPtr) *Viewport {
  return &Viewport{
    obj: obj,
  }
}

func (me *Viewport) BaseClass() string {
  return "Viewport"
}
