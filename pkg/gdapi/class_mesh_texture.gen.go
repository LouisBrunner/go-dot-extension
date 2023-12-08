// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshTexture struct {
  obj gdc.ObjectPtr
}

func createMeshTexture(obj gdc.ObjectPtr) *MeshTexture {
  return &MeshTexture{
    obj: obj,
  }
}

func (me *MeshTexture) BaseClass() string {
  return "MeshTexture"
}
