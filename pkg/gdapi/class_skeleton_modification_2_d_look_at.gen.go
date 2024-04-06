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

type ptrsForSkeletonModification2DLookAtList struct {
	fnSetBone2DNode            gdc.MethodBindPtr
	fnGetBone2DNode            gdc.MethodBindPtr
	fnSetBoneIndex             gdc.MethodBindPtr
	fnGetBoneIndex             gdc.MethodBindPtr
	fnSetTargetNode            gdc.MethodBindPtr
	fnGetTargetNode            gdc.MethodBindPtr
	fnSetAdditionalRotation    gdc.MethodBindPtr
	fnGetAdditionalRotation    gdc.MethodBindPtr
	fnSetEnableConstraint      gdc.MethodBindPtr
	fnGetEnableConstraint      gdc.MethodBindPtr
	fnSetConstraintAngleMin    gdc.MethodBindPtr
	fnGetConstraintAngleMin    gdc.MethodBindPtr
	fnSetConstraintAngleMax    gdc.MethodBindPtr
	fnGetConstraintAngleMax    gdc.MethodBindPtr
	fnSetConstraintAngleInvert gdc.MethodBindPtr
	fnGetConstraintAngleInvert gdc.MethodBindPtr
}

var ptrsForSkeletonModification2DLookAt ptrsForSkeletonModification2DLookAtList

func initSkeletonModification2DLookAtPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SkeletonModification2DLookAt")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_bone2d_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnSetBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_bone2d_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnGetBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_bone_index")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnSetBoneIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_bone_index")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnGetBoneIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_target_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnSetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_target_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnGetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_additional_rotation")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnSetAdditionalRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_additional_rotation")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnGetAdditionalRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_enable_constraint")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnSetEnableConstraint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_enable_constraint")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnGetEnableConstraint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_constraint_angle_min")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnSetConstraintAngleMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_constraint_angle_min")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnGetConstraintAngleMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_constraint_angle_max")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnSetConstraintAngleMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_constraint_angle_max")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnGetConstraintAngleMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_constraint_angle_invert")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnSetConstraintAngleInvert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_constraint_angle_invert")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DLookAt.fnGetConstraintAngleInvert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type SkeletonModification2DLookAt struct {
	SkeletonModification2D
}

func (me *SkeletonModification2DLookAt) BaseClass() string {
	return "SkeletonModification2DLookAt"
}

func NewSkeletonModification2DLookAt() *SkeletonModification2DLookAt {
	str := StringNameFromStr("SkeletonModification2DLookAt") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SkeletonModification2DLookAt{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SkeletonModification2DLookAt) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SkeletonModification2DLookAt) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DLookAt) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SkeletonModification2DLookAt) SetBone2DNode(bone2d_nodepath NodePath) {
	cargs := []gdc.ConstTypePtr{bone2d_nodepath.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnSetBone2DNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DLookAt) GetBone2DNode() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnGetBone2DNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonModification2DLookAt) SetBoneIndex(bone_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnSetBoneIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DLookAt) GetBoneIndex() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnGetBoneIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DLookAt) SetTargetNode(target_nodepath NodePath) {
	cargs := []gdc.ConstTypePtr{target_nodepath.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnSetTargetNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DLookAt) GetTargetNode() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnGetTargetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonModification2DLookAt) SetAdditionalRotation(rotation float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnSetAdditionalRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DLookAt) GetAdditionalRotation() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnGetAdditionalRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DLookAt) SetEnableConstraint(enable_constraint bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable_constraint)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnSetEnableConstraint), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DLookAt) GetEnableConstraint() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnGetEnableConstraint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DLookAt) SetConstraintAngleMin(angle_min float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle_min)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnSetConstraintAngleMin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DLookAt) GetConstraintAngleMin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnGetConstraintAngleMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DLookAt) SetConstraintAngleMax(angle_max float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle_max)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnSetConstraintAngleMax), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DLookAt) GetConstraintAngleMax() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnGetConstraintAngleMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DLookAt) SetConstraintAngleInvert(invert bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnSetConstraintAngleInvert), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DLookAt) GetConstraintAngleInvert() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DLookAt.fnGetConstraintAngleInvert), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
