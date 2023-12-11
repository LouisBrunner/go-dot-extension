// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonProfile struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonProfile) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonProfile) BaseClass() string {
  return "SkeletonProfile"
}



// Enums

type SkeletonProfileTailDirection int
const (
  SkeletonProfileTailDirectionTailDirectionAverageChildren SkeletonProfileTailDirection = 0
  SkeletonProfileTailDirectionTailDirectionSpecificChild SkeletonProfileTailDirection = 1
  SkeletonProfileTailDirectionTailDirectionEnd SkeletonProfileTailDirection = 2
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

func  (me *SkeletonProfile) SetRootBone(bone_name StringName, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_root_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bone_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetRootBone() StringName {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_root_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2737447660) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetScaleBaseBone(bone_name StringName, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale_base_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bone_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetScaleBaseBone() StringName {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale_base_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2737447660) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetGroupSize(size int, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_group_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetGroupSize() int {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_group_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) GetGroupName(group_idx int, ) StringName {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_group_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetGroupName(group_idx int, group_name StringName, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_group_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_idx), gdc.ConstTypePtr(group_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetTexture(group_idx int, ) Texture2D {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3536238170) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetTexture(group_idx int, texture Texture2D, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 666127730) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&group_idx), gdc.ConstTypePtr(texture.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) SetBoneSize(size int, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&size), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetBoneSize() int {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_size")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) FindBone(bone_name StringName, ) int {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2458036349) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bone_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) GetBoneName(bone_idx int, ) StringName {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetBoneName(bone_idx int, bone_name StringName, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(bone_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetBoneParent(bone_idx int, ) StringName {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetBoneParent(bone_idx int, bone_parent StringName, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(bone_parent.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetTailDirection(bone_idx int, ) SkeletonProfileTailDirection {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tail_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2675997574) // FIXME: should cache?
  var ret SkeletonProfileTailDirection
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetTailDirection(bone_idx int, tail_direction SkeletonProfileTailDirection, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tail_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1231951015) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(&tail_direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetBoneTail(bone_idx int, ) StringName {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_tail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetBoneTail(bone_idx int, bone_tail StringName, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_tail")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(bone_tail.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetReferencePose(bone_idx int, ) Transform3D {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reference_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetReferencePose(bone_idx int, bone_name Transform3D, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reference_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(bone_name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetHandleOffset(bone_idx int, ) Vector2 {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_handle_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetHandleOffset(bone_idx int, handle_offset Vector2, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_handle_offset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(handle_offset.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonProfile) GetGroup(bone_idx int, ) StringName {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonProfile) SetGroup(bone_idx int, group StringName, )  {
  classNameV := StringNameFromStr("SkeletonProfile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_group")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(group.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type SkeletonProfileProfileUpdatedSignalFn func()

func (me *SkeletonProfile) ConnectProfileUpdated(subs SignalSubscribers, fn SkeletonProfileProfileUpdatedSignalFn) {
  sig := StringNameFromStr("profile_updated")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *SkeletonProfile) DisconnectProfileUpdated(subs SignalSubscribers, fn SkeletonProfileProfileUpdatedSignalFn) {
  sig := StringNameFromStr("profile_updated")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
