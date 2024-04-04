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

func  (me *BoneMap) GetProfile() SkeletonProfile {
  classNameV := StringNameFromStr("BoneMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4291782652) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkeletonProfile()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BoneMap) SetProfile(profile SkeletonProfile, )  {
  classNameV := StringNameFromStr("BoneMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3870374136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{profile.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoneMap) GetSkeletonBoneName(profile_bone_name StringName, ) StringName {
  classNameV := StringNameFromStr("BoneMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965194235) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{profile_bone_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BoneMap) SetSkeletonBoneName(profile_bone_name StringName, skeleton_bone_name StringName, )  {
  classNameV := StringNameFromStr("BoneMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skeleton_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{profile_bone_name.AsCTypePtr(), skeleton_bone_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoneMap) FindProfileBoneName(skeleton_bone_name StringName, ) StringName {
  classNameV := StringNameFromStr("BoneMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_profile_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965194235) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{skeleton_bone_name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
