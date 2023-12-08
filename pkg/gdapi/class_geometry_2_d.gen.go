// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Geometry2D struct {
  obj gdc.ObjectPtr
}

func createGeometry2D(obj gdc.ObjectPtr) *Geometry2D {
  return &Geometry2D{
    obj: obj,
  }
}

func (me *Geometry2D) BaseClass() string {
  return "Geometry2D"
}
