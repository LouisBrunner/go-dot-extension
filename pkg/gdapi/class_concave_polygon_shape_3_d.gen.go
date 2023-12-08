// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConcavePolygonShape3D struct {
  obj gdc.ObjectPtr
}

func createConcavePolygonShape3D(obj gdc.ObjectPtr) *ConcavePolygonShape3D {
  return &ConcavePolygonShape3D{
    obj: obj,
  }
}

func (me *ConcavePolygonShape3D) BaseClass() string {
  return "ConcavePolygonShape3D"
}
