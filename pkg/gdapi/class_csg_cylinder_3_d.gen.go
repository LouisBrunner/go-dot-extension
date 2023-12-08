// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGCylinder3D struct {
  obj gdc.ObjectPtr
}

func createCSGCylinder3D(obj gdc.ObjectPtr) *CSGCylinder3D {
  return &CSGCylinder3D{
    obj: obj,
  }
}

func (me *CSGCylinder3D) BaseClass() string {
  return "CSGCylinder3D"
}
