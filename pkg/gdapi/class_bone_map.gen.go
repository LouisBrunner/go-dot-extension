// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type BoneMap struct {
  obj gdc.ObjectPtr
}

func (me *BoneMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BoneMap) BaseClass() string {
  return "BoneMap"
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
  var ret SkeletonProfile
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BoneMap) SetProfile(profile SkeletonProfile, )  {
  classNameV := StringNameFromStr("BoneMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_profile")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3870374136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(profile.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BoneMap) GetSkeletonBoneName(profile_bone_name StringName, ) StringName {
  classNameV := StringNameFromStr("BoneMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965194235) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(profile_bone_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *BoneMap) SetSkeletonBoneName(profile_bone_name StringName, skeleton_bone_name StringName, )  {
  classNameV := StringNameFromStr("BoneMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skeleton_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3740211285) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(profile_bone_name.AsCTypePtr()), gdc.ConstTypePtr(skeleton_bone_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *BoneMap) FindProfileBoneName(skeleton_bone_name StringName, ) StringName {
  classNameV := StringNameFromStr("BoneMap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_profile_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965194235) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(skeleton_bone_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *BoneMap) GetPropProfile() SkeletonProfile {
  panic("TODO: implement")
}

func (me *BoneMap) SetPropProfile(value SkeletonProfile) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API