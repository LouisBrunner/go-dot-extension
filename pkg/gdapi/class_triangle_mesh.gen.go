// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type TriangleMesh struct {
  obj gdc.ObjectPtr
}

func createTriangleMesh(obj gdc.ObjectPtr) *TriangleMesh {
  return &TriangleMesh{
    obj: obj,
  }
}

func (me *TriangleMesh) BaseClass() string {
  return "TriangleMesh"
}
