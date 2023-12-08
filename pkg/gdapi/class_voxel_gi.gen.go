// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VoxelGI struct {
  obj gdc.ObjectPtr
}

func createVoxelGI(obj gdc.ObjectPtr) *VoxelGI {
  return &VoxelGI{
    obj: obj,
  }
}

func (me *VoxelGI) BaseClass() string {
  return "VoxelGI"
}
