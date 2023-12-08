// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGCombiner3D struct {
  obj gdc.ObjectPtr
}

func createCSGCombiner3D(obj gdc.ObjectPtr) *CSGCombiner3D {
  return &CSGCombiner3D{
    obj: obj,
  }
}

func (me *CSGCombiner3D) BaseClass() string {
  return "CSGCombiner3D"
}
