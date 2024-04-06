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

type ptrsForPhysicsRayQueryParameters2DList struct {
	fnCreate                     gdc.MethodBindPtr
	fnSetFrom                    gdc.MethodBindPtr
	fnGetFrom                    gdc.MethodBindPtr
	fnSetTo                      gdc.MethodBindPtr
	fnGetTo                      gdc.MethodBindPtr
	fnSetCollisionMask           gdc.MethodBindPtr
	fnGetCollisionMask           gdc.MethodBindPtr
	fnSetExclude                 gdc.MethodBindPtr
	fnGetExclude                 gdc.MethodBindPtr
	fnSetCollideWithBodies       gdc.MethodBindPtr
	fnIsCollideWithBodiesEnabled gdc.MethodBindPtr
	fnSetCollideWithAreas        gdc.MethodBindPtr
	fnIsCollideWithAreasEnabled  gdc.MethodBindPtr
	fnSetHitFromInside           gdc.MethodBindPtr
	fnIsHitFromInsideEnabled     gdc.MethodBindPtr
}

var ptrsForPhysicsRayQueryParameters2D ptrsForPhysicsRayQueryParameters2DList

func initPhysicsRayQueryParameters2DPtrs(iface gdc.Interface) {

	className := StringNameFromStr("PhysicsRayQueryParameters2D")
	defer className.Destroy()
	{
		methodName := StringNameFromStr("create")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnCreate = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3196569324))
	}
	{
		methodName := StringNameFromStr("set_from")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnSetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_from")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnGetFrom = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_to")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnSetTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 743155724))
	}
	{
		methodName := StringNameFromStr("get_to")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnGetTo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3341600327))
	}
	{
		methodName := StringNameFromStr("set_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnSetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
	}
	{
		methodName := StringNameFromStr("get_collision_mask")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnGetCollisionMask = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
	}
	{
		methodName := StringNameFromStr("set_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnSetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 381264803))
	}
	{
		methodName := StringNameFromStr("get_exclude")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnGetExclude = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3995934104))
	}
	{
		methodName := StringNameFromStr("set_collide_with_bodies")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnSetCollideWithBodies = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_bodies_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnIsCollideWithBodiesEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_collide_with_areas")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnSetCollideWithAreas = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_collide_with_areas_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnIsCollideWithAreasEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
	{
		methodName := StringNameFromStr("set_hit_from_inside")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnSetHitFromInside = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
	}
	{
		methodName := StringNameFromStr("is_hit_from_inside_enabled")
		defer methodName.Destroy()
		ptrsForPhysicsRayQueryParameters2D.fnIsHitFromInsideEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
	}
}

type PhysicsRayQueryParameters2D struct {
	RefCounted
}

func (me *PhysicsRayQueryParameters2D) BaseClass() string {
	return "PhysicsRayQueryParameters2D"
}

func NewPhysicsRayQueryParameters2D() *PhysicsRayQueryParameters2D {
	str := StringNameFromStr("PhysicsRayQueryParameters2D") // FIXME: should cache?
	defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
	obj := &PhysicsRayQueryParameters2D{}
	obj.SetBaseObject(objPtr)
	return obj
}

// Enums

func (me *PhysicsRayQueryParameters2D) Type() gdc.VariantType {
	return gdc.VariantTypeObject
}

func (me *PhysicsRayQueryParameters2D) AsTypePtr() gdc.TypePtr {
	return gdc.TypePtr(me.obj)
}

func (me *PhysicsRayQueryParameters2D) AsCTypePtr() gdc.ConstTypePtr {
	return gdc.ConstTypePtr(me.obj)
}

// Methods

func PhysicsRayQueryParameters2DCreate(from Vector2, to Vector2, collision_mask int64, exclude []RID) PhysicsRayQueryParameters2D {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), gdc.ConstTypePtr(&collision_mask), gdc.ConstTypePtr(&exclude)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewPhysicsRayQueryParameters2D()
	pinner.Pin(&collision_mask)
	pinner.Pin(&exclude)

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnCreate), nil, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsRayQueryParameters2D) SetFrom(from Vector2) {
	cargs := []gdc.ConstTypePtr{from.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnSetFrom), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters2D) GetFrom() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnGetFrom), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsRayQueryParameters2D) SetTo(to Vector2) {
	cargs := []gdc.ConstTypePtr{to.AsCTypePtr()}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnSetTo), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters2D) GetTo() Vector2 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewVector2()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnGetTo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return *ret
}

func (me *PhysicsRayQueryParameters2D) SetCollisionMask(collision_mask int64) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnSetCollisionMask), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters2D) GetCollisionMask() int64 {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewInt()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnGetCollisionMask), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsRayQueryParameters2D) SetExclude(exclude []RID) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&exclude)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnSetExclude), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters2D) GetExclude() []RID {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewArray()
	defer ret.Destroy()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnGetExclude), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	sliceRet, err := ConvertArrayToSlice[RID](ret)
	if err != nil {
		log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
		return nil
	}
	return sliceRet
}

func (me *PhysicsRayQueryParameters2D) SetCollideWithBodies(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnSetCollideWithBodies), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters2D) IsCollideWithBodiesEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnIsCollideWithBodiesEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsRayQueryParameters2D) SetCollideWithAreas(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnSetCollideWithAreas), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters2D) IsCollideWithAreasEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnIsCollideWithAreasEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

func (me *PhysicsRayQueryParameters2D) SetHitFromInside(enable bool) {
	cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable)}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnSetHitFromInside), me.obj, unsafe.SliceData(cargs), nil)

}

func (me *PhysicsRayQueryParameters2D) IsHitFromInsideEnabled() bool {
	cargs := []gdc.ConstTypePtr{}
	pinner := runtime.Pinner{}
	defer pinner.Unpin()
	ret := NewBool()

	giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsRayQueryParameters2D.fnIsHitFromInsideEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
	return ret.Get()
}

// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
