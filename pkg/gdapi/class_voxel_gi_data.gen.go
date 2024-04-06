// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForVoxelGIDataList struct {
	fnAllocate          gdc.MethodBindPtr
	fnGetBounds         gdc.MethodBindPtr
	fnGetOctreeSize     gdc.MethodBindPtr
	fnGetToCellXform    gdc.MethodBindPtr
	fnGetOctreeCells    gdc.MethodBindPtr
	fnGetDataCells      gdc.MethodBindPtr
	fnGetLevelCounts    gdc.MethodBindPtr
	fnSetDynamicRange   gdc.MethodBindPtr
	fnGetDynamicRange   gdc.MethodBindPtr
	fnSetEnergy         gdc.MethodBindPtr
	fnGetEnergy         gdc.MethodBindPtr
	fnSetBias           gdc.MethodBindPtr
	fnGetBias           gdc.MethodBindPtr
	fnSetNormalBias     gdc.MethodBindPtr
	fnGetNormalBias     gdc.MethodBindPtr
	fnSetPropagation    gdc.MethodBindPtr
	fnGetPropagation    gdc.MethodBindPtr
	fnSetInterior       gdc.MethodBindPtr
	fnIsInterior        gdc.MethodBindPtr
	fnSetUseTwoBounces  gdc.MethodBindPtr
	fnIsUsingTwoBounces gdc.MethodBindPtr
}

var ptrsForVoxelGIData ptrsForVoxelGIDataList

func initVoxelGIDataPtrs(iface gdc.Interface) {

	className := StringNameFromStr("VoxelGIData")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("allocate")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnAllocate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4041601946))
	}
	{
		methodName := StringNameFromStr("get_bounds")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetBounds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1068685055))
	}
	{
		methodName := StringNameFromStr("get_octree_size")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetOctreeSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_to_cell_xform")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetToCellXform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("get_octree_cells")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetOctreeCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2362200018))
	}
	{
		methodName := StringNameFromStr("get_data_cells")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetDataCells = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2362200018))
	}
	{
		methodName := StringNameFromStr("get_level_counts")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetLevelCounts = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
	}
	{
		methodName := StringNameFromStr("set_dynamic_range")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnSetDynamicRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_dynamic_range")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetDynamicRange = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_energy")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnSetEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_energy")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetEnergy = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_bias")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnSetBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_bias")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_normal_bias")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnSetNormalBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_normal_bias")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetNormalBias = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_propagation")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnSetPropagation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_propagation")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnGetPropagation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_interior")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnSetInterior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_interior")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnIsInterior = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_two_bounces")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnSetUseTwoBounces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_two_bounces")
		defer methodName.Destroy()
		ptrsForVoxelGIData.fnIsUsingTwoBounces = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

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

func (me *VoxelGIData) Allocate(to_cell_xform Transform3D, aabb AABB, octree_size Vector3, octree_cells PackedByteArray, data_cells PackedByteArray, distance_field PackedByteArray, level_counts PackedInt32Array) {
	cargs := []gdc.ConstTypePtr{to_cell_xform.AsCTypePtr(), aabb.AsCTypePtr(), octree_size.AsCTypePtr(), octree_cells.AsCTypePtr(), data_cells.AsCTypePtr(), distance_field.AsCTypePtr(), level_counts.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnAllocate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VoxelGIData) GetBounds() AABB {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewAABB()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetBounds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VoxelGIData) GetOctreeSize() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetOctreeSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VoxelGIData) GetToCellXform() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetToCellXform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VoxelGIData) GetOctreeCells() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetOctreeCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VoxelGIData) GetDataCells() PackedByteArray {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedByteArray()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetDataCells), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VoxelGIData) GetLevelCounts() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetLevelCounts), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *VoxelGIData) SetDynamicRange(dynamic_range float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&dynamic_range)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnSetDynamicRange), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VoxelGIData) GetDynamicRange() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetDynamicRange), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VoxelGIData) SetEnergy(energy float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&energy)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnSetEnergy), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VoxelGIData) GetEnergy() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetEnergy), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VoxelGIData) SetBias(bias float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnSetBias), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VoxelGIData) GetBias() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VoxelGIData) SetNormalBias(bias float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnSetNormalBias), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VoxelGIData) GetNormalBias() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetNormalBias), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VoxelGIData) SetPropagation(propagation float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&propagation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnSetPropagation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VoxelGIData) GetPropagation() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnGetPropagation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VoxelGIData) SetInterior(interior bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interior)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnSetInterior), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VoxelGIData) IsInterior() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnIsInterior), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *VoxelGIData) SetUseTwoBounces(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnSetUseTwoBounces), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *VoxelGIData) IsUsingTwoBounces() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVoxelGIData.fnIsUsingTwoBounces), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
