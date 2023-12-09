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



// Constants

var (
  Skeleton3DNotificationUpdateSkeleton = "50" // TODO: construct correctly
)

// Enums

func (me *Skeleton3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Skeleton3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Skeleton3D) AddBone(name String, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) FindBone(name String, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBoneName(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetBoneName(bone_idx int, name String, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBoneParent(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetBoneParent(bone_idx int, parent_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBoneCount()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetVersion()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) UnparentBoneAndRest(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBoneChildren(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetParentlessBones()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBoneRest(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetBoneRest(bone_idx int, rest Transform3D, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBoneGlobalRest(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) CreateSkinFromRestTransforms()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) RegisterSkin(skin Skin, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) LocalizeRests()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) ClearBones()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBonePose(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetBonePosePosition(bone_idx int, position Vector3, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetBonePoseRotation(bone_idx int, rotation Quaternion, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetBonePoseScale(bone_idx int, scale Vector3, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBonePosePosition(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBonePoseRotation(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBonePoseScale(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) ResetBonePose(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) ResetBonePoses()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) IsBoneEnabled(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetBoneEnabled(bone_idx int, enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) ClearBonesGlobalPoseOverride()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetBoneGlobalPoseOverride(bone_idx int, pose Transform3D, amount float32, persistent bool, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBoneGlobalPoseOverride(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBoneGlobalPose(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetBoneGlobalPoseNoOverride(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) ForceUpdateAllBoneTransforms()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) ForceUpdateBoneChildTransform(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetMotionScale(motion_scale float32, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetMotionScale()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetShowRestOnly(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) IsShowRestOnly()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) SetAnimatePhysicalBones(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) GetAnimatePhysicalBones()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) PhysicalBonesStopSimulation()  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) PhysicalBonesStartSimulation(bones StringName, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) PhysicalBonesAddCollisionException(exception RID, )  {
  panic("TODO: implement")
}

func  (me *Skeleton3D) PhysicalBonesRemoveCollisionException(exception RID, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
