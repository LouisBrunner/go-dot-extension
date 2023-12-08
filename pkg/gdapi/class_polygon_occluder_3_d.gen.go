// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PolygonOccluder3D struct {
  obj gdc.ObjectPtr
}

func createPolygonOccluder3D(obj gdc.ObjectPtr) *PolygonOccluder3D {
  return &PolygonOccluder3D{
    obj: obj,
  }
}

func (me *PolygonOccluder3D) BaseClass() string {
  return "PolygonOccluder3D"
}
