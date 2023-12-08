// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MultiMeshInstance3D struct {
  obj gdc.ObjectPtr
}

func createMultiMeshInstance3D(obj gdc.ObjectPtr) *MultiMeshInstance3D {
  return &MultiMeshInstance3D{
    obj: obj,
  }
}

func (me *MultiMeshInstance3D) BaseClass() string {
  return "MultiMeshInstance3D"
}
