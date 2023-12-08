// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GeometryInstance3D struct {
  obj gdc.ObjectPtr
}

func createGeometryInstance3D(obj gdc.ObjectPtr) *GeometryInstance3D {
  return &GeometryInstance3D{
    obj: obj,
  }
}

func (me *GeometryInstance3D) BaseClass() string {
  return "GeometryInstance3D"
}
