// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type QuadOccluder3D struct {
  obj gdc.ObjectPtr
}

func createQuadOccluder3D(obj gdc.ObjectPtr) *QuadOccluder3D {
  return &QuadOccluder3D{
    obj: obj,
  }
}

func (me *QuadOccluder3D) BaseClass() string {
  return "QuadOccluder3D"
}
