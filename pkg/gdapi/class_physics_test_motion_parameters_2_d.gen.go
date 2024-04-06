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

type ptrsForPhysicsTestMotionParameters2DList struct {
	fnGetFrom                        gdc.MethodBindPtr
	fnSetFrom                        gdc.MethodBindPtr
	fnGetMotion                      gdc.MethodBindPtr
	fnSetMotion                      gdc.MethodBindPtr
	fnGetMargin                      gdc.MethodBindPtr
	fnSetMargin                      gdc.MethodBindPtr
	fnIsCollideSeparationRayEnabled  gdc.MethodBindPtr
	fnSetCollideSeparationRayEnabled gdc.MethodBindPtr
	fnGetExcludeBodies               gdc.MethodBindPtr
	fnSetExcludeBodies               gdc.MethodBindPtr
	fnGetExcludeObjects              gdc.MethodBindPtr
	fnSetExcludeObjects              gdc.MethodBindPtr
	fnIsRecoveryAsCollisionEnabled   gdc.MethodBindPtr
	fnSetRecoveryAsCollisionEnabled  gdc.MethodBindPtr
}

var ptrsForPhysicsTestMotionParameters2D ptrsForPhysicsTestMotionParameters2DList

func initPhysicsTestMotionParameters2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsTestMotionParameters2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("get_from")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnGetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3814499831))
	}
	{
		methodName := StringNameFromStr("set_from")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnSetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761652528))
	}
	{
		methodName := StringNameFromStr("get_motion")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnGetMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_motion")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnSetMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_margin")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnGetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
	}
	{
		methodName := StringNameFromStr("set_margin")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnSetMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
	}
	{
		methodName := StringNameFromStr("is_collide_separation_ray_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnIsCollideSeparationRayEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collide_separation_ray_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnSetCollideSeparationRayEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("get_exclude_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnGetExcludeBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_exclude_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnSetExcludeBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_exclude_objects")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnGetExcludeObjects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_exclude_objects")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnSetExcludeObjects = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("is_recovery_as_collision_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnIsRecoveryAsCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_recovery_as_collision_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsTestMotionParameters2D.fnSetRecoveryAsCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}

}

type PhysicsTestMotionParameters2D struct {
	RefCounted
}

func (me *PhysicsTestMotionParameters2D) BaseClass() string {
	return "PhysicsTestMotionParameters2D"
}

func NewPhysicsTestMotionParameters2D() *PhysicsTestMotionParameters2D {
	str := StringNameFromStr("PhysicsTestMotionParameters2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsTestMotionParameters2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsTestMotionParameters2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsTestMotionParameters2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsTestMotionParameters2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func (me *PhysicsTestMotionParameters2D) GetFrom() Transform2D {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewTransform2D()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnGetFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsTestMotionParameters2D) SetFrom(from Transform2D) {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnSetFrom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsTestMotionParameters2D) GetMotion() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnGetMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsTestMotionParameters2D) SetMotion(motion Vector2) {
	cargs := []gdc.ConstTypePtr{motion.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnSetMotion), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsTestMotionParameters2D) GetMargin() float64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewFloat()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnGetMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsTestMotionParameters2D) SetMargin(margin float64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnSetMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsTestMotionParameters2D) IsCollideSeparationRayEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnIsCollideSeparationRayEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsTestMotionParameters2D) SetCollideSeparationRayEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnSetCollideSeparationRayEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsTestMotionParameters2D) GetExcludeBodies() []RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnGetExcludeBodies), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsTestMotionParameters2D) SetExcludeBodies(exclude_list []RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_list)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnSetExcludeBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsTestMotionParameters2D) GetExcludeObjects() []int {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnGetExcludeObjects), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[int](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsTestMotionParameters2D) SetExcludeObjects(exclude_list []int) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude_list)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnSetExcludeObjects), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsTestMotionParameters2D) IsRecoveryAsCollisionEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnIsRecoveryAsCollisionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsTestMotionParameters2D) SetRecoveryAsCollisionEnabled(enabled bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsTestMotionParameters2D.fnSetRecoveryAsCollisionEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
