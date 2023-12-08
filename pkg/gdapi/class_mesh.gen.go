// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Mesh struct {
  obj gdc.ObjectPtr
}

func createMesh(obj gdc.ObjectPtr) *Mesh {
  return &Mesh{
    obj: obj,
  }
}

func (me *Mesh) BaseClass() string {
  return "Mesh"
}
