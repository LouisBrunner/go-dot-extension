// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGBox3D struct {
  obj gdc.ObjectPtr
}

func createCSGBox3D(obj gdc.ObjectPtr) *CSGBox3D {
  return &CSGBox3D{
    obj: obj,
  }
}

func (me *CSGBox3D) BaseClass() string {
  return "CSGBox3D"
}
