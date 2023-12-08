// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RayCast2D struct {
  obj gdc.ObjectPtr
}

func createRayCast2D(obj gdc.ObjectPtr) *RayCast2D {
  return &RayCast2D{
    obj: obj,
  }
}

func (me *RayCast2D) BaseClass() string {
  return "RayCast2D"
}
