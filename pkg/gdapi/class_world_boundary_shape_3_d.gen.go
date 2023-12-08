// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type WorldBoundaryShape3D struct {
  obj gdc.ObjectPtr
}

func createWorldBoundaryShape3D(obj gdc.ObjectPtr) *WorldBoundaryShape3D {
  return &WorldBoundaryShape3D{
    obj: obj,
  }
}

func (me *WorldBoundaryShape3D) BaseClass() string {
  return "WorldBoundaryShape3D"
}
