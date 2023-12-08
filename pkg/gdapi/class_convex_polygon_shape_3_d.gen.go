// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConvexPolygonShape3D struct {
  obj gdc.ObjectPtr
}

func createConvexPolygonShape3D(obj gdc.ObjectPtr) *ConvexPolygonShape3D {
  return &ConvexPolygonShape3D{
    obj: obj,
  }
}

func (me *ConvexPolygonShape3D) BaseClass() string {
  return "ConvexPolygonShape3D"
}
