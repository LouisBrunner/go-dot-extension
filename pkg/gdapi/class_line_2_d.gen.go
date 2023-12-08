// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Line2D struct {
  obj gdc.ObjectPtr
}

func createLine2D(obj gdc.ObjectPtr) *Line2D {
  return &Line2D{
    obj: obj,
  }
}

func (me *Line2D) BaseClass() string {
  return "Line2D"
}
