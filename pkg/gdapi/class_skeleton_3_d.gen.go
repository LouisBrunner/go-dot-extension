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

type ptrsForSkeleton3DList struct {
	fnAddBone                               gdc.MethodBindPtr
	fnFindBone                              gdc.MethodBindPtr
	fnGetBoneName                           gdc.MethodBindPtr
	fnSetBoneName                           gdc.MethodBindPtr
	fnGetBoneParent                         gdc.MethodBindPtr
	fnSetBoneParent                         gdc.MethodBindPtr
	fnGetBoneCount                          gdc.MethodBindPtr
	fnGetVersion                            gdc.MethodBindPtr
	fnUnparentBoneAndRest                   gdc.MethodBindPtr
	fnGetBoneChildren                       gdc.MethodBindPtr
	fnGetParentlessBones                    gdc.MethodBindPtr
	fnGetBoneRest                           gdc.MethodBindPtr
	fnSetBoneRest                           gdc.MethodBindPtr
	fnGetBoneGlobalRest                     gdc.MethodBindPtr
	fnCreateSkinFromRestTransforms          gdc.MethodBindPtr
	fnRegisterSkin                          gdc.MethodBindPtr
	fnLocalizeRests                         gdc.MethodBindPtr
	fnClearBones                            gdc.MethodBindPtr
	fnGetBonePose                           gdc.MethodBindPtr
	fnSetBonePose                           gdc.MethodBindPtr
	fnSetBonePosePosition                   gdc.MethodBindPtr
	fnSetBonePoseRotation                   gdc.MethodBindPtr
	fnSetBonePoseScale                      gdc.MethodBindPtr
	fnGetBonePosePosition                   gdc.MethodBindPtr
	fnGetBonePoseRotation                   gdc.MethodBindPtr
	fnGetBonePoseScale                      gdc.MethodBindPtr
	fnResetBonePose                         gdc.MethodBindPtr
	fnResetBonePoses                        gdc.MethodBindPtr
	fnIsBoneEnabled                         gdc.MethodBindPtr
	fnSetBoneEnabled                        gdc.MethodBindPtr
	fnGetBoneGlobalPose                     gdc.MethodBindPtr
	fnSetBoneGlobalPose                     gdc.MethodBindPtr
	fnForceUpdateAllBoneTransforms          gdc.MethodBindPtr
	fnForceUpdateBoneChildTransform         gdc.MethodBindPtr
	fnSetMotionScale                        gdc.MethodBindPtr
	fnGetMotionScale                        gdc.MethodBindPtr
	fnSetShowRestOnly                       gdc.MethodBindPtr
	fnIsShowRestOnly                        gdc.MethodBindPtr
	fnSetModifierCallbackModeProcess        gdc.MethodBindPtr
	fnGetModifierCallbackModeProcess        gdc.MethodBindPtr
	fnClearBonesGlobalPoseOverride          gdc.MethodBindPtr
	fnSetBoneGlobalPoseOverride             gdc.MethodBindPtr
	fnGetBoneGlobalPoseOverride             gdc.MethodBindPtr
	fnGetBoneGlobalPoseNoOverride           gdc.MethodBindPtr
	fnSetAnimatePhysicalBones               gdc.MethodBindPtr
	fnGetAnimatePhysicalBones               gdc.MethodBindPtr
	fnPhysicalBonesStopSimulation           gdc.MethodBindPtr
	fnPhysicalBonesStartSimulation          gdc.MethodBindPtr
	fnPhysicalBonesAddCollisionException    gdc.MethodBindPtr
	fnPhysicalBonesRemoveCollisionException gdc.MethodBindPtr
}

var ptrsForSkeleton3D ptrsForSkeleton3DList

