// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConcavePolygonShape2D struct {
  obj gdc.ObjectPtr
}

func createConcavePolygonShape2D(obj gdc.ObjectPtr) *ConcavePolygonShape2D {
  return &ConcavePolygonShape2D{
    obj: obj,
  }
}

func (me *ConcavePolygonShape2D) BaseClass() string {
  return "ConcavePolygonShape2D"
}
