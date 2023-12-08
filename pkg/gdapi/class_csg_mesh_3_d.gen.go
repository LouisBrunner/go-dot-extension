// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CSGMesh3D struct {
  obj gdc.ObjectPtr
}

func createCSGMesh3D(obj gdc.ObjectPtr) *CSGMesh3D {
  return &CSGMesh3D{
    obj: obj,
  }
}

func (me *CSGMesh3D) BaseClass() string {
  return "CSGMesh3D"
}
