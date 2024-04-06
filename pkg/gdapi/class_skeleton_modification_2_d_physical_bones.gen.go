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

type ptrsForSkeletonModification2DPhysicalBonesList struct {
	fnSetPhysicalBoneChainLength gdc.MethodBindPtr
	fnGetPhysicalBoneChainLength gdc.MethodBindPtr
	fnSetPhysicalBoneNode        gdc.MethodBindPtr
	fnGetPhysicalBoneNode        gdc.MethodBindPtr
	fnFetchPhysicalBones         gdc.MethodBindPtr
	fnStartSimulation            gdc.MethodBindPtr
	fnStopSimulation             gdc.MethodBindPtr
}

var ptrsForSkeletonModification2DPhysicalBones ptrsForSkeletonModification2DPhysicalBonesList

func initSkeletonModification2DPhysicalBonesPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SkeletonModification2DPhysicalBones")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_physical_bone_chain_length")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DPhysicalBones.fnSetPhysicalBoneChainLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_physical_bone_chain_length")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DPhysicalBones.fnGetPhysicalBoneChainLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_physical_bone_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DPhysicalBones.fnSetPhysicalBoneNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761262315))
	}
	{
		methodName := StringNameFromStr("get_physical_bone_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DPhysicalBones.fnGetPhysicalBoneNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
	}
	{
		methodName := StringNameFromStr("fetch_physical_bones")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DPhysicalBones.fnFetchPhysicalBones = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("start_simulation")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DPhysicalBones.fnStartSimulation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2787316981))
	}
	{
		methodName := StringNameFromStr("stop_simulation")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DPhysicalBones.fnStopSimulation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2787316981))
	}
}

type SkeletonModification2DPhysicalBones struct {
	SkeletonModification2D
}

func (me *SkeletonModification2DPhysicalBones) BaseClass() string {
	return "SkeletonModification2DPhysicalBones"
}

func NewSkeletonModification2DPhysicalBones() *SkeletonModification2DPhysicalBones {
	str := StringNameFromStr("SkeletonModification2DPhysicalBones") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SkeletonModification2DPhysicalBones{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SkeletonModification2DPhysicalBones) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SkeletonModification2DPhysicalBones) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DPhysicalBones) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SkeletonModification2DPhysicalBones) SetPhysicalBoneChainLength(length int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DPhysicalBones.fnSetPhysicalBoneChainLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DPhysicalBones) GetPhysicalBoneChainLength() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DPhysicalBones.fnGetPhysicalBoneChainLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DPhysicalBones) SetPhysicalBoneNode(joint_idx int64, physicalbone2d_node NodePath) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), physicalbone2d_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DPhysicalBones.fnSetPhysicalBoneNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DPhysicalBones) GetPhysicalBoneNode(joint_idx int64) NodePath {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()
	pinner.Pin(&joint_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DPhysicalBones.fnGetPhysicalBoneNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonModification2DPhysicalBones) FetchPhysicalBones() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DPhysicalBones.fnFetchPhysicalBones), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DPhysicalBones) StartSimulation(bones []StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bones)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DPhysicalBones.fnStartSimulation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DPhysicalBones) StopSimulation(bones []StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bones)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DPhysicalBones.fnStopSimulation), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
