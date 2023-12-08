// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorldBoundaryShape2D struct {
  obj gdc.ObjectPtr
}

func createWorldBoundaryShape2D(obj gdc.ObjectPtr) *WorldBoundaryShape2D {
  return &WorldBoundaryShape2D{
    obj: obj,
  }
}

func (me *WorldBoundaryShape2D) BaseClass() string {
  return "WorldBoundaryShape2D"
}
