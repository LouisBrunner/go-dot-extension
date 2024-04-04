// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type VoxelGIData struct {
  Resource
}

func (me *VoxelGIData) BaseClass() string {
  return "VoxelGIData"
}

func NewVoxelGIData() *VoxelGIData {
  str := StringNameFromStr("VoxelGIData") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VoxelGIData{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VoxelGIData) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VoxelGIData) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VoxelGIData) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VoxelGIData) Allocate(to_cell_xform Transform3D, aabb AABB, octree_size Vector3, octree_cells PackedByteArray, data_cells PackedByteArray, distance_field PackedByteArray, level_counts PackedInt32Array, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("allocate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4041601946) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{to_cell_xform.AsCTypePtr(), aabb.AsCTypePtr(), octree_size.AsCTypePtr(), octree_cells.AsCTypePtr(), data_cells.AsCTypePtr(), distance_field.AsCTypePtr(), level_counts.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGIData) GetBounds() AABB {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bounds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewAABB()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VoxelGIData) GetOctreeSize() Vector3 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_octree_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VoxelGIData) GetToCellXform() Transform3D {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_to_cell_xform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VoxelGIData) GetOctreeCells() PackedByteArray {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_octree_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2362200018) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VoxelGIData) GetDataCells() PackedByteArray {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_data_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2362200018) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VoxelGIData) GetLevelCounts() PackedInt32Array {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_level_counts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VoxelGIData) SetDynamicRange(dynamic_range float64, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dynamic_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&dynamic_range) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGIData) GetDynamicRange() float64 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dynamic_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VoxelGIData) SetEnergy(energy float64, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGIData) GetEnergy() float64 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VoxelGIData) SetBias(bias float64, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGIData) GetBias() float64 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VoxelGIData) SetNormalBias(bias float64, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGIData) GetNormalBias() float64 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normal_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VoxelGIData) SetPropagation(propagation float64, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_propagation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&propagation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGIData) GetPropagation() float64 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_propagation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VoxelGIData) SetInterior(interior bool, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_interior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interior) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGIData) IsInterior() bool {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_interior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VoxelGIData) SetUseTwoBounces(enable bool, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_two_bounces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VoxelGIData) IsUsingTwoBounces() bool {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_two_bounces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
