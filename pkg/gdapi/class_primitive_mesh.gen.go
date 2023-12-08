// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PrimitiveMesh struct {
  obj gdc.ObjectPtr
}

func createPrimitiveMesh(obj gdc.ObjectPtr) *PrimitiveMesh {
  return &PrimitiveMesh{
    obj: obj,
  }
}

func (me *PrimitiveMesh) BaseClass() string {
  return "PrimitiveMesh"
}
