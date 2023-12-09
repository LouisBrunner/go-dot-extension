// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VoxelGIData struct {
  obj gdc.ObjectPtr
}

func (me *VoxelGIData) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VoxelGIData) BaseClass() string {
  return "VoxelGIData"
}



// Enums

func (me *VoxelGIData) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VoxelGIData) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VoxelGIData) Allocate(to_cell_xform Transform3D, aabb AABB, octree_size Vector3, octree_cells PackedByteArray, data_cells PackedByteArray, distance_field PackedByteArray, level_counts PackedInt32Array, )  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetBounds()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetOctreeSize()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetToCellXform()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetOctreeCells()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetDataCells()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetLevelCounts()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) SetDynamicRange(dynamic_range float32, )  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetDynamicRange()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) SetEnergy(energy float32, )  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetEnergy()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) SetBias(bias float32, )  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetBias()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) SetNormalBias(bias float32, )  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetNormalBias()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) SetPropagation(propagation float32, )  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) GetPropagation()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) SetInterior(interior bool, )  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) IsInterior()  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) SetUseTwoBounces(enable bool, )  {
  panic("TODO: implement")
}

func  (me *VoxelGIData) IsUsingTwoBounces()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
