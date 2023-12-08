// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Polygon2D struct {
  obj gdc.ObjectPtr
}

func createPolygon2D(obj gdc.ObjectPtr) *Polygon2D {
  return &Polygon2D{
    obj: obj,
  }
}

func (me *Polygon2D) BaseClass() string {
  return "Polygon2D"
}
