// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiMeshInstance2D struct {
  obj gdc.ObjectPtr
}

func createMultiMeshInstance2D(obj gdc.ObjectPtr) *MultiMeshInstance2D {
  return &MultiMeshInstance2D{
    obj: obj,
  }
}

func (me *MultiMeshInstance2D) BaseClass() string {
  return "MultiMeshInstance2D"
}
