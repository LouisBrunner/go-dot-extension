// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGSphere3D struct {
  obj gdc.ObjectPtr
}

func createCSGSphere3D(obj gdc.ObjectPtr) *CSGSphere3D {
  return &CSGSphere3D{
    obj: obj,
  }
}

func (me *CSGSphere3D) BaseClass() string {
  return "CSGSphere3D"
}
