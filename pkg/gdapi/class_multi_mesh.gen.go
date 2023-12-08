// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiMesh struct {
  obj gdc.ObjectPtr
}

func createMultiMesh(obj gdc.ObjectPtr) *MultiMesh {
  return &MultiMesh{
    obj: obj,
  }
}

func (me *MultiMesh) BaseClass() string {
  return "MultiMesh"
}
