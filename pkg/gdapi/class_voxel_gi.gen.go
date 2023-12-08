// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VoxelGISubdivSubdiv64 VoxelGISubdiv = 0
  VoxelGISubdivSubdiv128 VoxelGISubdiv = 1
  VoxelGISubdivSubdiv256 VoxelGISubdiv = 2
  VoxelGISubdivSubdiv512 VoxelGISubdiv = 3
  VoxelGISubdivSubdivMax VoxelGISubdiv = 4
)

func  (me *VoxelGI) SetProbeData(data VoxelGIData, ) { // TODO: return value
  // TODO: implement
}

func  (me *VoxelGI) GetProbeData() { // TODO: return value
  // TODO: implement
}

func  (me *VoxelGI) SetSubdiv(subdiv VoxelGISubdiv, ) { // TODO: return value
  // TODO: implement
}

func  (me *VoxelGI) GetSubdiv() { // TODO: return value
  // TODO: implement
}

func  (me *VoxelGI) SetSize(size Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *VoxelGI) GetSize() { // TODO: return value
  // TODO: implement
}

func  (me *VoxelGI) SetCameraAttributes(camera_attributes CameraAttributes, ) { // TODO: return value
  // TODO: implement
}

func  (me *VoxelGI) GetCameraAttributes() { // TODO: return value
  // TODO: implement
}

func  (me *VoxelGI) Bake(from_node Node, create_visual_debug bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *VoxelGI) DebugBake() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
