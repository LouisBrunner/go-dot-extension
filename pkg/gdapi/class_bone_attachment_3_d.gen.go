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

type ptrsForBoneAttachment3DList struct {
	fnSetBoneName            gdc.MethodBindPtr
	fnGetBoneName            gdc.MethodBindPtr
	fnSetBoneIdx             gdc.MethodBindPtr
	fnGetBoneIdx             gdc.MethodBindPtr
	fnOnBonePoseUpdate       gdc.MethodBindPtr
	fnSetOverridePose        gdc.MethodBindPtr
	fnGetOverridePose        gdc.MethodBindPtr
	fnSetUseExternalSkeleton gdc.MethodBindPtr
	fnGetUseExternalSkeleton gdc.MethodBindPtr
	fnSetExternalSkeleton    gdc.MethodBindPtr
	fnGetExternalSkeleton    gdc.MethodBindPtr
}

var ptrsForBoneAttachment3D ptrsForBoneAttachment3DList

func initBoneAttachment3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("BoneAttachment3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_bone_name")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnSetBoneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 83702148))
	}
	{
		methodName := StringNameFromStr("get_bone_name")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnGetBoneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 201670096))
	}
	{
		methodName := StringNameFromStr("set_bone_idx")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnSetBoneIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_bone_idx")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnGetBoneIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("on_bone_pose_update")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnOnBonePoseUpdate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_override_pose")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnSetOverridePose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_override_pose")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnGetOverridePose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_external_skeleton")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnSetUseExternalSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_use_external_skeleton")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnGetUseExternalSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_external_skeleton")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnSetExternalSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_external_skeleton")
		defer methodName.Destroy()
		ptrsForBoneAttachment3D.fnGetExternalSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
}

type BoneAttachment3D struct {
	Node3D
}

func (me *BoneAttachment3D) BaseClass() string {
	return "BoneAttachment3D"
}

func NewBoneAttachment3D() *BoneAttachment3D {
	str := StringNameFromStr("BoneAttachment3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &BoneAttachment3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *BoneAttachment3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *BoneAttachment3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *BoneAttachment3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *BoneAttachment3D) SetBoneName(bone_name String) {
	cargs := []gdc.ConstTypePtr{bone_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnSetBoneName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoneAttachment3D) GetBoneName() String {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnGetBoneName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BoneAttachment3D) SetBoneIdx(bone_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnSetBoneIdx), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoneAttachment3D) GetBoneIdx() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnGetBoneIdx), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BoneAttachment3D) OnBonePoseUpdate(bone_index int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnOnBonePoseUpdate), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoneAttachment3D) SetOverridePose(override_pose bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&override_pose)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnSetOverridePose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoneAttachment3D) GetOverridePose() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnGetOverridePose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BoneAttachment3D) SetUseExternalSkeleton(use_external_skeleton bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_external_skeleton)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnSetUseExternalSkeleton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoneAttachment3D) GetUseExternalSkeleton() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnGetUseExternalSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *BoneAttachment3D) SetExternalSkeleton(external_skeleton NodePath) {
	cargs := []gdc.ConstTypePtr{external_skeleton.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnSetExternalSkeleton), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoneAttachment3D) GetExternalSkeleton() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneAttachment3D.fnGetExternalSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
