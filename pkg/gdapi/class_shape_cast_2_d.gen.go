// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ShapeCast2D struct {
  obj gdc.ObjectPtr
}

func createShapeCast2D(obj gdc.ObjectPtr) *ShapeCast2D {
  return &ShapeCast2D{
    obj: obj,
  }
}

func (me *ShapeCast2D) BaseClass() string {
  return "ShapeCast2D"
}
