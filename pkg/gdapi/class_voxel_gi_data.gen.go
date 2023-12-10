// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(to_cell_xform.AsCTypePtr()), gdc.ConstTypePtr(aabb.AsCTypePtr()), gdc.ConstTypePtr(octree_size.AsCTypePtr()), gdc.ConstTypePtr(octree_cells.AsCTypePtr()), gdc.ConstTypePtr(data_cells.AsCTypePtr()), gdc.ConstTypePtr(distance_field.AsCTypePtr()), gdc.ConstTypePtr(level_counts.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGIData) GetBounds() AABB {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bounds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1068685055) // FIXME: should cache?
  var ret AABB
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) GetOctreeSize() Vector3 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_octree_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) GetToCellXform() Transform3D {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_to_cell_xform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3229777777) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) GetOctreeCells() PackedByteArray {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_octree_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2362200018) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) GetDataCells() PackedByteArray {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_data_cells")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2362200018) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) GetLevelCounts() PackedInt32Array {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_level_counts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  var ret PackedInt32Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) SetDynamicRange(dynamic_range float32, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_dynamic_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&dynamic_range), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGIData) GetDynamicRange() float32 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_dynamic_range")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) SetEnergy(energy float32, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGIData) GetEnergy() float32 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_energy")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) SetBias(bias float32, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGIData) GetBias() float32 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) SetNormalBias(bias float32, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGIData) GetNormalBias() float32 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normal_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) SetPropagation(propagation float32, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_propagation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&propagation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGIData) GetPropagation() float32 {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_propagation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) SetInterior(interior bool, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_interior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interior), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGIData) IsInterior() bool {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_interior")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VoxelGIData) SetUseTwoBounces(enable bool, )  {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_two_bounces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VoxelGIData) IsUsingTwoBounces() bool {
  classNameV := StringNameFromStr("VoxelGIData")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_using_two_bounces")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VoxelGIData) GetPropDynamicRange() float32 {
  panic("TODO: implement")
}

func (me *VoxelGIData) SetPropDynamicRange(value float32) {
  panic("TODO: implement")
}

func (me *VoxelGIData) GetPropEnergy() float32 {
  panic("TODO: implement")
}

func (me *VoxelGIData) SetPropEnergy(value float32) {
  panic("TODO: implement")
}

func (me *VoxelGIData) GetPropBias() float32 {
  panic("TODO: implement")
}

func (me *VoxelGIData) SetPropBias(value float32) {
  panic("TODO: implement")
}

func (me *VoxelGIData) GetPropNormalBias() float32 {
  panic("TODO: implement")
}

func (me *VoxelGIData) SetPropNormalBias(value float32) {
  panic("TODO: implement")
}

func (me *VoxelGIData) GetPropPropagation() float32 {
  panic("TODO: implement")
}

func (me *VoxelGIData) SetPropPropagation(value float32) {
  panic("TODO: implement")
}

func (me *VoxelGIData) GetPropUseTwoBounces() bool {
  panic("TODO: implement")
}

func (me *VoxelGIData) SetPropUseTwoBounces(value bool) {
  panic("TODO: implement")
}

func (me *VoxelGIData) GetPropInterior() bool {
  panic("TODO: implement")
}

func (me *VoxelGIData) SetPropInterior(value bool) {
  panic("TODO: implement")
}