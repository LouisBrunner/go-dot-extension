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

type ptrsForBoneMapList struct {
	fnGetProfile          gdc.MethodBindPtr
	fnSetProfile          gdc.MethodBindPtr
	fnGetSkeletonBoneName gdc.MethodBindPtr
	fnSetSkeletonBoneName gdc.MethodBindPtr
	fnFindProfileBoneName gdc.MethodBindPtr
}

var ptrsForBoneMap ptrsForBoneMapList

func initBoneMapPtrs(iface gdc.Interface) {

	className := StringNameFromStr("BoneMap")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_profile")
		defer methodName.Destroy()
		ptrsForBoneMap.fnGetProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4291782652))
	}
	{
		methodName := StringNameFromStr("set_profile")
		defer methodName.Destroy()
		ptrsForBoneMap.fnSetProfile = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3870374136))
	}
	{
		methodName := StringNameFromStr("get_skeleton_bone_name")
		defer methodName.Destroy()
		ptrsForBoneMap.fnGetSkeletonBoneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965194235))
	}
	{
		methodName := StringNameFromStr("set_skeleton_bone_name")
		defer methodName.Destroy()
		ptrsForBoneMap.fnSetSkeletonBoneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3740211285))
	}
	{
		methodName := StringNameFromStr("find_profile_bone_name")
		defer methodName.Destroy()
		ptrsForBoneMap.fnFindProfileBoneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965194235))
	}

}

type BoneMap struct {
	Resource
}

func (me *BoneMap) BaseClass() string {
	return "BoneMap"
}

func NewBoneMap() *BoneMap {
	str := StringNameFromStr("BoneMap") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &BoneMap{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *BoneMap) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *BoneMap) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *BoneMap) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *BoneMap) GetProfile() SkeletonProfile {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkeletonProfile()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneMap.fnGetProfile), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BoneMap) SetProfile(profile SkeletonProfile) {
	cargs := []gdc.ConstTypePtr{profile.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneMap.fnSetProfile), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoneMap) GetSkeletonBoneName(profile_bone_name StringName) StringName {
	cargs := []gdc.ConstTypePtr{profile_bone_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneMap.fnGetSkeletonBoneName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *BoneMap) SetSkeletonBoneName(profile_bone_name StringName, skeleton_bone_name StringName) {
	cargs := []gdc.ConstTypePtr{profile_bone_name.AsCTypePtr(), skeleton_bone_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneMap.fnSetSkeletonBoneName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *BoneMap) FindProfileBoneName(skeleton_bone_name StringName) StringName {
	cargs := []gdc.ConstTypePtr{skeleton_bone_name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForBoneMap.fnFindProfileBoneName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type BoneMapBoneMapUpdatedSignalFn func()

func (me *BoneMap) ConnectBoneMapUpdated(subs SignalSubscribers, fn BoneMapBoneMapUpdatedSignalFn) {
	sig := StringNameFromStr("bone_map_updated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *BoneMap) DisconnectBoneMapUpdated(subs SignalSubscribers, fn BoneMapBoneMapUpdatedSignalFn) {
	sig := StringNameFromStr("bone_map_updated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type BoneMapProfileUpdatedSignalFn func()

func (me *BoneMap) ConnectProfileUpdated(subs SignalSubscribers, fn BoneMapProfileUpdatedSignalFn) {
	sig := StringNameFromStr("profile_updated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *BoneMap) DisconnectProfileUpdated(subs SignalSubscribers, fn BoneMapProfileUpdatedSignalFn) {
	sig := StringNameFromStr("profile_updated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}
