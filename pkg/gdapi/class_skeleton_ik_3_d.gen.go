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

type ptrsForSkeletonIK3DList struct {
	fnSetRootBone         gdc.MethodBindPtr
	fnGetRootBone         gdc.MethodBindPtr
	fnSetTipBone          gdc.MethodBindPtr
	fnGetTipBone          gdc.MethodBindPtr
	fnSetInterpolation    gdc.MethodBindPtr
	fnGetInterpolation    gdc.MethodBindPtr
	fnSetTargetTransform  gdc.MethodBindPtr
	fnGetTargetTransform  gdc.MethodBindPtr
	fnSetTargetNode       gdc.MethodBindPtr
	fnGetTargetNode       gdc.MethodBindPtr
	fnSetOverrideTipBasis gdc.MethodBindPtr
	fnIsOverrideTipBasis  gdc.MethodBindPtr
	fnSetUseMagnet        gdc.MethodBindPtr
	fnIsUsingMagnet       gdc.MethodBindPtr
	fnSetMagnetPosition   gdc.MethodBindPtr
	fnGetMagnetPosition   gdc.MethodBindPtr
	fnGetParentSkeleton   gdc.MethodBindPtr
	fnIsRunning           gdc.MethodBindPtr
	fnSetMinDistance      gdc.MethodBindPtr
	fnGetMinDistance      gdc.MethodBindPtr
	fnSetMaxIterations    gdc.MethodBindPtr
	fnGetMaxIterations    gdc.MethodBindPtr
	fnStart               gdc.MethodBindPtr
	fnStop                gdc.MethodBindPtr
}

var ptrsForSkeletonIK3D ptrsForSkeletonIK3DList

func initSkeletonIK3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("SkeletonIK3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("set_root_bone")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetRootBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_root_bone")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnGetRootBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_tip_bone")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetTipBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3304788590))
	}
	{
		methodName := StringNameFromStr("get_tip_bone")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnGetTipBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2002593661))
	}
	{
		methodName := StringNameFromStr("set_interpolation")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetInterpolation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_interpolation")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnGetInterpolation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_target_transform")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetTargetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2952846383))
	}
	{
		methodName := StringNameFromStr("get_target_transform")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnGetTargetTransform = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3229777777))
	}
	{
		methodName := StringNameFromStr("set_target_node")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
	}
	{
		methodName := StringNameFromStr("get_target_node")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnGetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 277076166))
	}
	{
		methodName := StringNameFromStr("set_override_tip_basis")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetOverrideTipBasis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_override_tip_basis")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnIsOverrideTipBasis = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_use_magnet")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetUseMagnet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_using_magnet")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnIsUsingMagnet = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_magnet_position")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetMagnetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3460891852))
	}
	{
		methodName := StringNameFromStr("get_magnet_position")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnGetMagnetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_parent_skeleton")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnGetParentSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1488626673))
	}
	{
		methodName := StringNameFromStr("is_running")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnIsRunning = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
	}
	{
		methodName := StringNameFromStr("set_min_distance")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetMinDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("get_min_distance")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnGetMinDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_max_iterations")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnSetMaxIterations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_max_iterations")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnGetMaxIterations = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("start")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnStart = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 107499316))
	}
	{
		methodName := StringNameFromStr("stop")
		defer methodName.Destroy()
		ptrsForSkeletonIK3D.fnStop = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3218959716))
	}

}

type SkeletonIK3D struct {
	Node
}

func (me *SkeletonIK3D) BaseClass() string {
	return "SkeletonIK3D"
}

func NewSkeletonIK3D() *SkeletonIK3D {
	str := StringNameFromStr("SkeletonIK3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &SkeletonIK3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *SkeletonIK3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *SkeletonIK3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *SkeletonIK3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *SkeletonIK3D) SetRootBone(root_bone StringName) {
	cargs := []gdc.ConstTypePtr{root_bone.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetRootBone), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) GetRootBone() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnGetRootBone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonIK3D) SetTipBone(tip_bone StringName) {
	cargs := []gdc.ConstTypePtr{tip_bone.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetTipBone), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) GetTipBone() StringName {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewStringName()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnGetTipBone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonIK3D) SetInterpolation(interpolation float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&interpolation)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetInterpolation), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) GetInterpolation() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnGetInterpolation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonIK3D) SetTargetTransform(target Transform3D) {
	cargs := []gdc.ConstTypePtr{target.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetTargetTransform), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) GetTargetTransform() Transform3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnGetTargetTransform), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonIK3D) SetTargetNode(node NodePath) {
	cargs := []gdc.ConstTypePtr{node.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetTargetNode), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) GetTargetNode() NodePath {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewNodePath()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnGetTargetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonIK3D) SetOverrideTipBasis(override bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&override)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetOverrideTipBasis), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) IsOverrideTipBasis() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnIsOverrideTipBasis), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonIK3D) SetUseMagnet(use bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetUseMagnet), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) IsUsingMagnet() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnIsUsingMagnet), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonIK3D) SetMagnetPosition(local_position Vector3) {
	cargs := []gdc.ConstTypePtr{local_position.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetMagnetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) GetMagnetPosition() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnGetMagnetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonIK3D) GetParentSkeleton() Skeleton3D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewSkeleton3D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnGetParentSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *SkeletonIK3D) IsRunning() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnIsRunning), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonIK3D) SetMinDistance(min_distance float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&min_distance)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetMinDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) GetMinDistance() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnGetMinDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonIK3D) SetMaxIterations(iterations int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&iterations)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnSetMaxIterations), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) GetMaxIterations() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnGetMaxIterations), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *SkeletonIK3D) Start(one_time bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&one_time)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnStart), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *SkeletonIK3D) Stop() {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonIK3D.fnStop), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
