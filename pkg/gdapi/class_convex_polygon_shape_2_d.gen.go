// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConvexPolygonShape2D struct {
  obj gdc.ObjectPtr
}

func createConvexPolygonShape2D(obj gdc.ObjectPtr) *ConvexPolygonShape2D {
  return &ConvexPolygonShape2D{
    obj: obj,
  }
}

func (me *ConvexPolygonShape2D) BaseClass() string {
  return "ConvexPolygonShape2D"
}
