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

type ptrsForSkeletonModification2DJiggleList struct {
	fnSetTargetNode            gdc.MethodBindPtr
	fnGetTargetNode            gdc.MethodBindPtr
	fnSetJiggleDataChainLength gdc.MethodBindPtr
	fnGetJiggleDataChainLength gdc.MethodBindPtr
	fnSetStiffness             gdc.MethodBindPtr
	fnGetStiffness             gdc.MethodBindPtr
	fnSetMass                  gdc.MethodBindPtr
	fnGetMass                  gdc.MethodBindPtr
	fnSetDamping               gdc.MethodBindPtr
	fnGetDamping               gdc.MethodBindPtr
	fnSetUseGravity            gdc.MethodBindPtr
	fnGetUseGravity            gdc.MethodBindPtr
	fnSetGravity               gdc.MethodBindPtr
	fnGetGravity               gdc.MethodBindPtr
	fnSetUseColliders          gdc.MethodBindPtr
	fnGetUseColliders          gdc.MethodBindPtr
	fnSetCollisionMask         gdc.MethodBindPtr
	fnGetCollisionMask         gdc.MethodBindPtr
	fnSetJiggleJointBone2DNode gdc.MethodBindPtr
	fnGetJiggleJointBone2DNode gdc.MethodBindPtr
	fnSetJiggleJointBoneIndex  gdc.MethodBindPtr
	fnGetJiggleJointBoneIndex  gdc.MethodBindPtr
	fnSetJiggleJointOverride   gdc.MethodBindPtr
	fnGetJiggleJointOverride   gdc.MethodBindPtr
	fnSetJiggleJointStiffness  gdc.MethodBindPtr
	fnGetJiggleJointStiffness  gdc.MethodBindPtr
	fnSetJiggleJointMass       gdc.MethodBindPtr
	fnGetJiggleJointMass       gdc.MethodBindPtr
	fnSetJiggleJointDamping    gdc.MethodBindPtr
	fnGetJiggleJointDamping    gdc.MethodBindPtr
	fnSetJiggleJointUseGravity gdc.MethodBindPtr
	fnGetJiggleJointUseGravity gdc.MethodBindPtr
	fnSetJiggleJointGravity    gdc.MethodBindPtr
	fnGetJiggleJointGravity    gdc.MethodBindPtr
}

var ptrsForSkeletonModification2DJiggle ptrsForSkeletonModification2DJiggleList

func initSkeletonModification2DJigglePtrs(iface gdc.Interface) {

	className := StringNameFromStr("SkeletonModification2DJiggle")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_target_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_target_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
	}
	{
		methodName := StringNameFromStr("set_jiggle_data_chain_length")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetJiggleDataChainLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_jiggle_data_chain_length")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetJiggleDataChainLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
	}
	{
		methodName := StringNameFromStr("set_stiffness")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_stiffness")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_mass")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_mass")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_damping")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetDamping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_damping")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetDamping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_use_gravity")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetUseGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_use_gravity")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetUseGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_gravity")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_gravity")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_use_colliders")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetUseColliders = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_use_colliders")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetUseColliders = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_jiggle_joint_bone2d_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetJiggleJointBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761262315))
	}
	{
		methodName := StringNameFromStr("get_jiggle_joint_bone2d_node")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetJiggleJointBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
	}
	{
		methodName := StringNameFromStr("set_jiggle_joint_bone_index")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetJiggleJointBoneIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
	}
	{
		methodName := StringNameFromStr("get_jiggle_joint_bone_index")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetJiggleJointBoneIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
	}
	{
		methodName := StringNameFromStr("set_jiggle_joint_override")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetJiggleJointOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_jiggle_joint_override")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetJiggleJointOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_jiggle_joint_stiffness")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetJiggleJointStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_jiggle_joint_stiffness")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetJiggleJointStiffness = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_jiggle_joint_mass")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetJiggleJointMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_jiggle_joint_mass")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetJiggleJointMass = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_jiggle_joint_damping")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetJiggleJointDamping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
	}
	{
		methodName := StringNameFromStr("get_jiggle_joint_damping")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetJiggleJointDamping = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
	}
	{
		methodName := StringNameFromStr("set_jiggle_joint_use_gravity")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetJiggleJointUseGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
	}
	{
		methodName := StringNameFromStr("get_jiggle_joint_use_gravity")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetJiggleJointUseGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
	}
	{
		methodName := StringNameFromStr("set_jiggle_joint_gravity")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnSetJiggleJointGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
	}
	{
		methodName := StringNameFromStr("get_jiggle_joint_gravity")
		defer methodName.Destroy()
		ptrsForSkeletonModification2DJiggle.fnGetJiggleJointGravity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
	}
}

