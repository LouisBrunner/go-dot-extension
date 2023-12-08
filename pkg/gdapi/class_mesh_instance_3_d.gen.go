// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshInstance3D struct {
  obj gdc.ObjectPtr
}

func createMeshInstance3D(obj gdc.ObjectPtr) *MeshInstance3D {
  return &MeshInstance3D{
    obj: obj,
  }
}

func (me *MeshInstance3D) BaseClass() string {
  return "MeshInstance3D"
}
