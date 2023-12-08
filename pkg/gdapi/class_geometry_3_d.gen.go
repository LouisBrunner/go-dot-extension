// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Geometry3D struct {
  obj gdc.ObjectPtr
}

func createGeometry3D(obj gdc.ObjectPtr) *Geometry3D {
  return &Geometry3D{
    obj: obj,
  }
}

func (me *Geometry3D) BaseClass() string {
  return "Geometry3D"
}
