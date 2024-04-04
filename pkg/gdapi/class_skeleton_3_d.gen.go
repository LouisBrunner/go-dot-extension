// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type Skeleton3D struct {
  Node3D
}

func (me *Skeleton3D) BaseClass() string {
  return "Skeleton3D"
}

func NewSkeleton3D() *Skeleton3D {
  str := StringNameFromStr("Skeleton3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Skeleton3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Constants

var (
  Skeleton3DNotificationUpdateSkeleton = "50" // TODO: construct correctly
)

// Enums

func (me *Skeleton3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Skeleton3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Skeleton3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Skeleton3D) AddBone(name String, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) FindBone(name String, ) int64 {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1321353865) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton3D) GetBoneName(bone_idx int64, ) String {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 844755477) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) SetBoneName(bone_idx int64, name String, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 501894301) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) GetBoneParent(bone_idx int64, ) int64 {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton3D) SetBoneParent(bone_idx int64, parent_idx int64, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , gdc.ConstTypePtr(&parent_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) GetBoneCount() int64 {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton3D) GetVersion() int64 {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_version")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton3D) UnparentBoneAndRest(bone_idx int64, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unparent_bone_and_rest")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) GetBoneChildren(bone_idx int64, ) PackedInt32Array {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_children")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1706082319) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) GetParentlessBones() PackedInt32Array {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parentless_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1930428628) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedInt32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) GetBoneRest(bone_idx int64, ) Transform3D {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_rest")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) SetBoneRest(bone_idx int64, rest Transform3D, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_rest")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , rest.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) GetBoneGlobalRest(bone_idx int64, ) Transform3D {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_global_rest")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) CreateSkinFromRestTransforms() Skin {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("create_skin_from_rest_transforms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1032037385) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) RegisterSkin(skin Skin, ) SkinReference {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("register_skin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3405789568) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{skin.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkinReference()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) LocalizeRests()  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("localize_rests")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) ClearBones()  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) GetBonePose(bone_idx int64, ) Transform3D {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) SetBonePosePosition(bone_idx int64, position Vector3, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_pose_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) SetBonePoseRotation(bone_idx int64, rotation Quaternion, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_pose_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2823819782) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , rotation.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) SetBonePoseScale(bone_idx int64, scale Vector3, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_pose_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1530502735) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) GetBonePosePosition(bone_idx int64, ) Vector3 {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_pose_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) GetBonePoseRotation(bone_idx int64, ) Quaternion {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_pose_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 476865136) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewQuaternion()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) GetBonePoseScale(bone_idx int64, ) Vector3 {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_pose_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711720468) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector3()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) ResetBonePose(bone_idx int64, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reset_bone_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) ResetBonePoses()  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reset_bone_poses")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) IsBoneEnabled(bone_idx int64, ) bool {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_bone_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton3D) SetBoneEnabled(bone_idx int64, enabled bool, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 972357352) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) ClearBonesGlobalPoseOverride()  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_bones_global_pose_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) SetBoneGlobalPoseOverride(bone_idx int64, pose Transform3D, amount float64, persistent bool, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_global_pose_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3483398371) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , pose.AsCTypePtr(), gdc.ConstTypePtr(&amount) , gdc.ConstTypePtr(&persistent) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) GetBoneGlobalPoseOverride(bone_idx int64, ) Transform3D {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_global_pose_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) GetBoneGlobalPose(bone_idx int64, ) Transform3D {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_global_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) GetBoneGlobalPoseNoOverride(bone_idx int64, ) Transform3D {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_global_pose_no_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton3D) ForceUpdateAllBoneTransforms()  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update_all_bone_transforms")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) ForceUpdateBoneChildTransform(bone_idx int64, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("force_update_bone_child_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) SetMotionScale(motion_scale float64, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_motion_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&motion_scale) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) GetMotionScale() float64 {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_motion_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton3D) SetShowRestOnly(enabled bool, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_show_rest_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) IsShowRestOnly() bool {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_show_rest_only")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton3D) SetAnimatePhysicalBones(enabled bool, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_animate_physical_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) GetAnimatePhysicalBones() bool {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_animate_physical_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton3D) PhysicalBonesStopSimulation()  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("physical_bones_stop_simulation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) PhysicalBonesStartSimulation(bones []StringName, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("physical_bones_start_simulation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2787316981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bones) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) PhysicalBonesAddCollisionException(exception RID, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("physical_bones_add_collision_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{exception.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton3D) PhysicalBonesRemoveCollisionException(exception RID, )  {
  classNameV := StringNameFromStr("Skeleton3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("physical_bones_remove_collision_exception")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2722037293) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{exception.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type Skeleton3DPoseUpdatedSignalFn func()

func (me *Skeleton3D) ConnectPoseUpdated(subs SignalSubscribers, fn Skeleton3DPoseUpdatedSignalFn) {
  sig := StringNameFromStr("pose_updated")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Skeleton3D) DisconnectPoseUpdated(subs SignalSubscribers, fn Skeleton3DPoseUpdatedSignalFn) {
  sig := StringNameFromStr("pose_updated")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Skeleton3DBonePoseChangedSignalFn func(bone_idx int, )

func (me *Skeleton3D) ConnectBonePoseChanged(subs SignalSubscribers, fn Skeleton3DBonePoseChangedSignalFn) {
  sig := StringNameFromStr("bone_pose_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Skeleton3D) DisconnectBonePoseChanged(subs SignalSubscribers, fn Skeleton3DBonePoseChangedSignalFn) {
  sig := StringNameFromStr("bone_pose_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Skeleton3DBoneEnabledChangedSignalFn func(bone_idx int, )

func (me *Skeleton3D) ConnectBoneEnabledChanged(subs SignalSubscribers, fn Skeleton3DBoneEnabledChangedSignalFn) {
  sig := StringNameFromStr("bone_enabled_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Skeleton3D) DisconnectBoneEnabledChanged(subs SignalSubscribers, fn Skeleton3DBoneEnabledChangedSignalFn) {
  sig := StringNameFromStr("bone_enabled_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}

type Skeleton3DShowRestOnlyChangedSignalFn func()

func (me *Skeleton3D) ConnectShowRestOnlyChanged(subs SignalSubscribers, fn Skeleton3DShowRestOnlyChangedSignalFn) {
  sig := StringNameFromStr("show_rest_only_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Skeleton3D) DisconnectShowRestOnlyChanged(subs SignalSubscribers, fn Skeleton3DShowRestOnlyChangedSignalFn) {
  sig := StringNameFromStr("show_rest_only_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