func initSkeleton3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("Skeleton3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("add_bone")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnAddBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1597066294))
	}
	{
		methodName := StringNameFromStr("find_bone")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnFindBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1321353865))
	}
	{
		methodName := StringNameFromStr("get_bone_name")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBoneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 844755477))
	}
	{
		methodName := StringNameFromStr("set_bone_name")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBoneName = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 501894301))
	}
	{
		methodName := StringNameFromStr("get_bone_parent")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBoneParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_bone_parent")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBoneParent = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_bone_count")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBoneCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_version")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetVersion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("unparent_bone_and_rest")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnUnparentBoneAndRest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_bone_children")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBoneChildren = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1706082319))
	}
	{
		methodName := StringNameFromStr("get_parentless_bones")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetParentlessBones = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1930428628))
	}
	{
		methodName := StringNameFromStr("get_bone_rest")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBoneRest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("set_bone_rest")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBoneRest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("get_bone_global_rest")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBoneGlobalRest = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("create_skin_from_rest_transforms")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnCreateSkinFromRestTransforms = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1032037385))
	}
	{
		methodName := StringNameFromStr("register_skin")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnRegisterSkin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3405789568))
	}
	{
		methodName := StringNameFromStr("localize_rests")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnLocalizeRests = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("clear_bones")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnClearBones = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("get_bone_pose")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBonePose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("set_bone_pose")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBonePose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("set_bone_pose_position")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBonePosePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
	}
	{
		methodName := StringNameFromStr("set_bone_pose_rotation")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBonePoseRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2823819782))
	}
	{
		methodName := StringNameFromStr("set_bone_pose_scale")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBonePoseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1530502735))
	}
	{
		methodName := StringNameFromStr("get_bone_pose_position")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBonePosePosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("get_bone_pose_rotation")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBonePoseRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 476865136))
	}
	{
		methodName := StringNameFromStr("get_bone_pose_scale")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBonePoseScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 711720468))
	}
	{
		methodName := StringNameFromStr("reset_bone_pose")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnResetBonePose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("reset_bone_poses")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnResetBonePoses = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("is_bone_enabled")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnIsBoneEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_bone_enabled")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBoneEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 972357352))
	}
	{
		methodName := StringNameFromStr("get_bone_global_pose")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBoneGlobalPose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("set_bone_global_pose")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBoneGlobalPose = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3616898986))
	}
	{
		methodName := StringNameFromStr("force_update_all_bone_transforms")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnForceUpdateAllBoneTransforms = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("force_update_bone_child_transform")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnForceUpdateBoneChildTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("set_motion_scale")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetMotionScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_motion_scale")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetMotionScale = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_show_rest_only")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetShowRestOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_show_rest_only")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnIsShowRestOnly = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_modifier_callback_mode_process")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetModifierCallbackModeProcess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3916362634))
	}
	{
		methodName := StringNameFromStr("get_modifier_callback_mode_process")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetModifierCallbackModeProcess = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 997182536))
	}
	{
		methodName := StringNameFromStr("clear_bones_global_pose_override")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnClearBonesGlobalPoseOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("set_bone_global_pose_override")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetBoneGlobalPoseOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3483398371))
	}
	{
		methodName := StringNameFromStr("get_bone_global_pose_override")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBoneGlobalPoseOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("get_bone_global_pose_no_override")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetBoneGlobalPoseNoOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1965739696))
	}
	{
		methodName := StringNameFromStr("set_animate_physical_bones")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnSetAnimatePhysicalBones = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_animate_physical_bones")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnGetAnimatePhysicalBones = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("physical_bones_stop_simulation")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnPhysicalBonesStopSimulation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}
	{
		methodName := StringNameFromStr("physical_bones_start_simulation")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnPhysicalBonesStartSimulation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2787316981))
	}
	{
		methodName := StringNameFromStr("physical_bones_add_collision_exception")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnPhysicalBonesAddCollisionException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}
	{
		methodName := StringNameFromStr("physical_bones_remove_collision_exception")
		defer methodName.Destroy()
		ptrsForSkeleton3D.fnPhysicalBonesRemoveCollisionException = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2722037293))
	}

}

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
	Skeleton3DNotificationUpdateSkeleton = 50
)

// Enums

type Skeleton3DModifierCallbackModeProcess int

const (
	Skeleton3DModifierCallbackModeProcessModifierCallbackModeProcessPhysics Skeleton3DModifierCallbackModeProcess = 0
	Skeleton3DModifierCallbackModeProcessModifierCallbackModeProcessIdle    Skeleton3DModifierCallbackModeProcess = 1
)

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