type SkeletonModification2DJiggle struct {
	SkeletonModification2D
}

func (me *SkeletonModification2DJiggle) BaseClass() string {
	return "SkeletonModification2DJiggle"
}

func NewSkeletonModification2DJiggle() *SkeletonModification2DJiggle {
	str := StringNameFromStr("SkeletonModification2DJiggle") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SkeletonModification2DJiggle{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SkeletonModification2DJiggle) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SkeletonModification2DJiggle) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DJiggle) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SkeletonModification2DJiggle) SetTargetNode(target_nodepath NodePath) {
	cargs := []gdc.ConstTypePtr{target_nodepath.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetTargetNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetTargetNode() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetTargetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonModification2DJiggle) SetJiggleDataChainLength(length int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetJiggleDataChainLength), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetJiggleDataChainLength() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetJiggleDataChainLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetStiffness(stiffness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stiffness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetStiffness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetStiffness() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetStiffness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetMass(mass float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetMass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetMass() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetDamping(damping float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&damping)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetDamping), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetDamping() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetDamping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetUseGravity(use_gravity bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_gravity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetUseGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetUseGravity() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetUseGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetGravity(gravity Vector2) {
	cargs := []gdc.ConstTypePtr{gravity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetGravity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonModification2DJiggle) SetUseColliders(use_colliders bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_colliders)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetUseColliders), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetUseColliders() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetUseColliders), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetCollisionMask(collision_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetJiggleJointBone2DNode(joint_idx int64, bone2d_node NodePath) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), bone2d_node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetJiggleJointBone2DNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetJiggleJointBone2DNode(joint_idx int64) NodePath {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()
	pinner.Pin(&joint_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetJiggleJointBone2DNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonModification2DJiggle) SetJiggleJointBoneIndex(joint_idx int64, bone_idx int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&bone_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetJiggleJointBoneIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetJiggleJointBoneIndex(joint_idx int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&joint_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetJiggleJointBoneIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetJiggleJointOverride(joint_idx int64, override bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&override)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetJiggleJointOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetJiggleJointOverride(joint_idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&joint_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetJiggleJointOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetJiggleJointStiffness(joint_idx int64, stiffness float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&stiffness)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetJiggleJointStiffness), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetJiggleJointStiffness(joint_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&joint_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetJiggleJointStiffness), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetJiggleJointMass(joint_idx int64, mass float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&mass)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetJiggleJointMass), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetJiggleJointMass(joint_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&joint_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetJiggleJointMass), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetJiggleJointDamping(joint_idx int64, damping float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&damping)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetJiggleJointDamping), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetJiggleJointDamping(joint_idx int64) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&joint_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetJiggleJointDamping), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetJiggleJointUseGravity(joint_idx int64, use_gravity bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&use_gravity)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetJiggleJointUseGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetJiggleJointUseGravity(joint_idx int64) bool {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()
	pinner.Pin(&joint_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetJiggleJointUseGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonModification2DJiggle) SetJiggleJointGravity(joint_idx int64, gravity Vector2) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gravity.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnSetJiggleJointGravity), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonModification2DJiggle) GetJiggleJointGravity(joint_idx int64) Vector2 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()
	pinner.Pin(&joint_idx)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DJiggle.fnGetJiggleJointGravity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
