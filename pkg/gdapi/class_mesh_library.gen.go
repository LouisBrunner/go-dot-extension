// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshLibrary struct {
  obj gdc.ObjectPtr
}

func createMeshLibrary(obj gdc.ObjectPtr) *MeshLibrary {
  return &MeshLibrary{
    obj: obj,
  }
}

func (me *MeshLibrary) BaseClass() string {
  return "MeshLibrary"
}
