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

type SkeletonProfileTailDirection int
const (
  SkeletonProfileTailDirectionTailDirectionAverageChildren SkeletonProfileTailDirection = 0
  SkeletonProfileTailDirectionTailDirectionSpecificChild SkeletonProfileTailDirection = 1
  SkeletonProfileTailDirectionTailDirectionEnd SkeletonProfileTailDirection = 2
)

func  (me *SkeletonProfile) SetRootBone(bone_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetRootBone() { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetScaleBaseBone(bone_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetScaleBaseBone() { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetGroupSize(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetGroupSize() { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetGroupName(group_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetGroupName(group_idx int, group_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetTexture(group_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetTexture(group_idx int, texture Texture2D, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetBoneSize(size int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetBoneSize() { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) FindBone(bone_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetBoneName(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetBoneName(bone_idx int, bone_name StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetBoneParent(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetBoneParent(bone_idx int, bone_parent StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetTailDirection(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetTailDirection(bone_idx int, tail_direction SkeletonProfileTailDirection, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetBoneTail(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetBoneTail(bone_idx int, bone_tail StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetReferencePose(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetReferencePose(bone_idx int, bone_name Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetHandleOffset(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetHandleOffset(bone_idx int, handle_offset Vector2, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) GetGroup(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *SkeletonProfile) SetGroup(bone_idx int, group StringName, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
