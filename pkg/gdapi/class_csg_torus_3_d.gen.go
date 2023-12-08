// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGTorus3D struct {
  obj gdc.ObjectPtr
}

func createCSGTorus3D(obj gdc.ObjectPtr) *CSGTorus3D {
  return &CSGTorus3D{
    obj: obj,
  }
}

func (me *CSGTorus3D) BaseClass() string {
  return "CSGTorus3D"
}
