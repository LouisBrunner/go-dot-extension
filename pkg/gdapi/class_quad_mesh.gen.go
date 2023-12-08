// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type QuadMesh struct {
  obj gdc.ObjectPtr
}

func createQuadMesh(obj gdc.ObjectPtr) *QuadMesh {
  return &QuadMesh{
    obj: obj,
  }
}

func (me *QuadMesh) BaseClass() string {
  return "QuadMesh"
}
