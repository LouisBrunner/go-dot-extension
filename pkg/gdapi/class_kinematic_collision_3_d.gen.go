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

type ptrsForKinematicCollision3DList struct {
	fnGetTravel             gdc.MethodBindPtr
	fnGetRemainder          gdc.MethodBindPtr
	fnGetDepth              gdc.MethodBindPtr
	fnGetCollisionCount     gdc.MethodBindPtr
	fnGetPosition           gdc.MethodBindPtr
	fnGetNormal             gdc.MethodBindPtr
	fnGetAngle              gdc.MethodBindPtr
	fnGetLocalShape         gdc.MethodBindPtr
	fnGetCollider           gdc.MethodBindPtr
	fnGetColliderId         gdc.MethodBindPtr
	fnGetColliderRid        gdc.MethodBindPtr
	fnGetColliderShape      gdc.MethodBindPtr
	fnGetColliderShapeIndex gdc.MethodBindPtr
	fnGetColliderVelocity   gdc.MethodBindPtr
}

var ptrsForKinematicCollision3D ptrsForKinematicCollision3DList

func initKinematicCollision3DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("KinematicCollision3D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_travel")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetTravel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_remainder")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetRemainder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3360562783))
	}
	{
		methodName := StringNameFromStr("get_depth")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_collision_count")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetCollisionCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_position")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1914908202))
	}
	{
		methodName := StringNameFromStr("get_normal")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1914908202))
	}
	{
		methodName := StringNameFromStr("get_angle")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1242741860))
	}
	{
		methodName := StringNameFromStr("get_local_shape")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetLocalShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2639523548))
	}
	{
		methodName := StringNameFromStr("get_collider")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2639523548))
	}
	{
		methodName := StringNameFromStr("get_collider_id")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetColliderId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("get_collider_rid")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetColliderRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1231817359))
	}
	{
		methodName := StringNameFromStr("get_collider_shape")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2639523548))
	}
	{
		methodName := StringNameFromStr("get_collider_shape_index")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetColliderShapeIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1591665591))
	}
	{
		methodName := StringNameFromStr("get_collider_velocity")
		defer methodName.Destroy()
		ptrsForKinematicCollision3D.fnGetColliderVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1914908202))
	}
}

type KinematicCollision3D struct {
	RefCounted
}

func (me *KinematicCollision3D) BaseClass() string {
	return "KinematicCollision3D"
}

func NewKinematicCollision3D() *KinematicCollision3D {
	str := StringNameFromStr("KinematicCollision3D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &KinematicCollision3D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *KinematicCollision3D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *KinematicCollision3D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *KinematicCollision3D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *KinematicCollision3D) GetTravel() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetTravel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision3D) GetRemainder() Vector3 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetRemainder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision3D) GetDepth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *KinematicCollision3D) GetCollisionCount() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetCollisionCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *KinematicCollision3D) GetPosition(collision_index int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision3D) GetNormal(collision_index int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision3D) GetAngle(collision_index int64, up_direction Vector3) float64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index), up_direction.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *KinematicCollision3D) GetLocalShape(collision_index int64) Object {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetLocalShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision3D) GetCollider(collision_index int64) Object {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision3D) GetColliderId(collision_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetColliderId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *KinematicCollision3D) GetColliderRid(collision_index int64) RID {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetColliderRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision3D) GetColliderShape(collision_index int64) Object {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *KinematicCollision3D) GetColliderShapeIndex(collision_index int64) int64 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetColliderShapeIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *KinematicCollision3D) GetColliderVelocity(collision_index int64) Vector3 {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_index)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector3()
	pinner.Pin(&collision_index)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForKinematicCollision3D.fnGetColliderVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

// Signals
