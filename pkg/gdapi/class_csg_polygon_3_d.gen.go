// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGPolygon3D struct {
  obj gdc.ObjectPtr
}

func createCSGPolygon3D(obj gdc.ObjectPtr) *CSGPolygon3D {
  return &CSGPolygon3D{
    obj: obj,
  }
}

func (me *CSGPolygon3D) BaseClass() string {
  return "CSGPolygon3D"
}