func (me *Skeleton3D) AddBone(name String) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnAddBone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skeleton3D) FindBone(name String) int64 {
	cargs := []gdc.ConstTypePtr{name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnFindBone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skeleton3D) GetBoneName(bone_idx int64) String {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewString()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBoneName), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) SetBoneName(bone_idx int64, name String) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), name.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBoneName), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetBoneParent(bone_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBoneParent), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skeleton3D) SetBoneParent(bone_idx int64, parent_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(&parent_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBoneParent), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetBoneCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBoneCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skeleton3D) GetVersion() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetVersion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skeleton3D) UnparentBoneAndRest(bone_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnUnparentBoneAndRest), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetBoneChildren(bone_idx int64) PackedInt32Array {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBoneChildren), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) GetParentlessBones() PackedInt32Array {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPackedInt32Array()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetParentlessBones), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) GetBoneRest(bone_idx int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBoneRest), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) SetBoneRest(bone_idx int64, rest Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), rest.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBoneRest), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetBoneGlobalRest(bone_idx int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBoneGlobalRest), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) CreateSkinFromRestTransforms() Skin {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnCreateSkinFromRestTransforms), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) RegisterSkin(skin Skin) SkinReference {
	cargs := []gdc.ConstTypePtr{skin.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkinReference()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnRegisterSkin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) LocalizeRests() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnLocalizeRests), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) ClearBones() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnClearBones), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetBonePose(bone_idx int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBonePose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) SetBonePose(bone_idx int64, pose Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), pose.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBonePose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) SetBonePosePosition(bone_idx int64, position Vector3) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBonePosePosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) SetBonePoseRotation(bone_idx int64, rotation Quaternion) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), rotation.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBonePoseRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) SetBonePoseScale(bone_idx int64, scale Vector3) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), scale.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBonePoseScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetBonePosePosition(bone_idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBonePosePosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) GetBonePoseRotation(bone_idx int64) Quaternion {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewQuaternion()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBonePoseRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) GetBonePoseScale(bone_idx int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBonePoseScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) ResetBonePose(bone_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnResetBonePose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) ResetBonePoses() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnResetBonePoses), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) IsBoneEnabled(bone_idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnIsBoneEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skeleton3D) SetBoneEnabled(bone_idx int64, enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBoneEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetBoneGlobalPose(bone_idx int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBoneGlobalPose), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) SetBoneGlobalPose(bone_idx int64, pose Transform3D) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), pose.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBoneGlobalPose), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) ForceUpdateAllBoneTransforms() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnForceUpdateAllBoneTransforms), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) ForceUpdateBoneChildTransform(bone_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnForceUpdateBoneChildTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) SetMotionScale(motion_scale float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&motion_scale)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetMotionScale), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetMotionScale() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetMotionScale), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skeleton3D) SetShowRestOnly(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetShowRestOnly), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) IsShowRestOnly() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnIsShowRestOnly), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skeleton3D) SetModifierCallbackModeProcess(mode Skeleton3DModifierCallbackModeProcess) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetModifierCallbackModeProcess), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetModifierCallbackModeProcess() Skeleton3DModifierCallbackModeProcess {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	var ret Skeleton3DModifierCallbackModeProcess

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetModifierCallbackModeProcess), me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
	return ret
}

func (me *Skeleton3D) ClearBonesGlobalPoseOverride() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnClearBonesGlobalPoseOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) SetBoneGlobalPoseOverride(bone_idx int64, pose Transform3D, amount float64, persistent bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), pose.AsCTypePtr(), gdc.ConstTypePtr(&amount), gdc.ConstTypePtr(&persistent)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetBoneGlobalPoseOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetBoneGlobalPoseOverride(bone_idx int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBoneGlobalPoseOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) GetBoneGlobalPoseNoOverride(bone_idx int64) Transform3D {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()
	pinner.Pin(&bone_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetBoneGlobalPoseNoOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *Skeleton3D) SetAnimatePhysicalBones(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnSetAnimatePhysicalBones), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) GetAnimatePhysicalBones() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnGetAnimatePhysicalBones), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *Skeleton3D) PhysicalBonesStopSimulation() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnPhysicalBonesStopSimulation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) PhysicalBonesStartSimulation(bones []StringName) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bones)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnPhysicalBonesStartSimulation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) PhysicalBonesAddCollisionException(exception RID) {
	cargs := []gdc.ConstTypePtr{exception.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnPhysicalBonesAddCollisionException), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *Skeleton3D) PhysicalBonesRemoveCollisionException(exception RID) {
	cargs := []gdc.ConstTypePtr{exception.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton3D.fnPhysicalBonesRemoveCollisionException), me.obj, unsafe.SliceData(cargs), nil)

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

type Skeleton3DSkeletonUpdatedSignalFn func()

func (me *Skeleton3D) ConnectSkeletonUpdated(subs SignalSubscribers, fn Skeleton3DSkeletonUpdatedSignalFn) {
	sig := StringNameFromStr("skeleton_updated")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Skeleton3D) DisconnectSkeletonUpdated(subs SignalSubscribers, fn Skeleton3DSkeletonUpdatedSignalFn) {
	sig := StringNameFromStr("skeleton_updated")
	defer sig.Destroy()
	me.Disconnect(*sig, *subs.remove(fn))
}

type Skeleton3DBoneEnabledChangedSignalFn func(bone_idx int)

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

type Skeleton3DBoneListChangedSignalFn func()

func (me *Skeleton3D) ConnectBoneListChanged(subs SignalSubscribers, fn Skeleton3DBoneListChangedSignalFn) {
	sig := StringNameFromStr("bone_list_changed")
	defer sig.Destroy()
	me.Connect(*sig, subs.add(fn), 0)
}

func (me *Skeleton3D) DisconnectBoneListChanged(subs SignalSubscribers, fn Skeleton3DBoneListChangedSignalFn) {
	sig := StringNameFromStr("bone_list_changed")
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
