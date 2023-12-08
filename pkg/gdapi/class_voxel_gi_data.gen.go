// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VoxelGIData struct {
  obj gdc.ObjectPtr
}

func createVoxelGIData(obj gdc.ObjectPtr) *VoxelGIData {
  return &VoxelGIData{
    obj: obj,
  }
}

func (me *VoxelGIData) BaseClass() string {
  return "VoxelGIData"
}
