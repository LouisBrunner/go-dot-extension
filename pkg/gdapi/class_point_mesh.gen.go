// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PointMesh struct {
  obj gdc.ObjectPtr
}

func createPointMesh(obj gdc.ObjectPtr) *PointMesh {
  return &PointMesh{
    obj: obj,
  }
}

func (me *PointMesh) BaseClass() string {
  return "PointMesh"
}
