// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type OccluderPolygon2D struct {
  obj gdc.ObjectPtr
}

func createOccluderPolygon2D(obj gdc.ObjectPtr) *OccluderPolygon2D {
  return &OccluderPolygon2D{
    obj: obj,
  }
}

func (me *OccluderPolygon2D) BaseClass() string {
  return "OccluderPolygon2D"
}
