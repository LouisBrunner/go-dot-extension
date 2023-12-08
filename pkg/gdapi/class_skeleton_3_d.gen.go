// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Skeleton3D struct {
  obj gdc.ObjectPtr
}

func (me *Skeleton3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Skeleton3D) BaseClass() string {
  return "Skeleton3D"
}

// TODO: needed?
// const (
// // )

var (
  Skeleton3DNotificationUpdateSkeleton = "50" // TODO: construct correctly
)

func  (me *Skeleton3D) AddBone(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) FindBone(name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBoneName(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetBoneName(bone_idx int, name String, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBoneParent(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetBoneParent(bone_idx int, parent_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBoneCount() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetVersion() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) UnparentBoneAndRest(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBoneChildren(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetParentlessBones() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBoneRest(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetBoneRest(bone_idx int, rest Transform3D, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBoneGlobalRest(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) CreateSkinFromRestTransforms() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) RegisterSkin(skin Skin, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) LocalizeRests() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) ClearBones() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBonePose(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetBonePosePosition(bone_idx int, position Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetBonePoseRotation(bone_idx int, rotation Quaternion, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetBonePoseScale(bone_idx int, scale Vector3, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBonePosePosition(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBonePoseRotation(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBonePoseScale(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) ResetBonePose(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) ResetBonePoses() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) IsBoneEnabled(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetBoneEnabled(bone_idx int, enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) ClearBonesGlobalPoseOverride() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetBoneGlobalPoseOverride(bone_idx int, pose Transform3D, amount float32, persistent bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBoneGlobalPoseOverride(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBoneGlobalPose(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetBoneGlobalPoseNoOverride(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) ForceUpdateAllBoneTransforms() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) ForceUpdateBoneChildTransform(bone_idx int, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetMotionScale(motion_scale float32, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetMotionScale() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetShowRestOnly(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) IsShowRestOnly() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) SetAnimatePhysicalBones(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) GetAnimatePhysicalBones() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) PhysicalBonesStopSimulation() { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) PhysicalBonesStartSimulation(bones StringName, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) PhysicalBonesAddCollisionException(exception RID, ) { // TODO: return value
  // TODO: implement
}

func  (me *Skeleton3D) PhysicalBonesRemoveCollisionException(exception RID, ) { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
