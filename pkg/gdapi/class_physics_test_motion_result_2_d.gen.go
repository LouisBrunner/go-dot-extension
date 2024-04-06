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

type ptrsForPhysicsTestMotionResult2DList struct {
	fnGetTravel                  gdc.MethodBindPtr
	fnGetRemainder               gdc.MethodBindPtr
	fnGetCollisionPoint          gdc.MethodBindPtr
	fnGetCollisionNormal         gdc.MethodBindPtr
	fnGetColliderVelocity        gdc.MethodBindPtr
	fnGetColliderId              gdc.MethodBindPtr
	fnGetColliderRid             gdc.MethodBindPtr
	fnGetCollider                gdc.MethodBindPtr
	fnGetColliderShape           gdc.MethodBindPtr
	fnGetCollisionLocalShape     gdc.MethodBindPtr
	fnGetCollisionDepth          gdc.MethodBindPtr
	fnGetCollisionSafeFraction   gdc.MethodBindPtr
	fnGetCollisionUnsafeFraction gdc.MethodBindPtr
}

var ptrsForPhysicsTestMotionResult2D ptrsForPhysicsTestMotionResult2DList

func initPhysicsTestMotionResult2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsTestMotionResult2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_travel")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetTravel = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_remainder")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetRemainder = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_collision_point")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetCollisionPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_collision_normal")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetCollisionNormal = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_collider_velocity")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetColliderVelocity = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("get_collider_id")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetColliderId = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_collider_rid")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetColliderRid = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
	}
	{
		methodName := StringNameFromStr("get_collider")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetCollider = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1981248198))
	}
	{
		methodName := StringNameFromStr("get_collider_shape")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetColliderShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_collision_local_shape")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetCollisionLocalShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("get_collision_depth")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetCollisionDepth = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_collision_safe_fraction")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetCollisionSafeFraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("get_collision_unsafe_fraction")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionResult2D.fnGetCollisionUnsafeFraction = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}

}

type PhysicsTestMotionResult2D struct {
	RefCounted
}

func (me *PhysicsTestMotionResult2D) BaseClass() string {
	return "PhysicsTestMotionResult2D"
}

func NewPhysicsTestMotionResult2D() *PhysicsTestMotionResult2D {
	str := StringNameFromStr("PhysicsTestMotionResult2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsTestMotionResult2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsTestMotionResult2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsTestMotionResult2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsTestMotionResult2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsTestMotionResult2D) GetTravel() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetTravel), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsTestMotionResult2D) GetRemainder() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetRemainder), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsTestMotionResult2D) GetCollisionPoint() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetCollisionPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsTestMotionResult2D) GetCollisionNormal() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetCollisionNormal), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsTestMotionResult2D) GetColliderVelocity() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetColliderVelocity), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsTestMotionResult2D) GetColliderId() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetColliderId), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsTestMotionResult2D) GetColliderRid() RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewRID()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetColliderRid), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsTestMotionResult2D) GetCollider() Object {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewObject()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetCollider), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsTestMotionResult2D) GetColliderShape() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetColliderShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsTestMotionResult2D) GetCollisionLocalShape() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetCollisionLocalShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsTestMotionResult2D) GetCollisionDepth() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetCollisionDepth), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsTestMotionResult2D) GetCollisionSafeFraction() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetCollisionSafeFraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsTestMotionResult2D) GetCollisionUnsafeFraction() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionResult2D.fnGetCollisionUnsafeFraction), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Signals
