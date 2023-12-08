// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type MeshDataTool struct {
  obj gdc.ObjectPtr
}

func createMeshDataTool(obj gdc.ObjectPtr) *MeshDataTool {
  return &MeshDataTool{
    obj: obj,
  }
}

func (me *MeshDataTool) BaseClass() string {
  return "MeshDataTool"
}
