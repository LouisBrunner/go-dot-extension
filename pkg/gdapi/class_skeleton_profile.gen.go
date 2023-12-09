// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetRootBone()  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetScaleBaseBone(bone_name StringName, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetScaleBaseBone()  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetGroupSize(size int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetGroupSize()  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetGroupName(group_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetGroupName(group_idx int, group_name StringName, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetTexture(group_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetTexture(group_idx int, texture Texture2D, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetBoneSize(size int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetBoneSize()  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) FindBone(bone_name StringName, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetBoneName(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetBoneName(bone_idx int, bone_name StringName, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetBoneParent(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetBoneParent(bone_idx int, bone_parent StringName, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetTailDirection(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetTailDirection(bone_idx int, tail_direction SkeletonProfileTailDirection, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetBoneTail(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetBoneTail(bone_idx int, bone_tail StringName, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetReferencePose(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetReferencePose(bone_idx int, bone_name Transform3D, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetHandleOffset(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetHandleOffset(bone_idx int, handle_offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) GetGroup(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonProfile) SetGroup(bone_idx int, group StringName, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
