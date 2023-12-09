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



// Enums

type VoxelGISubdiv int
const (
  VoxelGISubdivSubdiv64 VoxelGISubdiv = 0
  VoxelGISubdivSubdiv128 VoxelGISubdiv = 1
  VoxelGISubdivSubdiv256 VoxelGISubdiv = 2
  VoxelGISubdivSubdiv512 VoxelGISubdiv = 3
  VoxelGISubdivSubdivMax VoxelGISubdiv = 4
)

func (me *VoxelGI) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VoxelGI) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VoxelGI) SetProbeData(data VoxelGIData, )  {
  panic("TODO: implement")
}

func  (me *VoxelGI) GetProbeData()  {
  panic("TODO: implement")
}

func  (me *VoxelGI) SetSubdiv(subdiv VoxelGISubdiv, )  {
  panic("TODO: implement")
}

func  (me *VoxelGI) GetSubdiv()  {
  panic("TODO: implement")
}

func  (me *VoxelGI) SetSize(size Vector3, )  {
  panic("TODO: implement")
}

func  (me *VoxelGI) GetSize()  {
  panic("TODO: implement")
}

func  (me *VoxelGI) SetCameraAttributes(camera_attributes CameraAttributes, )  {
  panic("TODO: implement")
}

func  (me *VoxelGI) GetCameraAttributes()  {
  panic("TODO: implement")
}

func  (me *VoxelGI) Bake(from_node Node, create_visual_debug bool, )  {
  panic("TODO: implement")
}

func  (me *VoxelGI) DebugBake()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
