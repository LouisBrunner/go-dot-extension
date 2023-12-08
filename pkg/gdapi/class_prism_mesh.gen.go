// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PrismMesh struct {
  obj gdc.ObjectPtr
}

func createPrismMesh(obj gdc.ObjectPtr) *PrismMesh {
  return &PrismMesh{
    obj: obj,
  }
}

func (me *PrismMesh) BaseClass() string {
  return "PrismMesh"
}
