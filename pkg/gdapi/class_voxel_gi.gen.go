// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VoxelGI struct {
  obj gdc.ObjectPtr
}

func (me *VoxelGI) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VoxelGI) BaseClass() string {
  return "VoxelGI"
}

type VoxelGISubdiv int
const (
  VoxelGISubdiv64 VoxelGISubdiv = 0
  VoxelGISubdiv128 VoxelGISubdiv = 1
  VoxelGISubdiv256 VoxelGISubdiv = 2
  VoxelGISubdiv512 VoxelGISubdiv = 3
  VoxelGISubdivMax VoxelGISubdiv = 4
)
