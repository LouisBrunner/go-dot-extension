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

type ptrsForSkeletonProfileList struct {
	fnSetRootBone      gdc.MethodBindPtr
	fnGetRootBone      gdc.MethodBindPtr
	fnSetScaleBaseBone gdc.MethodBindPtr
	fnGetScaleBaseBone gdc.MethodBindPtr
	fnSetGroupSize     gdc.MethodBindPtr
	fnGetGroupSize     gdc.MethodBindPtr
	fnGetGroupName     gdc.MethodBindPtr
	fnSetGroupName     gdc.MethodBindPtr
	fnGetTexture       gdc.MethodBindPtr
	fnSetTexture       gdc.MethodBindPtr
	fnSetBoneSize      gdc.MethodBindPtr
	fnGetBoneSize      gdc.MethodBindPtr
	fnFindBone         gdc.MethodBindPtr
	fnGetBoneName      gdc.MethodBindPtr
	fnSetBoneName      gdc.MethodBindPtr
	fnGetBoneParent    gdc.MethodBindPtr
	fnSetBoneParent    gdc.MethodBindPtr
	fnGetTailDirection gdc.MethodBindPtr
	fnSetTailDirection gdc.MethodBindPtr
	fnGetBoneTail      gdc.MethodBindPtr
	fnSetBoneTail      gdc.MethodBindPtr
	fnGetReferencePose gdc.MethodBindPtr
	fnSetReferencePose gdc.MethodBindPtr
	fnGetHandleOffset  gdc.MethodBindPtr
	fnSetHandleOffset  gdc.MethodBindPtr
	fnGetGroup         gdc.MethodBindPtr
	fnSetGroup         gdc.MethodBindPtr
}

var ptrsForSkeletonProfile ptrsForSkeletonProfileList

func initSkeletonProfilePtrs(iface gdc.Interface) {

	className := StringNameFromStr("SkeletonProfile")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_root_bone")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetRootBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_root_bone")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetRootBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2737447660))
	}
	{
		methodName := StringNameFromStr("set_scale_base_bone")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetScaleBaseBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_scale_base_bone")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetScaleBaseBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2737447660))
	}
	{
		methodName := StringNameFromStr("set_group_size")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetGroupSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_group_size")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetGroupSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("get_group_name")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetGroupName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("set_group_name")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetGroupName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780747571))
	}
	{
		methodName := StringNameFromStr("get_texture")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3536238170))
	}
	{
		methodName := StringNameFromStr("set_texture")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetTexture = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 666127730))
	}
	{
		methodName := StringNameFromStr("set_bone_size")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetBoneSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_bone_size")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetBoneSize = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("find_bone")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnFindBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2458036349))
	}
	{
		methodName := StringNameFromStr("get_bone_name")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetBoneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("set_bone_name")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetBoneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780747571))
	}
	{
		methodName := StringNameFromStr("get_bone_parent")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetBoneParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("set_bone_parent")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetBoneParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780747571))
	}
	{
		methodName := StringNameFromStr("get_tail_direction")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetTailDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2675997574))
	}
	{
		methodName := StringNameFromStr("set_tail_direction")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetTailDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1231951015))
	}
	{
		methodName := StringNameFromStr("get_bone_tail")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetBoneTail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("set_bone_tail")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetBoneTail = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780747571))
	}
	{
		methodName := StringNameFromStr("get_reference_pose")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetReferencePose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("set_reference_pose")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetReferencePose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("get_handle_offset")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetHandleOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
	{
		methodName := StringNameFromStr("set_handle_offset")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetHandleOffset = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
	}
	{
		methodName := StringNameFromStr("get_group")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnGetGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659327637))
	}
	{
		methodName := StringNameFromStr("set_group")
		defer methodName.Destroy()
		ptrsForSkeletonProfile.fnSetGroup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3780747571))
	}
}

type SkeletonProfile struct {
	Resource
}

func (me *SkeletonProfile) BaseClass() string {
	return "SkeletonProfile"
}

func NewSkeletonProfile() *SkeletonProfile {
	str := StringNameFromStr("SkeletonProfile") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SkeletonProfile{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

type SkeletonProfileTailDirection int

const (
	SkeletonProfileTailDirectionTailDirectionAverageChildren SkeletonProfileTailDirection = 0
	SkeletonProfileTailDirectionTailDirectionSpecificChild   SkeletonProfileTailDirection = 1
	SkeletonProfileTailDirectionTailDirectionEnd             SkeletonProfileTailDirection = 2
)

func (me *SkeletonProfile) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SkeletonProfile) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SkeletonProfile) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SkeletonProfile) SetRootBone(bone_name StringName) {
	cargs := []gdc.ConstTypePtr{bone_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetRootBone), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetRootBone() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetRootBone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetScaleBaseBone(bone_name StringName) {
	cargs := []gdc.ConstTypePtr{bone_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetScaleBaseBone), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetScaleBaseBone() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetScaleBaseBone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetGroupSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetGroupSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetGroupSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetGroupSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonProfile) GetGroupName(group_idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&group_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetGroupName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetGroupName(group_idx int64, group_name StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_idx), group_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetGroupName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetTexture(group_idx int64) Texture2D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTexture2D()
	pinner.Pin(&group_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetTexture), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetTexture(group_idx int64, texture Texture2D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_idx), texture.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetTexture), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) SetBoneSize(size int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetBoneSize), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetBoneSize() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetBoneSize), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonProfile) FindBone(bone_name StringName) int64 {
	cargs := []gdc.ConstTypePtr{bone_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnFindBone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonProfile) GetBoneName(bone_idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetBoneName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetBoneName(bone_idx int64, bone_name StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), bone_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetBoneName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetBoneParent(bone_idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetBoneParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetBoneParent(bone_idx int64, bone_parent StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), bone_parent.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetBoneParent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetTailDirection(bone_idx int64) SkeletonProfileTailDirection {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret SkeletonProfileTailDirection
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetTailDirection), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *SkeletonProfile) SetTailDirection(bone_idx int64, tail_direction SkeletonProfileTailDirection) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(&tail_direction)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetTailDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetBoneTail(bone_idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetBoneTail), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetBoneTail(bone_idx int64, bone_tail StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), bone_tail.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetBoneTail), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetReferencePose(bone_idx int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetReferencePose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetReferencePose(bone_idx int64, bone_name Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), bone_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetReferencePose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetHandleOffset(bone_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetHandleOffset), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetHandleOffset(bone_idx int64, handle_offset Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), handle_offset.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetHandleOffset), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonProfile) GetGroup(bone_idx int64) StringName {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnGetGroup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonProfile) SetGroup(bone_idx int64, group StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), group.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonProfile.fnSetGroup), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SkeletonProfileProfileUpdatedSignalFn func()

func (me *SkeletonProfile) ConnectProfileUpdated(subs SignalSubscribers, fn SkeletonProfileProfileUpdatedSignalFn) {
	sig := StringNameFromStr("profile_updated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *SkeletonProfile) DisconnectProfileUpdated(subs SignalSubscribers, fn SkeletonProfileProfileUpdatedSignalFn) {
	sig := StringNameFromStr("profile_updated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
