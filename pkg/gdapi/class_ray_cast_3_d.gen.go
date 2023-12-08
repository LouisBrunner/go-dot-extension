// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RayCast3D struct {
  obj gdc.ObjectPtr
}

func createRayCast3D(obj gdc.ObjectPtr) *RayCast3D {
  return &RayCast3D{
    obj: obj,
  }
}

func (me *RayCast3D) BaseClass() string {
  return "RayCast3D"
}
